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
	t.Log(client.SubscribeDepth(goex.BINANCE, "EOS_USDT", 100))
}

func TestClient_GetDepth(t *testing.T) {
	client.SubscribeDepth(goex.BINANCE, "EOS_USDT", 100)
	time.Sleep(time.Second)
	t.Log(client.GetDepth(goex.BINANCE, "EOS_USDT"))
}

func TestClient_GetTicker(t *testing.T) {
	client.SubscribeTicker(goex.BINANCE, "EOS_USDT", 100)
	time.Sleep(time.Second)
	t.Log(client.GetTicker(goex.BINANCE, "EOS_USDT"))
}
