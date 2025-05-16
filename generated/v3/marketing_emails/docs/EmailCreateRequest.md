# EmailCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackSurveyId** | Pointer to **string** | The ID of the feedback survey linked to the email. | [optional] 
**RssData** | Pointer to [**PublicRssEmailDetails**](PublicRssEmailDetails.md) |  | [optional] 
**Subject** | Pointer to **string** | The subject of the email. | [optional] 
**Testing** | Pointer to [**PublicEmailTestingDetails**](PublicEmailTestingDetails.md) |  | [optional] 
**PublishDate** | Pointer to **time.Time** | The date and time the email is scheduled for, in ISO8601 representation. This is only used in local time or scheduled emails. | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**BusinessUnitId** | Pointer to **int64** |  | [optional] 
**Content** | Pointer to [**PublicEmailContent**](PublicEmailContent.md) |  | [optional] 
**Webversion** | Pointer to [**PublicWebversionDetails**](PublicWebversionDetails.md) |  | [optional] 
**Archived** | Pointer to **bool** | Determines if the email is archived or not. | [optional] 
**SubscriptionDetails** | Pointer to [**PublicEmailSubscriptionDetails**](PublicEmailSubscriptionDetails.md) |  | [optional] 
**ActiveDomain** | Pointer to **string** | The active domain of the email. | [optional] 
**Name** | **string** | The name of the email, as displayed on the email dashboard. | 
**Campaign** | Pointer to **string** | The ID of the campaign this email is associated to. | [optional] 
**From** | Pointer to [**PublicEmailFromDetails**](PublicEmailFromDetails.md) |  | [optional] 
**JitterSendTime** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** | The email state. | [optional] 
**To** | Pointer to [**PublicEmailToDetails**](PublicEmailToDetails.md) |  | [optional] 
**Subcategory** | Pointer to **string** | The email subcategory. | [optional] 
**SendOnPublish** | Pointer to **bool** | Determines whether the email will be sent immediately on publish. | [optional] 

## Methods

### NewEmailCreateRequest

`func NewEmailCreateRequest(name string, ) *EmailCreateRequest`

NewEmailCreateRequest instantiates a new EmailCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailCreateRequestWithDefaults

`func NewEmailCreateRequestWithDefaults() *EmailCreateRequest`

NewEmailCreateRequestWithDefaults instantiates a new EmailCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackSurveyId

`func (o *EmailCreateRequest) GetFeedbackSurveyId() string`

GetFeedbackSurveyId returns the FeedbackSurveyId field if non-nil, zero value otherwise.

### GetFeedbackSurveyIdOk

`func (o *EmailCreateRequest) GetFeedbackSurveyIdOk() (*string, bool)`

GetFeedbackSurveyIdOk returns a tuple with the FeedbackSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackSurveyId

`func (o *EmailCreateRequest) SetFeedbackSurveyId(v string)`

SetFeedbackSurveyId sets FeedbackSurveyId field to given value.

### HasFeedbackSurveyId

`func (o *EmailCreateRequest) HasFeedbackSurveyId() bool`

HasFeedbackSurveyId returns a boolean if a field has been set.

### GetRssData

`func (o *EmailCreateRequest) GetRssData() PublicRssEmailDetails`

GetRssData returns the RssData field if non-nil, zero value otherwise.

### GetRssDataOk

`func (o *EmailCreateRequest) GetRssDataOk() (*PublicRssEmailDetails, bool)`

GetRssDataOk returns a tuple with the RssData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssData

`func (o *EmailCreateRequest) SetRssData(v PublicRssEmailDetails)`

SetRssData sets RssData field to given value.

### HasRssData

`func (o *EmailCreateRequest) HasRssData() bool`

HasRssData returns a boolean if a field has been set.

### GetSubject

`func (o *EmailCreateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailCreateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailCreateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailCreateRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTesting

`func (o *EmailCreateRequest) GetTesting() PublicEmailTestingDetails`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *EmailCreateRequest) GetTestingOk() (*PublicEmailTestingDetails, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *EmailCreateRequest) SetTesting(v PublicEmailTestingDetails)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *EmailCreateRequest) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetPublishDate

`func (o *EmailCreateRequest) GetPublishDate() time.Time`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *EmailCreateRequest) GetPublishDateOk() (*time.Time, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *EmailCreateRequest) SetPublishDate(v time.Time)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *EmailCreateRequest) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### GetLanguage

`func (o *EmailCreateRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EmailCreateRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EmailCreateRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *EmailCreateRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *EmailCreateRequest) GetBusinessUnitId() int64`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *EmailCreateRequest) GetBusinessUnitIdOk() (*int64, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *EmailCreateRequest) SetBusinessUnitId(v int64)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *EmailCreateRequest) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetContent

`func (o *EmailCreateRequest) GetContent() PublicEmailContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EmailCreateRequest) GetContentOk() (*PublicEmailContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EmailCreateRequest) SetContent(v PublicEmailContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *EmailCreateRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetWebversion

`func (o *EmailCreateRequest) GetWebversion() PublicWebversionDetails`

GetWebversion returns the Webversion field if non-nil, zero value otherwise.

### GetWebversionOk

`func (o *EmailCreateRequest) GetWebversionOk() (*PublicWebversionDetails, bool)`

GetWebversionOk returns a tuple with the Webversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebversion

`func (o *EmailCreateRequest) SetWebversion(v PublicWebversionDetails)`

SetWebversion sets Webversion field to given value.

### HasWebversion

`func (o *EmailCreateRequest) HasWebversion() bool`

HasWebversion returns a boolean if a field has been set.

### GetArchived

`func (o *EmailCreateRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *EmailCreateRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *EmailCreateRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *EmailCreateRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetSubscriptionDetails

`func (o *EmailCreateRequest) GetSubscriptionDetails() PublicEmailSubscriptionDetails`

GetSubscriptionDetails returns the SubscriptionDetails field if non-nil, zero value otherwise.

### GetSubscriptionDetailsOk

`func (o *EmailCreateRequest) GetSubscriptionDetailsOk() (*PublicEmailSubscriptionDetails, bool)`

GetSubscriptionDetailsOk returns a tuple with the SubscriptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDetails

`func (o *EmailCreateRequest) SetSubscriptionDetails(v PublicEmailSubscriptionDetails)`

SetSubscriptionDetails sets SubscriptionDetails field to given value.

### HasSubscriptionDetails

`func (o *EmailCreateRequest) HasSubscriptionDetails() bool`

HasSubscriptionDetails returns a boolean if a field has been set.

### GetActiveDomain

`func (o *EmailCreateRequest) GetActiveDomain() string`

GetActiveDomain returns the ActiveDomain field if non-nil, zero value otherwise.

### GetActiveDomainOk

`func (o *EmailCreateRequest) GetActiveDomainOk() (*string, bool)`

GetActiveDomainOk returns a tuple with the ActiveDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDomain

`func (o *EmailCreateRequest) SetActiveDomain(v string)`

SetActiveDomain sets ActiveDomain field to given value.

### HasActiveDomain

`func (o *EmailCreateRequest) HasActiveDomain() bool`

HasActiveDomain returns a boolean if a field has been set.

### GetName

`func (o *EmailCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCampaign

`func (o *EmailCreateRequest) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *EmailCreateRequest) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *EmailCreateRequest) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *EmailCreateRequest) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetFrom

`func (o *EmailCreateRequest) GetFrom() PublicEmailFromDetails`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailCreateRequest) GetFromOk() (*PublicEmailFromDetails, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailCreateRequest) SetFrom(v PublicEmailFromDetails)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailCreateRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetJitterSendTime

`func (o *EmailCreateRequest) GetJitterSendTime() bool`

GetJitterSendTime returns the JitterSendTime field if non-nil, zero value otherwise.

### GetJitterSendTimeOk

`func (o *EmailCreateRequest) GetJitterSendTimeOk() (*bool, bool)`

GetJitterSendTimeOk returns a tuple with the JitterSendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitterSendTime

`func (o *EmailCreateRequest) SetJitterSendTime(v bool)`

SetJitterSendTime sets JitterSendTime field to given value.

### HasJitterSendTime

`func (o *EmailCreateRequest) HasJitterSendTime() bool`

HasJitterSendTime returns a boolean if a field has been set.

### GetState

`func (o *EmailCreateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EmailCreateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EmailCreateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EmailCreateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTo

`func (o *EmailCreateRequest) GetTo() PublicEmailToDetails`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailCreateRequest) GetToOk() (*PublicEmailToDetails, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailCreateRequest) SetTo(v PublicEmailToDetails)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailCreateRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetSubcategory

`func (o *EmailCreateRequest) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *EmailCreateRequest) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *EmailCreateRequest) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.

### HasSubcategory

`func (o *EmailCreateRequest) HasSubcategory() bool`

HasSubcategory returns a boolean if a field has been set.

### GetSendOnPublish

`func (o *EmailCreateRequest) GetSendOnPublish() bool`

GetSendOnPublish returns the SendOnPublish field if non-nil, zero value otherwise.

### GetSendOnPublishOk

`func (o *EmailCreateRequest) GetSendOnPublishOk() (*bool, bool)`

GetSendOnPublishOk returns a tuple with the SendOnPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnPublish

`func (o *EmailCreateRequest) SetSendOnPublish(v bool)`

SetSendOnPublish sets SendOnPublish field to given value.

### HasSendOnPublish

`func (o *EmailCreateRequest) HasSendOnPublish() bool`

HasSendOnPublish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


