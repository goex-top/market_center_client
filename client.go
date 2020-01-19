package market_center_client

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/goex-top/market_center"
	goex "github.com/nntaoli-project/GoEx"
	"log"
	"net"
)

type Client struct {
	con net.Conn
}

func NewClient() *Client {
	c, err := net.Dial("unix", UDS_PATH)
	if err != nil {
		log.Fatal("Dial error: ", err)
	}

	return &Client{con: c}
}

func (c *Client) Close() {
	c.con.Close()
}

func (c *Client) newUdsRequest(req *Request) (*Response, error) {
	r, err := json.Marshal(req)
	fmt.Println(err, string(r))
	c.con.Write(r)

	buf := make([]byte, 1024)
	n, err := c.con.Read(buf[:])
	if err != nil {
		fmt.Println(err, n)
		return nil, err
	}
	fmt.Println("Client got:", string(buf[:n]))
	rsp := Response{}
	err = json.Unmarshal(buf[:n], &rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Status != 0 {
		return nil, errors.New(rsp.ErrorMessage)
	}
	return &rsp, err
}

func (c *Client) GetSupportList() ([]string, error) {
	req := &Request{}
	req.Type = Type_GetSupportList
	rsp, err := c.newUdsRequest(req)
	if err != nil {
		return nil, err
	}
	l := rsp.Data.([]interface{})
	list := make([]string, 0)
	for _, v := range l {
		list = append(list, v.(string))
	}
	return list, nil
}

func (c *Client) GetDepth(exchange, pair string) (*goex.Depth, error) {
	req := &Request{}
	req.Type = Type_GetDepth
	req.ExchangeName = exchange
	req.CurrencyPair = pair
	rsp, err := c.newUdsRequest(req)
	if err != nil {
		return nil, err
	}

	return rsp.Data.(*goex.Depth), nil
}

func (c *Client) GetTicker(exchange, pair string) (*goex.Ticker, error) {
	req := &Request{}
	req.Type = Type_GetDepth
	req.ExchangeName = exchange
	req.CurrencyPair = pair
	rsp, err := c.newUdsRequest(req)
	if err != nil {
		return nil, err
	}

	return rsp.Data.(*goex.Ticker), nil
}

func (c *Client) SubscribeDepth(exchange, pair string, period int64) error {
	req := &Request{}
	req.Type = Type_SubscribeDepth
	req.ExchangeName = exchange
	req.CurrencyPair = pair
	req.Period = period
	_, err := c.newUdsRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SubscribeTicker(exchange, pair string, period int64) error {
	req := &Request{}
	req.Type = Type_SubscribeTicker
	req.ExchangeName = exchange
	req.CurrencyPair = pair
	req.Period = period
	_, err := c.newUdsRequest(req)
	if err != nil {
		return err
	}

	return nil
}
