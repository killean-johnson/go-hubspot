# IndexedTimePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to [**IndexOffset**](IndexOffset.md) |  | [optional] 
**TimezoneSource** | **string** |  | 
**IndexReference** | [**IndexedTimePointIndexReference**](IndexedTimePointIndexReference.md) |  | 
**TimeType** | **string** |  | [default to "INDEXED"]
**ZoneId** | **string** |  | 

## Methods

### NewIndexedTimePoint

`func NewIndexedTimePoint(timezoneSource string, indexReference IndexedTimePointIndexReference, timeType string, zoneId string, ) *IndexedTimePoint`

NewIndexedTimePoint instantiates a new IndexedTimePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexedTimePointWithDefaults

`func NewIndexedTimePointWithDefaults() *IndexedTimePoint`

NewIndexedTimePointWithDefaults instantiates a new IndexedTimePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *IndexedTimePoint) GetOffset() IndexOffset`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IndexedTimePoint) GetOffsetOk() (*IndexOffset, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IndexedTimePoint) SetOffset(v IndexOffset)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *IndexedTimePoint) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTimezoneSource

`func (o *IndexedTimePoint) GetTimezoneSource() string`

GetTimezoneSource returns the TimezoneSource field if non-nil, zero value otherwise.

### GetTimezoneSourceOk

`func (o *IndexedTimePoint) GetTimezoneSourceOk() (*string, bool)`

GetTimezoneSourceOk returns a tuple with the TimezoneSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneSource

`func (o *IndexedTimePoint) SetTimezoneSource(v string)`

SetTimezoneSource sets TimezoneSource field to given value.


### GetIndexReference

`func (o *IndexedTimePoint) GetIndexReference() IndexedTimePointIndexReference`

GetIndexReference returns the IndexReference field if non-nil, zero value otherwise.

### GetIndexReferenceOk

`func (o *IndexedTimePoint) GetIndexReferenceOk() (*IndexedTimePointIndexReference, bool)`

GetIndexReferenceOk returns a tuple with the IndexReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexReference

`func (o *IndexedTimePoint) SetIndexReference(v IndexedTimePointIndexReference)`

SetIndexReference sets IndexReference field to given value.


### GetTimeType

`func (o *IndexedTimePoint) GetTimeType() string`

GetTimeType returns the TimeType field if non-nil, zero value otherwise.

### GetTimeTypeOk

`func (o *IndexedTimePoint) GetTimeTypeOk() (*string, bool)`

GetTimeTypeOk returns a tuple with the TimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeType

`func (o *IndexedTimePoint) SetTimeType(v string)`

SetTimeType sets TimeType field to given value.


### GetZoneId

`func (o *IndexedTimePoint) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *IndexedTimePoint) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *IndexedTimePoint) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


