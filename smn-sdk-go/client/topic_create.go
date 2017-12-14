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

// the request data of create topic
type CreateTopicRequest struct {
	*BaseRequest
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

// the response data of create topic
type CreateTopicResponse struct {
	*BaseResponse
	TopicUrn string `json:"topic_urn"`
}

// send request to create topic
func (client *SmnClient) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
	response = &CreateTopicResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

// create a new create topic request struct
func (client *SmnClient) NewCreateTopicRequest() (request *CreateTopicRequest) {
	request = &CreateTopicRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the create topic request
func (request *CreateTopicRequest) GetUrl() (string, error) {
	if !util.ValidateTopicName(request.Name) {
		return "", fmt.Errorf("topic name is invalid")
	}
	if !util.ValidateTopicDiplayName(request.DisplayName) {
		return "", fmt.Errorf("topic displayName is overSize")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics, nil
}

// get the http method of the create topic request
func (request *CreateTopicRequest) GetMethod() string {
	return util.POST
}

// get the body params of the create topic request
func (request *CreateTopicRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
