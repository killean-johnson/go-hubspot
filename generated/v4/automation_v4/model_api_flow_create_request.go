/*
Automation V4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package automation_v4

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// ApiFlowCreateRequest - struct for ApiFlowCreateRequest
type ApiFlowCreateRequest struct {
	ApiContactFlowCreateRequest *ApiContactFlowCreateRequest
	ApiPlatformFlowCreateRequest *ApiPlatformFlowCreateRequest
}

// ApiContactFlowCreateRequestAsApiFlowCreateRequest is a convenience function that returns ApiContactFlowCreateRequest wrapped in ApiFlowCreateRequest
func ApiContactFlowCreateRequestAsApiFlowCreateRequest(v *ApiContactFlowCreateRequest) ApiFlowCreateRequest {
	return ApiFlowCreateRequest{
		ApiContactFlowCreateRequest: v,
	}
}

// ApiPlatformFlowCreateRequestAsApiFlowCreateRequest is a convenience function that returns ApiPlatformFlowCreateRequest wrapped in ApiFlowCreateRequest
func ApiPlatformFlowCreateRequestAsApiFlowCreateRequest(v *ApiPlatformFlowCreateRequest) ApiFlowCreateRequest {
	return ApiFlowCreateRequest{
		ApiPlatformFlowCreateRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApiFlowCreateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApiContactFlowCreateRequest
	err = newStrictDecoder(data).Decode(&dst.ApiContactFlowCreateRequest)
	if err == nil {
		jsonApiContactFlowCreateRequest, _ := json.Marshal(dst.ApiContactFlowCreateRequest)
		if string(jsonApiContactFlowCreateRequest) == "{}" { // empty struct
			dst.ApiContactFlowCreateRequest = nil
		} else {
			if err = validator.Validate(dst.ApiContactFlowCreateRequest); err != nil {
				dst.ApiContactFlowCreateRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiContactFlowCreateRequest = nil
	}

	// try to unmarshal data into ApiPlatformFlowCreateRequest
	err = newStrictDecoder(data).Decode(&dst.ApiPlatformFlowCreateRequest)
	if err == nil {
		jsonApiPlatformFlowCreateRequest, _ := json.Marshal(dst.ApiPlatformFlowCreateRequest)
		if string(jsonApiPlatformFlowCreateRequest) == "{}" { // empty struct
			dst.ApiPlatformFlowCreateRequest = nil
		} else {
			if err = validator.Validate(dst.ApiPlatformFlowCreateRequest); err != nil {
				dst.ApiPlatformFlowCreateRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiPlatformFlowCreateRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApiContactFlowCreateRequest = nil
		dst.ApiPlatformFlowCreateRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApiFlowCreateRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApiFlowCreateRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApiFlowCreateRequest) MarshalJSON() ([]byte, error) {
	if src.ApiContactFlowCreateRequest != nil {
		return json.Marshal(&src.ApiContactFlowCreateRequest)
	}

	if src.ApiPlatformFlowCreateRequest != nil {
		return json.Marshal(&src.ApiPlatformFlowCreateRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApiFlowCreateRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApiContactFlowCreateRequest != nil {
		return obj.ApiContactFlowCreateRequest
	}

	if obj.ApiPlatformFlowCreateRequest != nil {
		return obj.ApiPlatformFlowCreateRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ApiFlowCreateRequest) GetActualInstanceValue() (interface{}) {
	if obj.ApiContactFlowCreateRequest != nil {
		return *obj.ApiContactFlowCreateRequest
	}

	if obj.ApiPlatformFlowCreateRequest != nil {
		return *obj.ApiPlatformFlowCreateRequest
	}

	// all schemas are nil
	return nil
}

type NullableApiFlowCreateRequest struct {
	value *ApiFlowCreateRequest
	isSet bool
}

func (v NullableApiFlowCreateRequest) Get() *ApiFlowCreateRequest {
	return v.value
}

func (v *NullableApiFlowCreateRequest) Set(val *ApiFlowCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFlowCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFlowCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFlowCreateRequest(val *ApiFlowCreateRequest) *NullableApiFlowCreateRequest {
	return &NullableApiFlowCreateRequest{value: val, isSet: true}
}

func (v NullableApiFlowCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFlowCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


