package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type EmailTemplatesCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewEmailTemplatesCommandableHttpClientV1() *EmailTemplatesCommandableHttpClientV1 {
	return &EmailTemplatesCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/email_templates"),
	}
}

func (c *EmailTemplatesCommandableHttpClientV1) GetTemplates(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*EmailTemplateV1], error) {
	params := data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_templates", correlationId, params)
	if err != nil {
		return *data.NewEmptyDataPage[*EmailTemplateV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*EmailTemplateV1]](res, correlationId)
}

func (c *EmailTemplatesCommandableHttpClientV1) GetTemplateById(ctx context.Context, correlationId string, id string) (*EmailTemplateV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"template_id", id,
	)

	res, err := c.CallCommand(ctx, "get_template_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailTemplateV1](res, correlationId)
}

func (c *EmailTemplatesCommandableHttpClientV1) GetTemplateByIdOrName(ctx context.Context, correlationId string, idOrName string) (*EmailTemplateV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"id_or_name", idOrName,
	)

	res, err := c.CallCommand(ctx, "get_template_by_id_or_name", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailTemplateV1](res, correlationId)
}

func (c *EmailTemplatesCommandableHttpClientV1) CreateTemplate(ctx context.Context, correlationId string, template *EmailTemplateV1) (*EmailTemplateV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"template", template,
	)

	res, err := c.CallCommand(ctx, "create_template", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailTemplateV1](res, correlationId)
}

func (c *EmailTemplatesCommandableHttpClientV1) UpdateTemplate(ctx context.Context, correlationId string, template *EmailTemplateV1) (*EmailTemplateV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"template", template,
	)

	res, err := c.CallCommand(ctx, "update_template", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailTemplateV1](res, correlationId)
}

func (c *EmailTemplatesCommandableHttpClientV1) DeleteTemplateById(ctx context.Context, correlationId string, id string) (*EmailTemplateV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"template_id", id,
	)

	res, err := c.CallCommand(ctx, "delete_template_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailTemplateV1](res, correlationId)
}
