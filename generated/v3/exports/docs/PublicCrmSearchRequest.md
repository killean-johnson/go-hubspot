# PublicCrmSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**Filters** | [**[]Filter**](Filter.md) |  | 
**Sorts** | **[]string** |  | 

## Methods

### NewPublicCrmSearchRequest

`func NewPublicCrmSearchRequest(query string, filters []Filter, sorts []string, ) *PublicCrmSearchRequest`

NewPublicCrmSearchRequest instantiates a new PublicCrmSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCrmSearchRequestWithDefaults

`func NewPublicCrmSearchRequestWithDefaults() *PublicCrmSearchRequest`

NewPublicCrmSearchRequestWithDefaults instantiates a new PublicCrmSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *PublicCrmSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PublicCrmSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PublicCrmSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetFilters

`func (o *PublicCrmSearchRequest) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PublicCrmSearchRequest) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PublicCrmSearchRequest) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.


### GetSorts

`func (o *PublicCrmSearchRequest) GetSorts() []string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *PublicCrmSearchRequest) GetSortsOk() (*[]string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *PublicCrmSearchRequest) SetSorts(v []string)`

SetSorts sets Sorts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


