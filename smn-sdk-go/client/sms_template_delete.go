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

//the request data of delete sms template
type DeleteSmsTemplateRequest struct {
	*BaseRequest
	SmsTemplateId string `json:"-"`
}

//the response data of delete sms template
type DeleteSmsTemplateResponse struct {
	*BaseResponse
}

// send request to delete sms template
func (client *SmnClient) DeleteSmsTemplate(request *DeleteSmsTemplateRequest) (response *DeleteSmsTemplateResponse, err error) {
	response = &DeleteSmsTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list sms template request
func (client *SmnClient) NewDeleteSmsTemplateRequest() (request *DeleteSmsTemplateRequest) {
	request = &DeleteSmsTemplateRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the delete sms template request
func (request *DeleteSmsTemplateRequest) GetUrl() (urlStr string, err error) {
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

// get the http method of the delete sms template request
func (request *DeleteSmsTemplateRequest) GetMethod() string {
	return util.DELETE
}

// no body params
func (request *DeleteSmsTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
