/*
Events Manage Event Definitions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package manage_event_definitions

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EnumerationPropertyOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnumerationPropertyOperation{}

// EnumerationPropertyOperation struct for EnumerationPropertyOperation
type EnumerationPropertyOperation struct {
	IncludeObjectsWithNoValueSet bool `json:"includeObjectsWithNoValueSet"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	PropertyType string `json:"propertyType"`
	Values []string `json:"values"`
	OperationType string `json:"operationType"`
	OperatorName string `json:"operatorName"`
	Operator string `json:"operator"`
}

type _EnumerationPropertyOperation EnumerationPropertyOperation

// NewEnumerationPropertyOperation instantiates a new EnumerationPropertyOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnumerationPropertyOperation(includeObjectsWithNoValueSet bool, propertyType string, values []string, operationType string, operatorName string, operator string) *EnumerationPropertyOperation {
	this := EnumerationPropertyOperation{}
	this.IncludeObjectsWithNoValueSet = includeObjectsWithNoValueSet
	this.PropertyType = propertyType
	this.Values = values
	this.OperationType = operationType
	this.OperatorName = operatorName
	this.Operator = operator
	return &this
}

// NewEnumerationPropertyOperationWithDefaults instantiates a new EnumerationPropertyOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnumerationPropertyOperationWithDefaults() *EnumerationPropertyOperation {
	this := EnumerationPropertyOperation{}
	var propertyType string = "enumeration"
	this.PropertyType = propertyType
	return &this
}

// GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field value
func (o *EnumerationPropertyOperation) GetIncludeObjectsWithNoValueSet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeObjectsWithNoValueSet
}

// GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field value
// and a boolean to check if the value has been set.
func (o *EnumerationPropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeObjectsWithNoValueSet, true
}

// SetIncludeObjectsWithNoValueSet sets field value
func (o *EnumerationPropertyOperation) SetIncludeObjectsWithNoValueSet(v bool) {
	o.IncludeObjectsWithNoValueSet = v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *EnumerationPropertyOperation) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumerationPropertyOperation) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *EnumerationPropertyOperation) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *EnumerationPropertyOperation) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetPropertyType returns the PropertyType field value
func (o *EnumerationPropertyOperation) GetPropertyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropertyType
}

// GetPropertyTypeOk returns a tuple with the PropertyType field value
// and a boolean to check if the value has been set.
func (o *EnumerationPropertyOperation) GetPropertyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyType, true
}

// SetPropertyType sets field value
func (o *EnumerationPropertyOperation) SetPropertyType(v string) {
	o.PropertyType = v
}

// GetValues returns the Values field value
func (o *EnumerationPropertyOperation) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *EnumerationPropertyOperation) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *EnumerationPropertyOperation) SetValues(v []string) {
	o.Values = v
}

// GetOperationType returns the OperationType field value
func (o *EnumerationPropertyOperation) GetOperationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value
// and a boolean to check if the value has been set.
func (o *EnumerationPropertyOperation) GetOperationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// SetOperationType sets field value
func (o *EnumerationPropertyOperation) SetOperationType(v string) {
	o.OperationType = v
}

// GetOperatorName returns the OperatorName field value
func (o *EnumerationPropertyOperation) GetOperatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatorName
}

// GetOperatorNameOk returns a tuple with the OperatorName field value
// and a boolean to check if the value has been set.
func (o *EnumerationPropertyOperation) GetOperatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatorName, true
}

// SetOperatorName sets field value
func (o *EnumerationPropertyOperation) SetOperatorName(v string) {
	o.OperatorName = v
}

// GetOperator returns the Operator field value
func (o *EnumerationPropertyOperation) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *EnumerationPropertyOperation) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *EnumerationPropertyOperation) SetOperator(v string) {
	o.Operator = v
}

func (o EnumerationPropertyOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnumerationPropertyOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["includeObjectsWithNoValueSet"] = o.IncludeObjectsWithNoValueSet
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	toSerialize["propertyType"] = o.PropertyType
	toSerialize["values"] = o.Values
	toSerialize["operationType"] = o.OperationType
	toSerialize["operatorName"] = o.OperatorName
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *EnumerationPropertyOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"includeObjectsWithNoValueSet",
		"propertyType",
		"values",
		"operationType",
		"operatorName",
		"operator",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEnumerationPropertyOperation := _EnumerationPropertyOperation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnumerationPropertyOperation)

	if err != nil {
		return err
	}

	*o = EnumerationPropertyOperation(varEnumerationPropertyOperation)

	return err
}

type NullableEnumerationPropertyOperation struct {
	value *EnumerationPropertyOperation
	isSet bool
}

func (v NullableEnumerationPropertyOperation) Get() *EnumerationPropertyOperation {
	return v.value
}

func (v *NullableEnumerationPropertyOperation) Set(val *EnumerationPropertyOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumerationPropertyOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumerationPropertyOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumerationPropertyOperation(val *EnumerationPropertyOperation) *NullableEnumerationPropertyOperation {
	return &NullableEnumerationPropertyOperation{value: val, isSet: true}
}

func (v NullableEnumerationPropertyOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumerationPropertyOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


