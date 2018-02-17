package kafkaconnect

import (
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("KAFKACONNECT_URL", nil),
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"kafkaconnect_connector": resourceConnector(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := &KafkaConnectClient{
		Url:        d.Get("url").(string),
		HTTPClient: &http.Client{},
	}
	return client, nil
}
