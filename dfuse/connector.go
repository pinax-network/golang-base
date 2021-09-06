package dfuse

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/eosnationftw/eosn-base-api/log"
	"go.uber.org/zap"
	"sync"
)

type Connector struct {
	identifier      string
	dfuseClient     *DfuseClient
	actionHandler   map[string]ActionHandler
	docHandler      DocumentHandler
	shutdownHandler ShutdownHandler
}

func NewConnector(identifier string, dfuseClient *DfuseClient, actionHandler map[string]ActionHandler, docHandler DocumentHandler, shutdownHandler ShutdownHandler) *Connector {
	return &Connector{
		identifier, dfuseClient, actionHandler, docHandler, shutdownHandler,
	}
}

func (c *Connector) Run(wg *sync.WaitGroup, ctx context.Context, query string) error {
	defer wg.Done()

	executor, err := c.dfuseClient.GraphQLClient.Execute(ctx, &Request{Query: query})
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			// todo shutdown routine if necessary
			if c.shutdownHandler != nil {
				c.shutdownHandler.Shutdown(ctx)
			}
			log.Info("shut down accounts dfuse conector")
			return nil
		default:
			resp, err := executor.Recv()
			// todo this is not an error if context is canceled due to shutdown, handle this case
			if err != nil {
				return fmt.Errorf("failed on Recv(): %w", err)
			}

			if len(resp.Errors) > 0 {
				log.Error("Request failed", zap.Any("errors", resp.Errors))
				return fmt.Errorf("request failed: %v", resp.Errors)
			}

			document := &SearchTransactionsForwardResponse{}
			err = json.Unmarshal([]byte(resp.Data), document)
			if err != nil {
				return fmt.Errorf("failed to unmarshal: %w", err)
			}

			log.Debug("received new dfuse result", zap.Any("document", document))

			result := document.SearchTransactionsForwardDoc
			blockNum := result.Trace.Block.Num
			trxId, err := hex.DecodeString(result.Trace.ID)
			if err != nil {
				return fmt.Errorf("failed to decode transaction id: %w", err)
			}

			for _, action := range result.Trace.MatchingActions {
				actionHandler, ok := c.actionHandler[action.Name]
				if !ok {
					// check if we have a general action handler instead
					if actionHandler, ok = c.actionHandler["*"]; !ok {
						log.Debug("no action handler available for action", zap.String("action", action.Name))
						continue
					}
				}
				err = actionHandler.HandleAction(ctx, result.Undo, trxId, &result.Trace.Block, &action)
				log.CriticalIfError("failed to execute action handler", err, zap.String("action", action.Name), zap.Any("action", action))
			}

			if c.docHandler != nil {
				err = c.docHandler.HandleDocument(ctx, document)
				log.CriticalIfError("failed to execute document handler", err)
			}

			reportLastHeadBlockTime(c.identifier, result.Trace.Block.Timestamp)
			reportLastHeadBlockNumber(c.identifier, blockNum)
		}
	}
}
