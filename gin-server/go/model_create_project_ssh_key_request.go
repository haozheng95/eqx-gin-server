/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package startapi

type CreateProjectSshKeyRequest struct {

	// List of instance UUIDs to associate SSH key with, when empty array is sent all instances belonging       to entity will be included
	InstancesIds []string `json:"instances_ids,omitempty"`

	Key string `json:"key,omitempty"`

	Label string `json:"label,omitempty"`
}
