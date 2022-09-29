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

type ProjectUsage struct {
	Facility string `json:"facility,omitempty"`

	Name string `json:"name,omitempty"`

	Plan string `json:"plan,omitempty"`

	PlanVersion string `json:"plan_version,omitempty"`

	Price string `json:"price,omitempty"`

	Quantity string `json:"quantity,omitempty"`

	Total string `json:"total,omitempty"`

	Type string `json:"type,omitempty"`

	Unit string `json:"unit,omitempty"`
}