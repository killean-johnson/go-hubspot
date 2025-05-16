# TimePointOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointBehavior** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PropertyType** | **string** |  | [default to "timepoint"]
**PropertyParser** | **string** |  | 
**TimePoint** | [**RangedTimeOperationUpperBoundTimePoint**](RangedTimeOperationUpperBoundTimePoint.md) |  | 
**OperationType** | **string** |  | 
**Type** | **string** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 

## Methods

### NewTimePointOperation

`func NewTimePointOperation(endpointBehavior string, includeObjectsWithNoValueSet bool, propertyType string, propertyParser string, timePoint RangedTimeOperationUpperBoundTimePoint, operationType string, type_ string, operatorName string, operator string, ) *TimePointOperation`

NewTimePointOperation instantiates a new TimePointOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimePointOperationWithDefaults

`func NewTimePointOperationWithDefaults() *TimePointOperation`

NewTimePointOperationWithDefaults instantiates a new TimePointOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointBehavior

`func (o *TimePointOperation) GetEndpointBehavior() string`

GetEndpointBehavior returns the EndpointBehavior field if non-nil, zero value otherwise.

### GetEndpointBehaviorOk

`func (o *TimePointOperation) GetEndpointBehaviorOk() (*string, bool)`

GetEndpointBehaviorOk returns a tuple with the EndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointBehavior

`func (o *TimePointOperation) SetEndpointBehavior(v string)`

SetEndpointBehavior sets EndpointBehavior field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *TimePointOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *TimePointOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *TimePointOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *TimePointOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *TimePointOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *TimePointOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *TimePointOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPropertyType

`func (o *TimePointOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *TimePointOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *TimePointOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetPropertyParser

`func (o *TimePointOperation) GetPropertyParser() string`

GetPropertyParser returns the PropertyParser field if non-nil, zero value otherwise.

### GetPropertyParserOk

`func (o *TimePointOperation) GetPropertyParserOk() (*string, bool)`

GetPropertyParserOk returns a tuple with the PropertyParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParser

`func (o *TimePointOperation) SetPropertyParser(v string)`

SetPropertyParser sets PropertyParser field to given value.


### GetTimePoint

`func (o *TimePointOperation) GetTimePoint() RangedTimeOperationUpperBoundTimePoint`

GetTimePoint returns the TimePoint field if non-nil, zero value otherwise.

### GetTimePointOk

`func (o *TimePointOperation) GetTimePointOk() (*RangedTimeOperationUpperBoundTimePoint, bool)`

GetTimePointOk returns a tuple with the TimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePoint

`func (o *TimePointOperation) SetTimePoint(v RangedTimeOperationUpperBoundTimePoint)`

SetTimePoint sets TimePoint field to given value.


### GetOperationType

`func (o *TimePointOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TimePointOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TimePointOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetType

`func (o *TimePointOperation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimePointOperation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimePointOperation) SetType(v string)`

SetType sets Type field to given value.


### GetOperatorName

`func (o *TimePointOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *TimePointOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *TimePointOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *TimePointOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TimePointOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TimePointOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


