# ImportFromUrlTaskLocator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | **map[string]string** | Links for where to check information related to the task. The &#x60;status&#x60; link gives the URL for where to check the status of the task. | 
**Id** | **string** | ID of the task | 

## Methods

### NewImportFromUrlTaskLocator

`func NewImportFromUrlTaskLocator(links map[string]string, id string, ) *ImportFromUrlTaskLocator`

NewImportFromUrlTaskLocator instantiates a new ImportFromUrlTaskLocator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportFromUrlTaskLocatorWithDefaults

`func NewImportFromUrlTaskLocatorWithDefaults() *ImportFromUrlTaskLocator`

NewImportFromUrlTaskLocatorWithDefaults instantiates a new ImportFromUrlTaskLocator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ImportFromUrlTaskLocator) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ImportFromUrlTaskLocator) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ImportFromUrlTaskLocator) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.


### GetId

`func (o *ImportFromUrlTaskLocator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportFromUrlTaskLocator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportFromUrlTaskLocator) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


