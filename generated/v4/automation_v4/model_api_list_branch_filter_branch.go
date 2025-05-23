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

// ApiListBranchFilterBranch - The list criteria that determine when to execute this branch. The first matching branch will execute.
type ApiListBranchFilterBranch struct {
	PublicAndFilterBranch *PublicAndFilterBranch
	PublicAssociationFilterBranch *PublicAssociationFilterBranch
	PublicNotAllFilterBranch *PublicNotAllFilterBranch
	PublicNotAnyFilterBranch *PublicNotAnyFilterBranch
	PublicOrFilterBranch *PublicOrFilterBranch
	PublicPropertyAssociationFilterBranch *PublicPropertyAssociationFilterBranch
	PublicRestrictedFilterBranch *PublicRestrictedFilterBranch
	PublicUnifiedEventsFilterBranch *PublicUnifiedEventsFilterBranch
}

// PublicAndFilterBranchAsApiListBranchFilterBranch is a convenience function that returns PublicAndFilterBranch wrapped in ApiListBranchFilterBranch
func PublicAndFilterBranchAsApiListBranchFilterBranch(v *PublicAndFilterBranch) ApiListBranchFilterBranch {
	return ApiListBranchFilterBranch{
		PublicAndFilterBranch: v,
	}
}

// PublicAssociationFilterBranchAsApiListBranchFilterBranch is a convenience function that returns PublicAssociationFilterBranch wrapped in ApiListBranchFilterBranch
func PublicAssociationFilterBranchAsApiListBranchFilterBranch(v *PublicAssociationFilterBranch) ApiListBranchFilterBranch {
	return ApiListBranchFilterBranch{
		PublicAssociationFilterBranch: v,
	}
}

// PublicNotAllFilterBranchAsApiListBranchFilterBranch is a convenience function that returns PublicNotAllFilterBranch wrapped in ApiListBranchFilterBranch
func PublicNotAllFilterBranchAsApiListBranchFilterBranch(v *PublicNotAllFilterBranch) ApiListBranchFilterBranch {
	return ApiListBranchFilterBranch{
		PublicNotAllFilterBranch: v,
	}
}

// PublicNotAnyFilterBranchAsApiListBranchFilterBranch is a convenience function that returns PublicNotAnyFilterBranch wrapped in ApiListBranchFilterBranch
func PublicNotAnyFilterBranchAsApiListBranchFilterBranch(v *PublicNotAnyFilterBranch) ApiListBranchFilterBranch {
	return ApiListBranchFilterBranch{
		PublicNotAnyFilterBranch: v,
	}
}

// PublicOrFilterBranchAsApiListBranchFilterBranch is a convenience function that returns PublicOrFilterBranch wrapped in ApiListBranchFilterBranch
func PublicOrFilterBranchAsApiListBranchFilterBranch(v *PublicOrFilterBranch) ApiListBranchFilterBranch {
	return ApiListBranchFilterBranch{
		PublicOrFilterBranch: v,
	}
}

// PublicPropertyAssociationFilterBranchAsApiListBranchFilterBranch is a convenience function that returns PublicPropertyAssociationFilterBranch wrapped in ApiListBranchFilterBranch
func PublicPropertyAssociationFilterBranchAsApiListBranchFilterBranch(v *PublicPropertyAssociationFilterBranch) ApiListBranchFilterBranch {
	return ApiListBranchFilterBranch{
		PublicPropertyAssociationFilterBranch: v,
	}
}

// PublicRestrictedFilterBranchAsApiListBranchFilterBranch is a convenience function that returns PublicRestrictedFilterBranch wrapped in ApiListBranchFilterBranch
func PublicRestrictedFilterBranchAsApiListBranchFilterBranch(v *PublicRestrictedFilterBranch) ApiListBranchFilterBranch {
	return ApiListBranchFilterBranch{
		PublicRestrictedFilterBranch: v,
	}
}

// PublicUnifiedEventsFilterBranchAsApiListBranchFilterBranch is a convenience function that returns PublicUnifiedEventsFilterBranch wrapped in ApiListBranchFilterBranch
func PublicUnifiedEventsFilterBranchAsApiListBranchFilterBranch(v *PublicUnifiedEventsFilterBranch) ApiListBranchFilterBranch {
	return ApiListBranchFilterBranch{
		PublicUnifiedEventsFilterBranch: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApiListBranchFilterBranch) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ApiListBranchFilterBranch)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApiListBranchFilterBranch)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApiListBranchFilterBranch) MarshalJSON() ([]byte, error) {
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
func (obj *ApiListBranchFilterBranch) GetActualInstance() (interface{}) {
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
func (obj ApiListBranchFilterBranch) GetActualInstanceValue() (interface{}) {
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

type NullableApiListBranchFilterBranch struct {
	value *ApiListBranchFilterBranch
	isSet bool
}

func (v NullableApiListBranchFilterBranch) Get() *ApiListBranchFilterBranch {
	return v.value
}

func (v *NullableApiListBranchFilterBranch) Set(val *ApiListBranchFilterBranch) {
	v.value = val
	v.isSet = true
}

func (v NullableApiListBranchFilterBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableApiListBranchFilterBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiListBranchFilterBranch(val *ApiListBranchFilterBranch) *NullableApiListBranchFilterBranch {
	return &NullableApiListBranchFilterBranch{value: val, isSet: true}
}

func (v NullableApiListBranchFilterBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiListBranchFilterBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


