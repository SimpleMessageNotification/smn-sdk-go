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
	"fmt"
	"io"
)

// the request data of create message template
type CreateMessageTemplateRequest struct {
	*BaseRequest
	MessageTemplateName string `json:"message_template_name"`
	Content             string `json:"content"`
	Protocol            string `json:"protocol"`
}

// the response data of create message template
type CreateMessageTemplateResponse struct {
	*BaseResponse
	MessageTemplateId string `json:"message_template_id"`
}

// send request to create message template
func (client *SmnClient) CreateMessageTemplate(request *CreateMessageTemplateRequest) (response *CreateMessageTemplateResponse, err error) {
	response = &CreateMessageTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

// create a new create message template request struct
func (client *SmnClient) NewCreateMessageTemplateRequest() (request *CreateMessageTemplateRequest) {
	request = &CreateMessageTemplateRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the create message template request
func (request *CreateMessageTemplateRequest) GetUrl() (string, error) {
	if request.MessageTemplateName == "" {
		return "", fmt.Errorf("template name is null")

	}
	if !util.ValidateMessageTemplateContent(request.Content) {
		return "", fmt.Errorf("template content is null or oversize 256KB")
	}
	if !util.ValidateMessageTemplateName(request.MessageTemplateName) {
		return "", fmt.Errorf("template name is invalid")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.MessageTemplate, nil
}

// get the http method of create message template request
func (request *CreateMessageTemplateRequest) GetMethod() string {
	return util.POST
}

// get the body params of the create message template request
func (request *CreateMessageTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
