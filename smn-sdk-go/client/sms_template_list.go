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
	"github.com/SimpleMessageNotification/smn-sdk-go/smn-sdk-go/util"
	"io"
)

//the request data of list sms template
type ListSmsTemplatesRequest struct {
	*BaseRequest
	SmsTemplateName string `json:"sms_template_name"`
	SmsTemplateType int `json:"sms_template_type"`
	Limit           string `json:"limit"`
	Offset          string `json:"offset"`
}

//the response data of list sms template
type ListSmsTemplatesResponse struct {
	*BaseResponse
	Count        int                   `json:"count"`
	SmsTemplates []SmsTemplate `json:"sms_templates"`
}

// message template data info
type SmsTemplate struct {
	SmsTemplateName string   `json:"sms_template_name"`
	SmsTemplateType int   `json:"sms_template_type"`
	SmsTemplateId   string   `json:"sms_template_id"`
	Reply           string `json:"reply"`
	Status          int `json:"status"`
	CreateTime      string   `json:"create_time"`
	ValidityEndTime string   `json:"validity_end_time"`
}

// send request to list message template
func (client *SmnClient) ListSmsTemplates(request *ListSmsTemplatesRequest) (response *ListSmsTemplatesResponse, err error) {
	response = &ListSmsTemplatesResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list sms template request
func (client *SmnClient) NewListSmsTemplatesRequest() (request *ListSmsTemplatesRequest) {
	request = &ListSmsTemplatesRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the list sms template request
func (request *ListSmsTemplatesRequest) GetUrl() (urlStr string, err error) {
	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmnSmsTemplate

	urlEncoded, err := util.GetQueryParams(request)
	if err != nil {
		return
	}
	if urlEncoded != "" {
		urlStr += "?" + urlEncoded
	}
	return
}

// get the http method of the list sms template request
func (request *ListSmsTemplatesRequest) GetMethod() string {
	return util.GET
}

// no body params
func (request *ListSmsTemplatesRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
