package version1

import (
	"context"
	"strings"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type EmailTemplatesMockClientV1 struct {
	templates []*EmailTemplateV1
}

func NewEmailTemplatesMockClientV1() *EmailTemplatesMockClientV1 {
	return &EmailTemplatesMockClientV1{
		templates: make([]*EmailTemplateV1, 0),
	}
}

func (c *EmailTemplatesMockClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(search))
}

func (c *EmailTemplatesMockClientV1) matchMultilanguageString(value map[string]string, search string) bool {
	for text, _ := range value {
		if c.matchString(text, search) {
			return true
		}
	}

	return false
}

func (c *EmailTemplatesMockClientV1) matchSearch(item *EmailTemplateV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchString(item.Name, search) {
		return true
	}
	if c.matchMultilanguageString(item.Subject, search) {
		return true
	}
	if c.matchMultilanguageString(item.Text, search) {
		return true
	}
	if c.matchMultilanguageString(item.Html, search) {
		return true
	}
	if c.matchString(item.Status, search) {
		return true
	}
	return false
}

func (c *EmailTemplatesMockClientV1) composeFilter(filter *data.FilterParams) func(item *EmailTemplateV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	search, searchOk := filter.GetAsNullableString("search")
	id, idOk := filter.GetAsNullableString("id")
	status, statusOk := filter.GetAsNullableString("status")
	name, nameOk := filter.GetAsNullableString("name")

	return func(item *EmailTemplateV1) bool {
		if idOk && id != item.Id {
			return false
		}
		if nameOk && name != item.Name {
			return false
		}
		if statusOk && status != item.Status {
			return false
		}
		if searchOk && !c.matchSearch(item, search) {
			return false
		}

		return true
	}
}

func (c *EmailTemplatesMockClientV1) GetTemplates(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*EmailTemplateV1], error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*EmailTemplateV1, 0)
	for _, v := range c.templates {
		item := *v
		if filterFunc(&item) {
			items = append(items, &item)
		}
	}
	return *data.NewDataPage(items, len(c.templates)), nil
}

func (c *EmailTemplatesMockClientV1) GetTemplateById(ctx context.Context, correlationId string, id string) (*EmailTemplateV1, error) {
	var item *EmailTemplateV1

	for _, template := range c.templates {
		if template.Id == id {
			buf := *template
			item = &buf
			break
		}
	}

	return item, nil
}

func (c *EmailTemplatesMockClientV1) GetTemplateByIdOrName(ctx context.Context, correlationId string, idOrName string) (*EmailTemplateV1, error) {
	var item *EmailTemplateV1

	for _, template := range c.templates {
		if template.Id == idOrName || template.Name == idOrName {
			buf := *template
			item = &buf
			break
		}
	}

	return item, nil
}

func (c *EmailTemplatesMockClientV1) CreateTemplate(ctx context.Context, correlationId string, template *EmailTemplateV1) (*EmailTemplateV1, error) {
	buf := *template
	c.templates = append(c.templates, &buf)
	return template, nil
}

func (c *EmailTemplatesMockClientV1) UpdateTemplate(ctx context.Context, correlationId string, template *EmailTemplateV1) (*EmailTemplateV1, error) {
	var item *EmailTemplateV1

	for _, el := range c.templates {
		if template.Id == el.Id {
			buf := *template
			item = &buf
			break
		}
	}

	return item, nil
}

func (c *EmailTemplatesMockClientV1) DeleteTemplateById(ctx context.Context, correlationId string, id string) (*EmailTemplateV1, error) {
	var item *EmailTemplateV1

	for index, template := range c.templates {
		if template.Id == id {
			buf := *template
			item = &buf
			if index < len(c.templates) {
				c.templates = append(c.templates[:index], c.templates[index+1:]...)
			} else {
				c.templates = c.templates[:index]
			}
			break
		}
	}

	return item, nil
}
