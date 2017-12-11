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
)

type ListSubscriptionsRequest struct {
	*BaseRequest
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

type ListSubscriptionsResponse struct {
	*BaseResponse
	SubscriptionCount int                `json:"subscription_count"`
	Subscriptions     []SubscriptionInfo `json:"subscriptions"`
}

type SubscriptionInfo struct {
	TopicUrn        string `json:"topic_urn"`
	Protocol        string `json:"protocol"`
	SubscriptionUrn string `json:"subscription_urn"`
	Owner           string `json:"owner"`
	Endpoint        string `json:"endpoint"`
	Remark          string `json:"remark"`
	Status          int    `json:"status"`
}

func (client *SmnClient) ListSubscriptions(request *ListSubscriptionsRequest) (response *ListSubscriptionsResponse, err error) {
	response = &ListSubscriptionsResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewListSubscriptionsRequest() (request *ListSubscriptionsRequest) {
	request = &ListSubscriptionsRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
		Offset:      "0",
		Limit:       "100",
	}
	return
}

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

func (request *ListSubscriptionsRequest) GetMethod() string {
	return util.GET
}

func (request *ListSubscriptionsRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
