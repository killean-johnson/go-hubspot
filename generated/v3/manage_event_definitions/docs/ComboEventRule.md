# ComboEventRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LookbackWindowDays** | Pointer to **int32** |  | [optional] 
**EventTypeId** | **string** |  | 
**PropertyFilters** | [**[]PropertyFilter**](PropertyFilter.md) |  | 
**Count** | **int32** |  | 

## Methods

### NewComboEventRule

`func NewComboEventRule(eventTypeId string, propertyFilters []PropertyFilter, count int32, ) *ComboEventRule`

NewComboEventRule instantiates a new ComboEventRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComboEventRuleWithDefaults

`func NewComboEventRuleWithDefaults() *ComboEventRule`

NewComboEventRuleWithDefaults instantiates a new ComboEventRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLookbackWindowDays

`func (o *ComboEventRule) GetLookbackWindowDays() int32`

GetLookbackWindowDays returns the LookbackWindowDays field if non-nil, zero value otherwise.

### GetLookbackWindowDaysOk

`func (o *ComboEventRule) GetLookbackWindowDaysOk() (*int32, bool)`

GetLookbackWindowDaysOk returns a tuple with the LookbackWindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookbackWindowDays

`func (o *ComboEventRule) SetLookbackWindowDays(v int32)`

SetLookbackWindowDays sets LookbackWindowDays field to given value.

### HasLookbackWindowDays

`func (o *ComboEventRule) HasLookbackWindowDays() bool`

HasLookbackWindowDays returns a boolean if a field has been set.

### GetEventTypeId

`func (o *ComboEventRule) GetEventTypeId() string`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *ComboEventRule) GetEventTypeIdOk() (*string, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *ComboEventRule) SetEventTypeId(v string)`

SetEventTypeId sets EventTypeId field to given value.


### GetPropertyFilters

`func (o *ComboEventRule) GetPropertyFilters() []PropertyFilter`

GetPropertyFilters returns the PropertyFilters field if non-nil, zero value otherwise.

### GetPropertyFiltersOk

`func (o *ComboEventRule) GetPropertyFiltersOk() (*[]PropertyFilter, bool)`

GetPropertyFiltersOk returns a tuple with the PropertyFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyFilters

`func (o *ComboEventRule) SetPropertyFilters(v []PropertyFilter)`

SetPropertyFilters sets PropertyFilters field to given value.


### GetCount

`func (o *ComboEventRule) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ComboEventRule) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ComboEventRule) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


