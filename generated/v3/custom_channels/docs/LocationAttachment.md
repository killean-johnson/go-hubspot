# LocationAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Latitude** | **float32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | [default to "LOCATION"]
**Url** | Pointer to **string** |  | [optional] 
**Longitude** | **float32** |  | 

## Methods

### NewLocationAttachment

`func NewLocationAttachment(latitude float32, type_ string, longitude float32, ) *LocationAttachment`

NewLocationAttachment instantiates a new LocationAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAttachmentWithDefaults

`func NewLocationAttachmentWithDefaults() *LocationAttachment`

NewLocationAttachmentWithDefaults instantiates a new LocationAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *LocationAttachment) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LocationAttachment) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LocationAttachment) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LocationAttachment) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *LocationAttachment) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocationAttachment) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocationAttachment) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetName

`func (o *LocationAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *LocationAttachment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocationAttachment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocationAttachment) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *LocationAttachment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LocationAttachment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LocationAttachment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LocationAttachment) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLongitude

`func (o *LocationAttachment) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocationAttachment) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocationAttachment) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


