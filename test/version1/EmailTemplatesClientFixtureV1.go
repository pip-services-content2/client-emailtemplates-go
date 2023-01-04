package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-emailtemplates-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
)

type EmailTemplatesClientFixtureV1 struct {
	Client version1.IEmailTemplatesClientV1

	TEMPLATE1 *version1.EmailTemplateV1
	TEMPLATE2 *version1.EmailTemplateV1
}

func NewEmailTemplatesClientFixtureV1(client version1.IEmailTemplatesClientV1) *EmailTemplatesClientFixtureV1 {
	return &EmailTemplatesClientFixtureV1{
		Client: client,
		TEMPLATE1: &version1.EmailTemplateV1{
			Id:      "1",
			Name:    "template1",
			From:    "",
			ReplyTo: "",
			Subject: map[string]string{"en": "Text 1"},
			Text:    map[string]string{"en": "Text 1"},
			Html:    map[string]string{"en": "Text 1"},
			Status:  version1.EmailTemplateStatusCompleted,
		},
		TEMPLATE2: &version1.EmailTemplateV1{
			Id:      "2",
			Name:    "template2",
			From:    "",
			ReplyTo: "",
			Subject: map[string]string{"en": "Text 2"},
			Text:    map[string]string{"en": "Text 2"},
			Html:    map[string]string{"en": "Text 2"},
			Status:  version1.EmailTemplateStatusCompleted,
		},
	}
}

func (c *EmailTemplatesClientFixtureV1) clear() {
	c.Client.DeleteTemplateById(context.Background(), "", c.TEMPLATE1.Id)
	c.Client.DeleteTemplateById(context.Background(), "", c.TEMPLATE2.Id)
}

func (c *EmailTemplatesClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one template
	template1, err := c.Client.CreateTemplate(context.Background(), "123", c.TEMPLATE1)
	assert.Nil(t, err)

	assert.NotNil(t, template1)
	assert.Equal(t, template1.Text["en"], c.TEMPLATE1.Text["en"])
	assert.Equal(t, template1.Name, c.TEMPLATE1.Name)

	// Create another template
	template2, err := c.Client.CreateTemplate(context.Background(), "123", c.TEMPLATE2)
	assert.Nil(t, err)

	assert.NotNil(t, template2)
	assert.Equal(t, template2.Text["en"], c.TEMPLATE2.Text["en"])
	assert.Equal(t, template2.Name, c.TEMPLATE2.Name)

	// Get all templates
	page, err := c.Client.GetTemplates(context.Background(), "123", nil, data.NewPagingParams(0, 5, false))
	assert.Nil(t, err)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 2)

	// Get template by name
	template, err := c.Client.GetTemplateByIdOrName(context.Background(), "123", c.TEMPLATE1.Name)
	assert.Nil(t, err)

	assert.NotNil(t, template)

	// Update the template
	template1.Text["en"] = "Updated Content 1"

	template, err = c.Client.UpdateTemplate(context.Background(), "123", template1)
	assert.Nil(t, err)

	assert.NotNil(t, template)
	assert.Equal(t, template.Text["en"], "Updated Content 1")
	assert.Equal(t, template.Name, c.TEMPLATE1.Name)

	// Delete template
	_, err = c.Client.DeleteTemplateById(context.Background(), "123", template1.Id)
	assert.Nil(t, err)

	// Try to get delete template
	template, err = c.Client.GetTemplateById(context.Background(), "123", template1.Id)
	assert.Nil(t, err)
	assert.Nil(t, template)
}
