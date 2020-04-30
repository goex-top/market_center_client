# move to https://github.com/goex-top/market_center/tree/master/client

# Market Center Client
![HitCount](http://hits.dwyl.io/goex-top/market_center_client.svg)

Client for [market center](https://github.com/goex-top/market_center)

## APIs

* SubscribeTicker(exchangeName, currencyPair, period) (error)
* SubscribeDepth(exchangeName, currencyPair, period) (error)
* SubscribeTrade(exchangeName, currencyPair, period) (error)
* GetTicker(exchangeName, currencyPair) (*Ticker, error)
* GetDepth(exchangeName, currencyPair) (*Depth, error)
* GetTrade(exchangeName, currencyPair) (*Trade, error)
* GetSupportList() []

### 观星者

![观星者](https://starchart.cc/goex-top/market_center_client.svg)
