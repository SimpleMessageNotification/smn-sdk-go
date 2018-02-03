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
	"io"
)

// the request data of list sms msg report
type ListSmsMsgReportRequest struct {
	*BaseRequest
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	SignId    string `json:"sign_id"`
	Mobile    string `json:"mobile"`
	Status    string `json:"status"`
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
}

// the response data of list sms msg report
type ListSmsMsgReportResponse struct {
	*BaseResponse
	Count int                `json:"count"`
	Data  []SmsMsgReportInfo `json:"data"`
}

type SmsMsgReportInfo struct {
	MessageId   string `json:"message_id"`
	Status      int    `json:"status"`
	SignId      string `json:"sign_id"`
	StatusDesc  string `json:"status_desc"`
	FeeNum      int    `json:"fee_num"`
	ExtendCode  string `json:"extend_code"`
	NationCode  string `json:"nation_code"`
	Mobile      string `json:"mobile"`
	SubmitTime  string `json:"submit_time"`
	DeliverTime string `json:"deliver_time"`
}

// send request to list sms msg report
func (client *SmnClient) ListSmsMsgReport(request *ListSmsMsgReportRequest) (response *ListSmsMsgReportResponse, err error) {
	response = &ListSmsMsgReportResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.SendRequest(request, response)
	return
}

// create a new list sms msg report request struct
func (client *SmnClient) NewListSmsMsgReportRequest() (request *ListSmsMsgReportRequest) {
	request = &ListSmsMsgReportRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
		Offset:      0,
		Limit:       100,
	}
	return
}

// get the url of the list sms msg report request
func (request *ListSmsMsgReportRequest) GetUrl() (urlStr string, err error) {
	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.SmnProtocolSms +
		util.UrlDelimiter + util.Report
	urlEncoded, err := util.GetQueryParams(request)
	if err != nil {
		return
	}

	urlStr += "?" + urlEncoded
	return
}

// get the http method of the list sms msg report request
func (request *ListSmsMsgReportRequest) GetMethod() string {
	return util.GET
}

// no body params
func (request *ListSmsMsgReportRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
