# PublicEmailTestingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbSamplingDefault** | Pointer to **string** | Version of the email that should be sent if the results are inconclusive after the test period, master or variant. | [optional] 
**AbSampleSizeDefault** | Pointer to **string** | Version of the email that should be sent if there are too few recipients to conduct an AB test. | [optional] 
**AbStatus** | Pointer to **string** | Status of the AB test. | [optional] 
**AbTestPercentage** | Pointer to **int32** | The size of your test group. | [optional] 
**HoursToWait** | Pointer to **int32** | Time limit on gathering test results. After this time is up, the winning version will be sent to the remaining contacts. | [optional] 
**TestId** | Pointer to **string** | The ID of the AB test. | [optional] 
**AbSuccessMetric** | Pointer to **string** | Metric to determine the version that will be sent to the remaining contacts. | [optional] 

## Methods

### NewPublicEmailTestingDetails

`func NewPublicEmailTestingDetails() *PublicEmailTestingDetails`

NewPublicEmailTestingDetails instantiates a new PublicEmailTestingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEmailTestingDetailsWithDefaults

`func NewPublicEmailTestingDetailsWithDefaults() *PublicEmailTestingDetails`

NewPublicEmailTestingDetailsWithDefaults instantiates a new PublicEmailTestingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbSamplingDefault

`func (o *PublicEmailTestingDetails) GetAbSamplingDefault() string`

GetAbSamplingDefault returns the AbSamplingDefault field if non-nil, zero value otherwise.

### GetAbSamplingDefaultOk

`func (o *PublicEmailTestingDetails) GetAbSamplingDefaultOk() (*string, bool)`

GetAbSamplingDefaultOk returns a tuple with the AbSamplingDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbSamplingDefault

`func (o *PublicEmailTestingDetails) SetAbSamplingDefault(v string)`

SetAbSamplingDefault sets AbSamplingDefault field to given value.

### HasAbSamplingDefault

`func (o *PublicEmailTestingDetails) HasAbSamplingDefault() bool`

HasAbSamplingDefault returns a boolean if a field has been set.

### GetAbSampleSizeDefault

`func (o *PublicEmailTestingDetails) GetAbSampleSizeDefault() string`

GetAbSampleSizeDefault returns the AbSampleSizeDefault field if non-nil, zero value otherwise.

### GetAbSampleSizeDefaultOk

`func (o *PublicEmailTestingDetails) GetAbSampleSizeDefaultOk() (*string, bool)`

GetAbSampleSizeDefaultOk returns a tuple with the AbSampleSizeDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbSampleSizeDefault

`func (o *PublicEmailTestingDetails) SetAbSampleSizeDefault(v string)`

SetAbSampleSizeDefault sets AbSampleSizeDefault field to given value.

### HasAbSampleSizeDefault

`func (o *PublicEmailTestingDetails) HasAbSampleSizeDefault() bool`

HasAbSampleSizeDefault returns a boolean if a field has been set.

### GetAbStatus

`func (o *PublicEmailTestingDetails) GetAbStatus() string`

GetAbStatus returns the AbStatus field if non-nil, zero value otherwise.

### GetAbStatusOk

`func (o *PublicEmailTestingDetails) GetAbStatusOk() (*string, bool)`

GetAbStatusOk returns a tuple with the AbStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbStatus

`func (o *PublicEmailTestingDetails) SetAbStatus(v string)`

SetAbStatus sets AbStatus field to given value.

### HasAbStatus

`func (o *PublicEmailTestingDetails) HasAbStatus() bool`

HasAbStatus returns a boolean if a field has been set.

### GetAbTestPercentage

`func (o *PublicEmailTestingDetails) GetAbTestPercentage() int32`

GetAbTestPercentage returns the AbTestPercentage field if non-nil, zero value otherwise.

### GetAbTestPercentageOk

`func (o *PublicEmailTestingDetails) GetAbTestPercentageOk() (*int32, bool)`

GetAbTestPercentageOk returns a tuple with the AbTestPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbTestPercentage

`func (o *PublicEmailTestingDetails) SetAbTestPercentage(v int32)`

SetAbTestPercentage sets AbTestPercentage field to given value.

### HasAbTestPercentage

`func (o *PublicEmailTestingDetails) HasAbTestPercentage() bool`

HasAbTestPercentage returns a boolean if a field has been set.

### GetHoursToWait

`func (o *PublicEmailTestingDetails) GetHoursToWait() int32`

GetHoursToWait returns the HoursToWait field if non-nil, zero value otherwise.

### GetHoursToWaitOk

`func (o *PublicEmailTestingDetails) GetHoursToWaitOk() (*int32, bool)`

GetHoursToWaitOk returns a tuple with the HoursToWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursToWait

`func (o *PublicEmailTestingDetails) SetHoursToWait(v int32)`

SetHoursToWait sets HoursToWait field to given value.

### HasHoursToWait

`func (o *PublicEmailTestingDetails) HasHoursToWait() bool`

HasHoursToWait returns a boolean if a field has been set.

### GetTestId

`func (o *PublicEmailTestingDetails) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *PublicEmailTestingDetails) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *PublicEmailTestingDetails) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *PublicEmailTestingDetails) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetAbSuccessMetric

`func (o *PublicEmailTestingDetails) GetAbSuccessMetric() string`

GetAbSuccessMetric returns the AbSuccessMetric field if non-nil, zero value otherwise.

### GetAbSuccessMetricOk

`func (o *PublicEmailTestingDetails) GetAbSuccessMetricOk() (*string, bool)`

GetAbSuccessMetricOk returns a tuple with the AbSuccessMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbSuccessMetric

`func (o *PublicEmailTestingDetails) SetAbSuccessMetric(v string)`

SetAbSuccessMetric sets AbSuccessMetric field to given value.

### HasAbSuccessMetric

`func (o *PublicEmailTestingDetails) HasAbSuccessMetric() bool`

HasAbSuccessMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


