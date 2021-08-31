// This allows scc_si_note data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_scc_si_note" {
  value       = ibm_scc_si_note.scc_si_note_instance
  description = "scc_si_note resource instance"
}
