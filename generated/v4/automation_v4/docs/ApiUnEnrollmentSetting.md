# ApiUnEnrollmentSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowIds** | **[]string** | The IDs of the flows to unenroll an object in if it&#39;s enrolled in this flow. | 
**Type** | **string** | The type of unenrollment to perform:  \&quot;ALL\&quot; - unenroll the object from all other flows  \&quot;SELECTIVE\&quot; - only unenroll the object from the flows specified in &#x60;flowIds&#x60; | 

## Methods

### NewApiUnEnrollmentSetting

`func NewApiUnEnrollmentSetting(flowIds []string, type_ string, ) *ApiUnEnrollmentSetting`

NewApiUnEnrollmentSetting instantiates a new ApiUnEnrollmentSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUnEnrollmentSettingWithDefaults

`func NewApiUnEnrollmentSettingWithDefaults() *ApiUnEnrollmentSetting`

NewApiUnEnrollmentSettingWithDefaults instantiates a new ApiUnEnrollmentSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowIds

`func (o *ApiUnEnrollmentSetting) GetFlowIds() []string`

GetFlowIds returns the FlowIds field if non-nil, zero value otherwise.

### GetFlowIdsOk

`func (o *ApiUnEnrollmentSetting) GetFlowIdsOk() (*[]string, bool)`

GetFlowIdsOk returns a tuple with the FlowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowIds

`func (o *ApiUnEnrollmentSetting) SetFlowIds(v []string)`

SetFlowIds sets FlowIds field to given value.


### GetType

`func (o *ApiUnEnrollmentSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiUnEnrollmentSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiUnEnrollmentSetting) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


