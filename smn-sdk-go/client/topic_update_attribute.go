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

type UpdateTopicAttributeRequest struct {
	*BaseRequest
	TopicUrn string `json:"-"`
	Name     string `json:"-"`
	Value    string `json:"value"`
}

type UpdateTopicAttributeResponse struct {
	*BaseResponse
}

func (client *SmnClient) UpdateTopicAttribute(request *UpdateTopicAttributeRequest) (response *UpdateTopicAttributeResponse, err error) {
	response = &UpdateTopicAttributeResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewUpdateTopicAttributeRequest() (request *UpdateTopicAttributeRequest) {
	request = &UpdateTopicAttributeRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *UpdateTopicAttributeRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn + util.UrlDelimiter +
		util.Attributes + util.UrlDelimiter + request.Name, nil
}

func (request *UpdateTopicAttributeRequest) GetMethod() string {
	return util.PUT
}

func (request *UpdateTopicAttributeRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
