/*
 * Copyright (C) 2018. Huawei Technologies Co., LTD. All rights reserved.
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

// the request data of create sms template
type CreateSmsTemplateRequest struct {
	*BaseRequest
	SmsTemplateName    string `json:"sms_template_name""`
	SmsTemplateContent string `json:"sms_template_content"`
	Remark             string `json:"remark"`
	SmsTemplateType    int `json:"sms_template_type"`
}

// the response data of create sms template
type CreateSmsTemplateResponse struct {
	*BaseResponse
	SmsTemplateId string `json:"sms_template_id""`
}

// send request to create sms template
func (client *SmnClient) CreateSmsTemplate(request *CreateSmsTemplateRequest) (response *CreateSmsTemplateResponse, err error) {
	response = &CreateSmsTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new create sms template request struct
func (client *SmnClient) NewCreateSmsTemplateRequest() (request *CreateSmsTemplateRequest) {
	request = &CreateSmsTemplateRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of create sms template request
func (request *CreateSmsTemplateRequest) GetUrl() (string, error) {
	if request.SmsTemplateName == "" {
		return "", fmt.Errorf("sms template name is null")
	}

	if request.SmsTemplateContent == "" {
		return "", fmt.Errorf("sms template content is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter +
		util.SmnSmsTemplate, nil
}

// get the http method of the promotion sms publish request
func (request *CreateSmsTemplateRequest) GetMethod() string {
	return util.POST
}

// get the body params of the promotion sms publish request
func (request *CreateSmsTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
