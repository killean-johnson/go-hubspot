/*
Posts

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posts

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the BatchResponseBlogPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchResponseBlogPost{}

// BatchResponseBlogPost Response object for batch operations on blog posts.
type BatchResponseBlogPost struct {
	// Time of batch operation completion.
	CompletedAt time.Time `json:"completedAt"`
	// Time of batch operation request.
	RequestedAt *time.Time `json:"requestedAt,omitempty"`
	// Time of batch operation start.
	StartedAt time.Time `json:"startedAt"`
	// Links associated with batch operation.
	Links *map[string]string `json:"links,omitempty"`
	// Results of batch operation.
	Results []BlogPost `json:"results"`
	// Status of batch operation.
	Status string `json:"status"`
}

type _BatchResponseBlogPost BatchResponseBlogPost

// NewBatchResponseBlogPost instantiates a new BatchResponseBlogPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseBlogPost(completedAt time.Time, startedAt time.Time, results []BlogPost, status string) *BatchResponseBlogPost {
	this := BatchResponseBlogPost{}
	this.CompletedAt = completedAt
	this.StartedAt = startedAt
	this.Results = results
	this.Status = status
	return &this
}

// NewBatchResponseBlogPostWithDefaults instantiates a new BatchResponseBlogPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseBlogPostWithDefaults() *BatchResponseBlogPost {
	this := BatchResponseBlogPost{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponseBlogPost) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPost) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponseBlogPost) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponseBlogPost) GetRequestedAt() time.Time {
	if o == nil || IsNil(o.RequestedAt) {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPost) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestedAt) {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponseBlogPost) HasRequestedAt() bool {
	if o != nil && !IsNil(o.RequestedAt) {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponseBlogPost) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponseBlogPost) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPost) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponseBlogPost) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponseBlogPost) GetLinks() map[string]string {
	if o == nil || IsNil(o.Links) {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPost) GetLinksOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponseBlogPost) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponseBlogPost) SetLinks(v map[string]string) {
	o.Links = &v
}

// GetResults returns the Results field value
func (o *BatchResponseBlogPost) GetResults() []BlogPost {
	if o == nil {
		var ret []BlogPost
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPost) GetResultsOk() ([]BlogPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchResponseBlogPost) SetResults(v []BlogPost) {
	o.Results = v
}

// GetStatus returns the Status field value
func (o *BatchResponseBlogPost) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponseBlogPost) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponseBlogPost) SetStatus(v string) {
	o.Status = v
}

func (o BatchResponseBlogPost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchResponseBlogPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completedAt"] = o.CompletedAt
	if !IsNil(o.RequestedAt) {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	toSerialize["startedAt"] = o.StartedAt
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["results"] = o.Results
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *BatchResponseBlogPost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completedAt",
		"startedAt",
		"results",
		"status",
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

	varBatchResponseBlogPost := _BatchResponseBlogPost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchResponseBlogPost)

	if err != nil {
		return err
	}

	*o = BatchResponseBlogPost(varBatchResponseBlogPost)

	return err
}

type NullableBatchResponseBlogPost struct {
	value *BatchResponseBlogPost
	isSet bool
}

func (v NullableBatchResponseBlogPost) Get() *BatchResponseBlogPost {
	return v.value
}

func (v *NullableBatchResponseBlogPost) Set(val *BatchResponseBlogPost) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseBlogPost) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseBlogPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseBlogPost(val *BatchResponseBlogPost) *NullableBatchResponseBlogPost {
	return &NullableBatchResponseBlogPost{value: val, isSet: true}
}

func (v NullableBatchResponseBlogPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseBlogPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


