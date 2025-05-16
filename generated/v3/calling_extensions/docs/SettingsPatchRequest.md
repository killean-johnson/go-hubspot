# SettingsPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsCustomObjects** | Pointer to **bool** | When true, users will be able to click to dial from custom objects. | [optional] 
**UsesRemote** | Pointer to **bool** | When false, this indicates that your calling app does not use the anchored calling remote within the HubSpot app.  | [optional] 
**IsReady** | Pointer to **bool** | When true, this indicates that your calling app is ready for production. Users will be able to select your calling app as their provider and can then click to dial within HubSpot. | [optional] 
**Name** | Pointer to **string** | The name of your calling service to display to users. | [optional] 
**Width** | Pointer to **int32** | The target width of the iframe that will contain your phone/calling UI. | [optional] 
**UsesCallingWindow** | Pointer to **bool** | When false, this indicates that your calling app does not require the use of the separate calling window to hold the call connection.  | [optional] 
**SupportsInboundCalling** | Pointer to **bool** | When true, this indicates that your calling app supports inbound calling within HubSpot. | [optional] 
**Url** | Pointer to **string** | The URL to your phone/calling UI, built with the [Calling SDK](#). | [optional] 
**Height** | Pointer to **int32** | The target height of the iframe that will contain your phone/calling UI. | [optional] 

## Methods

### NewSettingsPatchRequest

`func NewSettingsPatchRequest() *SettingsPatchRequest`

NewSettingsPatchRequest instantiates a new SettingsPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsPatchRequestWithDefaults

`func NewSettingsPatchRequestWithDefaults() *SettingsPatchRequest`

NewSettingsPatchRequestWithDefaults instantiates a new SettingsPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsCustomObjects

`func (o *SettingsPatchRequest) GetSupportsCustomObjects() bool`

GetSupportsCustomObjects returns the SupportsCustomObjects field if non-nil, zero value otherwise.

### GetSupportsCustomObjectsOk

`func (o *SettingsPatchRequest) GetSupportsCustomObjectsOk() (*bool, bool)`

GetSupportsCustomObjectsOk returns a tuple with the SupportsCustomObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCustomObjects

`func (o *SettingsPatchRequest) SetSupportsCustomObjects(v bool)`

SetSupportsCustomObjects sets SupportsCustomObjects field to given value.

### HasSupportsCustomObjects

`func (o *SettingsPatchRequest) HasSupportsCustomObjects() bool`

HasSupportsCustomObjects returns a boolean if a field has been set.

### GetUsesRemote

`func (o *SettingsPatchRequest) GetUsesRemote() bool`

GetUsesRemote returns the UsesRemote field if non-nil, zero value otherwise.

### GetUsesRemoteOk

`func (o *SettingsPatchRequest) GetUsesRemoteOk() (*bool, bool)`

GetUsesRemoteOk returns a tuple with the UsesRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesRemote

`func (o *SettingsPatchRequest) SetUsesRemote(v bool)`

SetUsesRemote sets UsesRemote field to given value.

### HasUsesRemote

`func (o *SettingsPatchRequest) HasUsesRemote() bool`

HasUsesRemote returns a boolean if a field has been set.

### GetIsReady

`func (o *SettingsPatchRequest) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *SettingsPatchRequest) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *SettingsPatchRequest) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.

### HasIsReady

`func (o *SettingsPatchRequest) HasIsReady() bool`

HasIsReady returns a boolean if a field has been set.

### GetName

`func (o *SettingsPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingsPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingsPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SettingsPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWidth

`func (o *SettingsPatchRequest) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *SettingsPatchRequest) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *SettingsPatchRequest) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *SettingsPatchRequest) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetUsesCallingWindow

`func (o *SettingsPatchRequest) GetUsesCallingWindow() bool`

GetUsesCallingWindow returns the UsesCallingWindow field if non-nil, zero value otherwise.

### GetUsesCallingWindowOk

`func (o *SettingsPatchRequest) GetUsesCallingWindowOk() (*bool, bool)`

GetUsesCallingWindowOk returns a tuple with the UsesCallingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCallingWindow

`func (o *SettingsPatchRequest) SetUsesCallingWindow(v bool)`

SetUsesCallingWindow sets UsesCallingWindow field to given value.

### HasUsesCallingWindow

`func (o *SettingsPatchRequest) HasUsesCallingWindow() bool`

HasUsesCallingWindow returns a boolean if a field has been set.

### GetSupportsInboundCalling

`func (o *SettingsPatchRequest) GetSupportsInboundCalling() bool`

GetSupportsInboundCalling returns the SupportsInboundCalling field if non-nil, zero value otherwise.

### GetSupportsInboundCallingOk

`func (o *SettingsPatchRequest) GetSupportsInboundCallingOk() (*bool, bool)`

GetSupportsInboundCallingOk returns a tuple with the SupportsInboundCalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsInboundCalling

`func (o *SettingsPatchRequest) SetSupportsInboundCalling(v bool)`

SetSupportsInboundCalling sets SupportsInboundCalling field to given value.

### HasSupportsInboundCalling

`func (o *SettingsPatchRequest) HasSupportsInboundCalling() bool`

HasSupportsInboundCalling returns a boolean if a field has been set.

### GetUrl

`func (o *SettingsPatchRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SettingsPatchRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SettingsPatchRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SettingsPatchRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHeight

`func (o *SettingsPatchRequest) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SettingsPatchRequest) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SettingsPatchRequest) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *SettingsPatchRequest) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


