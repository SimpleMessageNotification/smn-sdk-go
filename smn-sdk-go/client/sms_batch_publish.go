/*
 * Copyright (C) 2018. Huawei Technologies Co., LTD. All rights reserved.
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

// the request data of batch publish sms
type SmsBatchPublishRequest struct {
	*BaseRequest
	Message                string   `json:"message""`
	EndPoints              []string `json:"endpoints"`
	SignId                 string   `json:"sign_id,omitempty"`
	MessageIncludeSignFlag bool     `json:"message_include_sign_flag"`
}

//the response data of batch publish sms
type SmsBatchPublishResponse struct {
	*BaseResponse
	Result []SmsBatchPublishResult `json:"result"`
}

type SmsBatchPublishResult struct {
	MessageId string `json:"message_id"`
	endpoint  string `json:"endpoint"`
	code      string `json:"code"`
	message   string `json:"message"`
}

// send request to batch publish sms
func (client *SmnClient) SmsBatchPublish(request *SmsBatchPublishRequest) (response *SmsBatchPublishResponse, err error) {
	response = &SmsBatchPublishResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list sms batch publish request struct
func (client *SmnClient) NewSmsBatchPublishRequest() (request *SmsBatchPublishRequest) {
	request = &SmsBatchPublishRequest{
		BaseRequest:            &BaseRequest{Headers: make(map[string]string)},
		MessageIncludeSignFlag: false,
	}
	return
}

// get the url of the list sms batch publish request
func (request *SmsBatchPublishRequest) GetUrl() (string, error) {
	if request.Message == "" {
		return "", fmt.Errorf("message is null")
	}

	if request.EndPoints == nil || len(request.EndPoints) == 0 {
		return "", fmt.Errorf("endpoints is empty")
	}

	if !request.MessageIncludeSignFlag && request.SignId == "" {
		return "", fmt.Errorf("signId is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmnProtocolSms, nil
}

// get the http method of the sms batch publish request
func (request *SmsBatchPublishRequest) GetMethod() string {
	return util.POST
}

// get the body params of the sms batch publish request
func (request *SmsBatchPublishRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
