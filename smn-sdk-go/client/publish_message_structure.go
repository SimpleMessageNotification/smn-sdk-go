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

//the request data of publish message with structure
type PublishMessageStructureRequest struct {
	*BaseRequest
	TopicUrn         string            `json:"-"`
	Subject          string            `json:"subject"`
	MessageStructure string `json:"message_structure"`
}

//the response data of publish message with structure
type PublishMessageStructureResponse struct {
	*BaseResponse
	MessageId string `json:"message_id"`
}

// send request to publish message with structure
func (client *SmnClient) PublishMessageStructure(request *PublishMessageStructureRequest) (response *PublishMessageStructureResponse, err error) {
	response = &PublishMessageStructureResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new publish message structure request struct
func (client *SmnClient) NewPublishMessageStructureRequest() (request *PublishMessageStructureRequest) {
	request = &PublishMessageStructureRequest{
		BaseRequest:      &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the publish message structure request
func (request *PublishMessageStructureRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	if !util.ValidateSubject(request.Subject) {
		return "", fmt.Errorf("subject is invalid or oversize 512bytes")
	}

	if len(request.MessageStructure) == 0 {
		return "", fmt.Errorf("message structure is null")
	}

	if !util.ValidateMessageStructureDefault(request.MessageStructure) {
		return "", fmt.Errorf("message structure not contain defalult message")

	}

	if !util.ValidateMessageStructureLength(request.MessageStructure) {
		return "", fmt.Errorf("message structure is oversize")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn + util.UrlDelimiter + util.Publish, nil
}

// get the http method of the publish message structure request
func (request *PublishMessageStructureRequest) GetMethod() string {
	return util.POST
}

// get the body params of the publish message structure request
func (request *PublishMessageStructureRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
