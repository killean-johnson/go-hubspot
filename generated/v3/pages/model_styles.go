/*
Pages

Use these endpoints for interacting with Landing Pages and Site Pages

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pages

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Styles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Styles{}

// Styles struct for Styles
type Styles struct {
	BackgroundColor RGBAColor `json:"backgroundColor"`
	// 
	FlexboxPositioning string `json:"flexboxPositioning"`
	BackgroundImage BackgroundImage `json:"backgroundImage"`
	// 
	ForceFullWidthSection bool `json:"forceFullWidthSection"`
	// 
	VerticalAlignment string `json:"verticalAlignment"`
	// 
	MaxWidthSectionCentering int32 `json:"maxWidthSectionCentering"`
	BackgroundGradient Gradient `json:"backgroundGradient"`
}

type _Styles Styles

// NewStyles instantiates a new Styles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStyles(backgroundColor RGBAColor, flexboxPositioning string, backgroundImage BackgroundImage, forceFullWidthSection bool, verticalAlignment string, maxWidthSectionCentering int32, backgroundGradient Gradient) *Styles {
	this := Styles{}
	this.BackgroundColor = backgroundColor
	this.FlexboxPositioning = flexboxPositioning
	this.BackgroundImage = backgroundImage
	this.ForceFullWidthSection = forceFullWidthSection
	this.VerticalAlignment = verticalAlignment
	this.MaxWidthSectionCentering = maxWidthSectionCentering
	this.BackgroundGradient = backgroundGradient
	return &this
}

// NewStylesWithDefaults instantiates a new Styles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStylesWithDefaults() *Styles {
	this := Styles{}
	return &this
}

// GetBackgroundColor returns the BackgroundColor field value
func (o *Styles) GetBackgroundColor() RGBAColor {
	if o == nil {
		var ret RGBAColor
		return ret
	}

	return o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value
// and a boolean to check if the value has been set.
func (o *Styles) GetBackgroundColorOk() (*RGBAColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackgroundColor, true
}

// SetBackgroundColor sets field value
func (o *Styles) SetBackgroundColor(v RGBAColor) {
	o.BackgroundColor = v
}

// GetFlexboxPositioning returns the FlexboxPositioning field value
func (o *Styles) GetFlexboxPositioning() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlexboxPositioning
}

// GetFlexboxPositioningOk returns a tuple with the FlexboxPositioning field value
// and a boolean to check if the value has been set.
func (o *Styles) GetFlexboxPositioningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlexboxPositioning, true
}

// SetFlexboxPositioning sets field value
func (o *Styles) SetFlexboxPositioning(v string) {
	o.FlexboxPositioning = v
}

// GetBackgroundImage returns the BackgroundImage field value
func (o *Styles) GetBackgroundImage() BackgroundImage {
	if o == nil {
		var ret BackgroundImage
		return ret
	}

	return o.BackgroundImage
}

// GetBackgroundImageOk returns a tuple with the BackgroundImage field value
// and a boolean to check if the value has been set.
func (o *Styles) GetBackgroundImageOk() (*BackgroundImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackgroundImage, true
}

// SetBackgroundImage sets field value
func (o *Styles) SetBackgroundImage(v BackgroundImage) {
	o.BackgroundImage = v
}

// GetForceFullWidthSection returns the ForceFullWidthSection field value
func (o *Styles) GetForceFullWidthSection() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForceFullWidthSection
}

// GetForceFullWidthSectionOk returns a tuple with the ForceFullWidthSection field value
// and a boolean to check if the value has been set.
func (o *Styles) GetForceFullWidthSectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForceFullWidthSection, true
}

// SetForceFullWidthSection sets field value
func (o *Styles) SetForceFullWidthSection(v bool) {
	o.ForceFullWidthSection = v
}

// GetVerticalAlignment returns the VerticalAlignment field value
func (o *Styles) GetVerticalAlignment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerticalAlignment
}

// GetVerticalAlignmentOk returns a tuple with the VerticalAlignment field value
// and a boolean to check if the value has been set.
func (o *Styles) GetVerticalAlignmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerticalAlignment, true
}

// SetVerticalAlignment sets field value
func (o *Styles) SetVerticalAlignment(v string) {
	o.VerticalAlignment = v
}

// GetMaxWidthSectionCentering returns the MaxWidthSectionCentering field value
func (o *Styles) GetMaxWidthSectionCentering() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxWidthSectionCentering
}

// GetMaxWidthSectionCenteringOk returns a tuple with the MaxWidthSectionCentering field value
// and a boolean to check if the value has been set.
func (o *Styles) GetMaxWidthSectionCenteringOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxWidthSectionCentering, true
}

// SetMaxWidthSectionCentering sets field value
func (o *Styles) SetMaxWidthSectionCentering(v int32) {
	o.MaxWidthSectionCentering = v
}

// GetBackgroundGradient returns the BackgroundGradient field value
func (o *Styles) GetBackgroundGradient() Gradient {
	if o == nil {
		var ret Gradient
		return ret
	}

	return o.BackgroundGradient
}

// GetBackgroundGradientOk returns a tuple with the BackgroundGradient field value
// and a boolean to check if the value has been set.
func (o *Styles) GetBackgroundGradientOk() (*Gradient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackgroundGradient, true
}

// SetBackgroundGradient sets field value
func (o *Styles) SetBackgroundGradient(v Gradient) {
	o.BackgroundGradient = v
}

func (o Styles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Styles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["backgroundColor"] = o.BackgroundColor
	toSerialize["flexboxPositioning"] = o.FlexboxPositioning
	toSerialize["backgroundImage"] = o.BackgroundImage
	toSerialize["forceFullWidthSection"] = o.ForceFullWidthSection
	toSerialize["verticalAlignment"] = o.VerticalAlignment
	toSerialize["maxWidthSectionCentering"] = o.MaxWidthSectionCentering
	toSerialize["backgroundGradient"] = o.BackgroundGradient
	return toSerialize, nil
}

func (o *Styles) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"backgroundColor",
		"flexboxPositioning",
		"backgroundImage",
		"forceFullWidthSection",
		"verticalAlignment",
		"maxWidthSectionCentering",
		"backgroundGradient",
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

	varStyles := _Styles{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStyles)

	if err != nil {
		return err
	}

	*o = Styles(varStyles)

	return err
}

type NullableStyles struct {
	value *Styles
	isSet bool
}

func (v NullableStyles) Get() *Styles {
	return v.value
}

func (v *NullableStyles) Set(val *Styles) {
	v.value = val
	v.isSet = true
}

func (v NullableStyles) IsSet() bool {
	return v.isSet
}

func (v *NullableStyles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStyles(val *Styles) *NullableStyles {
	return &NullableStyles{value: val, isSet: true}
}

func (v NullableStyles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStyles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


