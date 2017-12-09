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

type ListSmsEventRequest struct {
	*BaseRequest
	EventType string `json:"event_type"`
}

type ListSmsEventResponse struct {
	*BaseResponse
	TopicUrn string         `json:"topic_urn"`
	Callback []CallbackInfo `json:"callback"`
}

type CallbackInfo struct {
	EventType string `json:"event_type"`
	TopicUrn  string `json:"topic_urn"`
}

func (client *SmnClient) ListSmsEvent(request *ListSmsEventRequest) (response *ListSmsEventResponse, err error) {
	response = &ListSmsEventResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewListSmsEventRequest() (request *ListSmsEventRequest) {
	request = &ListSmsEventRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

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

func (request *ListSmsEventRequest) GetMethod() string {
	return util.GET
}

func (request *ListSmsEventRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
