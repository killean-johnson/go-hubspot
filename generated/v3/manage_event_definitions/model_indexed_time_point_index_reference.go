/*
Events Manage Event Definitions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package manage_event_definitions

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// IndexedTimePointIndexReference - struct for IndexedTimePointIndexReference
type IndexedTimePointIndexReference struct {
	FiscalQuarter *FiscalQuarter
	FiscalYear *FiscalYear
	MonthReference *MonthReference
	NowReference *NowReference
	QuarterReference *QuarterReference
	TodayReference *TodayReference
	WeekReference *WeekReference
	YearReference *YearReference
}

// FiscalQuarterAsIndexedTimePointIndexReference is a convenience function that returns FiscalQuarter wrapped in IndexedTimePointIndexReference
func FiscalQuarterAsIndexedTimePointIndexReference(v *FiscalQuarter) IndexedTimePointIndexReference {
	return IndexedTimePointIndexReference{
		FiscalQuarter: v,
	}
}

// FiscalYearAsIndexedTimePointIndexReference is a convenience function that returns FiscalYear wrapped in IndexedTimePointIndexReference
func FiscalYearAsIndexedTimePointIndexReference(v *FiscalYear) IndexedTimePointIndexReference {
	return IndexedTimePointIndexReference{
		FiscalYear: v,
	}
}

// MonthReferenceAsIndexedTimePointIndexReference is a convenience function that returns MonthReference wrapped in IndexedTimePointIndexReference
func MonthReferenceAsIndexedTimePointIndexReference(v *MonthReference) IndexedTimePointIndexReference {
	return IndexedTimePointIndexReference{
		MonthReference: v,
	}
}

// NowReferenceAsIndexedTimePointIndexReference is a convenience function that returns NowReference wrapped in IndexedTimePointIndexReference
func NowReferenceAsIndexedTimePointIndexReference(v *NowReference) IndexedTimePointIndexReference {
	return IndexedTimePointIndexReference{
		NowReference: v,
	}
}

// QuarterReferenceAsIndexedTimePointIndexReference is a convenience function that returns QuarterReference wrapped in IndexedTimePointIndexReference
func QuarterReferenceAsIndexedTimePointIndexReference(v *QuarterReference) IndexedTimePointIndexReference {
	return IndexedTimePointIndexReference{
		QuarterReference: v,
	}
}

// TodayReferenceAsIndexedTimePointIndexReference is a convenience function that returns TodayReference wrapped in IndexedTimePointIndexReference
func TodayReferenceAsIndexedTimePointIndexReference(v *TodayReference) IndexedTimePointIndexReference {
	return IndexedTimePointIndexReference{
		TodayReference: v,
	}
}

// WeekReferenceAsIndexedTimePointIndexReference is a convenience function that returns WeekReference wrapped in IndexedTimePointIndexReference
func WeekReferenceAsIndexedTimePointIndexReference(v *WeekReference) IndexedTimePointIndexReference {
	return IndexedTimePointIndexReference{
		WeekReference: v,
	}
}

// YearReferenceAsIndexedTimePointIndexReference is a convenience function that returns YearReference wrapped in IndexedTimePointIndexReference
func YearReferenceAsIndexedTimePointIndexReference(v *YearReference) IndexedTimePointIndexReference {
	return IndexedTimePointIndexReference{
		YearReference: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IndexedTimePointIndexReference) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FiscalQuarter
	err = newStrictDecoder(data).Decode(&dst.FiscalQuarter)
	if err == nil {
		jsonFiscalQuarter, _ := json.Marshal(dst.FiscalQuarter)
		if string(jsonFiscalQuarter) == "{}" { // empty struct
			dst.FiscalQuarter = nil
		} else {
			if err = validator.Validate(dst.FiscalQuarter); err != nil {
				dst.FiscalQuarter = nil
			} else {
				match++
			}
		}
	} else {
		dst.FiscalQuarter = nil
	}

	// try to unmarshal data into FiscalYear
	err = newStrictDecoder(data).Decode(&dst.FiscalYear)
	if err == nil {
		jsonFiscalYear, _ := json.Marshal(dst.FiscalYear)
		if string(jsonFiscalYear) == "{}" { // empty struct
			dst.FiscalYear = nil
		} else {
			if err = validator.Validate(dst.FiscalYear); err != nil {
				dst.FiscalYear = nil
			} else {
				match++
			}
		}
	} else {
		dst.FiscalYear = nil
	}

	// try to unmarshal data into MonthReference
	err = newStrictDecoder(data).Decode(&dst.MonthReference)
	if err == nil {
		jsonMonthReference, _ := json.Marshal(dst.MonthReference)
		if string(jsonMonthReference) == "{}" { // empty struct
			dst.MonthReference = nil
		} else {
			if err = validator.Validate(dst.MonthReference); err != nil {
				dst.MonthReference = nil
			} else {
				match++
			}
		}
	} else {
		dst.MonthReference = nil
	}

	// try to unmarshal data into NowReference
	err = newStrictDecoder(data).Decode(&dst.NowReference)
	if err == nil {
		jsonNowReference, _ := json.Marshal(dst.NowReference)
		if string(jsonNowReference) == "{}" { // empty struct
			dst.NowReference = nil
		} else {
			if err = validator.Validate(dst.NowReference); err != nil {
				dst.NowReference = nil
			} else {
				match++
			}
		}
	} else {
		dst.NowReference = nil
	}

	// try to unmarshal data into QuarterReference
	err = newStrictDecoder(data).Decode(&dst.QuarterReference)
	if err == nil {
		jsonQuarterReference, _ := json.Marshal(dst.QuarterReference)
		if string(jsonQuarterReference) == "{}" { // empty struct
			dst.QuarterReference = nil
		} else {
			if err = validator.Validate(dst.QuarterReference); err != nil {
				dst.QuarterReference = nil
			} else {
				match++
			}
		}
	} else {
		dst.QuarterReference = nil
	}

	// try to unmarshal data into TodayReference
	err = newStrictDecoder(data).Decode(&dst.TodayReference)
	if err == nil {
		jsonTodayReference, _ := json.Marshal(dst.TodayReference)
		if string(jsonTodayReference) == "{}" { // empty struct
			dst.TodayReference = nil
		} else {
			if err = validator.Validate(dst.TodayReference); err != nil {
				dst.TodayReference = nil
			} else {
				match++
			}
		}
	} else {
		dst.TodayReference = nil
	}

	// try to unmarshal data into WeekReference
	err = newStrictDecoder(data).Decode(&dst.WeekReference)
	if err == nil {
		jsonWeekReference, _ := json.Marshal(dst.WeekReference)
		if string(jsonWeekReference) == "{}" { // empty struct
			dst.WeekReference = nil
		} else {
			if err = validator.Validate(dst.WeekReference); err != nil {
				dst.WeekReference = nil
			} else {
				match++
			}
		}
	} else {
		dst.WeekReference = nil
	}

	// try to unmarshal data into YearReference
	err = newStrictDecoder(data).Decode(&dst.YearReference)
	if err == nil {
		jsonYearReference, _ := json.Marshal(dst.YearReference)
		if string(jsonYearReference) == "{}" { // empty struct
			dst.YearReference = nil
		} else {
			if err = validator.Validate(dst.YearReference); err != nil {
				dst.YearReference = nil
			} else {
				match++
			}
		}
	} else {
		dst.YearReference = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FiscalQuarter = nil
		dst.FiscalYear = nil
		dst.MonthReference = nil
		dst.NowReference = nil
		dst.QuarterReference = nil
		dst.TodayReference = nil
		dst.WeekReference = nil
		dst.YearReference = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IndexedTimePointIndexReference)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IndexedTimePointIndexReference)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IndexedTimePointIndexReference) MarshalJSON() ([]byte, error) {
	if src.FiscalQuarter != nil {
		return json.Marshal(&src.FiscalQuarter)
	}

	if src.FiscalYear != nil {
		return json.Marshal(&src.FiscalYear)
	}

	if src.MonthReference != nil {
		return json.Marshal(&src.MonthReference)
	}

	if src.NowReference != nil {
		return json.Marshal(&src.NowReference)
	}

	if src.QuarterReference != nil {
		return json.Marshal(&src.QuarterReference)
	}

	if src.TodayReference != nil {
		return json.Marshal(&src.TodayReference)
	}

	if src.WeekReference != nil {
		return json.Marshal(&src.WeekReference)
	}

	if src.YearReference != nil {
		return json.Marshal(&src.YearReference)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IndexedTimePointIndexReference) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.FiscalQuarter != nil {
		return obj.FiscalQuarter
	}

	if obj.FiscalYear != nil {
		return obj.FiscalYear
	}

	if obj.MonthReference != nil {
		return obj.MonthReference
	}

	if obj.NowReference != nil {
		return obj.NowReference
	}

	if obj.QuarterReference != nil {
		return obj.QuarterReference
	}

	if obj.TodayReference != nil {
		return obj.TodayReference
	}

	if obj.WeekReference != nil {
		return obj.WeekReference
	}

	if obj.YearReference != nil {
		return obj.YearReference
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj IndexedTimePointIndexReference) GetActualInstanceValue() (interface{}) {
	if obj.FiscalQuarter != nil {
		return *obj.FiscalQuarter
	}

	if obj.FiscalYear != nil {
		return *obj.FiscalYear
	}

	if obj.MonthReference != nil {
		return *obj.MonthReference
	}

	if obj.NowReference != nil {
		return *obj.NowReference
	}

	if obj.QuarterReference != nil {
		return *obj.QuarterReference
	}

	if obj.TodayReference != nil {
		return *obj.TodayReference
	}

	if obj.WeekReference != nil {
		return *obj.WeekReference
	}

	if obj.YearReference != nil {
		return *obj.YearReference
	}

	// all schemas are nil
	return nil
}

type NullableIndexedTimePointIndexReference struct {
	value *IndexedTimePointIndexReference
	isSet bool
}

func (v NullableIndexedTimePointIndexReference) Get() *IndexedTimePointIndexReference {
	return v.value
}

func (v *NullableIndexedTimePointIndexReference) Set(val *IndexedTimePointIndexReference) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexedTimePointIndexReference) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexedTimePointIndexReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexedTimePointIndexReference(val *IndexedTimePointIndexReference) *NullableIndexedTimePointIndexReference {
	return &NullableIndexedTimePointIndexReference{value: val, isSet: true}
}

func (v NullableIndexedTimePointIndexReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexedTimePointIndexReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


