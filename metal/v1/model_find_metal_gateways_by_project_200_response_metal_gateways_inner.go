/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
	_ "time"
)

// FindMetalGatewaysByProject200ResponseMetalGatewaysInner struct for FindMetalGatewaysByProject200ResponseMetalGatewaysInner
type FindMetalGatewaysByProject200ResponseMetalGatewaysInner struct {
	FindMetalGatewayById200ResponseOneOf  *FindMetalGatewayById200ResponseOneOf
	FindMetalGatewayById200ResponseOneOf1 *FindMetalGatewayById200ResponseOneOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FindMetalGatewayById200ResponseOneOf
	err = json.Unmarshal(data, &dst.FindMetalGatewayById200ResponseOneOf)
	if err == nil {
		jsonFindMetalGatewayById200ResponseOneOf, _ := json.Marshal(dst.FindMetalGatewayById200ResponseOneOf)
		if string(jsonFindMetalGatewayById200ResponseOneOf) == "{}" { // empty struct
			dst.FindMetalGatewayById200ResponseOneOf = nil
		} else {
			return nil // data stored in dst.FindMetalGatewayById200ResponseOneOf, return on the first match
		}
	} else {
		dst.FindMetalGatewayById200ResponseOneOf = nil
	}

	// try to unmarshal JSON data into FindMetalGatewayById200ResponseOneOf1
	err = json.Unmarshal(data, &dst.FindMetalGatewayById200ResponseOneOf1)
	if err == nil {
		jsonFindMetalGatewayById200ResponseOneOf1, _ := json.Marshal(dst.FindMetalGatewayById200ResponseOneOf1)
		if string(jsonFindMetalGatewayById200ResponseOneOf1) == "{}" { // empty struct
			dst.FindMetalGatewayById200ResponseOneOf1 = nil
		} else {
			return nil // data stored in dst.FindMetalGatewayById200ResponseOneOf1, return on the first match
		}
	} else {
		dst.FindMetalGatewayById200ResponseOneOf1 = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(FindMetalGatewaysByProject200ResponseMetalGatewaysInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) MarshalJSON() ([]byte, error) {
	if src.FindMetalGatewayById200ResponseOneOf != nil {
		return json.Marshal(&src.FindMetalGatewayById200ResponseOneOf)
	}

	if src.FindMetalGatewayById200ResponseOneOf1 != nil {
		return json.Marshal(&src.FindMetalGatewayById200ResponseOneOf1)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFindMetalGatewaysByProject200ResponseMetalGatewaysInner struct {
	value *FindMetalGatewaysByProject200ResponseMetalGatewaysInner
	isSet bool
}

func (v NullableFindMetalGatewaysByProject200ResponseMetalGatewaysInner) Get() *FindMetalGatewaysByProject200ResponseMetalGatewaysInner {
	return v.value
}

func (v *NullableFindMetalGatewaysByProject200ResponseMetalGatewaysInner) Set(val *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFindMetalGatewaysByProject200ResponseMetalGatewaysInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFindMetalGatewaysByProject200ResponseMetalGatewaysInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindMetalGatewaysByProject200ResponseMetalGatewaysInner(val *FindMetalGatewaysByProject200ResponseMetalGatewaysInner) *NullableFindMetalGatewaysByProject200ResponseMetalGatewaysInner {
	return &NullableFindMetalGatewaysByProject200ResponseMetalGatewaysInner{value: val, isSet: true}
}

func (v NullableFindMetalGatewaysByProject200ResponseMetalGatewaysInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindMetalGatewaysByProject200ResponseMetalGatewaysInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
