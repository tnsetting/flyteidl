/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type CoreContainer struct {
	Image string `json:"image,omitempty"`
	// Command to be executed, if not provided, the default entrypoint in the container image will be used.
	Command []string `json:"command,omitempty"`
	// These will default to Flyte given paths. If provided, the system will not append known paths. If the task still needs flyte's inputs and outputs path, add $(FLYTE_INPUT_FILE), $(FLYTE_OUTPUT_FILE) wherever makes sense and the system will populate these before executing the container.
	Args []string `json:"args,omitempty"`
	// Container resources requirement as specified by the container engine.
	Resources *CoreResources `json:"resources,omitempty"`
	// Environment variables will be set as the container is starting up.
	Env []CoreKeyValuePair `json:"env,omitempty"`
	// Allows extra configs to be available for the container. TODO: elaborate on how configs will become available.
	Config []CoreKeyValuePair `json:"config,omitempty"`
	Ports []CoreContainerPort `json:"ports,omitempty"`
	// BETA: This enables use of CoPilot. This makes it possible to to run a completely portable container, that uses inputs and outputs only from the local file-system and without having any reference to flyteidl. This is supported only on K8s at the moment. If CoPilot is enabled, then data will be mounted in accompanying directories specified in the CoPilot settings. If the directories  are not specified, inputs will be mounted onto and outputs will be uploaded from a pre-determined file-system path. Refer to the documentation to understand the default paths.
	UseCopilot bool `json:"use_copilot,omitempty"`
	// Optional configuration for CoPilot. If not specified, then default values are used.
	CopilotConfig *CoreCoPilot `json:"copilot_config,omitempty"`
}
