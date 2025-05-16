# ListSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListIds** | Pointer to **[]string** | The &#x60;listIds&#x60; that will be used to filter results by &#x60;listId&#x60;. If values are provided, then the response will only include results that have a &#x60;listId&#x60; in this array.  If no value is provided, or if an empty list is provided, then the results will not be filtered by &#x60;listId&#x60;. | [optional] 
**Offset** | Pointer to **int32** | Value used to paginate through lists. The &#x60;offset&#x60; provided in the response can be used in the next request to fetch the next page of results. Defaults to &#x60;0&#x60; if no offset is provided. | [optional] 
**Query** | Pointer to **string** | The &#x60;query&#x60; that will be used to search for lists by list name. If no &#x60;query&#x60; is provided, then the results will include all lists. | [optional] 
**Count** | Pointer to **int32** | The number of lists to include in the response. Defaults to &#x60;20&#x60; if no value is provided. The max &#x60;count&#x60; is &#x60;500&#x60;. | [optional] 
**ProcessingTypes** | Pointer to **[]string** | The &#x60;processingTypes&#x60; that will be used to filter results by &#x60;processingType&#x60;. If values are provided, then the response will only include results that have a &#x60;processingType&#x60; in this array.  If no value is provided, or if an empty list is provided, then results will not be filtered by &#x60;processingType&#x60;.  Valid &#x60;processingTypes&#x60; are: &#x60;MANUAL&#x60;, &#x60;SNAPSHOT&#x60;, or &#x60;DYNAMIC&#x60;. | [optional] 
**AdditionalPropertiesField** | Pointer to **[]string** | The property names of any additional list properties to include in the response. Properties that do not exist or that are empty for a particular list are not included in the response.  By default, all requests will fetch the following properties for each list: &#x60;hs_list_size&#x60;, &#x60;hs_last_record_added_at&#x60;, &#x60;hs_last_record_removed_at&#x60;, &#x60;hs_folder_name&#x60;, and &#x60;hs_list_reference_count&#x60;. | [optional] 
**Sort** | Pointer to **string** |  | [optional] 

## Methods

### NewListSearchRequest

`func NewListSearchRequest() *ListSearchRequest`

NewListSearchRequest instantiates a new ListSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSearchRequestWithDefaults

`func NewListSearchRequestWithDefaults() *ListSearchRequest`

NewListSearchRequestWithDefaults instantiates a new ListSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListIds

`func (o *ListSearchRequest) GetListIds() []string`

GetListIds returns the ListIds field if non-nil, zero value otherwise.

### GetListIdsOk

`func (o *ListSearchRequest) GetListIdsOk() (*[]string, bool)`

GetListIdsOk returns a tuple with the ListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListIds

`func (o *ListSearchRequest) SetListIds(v []string)`

SetListIds sets ListIds field to given value.

### HasListIds

`func (o *ListSearchRequest) HasListIds() bool`

HasListIds returns a boolean if a field has been set.

### GetOffset

`func (o *ListSearchRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListSearchRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListSearchRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListSearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQuery

`func (o *ListSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListSearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetCount

`func (o *ListSearchRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListSearchRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListSearchRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListSearchRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetProcessingTypes

`func (o *ListSearchRequest) GetProcessingTypes() []string`

GetProcessingTypes returns the ProcessingTypes field if non-nil, zero value otherwise.

### GetProcessingTypesOk

`func (o *ListSearchRequest) GetProcessingTypesOk() (*[]string, bool)`

GetProcessingTypesOk returns a tuple with the ProcessingTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTypes

`func (o *ListSearchRequest) SetProcessingTypes(v []string)`

SetProcessingTypes sets ProcessingTypes field to given value.

### HasProcessingTypes

`func (o *ListSearchRequest) HasProcessingTypes() bool`

HasProcessingTypes returns a boolean if a field has been set.

### GetAdditionalPropertiesField

`func (o *ListSearchRequest) GetAdditionalPropertiesField() []string`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *ListSearchRequest) GetAdditionalPropertiesFieldOk() (*[]string, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *ListSearchRequest) SetAdditionalPropertiesField(v []string)`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *ListSearchRequest) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### GetSort

`func (o *ListSearchRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ListSearchRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ListSearchRequest) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ListSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


