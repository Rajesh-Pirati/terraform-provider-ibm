// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/scc-go-sdk/findingsv1"
)

func dataSourceIBMSccSiProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSccSiProviderRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the provider.",
			},
			"providers": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The providers requested.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the provider in the form '{account_id}/providers/{provider_id}'.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the provider.",
						},
					},
				},
			},
			"limit": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of elements returned in the current instance. The default is 200.",
			},
			"skip": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The offset is the index of the item from which you want to start returning data from. The default is 0.",
			},
			"total_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of providers available.",
			},
		},
	}
}

func dataSourceIBMSccSiProviderRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	findingsClient, err := meta.(ClientSession).FindingsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	listProvidersOptions := &findingsv1.ListProvidersOptions{}

	apiListProvidersResponse, response, err := findingsClient.ListProvidersWithContext(context, listProvidersOptions)
	if err != nil {
		log.Printf("[DEBUG] ListProvidersWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListProvidersWithContext failed %s\n%s", err, response))
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchProviders []findingsv1.APIProvider
	var id string
	var suppliedFilter bool

	if v, ok := d.GetOk("id"); ok {
		id = v.(string)
		suppliedFilter = true
		for _, data := range apiListProvidersResponse.Providers {
			if *data.ID == id {
				matchProviders = append(matchProviders, data)
			}
		}
	} else {
		matchProviders = apiListProvidersResponse.Providers
	}
	apiListProvidersResponse.Providers = matchProviders

	if suppliedFilter {
		if len(apiListProvidersResponse.Providers) == 0 {
			return diag.FromErr(fmt.Errorf("no Providers found with id %s", id))
		}
		d.SetId(id)
	} else {
		d.SetId(dataSourceIBMSccSiProviderID(d))
	}

	if apiListProvidersResponse.Providers != nil {
		err = d.Set("providers", dataSourceAPIListProvidersResponseFlattenProviders(apiListProvidersResponse.Providers))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting providers %s", err))
		}
	}
	if err = d.Set("limit", intValue(apiListProvidersResponse.Limit)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting limit: %s", err))
	}
	if err = d.Set("skip", intValue(apiListProvidersResponse.Skip)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting skip: %s", err))
	}
	if err = d.Set("total_count", intValue(apiListProvidersResponse.TotalCount)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting total_count: %s", err))
	}

	return nil
}

// dataSourceIBMSccSiProviderID returns a reasonable ID for the list.
func dataSourceIBMSccSiProviderID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceAPIListProvidersResponseFlattenProviders(result []findingsv1.APIProvider) (providers []map[string]interface{}) {
	for _, providersItem := range result {
		providers = append(providers, dataSourceAPIListProvidersResponseProvidersToMap(providersItem))
	}

	return providers
}

func dataSourceAPIListProvidersResponseProvidersToMap(providersItem findingsv1.APIProvider) (providersMap map[string]interface{}) {
	providersMap = map[string]interface{}{}

	if providersItem.Name != nil {
		providersMap["name"] = providersItem.Name
	}
	if providersItem.ID != nil {
		providersMap["id"] = providersItem.ID
	}

	return providersMap
}
