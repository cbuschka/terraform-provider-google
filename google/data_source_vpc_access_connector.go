package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVPCAccessConnector() *schema.Resource {
	dsSchema := datasourceSchemaFromResourceSchema(resourceSqlDatabaseInstance().Schema)
	addRequiredFieldsToSchema(dsSchema, "name")
	addOptionalFieldsToSchema(dsSchema, "project")
	addOptionalFieldsToSchema(dsSchema, "region")

	return &schema.Resource{
		Read:   dataSourceSqlDatabaseInstanceRead,
		Schema: dsSchema,
	}
}

func dataSourceVPCAccessConnectorRead(d *schema.ResourceData, meta interface{}) error {

	return resourceVPCAccessConnectorRead(d, meta)

}
