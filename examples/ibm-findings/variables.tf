
// Resource arguments for scc_si_note
variable "scc_si_note_provider_id" {
  description = "Part of the parent. This field contains the provider ID. For example: providers/{provider_id}."
  type        = string
  default     = "provider_id"
}
variable "scc_si_note_short_description" {
  description = "A one sentence description of your note."
  type        = string
  default     = "short_description"
}
variable "scc_si_note_long_description" {
  description = "A more detailed description of your note."
  type        = string
  default     = "long_description"
}
variable "scc_si_note_kind" {
  description = "The type of note. Use this field to filter notes and occurences by kind. - FINDING&#58; The note and occurrence represent a finding. - KPI&#58; The note and occurrence represent a KPI value. - CARD&#58; The note represents a card showing findings and related metric values. - CARD_CONFIGURED&#58; The note represents a card configured for a user account. - SECTION&#58; The note represents a section in a dashboard."
  type        = string
  default     = "FINDING"
}
variable "scc_si_note_note_id" {
  description = "The ID of the note."
  type        = string
  default     = "note_id"
}
variable "scc_si_note_related_url" {
  description = ""
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "scc_si_note_expiration_time" {
  description = "Time of expiration for this note, null if note does not expire."
  type        = string
  default     = "2021-01-31T09:44:12Z"
}
variable "scc_si_note_shared" {
  description = "True if this note can be shared by multiple accounts."
  type        = bool
  default     = false
}

// Data source arguments for scc_si_provider
variable "scc_si_provider_id" {
  description = "The ID of the provider."
  type        = string
  default     = "provider_id"
}

// Data source arguments for scc_si_note
