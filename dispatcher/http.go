package dispatcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/komand/plugin-sdk-go2/message"
)

// HTTP will dispatch via HTTP
type HTTP struct {
	URL    string `json:"url"`
	client *http.Client
}

// NewHTTP will return a new HTTP Dispatcher
func NewHTTP(url string) *HTTP {
	// TODO build a proper http.Client, don't rely on the 0val default like we did in the v1 sdk
	return &HTTP{
		URL: url,
	}
}

// Send dispatches a trigger event
func (d *HTTP) Send(e *message.Response) error {
	messageBytes, err := json.Marshal(e)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", d.URL, bytes.NewBuffer(messageBytes))
	if err != nil {
		return fmt.Errorf("Unable to POST to dispatcher: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := d.client.Do(req)
	// Since we don't even touch the body, the Close call will fast-track an ioutil.Discard copy
	// this will never return an error we need to care about here afaik. Explicitly _-ing the error by design
	_ = resp.Body.Close()
	if err != nil {
		return fmt.Errorf("Unable to send event to http dispatcher: %s", err.Error())
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Response failed, stopping trigger: %s %+v", resp.Status, resp.Header)
	}
	return nil
}
