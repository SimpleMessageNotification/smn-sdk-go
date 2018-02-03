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

//the request data of list message template
type ListMessageTemplateRequest struct {
	*BaseRequest
	MessageTemplateName string `json:"message_template_name"`
	Protocol            string `json:"protocol"`
	Limit               int `json:"limit"`
	Offset              int `json:"offset"`
}

//the response data of list message template
type ListMessageTemplateResponse struct {
	*BaseResponse
	MessageTemplateCount int                   `json:"message_template_count"`
	MessageTemplates     []MessageTemplateInfo `json:"message_templates"`
}

// message template data info
type MessageTemplateInfo struct {
	MessageTemplateName string   `json:"message_template_name"`
	MessageTemplateId   string   `json:"message_template_id"`
	TagNames            []string `json:"tag_names"`
	CreateTime          string   `json:"create_time"`
	UpdateTime          string   `json:"update_time"`
	Protocol            string   `json:"protocol"`
}

// send request to list message template
func (client *SmnClient) ListMessageTemplate(request *ListMessageTemplateRequest) (response *ListMessageTemplateResponse, err error) {
	response = &ListMessageTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list message template request
func (client *SmnClient) NewListMessageTemplateRequest() (request *ListMessageTemplateRequest) {
	request = &ListMessageTemplateRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
		Offset:      0,
		Limit:       100,
	}
	return
}

// get the url of the list message template request
func (request *ListMessageTemplateRequest) GetUrl() (urlStr string, err error) {

	if request.MessageTemplateName != "" && !util.ValidateMessageTemplateName(request.MessageTemplateName) {
		return "", fmt.Errorf("template name is invalid")
	}

	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.MessageTemplate

	urlEncoded, err := util.GetQueryParams(request)
	if err != nil {
		return
	}
	if urlEncoded != "" {
		urlStr += "?" + urlEncoded
	}
	return
}

// get the http method of the list message template request
func (request *ListMessageTemplateRequest) GetMethod() string {
	return util.GET
}

// no body params
func (request *ListMessageTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
