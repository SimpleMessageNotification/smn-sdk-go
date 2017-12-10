package client

import (
	"smn-sdk-go/util"
	"fmt"
	"io"
)

type ListMessageTemplateRequest struct {
	*BaseRequest
	MessageTemplateName string `json:"message_template_name"`
	Protocol            string `json:"protocol"`
	Limit               string `json:"limit"`
	Offset              string `json:"offset"`
}

type ListMessageTemplateResponse struct {
	*BaseResponse
	MessageTemplateCount int                   `json:"message_template_count"`
	MessageTemplates     []MessageTemplateInfo `json:"message_templates"`
}
type MessageTemplateInfo struct {
	MessageTemplateName string   `json:"message_template_name"`
	MessageTemplateId   string   `json:"message_template_id"`
	TagNames            []string `json:"tag_names"`
	CreateTime          string   `json:"create_time"`
	UpdateTime          string   `json:"update_time"`
	Protocol            string   `json:"protocol"`
}

func (client *SmnClient) ListMessageTemplate(request *ListMessageTemplateRequest) (response *ListMessageTemplateResponse, err error) {
	response = &ListMessageTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewListMessageTemplateRequest() (request *ListMessageTemplateRequest) {
	request = &ListMessageTemplateRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

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

func (request *ListMessageTemplateRequest) GetMethod() string {
	return util.GET
}

func (request *ListMessageTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
