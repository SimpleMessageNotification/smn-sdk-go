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
	"fmt"
	"io"
)

//the request data of delete message template
type DeleteMessageTemplateRequest struct {
	*BaseRequest
	MessageTemplateId string
}

// the response data of delete message template
type DeleteMessageTemplateResponse struct {
	*BaseResponse
}

// send request to delete message template
func (client *SmnClient) DeleteMessageTemplate(request *DeleteMessageTemplateRequest) (response *DeleteMessageTemplateResponse, err error) {
	response = &DeleteMessageTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

// create a new delete message template request
func (client *SmnClient) NewDeleteMessageTemplateRequest() (request *DeleteMessageTemplateRequest) {
	request = &DeleteMessageTemplateRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the delete message template request
func (request *DeleteMessageTemplateRequest) GetUrl() (string, error) {

	if request.MessageTemplateId == "" {
		return "", fmt.Errorf("template id is nul")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.MessageTemplate +
		util.UrlDelimiter + request.MessageTemplateId, nil
}

// get the http method of the delete message template request
func (request *DeleteMessageTemplateRequest) GetMethod() string {
	return util.DELETE
}

// no body params
func (request *DeleteMessageTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
