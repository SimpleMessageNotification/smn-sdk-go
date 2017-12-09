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

type ListSubscriptionsByTopicRequest struct {
	*BaseRequest
	TopicUrn string `json:"-"`
	Limit    string `json:"limit"`
	Offset   string `json:"offset"`
}

type ListSubscriptionsByTopicResponse struct {
	*BaseResponse
	SubscriptionCount int                `json:"subscription_count"`
	Subscriptions     []SubscriptionInfo `json:"subscriptions"`
}

func (client *SmnClient) ListSubscriptionsByTopic(request *ListSubscriptionsByTopicRequest) (response *ListSubscriptionsByTopicResponse, err error) {
	response = &ListSubscriptionsByTopicResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewListSubscriptionsByTopicRequest() (request *ListSubscriptionsByTopicRequest) {
	request = &ListSubscriptionsByTopicRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
		Offset:      "0",
		Limit:       "100",
	}
	return
}

func (request *ListSubscriptionsByTopicRequest) GetUrl() (urlStr string, err error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}
	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics + util.UrlDelimiter +
		request.TopicUrn + util.UrlDelimiter + util.Subscriptions
	urlEncoded, err := util.GetQueryParams(request)
	if err != nil {
		return
	}

	urlStr += "?" + urlEncoded
	return
}

func (request *ListSubscriptionsByTopicRequest) GetMethod() string {
	return util.GET
}

func (request *ListSubscriptionsByTopicRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
