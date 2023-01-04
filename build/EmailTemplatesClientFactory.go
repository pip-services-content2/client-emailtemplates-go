package build

import (
	clients1 "github.com/pip-services-content2/client-emailtemplates-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type EmailTemplatesClientFactory struct {
	*cbuild.Factory
}

func NewEmailTemplatesClientFactory() *EmailTemplatesClientFactory {
	c := &EmailTemplatesClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-emailtemplates", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-emailtemplates", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-emailtemplates", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewEmailTemplatesNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewEmailTemplatesMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewEmailTemplatesCommandableHttpClientV1)

	return c
}
