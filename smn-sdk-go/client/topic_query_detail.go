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

// the request data of query topic detail
type QueryTopicDetailRequest struct {
	*BaseRequest
	TopicUrn string
}

// the response data of query topic detail
type QueryTopicDetailResponse struct {
	*BaseResponse
	TopicUrn    string `json:"topic_urn"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	PushPolicy  int    `json:"push_policy"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

// send request to query topic detail
func (client *SmnClient) QueryTopicDetail(request *QueryTopicDetailRequest) (response *QueryTopicDetailResponse, err error) {
	response = &QueryTopicDetailResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

// create a new query topic detail request struct
func (client *SmnClient) NewQueryTopicDetailRequest() (request *QueryTopicDetailRequest) {
	request = &QueryTopicDetailRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the query topic detail request
func (request *QueryTopicDetailRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn, nil
}

// get the http method of the query topic detail request
func (request *QueryTopicDetailRequest) GetMethod() string {
	return util.GET
}

// no body param
func (request *QueryTopicDetailRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
