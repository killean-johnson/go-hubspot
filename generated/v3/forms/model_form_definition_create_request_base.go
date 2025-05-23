/*
Forms

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package forms

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// FormDefinitionCreateRequestBase - struct for FormDefinitionCreateRequestBase
type FormDefinitionCreateRequestBase struct {
	HubSpotFormDefinitionCreateRequest *HubSpotFormDefinitionCreateRequest
}

// HubSpotFormDefinitionCreateRequestAsFormDefinitionCreateRequestBase is a convenience function that returns HubSpotFormDefinitionCreateRequest wrapped in FormDefinitionCreateRequestBase
func HubSpotFormDefinitionCreateRequestAsFormDefinitionCreateRequestBase(v *HubSpotFormDefinitionCreateRequest) FormDefinitionCreateRequestBase {
	return FormDefinitionCreateRequestBase{
		HubSpotFormDefinitionCreateRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *FormDefinitionCreateRequestBase) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HubSpotFormDefinitionCreateRequest
	err = newStrictDecoder(data).Decode(&dst.HubSpotFormDefinitionCreateRequest)
	if err == nil {
		jsonHubSpotFormDefinitionCreateRequest, _ := json.Marshal(dst.HubSpotFormDefinitionCreateRequest)
		if string(jsonHubSpotFormDefinitionCreateRequest) == "{}" { // empty struct
			dst.HubSpotFormDefinitionCreateRequest = nil
		} else {
			if err = validator.Validate(dst.HubSpotFormDefinitionCreateRequest); err != nil {
				dst.HubSpotFormDefinitionCreateRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.HubSpotFormDefinitionCreateRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.HubSpotFormDefinitionCreateRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(FormDefinitionCreateRequestBase)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(FormDefinitionCreateRequestBase)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FormDefinitionCreateRequestBase) MarshalJSON() ([]byte, error) {
	if src.HubSpotFormDefinitionCreateRequest != nil {
		return json.Marshal(&src.HubSpotFormDefinitionCreateRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FormDefinitionCreateRequestBase) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.HubSpotFormDefinitionCreateRequest != nil {
		return obj.HubSpotFormDefinitionCreateRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj FormDefinitionCreateRequestBase) GetActualInstanceValue() (interface{}) {
	if obj.HubSpotFormDefinitionCreateRequest != nil {
		return *obj.HubSpotFormDefinitionCreateRequest
	}

	// all schemas are nil
	return nil
}

type NullableFormDefinitionCreateRequestBase struct {
	value *FormDefinitionCreateRequestBase
	isSet bool
}

func (v NullableFormDefinitionCreateRequestBase) Get() *FormDefinitionCreateRequestBase {
	return v.value
}

func (v *NullableFormDefinitionCreateRequestBase) Set(val *FormDefinitionCreateRequestBase) {
	v.value = val
	v.isSet = true
}

func (v NullableFormDefinitionCreateRequestBase) IsSet() bool {
	return v.isSet
}

func (v *NullableFormDefinitionCreateRequestBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormDefinitionCreateRequestBase(val *FormDefinitionCreateRequestBase) *NullableFormDefinitionCreateRequestBase {
	return &NullableFormDefinitionCreateRequestBase{value: val, isSet: true}
}

func (v NullableFormDefinitionCreateRequestBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormDefinitionCreateRequestBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


