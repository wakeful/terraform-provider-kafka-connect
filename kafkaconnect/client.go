package kafkaconnect

import (
	"bytes"
	"io"
	"net/http"
)

type KafkaConnectClient struct {
	Url        string
	HTTPClient *http.Client
}

func (c *KafkaConnectClient) Do(method string, endpoint string, payload []byte) (*http.Response, error) {

	var body io.Reader
	if payload != nil {
		body = bytes.NewReader(payload)
	}

	req, err := http.NewRequest(method, c.Url+"/connectors/"+endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Close = true
	resp, err := c.HTTPClient.Do(req)

	return resp, err
}

func (c *KafkaConnectClient) Get(connector string) (*http.Response, error) {
	return c.Do("GET", connector, nil)
}

func (c *KafkaConnectClient) Post(payload []byte) (*http.Response, error) {
	return c.Do("POST", "", payload)
}

func (c *KafkaConnectClient) Delete(connector string) (*http.Response, error) {
	return c.Do("DELETE", connector, nil)
}
