package client

import (
	"smn-sdk-go/util"
	"fmt"
	"io"
)

type QueryMessageTemplateDetailRequest struct {
	*BaseRequest
	MessageTemplateId string
}

type QueryMessageTemplateDetailResponse struct {
	*BaseResponse
	MessageTemplateName string   `json:"message_template_name"`
	MessageTemplateId   string   `json:"message_template_id"`
	TagNames            []string `json:"tag_names"`
	CreateTime          string   `json:"create_time"`
	UpdateTime          string   `json:"update_time"`
	Content             string   `json:"content"`
}

func (client *SmnClient) QueryMessageTemplateDetail(request *QueryMessageTemplateDetailRequest) (response *QueryMessageTemplateDetailResponse, err error) {
	response = &QueryMessageTemplateDetailResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewQueryMessageTemplateDetailRequest() (request *QueryMessageTemplateDetailRequest) {
	request = &QueryMessageTemplateDetailRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *QueryMessageTemplateDetailRequest) GetUrl() (string, error) {

	if request.MessageTemplateId == "" {
		return "", fmt.Errorf("template id is nul")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.MessageTemplate +
		util.UrlDelimiter + request.MessageTemplateId, nil
}

func (request *QueryMessageTemplateDetailRequest) GetMethod() string {
	return util.GET
}

func (request *QueryMessageTemplateDetailRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
