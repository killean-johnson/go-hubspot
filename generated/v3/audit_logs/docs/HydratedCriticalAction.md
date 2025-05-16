# HydratedCriticalAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The time the activity took place. | 
**ActingUser** | Pointer to **string** | Email address of the user associated with the activity. | [optional] 
**RegionCode** | Pointer to **string** | The approximate region code. | [optional] 
**InfoUrl** | Pointer to **string** | A link to the URL where the action was taken in the account. | [optional] 
**CountryCode** | Pointer to **string** | The approximate country code. | [optional] 
**IpAddress** | Pointer to **string** | IP address where the activity originated. | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Id** | **string** | The unique ID of the activity. | 
**Type** | **string** | The type of activity. | 
**UserId** | **int32** | The user&#39;s unique ID. | 
**ObjectId** | Pointer to **string** | The ID of the affected object. | [optional] 

## Methods

### NewHydratedCriticalAction

`func NewHydratedCriticalAction(createdAt time.Time, id string, type_ string, userId int32, ) *HydratedCriticalAction`

NewHydratedCriticalAction instantiates a new HydratedCriticalAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHydratedCriticalActionWithDefaults

`func NewHydratedCriticalActionWithDefaults() *HydratedCriticalAction`

NewHydratedCriticalActionWithDefaults instantiates a new HydratedCriticalAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HydratedCriticalAction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HydratedCriticalAction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HydratedCriticalAction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActingUser

`func (o *HydratedCriticalAction) GetActingUser() string`

GetActingUser returns the ActingUser field if non-nil, zero value otherwise.

### GetActingUserOk

`func (o *HydratedCriticalAction) GetActingUserOk() (*string, bool)`

GetActingUserOk returns a tuple with the ActingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUser

`func (o *HydratedCriticalAction) SetActingUser(v string)`

SetActingUser sets ActingUser field to given value.

### HasActingUser

`func (o *HydratedCriticalAction) HasActingUser() bool`

HasActingUser returns a boolean if a field has been set.

### GetRegionCode

`func (o *HydratedCriticalAction) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *HydratedCriticalAction) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *HydratedCriticalAction) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *HydratedCriticalAction) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetInfoUrl

`func (o *HydratedCriticalAction) GetInfoUrl() string`

GetInfoUrl returns the InfoUrl field if non-nil, zero value otherwise.

### GetInfoUrlOk

`func (o *HydratedCriticalAction) GetInfoUrlOk() (*string, bool)`

GetInfoUrlOk returns a tuple with the InfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrl

`func (o *HydratedCriticalAction) SetInfoUrl(v string)`

SetInfoUrl sets InfoUrl field to given value.

### HasInfoUrl

`func (o *HydratedCriticalAction) HasInfoUrl() bool`

HasInfoUrl returns a boolean if a field has been set.

### GetCountryCode

`func (o *HydratedCriticalAction) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *HydratedCriticalAction) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *HydratedCriticalAction) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *HydratedCriticalAction) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetIpAddress

`func (o *HydratedCriticalAction) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *HydratedCriticalAction) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *HydratedCriticalAction) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *HydratedCriticalAction) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLocation

`func (o *HydratedCriticalAction) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *HydratedCriticalAction) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *HydratedCriticalAction) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *HydratedCriticalAction) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetId

`func (o *HydratedCriticalAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HydratedCriticalAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HydratedCriticalAction) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *HydratedCriticalAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HydratedCriticalAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HydratedCriticalAction) SetType(v string)`

SetType sets Type field to given value.


### GetUserId

`func (o *HydratedCriticalAction) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HydratedCriticalAction) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HydratedCriticalAction) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetObjectId

`func (o *HydratedCriticalAction) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *HydratedCriticalAction) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *HydratedCriticalAction) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *HydratedCriticalAction) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


