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
package client

import (
	"net/http"
	"io/ioutil"
)

type BaseResponse struct {
	httpStatus    int    `json:"-"`
	contentString string `json:"-"`
	contentBytes  []byte `json:"-"`
	RequestId     string `json:"request_id""`
	*ErrorResponse
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SmnResponse interface {
	parseHttpResponse(httpResponse *http.Response) error
	GetContentBytes() []byte
}

func (response *BaseResponse) parseHttpResponse(httpResponse *http.Response) (err error) {
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}
	response.httpStatus = httpResponse.StatusCode
	response.contentBytes = body
	response.contentString = string(body)
	return
}

func (response *BaseResponse) GetContentBytes() []byte {
	return response.contentBytes
}

func (response *BaseResponse) GetContentString() string {
	return response.contentString
}

func (response *BaseResponse) GetHttpStatus() int {
	return response.httpStatus
}

func (response *BaseResponse) IsSuccess() bool {
	if response.GetHttpStatus() >= 200 && response.GetHttpStatus() < 300 {
		return true
	}

	return false
}
