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

// ApiPlatformFlowAllOfSuppressionFilterBranch - struct for ApiPlatformFlowAllOfSuppressionFilterBranch
type ApiPlatformFlowAllOfSuppressionFilterBranch struct {
	PublicAndFilterBranch *PublicAndFilterBranch
	PublicAssociationFilterBranch *PublicAssociationFilterBranch
	PublicNotAllFilterBranch *PublicNotAllFilterBranch
	PublicNotAnyFilterBranch *PublicNotAnyFilterBranch
	PublicOrFilterBranch *PublicOrFilterBranch
	PublicPropertyAssociationFilterBranch *PublicPropertyAssociationFilterBranch
	PublicRestrictedFilterBranch *PublicRestrictedFilterBranch
	PublicUnifiedEventsFilterBranch *PublicUnifiedEventsFilterBranch
}

// PublicAndFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch is a convenience function that returns PublicAndFilterBranch wrapped in ApiPlatformFlowAllOfSuppressionFilterBranch
func PublicAndFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch(v *PublicAndFilterBranch) ApiPlatformFlowAllOfSuppressionFilterBranch {
	return ApiPlatformFlowAllOfSuppressionFilterBranch{
		PublicAndFilterBranch: v,
	}
}

// PublicAssociationFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch is a convenience function that returns PublicAssociationFilterBranch wrapped in ApiPlatformFlowAllOfSuppressionFilterBranch
func PublicAssociationFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch(v *PublicAssociationFilterBranch) ApiPlatformFlowAllOfSuppressionFilterBranch {
	return ApiPlatformFlowAllOfSuppressionFilterBranch{
		PublicAssociationFilterBranch: v,
	}
}

// PublicNotAllFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch is a convenience function that returns PublicNotAllFilterBranch wrapped in ApiPlatformFlowAllOfSuppressionFilterBranch
func PublicNotAllFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch(v *PublicNotAllFilterBranch) ApiPlatformFlowAllOfSuppressionFilterBranch {
	return ApiPlatformFlowAllOfSuppressionFilterBranch{
		PublicNotAllFilterBranch: v,
	}
}

// PublicNotAnyFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch is a convenience function that returns PublicNotAnyFilterBranch wrapped in ApiPlatformFlowAllOfSuppressionFilterBranch
func PublicNotAnyFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch(v *PublicNotAnyFilterBranch) ApiPlatformFlowAllOfSuppressionFilterBranch {
	return ApiPlatformFlowAllOfSuppressionFilterBranch{
		PublicNotAnyFilterBranch: v,
	}
}

// PublicOrFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch is a convenience function that returns PublicOrFilterBranch wrapped in ApiPlatformFlowAllOfSuppressionFilterBranch
func PublicOrFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch(v *PublicOrFilterBranch) ApiPlatformFlowAllOfSuppressionFilterBranch {
	return ApiPlatformFlowAllOfSuppressionFilterBranch{
		PublicOrFilterBranch: v,
	}
}

// PublicPropertyAssociationFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch is a convenience function that returns PublicPropertyAssociationFilterBranch wrapped in ApiPlatformFlowAllOfSuppressionFilterBranch
func PublicPropertyAssociationFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch(v *PublicPropertyAssociationFilterBranch) ApiPlatformFlowAllOfSuppressionFilterBranch {
	return ApiPlatformFlowAllOfSuppressionFilterBranch{
		PublicPropertyAssociationFilterBranch: v,
	}
}

// PublicRestrictedFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch is a convenience function that returns PublicRestrictedFilterBranch wrapped in ApiPlatformFlowAllOfSuppressionFilterBranch
func PublicRestrictedFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch(v *PublicRestrictedFilterBranch) ApiPlatformFlowAllOfSuppressionFilterBranch {
	return ApiPlatformFlowAllOfSuppressionFilterBranch{
		PublicRestrictedFilterBranch: v,
	}
}

// PublicUnifiedEventsFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch is a convenience function that returns PublicUnifiedEventsFilterBranch wrapped in ApiPlatformFlowAllOfSuppressionFilterBranch
func PublicUnifiedEventsFilterBranchAsApiPlatformFlowAllOfSuppressionFilterBranch(v *PublicUnifiedEventsFilterBranch) ApiPlatformFlowAllOfSuppressionFilterBranch {
	return ApiPlatformFlowAllOfSuppressionFilterBranch{
		PublicUnifiedEventsFilterBranch: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApiPlatformFlowAllOfSuppressionFilterBranch) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PublicAndFilterBranch
	err = newStrictDecoder(data).Decode(&dst.PublicAndFilterBranch)
	if err == nil {
		jsonPublicAndFilterBranch, _ := json.Marshal(dst.PublicAndFilterBranch)
		if string(jsonPublicAndFilterBranch) == "{}" { // empty struct
			dst.PublicAndFilterBranch = nil
		} else {
			if err = validator.Validate(dst.PublicAndFilterBranch); err != nil {
				dst.PublicAndFilterBranch = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicAndFilterBranch = nil
	}

	// try to unmarshal data into PublicAssociationFilterBranch
	err = newStrictDecoder(data).Decode(&dst.PublicAssociationFilterBranch)
	if err == nil {
		jsonPublicAssociationFilterBranch, _ := json.Marshal(dst.PublicAssociationFilterBranch)
		if string(jsonPublicAssociationFilterBranch) == "{}" { // empty struct
			dst.PublicAssociationFilterBranch = nil
		} else {
			if err = validator.Validate(dst.PublicAssociationFilterBranch); err != nil {
				dst.PublicAssociationFilterBranch = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicAssociationFilterBranch = nil
	}

	// try to unmarshal data into PublicNotAllFilterBranch
	err = newStrictDecoder(data).Decode(&dst.PublicNotAllFilterBranch)
	if err == nil {
		jsonPublicNotAllFilterBranch, _ := json.Marshal(dst.PublicNotAllFilterBranch)
		if string(jsonPublicNotAllFilterBranch) == "{}" { // empty struct
			dst.PublicNotAllFilterBranch = nil
		} else {
			if err = validator.Validate(dst.PublicNotAllFilterBranch); err != nil {
				dst.PublicNotAllFilterBranch = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicNotAllFilterBranch = nil
	}

	// try to unmarshal data into PublicNotAnyFilterBranch
	err = newStrictDecoder(data).Decode(&dst.PublicNotAnyFilterBranch)
	if err == nil {
		jsonPublicNotAnyFilterBranch, _ := json.Marshal(dst.PublicNotAnyFilterBranch)
		if string(jsonPublicNotAnyFilterBranch) == "{}" { // empty struct
			dst.PublicNotAnyFilterBranch = nil
		} else {
			if err = validator.Validate(dst.PublicNotAnyFilterBranch); err != nil {
				dst.PublicNotAnyFilterBranch = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicNotAnyFilterBranch = nil
	}

	// try to unmarshal data into PublicOrFilterBranch
	err = newStrictDecoder(data).Decode(&dst.PublicOrFilterBranch)
	if err == nil {
		jsonPublicOrFilterBranch, _ := json.Marshal(dst.PublicOrFilterBranch)
		if string(jsonPublicOrFilterBranch) == "{}" { // empty struct
			dst.PublicOrFilterBranch = nil
		} else {
			if err = validator.Validate(dst.PublicOrFilterBranch); err != nil {
				dst.PublicOrFilterBranch = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicOrFilterBranch = nil
	}

	// try to unmarshal data into PublicPropertyAssociationFilterBranch
	err = newStrictDecoder(data).Decode(&dst.PublicPropertyAssociationFilterBranch)
	if err == nil {
		jsonPublicPropertyAssociationFilterBranch, _ := json.Marshal(dst.PublicPropertyAssociationFilterBranch)
		if string(jsonPublicPropertyAssociationFilterBranch) == "{}" { // empty struct
			dst.PublicPropertyAssociationFilterBranch = nil
		} else {
			if err = validator.Validate(dst.PublicPropertyAssociationFilterBranch); err != nil {
				dst.PublicPropertyAssociationFilterBranch = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicPropertyAssociationFilterBranch = nil
	}

	// try to unmarshal data into PublicRestrictedFilterBranch
	err = newStrictDecoder(data).Decode(&dst.PublicRestrictedFilterBranch)
	if err == nil {
		jsonPublicRestrictedFilterBranch, _ := json.Marshal(dst.PublicRestrictedFilterBranch)
		if string(jsonPublicRestrictedFilterBranch) == "{}" { // empty struct
			dst.PublicRestrictedFilterBranch = nil
		} else {
			if err = validator.Validate(dst.PublicRestrictedFilterBranch); err != nil {
				dst.PublicRestrictedFilterBranch = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicRestrictedFilterBranch = nil
	}

	// try to unmarshal data into PublicUnifiedEventsFilterBranch
	err = newStrictDecoder(data).Decode(&dst.PublicUnifiedEventsFilterBranch)
	if err == nil {
		jsonPublicUnifiedEventsFilterBranch, _ := json.Marshal(dst.PublicUnifiedEventsFilterBranch)
		if string(jsonPublicUnifiedEventsFilterBranch) == "{}" { // empty struct
			dst.PublicUnifiedEventsFilterBranch = nil
		} else {
			if err = validator.Validate(dst.PublicUnifiedEventsFilterBranch); err != nil {
				dst.PublicUnifiedEventsFilterBranch = nil
			} else {
				match++
			}
		}
	} else {
		dst.PublicUnifiedEventsFilterBranch = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PublicAndFilterBranch = nil
		dst.PublicAssociationFilterBranch = nil
		dst.PublicNotAllFilterBranch = nil
		dst.PublicNotAnyFilterBranch = nil
		dst.PublicOrFilterBranch = nil
		dst.PublicPropertyAssociationFilterBranch = nil
		dst.PublicRestrictedFilterBranch = nil
		dst.PublicUnifiedEventsFilterBranch = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApiPlatformFlowAllOfSuppressionFilterBranch)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApiPlatformFlowAllOfSuppressionFilterBranch)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApiPlatformFlowAllOfSuppressionFilterBranch) MarshalJSON() ([]byte, error) {
	if src.PublicAndFilterBranch != nil {
		return json.Marshal(&src.PublicAndFilterBranch)
	}

	if src.PublicAssociationFilterBranch != nil {
		return json.Marshal(&src.PublicAssociationFilterBranch)
	}

	if src.PublicNotAllFilterBranch != nil {
		return json.Marshal(&src.PublicNotAllFilterBranch)
	}

	if src.PublicNotAnyFilterBranch != nil {
		return json.Marshal(&src.PublicNotAnyFilterBranch)
	}

	if src.PublicOrFilterBranch != nil {
		return json.Marshal(&src.PublicOrFilterBranch)
	}

	if src.PublicPropertyAssociationFilterBranch != nil {
		return json.Marshal(&src.PublicPropertyAssociationFilterBranch)
	}

	if src.PublicRestrictedFilterBranch != nil {
		return json.Marshal(&src.PublicRestrictedFilterBranch)
	}

	if src.PublicUnifiedEventsFilterBranch != nil {
		return json.Marshal(&src.PublicUnifiedEventsFilterBranch)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApiPlatformFlowAllOfSuppressionFilterBranch) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PublicAndFilterBranch != nil {
		return obj.PublicAndFilterBranch
	}

	if obj.PublicAssociationFilterBranch != nil {
		return obj.PublicAssociationFilterBranch
	}

	if obj.PublicNotAllFilterBranch != nil {
		return obj.PublicNotAllFilterBranch
	}

	if obj.PublicNotAnyFilterBranch != nil {
		return obj.PublicNotAnyFilterBranch
	}

	if obj.PublicOrFilterBranch != nil {
		return obj.PublicOrFilterBranch
	}

	if obj.PublicPropertyAssociationFilterBranch != nil {
		return obj.PublicPropertyAssociationFilterBranch
	}

	if obj.PublicRestrictedFilterBranch != nil {
		return obj.PublicRestrictedFilterBranch
	}

	if obj.PublicUnifiedEventsFilterBranch != nil {
		return obj.PublicUnifiedEventsFilterBranch
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ApiPlatformFlowAllOfSuppressionFilterBranch) GetActualInstanceValue() (interface{}) {
	if obj.PublicAndFilterBranch != nil {
		return *obj.PublicAndFilterBranch
	}

	if obj.PublicAssociationFilterBranch != nil {
		return *obj.PublicAssociationFilterBranch
	}

	if obj.PublicNotAllFilterBranch != nil {
		return *obj.PublicNotAllFilterBranch
	}

	if obj.PublicNotAnyFilterBranch != nil {
		return *obj.PublicNotAnyFilterBranch
	}

	if obj.PublicOrFilterBranch != nil {
		return *obj.PublicOrFilterBranch
	}

	if obj.PublicPropertyAssociationFilterBranch != nil {
		return *obj.PublicPropertyAssociationFilterBranch
	}

	if obj.PublicRestrictedFilterBranch != nil {
		return *obj.PublicRestrictedFilterBranch
	}

	if obj.PublicUnifiedEventsFilterBranch != nil {
		return *obj.PublicUnifiedEventsFilterBranch
	}

	// all schemas are nil
	return nil
}

type NullableApiPlatformFlowAllOfSuppressionFilterBranch struct {
	value *ApiPlatformFlowAllOfSuppressionFilterBranch
	isSet bool
}

func (v NullableApiPlatformFlowAllOfSuppressionFilterBranch) Get() *ApiPlatformFlowAllOfSuppressionFilterBranch {
	return v.value
}

func (v *NullableApiPlatformFlowAllOfSuppressionFilterBranch) Set(val *ApiPlatformFlowAllOfSuppressionFilterBranch) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPlatformFlowAllOfSuppressionFilterBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPlatformFlowAllOfSuppressionFilterBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPlatformFlowAllOfSuppressionFilterBranch(val *ApiPlatformFlowAllOfSuppressionFilterBranch) *NullableApiPlatformFlowAllOfSuppressionFilterBranch {
	return &NullableApiPlatformFlowAllOfSuppressionFilterBranch{value: val, isSet: true}
}

func (v NullableApiPlatformFlowAllOfSuppressionFilterBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPlatformFlowAllOfSuppressionFilterBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


