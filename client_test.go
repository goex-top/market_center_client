package market_center_client

import (
	goex "github.com/nntaoli-project/GoEx"
	"testing"
	"time"
)

var client = NewClient()

func TestClient_GetSupportList(t *testing.T) {
	t.Log(client.GetSupportList())
}

func TestClient_SubscribeDepth(t *testing.T) {
	t.Log(client.SubscribeDepth(goex.BINANCE, "BTC_USDT", 200))
}

func TestClient_GetDepth(t *testing.T) {
	client.SubscribeDepth(goex.BINANCE, "BTC_USDT", 200)
	time.Sleep(time.Second)
	t.Log(client.GetDepth(goex.BINANCE, "BTC_USDT"))
}

func TestClient_GetTicker(t *testing.T) {
	client.SubscribeTicker(goex.BINANCE, "BTC_USDT", 200)
	time.Sleep(time.Second)
	t.Log(client.GetTicker(goex.BINANCE, "BTC_USDT"))
}
