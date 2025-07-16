package domain

type RateRequest struct {
	Market string
}

type RateResponse struct {
	Market    string
	Ask       AskBidData
	Bid       AskBidData
	TimeStamp uint64
}

type AskBidData struct {
	Price   string `json:"price"`
	Volume  string `json:"volume"`
	Amount  string `json:"amount"`
	Factor  string `json:"factor"`
	TypeAsk string `json:"type"`
}

type GrinexResponse struct {
	TimeStamp uint64       `json:"timestamp"`
	Asks      []AskBidData `json:"asks"`
	Bids      []AskBidData `json:"bids"`
}
