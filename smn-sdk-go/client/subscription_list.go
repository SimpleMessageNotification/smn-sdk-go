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

// the request data of list subscriptions
type ListSubscriptionsRequest struct {
	*BaseRequest
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// the response data of list subscriptions
type ListSubscriptionsResponse struct {
	*BaseResponse
	SubscriptionCount int                `json:"subscription_count"`
	Subscriptions     []SubscriptionInfo `json:"subscriptions"`
}

// the subscription data info
type SubscriptionInfo struct {
	TopicUrn        string `json:"topic_urn"`
	Protocol        string `json:"protocol"`
	SubscriptionUrn string `json:"subscription_urn"`
	Owner           string `json:"owner"`
	Endpoint        string `json:"endpoint"`
	Remark          string `json:"remark"`
	Status          int    `json:"status"`
}

// send request to list subscriptions
func (client *SmnClient) ListSubscriptions(request *ListSubscriptionsRequest) (response *ListSubscriptionsResponse, err error) {
	response = &ListSubscriptionsResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list subscriptions request struct
func (client *SmnClient) NewListSubscriptionsRequest() (request *ListSubscriptionsRequest) {
	request = &ListSubscriptionsRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
		Offset:      0,
		Limit:       100,
	}
	return
}

// get the url of the list subscriptions request
func (request *ListSubscriptionsRequest) GetUrl() (urlStr string, err error) {
	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Subscriptions
	urlEncoded, err := util.GetQueryParams(request)
	if err != nil {
		return
	}

	urlStr += "?" + urlEncoded
	return
}

// get the http method of the list subscriptions request
func (request *ListSubscriptionsRequest) GetMethod() string {
	return util.GET
}

// no body param
func (request *ListSubscriptionsRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
