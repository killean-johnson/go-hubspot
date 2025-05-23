/*
Forms

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package forms

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HubSpotFormConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HubSpotFormConfiguration{}

// HubSpotFormConfiguration struct for HubSpotFormConfiguration
type HubSpotFormConfiguration struct {
	// Whether to create a new contact when a form is submitted with an email address that doesn’t match any in your existing contacts records.
	CreateNewContactForNewEmail bool `json:"createNewContactForNewEmail"`
	// Whether the form can be edited.
	Editable bool `json:"editable"`
	// Whether to add a reset link to the form. This removes any pre-populated content on the form and creates a new contact on submission.
	AllowLinkToResetKnownValues bool `json:"allowLinkToResetKnownValues"`
	LifecycleStages []LifecycleStage `json:"lifecycleStages,omitempty"`
	PostSubmitAction FormPostSubmitAction `json:"postSubmitAction"`
	// The language of the form.
	Language string `json:"language"`
	// Whether contact fields should pre-populate with known information when a contact returns to your site.
	PrePopulateKnownValues bool `json:"prePopulateKnownValues"`
	// Whether the form can be cloned.
	Cloneable bool `json:"cloneable"`
	// Whether to send a notification email to the contact owner when a submission is received.
	NotifyContactOwner bool `json:"notifyContactOwner"`
	// Whether CAPTCHA (spam prevention) is enabled.
	RecaptchaEnabled bool `json:"recaptchaEnabled"`
	// Whether the form can be archived.
	Archivable bool `json:"archivable"`
	// The list of user IDs to receive a notification email when a submission is received.
	NotifyRecipients []string `json:"notifyRecipients"`
}

type _HubSpotFormConfiguration HubSpotFormConfiguration

// NewHubSpotFormConfiguration instantiates a new HubSpotFormConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHubSpotFormConfiguration(createNewContactForNewEmail bool, editable bool, allowLinkToResetKnownValues bool, postSubmitAction FormPostSubmitAction, language string, prePopulateKnownValues bool, cloneable bool, notifyContactOwner bool, recaptchaEnabled bool, archivable bool, notifyRecipients []string) *HubSpotFormConfiguration {
	this := HubSpotFormConfiguration{}
	this.CreateNewContactForNewEmail = createNewContactForNewEmail
	this.Editable = editable
	this.AllowLinkToResetKnownValues = allowLinkToResetKnownValues
	this.PostSubmitAction = postSubmitAction
	this.Language = language
	this.PrePopulateKnownValues = prePopulateKnownValues
	this.Cloneable = cloneable
	this.NotifyContactOwner = notifyContactOwner
	this.RecaptchaEnabled = recaptchaEnabled
	this.Archivable = archivable
	this.NotifyRecipients = notifyRecipients
	return &this
}

// NewHubSpotFormConfigurationWithDefaults instantiates a new HubSpotFormConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHubSpotFormConfigurationWithDefaults() *HubSpotFormConfiguration {
	this := HubSpotFormConfiguration{}
	return &this
}

// GetCreateNewContactForNewEmail returns the CreateNewContactForNewEmail field value
func (o *HubSpotFormConfiguration) GetCreateNewContactForNewEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CreateNewContactForNewEmail
}

// GetCreateNewContactForNewEmailOk returns a tuple with the CreateNewContactForNewEmail field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetCreateNewContactForNewEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateNewContactForNewEmail, true
}

// SetCreateNewContactForNewEmail sets field value
func (o *HubSpotFormConfiguration) SetCreateNewContactForNewEmail(v bool) {
	o.CreateNewContactForNewEmail = v
}

// GetEditable returns the Editable field value
func (o *HubSpotFormConfiguration) GetEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Editable
}

// GetEditableOk returns a tuple with the Editable field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editable, true
}

// SetEditable sets field value
func (o *HubSpotFormConfiguration) SetEditable(v bool) {
	o.Editable = v
}

// GetAllowLinkToResetKnownValues returns the AllowLinkToResetKnownValues field value
func (o *HubSpotFormConfiguration) GetAllowLinkToResetKnownValues() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowLinkToResetKnownValues
}

// GetAllowLinkToResetKnownValuesOk returns a tuple with the AllowLinkToResetKnownValues field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetAllowLinkToResetKnownValuesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowLinkToResetKnownValues, true
}

// SetAllowLinkToResetKnownValues sets field value
func (o *HubSpotFormConfiguration) SetAllowLinkToResetKnownValues(v bool) {
	o.AllowLinkToResetKnownValues = v
}

// GetLifecycleStages returns the LifecycleStages field value if set, zero value otherwise.
func (o *HubSpotFormConfiguration) GetLifecycleStages() []LifecycleStage {
	if o == nil || IsNil(o.LifecycleStages) {
		var ret []LifecycleStage
		return ret
	}
	return o.LifecycleStages
}

// GetLifecycleStagesOk returns a tuple with the LifecycleStages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetLifecycleStagesOk() ([]LifecycleStage, bool) {
	if o == nil || IsNil(o.LifecycleStages) {
		return nil, false
	}
	return o.LifecycleStages, true
}

// HasLifecycleStages returns a boolean if a field has been set.
func (o *HubSpotFormConfiguration) HasLifecycleStages() bool {
	if o != nil && !IsNil(o.LifecycleStages) {
		return true
	}

	return false
}

// SetLifecycleStages gets a reference to the given []LifecycleStage and assigns it to the LifecycleStages field.
func (o *HubSpotFormConfiguration) SetLifecycleStages(v []LifecycleStage) {
	o.LifecycleStages = v
}

// GetPostSubmitAction returns the PostSubmitAction field value
func (o *HubSpotFormConfiguration) GetPostSubmitAction() FormPostSubmitAction {
	if o == nil {
		var ret FormPostSubmitAction
		return ret
	}

	return o.PostSubmitAction
}

// GetPostSubmitActionOk returns a tuple with the PostSubmitAction field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetPostSubmitActionOk() (*FormPostSubmitAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostSubmitAction, true
}

// SetPostSubmitAction sets field value
func (o *HubSpotFormConfiguration) SetPostSubmitAction(v FormPostSubmitAction) {
	o.PostSubmitAction = v
}

// GetLanguage returns the Language field value
func (o *HubSpotFormConfiguration) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *HubSpotFormConfiguration) SetLanguage(v string) {
	o.Language = v
}

// GetPrePopulateKnownValues returns the PrePopulateKnownValues field value
func (o *HubSpotFormConfiguration) GetPrePopulateKnownValues() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PrePopulateKnownValues
}

// GetPrePopulateKnownValuesOk returns a tuple with the PrePopulateKnownValues field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetPrePopulateKnownValuesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrePopulateKnownValues, true
}

// SetPrePopulateKnownValues sets field value
func (o *HubSpotFormConfiguration) SetPrePopulateKnownValues(v bool) {
	o.PrePopulateKnownValues = v
}

// GetCloneable returns the Cloneable field value
func (o *HubSpotFormConfiguration) GetCloneable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cloneable
}

// GetCloneableOk returns a tuple with the Cloneable field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetCloneableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloneable, true
}

// SetCloneable sets field value
func (o *HubSpotFormConfiguration) SetCloneable(v bool) {
	o.Cloneable = v
}

// GetNotifyContactOwner returns the NotifyContactOwner field value
func (o *HubSpotFormConfiguration) GetNotifyContactOwner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NotifyContactOwner
}

// GetNotifyContactOwnerOk returns a tuple with the NotifyContactOwner field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetNotifyContactOwnerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyContactOwner, true
}

// SetNotifyContactOwner sets field value
func (o *HubSpotFormConfiguration) SetNotifyContactOwner(v bool) {
	o.NotifyContactOwner = v
}

// GetRecaptchaEnabled returns the RecaptchaEnabled field value
func (o *HubSpotFormConfiguration) GetRecaptchaEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecaptchaEnabled
}

// GetRecaptchaEnabledOk returns a tuple with the RecaptchaEnabled field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetRecaptchaEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecaptchaEnabled, true
}

// SetRecaptchaEnabled sets field value
func (o *HubSpotFormConfiguration) SetRecaptchaEnabled(v bool) {
	o.RecaptchaEnabled = v
}

// GetArchivable returns the Archivable field value
func (o *HubSpotFormConfiguration) GetArchivable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archivable
}

// GetArchivableOk returns a tuple with the Archivable field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetArchivableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archivable, true
}

// SetArchivable sets field value
func (o *HubSpotFormConfiguration) SetArchivable(v bool) {
	o.Archivable = v
}

// GetNotifyRecipients returns the NotifyRecipients field value
func (o *HubSpotFormConfiguration) GetNotifyRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NotifyRecipients
}

// GetNotifyRecipientsOk returns a tuple with the NotifyRecipients field value
// and a boolean to check if the value has been set.
func (o *HubSpotFormConfiguration) GetNotifyRecipientsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifyRecipients, true
}

// SetNotifyRecipients sets field value
func (o *HubSpotFormConfiguration) SetNotifyRecipients(v []string) {
	o.NotifyRecipients = v
}

func (o HubSpotFormConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HubSpotFormConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createNewContactForNewEmail"] = o.CreateNewContactForNewEmail
	toSerialize["editable"] = o.Editable
	toSerialize["allowLinkToResetKnownValues"] = o.AllowLinkToResetKnownValues
	if !IsNil(o.LifecycleStages) {
		toSerialize["lifecycleStages"] = o.LifecycleStages
	}
	toSerialize["postSubmitAction"] = o.PostSubmitAction
	toSerialize["language"] = o.Language
	toSerialize["prePopulateKnownValues"] = o.PrePopulateKnownValues
	toSerialize["cloneable"] = o.Cloneable
	toSerialize["notifyContactOwner"] = o.NotifyContactOwner
	toSerialize["recaptchaEnabled"] = o.RecaptchaEnabled
	toSerialize["archivable"] = o.Archivable
	toSerialize["notifyRecipients"] = o.NotifyRecipients
	return toSerialize, nil
}

func (o *HubSpotFormConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createNewContactForNewEmail",
		"editable",
		"allowLinkToResetKnownValues",
		"postSubmitAction",
		"language",
		"prePopulateKnownValues",
		"cloneable",
		"notifyContactOwner",
		"recaptchaEnabled",
		"archivable",
		"notifyRecipients",
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

	varHubSpotFormConfiguration := _HubSpotFormConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHubSpotFormConfiguration)

	if err != nil {
		return err
	}

	*o = HubSpotFormConfiguration(varHubSpotFormConfiguration)

	return err
}

type NullableHubSpotFormConfiguration struct {
	value *HubSpotFormConfiguration
	isSet bool
}

func (v NullableHubSpotFormConfiguration) Get() *HubSpotFormConfiguration {
	return v.value
}

func (v *NullableHubSpotFormConfiguration) Set(val *HubSpotFormConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableHubSpotFormConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableHubSpotFormConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHubSpotFormConfiguration(val *HubSpotFormConfiguration) *NullableHubSpotFormConfiguration {
	return &NullableHubSpotFormConfiguration{value: val, isSet: true}
}

func (v NullableHubSpotFormConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHubSpotFormConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


