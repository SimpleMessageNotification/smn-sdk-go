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

// the request data of update topic
type UpdateTopicRequest struct {
	*BaseRequest
	TopicUrn    string `json:"-"`
	DisplayName string `json:"display_name"`
}

// the response data of update topic
type UpdateTopicResponse struct {
	*BaseResponse
}

// send request to update topic
func (client *SmnClient) UpdateTopic(request *UpdateTopicRequest) (response *UpdateTopicResponse, err error) {
	response = &UpdateTopicResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

// create a new update topic request struct
func (client *SmnClient) NewUpdateTopicRequest() (request *UpdateTopicRequest) {
	request = &UpdateTopicRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the update topic request
func (request *UpdateTopicRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	if !util.ValidateTopicDiplayName(request.DisplayName) {
		return "", fmt.Errorf("topic displayName is overSize")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn, nil
}

// get the http method of the update topic request
func (request *UpdateTopicRequest) GetMethod() string {
	return util.PUT
}

// get the body params of the update topic request
func (request *UpdateTopicRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
