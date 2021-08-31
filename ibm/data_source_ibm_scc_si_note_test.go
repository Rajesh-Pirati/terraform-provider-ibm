// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccSiNoteDataSourceBasic(t *testing.T) {
	apiNoteProviderID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	apiNoteShortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	apiNoteLongDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	apiNoteKind := "FINDING"
	apiNoteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiNoteDataSourceConfigBasic(apiNoteProviderID, apiNoteShortDescription, apiNoteLongDescription, apiNoteKind, apiNoteID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "provider_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "note_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "kind"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "reported_by.#"),
				),
			},
		},
	})
}

func TestAccIBMSccSiNoteDataSourceAllArgs(t *testing.T) {
	apiNoteProviderID := fmt.Sprintf("tf_provider_id_%d", acctest.RandIntRange(10, 100))
	apiNoteShortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	apiNoteLongDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	apiNoteKind := "FINDING"
	apiNoteID := fmt.Sprintf("tf_note_id_%d", acctest.RandIntRange(10, 100))
	apiNoteShared := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccSiNoteDataSourceConfig(apiNoteProviderID, apiNoteShortDescription, apiNoteLongDescription, apiNoteKind, apiNoteID, apiNoteShared),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "provider_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "note_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "kind"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "related_url.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "related_url.0.label"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "related_url.0.url"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "expiration_time"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "create_time"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "update_time"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "shared"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "reported_by.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "finding.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "kpi.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "card.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_si_note.scc_si_note", "section.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccSiNoteDataSourceConfigBasic(apiNoteProviderID string, apiNoteShortDescription string, apiNoteLongDescription string, apiNoteKind string, apiNoteID string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_si_note" "scc_si_note" {
			provider_id = "%s"
			short_description = "%s"
			long_description = "%s"
			kind = "%s"
			note_id = "%s"
			reported_by {
				id = "id"
				title = "title"
				url = "url"
			}
		}

		data "ibm_scc_si_note" "scc_si_note" {
			provider_id = ibm_scc_si_note.scc_si_note.provider_id
			note_id = ibm_scc_si_note.scc_si_note.note_id
		}
	`, apiNoteProviderID, apiNoteShortDescription, apiNoteLongDescription, apiNoteKind, apiNoteID)
}

func testAccCheckIBMSccSiNoteDataSourceConfig(apiNoteProviderID string, apiNoteShortDescription string, apiNoteLongDescription string, apiNoteKind string, apiNoteID string, apiNoteShared string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_si_note" "scc_si_note" {
			provider_id = "%s"
			short_description = "%s"
			long_description = "%s"
			kind = "%s"
			note_id = "%s"
			reported_by {
				id = "id"
				title = "title"
				url = "url"
			}
			related_url {
				label = "label"
				url = "url"
			}
			expiration_time = "2004-10-28T04:39:00.000Z"
			shared = %s
			finding {
				severity = "LOW"
				next_steps {
					title = "title"
					url = "url"
				}
			}
			kpi {
				aggregation_type = "SUM"
			}
			card {
				section = "section"
				title = "title"
				subtitle = "subtitle"
				order = 1
				finding_note_names = [ "finding_note_names" ]
				badge_text = "badge_text"
				badge_image = "badge_image"
				elements {
					text = "text"
					kind = "NUMERIC"
					default_time_range = "1d"
					value_type {
						kind = "KPI"
						kpi_note_name = "kpi_note_name"
						text = "text"
						finding_note_names = [ "finding_note_names" ]
					}
					value_types {
						kind = "KPI"
						kpi_note_name = "kpi_note_name"
						text = "text"
						finding_note_names = [ "finding_note_names" ]
					}
					default_interval = "default_interval"
				}
			}
			section {
				title = "title"
				image = "image"
			}
		}

		data "ibm_scc_si_note" "scc_si_note" {
			provider_id = ibm_scc_si_note.scc_si_note.provider_id
			note_id = ibm_scc_si_note.scc_si_note.note_id
		}
	`, apiNoteProviderID, apiNoteShortDescription, apiNoteLongDescription, apiNoteKind, apiNoteID, apiNoteShared)
}
