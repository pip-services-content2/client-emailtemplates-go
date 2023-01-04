package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-emailtemplates-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type EmailTemplatesCommandableHttpClientV1 struct {
	client  *version1.EmailTemplatesCommandableHttpClientV1
	fixture *EmailTemplatesClientFixtureV1
}

func newEmailTemplatesCommandableHttpClientV1() *EmailTemplatesCommandableHttpClientV1 {
	return &EmailTemplatesCommandableHttpClientV1{}
}

func (c *EmailTemplatesCommandableHttpClientV1) setup(t *testing.T) *EmailTemplatesClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewEmailTemplatesCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewEmailTemplatesClientFixtureV1(c.client)

	return c.fixture
}

func (c *EmailTemplatesCommandableHttpClientV1) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newEmailTemplatesCommandableHttpClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
