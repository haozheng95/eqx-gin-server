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

import (
	"time"
)

type PaymentMethod struct {
	BillingAddress FindOrganizationPaymentMethods200ResponsePaymentMethodsInnerBillingAddress `json:"billing_address,omitempty"`

	CardType string `json:"card_type,omitempty"`

	CardholderName string `json:"cardholder_name,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	CreatedByUser FindBatchById200ResponseDevicesInner `json:"created_by_user,omitempty"`

	Default bool `json:"default,omitempty"`

	Email string `json:"email,omitempty"`

	ExpirationMonth string `json:"expiration_month,omitempty"`

	ExpirationYear string `json:"expiration_year,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Organization FindBatchById200ResponseDevicesInner `json:"organization,omitempty"`

	Projects []FindBatchById200ResponseDevicesInner `json:"projects,omitempty"`

	Type string `json:"type,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
