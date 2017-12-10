package client

import (
	"smn-sdk-go/util"
	"fmt"
	"io"
)

type CreateMessageTemplateRequest struct {
	*BaseRequest
	MessageTemplateName string `json:"message_template_name"`
	Content             string `json:"content"`
	Protocol            string `json:"protocol"`
}

type CreateMessageTemplateResponse struct {
	*BaseResponse
	MessageTemplateId string `json:"message_template_id"`
}

func (client *SmnClient) CreateMessageTemplate(request *CreateMessageTemplateRequest) (response *CreateMessageTemplateResponse, err error) {
	response = &CreateMessageTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewCreateMessageTemplateRequest() (request *CreateMessageTemplateRequest) {
	request = &CreateMessageTemplateRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

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

func (request *CreateMessageTemplateRequest) GetMethod() string {
	return util.POST
}

func (request *CreateMessageTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
