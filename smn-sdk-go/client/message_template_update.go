package client

import (
	"smn-sdk-go/util"
	"fmt"
	"io"
)

type UpdateMessageTemplateRequest struct {
	*BaseRequest
	Content           string `json:"content"`
	MessageTemplateId string `json:"-"`
}

type UpdateMessageTemplateResponse struct {
	*BaseResponse
}

func (client *SmnClient) UpdateMessageTemplate(request *UpdateMessageTemplateRequest) (response *UpdateMessageTemplateResponse, err error) {
	response = &UpdateMessageTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewUpdateMessageTemplateRequest() (request *UpdateMessageTemplateRequest) {
	request = &UpdateMessageTemplateRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *UpdateMessageTemplateRequest) GetUrl() (string, error) {

	if request.MessageTemplateId == "" {
		return "", fmt.Errorf("template id is nul")
	}

	if !util.ValidateMessageTemplateContent(request.Content) {
		return "", fmt.Errorf("template content is null or oversize 256KB")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.MessageTemplate +
		util.UrlDelimiter + request.MessageTemplateId, nil
}

func (request *UpdateMessageTemplateRequest) GetMethod() string {
	return util.PUT
}

func (request *UpdateMessageTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}
