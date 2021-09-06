package dfuse

import (
	"context"
)

type ActionHandler interface {
	HandleAction(ctx context.Context, undo bool, trxId []byte, block *BlockResponse, action *ActionResponse) error
}

type DocumentHandler interface {
	HandleDocument(ctx context.Context, doc *SearchTransactionsForwardResponse) error
}

type ShutdownHandler interface {
	Shutdown(ctx context.Context)
}
