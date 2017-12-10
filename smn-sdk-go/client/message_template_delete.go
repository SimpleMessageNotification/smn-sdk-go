package client

import (
	"smn-sdk-go/util"
	"fmt"
	"io"
)

type DeleteMessageTemplateRequest struct {
	*BaseRequest
	MessageTemplateId string
}

type DeleteMessageTemplateResponse struct {
	*BaseResponse
}

func (client *SmnClient) DeleteMessageTemplate(request *DeleteMessageTemplateRequest) (response *DeleteMessageTemplateResponse, err error) {
	response = &DeleteMessageTemplateResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewDeleteMessageTemplateRequest() (request *DeleteMessageTemplateRequest) {
	request = &DeleteMessageTemplateRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *DeleteMessageTemplateRequest) GetUrl() (string, error) {

	if request.MessageTemplateId == "" {
		return "", fmt.Errorf("template id is nul")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.MessageTemplate +
		util.UrlDelimiter + request.MessageTemplateId, nil
}

func (request *DeleteMessageTemplateRequest) GetMethod() string {
	return util.DELETE
}

func (request *DeleteMessageTemplateRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
