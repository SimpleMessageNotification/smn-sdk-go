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

// the request data of update topic attribute
type UpdateTopicAttributeRequest struct {
	*BaseRequest
	TopicUrn string `json:"-"`
	Name     string `json:"-"`
	Value    string `json:"value"`
}

// the response data of update topic attribute
type UpdateTopicAttributeResponse struct {
	*BaseResponse
}

// send request to update topic
func (client *SmnClient) UpdateTopicAttribute(request *UpdateTopicAttributeRequest) (response *UpdateTopicAttributeResponse, err error) {
	response = &UpdateTopicAttributeResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new update topic attribute request struct
func (client *SmnClient) NewUpdateTopicAttributeRequest() (request *UpdateTopicAttributeRequest) {
	request = &UpdateTopicAttributeRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the update topic attribute request
func (request *UpdateTopicAttributeRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn + util.UrlDelimiter +
		util.Attributes + util.UrlDelimiter + request.Name, nil
}

// get the http method of the update topic attribute request
func (request *UpdateTopicAttributeRequest) GetMethod() string {
	return util.PUT
}

// get the body params of the update topic attribute request
func (request *UpdateTopicAttributeRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
