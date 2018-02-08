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
	"fmt"
)

//the request data of get sms template detail
type GetSmsTemplateDetailRequest struct {
	*BaseRequest
	SmsTemplateId string `json:"-"`
}

//the response data of get sms template detail
type GetSmsTemplateDetailResponse struct {
	*BaseResponse
	SmsTemplateName    string   `json:"sms_template_name"`
	SmsTemplateType    int   `json:"sms_template_type"`
	SmsTemplateId      string   `json:"sms_template_id"`
	SmsTemplateContent string   `json:"sms_template_content"`
	Reply              string `json:"reply"`
	Status             int `json:"status"`
	CreateTime         string   `json:"create_time"`
	ValidityEndTime    string   `json:"validity_end_time"`
}

// send request to get sms template detail
func (client *SmnClient) GetSmsTemplateDetailRequest(request *GetSmsTemplateDetailRequest) (response *GetSmsTemplateDetailResponse, err error) {
	response = &GetSmsTemplateDetailResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list sms template request
func (client *SmnClient) NewGetSmsTemplateDetailRequest() (request *GetSmsTemplateDetailRequest) {
	request = &GetSmsTemplateDetailRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of get sms template detail request
func (request *GetSmsTemplateDetailRequest) GetUrl() (urlStr string, err error) {
	if request.SmsTemplateId == "" {
		return "", fmt.Errorf("sms template id is null")
	}

	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmnSmsTemplate +
		util.UrlDelimiter + request.SmsTemplateId

	urlEncoded, err := util.GetQueryParams(request)
	if err != nil {
		return
	}
	if urlEncoded != "" {
		urlStr += "?" + urlEncoded
	}
	return
}

// get the http method of get sms template detail request
func (request *GetSmsTemplateDetailRequest) GetMethod() string {
	return util.GET
}

// no body params
func (request *GetSmsTemplateDetailRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
