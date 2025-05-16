# StreamingCollectionResponseWithTotalHubDbTableRowV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Type** | **string** |  | [default to "STREAMING"]
**Results** | [**[]HubDbTableRowV3**](HubDbTableRowV3.md) |  | 

## Methods

### NewStreamingCollectionResponseWithTotalHubDbTableRowV3

`func NewStreamingCollectionResponseWithTotalHubDbTableRowV3(total int32, type_ string, results []HubDbTableRowV3, ) *StreamingCollectionResponseWithTotalHubDbTableRowV3`

NewStreamingCollectionResponseWithTotalHubDbTableRowV3 instantiates a new StreamingCollectionResponseWithTotalHubDbTableRowV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamingCollectionResponseWithTotalHubDbTableRowV3WithDefaults

`func NewStreamingCollectionResponseWithTotalHubDbTableRowV3WithDefaults() *StreamingCollectionResponseWithTotalHubDbTableRowV3`

NewStreamingCollectionResponseWithTotalHubDbTableRowV3WithDefaults instantiates a new StreamingCollectionResponseWithTotalHubDbTableRowV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetPaging

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetType

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) SetType(v string)`

SetType sets Type field to given value.


### GetResults

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) GetResults() []HubDbTableRowV3`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) GetResultsOk() (*[]HubDbTableRowV3, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *StreamingCollectionResponseWithTotalHubDbTableRowV3) SetResults(v []HubDbTableRowV3)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


