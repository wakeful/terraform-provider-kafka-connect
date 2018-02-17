package kafkaconnect

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

var validStatusCode = map[int]bool{
	201: true,
	409: true,
}

type connectorConfig map[string]string

type connectorMessage struct {
	Name   string          `json:"name"`
	Config connectorConfig `json:"config"`
}

func resourceConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceConnectorCreate,
		Read:   resourceConnectorRead,
		Delete: resourceConnectorDelete,
		Exists: resourceConnectorExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"config": {
				Type:      schema.TypeMap,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
		},
		SchemaVersion: 1,
	}
}

func resourceConnectorCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*KafkaConnectClient)
	connectorName := d.Get("name").(string)
	config := d.Get("config").(map[string]interface{})

	resourceConnectorConfig := make(connectorConfig)

	for k, v := range config {
		switch v := v.(type) {
		case string:
			resourceConnectorConfig[k] = v
		}
	}

	payload, err := json.Marshal(connectorMessage{
		Name:   connectorName,
		Config: resourceConnectorConfig,
	})

	if err != nil {
		return err
	}

	connectorResp, err := client.Post(payload)

	if err != nil {
		return err
	}

	if !validStatusCode[connectorResp.StatusCode] {
		return fmt.Errorf("failed to create connector %s, got response code %d", connectorName, connectorResp.StatusCode)
	}

	defer connectorResp.Body.Close()

	d.SetId(connectorName)
	return resourceConnectorRead(d, m)
}

func resourceConnectorRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*KafkaConnectClient)
	response, err := client.Get(d.Id())

	if err != nil {
		return err
	}

	if response.StatusCode == 200 {

		var connector connectorMessage

		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&connector)
		if err != nil {
			return err
		}
		d.SetId(connector.Name)
		d.Set("name", connector.Name)
		d.Set("config", connector.Config)
	}

	return nil
}

func resourceConnectorDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*KafkaConnectClient)

	_, err := client.Delete(d.Get("name").(string))
	if err != nil {
		return err
	}

	return nil
}

func resourceConnectorExists(d *schema.ResourceData, m interface{}) (bool, error) {

	client := m.(*KafkaConnectClient)
	response, err := client.Get(d.Id())

	if err != nil {
		return false, err
	}

	if response.StatusCode != 200 {
		return false, nil
	}

	return true, nil
}
