# NextPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to **string** | The link to the next page. | [optional] 
**After** | **string** | The cursor token value to get the next set of results. Use this value as query parameter (&amp;after&#x3D;...) to obtain the next page. | 

## Methods

### NewNextPage

`func NewNextPage(after string, ) *NextPage`

NewNextPage instantiates a new NextPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextPageWithDefaults

`func NewNextPageWithDefaults() *NextPage`

NewNextPageWithDefaults instantiates a new NextPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *NextPage) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *NextPage) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *NextPage) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *NextPage) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetAfter

`func (o *NextPage) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *NextPage) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *NextPage) SetAfter(v string)`

SetAfter sets After field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


