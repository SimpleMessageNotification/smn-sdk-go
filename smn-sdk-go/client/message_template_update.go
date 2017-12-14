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

//the request data of update message template
type UpdateMessageTemplateRequest struct {
	*BaseRequest
	Content           string `json:"content"`
	MessageTemplateId string `json:"-"`
}

//the response data of update message template
type UpdateMessageTemplateResponse struct {
	*BaseResponse
}

// send request to update message template
func (client *SmnClient) UpdateMessageTemplate(request *UpdateMessageTemplateRequest) (response *UpdateMessageTemplateResponse, err error) {
	response = &UpdateMessageTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

// create a new update message template request struct
func (client *SmnClient) NewUpdateMessageTemplateRequest() (request *UpdateMessageTemplateRequest) {
	request = &UpdateMessageTemplateRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the update message template request
func (request *UpdateMessageTemplateRequest) GetUrl() (string, error) {

	if request.MessageTemplateId == "" {
		return "", fmt.Errorf("template id is nul")
	}

	if !util.ValidateMessageTemplateContent(request.Content) {
		return "", fmt.Errorf("template content is null or oversize 256KB")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.MessageTemplate +
		util.UrlDelimiter + request.MessageTemplateId, nil
}

// get the http method of the update message template request
func (request *UpdateMessageTemplateRequest) GetMethod() string {
	return util.PUT
}

// get the body params of the update message template request
func (request *UpdateMessageTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
