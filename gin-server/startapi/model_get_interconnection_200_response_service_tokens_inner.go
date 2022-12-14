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

type GetInterconnection200ResponseServiceTokensInner struct {

	// The expiration date and time of the Fabric service token. Once a service token is expired, it is no longer redeemable.
	ExpiresAt time.Time `json:"expires_at,omitempty"`

	// The UUID that can be used on the Fabric Portal to redeem either an A-Side or Z-Side Service Token. For Fabric VCs (Metal Billed), this UUID will represent an A-Side Service Token, which will allow interconnections to be made from Equinix Metal to other Service Providers on Fabric. For Fabric VCs (Fabric Billed), this UUID will represent a Z-Side Service Token, which will allow interconnections to be made to connect an owned Fabric Port or  Virtual Device to Equinix Metal.
	Id string `json:"id,omitempty"`

	// The maximum speed that can be selected on the Fabric Portal when configuring a interconnection with either  an A-Side or Z-Side Service Token. For Fabric VCs (Metal Billed), this is what the billing is based off of, and can be one of the following options, '50mbps', '200mbps', '500mbps', '1gbps', '2gbps', '5gbps' or '10gbps'. For Fabric VCs (Fabric Billed), this will default to 10Gbps.
	MaxAllowedSpeed int32 `json:"max_allowed_speed,omitempty"`

	// Either primary or secondary, depending on which interconnection the service token is associated to.
	Role string `json:"role,omitempty"`

	// Either 'a_side' or 'z_side', depending on which type of Fabric VC was requested.
	ServiceTokenType string `json:"service_token_type,omitempty"`

	// The state of the service token that corresponds with the service token state on Fabric. An 'inactive' state refers to a token that has not been redeemed yet on the Fabric side, an 'active' state refers to a token that has already been redeemed, and an 'expired' state refers to a token that has reached its expiry time.
	State string `json:"state,omitempty"`
}
