package commom

import (
	"time"
	"net/http"
)

type ClientConfiguration struct {
	Timeout   time.Duration
	Transport *http.Transport
}

func NewClientConfig() (config *ClientConfiguration) {
	config = &ClientConfiguration{}
	return
}

func (config *ClientConfiguration) SetTimeout(duration time.Duration) {
	config.Timeout = duration
}

func (config *ClientConfiguration) SetTransport(transport *http.Transport) {
	config.Transport = transport
}
