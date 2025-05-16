# SignedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extension** | **string** | Extension of the requested file. | 
**Size** | **int64** | Size in bytes of the requested file. | 
**Name** | **string** | Name of the requested file. | 
**Width** | Pointer to **int32** | For image and video files. The width of the file. | [optional] 
**Type** | **string** | Type of the file. Can be IMG, DOCUMENT, AUDIO, MOVIE, or OTHER. | 
**Url** | **string** | Signed URL with access to the specified file. Anyone with this URL will be able to access the file until it expires. | 
**ExpiresAt** | **time.Time** | Timestamp of when the URL will no longer grant access to the file. | 
**Height** | Pointer to **int32** | For image and video files. The height of the file. | [optional] 

## Methods

### NewSignedUrl

`func NewSignedUrl(extension string, size int64, name string, type_ string, url string, expiresAt time.Time, ) *SignedUrl`

NewSignedUrl instantiates a new SignedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignedUrlWithDefaults

`func NewSignedUrlWithDefaults() *SignedUrl`

NewSignedUrlWithDefaults instantiates a new SignedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtension

`func (o *SignedUrl) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *SignedUrl) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *SignedUrl) SetExtension(v string)`

SetExtension sets Extension field to given value.


### GetSize

`func (o *SignedUrl) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SignedUrl) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SignedUrl) SetSize(v int64)`

SetSize sets Size field to given value.


### GetName

`func (o *SignedUrl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignedUrl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignedUrl) SetName(v string)`

SetName sets Name field to given value.


### GetWidth

`func (o *SignedUrl) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *SignedUrl) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *SignedUrl) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *SignedUrl) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetType

`func (o *SignedUrl) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignedUrl) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignedUrl) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *SignedUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SignedUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SignedUrl) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetExpiresAt

`func (o *SignedUrl) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SignedUrl) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SignedUrl) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetHeight

`func (o *SignedUrl) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SignedUrl) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SignedUrl) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *SignedUrl) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


