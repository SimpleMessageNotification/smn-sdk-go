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

// the request data of delete topic attribute by name
type DeleteTopicAttributeByNameRequest struct {
	*BaseRequest
	TopicUrn string
	Name     string
}

// the response data of delete topic attribute by name
type DeleteTopicAttributeByNameResponse struct {
	*BaseResponse
}

// send request to delete topic attribute by name
func (client *SmnClient) DeleteTopicAttributeByName(request *DeleteTopicAttributeByNameRequest) (response *DeleteTopicAttributeByNameResponse, err error) {
	response = &DeleteTopicAttributeByNameResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new delete topic attribute by name request struct
func (client *SmnClient) NewDeleteTopicAttributeByNameRequest() (request *DeleteTopicAttributeByNameRequest) {
	request = &DeleteTopicAttributeByNameRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the delete topic attribute by name request
func (request *DeleteTopicAttributeByNameRequest) GetUrl() (urlStr string, err error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}
	if request.Name == "" {
		return "", fmt.Errorf("name is null")
	}
	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn + util.UrlDelimiter +
		util.Attributes + util.UrlDelimiter + request.Name
	return
}

// get the http method of the delete topic attribute by name request
func (request *DeleteTopicAttributeByNameRequest) GetMethod() string {
	return util.DELETE
}

// nobody param
func (request *DeleteTopicAttributeByNameRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
