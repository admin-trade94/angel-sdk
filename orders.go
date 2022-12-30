package smartapigo

import (
	"net/http"
)

// Order represents a individual order response.
type Order struct {
	Variety                 string  `json:"variety"`
	OrderType               string  `json:"ordertype"`
	ProductType             string  `json:"producttype"`
	Duration                string  `json:"duration"`
	Price                   float64 `json:"price"`
	TriggerPrice            float64 `json:"triggerprice"`
	Quantity                string  `json:"quantity"`
	DisclosedQuantity       string  `json:"disclosedquantity"`
	SquareOff               float64 `json:"squareoff"`
	StopLoss                float64 `json:"stoploss"`
	TrailingStopLoss        float64 `json:"trailingstoploss"`
	TrailingSymbol          string  `json:"trailingsymbol"`
	TransactionType         string  `json:"transactiontype"`
	Exchange                string  `json:"exchange"`
	SymbolToken             string  `json:"symboltoken"`
	InstrumentType          string  `json:"instrumenttype"`
	StrikePrice             float64 `json:"strikeprice"`
	OptionType              string  `json:"optiontype"`
	ExpiryDate              string  `json:"expirydate"`
	LotSize                 string  `json:"lotsize"`
	CancelSize              string  `json:"cancelsize"`
	AveragePrice            float64 `json:"averageprice"`
	FilledShares            string  `json:"filledshares"`
	UnfilledShares          string  `json:"unfilledshares"`
	OrderID                 string  `json:"orderid"`
	Text                    string  `json:"text"`
	Status                  string  `json:"status"`
	OrderStatus             string  `json:"orderstatus"`
	UpdateTime              string  `json:"updatetime"`
	ExchangeTime            string  `json:"exchtime"`
	ExchangeOrderUpdateTime string  `json:"exchorderupdatetime"`
	FillID                  string  `json:"fillid"`
	FillTime                string  `json:"filltime"`
}

// Orders is a list of orders.
type Orders []Order

// OrderParams represents parameters for placing an order.
type OrderParams struct {
	Variety         string `json:"variety"`
	TradingSymbol   string `json:"tradingsymbol"`
	SymbolToken     string `json:"symboltoken"`
	TransactionType string `json:"transactiontype"`
	Exchange        string `json:"exchange"`
	OrderType       string `json:"ordertype"`
	ProductType     string `json:"producttype"`
	Duration        string `json:"duration"`
	Price           string `json:"price"`
	SquareOff       string `json:"squareoff"`
	StopLoss        string `json:"stoploss"`
	Quantity        string `json:"quantity"`
}

// OrderParams represents parameters for modifying an order.
type ModifyOrderParams struct {
	Variety       string `json:"variety"`
	OrderID       string `json:"orderid"`
	OrderType     string `json:"ordertype"`
	ProductType   string `json:"producttype"`
	Duration      string `json:"duration"`
	Price         string `json:"price"`
	Quantity      string `json:"quantity"`
	TradingSymbol string `json:"tradingsymbol"`
	SymbolToken   string `json:"symboltoken"`
	Exchange      string `json:"exchange"`
}

// OrderResponse represents the order place success response.
type OrderResponse struct {
	Script  string `json:"script"`
	OrderID string `json:"orderid"`
}

// Trade represents an individual trade response.
type Trade struct {
	Exchange        string `json:"exchange"`
	ProductType     string `json:"producttype"`
	TradingSymbol   string `json:"tradingsymbol"`
	InstrumentType  string `json:"instrumenttype"`
	SymbolGroup     string `json:"symbolgroup"`
	StrikePrice     string `json:"strikeprice"`
	OptionType      string `json:"optiontype"`
	ExpiryDate      string `json:"expirydate"`
	MarketLot       string `json:"marketlot"`
	Precision       string `json:"precision"`
	Multiplier      string `json:"multiplier"`
	TradeValue      string `json:"tradevalue"`
	TransactionType string `json:"transactiontype"`
	FillPrice       string `json:"fillprice"`
	FillSize        string `json:"fillsize"`
	OrderID         string `json:"orderid"`
	FillID          string `json:"fillid"`
	FillTime        string `json:"filltime"`
}

// Trades is a list of trades.
type Trades []Trade

// Position represents an individual position response.
type Position struct {
	Exchange              string `json:"exchange"`
	SymbolToken           string `json:"symboltoken"`
	ProductType           string `json:"producttype"`
	Tradingsymbol         string `json:"tradingsymbol"`
	SymbolName            string `json:"symbolname"`
	InstrumentType        string `json:"instrumenttype"`
	PriceDen              string `json:"priceden"`
	PriceNum              string `json:"pricenum"`
	GenDen                string `json:"genden"`
	GenNum                string `json:"gennum"`
	Precision             string `json:"precision"`
	Multiplier            string `json:"multiplier"`
	BoardLotSize          string `json:"boardlotsize"`
	BuyQuantity           string `json:"buyquantity"`
	SellQuantity          string `json:"sellquantity"`
	BuyAmount             string `json:"buyamount"`
	SellAmount            string `json:"sellamount"`
	SymbolGroup           string `json:"symbolgroup"`
	StrikePrice           string `json:"strikeprice"`
	OptionType            string `json:"optiontype"`
	ExpiryDate            string `json:"expirydate"`
	LotSize               string `json:"lotsize"`
	CfBuyQty              string `json:"cfbuyqty"`
	CfSellQty             string `json:"cfsellqty"`
	CfBuyAmount           string `json:"cfbuyamount"`
	CfSellAmount          string `json:"cfsellamount"`
	BuyAveragePrice       string `json:"buyavgprice"`
	SellAveragePrice      string `json:"sellavgprice"`
	AverageNetPrice       string `json:"avgnetprice"`
	NetValue              string `json:"netvalue"`
	NetQty                string `json:"netqty"`
	TotalBuyValue         string `json:"totalbuyvalue"`
	TotalSellValue        string `json:"totalsellvalue"`
	CfBuyAveragePrice     string `json:"cfbuyavgprice"`
	CfSellAveragePrice    string `json:"cfsellavgprice"`
	TotalBuyAveragePrice  string `json:"totalbuyavgprice"`
	TotalSellAveragePrice string `json:"totalsellavgprice"`
	NetPrice              string `json:"netprice"`
}

// Positions represents a list of net and day positions.
type Positions []Position

// ConvertPositionParams represents the input params for a position conversion.
type ConvertPositionParams struct {
	Exchange        string `url:"exchange"`
	TradingSymbol   string `url:"tradingsymbol"`
	OldProductType  string `url:"oldproducttype"`
	NewProductType  string `url:"newproducttype"`
	TransactionType string `url:"transactiontype"`
	Quantity        int    `url:"quantity"`
	Type            string `json:"type"`
}

// GetOrderBook gets user orders.
func (c *Client) GetOrderBook() (Orders, error) {
	var orders Orders
	err := c.doEnvelope(http.MethodGet, URIGetOrderBook, nil, nil, &orders, true)
	return orders, err
}

// PlaceOrder places an order.
func (c *Client) PlaceOrder(orderParams OrderParams) (OrderResponse, error) {
	var (
		orderResponse OrderResponse
		params        map[string]interface{}
		err           error
	)

	params = structToMap(orderParams, "json")

	err = c.doEnvelope(http.MethodPost, URIPlaceOrder, params, nil, &orderResponse, true)
	return orderResponse, err
}

// ModifyOrder for modifying an order.
func (c *Client) ModifyOrder(modifyOrderParams ModifyOrderParams) (OrderResponse, error) {
	var (
		orderResponse OrderResponse
		params        map[string]interface{}
		err           error
	)

	params = structToMap(modifyOrderParams, "json")

	err = c.doEnvelope(http.MethodPost, URIModifyOrder, params, nil, &orderResponse, true)
	return orderResponse, err
}

// CancelOrder for cancellation of an order.
func (c *Client) CancelOrder(variety string, orderid string) (OrderResponse, error) {
	var (
		orderResponse OrderResponse
		err           error
	)

	params := make(map[string]interface{})
	params["variety"] = variety
	params["orderid"] = orderid

	err = c.doEnvelope(http.MethodPost, URICancelOrder, params, nil, &orderResponse, true)
	return orderResponse, err
}

// GetPositions gets user positions.
func (c *Client) GetPositions() (Positions, error) {
	var positions Positions
	err := c.doEnvelope(http.MethodGet, URIGetPositions, nil, nil, &positions, true)
	return positions, err
}

// GetTradeBook gets user trades.
func (c *Client) GetTradeBook() (Trades, error) {
	var trades Trades
	err := c.doEnvelope(http.MethodGet, URIGetTradeBook, nil, nil, &trades, true)
	return trades, err
}

// ConvertPosition converts position's product type.
func (c *Client) ConvertPosition(convertPositionParams ConvertPositionParams) error {
	var (
		params map[string]interface{}
		err    error
	)

	params = structToMap(convertPositionParams, "json")

	err = c.doEnvelope(http.MethodPost, URIConvertPosition, params, nil, nil, true)
	return err
}
