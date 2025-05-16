# PropertyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** |  | 
**SelectedByUser** | **bool** |  | 
**SourceLabel** | **string** |  | 
**Source** | **string** |  | 
**UpdatedByUserId** | **int32** |  | 
**PersistenceTimestamp** | **int64** |  | 
**SourceMetadata** | **string** |  | 
**DataSensitivity** | **string** |  | 
**SourceVid** | **[]int64** |  | 
**Unit** | **string** |  | 
**RequestId** | **string** |  | 
**IsEncrypted** | **bool** |  | 
**Name** | **string** |  | 
**UseTimestampAsPersistenceTimestamp** | **bool** |  | 
**Value** | **string** |  | 
**SelectedByUserTimestamp** | **int64** |  | 
**Timestamp** | **int64** |  | 
**IsLargeValue** | **bool** |  | 

## Methods

### NewPropertyValue

`func NewPropertyValue(sourceId string, selectedByUser bool, sourceLabel string, source string, updatedByUserId int32, persistenceTimestamp int64, sourceMetadata string, dataSensitivity string, sourceVid []int64, unit string, requestId string, isEncrypted bool, name string, useTimestampAsPersistenceTimestamp bool, value string, selectedByUserTimestamp int64, timestamp int64, isLargeValue bool, ) *PropertyValue`

NewPropertyValue instantiates a new PropertyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyValueWithDefaults

`func NewPropertyValueWithDefaults() *PropertyValue`

NewPropertyValueWithDefaults instantiates a new PropertyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *PropertyValue) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *PropertyValue) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *PropertyValue) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSelectedByUser

`func (o *PropertyValue) GetSelectedByUser() bool`

GetSelectedByUser returns the SelectedByUser field if non-nil, zero value otherwise.

### GetSelectedByUserOk

`func (o *PropertyValue) GetSelectedByUserOk() (*bool, bool)`

GetSelectedByUserOk returns a tuple with the SelectedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedByUser

`func (o *PropertyValue) SetSelectedByUser(v bool)`

SetSelectedByUser sets SelectedByUser field to given value.


### GetSourceLabel

`func (o *PropertyValue) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *PropertyValue) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *PropertyValue) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.


### GetSource

`func (o *PropertyValue) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PropertyValue) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PropertyValue) SetSource(v string)`

SetSource sets Source field to given value.


### GetUpdatedByUserId

`func (o *PropertyValue) GetUpdatedByUserId() int32`

GetUpdatedByUserId returns the UpdatedByUserId field if non-nil, zero value otherwise.

### GetUpdatedByUserIdOk

`func (o *PropertyValue) GetUpdatedByUserIdOk() (*int32, bool)`

GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserId

`func (o *PropertyValue) SetUpdatedByUserId(v int32)`

SetUpdatedByUserId sets UpdatedByUserId field to given value.


### GetPersistenceTimestamp

`func (o *PropertyValue) GetPersistenceTimestamp() int64`

GetPersistenceTimestamp returns the PersistenceTimestamp field if non-nil, zero value otherwise.

### GetPersistenceTimestampOk

`func (o *PropertyValue) GetPersistenceTimestampOk() (*int64, bool)`

GetPersistenceTimestampOk returns a tuple with the PersistenceTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceTimestamp

`func (o *PropertyValue) SetPersistenceTimestamp(v int64)`

SetPersistenceTimestamp sets PersistenceTimestamp field to given value.


### GetSourceMetadata

`func (o *PropertyValue) GetSourceMetadata() string`

GetSourceMetadata returns the SourceMetadata field if non-nil, zero value otherwise.

### GetSourceMetadataOk

`func (o *PropertyValue) GetSourceMetadataOk() (*string, bool)`

GetSourceMetadataOk returns a tuple with the SourceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMetadata

`func (o *PropertyValue) SetSourceMetadata(v string)`

SetSourceMetadata sets SourceMetadata field to given value.


### GetDataSensitivity

`func (o *PropertyValue) GetDataSensitivity() string`

GetDataSensitivity returns the DataSensitivity field if non-nil, zero value otherwise.

### GetDataSensitivityOk

`func (o *PropertyValue) GetDataSensitivityOk() (*string, bool)`

GetDataSensitivityOk returns a tuple with the DataSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSensitivity

`func (o *PropertyValue) SetDataSensitivity(v string)`

SetDataSensitivity sets DataSensitivity field to given value.


### GetSourceVid

`func (o *PropertyValue) GetSourceVid() []int64`

GetSourceVid returns the SourceVid field if non-nil, zero value otherwise.

### GetSourceVidOk

`func (o *PropertyValue) GetSourceVidOk() (*[]int64, bool)`

GetSourceVidOk returns a tuple with the SourceVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVid

`func (o *PropertyValue) SetSourceVid(v []int64)`

SetSourceVid sets SourceVid field to given value.


### GetUnit

`func (o *PropertyValue) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PropertyValue) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PropertyValue) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetRequestId

`func (o *PropertyValue) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PropertyValue) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PropertyValue) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetIsEncrypted

`func (o *PropertyValue) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *PropertyValue) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *PropertyValue) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.


### GetName

`func (o *PropertyValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyValue) SetName(v string)`

SetName sets Name field to given value.


### GetUseTimestampAsPersistenceTimestamp

`func (o *PropertyValue) GetUseTimestampAsPersistenceTimestamp() bool`

GetUseTimestampAsPersistenceTimestamp returns the UseTimestampAsPersistenceTimestamp field if non-nil, zero value otherwise.

### GetUseTimestampAsPersistenceTimestampOk

`func (o *PropertyValue) GetUseTimestampAsPersistenceTimestampOk() (*bool, bool)`

GetUseTimestampAsPersistenceTimestampOk returns a tuple with the UseTimestampAsPersistenceTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimestampAsPersistenceTimestamp

`func (o *PropertyValue) SetUseTimestampAsPersistenceTimestamp(v bool)`

SetUseTimestampAsPersistenceTimestamp sets UseTimestampAsPersistenceTimestamp field to given value.


### GetValue

`func (o *PropertyValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PropertyValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PropertyValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetSelectedByUserTimestamp

`func (o *PropertyValue) GetSelectedByUserTimestamp() int64`

GetSelectedByUserTimestamp returns the SelectedByUserTimestamp field if non-nil, zero value otherwise.

### GetSelectedByUserTimestampOk

`func (o *PropertyValue) GetSelectedByUserTimestampOk() (*int64, bool)`

GetSelectedByUserTimestampOk returns a tuple with the SelectedByUserTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedByUserTimestamp

`func (o *PropertyValue) SetSelectedByUserTimestamp(v int64)`

SetSelectedByUserTimestamp sets SelectedByUserTimestamp field to given value.


### GetTimestamp

`func (o *PropertyValue) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PropertyValue) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PropertyValue) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetIsLargeValue

`func (o *PropertyValue) GetIsLargeValue() bool`

GetIsLargeValue returns the IsLargeValue field if non-nil, zero value otherwise.

### GetIsLargeValueOk

`func (o *PropertyValue) GetIsLargeValueOk() (*bool, bool)`

GetIsLargeValueOk returns a tuple with the IsLargeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLargeValue

`func (o *PropertyValue) SetIsLargeValue(v bool)`

SetIsLargeValue sets IsLargeValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


