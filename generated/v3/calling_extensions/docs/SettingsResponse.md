# SettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | When this calling extension was created. | 
**SupportsCustomObjects** | **bool** | When true, users will be able to click to dial from custom objects. | 
**UsesRemote** | **bool** | When false, this indicates that your calling app does not use the anchored calling remote within the HubSpot app.  | 
**IsReady** | **bool** | When true, this indicates that your calling app is ready for production. Users will be able to select your calling app as their provider and can then click to dial within HubSpot. | 
**Name** | **string** | The name of your calling service to display to users. | 
**Width** | **int32** | The target width of the iframe that will contain your phone/calling UI. | 
**UsesCallingWindow** | **bool** | When false, this indicates that your calling app does not require the use of the separate calling window to hold the call connection.  | 
**SupportsInboundCalling** | **bool** | When true, this indicates that your calling app supports inbound calling within HubSpot.  | 
**Url** | **string** | The URL to your phone/calling UI, built with the [Calling SDK](#). | 
**Height** | **int32** | The target height of the iframe that will contain your phone/calling UI. | 
**UpdatedAt** | **time.Time** | The last time the settings for this calling extension were modified. | 

## Methods

### NewSettingsResponse

`func NewSettingsResponse(createdAt time.Time, supportsCustomObjects bool, usesRemote bool, isReady bool, name string, width int32, usesCallingWindow bool, supportsInboundCalling bool, url string, height int32, updatedAt time.Time, ) *SettingsResponse`

NewSettingsResponse instantiates a new SettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsResponseWithDefaults

`func NewSettingsResponseWithDefaults() *SettingsResponse`

NewSettingsResponseWithDefaults instantiates a new SettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SettingsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettingsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettingsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSupportsCustomObjects

`func (o *SettingsResponse) GetSupportsCustomObjects() bool`

GetSupportsCustomObjects returns the SupportsCustomObjects field if non-nil, zero value otherwise.

### GetSupportsCustomObjectsOk

`func (o *SettingsResponse) GetSupportsCustomObjectsOk() (*bool, bool)`

GetSupportsCustomObjectsOk returns a tuple with the SupportsCustomObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCustomObjects

`func (o *SettingsResponse) SetSupportsCustomObjects(v bool)`

SetSupportsCustomObjects sets SupportsCustomObjects field to given value.


### GetUsesRemote

`func (o *SettingsResponse) GetUsesRemote() bool`

GetUsesRemote returns the UsesRemote field if non-nil, zero value otherwise.

### GetUsesRemoteOk

`func (o *SettingsResponse) GetUsesRemoteOk() (*bool, bool)`

GetUsesRemoteOk returns a tuple with the UsesRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesRemote

`func (o *SettingsResponse) SetUsesRemote(v bool)`

SetUsesRemote sets UsesRemote field to given value.


### GetIsReady

`func (o *SettingsResponse) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *SettingsResponse) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *SettingsResponse) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.


### GetName

`func (o *SettingsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetWidth

`func (o *SettingsResponse) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *SettingsResponse) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *SettingsResponse) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetUsesCallingWindow

`func (o *SettingsResponse) GetUsesCallingWindow() bool`

GetUsesCallingWindow returns the UsesCallingWindow field if non-nil, zero value otherwise.

### GetUsesCallingWindowOk

`func (o *SettingsResponse) GetUsesCallingWindowOk() (*bool, bool)`

GetUsesCallingWindowOk returns a tuple with the UsesCallingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCallingWindow

`func (o *SettingsResponse) SetUsesCallingWindow(v bool)`

SetUsesCallingWindow sets UsesCallingWindow field to given value.


### GetSupportsInboundCalling

`func (o *SettingsResponse) GetSupportsInboundCalling() bool`

GetSupportsInboundCalling returns the SupportsInboundCalling field if non-nil, zero value otherwise.

### GetSupportsInboundCallingOk

`func (o *SettingsResponse) GetSupportsInboundCallingOk() (*bool, bool)`

GetSupportsInboundCallingOk returns a tuple with the SupportsInboundCalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsInboundCalling

`func (o *SettingsResponse) SetSupportsInboundCalling(v bool)`

SetSupportsInboundCalling sets SupportsInboundCalling field to given value.


### GetUrl

`func (o *SettingsResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SettingsResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SettingsResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeight

`func (o *SettingsResponse) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SettingsResponse) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SettingsResponse) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetUpdatedAt

`func (o *SettingsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SettingsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SettingsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


