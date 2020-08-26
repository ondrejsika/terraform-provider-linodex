package linodex

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodexInstanceIPs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodexInstanceIPsRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"public": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"private": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"reserved": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"shares": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func dataSourceLinodexInstanceIPsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	reqInstance, err := strconv.Atoi(d.Get("id").(string))

	if err != nil {
		return fmt.Errorf("Instance ID is not a number: %s", err)
	}

	instanceNetwork, err := client.GetInstanceIPAddresses(context.Background(), reqInstance)

	if err != nil {
		return fmt.Errorf("Error getting the IPs for Linode instance %s: %s", d.Id(), err)
	}

	public := make([]string, len(instanceNetwork.IPv4.Public))
	for i := 0; i < len(instanceNetwork.IPv4.Public); i++ {
		public[i] = instanceNetwork.IPv4.Public[i].Address
	}
	d.Set("public", public)

	private := make([]string, len(instanceNetwork.IPv4.Public))
	for i := 0; i < len(instanceNetwork.IPv4.Private); i++ {
		private[i] = instanceNetwork.IPv4.Private[i].Address
	}
	d.Set("private", private)

	shared := make([]string, len(instanceNetwork.IPv4.Shared))
	for i := 0; i < len(instanceNetwork.IPv4.Shared); i++ {
		shared[i] = instanceNetwork.IPv4.Shared[i].Address
	}
	d.Set("shared", shared)

	reserved := make([]string, len(instanceNetwork.IPv4.Reserved))
	for i := 0; i < len(instanceNetwork.IPv4.Reserved); i++ {
		reserved[i] = instanceNetwork.IPv4.Reserved[i].Address
	}
	d.Set("reserved", reserved)

	if instanceNetwork != nil {
		d.SetId(d.Get("id").(string))
		return nil
	}

	d.SetId("")
	return fmt.Errorf("Instance %d was not found", reqInstance)
}
