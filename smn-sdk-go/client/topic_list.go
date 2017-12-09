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

type ListTopicRequest struct {
	*BaseRequest
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

type ListTopicResponse struct {
	*BaseResponse
	TopicCount int         `json:"topic_count"`
	Topics     []TopicInfo `json:"topics"`
}

type TopicInfo struct {
	TopicUrn    string `json:"topic_urn"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	PushPolicy  int    `json:"push_policy"`
}

func (client *SmnClient) ListTopic(request *ListTopicRequest) (response *ListTopicResponse, err error) {
	response = &ListTopicResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewListTopicRequest() (request *ListTopicRequest) {
	request = &ListTopicRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
		Offset:      "0",
		Limit:       "100",
	}
	return
}

func (request *ListTopicRequest) GetUrl() (urlStr string, err error) {
	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics
	urlEncoded, err := util.GetQueryParams(request)
	if err != nil {
		return
	}

	urlStr += "?" + urlEncoded
	return
}

func (request *ListTopicRequest) GetMethod() string {
	return util.GET
}

func (request *ListTopicRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
