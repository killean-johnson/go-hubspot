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

// ApiInputVariableValue - struct for ApiInputVariableValue
type ApiInputVariableValue struct {
	ApiActionDataValue *ApiActionDataValue
	ApiAppendObjectPropertyValue *ApiAppendObjectPropertyValue
	ApiEnrollmentEventPropertyValue *ApiEnrollmentEventPropertyValue
	ApiFetchedObjectPropertyValue *ApiFetchedObjectPropertyValue
	ApiIncrementValue *ApiIncrementValue
	ApiObjectPropertyValue *ApiObjectPropertyValue
	ApiRelativeDateTimeValue *ApiRelativeDateTimeValue
	ApiStaticAppendValue *ApiStaticAppendValue
	ApiStaticValue *ApiStaticValue
	ApiTimestampValue *ApiTimestampValue
}

// ApiActionDataValueAsApiInputVariableValue is a convenience function that returns ApiActionDataValue wrapped in ApiInputVariableValue
func ApiActionDataValueAsApiInputVariableValue(v *ApiActionDataValue) ApiInputVariableValue {
	return ApiInputVariableValue{
		ApiActionDataValue: v,
	}
}

// ApiAppendObjectPropertyValueAsApiInputVariableValue is a convenience function that returns ApiAppendObjectPropertyValue wrapped in ApiInputVariableValue
func ApiAppendObjectPropertyValueAsApiInputVariableValue(v *ApiAppendObjectPropertyValue) ApiInputVariableValue {
	return ApiInputVariableValue{
		ApiAppendObjectPropertyValue: v,
	}
}

// ApiEnrollmentEventPropertyValueAsApiInputVariableValue is a convenience function that returns ApiEnrollmentEventPropertyValue wrapped in ApiInputVariableValue
func ApiEnrollmentEventPropertyValueAsApiInputVariableValue(v *ApiEnrollmentEventPropertyValue) ApiInputVariableValue {
	return ApiInputVariableValue{
		ApiEnrollmentEventPropertyValue: v,
	}
}

// ApiFetchedObjectPropertyValueAsApiInputVariableValue is a convenience function that returns ApiFetchedObjectPropertyValue wrapped in ApiInputVariableValue
func ApiFetchedObjectPropertyValueAsApiInputVariableValue(v *ApiFetchedObjectPropertyValue) ApiInputVariableValue {
	return ApiInputVariableValue{
		ApiFetchedObjectPropertyValue: v,
	}
}

// ApiIncrementValueAsApiInputVariableValue is a convenience function that returns ApiIncrementValue wrapped in ApiInputVariableValue
func ApiIncrementValueAsApiInputVariableValue(v *ApiIncrementValue) ApiInputVariableValue {
	return ApiInputVariableValue{
		ApiIncrementValue: v,
	}
}

// ApiObjectPropertyValueAsApiInputVariableValue is a convenience function that returns ApiObjectPropertyValue wrapped in ApiInputVariableValue
func ApiObjectPropertyValueAsApiInputVariableValue(v *ApiObjectPropertyValue) ApiInputVariableValue {
	return ApiInputVariableValue{
		ApiObjectPropertyValue: v,
	}
}

// ApiRelativeDateTimeValueAsApiInputVariableValue is a convenience function that returns ApiRelativeDateTimeValue wrapped in ApiInputVariableValue
func ApiRelativeDateTimeValueAsApiInputVariableValue(v *ApiRelativeDateTimeValue) ApiInputVariableValue {
	return ApiInputVariableValue{
		ApiRelativeDateTimeValue: v,
	}
}

// ApiStaticAppendValueAsApiInputVariableValue is a convenience function that returns ApiStaticAppendValue wrapped in ApiInputVariableValue
func ApiStaticAppendValueAsApiInputVariableValue(v *ApiStaticAppendValue) ApiInputVariableValue {
	return ApiInputVariableValue{
		ApiStaticAppendValue: v,
	}
}

// ApiStaticValueAsApiInputVariableValue is a convenience function that returns ApiStaticValue wrapped in ApiInputVariableValue
func ApiStaticValueAsApiInputVariableValue(v *ApiStaticValue) ApiInputVariableValue {
	return ApiInputVariableValue{
		ApiStaticValue: v,
	}
}

// ApiTimestampValueAsApiInputVariableValue is a convenience function that returns ApiTimestampValue wrapped in ApiInputVariableValue
func ApiTimestampValueAsApiInputVariableValue(v *ApiTimestampValue) ApiInputVariableValue {
	return ApiInputVariableValue{
		ApiTimestampValue: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApiInputVariableValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApiActionDataValue
	err = newStrictDecoder(data).Decode(&dst.ApiActionDataValue)
	if err == nil {
		jsonApiActionDataValue, _ := json.Marshal(dst.ApiActionDataValue)
		if string(jsonApiActionDataValue) == "{}" { // empty struct
			dst.ApiActionDataValue = nil
		} else {
			if err = validator.Validate(dst.ApiActionDataValue); err != nil {
				dst.ApiActionDataValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiActionDataValue = nil
	}

	// try to unmarshal data into ApiAppendObjectPropertyValue
	err = newStrictDecoder(data).Decode(&dst.ApiAppendObjectPropertyValue)
	if err == nil {
		jsonApiAppendObjectPropertyValue, _ := json.Marshal(dst.ApiAppendObjectPropertyValue)
		if string(jsonApiAppendObjectPropertyValue) == "{}" { // empty struct
			dst.ApiAppendObjectPropertyValue = nil
		} else {
			if err = validator.Validate(dst.ApiAppendObjectPropertyValue); err != nil {
				dst.ApiAppendObjectPropertyValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiAppendObjectPropertyValue = nil
	}

	// try to unmarshal data into ApiEnrollmentEventPropertyValue
	err = newStrictDecoder(data).Decode(&dst.ApiEnrollmentEventPropertyValue)
	if err == nil {
		jsonApiEnrollmentEventPropertyValue, _ := json.Marshal(dst.ApiEnrollmentEventPropertyValue)
		if string(jsonApiEnrollmentEventPropertyValue) == "{}" { // empty struct
			dst.ApiEnrollmentEventPropertyValue = nil
		} else {
			if err = validator.Validate(dst.ApiEnrollmentEventPropertyValue); err != nil {
				dst.ApiEnrollmentEventPropertyValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiEnrollmentEventPropertyValue = nil
	}

	// try to unmarshal data into ApiFetchedObjectPropertyValue
	err = newStrictDecoder(data).Decode(&dst.ApiFetchedObjectPropertyValue)
	if err == nil {
		jsonApiFetchedObjectPropertyValue, _ := json.Marshal(dst.ApiFetchedObjectPropertyValue)
		if string(jsonApiFetchedObjectPropertyValue) == "{}" { // empty struct
			dst.ApiFetchedObjectPropertyValue = nil
		} else {
			if err = validator.Validate(dst.ApiFetchedObjectPropertyValue); err != nil {
				dst.ApiFetchedObjectPropertyValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiFetchedObjectPropertyValue = nil
	}

	// try to unmarshal data into ApiIncrementValue
	err = newStrictDecoder(data).Decode(&dst.ApiIncrementValue)
	if err == nil {
		jsonApiIncrementValue, _ := json.Marshal(dst.ApiIncrementValue)
		if string(jsonApiIncrementValue) == "{}" { // empty struct
			dst.ApiIncrementValue = nil
		} else {
			if err = validator.Validate(dst.ApiIncrementValue); err != nil {
				dst.ApiIncrementValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiIncrementValue = nil
	}

	// try to unmarshal data into ApiObjectPropertyValue
	err = newStrictDecoder(data).Decode(&dst.ApiObjectPropertyValue)
	if err == nil {
		jsonApiObjectPropertyValue, _ := json.Marshal(dst.ApiObjectPropertyValue)
		if string(jsonApiObjectPropertyValue) == "{}" { // empty struct
			dst.ApiObjectPropertyValue = nil
		} else {
			if err = validator.Validate(dst.ApiObjectPropertyValue); err != nil {
				dst.ApiObjectPropertyValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiObjectPropertyValue = nil
	}

	// try to unmarshal data into ApiRelativeDateTimeValue
	err = newStrictDecoder(data).Decode(&dst.ApiRelativeDateTimeValue)
	if err == nil {
		jsonApiRelativeDateTimeValue, _ := json.Marshal(dst.ApiRelativeDateTimeValue)
		if string(jsonApiRelativeDateTimeValue) == "{}" { // empty struct
			dst.ApiRelativeDateTimeValue = nil
		} else {
			if err = validator.Validate(dst.ApiRelativeDateTimeValue); err != nil {
				dst.ApiRelativeDateTimeValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiRelativeDateTimeValue = nil
	}

	// try to unmarshal data into ApiStaticAppendValue
	err = newStrictDecoder(data).Decode(&dst.ApiStaticAppendValue)
	if err == nil {
		jsonApiStaticAppendValue, _ := json.Marshal(dst.ApiStaticAppendValue)
		if string(jsonApiStaticAppendValue) == "{}" { // empty struct
			dst.ApiStaticAppendValue = nil
		} else {
			if err = validator.Validate(dst.ApiStaticAppendValue); err != nil {
				dst.ApiStaticAppendValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiStaticAppendValue = nil
	}

	// try to unmarshal data into ApiStaticValue
	err = newStrictDecoder(data).Decode(&dst.ApiStaticValue)
	if err == nil {
		jsonApiStaticValue, _ := json.Marshal(dst.ApiStaticValue)
		if string(jsonApiStaticValue) == "{}" { // empty struct
			dst.ApiStaticValue = nil
		} else {
			if err = validator.Validate(dst.ApiStaticValue); err != nil {
				dst.ApiStaticValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiStaticValue = nil
	}

	// try to unmarshal data into ApiTimestampValue
	err = newStrictDecoder(data).Decode(&dst.ApiTimestampValue)
	if err == nil {
		jsonApiTimestampValue, _ := json.Marshal(dst.ApiTimestampValue)
		if string(jsonApiTimestampValue) == "{}" { // empty struct
			dst.ApiTimestampValue = nil
		} else {
			if err = validator.Validate(dst.ApiTimestampValue); err != nil {
				dst.ApiTimestampValue = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiTimestampValue = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApiActionDataValue = nil
		dst.ApiAppendObjectPropertyValue = nil
		dst.ApiEnrollmentEventPropertyValue = nil
		dst.ApiFetchedObjectPropertyValue = nil
		dst.ApiIncrementValue = nil
		dst.ApiObjectPropertyValue = nil
		dst.ApiRelativeDateTimeValue = nil
		dst.ApiStaticAppendValue = nil
		dst.ApiStaticValue = nil
		dst.ApiTimestampValue = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApiInputVariableValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApiInputVariableValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApiInputVariableValue) MarshalJSON() ([]byte, error) {
	if src.ApiActionDataValue != nil {
		return json.Marshal(&src.ApiActionDataValue)
	}

	if src.ApiAppendObjectPropertyValue != nil {
		return json.Marshal(&src.ApiAppendObjectPropertyValue)
	}

	if src.ApiEnrollmentEventPropertyValue != nil {
		return json.Marshal(&src.ApiEnrollmentEventPropertyValue)
	}

	if src.ApiFetchedObjectPropertyValue != nil {
		return json.Marshal(&src.ApiFetchedObjectPropertyValue)
	}

	if src.ApiIncrementValue != nil {
		return json.Marshal(&src.ApiIncrementValue)
	}

	if src.ApiObjectPropertyValue != nil {
		return json.Marshal(&src.ApiObjectPropertyValue)
	}

	if src.ApiRelativeDateTimeValue != nil {
		return json.Marshal(&src.ApiRelativeDateTimeValue)
	}

	if src.ApiStaticAppendValue != nil {
		return json.Marshal(&src.ApiStaticAppendValue)
	}

	if src.ApiStaticValue != nil {
		return json.Marshal(&src.ApiStaticValue)
	}

	if src.ApiTimestampValue != nil {
		return json.Marshal(&src.ApiTimestampValue)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApiInputVariableValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApiActionDataValue != nil {
		return obj.ApiActionDataValue
	}

	if obj.ApiAppendObjectPropertyValue != nil {
		return obj.ApiAppendObjectPropertyValue
	}

	if obj.ApiEnrollmentEventPropertyValue != nil {
		return obj.ApiEnrollmentEventPropertyValue
	}

	if obj.ApiFetchedObjectPropertyValue != nil {
		return obj.ApiFetchedObjectPropertyValue
	}

	if obj.ApiIncrementValue != nil {
		return obj.ApiIncrementValue
	}

	if obj.ApiObjectPropertyValue != nil {
		return obj.ApiObjectPropertyValue
	}

	if obj.ApiRelativeDateTimeValue != nil {
		return obj.ApiRelativeDateTimeValue
	}

	if obj.ApiStaticAppendValue != nil {
		return obj.ApiStaticAppendValue
	}

	if obj.ApiStaticValue != nil {
		return obj.ApiStaticValue
	}

	if obj.ApiTimestampValue != nil {
		return obj.ApiTimestampValue
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ApiInputVariableValue) GetActualInstanceValue() (interface{}) {
	if obj.ApiActionDataValue != nil {
		return *obj.ApiActionDataValue
	}

	if obj.ApiAppendObjectPropertyValue != nil {
		return *obj.ApiAppendObjectPropertyValue
	}

	if obj.ApiEnrollmentEventPropertyValue != nil {
		return *obj.ApiEnrollmentEventPropertyValue
	}

	if obj.ApiFetchedObjectPropertyValue != nil {
		return *obj.ApiFetchedObjectPropertyValue
	}

	if obj.ApiIncrementValue != nil {
		return *obj.ApiIncrementValue
	}

	if obj.ApiObjectPropertyValue != nil {
		return *obj.ApiObjectPropertyValue
	}

	if obj.ApiRelativeDateTimeValue != nil {
		return *obj.ApiRelativeDateTimeValue
	}

	if obj.ApiStaticAppendValue != nil {
		return *obj.ApiStaticAppendValue
	}

	if obj.ApiStaticValue != nil {
		return *obj.ApiStaticValue
	}

	if obj.ApiTimestampValue != nil {
		return *obj.ApiTimestampValue
	}

	// all schemas are nil
	return nil
}

type NullableApiInputVariableValue struct {
	value *ApiInputVariableValue
	isSet bool
}

func (v NullableApiInputVariableValue) Get() *ApiInputVariableValue {
	return v.value
}

func (v *NullableApiInputVariableValue) Set(val *ApiInputVariableValue) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInputVariableValue) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInputVariableValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInputVariableValue(val *ApiInputVariableValue) *NullableApiInputVariableValue {
	return &NullableApiInputVariableValue{value: val, isSet: true}
}

func (v NullableApiInputVariableValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInputVariableValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


