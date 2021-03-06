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

// the request data of subscribe
type SubscribeRequest struct {
	*BaseRequest
	TopicUrn string `json:"-"`
	Endpoint string `json:"endpoint"`
	Protocol string `json:"protocol"`
	Remark   string `json:"remark"`
}

// the response data of subscribe
type SubscribeResponse struct {
	*BaseResponse
	SubscriptionUrn string `json:"subscription_urn"`
}

// send request to subscribe
func (client *SmnClient) Subscribe(request *SubscribeRequest) (response *SubscribeResponse, err error) {
	response = &SubscribeResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new subscribe request struct
func (client *SmnClient) NewSubscribeRequest() (request *SubscribeRequest) {
	request = &SubscribeRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the subscribe request
func (request *SubscribeRequest) GetUrl() (urlStr string, err error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}
	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics + util.UrlDelimiter +
		request.TopicUrn + util.UrlDelimiter + util.Subscriptions, nil
}

// get the http method of the subscribe request
func (request *SubscribeRequest) GetMethod() string {
	return util.POST
}

// get the body params of the subscribe request
func (request *SubscribeRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
