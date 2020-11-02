package api

import "github.com/alpacahq/marketstore/v4/contrib/alpaca/enums"

type Message struct {
	EventType enums.Prefix `json:"ev"`
}

type AlpacaMessage struct {
	Data Message `json:"data"`
}

// Trade is a trade
type Trade struct {
	// event name, always “T”
	eventType string `json:"-"` // ev
	// symbol
	Symbol string `json:"T"`
	// trade ID
	tradeID int `json:"-"` // i
	// exchange code where the trade occurred
	exchange int `json:"-"` // x
	// trade price
	Price float64 `json:"p"`
	// trade size (shares)
	Size int64 `json:"s"`
	// epoch timestamp in nanoseconds
	Timestamp int64 `json:"t"`
	// condition flags
	Conditions []int `json:"c"`
	// tape ID
	tapeID int `json:"-"` // z
}

// AlpacaTrade is the message
// from Alpaca that contains the trade
type AlpacaTrade struct {
	Data Trade `json:"data"`
}

// Quote is a quote
type Quote struct {
	// event name, always “Q”
	eventType string `json:"-"` // ev
	// symbol
	Symbol string `json:"T"`
	// exchange code for bid quote
	bidExchange int `json:"-"` // x
	// bid price
	BidPrice float64 `json:"p"`
	// bid size
	BidSize int64 `json:"s"`
	// exchange code for ask quote
	askExchange int `json:"-"` // X
	// ask price
	AskPrice float64 `json:"P"`
	// ask size
	AskSize int64 `json:"S"`
	// condition flags
	condition int `json:"-"` // c
	// epoch timestamp in nanoseconds
	Timestamp int64 `json:"t"`
}

// AlpacaQuote is the message
// from Alpaca that contains the quote
type AlpacaQuote struct {
	Data Quote `json:"data"`
}

// Aggregate is an aggregate
type Aggregate struct {
	// event name, always “AM”
	eventType string `json:"-"` // ev
	// symbol
	Symbol string `json:"T"`
	// volume (shares)
	Volume int `json:"v"`
	// accumulated volume (shares)
	accumVolume int `json:"-"` // av
	//official open price of the bar
	officialOpen float64 `json:"-"` // op
	// VWAP (Volume Weighted Average Price)
	vWAP float64 `json:"-"` // vw
	// open price of the bar
	Open float64 `json:"o"`
	// close price of the bar
	Close float64 `json:"c"`
	// high price of the bar
	High float64 `json:"h"`
	// low price of the bar
	Low float64 `json:"l"`
	// average price of the bar
	average float64 `json:"-"` // a
	// epoch time at the beginning of the window in milliseconds
	EpochMillis int64 `json:"s"`
	// epoch time at the ending of the window in milliseconds
	endTime int64 `json:"-"` // e
}

// AlpacaAggregate is the message
// from Alpaca that contains the aggregate
type AlpacaAggregate struct {
	Data Aggregate `json:"data"`
}
