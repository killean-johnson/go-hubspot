# ApiContactPropertyAnchor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of event anchor this is, can be: \&quot;CONTACT_PROPERTY_ANCHOR\&quot; or \&quot;STATIC_DATE_ANCHOR\&quot; | [default to "CONTACT_PROPERTY_ANCHOR"]
**ContactProperty** | **string** | A date property on the contact to use as the anchor point of this workflow. | 

## Methods

### NewApiContactPropertyAnchor

`func NewApiContactPropertyAnchor(type_ string, contactProperty string, ) *ApiContactPropertyAnchor`

NewApiContactPropertyAnchor instantiates a new ApiContactPropertyAnchor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContactPropertyAnchorWithDefaults

`func NewApiContactPropertyAnchorWithDefaults() *ApiContactPropertyAnchor`

NewApiContactPropertyAnchorWithDefaults instantiates a new ApiContactPropertyAnchor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiContactPropertyAnchor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiContactPropertyAnchor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiContactPropertyAnchor) SetType(v string)`

SetType sets Type field to given value.


### GetContactProperty

`func (o *ApiContactPropertyAnchor) GetContactProperty() string`

GetContactProperty returns the ContactProperty field if non-nil, zero value otherwise.

### GetContactPropertyOk

`func (o *ApiContactPropertyAnchor) GetContactPropertyOk() (*string, bool)`

GetContactPropertyOk returns a tuple with the ContactProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactProperty

`func (o *ApiContactPropertyAnchor) SetContactProperty(v string)`

SetContactProperty sets ContactProperty field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


