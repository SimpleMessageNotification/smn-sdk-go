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

// the request data of publish sms
type SmsPublishRequest struct {
	*BaseRequest
	Message                string `json:"message""`
	EndPoint               string `json:"endpoint"`
	SignId                 string `json:"sign_id,omitempty"`
	MessageIncludeSignFlag bool   `json:"message_include_sign_flag"`
}

// the response data of publish sms
type SmsPublishResponse struct {
	*BaseResponse
	MessageId string `json:"message_id""`
}

// send request to publish sms
func (client *SmnClient) SmsPublish(request *SmsPublishRequest) (response *SmsPublishResponse, err error) {
	response = &SmsPublishResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list sms publish request struct
func (client *SmnClient) NewSmsPublishRequest() (request *SmsPublishRequest) {
	request = &SmsPublishRequest{
		BaseRequest:            &BaseRequest{Headers: make(map[string]string)},
		MessageIncludeSignFlag: false,
	}
	return
}

// get the url of the list sms publish request
func (request *SmsPublishRequest) GetUrl() (string, error) {
	if request.Message == "" {
		return "", fmt.Errorf("message is null")
	}

	if request.EndPoint == "" {
		return "", fmt.Errorf("endpoint is invalid")
	}

	if !request.MessageIncludeSignFlag && request.SignId == "" {
		return "", fmt.Errorf("signId is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmnProtocolSms, nil
}

// get the http method of the sms publish request
func (request *SmsPublishRequest) GetMethod() string {
	return util.POST
}

// get the body params of the sms publish request
func (request *SmsPublishRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
