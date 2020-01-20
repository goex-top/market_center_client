# Market Center Client
client for [market center](https://github.com/goex-top/market_center)

## APIs

* SubscribeTicker(exchangeName, currencyPair, period) (error)
* SubscribeDepth(exchangeName, currencyPair, period) (error)
* SubscribeTrade(exchangeName, currencyPair, period) (error)
* GetTicker(exchangeName, currencyPair) (*Ticker, error)
* GetDepth(exchangeName, currencyPair) (*Depth, error)
* GetTrade(exchangeName, currencyPair) (*Trade, error)
* GetSupportList() []

