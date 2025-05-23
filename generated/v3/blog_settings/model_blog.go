/*
Blog Settings

\"Use these endpoints for interacting with Blog objects\"

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blog_settings

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Blog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Blog{}

// Blog struct for Blog
type Blog struct {
	Created time.Time `json:"created"`
	AbsoluteUrl string `json:"absoluteUrl"`
	// The Description of this Blog.
	Description string `json:"description"`
	// The explicitly defined language of the Blog. If null, the Blog will default to the language of the Domain.
	Language string `json:"language"`
	// ID of the primary Blog this object was translated from.
	TranslatedFromId string `json:"translatedFromId"`
	// Rules for require member registration to access private content.
	PublicAccessRules []map[string]interface{} `json:"publicAccessRules"`
	// The public title of this Blog.
	PublicTitle string `json:"publicTitle"`
	// Boolean determining whether or not this blog allows public comments.
	AllowComments bool `json:"allowComments"`
	// The timestamp (ISO8601 format) when this Blog was deleted.
	DeletedAt time.Time `json:"deletedAt"`
	// The html title of this Blog.
	HtmlTitle string `json:"htmlTitle"`
	// Boolean to determine whether or not to respect publicAccessRules.
	PublicAccessRulesEnabled bool `json:"publicAccessRulesEnabled"`
	// The internal name of the blog.
	Name string `json:"name"`
	// The unique ID of the Blog.
	Id string `json:"id"`
	Updated time.Time `json:"updated"`
	// The path of the this blog. This field is appended to the domain to construct the url of this blog.
	Slug string `json:"slug"`
}

type _Blog Blog

// NewBlog instantiates a new Blog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlog(created time.Time, absoluteUrl string, description string, language string, translatedFromId string, publicAccessRules []map[string]interface{}, publicTitle string, allowComments bool, deletedAt time.Time, htmlTitle string, publicAccessRulesEnabled bool, name string, id string, updated time.Time, slug string) *Blog {
	this := Blog{}
	this.Created = created
	this.AbsoluteUrl = absoluteUrl
	this.Description = description
	this.Language = language
	this.TranslatedFromId = translatedFromId
	this.PublicAccessRules = publicAccessRules
	this.PublicTitle = publicTitle
	this.AllowComments = allowComments
	this.DeletedAt = deletedAt
	this.HtmlTitle = htmlTitle
	this.PublicAccessRulesEnabled = publicAccessRulesEnabled
	this.Name = name
	this.Id = id
	this.Updated = updated
	this.Slug = slug
	return &this
}

// NewBlogWithDefaults instantiates a new Blog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlogWithDefaults() *Blog {
	this := Blog{}
	return &this
}

// GetCreated returns the Created field value
func (o *Blog) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Blog) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Blog) SetCreated(v time.Time) {
	o.Created = v
}

// GetAbsoluteUrl returns the AbsoluteUrl field value
func (o *Blog) GetAbsoluteUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AbsoluteUrl
}

// GetAbsoluteUrlOk returns a tuple with the AbsoluteUrl field value
// and a boolean to check if the value has been set.
func (o *Blog) GetAbsoluteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AbsoluteUrl, true
}

// SetAbsoluteUrl sets field value
func (o *Blog) SetAbsoluteUrl(v string) {
	o.AbsoluteUrl = v
}

// GetDescription returns the Description field value
func (o *Blog) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Blog) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Blog) SetDescription(v string) {
	o.Description = v
}

// GetLanguage returns the Language field value
func (o *Blog) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *Blog) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *Blog) SetLanguage(v string) {
	o.Language = v
}

// GetTranslatedFromId returns the TranslatedFromId field value
func (o *Blog) GetTranslatedFromId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TranslatedFromId
}

// GetTranslatedFromIdOk returns a tuple with the TranslatedFromId field value
// and a boolean to check if the value has been set.
func (o *Blog) GetTranslatedFromIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TranslatedFromId, true
}

// SetTranslatedFromId sets field value
func (o *Blog) SetTranslatedFromId(v string) {
	o.TranslatedFromId = v
}

// GetPublicAccessRules returns the PublicAccessRules field value
func (o *Blog) GetPublicAccessRules() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.PublicAccessRules
}

// GetPublicAccessRulesOk returns a tuple with the PublicAccessRules field value
// and a boolean to check if the value has been set.
func (o *Blog) GetPublicAccessRulesOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicAccessRules, true
}

// SetPublicAccessRules sets field value
func (o *Blog) SetPublicAccessRules(v []map[string]interface{}) {
	o.PublicAccessRules = v
}

// GetPublicTitle returns the PublicTitle field value
func (o *Blog) GetPublicTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicTitle
}

// GetPublicTitleOk returns a tuple with the PublicTitle field value
// and a boolean to check if the value has been set.
func (o *Blog) GetPublicTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicTitle, true
}

// SetPublicTitle sets field value
func (o *Blog) SetPublicTitle(v string) {
	o.PublicTitle = v
}

// GetAllowComments returns the AllowComments field value
func (o *Blog) GetAllowComments() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowComments
}

// GetAllowCommentsOk returns a tuple with the AllowComments field value
// and a boolean to check if the value has been set.
func (o *Blog) GetAllowCommentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowComments, true
}

// SetAllowComments sets field value
func (o *Blog) SetAllowComments(v bool) {
	o.AllowComments = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *Blog) GetDeletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *Blog) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *Blog) SetDeletedAt(v time.Time) {
	o.DeletedAt = v
}

// GetHtmlTitle returns the HtmlTitle field value
func (o *Blog) GetHtmlTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlTitle
}

// GetHtmlTitleOk returns a tuple with the HtmlTitle field value
// and a boolean to check if the value has been set.
func (o *Blog) GetHtmlTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlTitle, true
}

// SetHtmlTitle sets field value
func (o *Blog) SetHtmlTitle(v string) {
	o.HtmlTitle = v
}

// GetPublicAccessRulesEnabled returns the PublicAccessRulesEnabled field value
func (o *Blog) GetPublicAccessRulesEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PublicAccessRulesEnabled
}

// GetPublicAccessRulesEnabledOk returns a tuple with the PublicAccessRulesEnabled field value
// and a boolean to check if the value has been set.
func (o *Blog) GetPublicAccessRulesEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicAccessRulesEnabled, true
}

// SetPublicAccessRulesEnabled sets field value
func (o *Blog) SetPublicAccessRulesEnabled(v bool) {
	o.PublicAccessRulesEnabled = v
}

// GetName returns the Name field value
func (o *Blog) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Blog) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Blog) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *Blog) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Blog) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Blog) SetId(v string) {
	o.Id = v
}

// GetUpdated returns the Updated field value
func (o *Blog) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *Blog) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *Blog) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetSlug returns the Slug field value
func (o *Blog) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Blog) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Blog) SetSlug(v string) {
	o.Slug = v
}

func (o Blog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Blog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["absoluteUrl"] = o.AbsoluteUrl
	toSerialize["description"] = o.Description
	toSerialize["language"] = o.Language
	toSerialize["translatedFromId"] = o.TranslatedFromId
	toSerialize["publicAccessRules"] = o.PublicAccessRules
	toSerialize["publicTitle"] = o.PublicTitle
	toSerialize["allowComments"] = o.AllowComments
	toSerialize["deletedAt"] = o.DeletedAt
	toSerialize["htmlTitle"] = o.HtmlTitle
	toSerialize["publicAccessRulesEnabled"] = o.PublicAccessRulesEnabled
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	toSerialize["updated"] = o.Updated
	toSerialize["slug"] = o.Slug
	return toSerialize, nil
}

func (o *Blog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"absoluteUrl",
		"description",
		"language",
		"translatedFromId",
		"publicAccessRules",
		"publicTitle",
		"allowComments",
		"deletedAt",
		"htmlTitle",
		"publicAccessRulesEnabled",
		"name",
		"id",
		"updated",
		"slug",
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

	varBlog := _Blog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlog)

	if err != nil {
		return err
	}

	*o = Blog(varBlog)

	return err
}

type NullableBlog struct {
	value *Blog
	isSet bool
}

func (v NullableBlog) Get() *Blog {
	return v.value
}

func (v *NullableBlog) Set(val *Blog) {
	v.value = val
	v.isSet = true
}

func (v NullableBlog) IsSet() bool {
	return v.isSet
}

func (v *NullableBlog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlog(val *Blog) *NullableBlog {
	return &NullableBlog{value: val, isSet: true}
}

func (v NullableBlog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


