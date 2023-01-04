package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type EmailTemplatesNullClientV1 struct {
}

func NewEmailTemplatesNullClientV1() *EmailTemplatesNullClientV1 {
	return &EmailTemplatesNullClientV1{}
}

func (c *EmailTemplatesNullClientV1) GetTemplates(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*EmailTemplateV1], error) {
	return *data.NewEmptyDataPage[*EmailTemplateV1](), nil
}

func (c *EmailTemplatesNullClientV1) GetTemplateById(ctx context.Context, correlationId string, id string) (*EmailTemplateV1, error) {
	return nil, nil
}

func (c *EmailTemplatesNullClientV1) GetTemplateByIdOrName(ctx context.Context, correlationId string, idOrName string) (*EmailTemplateV1, error) {
	return nil, nil
}

func (c *EmailTemplatesNullClientV1) CreateTemplate(ctx context.Context, correlationId string, template *EmailTemplateV1) (*EmailTemplateV1, error) {
	return template, nil
}

func (c *EmailTemplatesNullClientV1) UpdateTemplate(ctx context.Context, correlationId string, template *EmailTemplateV1) (*EmailTemplateV1, error) {
	return template, nil
}

func (c *EmailTemplatesNullClientV1) DeleteTemplateById(ctx context.Context, correlationId string, id string) (*EmailTemplateV1, error) {
	return nil, nil
}
