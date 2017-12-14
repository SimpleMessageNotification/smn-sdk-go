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
	"github.com/SimpleMessageNotification/smn-sdk-go/smn-sdk-go/util"
	"io"
	"fmt"
)

// the request data of get sms message
type GetSmsMessageRequest struct {
	*BaseRequest
	MessageId string
}

// the response data of get sms message
type GetSmsMessageResponse struct {
	*BaseResponse
	Message string `json:"message"`
}

// send request to get sms message
func (client *SmnClient) GetSmsMessage(request *GetSmsMessageRequest) (response *GetSmsMessageResponse, err error) {
	response = &GetSmsMessageResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

// create a new get sms message request struct
func (client *SmnClient) NewGetSmsMessageRequest() (request *GetSmsMessageRequest) {
	request = &GetSmsMessageRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the get sms message request
func (request *GetSmsMessageRequest) GetUrl() (urlStr string, err error) {
	if request.MessageId == "" {
		return "", fmt.Errorf("messageId is null")
	}
	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmnProtocolSms +
		util.UrlDelimiter + util.Message + util.UrlDelimiter + request.MessageId
	return
}

// get the http method of the get sms message request
func (request *GetSmsMessageRequest) GetMethod() string {
	return util.GET
}

// no body params
func (request *GetSmsMessageRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
