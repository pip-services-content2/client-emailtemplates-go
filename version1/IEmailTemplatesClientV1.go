package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IEmailTemplatesClientV1 interface {
	GetTemplates(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*EmailTemplateV1], error)

	GetTemplateById(ctx context.Context, correlationId string, id string) (*EmailTemplateV1, error)

	GetTemplateByIdOrName(ctx context.Context, correlationId string, idOrName string) (*EmailTemplateV1, error)

	CreateTemplate(ctx context.Context, correlationId string, template *EmailTemplateV1) (*EmailTemplateV1, error)

	UpdateTemplate(ctx context.Context, correlationId string, template *EmailTemplateV1) (*EmailTemplateV1, error)

	DeleteTemplateById(ctx context.Context, correlationId string, id string) (*EmailTemplateV1, error)
}
