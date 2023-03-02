// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_sesv2_dedicated_ip_pool": DataSourceDedicatedIPPool,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_sesv2_configuration_set":                   ResourceConfigurationSet,
		"aws_sesv2_configuration_set_event_destination": ResourceConfigurationSetEventDestination,
		"aws_sesv2_dedicated_ip_assignment":             ResourceDedicatedIPAssignment,
		"aws_sesv2_dedicated_ip_pool":                   ResourceDedicatedIPPool,
		"aws_sesv2_email_identity":                      ResourceEmailIdentity,
		"aws_sesv2_email_identity_feedback_attributes":  ResourceEmailIdentityFeedbackAttributes,
		"aws_sesv2_email_identity_mail_from_attributes": ResourceEmailIdentityMailFromAttributes,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SESV2
}

var ServicePackage = &servicePackage{}
