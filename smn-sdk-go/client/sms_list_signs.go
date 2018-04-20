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
)

// the request data of list sms signs
type ListSmsSignsRequest struct {
	*BaseRequest
}

// the response data of list sms signs
type ListSmsSignsResponse struct {
	*BaseResponse
	SmsSignCount int           `json:"sms_sign_count"`
	SmsSigns     []SmsSignInfo `json:"sms_signs"`
}

type SmsSignInfo struct {
	SignName   string `json:"sign_name"`
	CreateTime string `json:"create_time"`
	SignId     string `json:"sign_id"`
	Reply      string `json:"reply"`
	Status     int    `json:"status"`
	SignType   int    `json:"sign_type"`
}

// send request to list sms signs
func (client *SmnClient) ListSmsSigns(request *ListSmsSignsRequest) (response *ListSmsSignsResponse, err error) {
	response = &ListSmsSignsResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list sms signs request struct
func (client *SmnClient) NewListSmsSignsRequest() (request *ListSmsSignsRequest) {
	request = &ListSmsSignsRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the list sms signs request
func (request *ListSmsSignsRequest) GetUrl() (string, error) {
	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmsSignature, nil
}

// get the http method of the list sms signs request
func (request *ListSmsSignsRequest) GetMethod() string {
	return util.GET
}

// no body params
func (request *ListSmsSignsRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
