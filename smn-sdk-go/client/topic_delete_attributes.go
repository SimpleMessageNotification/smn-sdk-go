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

// the request data of delete topic attributes
type DeleteTopicAttributesRequest struct {
	*BaseRequest
	TopicUrn string
}

// the response data of delete topic attributes
type DeleteTopicAttributeResponse struct {
	*BaseResponse
}

// send request to delete topic attributes
func (client *SmnClient) DeleteTopicAttributes(request *DeleteTopicAttributesRequest) (response *DeleteTopicAttributeResponse, err error) {
	response = &DeleteTopicAttributeResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new delete topic attributes request struct
func (client *SmnClient) NewDeleteTopicAttributesRequest() (request *DeleteTopicAttributesRequest) {
	request = &DeleteTopicAttributesRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the delete topic attributes request
func (request *DeleteTopicAttributesRequest) GetUrl() (urlStr string, err error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn + util.UrlDelimiter +
		util.Attributes
	return
}

// get the http method of the delete topic attributes request
func (request *DeleteTopicAttributesRequest) GetMethod() string {
	return util.DELETE
}

// no body param
func (request *DeleteTopicAttributesRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
