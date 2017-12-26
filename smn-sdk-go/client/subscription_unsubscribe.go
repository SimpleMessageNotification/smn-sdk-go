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

// the request data of unSubscribe
type UnsubscribeRequest struct {
	*BaseRequest
	SubscriptionUrn string
}

// the response data of unSubscribe
type UnsubscribeResponse struct {
	*BaseResponse
}

// send request to unSubscribe
func (client *SmnClient) Unsubscribe(request *UnsubscribeRequest) (response *UnsubscribeResponse, err error) {
	response = &UnsubscribeResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new unSubscribe request struct
func (client *SmnClient) NewUnsubscribeRequest() (request *UnsubscribeRequest) {
	request = &UnsubscribeRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the unSubscribe request
func (request *UnsubscribeRequest) GetUrl() (urlStr string, err error) {
	if request.SubscriptionUrn == "" {
		return "", fmt.Errorf("subscription urn is null")
	}

	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter +
		util.Subscriptions + util.UrlDelimiter + request.SubscriptionUrn
	return
}

// get the http method of the unSubscribe request
func (request *UnsubscribeRequest) GetMethod() string {
	return util.DELETE
}

// get the body params of the unSubscribe request
func (request *UnsubscribeRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
