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

// the request data of list sms event
type ListSmsEventRequest struct {
	*BaseRequest
	EventType string `json:"event_type"`
}

// the response data of list sms event
type ListSmsEventResponse struct {
	*BaseResponse
	TopicUrn string         `json:"topic_urn"`
	Callback []CallbackInfo `json:"callback"`
}

// the  data of sms event callback info
type CallbackInfo struct {
	EventType string `json:"event_type"`
	TopicUrn  string `json:"topic_urn"`
}

// send request to list sms event
func (client *SmnClient) ListSmsEvent(request *ListSmsEventRequest) (response *ListSmsEventResponse, err error) {
	response = &ListSmsEventResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

// create a new list sms event request struct
func (client *SmnClient) NewListSmsEventRequest() (request *ListSmsEventRequest) {
	request = &ListSmsEventRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the list sms event request
func (request *ListSmsEventRequest) GetUrl() (urlStr string, err error) {
	if request.EventType != "" && !util.ValidateSmsEventType(request.EventType) {
		return "", fmt.Errorf("eventType is invalid")
	}

	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmnProtocolSms +
		util.UrlDelimiter + util.CallBack
	urlEncoded, err := util.GetQueryParams(request)
	if err != nil {
		return
	}

	urlStr += "?" + urlEncoded
	return
}

// get the http method of the list sms event request
func (request *ListSmsEventRequest) GetMethod() string {
	return util.GET
}

// no body params
func (request *ListSmsEventRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
