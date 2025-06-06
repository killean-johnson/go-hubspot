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

// CollectionResponseFormDefinitionBaseForwardPagingResultsInner - struct for CollectionResponseFormDefinitionBaseForwardPagingResultsInner
type CollectionResponseFormDefinitionBaseForwardPagingResultsInner struct {
	HubSpotFormDefinition *HubSpotFormDefinition
}

// HubSpotFormDefinitionAsCollectionResponseFormDefinitionBaseForwardPagingResultsInner is a convenience function that returns HubSpotFormDefinition wrapped in CollectionResponseFormDefinitionBaseForwardPagingResultsInner
func HubSpotFormDefinitionAsCollectionResponseFormDefinitionBaseForwardPagingResultsInner(v *HubSpotFormDefinition) CollectionResponseFormDefinitionBaseForwardPagingResultsInner {
	return CollectionResponseFormDefinitionBaseForwardPagingResultsInner{
		HubSpotFormDefinition: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CollectionResponseFormDefinitionBaseForwardPagingResultsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HubSpotFormDefinition
	err = newStrictDecoder(data).Decode(&dst.HubSpotFormDefinition)
	if err == nil {
		jsonHubSpotFormDefinition, _ := json.Marshal(dst.HubSpotFormDefinition)
		if string(jsonHubSpotFormDefinition) == "{}" { // empty struct
			dst.HubSpotFormDefinition = nil
		} else {
			if err = validator.Validate(dst.HubSpotFormDefinition); err != nil {
				dst.HubSpotFormDefinition = nil
			} else {
				match++
			}
		}
	} else {
		dst.HubSpotFormDefinition = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.HubSpotFormDefinition = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CollectionResponseFormDefinitionBaseForwardPagingResultsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CollectionResponseFormDefinitionBaseForwardPagingResultsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CollectionResponseFormDefinitionBaseForwardPagingResultsInner) MarshalJSON() ([]byte, error) {
	if src.HubSpotFormDefinition != nil {
		return json.Marshal(&src.HubSpotFormDefinition)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CollectionResponseFormDefinitionBaseForwardPagingResultsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.HubSpotFormDefinition != nil {
		return obj.HubSpotFormDefinition
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CollectionResponseFormDefinitionBaseForwardPagingResultsInner) GetActualInstanceValue() (interface{}) {
	if obj.HubSpotFormDefinition != nil {
		return *obj.HubSpotFormDefinition
	}

	// all schemas are nil
	return nil
}

type NullableCollectionResponseFormDefinitionBaseForwardPagingResultsInner struct {
	value *CollectionResponseFormDefinitionBaseForwardPagingResultsInner
	isSet bool
}

func (v NullableCollectionResponseFormDefinitionBaseForwardPagingResultsInner) Get() *CollectionResponseFormDefinitionBaseForwardPagingResultsInner {
	return v.value
}

func (v *NullableCollectionResponseFormDefinitionBaseForwardPagingResultsInner) Set(val *CollectionResponseFormDefinitionBaseForwardPagingResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseFormDefinitionBaseForwardPagingResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseFormDefinitionBaseForwardPagingResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseFormDefinitionBaseForwardPagingResultsInner(val *CollectionResponseFormDefinitionBaseForwardPagingResultsInner) *NullableCollectionResponseFormDefinitionBaseForwardPagingResultsInner {
	return &NullableCollectionResponseFormDefinitionBaseForwardPagingResultsInner{value: val, isSet: true}
}

func (v NullableCollectionResponseFormDefinitionBaseForwardPagingResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseFormDefinitionBaseForwardPagingResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


