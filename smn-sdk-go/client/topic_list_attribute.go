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
	"fmt"
	"github.com/SimpleMessageNotification/smn-sdk-go/smn-sdk-go/util"
	"io"
)

// the request data of list topic attribute
type ListTopicAttributesRequest struct {
	*BaseRequest
	TopicUrn string `json:"-"`
	Name     string `json:"name"`
}

// the response data of list topic attribute
type ListTopicAttributesResponse struct {
	*BaseResponse
	Attributes map[string]interface{} `json:"attributes"`
}

// send request to list topic attribute
func (client *SmnClient) ListTopicAttributes(request *ListTopicAttributesRequest) (response *ListTopicAttributesResponse, err error) {
	response = &ListTopicAttributesResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

// create a new list topic attribute request struct
func (client *SmnClient) NewListTopicAttributesRequest() (request *ListTopicAttributesRequest) {
	request = &ListTopicAttributesRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the list topic attribute request
func (request *ListTopicAttributesRequest) GetUrl() (urlStr string, err error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}
	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn + util.UrlDelimiter +
		util.Attributes

	urlEncoded, err := util.GetQueryParams(request)
	if err != nil {
		return
	}
	if urlEncoded != "" {
		urlStr += "?" + urlEncoded
	}
	return
}

// get the http method of the list topic attribute request
func (request *ListTopicAttributesRequest) GetMethod() string {
	return util.GET
}

// no body param
func (request *ListTopicAttributesRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
