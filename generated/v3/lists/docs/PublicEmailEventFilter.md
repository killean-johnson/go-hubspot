# PublicEmailEventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClickUrl** | Pointer to **string** |  | [optional] 
**Level** | **string** |  | 
**PruningRefineBy** | Pointer to [**PublicFormSubmissionFilterCoalescingRefineBy**](PublicFormSubmissionFilterCoalescingRefineBy.md) |  | [optional] 
**AppId** | **string** |  | 
**EmailId** | **string** |  | 
**FilterType** | **string** |  | [default to "EMAIL_EVENT"]
**Operator** | **string** |  | 

## Methods

### NewPublicEmailEventFilter

`func NewPublicEmailEventFilter(level string, appId string, emailId string, filterType string, operator string, ) *PublicEmailEventFilter`

NewPublicEmailEventFilter instantiates a new PublicEmailEventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEmailEventFilterWithDefaults

`func NewPublicEmailEventFilterWithDefaults() *PublicEmailEventFilter`

NewPublicEmailEventFilterWithDefaults instantiates a new PublicEmailEventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClickUrl

`func (o *PublicEmailEventFilter) GetClickUrl() string`

GetClickUrl returns the ClickUrl field if non-nil, zero value otherwise.

### GetClickUrlOk

`func (o *PublicEmailEventFilter) GetClickUrlOk() (*string, bool)`

GetClickUrlOk returns a tuple with the ClickUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickUrl

`func (o *PublicEmailEventFilter) SetClickUrl(v string)`

SetClickUrl sets ClickUrl field to given value.

### HasClickUrl

`func (o *PublicEmailEventFilter) HasClickUrl() bool`

HasClickUrl returns a boolean if a field has been set.

### GetLevel

`func (o *PublicEmailEventFilter) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PublicEmailEventFilter) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PublicEmailEventFilter) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetPruningRefineBy

`func (o *PublicEmailEventFilter) GetPruningRefineBy() PublicFormSubmissionFilterCoalescingRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PublicEmailEventFilter) GetPruningRefineByOk() (*PublicFormSubmissionFilterCoalescingRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PublicEmailEventFilter) SetPruningRefineBy(v PublicFormSubmissionFilterCoalescingRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *PublicEmailEventFilter) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.

### GetAppId

`func (o *PublicEmailEventFilter) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *PublicEmailEventFilter) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *PublicEmailEventFilter) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetEmailId

`func (o *PublicEmailEventFilter) GetEmailId() string`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *PublicEmailEventFilter) GetEmailIdOk() (*string, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *PublicEmailEventFilter) SetEmailId(v string)`

SetEmailId sets EmailId field to given value.


### GetFilterType

`func (o *PublicEmailEventFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicEmailEventFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicEmailEventFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetOperator

`func (o *PublicEmailEventFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicEmailEventFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicEmailEventFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


