# CalendarDatePropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseFiscalYear** | **bool** |  | 
**FiscalYearStart** | Pointer to **string** |  | [optional] 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PropertyType** | **string** |  | [default to "calendar-date"]
**TimeUnitCount** | **int32** |  | 
**OperationType** | **string** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 
**TimeUnit** | **string** |  | 

## Methods

### NewCalendarDatePropertyOperation

`func NewCalendarDatePropertyOperation(useFiscalYear bool, includeObjectsWithNoValueSet bool, propertyType string, timeUnitCount int32, operationType string, operatorName string, operator string, timeUnit string, ) *CalendarDatePropertyOperation`

NewCalendarDatePropertyOperation instantiates a new CalendarDatePropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarDatePropertyOperationWithDefaults

`func NewCalendarDatePropertyOperationWithDefaults() *CalendarDatePropertyOperation`

NewCalendarDatePropertyOperationWithDefaults instantiates a new CalendarDatePropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseFiscalYear

`func (o *CalendarDatePropertyOperation) GetUseFiscalYear() bool`

GetUseFiscalYear returns the UseFiscalYear field if non-nil, zero value otherwise.

### GetUseFiscalYearOk

`func (o *CalendarDatePropertyOperation) GetUseFiscalYearOk() (*bool, bool)`

GetUseFiscalYearOk returns a tuple with the UseFiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFiscalYear

`func (o *CalendarDatePropertyOperation) SetUseFiscalYear(v bool)`

SetUseFiscalYear sets UseFiscalYear field to given value.


### GetFiscalYearStart

`func (o *CalendarDatePropertyOperation) GetFiscalYearStart() string`

GetFiscalYearStart returns the FiscalYearStart field if non-nil, zero value otherwise.

### GetFiscalYearStartOk

`func (o *CalendarDatePropertyOperation) GetFiscalYearStartOk() (*string, bool)`

GetFiscalYearStartOk returns a tuple with the FiscalYearStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearStart

`func (o *CalendarDatePropertyOperation) SetFiscalYearStart(v string)`

SetFiscalYearStart sets FiscalYearStart field to given value.

### HasFiscalYearStart

`func (o *CalendarDatePropertyOperation) HasFiscalYearStart() bool`

HasFiscalYearStart returns a boolean if a field has been set.

### GetIncludeObjectsWithNoValueSet

`func (o *CalendarDatePropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *CalendarDatePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *CalendarDatePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *CalendarDatePropertyOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CalendarDatePropertyOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CalendarDatePropertyOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CalendarDatePropertyOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPropertyType

`func (o *CalendarDatePropertyOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *CalendarDatePropertyOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *CalendarDatePropertyOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetTimeUnitCount

`func (o *CalendarDatePropertyOperation) GetTimeUnitCount() int32`

GetTimeUnitCount returns the TimeUnitCount field if non-nil, zero value otherwise.

### GetTimeUnitCountOk

`func (o *CalendarDatePropertyOperation) GetTimeUnitCountOk() (*int32, bool)`

GetTimeUnitCountOk returns a tuple with the TimeUnitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnitCount

`func (o *CalendarDatePropertyOperation) SetTimeUnitCount(v int32)`

SetTimeUnitCount sets TimeUnitCount field to given value.


### GetOperationType

`func (o *CalendarDatePropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *CalendarDatePropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *CalendarDatePropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperatorName

`func (o *CalendarDatePropertyOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *CalendarDatePropertyOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *CalendarDatePropertyOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *CalendarDatePropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CalendarDatePropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CalendarDatePropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetTimeUnit

`func (o *CalendarDatePropertyOperation) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *CalendarDatePropertyOperation) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *CalendarDatePropertyOperation) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


