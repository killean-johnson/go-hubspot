# ApiManualEnrollmentCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldReEnroll** | **bool** | Whether or not the same object can enroll in this workflow twice. | 
**Type** | **string** | The type of enrollment criteria this is, this can be \&quot;LIST_BASED\&quot;, \&quot;EVENT_BASED\&quot;, or \&quot;MANUAL\&quot;. | [default to "MANUAL"]

## Methods

### NewApiManualEnrollmentCriteria

`func NewApiManualEnrollmentCriteria(shouldReEnroll bool, type_ string, ) *ApiManualEnrollmentCriteria`

NewApiManualEnrollmentCriteria instantiates a new ApiManualEnrollmentCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiManualEnrollmentCriteriaWithDefaults

`func NewApiManualEnrollmentCriteriaWithDefaults() *ApiManualEnrollmentCriteria`

NewApiManualEnrollmentCriteriaWithDefaults instantiates a new ApiManualEnrollmentCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldReEnroll

`func (o *ApiManualEnrollmentCriteria) GetShouldReEnroll() bool`

GetShouldReEnroll returns the ShouldReEnroll field if non-nil, zero value otherwise.

### GetShouldReEnrollOk

`func (o *ApiManualEnrollmentCriteria) GetShouldReEnrollOk() (*bool, bool)`

GetShouldReEnrollOk returns a tuple with the ShouldReEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldReEnroll

`func (o *ApiManualEnrollmentCriteria) SetShouldReEnroll(v bool)`

SetShouldReEnroll sets ShouldReEnroll field to given value.


### GetType

`func (o *ApiManualEnrollmentCriteria) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiManualEnrollmentCriteria) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiManualEnrollmentCriteria) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


