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

//the request data of query message template detail
type QueryMessageTemplateDetailRequest struct {
	*BaseRequest
	MessageTemplateId string
}

//the response data of query message template detail
type QueryMessageTemplateDetailResponse struct {
	*BaseResponse
	MessageTemplateName string   `json:"message_template_name"`
	MessageTemplateId   string   `json:"message_template_id"`
	TagNames            []string `json:"tag_names"`
	CreateTime          string   `json:"create_time"`
	UpdateTime          string   `json:"update_time"`
	Content             string   `json:"content"`
}

// send request to query message template detail
func (client *SmnClient) QueryMessageTemplateDetail(request *QueryMessageTemplateDetailRequest) (response *QueryMessageTemplateDetailResponse, err error) {
	response = &QueryMessageTemplateDetailResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new query message template detail request struct
func (client *SmnClient) NewQueryMessageTemplateDetailRequest() (request *QueryMessageTemplateDetailRequest) {
	request = &QueryMessageTemplateDetailRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

// get the url of the query message template detail request
func (request *QueryMessageTemplateDetailRequest) GetUrl() (string, error) {

	if request.MessageTemplateId == "" {
		return "", fmt.Errorf("template id is nul")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.MessageTemplate +
		util.UrlDelimiter + request.MessageTemplateId, nil
}

// get the http method of the query message template detail request
func (request *QueryMessageTemplateDetailRequest) GetMethod() string {
	return util.GET
}

// no body params
func (request *QueryMessageTemplateDetailRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
