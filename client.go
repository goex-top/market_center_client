package market_center_client

import (
	"encoding/json"
	"errors"
	. "github.com/goex-top/market_center"
	"github.com/mitchellh/mapstructure"
	goex "github.com/nntaoli-project/GoEx"
	log "github.com/sirupsen/logrus"
	"net"
)

type Client struct {
	con net.Conn
}

func NewClient() *Client {
	return NewClientWithPath(UDS_PATH)
}

func NewClientWithPath(udsPath string) *Client {
	c, err := net.Dial("unix", udsPath)
	if err != nil {
		log.Fatalf("Dial error: %v", err)
	}

	return &Client{con: c}
}

func (c *Client) EnableDebug() {
	log.SetLevel(log.DebugLevel)
}

func (c *Client) Close() {
	log.Debugln("Close")
	c.con.Close()
}

func (c *Client) newUdsRequest(req *Request) (*Response, error) {
	r, err := json.Marshal(req)
	c.con.Write(r)

	buf := make([]byte, 1024)
	n, err := c.con.Read(buf[:])
	if err != nil {
		log.Errorln(err, n)
		return nil, err
	}
	log.Debugln("Client got:", string(buf[:n]))
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
	log.Debugln("GetSupportList")
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
	log.Debugln("GetDepth")
	req := &Request{}
	req.Type = Type_GetDepth
	req.ExchangeName = exchange
	req.CurrencyPair = pair
	rsp, err := c.newUdsRequest(req)
	if err != nil {
		return nil, err
	}
	depth := &goex.Depth{}
	r := rsp.Data.(map[string]interface{})
	mapstructure.Decode(r, depth)
	return depth, nil
}

func (c *Client) GetTicker(exchange, pair string) (*goex.Ticker, error) {
	log.Debugln("GetTicker")
	req := &Request{}
	req.Type = Type_GetTicker
	req.ExchangeName = exchange
	req.CurrencyPair = pair
	rsp, err := c.newUdsRequest(req)
	if err != nil {
		return nil, err
	}
	ticker := &goex.Ticker{}
	r := rsp.Data.(map[string]interface{})

	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           ticker,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(r)
	if err != nil {
		return nil, err
	}

	return ticker, nil
}

func (c *Client) SubscribeDepth(exchange, pair string, period int64) error {
	log.Debugln("SubscribeDepth")
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
	log.Debugln("SubscribeTicker")
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
