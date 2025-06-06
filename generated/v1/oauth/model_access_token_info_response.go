/*
Auth Oauth

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oauth

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AccessTokenInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenInfoResponse{}

// AccessTokenInfoResponse struct for AccessTokenInfoResponse
type AccessTokenInfoResponse struct {
	// 
	HubId int32 `json:"hub_id"`
	// 
	UserId int32 `json:"user_id"`
	// 
	Scopes []string `json:"scopes"`
	// 
	TokenType string `json:"token_type"`
	// 
	User *string `json:"user,omitempty"`
	// 
	HubDomain *string `json:"hub_domain,omitempty"`
	// 
	AppId int32 `json:"app_id"`
	// 
	ExpiresIn int32 `json:"expires_in"`
	// 
	Token string `json:"token"`
}

type _AccessTokenInfoResponse AccessTokenInfoResponse

// NewAccessTokenInfoResponse instantiates a new AccessTokenInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenInfoResponse(hubId int32, userId int32, scopes []string, tokenType string, appId int32, expiresIn int32, token string) *AccessTokenInfoResponse {
	this := AccessTokenInfoResponse{}
	this.HubId = hubId
	this.UserId = userId
	this.Scopes = scopes
	this.TokenType = tokenType
	this.AppId = appId
	this.ExpiresIn = expiresIn
	this.Token = token
	return &this
}

// NewAccessTokenInfoResponseWithDefaults instantiates a new AccessTokenInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenInfoResponseWithDefaults() *AccessTokenInfoResponse {
	this := AccessTokenInfoResponse{}
	return &this
}

// GetHubId returns the HubId field value
func (o *AccessTokenInfoResponse) GetHubId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HubId
}

// GetHubIdOk returns a tuple with the HubId field value
// and a boolean to check if the value has been set.
func (o *AccessTokenInfoResponse) GetHubIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubId, true
}

// SetHubId sets field value
func (o *AccessTokenInfoResponse) SetHubId(v int32) {
	o.HubId = v
}

// GetUserId returns the UserId field value
func (o *AccessTokenInfoResponse) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AccessTokenInfoResponse) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AccessTokenInfoResponse) SetUserId(v int32) {
	o.UserId = v
}

// GetScopes returns the Scopes field value
func (o *AccessTokenInfoResponse) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *AccessTokenInfoResponse) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *AccessTokenInfoResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetTokenType returns the TokenType field value
func (o *AccessTokenInfoResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *AccessTokenInfoResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *AccessTokenInfoResponse) SetTokenType(v string) {
	o.TokenType = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AccessTokenInfoResponse) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenInfoResponse) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AccessTokenInfoResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *AccessTokenInfoResponse) SetUser(v string) {
	o.User = &v
}

// GetHubDomain returns the HubDomain field value if set, zero value otherwise.
func (o *AccessTokenInfoResponse) GetHubDomain() string {
	if o == nil || IsNil(o.HubDomain) {
		var ret string
		return ret
	}
	return *o.HubDomain
}

// GetHubDomainOk returns a tuple with the HubDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenInfoResponse) GetHubDomainOk() (*string, bool) {
	if o == nil || IsNil(o.HubDomain) {
		return nil, false
	}
	return o.HubDomain, true
}

// HasHubDomain returns a boolean if a field has been set.
func (o *AccessTokenInfoResponse) HasHubDomain() bool {
	if o != nil && !IsNil(o.HubDomain) {
		return true
	}

	return false
}

// SetHubDomain gets a reference to the given string and assigns it to the HubDomain field.
func (o *AccessTokenInfoResponse) SetHubDomain(v string) {
	o.HubDomain = &v
}

// GetAppId returns the AppId field value
func (o *AccessTokenInfoResponse) GetAppId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *AccessTokenInfoResponse) GetAppIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *AccessTokenInfoResponse) SetAppId(v int32) {
	o.AppId = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *AccessTokenInfoResponse) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *AccessTokenInfoResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *AccessTokenInfoResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

// GetToken returns the Token field value
func (o *AccessTokenInfoResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AccessTokenInfoResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AccessTokenInfoResponse) SetToken(v string) {
	o.Token = v
}

func (o AccessTokenInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hub_id"] = o.HubId
	toSerialize["user_id"] = o.UserId
	toSerialize["scopes"] = o.Scopes
	toSerialize["token_type"] = o.TokenType
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.HubDomain) {
		toSerialize["hub_domain"] = o.HubDomain
	}
	toSerialize["app_id"] = o.AppId
	toSerialize["expires_in"] = o.ExpiresIn
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *AccessTokenInfoResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hub_id",
		"user_id",
		"scopes",
		"token_type",
		"app_id",
		"expires_in",
		"token",
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

	varAccessTokenInfoResponse := _AccessTokenInfoResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessTokenInfoResponse)

	if err != nil {
		return err
	}

	*o = AccessTokenInfoResponse(varAccessTokenInfoResponse)

	return err
}

type NullableAccessTokenInfoResponse struct {
	value *AccessTokenInfoResponse
	isSet bool
}

func (v NullableAccessTokenInfoResponse) Get() *AccessTokenInfoResponse {
	return v.value
}

func (v *NullableAccessTokenInfoResponse) Set(val *AccessTokenInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenInfoResponse(val *AccessTokenInfoResponse) *NullableAccessTokenInfoResponse {
	return &NullableAccessTokenInfoResponse{value: val, isSet: true}
}

func (v NullableAccessTokenInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


