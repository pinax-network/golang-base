package dfuse

import "time"

type SearchTransactionsForwardResponse struct {
	SearchTransactionsForwardDoc `json:"searchTransactionsForward"`
}

type SearchTransactionsForwardDoc struct {
	Cursor string `json:"cursor"`
	Undo   bool   `json:"undo"`
	Trace  struct {
		Block           BlockResponse    `json:"block"`
		ID              string           `json:"id"`
		MatchingActions []ActionResponse `json:"matchingActions"`
	} `json:"trace"`
}

type BlockResponse struct {
	Num       int       `json:"num"`
	Id        string    `json:"id"`
	Confirmed int       `json:"confirmed"`
	Timestamp time.Time `json:"timestamp"`
	Previous  string    `json:"previous"`
}

type ActionResponse struct {
	Account string      `json:"block"`
	Name    string      `json:"name"`
	Json    interface{} `json:"json"`
	Seq     string      `json:"seq"`
}
