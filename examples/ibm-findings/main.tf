variable "ibmcloud_api_key" {}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision scc_si_note resource instance
resource "ibm_scc_si_note" "scc_si_note_instance" {
  provider_id = var.scc_si_note_provider_id
  short_description = var.scc_si_note_short_description
  long_description = var.scc_si_note_long_description
  kind = var.scc_si_note_kind
  note_id = var.scc_si_note_note_id
  reported_by {
    id = "id"
    title = "title"
    url = "url"
  }
  // related_url = var.scc_si_note_related_url
  expiration_time = var.scc_si_note_expiration_time
  shared = var.scc_si_note_shared
  finding {
    severity = "LOW"
    next_steps {
      title = "title"
      url = "url"
    }
  }
}

// Create scc_si_provider data source
  data "ibm_scc_si_provider" "scc_si_provider_instance" {
   id = var.scc_si_provider_id
 }

// Create scc_si_note data source
 data "ibm_scc_si_note" "scc_si_note_instance" {
  provider_id = var.scc_si_note_provider_id
  note_id = var.scc_si_note_note_id
  }
