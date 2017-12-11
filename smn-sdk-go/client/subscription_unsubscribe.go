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

type UnsubscribeRequest struct {
	*BaseRequest
	SubscriptionUrn string
}

type UnsubscribeResponse struct {
	*BaseResponse
}

func (client *SmnClient) Unsubscribe(request *UnsubscribeRequest) (response *UnsubscribeResponse, err error) {
	response = &UnsubscribeResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewUnsubscribeRequest() (request *UnsubscribeRequest) {
	request = &UnsubscribeRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *UnsubscribeRequest) GetUrl() (urlStr string, err error) {
	if request.SubscriptionUrn == "" {
		return "", fmt.Errorf("subscription urn is null")
	}

	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter +
		util.Subscriptions + util.UrlDelimiter + request.SubscriptionUrn
	return
}

func (request *UnsubscribeRequest) GetMethod() string {
	return util.DELETE
}

func (request *UnsubscribeRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
