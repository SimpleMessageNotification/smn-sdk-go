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

// the request data of publish promotion sms
type PromotionSmsPublishRequest struct {
	*BaseRequest
	SmsTemplateId string `json:"sms_template_id""`
	EndPoints     []string `json:"endpoints"`
	SignId        string `json:"sign_id"`
}

// the response data of publish promotion sms
type PromotionSmsPublishResponse struct {
	*BaseResponse
	Result []PromotionSmsPublishResult `json:"result""`
}

// publish promotion sms result data info
type PromotionSmsPublishResult struct {
	MessageId string `json:"message_id"`
	Endpoint  string `json:"endpoint"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

// send request to publish sms
func (client *SmnClient) PromotionSmsPublish(request *PromotionSmsPublishRequest) (response *PromotionSmsPublishResponse, err error) {
	response = &PromotionSmsPublishResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list sms publish request struct
func (client *SmnClient) NewPromotionSmsPublishRequest() (request *PromotionSmsPublishRequest) {
	request = &PromotionSmsPublishRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of promotion sms publish request
func (request *PromotionSmsPublishRequest) GetUrl() (string, error) {
	if request.SmsTemplateId == "" {
		return "", fmt.Errorf("sms template id is null")
	}

	if request.SignId == "" {
		return "", fmt.Errorf("signId is null")
	}

	if request.EndPoints == nil || len(request.EndPoints) == 0 {
		return "", fmt.Errorf("endpoints is invalid")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter +
		util.SmnProtocolSms + util.UrlDelimiter + util.SmnPromotionSms, nil
}

// get the http method of the promotion sms publish request
func (request *PromotionSmsPublishRequest) GetMethod() string {
	return util.POST
}

// get the body params of the promotion sms publish request
func (request *PromotionSmsPublishRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
