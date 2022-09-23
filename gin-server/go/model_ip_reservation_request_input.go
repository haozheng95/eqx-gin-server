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

type IpReservationRequestInput struct {
	Comments string `json:"comments,omitempty"`

	Customdata map[string]interface{} `json:"customdata,omitempty"`

	Details string `json:"details,omitempty"`

	Facility string `json:"facility,omitempty"`

	FailOnApprovalRequired bool `json:"fail_on_approval_required,omitempty"`

	// The code of the metro you are requesting the IP reservation in.
	Metro string `json:"metro,omitempty"`

	Quantity int32 `json:"quantity"`

	Tags []string `json:"tags,omitempty"`

	Type string `json:"type"`
}
