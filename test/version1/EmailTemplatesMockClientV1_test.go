package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-emailtemplates-go/version1"
)

type EmailTemplatesMockClientV1 struct {
	client  *version1.EmailTemplatesMockClientV1
	fixture *EmailTemplatesClientFixtureV1
}

func newEmailTemplatesMockClientV1() *EmailTemplatesMockClientV1 {
	return &EmailTemplatesMockClientV1{}
}

func (c *EmailTemplatesMockClientV1) setup(t *testing.T) *EmailTemplatesClientFixtureV1 {
	c.client = version1.NewEmailTemplatesMockClientV1()
	c.fixture = NewEmailTemplatesClientFixtureV1(c.client)
	return c.fixture
}

func (c *EmailTemplatesMockClientV1) teardown(t *testing.T) {
	c.client = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newEmailTemplatesMockClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
