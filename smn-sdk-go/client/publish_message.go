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
	"smn-sdk-go/util"
	"io"
)

type PublishMessageRequest struct {
	*BaseRequest
	TopicUrn            string            `json:"-"`
	Subject             string            `json:"subject"`
	Message             string            `json:"message"`
}
type PublishMessageResponse struct {
	*BaseResponse
	MessageId string `json:"message_id"`
}

func (client *SmnClient) PublishMessage(request *PublishMessageRequest) (response *PublishMessageResponse, err error) {
	response = &PublishMessageResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewPublishMessageRequest() (request *PublishMessageRequest) {
	request = &PublishMessageRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *PublishMessageRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	if !util.ValidateSubject(request.Subject) {
		return "", fmt.Errorf("subject is invalid or oversize 512bytes")
	}

	if !util.ValidateMessage(request.Message) {
		return "", fmt.Errorf("message is null or oversize 256KB")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn + util.UrlDelimiter + util.Publish, nil
}

func (request *PublishMessageRequest) GetMethod() string {
	return util.POST
}

func (request *PublishMessageRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
