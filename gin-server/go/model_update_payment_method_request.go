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

type UpdatePaymentMethodRequest struct {
	BillingAddress map[string]interface{} `json:"billing_address,omitempty"`

	CardholderName string `json:"cardholder_name,omitempty"`

	Default bool `json:"default,omitempty"`

	ExpirationMonth string `json:"expiration_month,omitempty"`

	ExpirationYear int32 `json:"expiration_year,omitempty"`

	Name string `json:"name,omitempty"`
}
