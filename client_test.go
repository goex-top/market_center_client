package market_center_client

import (
	"fmt"
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

func TestNewClient(t *testing.T) {
	c1 := NewClient()
	c2 := NewClient()
	c1.SubscribeDepth("binance.com", "BTC_USDT", 500)
	c2.SubscribeTicker("binance.com", "BTC_USDT", 600)
	c2.SubscribeTicker("binance.com", "BTC_USDT", 300)
	//c1.EnableDebug()
	//c2.EnableDebug()
	for {
		fmt.Println(c1.GetDepth("binance.com", "BTC_USDT"))
		fmt.Println(c2.GetDepth("binance.com", "BTC_USDT"))
		fmt.Println(c2.GetTicker("binance.com", "BTC_USDT"))
		time.Sleep(time.Second)
	}

}
