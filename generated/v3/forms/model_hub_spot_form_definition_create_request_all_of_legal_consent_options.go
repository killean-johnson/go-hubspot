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

// HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions - struct for HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions
type HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions struct {
	LegalConsentOptionsExplicitConsentToProcess *LegalConsentOptionsExplicitConsentToProcess
	LegalConsentOptionsImplicitConsentToProcess *LegalConsentOptionsImplicitConsentToProcess
	LegalConsentOptionsLegitimateInterest *LegalConsentOptionsLegitimateInterest
	LegalConsentOptionsNone *LegalConsentOptionsNone
}

// LegalConsentOptionsExplicitConsentToProcessAsHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions is a convenience function that returns LegalConsentOptionsExplicitConsentToProcess wrapped in HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions
func LegalConsentOptionsExplicitConsentToProcessAsHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions(v *LegalConsentOptionsExplicitConsentToProcess) HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions {
	return HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions{
		LegalConsentOptionsExplicitConsentToProcess: v,
	}
}

// LegalConsentOptionsImplicitConsentToProcessAsHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions is a convenience function that returns LegalConsentOptionsImplicitConsentToProcess wrapped in HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions
func LegalConsentOptionsImplicitConsentToProcessAsHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions(v *LegalConsentOptionsImplicitConsentToProcess) HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions {
	return HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions{
		LegalConsentOptionsImplicitConsentToProcess: v,
	}
}

// LegalConsentOptionsLegitimateInterestAsHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions is a convenience function that returns LegalConsentOptionsLegitimateInterest wrapped in HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions
func LegalConsentOptionsLegitimateInterestAsHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions(v *LegalConsentOptionsLegitimateInterest) HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions {
	return HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions{
		LegalConsentOptionsLegitimateInterest: v,
	}
}

// LegalConsentOptionsNoneAsHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions is a convenience function that returns LegalConsentOptionsNone wrapped in HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions
func LegalConsentOptionsNoneAsHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions(v *LegalConsentOptionsNone) HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions {
	return HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions{
		LegalConsentOptionsNone: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LegalConsentOptionsExplicitConsentToProcess
	err = newStrictDecoder(data).Decode(&dst.LegalConsentOptionsExplicitConsentToProcess)
	if err == nil {
		jsonLegalConsentOptionsExplicitConsentToProcess, _ := json.Marshal(dst.LegalConsentOptionsExplicitConsentToProcess)
		if string(jsonLegalConsentOptionsExplicitConsentToProcess) == "{}" { // empty struct
			dst.LegalConsentOptionsExplicitConsentToProcess = nil
		} else {
			if err = validator.Validate(dst.LegalConsentOptionsExplicitConsentToProcess); err != nil {
				dst.LegalConsentOptionsExplicitConsentToProcess = nil
			} else {
				match++
			}
		}
	} else {
		dst.LegalConsentOptionsExplicitConsentToProcess = nil
	}

	// try to unmarshal data into LegalConsentOptionsImplicitConsentToProcess
	err = newStrictDecoder(data).Decode(&dst.LegalConsentOptionsImplicitConsentToProcess)
	if err == nil {
		jsonLegalConsentOptionsImplicitConsentToProcess, _ := json.Marshal(dst.LegalConsentOptionsImplicitConsentToProcess)
		if string(jsonLegalConsentOptionsImplicitConsentToProcess) == "{}" { // empty struct
			dst.LegalConsentOptionsImplicitConsentToProcess = nil
		} else {
			if err = validator.Validate(dst.LegalConsentOptionsImplicitConsentToProcess); err != nil {
				dst.LegalConsentOptionsImplicitConsentToProcess = nil
			} else {
				match++
			}
		}
	} else {
		dst.LegalConsentOptionsImplicitConsentToProcess = nil
	}

	// try to unmarshal data into LegalConsentOptionsLegitimateInterest
	err = newStrictDecoder(data).Decode(&dst.LegalConsentOptionsLegitimateInterest)
	if err == nil {
		jsonLegalConsentOptionsLegitimateInterest, _ := json.Marshal(dst.LegalConsentOptionsLegitimateInterest)
		if string(jsonLegalConsentOptionsLegitimateInterest) == "{}" { // empty struct
			dst.LegalConsentOptionsLegitimateInterest = nil
		} else {
			if err = validator.Validate(dst.LegalConsentOptionsLegitimateInterest); err != nil {
				dst.LegalConsentOptionsLegitimateInterest = nil
			} else {
				match++
			}
		}
	} else {
		dst.LegalConsentOptionsLegitimateInterest = nil
	}

	// try to unmarshal data into LegalConsentOptionsNone
	err = newStrictDecoder(data).Decode(&dst.LegalConsentOptionsNone)
	if err == nil {
		jsonLegalConsentOptionsNone, _ := json.Marshal(dst.LegalConsentOptionsNone)
		if string(jsonLegalConsentOptionsNone) == "{}" { // empty struct
			dst.LegalConsentOptionsNone = nil
		} else {
			if err = validator.Validate(dst.LegalConsentOptionsNone); err != nil {
				dst.LegalConsentOptionsNone = nil
			} else {
				match++
			}
		}
	} else {
		dst.LegalConsentOptionsNone = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LegalConsentOptionsExplicitConsentToProcess = nil
		dst.LegalConsentOptionsImplicitConsentToProcess = nil
		dst.LegalConsentOptionsLegitimateInterest = nil
		dst.LegalConsentOptionsNone = nil

		return fmt.Errorf("data matches more than one schema in oneOf(HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) MarshalJSON() ([]byte, error) {
	if src.LegalConsentOptionsExplicitConsentToProcess != nil {
		return json.Marshal(&src.LegalConsentOptionsExplicitConsentToProcess)
	}

	if src.LegalConsentOptionsImplicitConsentToProcess != nil {
		return json.Marshal(&src.LegalConsentOptionsImplicitConsentToProcess)
	}

	if src.LegalConsentOptionsLegitimateInterest != nil {
		return json.Marshal(&src.LegalConsentOptionsLegitimateInterest)
	}

	if src.LegalConsentOptionsNone != nil {
		return json.Marshal(&src.LegalConsentOptionsNone)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LegalConsentOptionsExplicitConsentToProcess != nil {
		return obj.LegalConsentOptionsExplicitConsentToProcess
	}

	if obj.LegalConsentOptionsImplicitConsentToProcess != nil {
		return obj.LegalConsentOptionsImplicitConsentToProcess
	}

	if obj.LegalConsentOptionsLegitimateInterest != nil {
		return obj.LegalConsentOptionsLegitimateInterest
	}

	if obj.LegalConsentOptionsNone != nil {
		return obj.LegalConsentOptionsNone
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) GetActualInstanceValue() (interface{}) {
	if obj.LegalConsentOptionsExplicitConsentToProcess != nil {
		return *obj.LegalConsentOptionsExplicitConsentToProcess
	}

	if obj.LegalConsentOptionsImplicitConsentToProcess != nil {
		return *obj.LegalConsentOptionsImplicitConsentToProcess
	}

	if obj.LegalConsentOptionsLegitimateInterest != nil {
		return *obj.LegalConsentOptionsLegitimateInterest
	}

	if obj.LegalConsentOptionsNone != nil {
		return *obj.LegalConsentOptionsNone
	}

	// all schemas are nil
	return nil
}

type NullableHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions struct {
	value *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions
	isSet bool
}

func (v NullableHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) Get() *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions {
	return v.value
}

func (v *NullableHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) Set(val *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions(val *HubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) *NullableHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions {
	return &NullableHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions{value: val, isSet: true}
}

func (v NullableHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHubSpotFormDefinitionCreateRequestAllOfLegalConsentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


