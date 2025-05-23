/*
CMS Media Bridge

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package media_bridge

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RollupExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RollupExpression{}

// RollupExpression struct for RollupExpression
type RollupExpression struct {
	AssociationTypes []AssociationSpec `json:"associationTypes"`
	EmptyRollupValue *string `json:"emptyRollupValue,omitempty"`
	ConditionalFormula *string `json:"conditionalFormula,omitempty"`
	SourceObjectTypeId string `json:"sourceObjectTypeId"`
	ConditionalExpression *OrInputsInner `json:"conditionalExpression,omitempty"`
	RollupOperator string `json:"rollupOperator"`
	SourcePropertyName string `json:"sourcePropertyName"`
	SourceCompareByPropertyName *string `json:"sourceCompareByPropertyName,omitempty"`
}

type _RollupExpression RollupExpression

// NewRollupExpression instantiates a new RollupExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollupExpression(associationTypes []AssociationSpec, sourceObjectTypeId string, rollupOperator string, sourcePropertyName string) *RollupExpression {
	this := RollupExpression{}
	this.AssociationTypes = associationTypes
	this.SourceObjectTypeId = sourceObjectTypeId
	this.RollupOperator = rollupOperator
	this.SourcePropertyName = sourcePropertyName
	return &this
}

// NewRollupExpressionWithDefaults instantiates a new RollupExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollupExpressionWithDefaults() *RollupExpression {
	this := RollupExpression{}
	return &this
}

// GetAssociationTypes returns the AssociationTypes field value
func (o *RollupExpression) GetAssociationTypes() []AssociationSpec {
	if o == nil {
		var ret []AssociationSpec
		return ret
	}

	return o.AssociationTypes
}

// GetAssociationTypesOk returns a tuple with the AssociationTypes field value
// and a boolean to check if the value has been set.
func (o *RollupExpression) GetAssociationTypesOk() ([]AssociationSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssociationTypes, true
}

// SetAssociationTypes sets field value
func (o *RollupExpression) SetAssociationTypes(v []AssociationSpec) {
	o.AssociationTypes = v
}

// GetEmptyRollupValue returns the EmptyRollupValue field value if set, zero value otherwise.
func (o *RollupExpression) GetEmptyRollupValue() string {
	if o == nil || IsNil(o.EmptyRollupValue) {
		var ret string
		return ret
	}
	return *o.EmptyRollupValue
}

// GetEmptyRollupValueOk returns a tuple with the EmptyRollupValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollupExpression) GetEmptyRollupValueOk() (*string, bool) {
	if o == nil || IsNil(o.EmptyRollupValue) {
		return nil, false
	}
	return o.EmptyRollupValue, true
}

// HasEmptyRollupValue returns a boolean if a field has been set.
func (o *RollupExpression) HasEmptyRollupValue() bool {
	if o != nil && !IsNil(o.EmptyRollupValue) {
		return true
	}

	return false
}

// SetEmptyRollupValue gets a reference to the given string and assigns it to the EmptyRollupValue field.
func (o *RollupExpression) SetEmptyRollupValue(v string) {
	o.EmptyRollupValue = &v
}

// GetConditionalFormula returns the ConditionalFormula field value if set, zero value otherwise.
func (o *RollupExpression) GetConditionalFormula() string {
	if o == nil || IsNil(o.ConditionalFormula) {
		var ret string
		return ret
	}
	return *o.ConditionalFormula
}

// GetConditionalFormulaOk returns a tuple with the ConditionalFormula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollupExpression) GetConditionalFormulaOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionalFormula) {
		return nil, false
	}
	return o.ConditionalFormula, true
}

// HasConditionalFormula returns a boolean if a field has been set.
func (o *RollupExpression) HasConditionalFormula() bool {
	if o != nil && !IsNil(o.ConditionalFormula) {
		return true
	}

	return false
}

// SetConditionalFormula gets a reference to the given string and assigns it to the ConditionalFormula field.
func (o *RollupExpression) SetConditionalFormula(v string) {
	o.ConditionalFormula = &v
}

// GetSourceObjectTypeId returns the SourceObjectTypeId field value
func (o *RollupExpression) GetSourceObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceObjectTypeId
}

// GetSourceObjectTypeIdOk returns a tuple with the SourceObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *RollupExpression) GetSourceObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceObjectTypeId, true
}

// SetSourceObjectTypeId sets field value
func (o *RollupExpression) SetSourceObjectTypeId(v string) {
	o.SourceObjectTypeId = v
}

// GetConditionalExpression returns the ConditionalExpression field value if set, zero value otherwise.
func (o *RollupExpression) GetConditionalExpression() OrInputsInner {
	if o == nil || IsNil(o.ConditionalExpression) {
		var ret OrInputsInner
		return ret
	}
	return *o.ConditionalExpression
}

// GetConditionalExpressionOk returns a tuple with the ConditionalExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollupExpression) GetConditionalExpressionOk() (*OrInputsInner, bool) {
	if o == nil || IsNil(o.ConditionalExpression) {
		return nil, false
	}
	return o.ConditionalExpression, true
}

// HasConditionalExpression returns a boolean if a field has been set.
func (o *RollupExpression) HasConditionalExpression() bool {
	if o != nil && !IsNil(o.ConditionalExpression) {
		return true
	}

	return false
}

// SetConditionalExpression gets a reference to the given OrInputsInner and assigns it to the ConditionalExpression field.
func (o *RollupExpression) SetConditionalExpression(v OrInputsInner) {
	o.ConditionalExpression = &v
}

// GetRollupOperator returns the RollupOperator field value
func (o *RollupExpression) GetRollupOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RollupOperator
}

// GetRollupOperatorOk returns a tuple with the RollupOperator field value
// and a boolean to check if the value has been set.
func (o *RollupExpression) GetRollupOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RollupOperator, true
}

// SetRollupOperator sets field value
func (o *RollupExpression) SetRollupOperator(v string) {
	o.RollupOperator = v
}

// GetSourcePropertyName returns the SourcePropertyName field value
func (o *RollupExpression) GetSourcePropertyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourcePropertyName
}

// GetSourcePropertyNameOk returns a tuple with the SourcePropertyName field value
// and a boolean to check if the value has been set.
func (o *RollupExpression) GetSourcePropertyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcePropertyName, true
}

// SetSourcePropertyName sets field value
func (o *RollupExpression) SetSourcePropertyName(v string) {
	o.SourcePropertyName = v
}

// GetSourceCompareByPropertyName returns the SourceCompareByPropertyName field value if set, zero value otherwise.
func (o *RollupExpression) GetSourceCompareByPropertyName() string {
	if o == nil || IsNil(o.SourceCompareByPropertyName) {
		var ret string
		return ret
	}
	return *o.SourceCompareByPropertyName
}

// GetSourceCompareByPropertyNameOk returns a tuple with the SourceCompareByPropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollupExpression) GetSourceCompareByPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCompareByPropertyName) {
		return nil, false
	}
	return o.SourceCompareByPropertyName, true
}

// HasSourceCompareByPropertyName returns a boolean if a field has been set.
func (o *RollupExpression) HasSourceCompareByPropertyName() bool {
	if o != nil && !IsNil(o.SourceCompareByPropertyName) {
		return true
	}

	return false
}

// SetSourceCompareByPropertyName gets a reference to the given string and assigns it to the SourceCompareByPropertyName field.
func (o *RollupExpression) SetSourceCompareByPropertyName(v string) {
	o.SourceCompareByPropertyName = &v
}

func (o RollupExpression) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollupExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["associationTypes"] = o.AssociationTypes
	if !IsNil(o.EmptyRollupValue) {
		toSerialize["emptyRollupValue"] = o.EmptyRollupValue
	}
	if !IsNil(o.ConditionalFormula) {
		toSerialize["conditionalFormula"] = o.ConditionalFormula
	}
	toSerialize["sourceObjectTypeId"] = o.SourceObjectTypeId
	if !IsNil(o.ConditionalExpression) {
		toSerialize["conditionalExpression"] = o.ConditionalExpression
	}
	toSerialize["rollupOperator"] = o.RollupOperator
	toSerialize["sourcePropertyName"] = o.SourcePropertyName
	if !IsNil(o.SourceCompareByPropertyName) {
		toSerialize["sourceCompareByPropertyName"] = o.SourceCompareByPropertyName
	}
	return toSerialize, nil
}

func (o *RollupExpression) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"associationTypes",
		"sourceObjectTypeId",
		"rollupOperator",
		"sourcePropertyName",
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

	varRollupExpression := _RollupExpression{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRollupExpression)

	if err != nil {
		return err
	}

	*o = RollupExpression(varRollupExpression)

	return err
}

type NullableRollupExpression struct {
	value *RollupExpression
	isSet bool
}

func (v NullableRollupExpression) Get() *RollupExpression {
	return v.value
}

func (v *NullableRollupExpression) Set(val *RollupExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableRollupExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableRollupExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollupExpression(val *RollupExpression) *NullableRollupExpression {
	return &NullableRollupExpression{value: val, isSet: true}
}

func (v NullableRollupExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollupExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


