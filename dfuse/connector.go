package dfuse

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pinax-network/golang-base/log"
	"go.uber.org/zap"
	"sync"
)

type Connector struct {
	identifier      string
	dfuseClient     *DfuseClient
	actionHandler   map[string]ActionHandler
	docHandler      DocumentHandler
	shutdownHandler ShutdownHandler
	hadError        bool
}

func NewConnector(identifier string, dfuseClient *DfuseClient, actionHandler map[string]ActionHandler, docHandler DocumentHandler, shutdownHandler ShutdownHandler) *Connector {
	return &Connector{
		identifier, dfuseClient, actionHandler, docHandler, shutdownHandler, false,
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
				c.hadError = true
				continue
			}

			document := &SearchTransactionsForwardResponse{}
			err = json.Unmarshal([]byte(resp.Data), document)
			if err != nil {
				log.Error("failed to unmarshal document", zap.Error(err))
				c.hadError = true
				continue
			}

			result := document.SearchTransactionsForwardDoc
			blockNum := result.Trace.Block.Num

			// report the read block here
			reportLastHeadBlockTime(c.identifier, result.Trace.Block.Timestamp)
			reportLastHeadBlockNumber(c.identifier, blockNum)

			trxId, err := hex.DecodeString(result.Trace.ID)
			if err != nil {
				log.Error("failed to decode transaction id", zap.Error(err))
				c.hadError = true
				continue
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
				if err != nil {
					log.Error("failed to execute action handler", zap.Error(err), zap.String("action", action.Name), zap.Any("action", action))
					c.hadError = true
					continue
				}
			}

			if c.docHandler != nil {
				err = c.docHandler.HandleDocument(ctx, document)
				if err != nil {
					log.Error("failed to execute document handler", zap.Error(err))
					c.hadError = true
					continue
				}
			}

			// only report the last successful block time if no error occurred since start time
			if !c.hadError {
				reportLastSuccessfulBlockTime(c.identifier, result.Trace.Block.Timestamp)
				reportLastSuccessfulBlockNumber(c.identifier, blockNum)
			}
		}
	}
}

func (c *Connector) HadErrors() bool {
	return c.hadError
}
