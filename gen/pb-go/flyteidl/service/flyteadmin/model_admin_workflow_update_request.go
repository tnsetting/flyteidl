/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Request to set the referenced workflow state to the configured value.
type AdminWorkflowUpdateRequest struct {
	// Identifier of workflow for which to change state.
	Id *CoreIdentifier `json:"id,omitempty"`
	// Desired state to apply to the workflow.
	State *AdminWorkflowState `json:"state,omitempty"`
}