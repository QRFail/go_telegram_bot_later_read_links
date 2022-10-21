package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"go_telegram_bot_later_read_links/lib/e"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host string, token string) Client {

	return Client{
		host:     host,
		basePath: NewBasePath(token),
		client:   http.Client{},
	}

}

func NewBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(limit int, offset int) ([]Update, error) {

	query := url.Values{}

	query.Add("limit", strconv.Itoa(limit))
	query.Add("offset", strconv.Itoa(offset))

	data, err := c.doRequest(getUpdatesMethod, query)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil

}

func (c *Client) SendMessage(chatID int, text string) error {
	query := url.Values{}

	query.Add("chat_id", strconv.Itoa(chatID))
	query.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, query)
	if err != nil {
		return e.Wrap("can't do request %w", err)
	}

	return nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {

	defer func() {
		err = e.WrapIfErr("can't do request %w", err)
	}()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
