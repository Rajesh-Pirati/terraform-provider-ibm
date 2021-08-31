# Example for FindingsV1

This example illustrates how to use the FindingsV1

These types of resources are supported:

* scc_si_note

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## FindingsV1 resources

scc_si_note resource:

```hcl
resource "scc_si_note" "scc_si_note_instance" {
  provider_id = var.scc_si_note_provider_id
  short_description = var.scc_si_note_short_description
  long_description = var.scc_si_note_long_description
  kind = var.scc_si_note_kind
  note_id = var.scc_si_note_note_id
  reported_by = var.scc_si_note_reported_by
  related_url = var.scc_si_note_related_url
  expiration_time = var.scc_si_note_expiration_time
  shared = var.scc_si_note_shared
  finding = var.scc_si_note_finding
  kpi = var.scc_si_note_kpi
  card = var.scc_si_note_card
  section = var.scc_si_note_section
}
```

## FindingsV1 Data sources

scc_si_provider data source:

```hcl
data "scc_si_provider" "scc_si_provider_instance" {
  id = var.scc_si_provider_id
}
```
scc_si_note data source:

```hcl
data "scc_si_note" "scc_si_note_instance" {
  provider_id = var.scc_si_note_provider_id
  note_id = var.scc_si_note_note_id
}
```

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| provider_id | Part of the parent. This field contains the provider ID. For example: providers/{provider_id}. | `string` | true |
| short_description | A one sentence description of your note. | `string` | true |
| long_description | A more detailed description of your note. | `string` | true |
| kind | The type of note. Use this field to filter notes and occurences by kind. - FINDING&#58; The note and occurrence represent a finding. - KPI&#58; The note and occurrence represent a KPI value. - CARD&#58; The note represents a card showing findings and related metric values. - CARD_CONFIGURED&#58; The note represents a card configured for a user account. - SECTION&#58; The note represents a section in a dashboard. | `string` | true |
| note_id | The ID of the note. | `string` | true |
| reported_by | The entity reporting a note. | `` | true |
| related_url |  | `list()` | false |
| expiration_time | Time of expiration for this note, null if note does not expire. | `` | false |
| shared | True if this note can be shared by multiple accounts. | `bool` | false |
| finding | FindingType provides details about a finding note. | `` | false |
| kpi | KpiType provides details about a KPI note. | `` | false |
| card | Card provides details about a card kind of note. | `` | false |
| section | Card provides details about a card kind of note. | `` | false |
| id | The ID of the provider. | `string` | false |
| provider_id | Part of the parent. This field contains the provider ID. For example: providers/{provider_id}. | `string` | true |
| note_id | Second part of note `name`: providers/{provider_id}/notes/{note_id}. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| scc_si_note | scc_si_note object |
| scc_si_provider | scc_si_provider object |
| scc_si_note | scc_si_note object |
