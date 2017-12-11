package commom

import (
	"time"
	"net/http"
)

// the http client config
type ClientConfiguration struct {
	Timeout   time.Duration
	Transport *http.Transport
}

// create a new client config
func NewClientConfig() (config *ClientConfiguration) {
	config = &ClientConfiguration{}
	return
}

// set timeout
func (config *ClientConfiguration) SetTimeout(duration time.Duration) {
	config.Timeout = duration
}

// set transport
func (config *ClientConfiguration) SetTransport(transport *http.Transport) {
	config.Transport = transport
}
