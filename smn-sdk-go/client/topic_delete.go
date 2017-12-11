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
	"github.com/smn-sdk-go/smn-sdk-go/util"
	"io"
	"fmt"
)

// the request data of delete topic
type DeleteTopicRequest struct {
	*BaseRequest
	TopicUrn string
}

// the response data of delete topic
type DeleteTopicResponse struct {
	*BaseResponse
}

// send request to delete topic
func (client *SmnClient) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
	response = &DeleteTopicResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

// create a new delete topic request struct
func (client *SmnClient) NewDeleteTopicRequest() (request *DeleteTopicRequest) {
	request = &DeleteTopicRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the delete topic request
func (request *DeleteTopicRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn, nil
}

// get the http method of the delete topic request
func (request *DeleteTopicRequest) GetMethod() string {
	return util.DELETE
}

// nobody param
func (request *DeleteTopicRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
