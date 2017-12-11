/*
 * Copyright (C) 2017. Huawei Technologies Co., LTD. All rights reserved.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of Apache License, Version 2.0.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * Apache License, Version 2.0 for more details.
 */
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
