package models

type Order struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type Board struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []Order `json:"bids"`
	Asks     []Order `json:"asks"`
}

type Ticket struct {
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         int     `json:"best_bid"`
	BestAsk         int     `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     int     `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   int     `json:"total_ask_depth"`
	MarketBidSize   int     `json:"market_bid_size"`
	MarketAskSize   int     `json:"market_ask_size"`
	Ltp             int     `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

type Execution struct {
	ID                         int     `json:"id"`
	Side                       string  `json:"side"`
	Price                      int     `json:"price"`
	Size                       float64 `json:"size"`
	ExecDate                   string  `json:"exec_date"`
	BuyChildOrderAcceptanceID  string  `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceID string  `json:"sell_child_order_acceptance_id"`
}

type BoardStatus struct {
	Health string `json:"health"`
	State  string `json:"state"`
	Data   struct {
		SpecialQuotation int `json:"special_quotation"`
	} `json:"data"`
}

type Health struct {
	Status string `json:"status"`
}
