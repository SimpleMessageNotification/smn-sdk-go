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
	"smn-sdk-go/util"
	"io"
	"fmt"
)

type SmsPublishRequest struct {
	*BaseRequest
	Message  string `json:"message""`
	EndPoint string `json:"endpoint"`
	SignId   string `json:"sign_id"`
}

type SmsPublishResponse struct {
	*BaseResponse
	MessageId string `json:"message_id""`
}

func (client *SmnClient) SmsPublish(request *SmsPublishRequest) (response *SmsPublishResponse, err error) {
	response = &SmsPublishResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewSmsPublishRequest() (request *SmsPublishRequest) {
	request = &SmsPublishRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *SmsPublishRequest) GetUrl() (string, error) {
	if request.Message == "" {
		return "", fmt.Errorf("message is null")
	}

	if request.SignId == "" {
		return "", fmt.Errorf("signId is null")
	}

	if request.EndPoint == "" || !util.ValidatePhone(request.EndPoint) {
		return "", fmt.Errorf("endpoint is invalid")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmnProtocolSms, nil
}

func (request *SmsPublishRequest) GetMethod() string {
	return util.POST
}

func (request *SmsPublishRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
