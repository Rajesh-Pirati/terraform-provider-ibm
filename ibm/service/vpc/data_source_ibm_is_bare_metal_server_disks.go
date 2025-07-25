// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerID               = "bare_metal_server"
	isBareMetalServerDiskHref         = "href"
	isBareMetalServerDiskResourceType = "resource_type"
)

func DataSourceIBMIsBareMetalServerDisks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISBareMetalServerDisksRead,

		Schema: map[string]*schema.Schema{
			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The bare metal server identifier",
			},

			//disks

			isBareMetalServerDisks: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of bare metal server disks. Disk is a block device that is locally attached to the physical server. By default, the listed disks are sorted by their created_at property values, with the newest disk first.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerDiskHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this bare metal server disk",
						},
						isBareMetalServerDiskID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this bare metal server disk",
						},
						isBareMetalServerDiskInterfaceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The disk interface used for attaching the disk. Supported values are [ nvme, sata ]",
						},
						isBareMetalServerDiskName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this disk",
						},
						isBareMetalServerDiskResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type",
						},
						isBareMetalServerDiskSize: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the disk in GB (gigabytes)",
						},
						"allowed_use": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The usage constraints to be matched against the requested bare metal server properties to determine compatibility.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bare_metal_server": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this disk.",
									},
									"api_version": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The API version with which to evaluate the expressions.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISBareMetalServerDisksRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	bareMetalServerID := d.Get(isBareMetalServerID).(string)
	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_disks", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.ListBareMetalServerDisksOptions{
		BareMetalServerID: &bareMetalServerID,
	}

	diskCollection, _, err := sess.ListBareMetalServerDisksWithContext(context, options)
	disks := diskCollection.Disks
	if err != nil || disks == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListBareMetalServerDisksWithContext failed: %s", err.Error()), "(Data) ibm_is_bare_metal_server_disks", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	disksInfo := make([]map[string]interface{}, 0)
	for _, disk := range disks {
		l := map[string]interface{}{
			isBareMetalServerDiskHref:          disk.Href,
			isBareMetalServerDiskID:            disk.ID,
			isBareMetalServerDiskInterfaceType: disk.InterfaceType,
			isBareMetalServerDiskName:          disk.Name,
			isBareMetalServerDiskResourceType:  disk.ResourceType,
			isBareMetalServerDiskSize:          disk.Size,
		}
		if disk.AllowedUse != nil {
			usageConstraintList := []map[string]interface{}{}
			modelMap, err := ResourceceIBMIsBareMetalServerDiskAllowedUseToMap(disk.AllowedUse)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_use: %s", err), "(Data) ibm_is_bare_metal_server_disks", "read", "set-allowed_use").GetDiag()
			}
			usageConstraintList = append(usageConstraintList, modelMap)
			l["allowed_use"] = usageConstraintList
		}
		disksInfo = append(disksInfo, l)
	}

	d.SetId(dataSourceIBMISBMSDisksID(d))
	if err = d.Set(isBareMetalServerDisks, disksInfo); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting disks: %s", err), "(Data) ibm_is_bare_metal_server_disks", "read", "set-disks").GetDiag()
	}
	return nil
}

// dataSourceIBMISBMSProfilesID returns a reasonable ID for a Bare Metal Server Disks list.
func dataSourceIBMISBMSDisksID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
