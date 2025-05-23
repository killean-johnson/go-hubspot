/*
Automation V4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package automation_v4

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiCustomCodeAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCustomCodeAction{}

// ApiCustomCodeAction struct for ApiCustomCodeAction
type ApiCustomCodeAction struct {
	InputFields []ApiInputVariable `json:"inputFields"`
	// The list of output fields that this custom action makes available to the rest of the flow.
	OutputFields []ApiCustomCodeActionOutputFieldsInner `json:"outputFields"`
	// The source code to execute when this action executes.
	SourceCode string `json:"sourceCode"`
	// The ID for this action.
	ActionId string `json:"actionId"`
	// The runtime to use to execute the source code. Supported runtimes are: \"NODE16X\", \"NODE20X\", \"PYTHON39\"
	Runtime string `json:"runtime"`
	Connection *ApiConnection `json:"connection,omitempty"`
	// The names of any \"secrets\" setup in this portal that will be used in this action.
	SecretNames []string `json:"secretNames"`
	// The type of action this is, can be: \"STATIC_BRANCH\", \"LIST_BRANCH\", \"AB_TEST_BRANCH\", \"CUSTOM_CODE\", \"WEBHOOK\", or \"SINGLE_CONNECTION\"
	Type string `json:"type"`
}

type _ApiCustomCodeAction ApiCustomCodeAction

// NewApiCustomCodeAction instantiates a new ApiCustomCodeAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCustomCodeAction(inputFields []ApiInputVariable, outputFields []ApiCustomCodeActionOutputFieldsInner, sourceCode string, actionId string, runtime string, secretNames []string, type_ string) *ApiCustomCodeAction {
	this := ApiCustomCodeAction{}
	this.InputFields = inputFields
	this.OutputFields = outputFields
	this.SourceCode = sourceCode
	this.ActionId = actionId
	this.Runtime = runtime
	this.SecretNames = secretNames
	this.Type = type_
	return &this
}

// NewApiCustomCodeActionWithDefaults instantiates a new ApiCustomCodeAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCustomCodeActionWithDefaults() *ApiCustomCodeAction {
	this := ApiCustomCodeAction{}
	var type_ string = "CUSTOM_CODE"
	this.Type = type_
	return &this
}

// GetInputFields returns the InputFields field value
func (o *ApiCustomCodeAction) GetInputFields() []ApiInputVariable {
	if o == nil {
		var ret []ApiInputVariable
		return ret
	}

	return o.InputFields
}

// GetInputFieldsOk returns a tuple with the InputFields field value
// and a boolean to check if the value has been set.
func (o *ApiCustomCodeAction) GetInputFieldsOk() ([]ApiInputVariable, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputFields, true
}

// SetInputFields sets field value
func (o *ApiCustomCodeAction) SetInputFields(v []ApiInputVariable) {
	o.InputFields = v
}

// GetOutputFields returns the OutputFields field value
func (o *ApiCustomCodeAction) GetOutputFields() []ApiCustomCodeActionOutputFieldsInner {
	if o == nil {
		var ret []ApiCustomCodeActionOutputFieldsInner
		return ret
	}

	return o.OutputFields
}

// GetOutputFieldsOk returns a tuple with the OutputFields field value
// and a boolean to check if the value has been set.
func (o *ApiCustomCodeAction) GetOutputFieldsOk() ([]ApiCustomCodeActionOutputFieldsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutputFields, true
}

// SetOutputFields sets field value
func (o *ApiCustomCodeAction) SetOutputFields(v []ApiCustomCodeActionOutputFieldsInner) {
	o.OutputFields = v
}

// GetSourceCode returns the SourceCode field value
func (o *ApiCustomCodeAction) GetSourceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceCode
}

// GetSourceCodeOk returns a tuple with the SourceCode field value
// and a boolean to check if the value has been set.
func (o *ApiCustomCodeAction) GetSourceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceCode, true
}

// SetSourceCode sets field value
func (o *ApiCustomCodeAction) SetSourceCode(v string) {
	o.SourceCode = v
}

// GetActionId returns the ActionId field value
func (o *ApiCustomCodeAction) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *ApiCustomCodeAction) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value
func (o *ApiCustomCodeAction) SetActionId(v string) {
	o.ActionId = v
}

// GetRuntime returns the Runtime field value
func (o *ApiCustomCodeAction) GetRuntime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value
// and a boolean to check if the value has been set.
func (o *ApiCustomCodeAction) GetRuntimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Runtime, true
}

// SetRuntime sets field value
func (o *ApiCustomCodeAction) SetRuntime(v string) {
	o.Runtime = v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *ApiCustomCodeAction) GetConnection() ApiConnection {
	if o == nil || IsNil(o.Connection) {
		var ret ApiConnection
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCustomCodeAction) GetConnectionOk() (*ApiConnection, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *ApiCustomCodeAction) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ApiConnection and assigns it to the Connection field.
func (o *ApiCustomCodeAction) SetConnection(v ApiConnection) {
	o.Connection = &v
}

// GetSecretNames returns the SecretNames field value
func (o *ApiCustomCodeAction) GetSecretNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SecretNames
}

// GetSecretNamesOk returns a tuple with the SecretNames field value
// and a boolean to check if the value has been set.
func (o *ApiCustomCodeAction) GetSecretNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretNames, true
}

// SetSecretNames sets field value
func (o *ApiCustomCodeAction) SetSecretNames(v []string) {
	o.SecretNames = v
}

// GetType returns the Type field value
func (o *ApiCustomCodeAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiCustomCodeAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiCustomCodeAction) SetType(v string) {
	o.Type = v
}

func (o ApiCustomCodeAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCustomCodeAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputFields"] = o.InputFields
	toSerialize["outputFields"] = o.OutputFields
	toSerialize["sourceCode"] = o.SourceCode
	toSerialize["actionId"] = o.ActionId
	toSerialize["runtime"] = o.Runtime
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	toSerialize["secretNames"] = o.SecretNames
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ApiCustomCodeAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inputFields",
		"outputFields",
		"sourceCode",
		"actionId",
		"runtime",
		"secretNames",
		"type",
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

	varApiCustomCodeAction := _ApiCustomCodeAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiCustomCodeAction)

	if err != nil {
		return err
	}

	*o = ApiCustomCodeAction(varApiCustomCodeAction)

	return err
}

type NullableApiCustomCodeAction struct {
	value *ApiCustomCodeAction
	isSet bool
}

func (v NullableApiCustomCodeAction) Get() *ApiCustomCodeAction {
	return v.value
}

func (v *NullableApiCustomCodeAction) Set(val *ApiCustomCodeAction) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCustomCodeAction) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCustomCodeAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCustomCodeAction(val *ApiCustomCodeAction) *NullableApiCustomCodeAction {
	return &NullableApiCustomCodeAction{value: val, isSet: true}
}

func (v NullableApiCustomCodeAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCustomCodeAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


