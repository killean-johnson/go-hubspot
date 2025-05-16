# RangedTimeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpperBoundEndpointBehavior** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**UpperBoundTimePoint** | [**RangedTimeOperationUpperBoundTimePoint**](RangedTimeOperationUpperBoundTimePoint.md) |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PropertyType** | **string** |  | [default to "rangedtime"]
**PropertyParser** | **string** |  | 
**OperationType** | **string** |  | 
**LowerBoundEndpointBehavior** | **string** |  | 
**Type** | **string** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 
**LowerBoundTimePoint** | [**RangedTimeOperationUpperBoundTimePoint**](RangedTimeOperationUpperBoundTimePoint.md) |  | 

## Methods

### NewRangedTimeOperation

`func NewRangedTimeOperation(upperBoundEndpointBehavior string, includeObjectsWithNoValueSet bool, upperBoundTimePoint RangedTimeOperationUpperBoundTimePoint, propertyType string, propertyParser string, operationType string, lowerBoundEndpointBehavior string, type_ string, operatorName string, operator string, lowerBoundTimePoint RangedTimeOperationUpperBoundTimePoint, ) *RangedTimeOperation`

NewRangedTimeOperation instantiates a new RangedTimeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangedTimeOperationWithDefaults

`func NewRangedTimeOperationWithDefaults() *RangedTimeOperation`

NewRangedTimeOperationWithDefaults instantiates a new RangedTimeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpperBoundEndpointBehavior

`func (o *RangedTimeOperation) GetUpperBoundEndpointBehavior() string`

GetUpperBoundEndpointBehavior returns the UpperBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetUpperBoundEndpointBehaviorOk

`func (o *RangedTimeOperation) GetUpperBoundEndpointBehaviorOk() (*string, bool)`

GetUpperBoundEndpointBehaviorOk returns a tuple with the UpperBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundEndpointBehavior

`func (o *RangedTimeOperation) SetUpperBoundEndpointBehavior(v string)`

SetUpperBoundEndpointBehavior sets UpperBoundEndpointBehavior field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *RangedTimeOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *RangedTimeOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *RangedTimeOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetUpperBoundTimePoint

`func (o *RangedTimeOperation) GetUpperBoundTimePoint() RangedTimeOperationUpperBoundTimePoint`

GetUpperBoundTimePoint returns the UpperBoundTimePoint field if non-nil, zero value otherwise.

### GetUpperBoundTimePointOk

`func (o *RangedTimeOperation) GetUpperBoundTimePointOk() (*RangedTimeOperationUpperBoundTimePoint, bool)`

GetUpperBoundTimePointOk returns a tuple with the UpperBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundTimePoint

`func (o *RangedTimeOperation) SetUpperBoundTimePoint(v RangedTimeOperationUpperBoundTimePoint)`

SetUpperBoundTimePoint sets UpperBoundTimePoint field to given value.


### GetDefaultValue

`func (o *RangedTimeOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *RangedTimeOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *RangedTimeOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *RangedTimeOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPropertyType

`func (o *RangedTimeOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *RangedTimeOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *RangedTimeOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetPropertyParser

`func (o *RangedTimeOperation) GetPropertyParser() string`

GetPropertyParser returns the PropertyParser field if non-nil, zero value otherwise.

### GetPropertyParserOk

`func (o *RangedTimeOperation) GetPropertyParserOk() (*string, bool)`

GetPropertyParserOk returns a tuple with the PropertyParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParser

`func (o *RangedTimeOperation) SetPropertyParser(v string)`

SetPropertyParser sets PropertyParser field to given value.


### GetOperationType

`func (o *RangedTimeOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *RangedTimeOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *RangedTimeOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetLowerBoundEndpointBehavior

`func (o *RangedTimeOperation) GetLowerBoundEndpointBehavior() string`

GetLowerBoundEndpointBehavior returns the LowerBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetLowerBoundEndpointBehaviorOk

`func (o *RangedTimeOperation) GetLowerBoundEndpointBehaviorOk() (*string, bool)`

GetLowerBoundEndpointBehaviorOk returns a tuple with the LowerBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundEndpointBehavior

`func (o *RangedTimeOperation) SetLowerBoundEndpointBehavior(v string)`

SetLowerBoundEndpointBehavior sets LowerBoundEndpointBehavior field to given value.


### GetType

`func (o *RangedTimeOperation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RangedTimeOperation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RangedTimeOperation) SetType(v string)`

SetType sets Type field to given value.


### GetOperatorName

`func (o *RangedTimeOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *RangedTimeOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *RangedTimeOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *RangedTimeOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RangedTimeOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RangedTimeOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetLowerBoundTimePoint

`func (o *RangedTimeOperation) GetLowerBoundTimePoint() RangedTimeOperationUpperBoundTimePoint`

GetLowerBoundTimePoint returns the LowerBoundTimePoint field if non-nil, zero value otherwise.

### GetLowerBoundTimePointOk

`func (o *RangedTimeOperation) GetLowerBoundTimePointOk() (*RangedTimeOperationUpperBoundTimePoint, bool)`

GetLowerBoundTimePointOk returns a tuple with the LowerBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundTimePoint

`func (o *RangedTimeOperation) SetLowerBoundTimePoint(v RangedTimeOperationUpperBoundTimePoint)`

SetLowerBoundTimePoint sets LowerBoundTimePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


