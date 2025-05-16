# EmailUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Name** | Pointer to **string** | The name of the email, as displayed on the email dashboard. | [optional] 
**Campaign** | Pointer to **string** | The ID of the campaign this email is associated to. | [optional] 
**From** | Pointer to [**PublicEmailFromDetails**](PublicEmailFromDetails.md) |  | [optional] 
**JitterSendTime** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** | The email state. | [optional] 
**To** | Pointer to [**PublicEmailToDetails**](PublicEmailToDetails.md) |  | [optional] 
**Subcategory** | Pointer to **string** | The email subcategory. | [optional] 
**SendOnPublish** | Pointer to **bool** | Determines whether the email will be sent immediately on publish. | [optional] 

## Methods

### NewEmailUpdateRequest

`func NewEmailUpdateRequest() *EmailUpdateRequest`

NewEmailUpdateRequest instantiates a new EmailUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailUpdateRequestWithDefaults

`func NewEmailUpdateRequestWithDefaults() *EmailUpdateRequest`

NewEmailUpdateRequestWithDefaults instantiates a new EmailUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRssData

`func (o *EmailUpdateRequest) GetRssData() PublicRssEmailDetails`

GetRssData returns the RssData field if non-nil, zero value otherwise.

### GetRssDataOk

`func (o *EmailUpdateRequest) GetRssDataOk() (*PublicRssEmailDetails, bool)`

GetRssDataOk returns a tuple with the RssData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssData

`func (o *EmailUpdateRequest) SetRssData(v PublicRssEmailDetails)`

SetRssData sets RssData field to given value.

### HasRssData

`func (o *EmailUpdateRequest) HasRssData() bool`

HasRssData returns a boolean if a field has been set.

### GetSubject

`func (o *EmailUpdateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailUpdateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailUpdateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailUpdateRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTesting

`func (o *EmailUpdateRequest) GetTesting() PublicEmailTestingDetails`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *EmailUpdateRequest) GetTestingOk() (*PublicEmailTestingDetails, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *EmailUpdateRequest) SetTesting(v PublicEmailTestingDetails)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *EmailUpdateRequest) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetPublishDate

`func (o *EmailUpdateRequest) GetPublishDate() time.Time`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *EmailUpdateRequest) GetPublishDateOk() (*time.Time, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *EmailUpdateRequest) SetPublishDate(v time.Time)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *EmailUpdateRequest) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### GetLanguage

`func (o *EmailUpdateRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EmailUpdateRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EmailUpdateRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *EmailUpdateRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *EmailUpdateRequest) GetBusinessUnitId() int64`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *EmailUpdateRequest) GetBusinessUnitIdOk() (*int64, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *EmailUpdateRequest) SetBusinessUnitId(v int64)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *EmailUpdateRequest) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetContent

`func (o *EmailUpdateRequest) GetContent() PublicEmailContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EmailUpdateRequest) GetContentOk() (*PublicEmailContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EmailUpdateRequest) SetContent(v PublicEmailContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *EmailUpdateRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetWebversion

`func (o *EmailUpdateRequest) GetWebversion() PublicWebversionDetails`

GetWebversion returns the Webversion field if non-nil, zero value otherwise.

### GetWebversionOk

`func (o *EmailUpdateRequest) GetWebversionOk() (*PublicWebversionDetails, bool)`

GetWebversionOk returns a tuple with the Webversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebversion

`func (o *EmailUpdateRequest) SetWebversion(v PublicWebversionDetails)`

SetWebversion sets Webversion field to given value.

### HasWebversion

`func (o *EmailUpdateRequest) HasWebversion() bool`

HasWebversion returns a boolean if a field has been set.

### GetArchived

`func (o *EmailUpdateRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *EmailUpdateRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *EmailUpdateRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *EmailUpdateRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetSubscriptionDetails

`func (o *EmailUpdateRequest) GetSubscriptionDetails() PublicEmailSubscriptionDetails`

GetSubscriptionDetails returns the SubscriptionDetails field if non-nil, zero value otherwise.

### GetSubscriptionDetailsOk

`func (o *EmailUpdateRequest) GetSubscriptionDetailsOk() (*PublicEmailSubscriptionDetails, bool)`

GetSubscriptionDetailsOk returns a tuple with the SubscriptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDetails

`func (o *EmailUpdateRequest) SetSubscriptionDetails(v PublicEmailSubscriptionDetails)`

SetSubscriptionDetails sets SubscriptionDetails field to given value.

### HasSubscriptionDetails

`func (o *EmailUpdateRequest) HasSubscriptionDetails() bool`

HasSubscriptionDetails returns a boolean if a field has been set.

### GetActiveDomain

`func (o *EmailUpdateRequest) GetActiveDomain() string`

GetActiveDomain returns the ActiveDomain field if non-nil, zero value otherwise.

### GetActiveDomainOk

`func (o *EmailUpdateRequest) GetActiveDomainOk() (*string, bool)`

GetActiveDomainOk returns a tuple with the ActiveDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDomain

`func (o *EmailUpdateRequest) SetActiveDomain(v string)`

SetActiveDomain sets ActiveDomain field to given value.

### HasActiveDomain

`func (o *EmailUpdateRequest) HasActiveDomain() bool`

HasActiveDomain returns a boolean if a field has been set.

### GetName

`func (o *EmailUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCampaign

`func (o *EmailUpdateRequest) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *EmailUpdateRequest) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *EmailUpdateRequest) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *EmailUpdateRequest) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetFrom

`func (o *EmailUpdateRequest) GetFrom() PublicEmailFromDetails`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailUpdateRequest) GetFromOk() (*PublicEmailFromDetails, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailUpdateRequest) SetFrom(v PublicEmailFromDetails)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailUpdateRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetJitterSendTime

`func (o *EmailUpdateRequest) GetJitterSendTime() bool`

GetJitterSendTime returns the JitterSendTime field if non-nil, zero value otherwise.

### GetJitterSendTimeOk

`func (o *EmailUpdateRequest) GetJitterSendTimeOk() (*bool, bool)`

GetJitterSendTimeOk returns a tuple with the JitterSendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitterSendTime

`func (o *EmailUpdateRequest) SetJitterSendTime(v bool)`

SetJitterSendTime sets JitterSendTime field to given value.

### HasJitterSendTime

`func (o *EmailUpdateRequest) HasJitterSendTime() bool`

HasJitterSendTime returns a boolean if a field has been set.

### GetState

`func (o *EmailUpdateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EmailUpdateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EmailUpdateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EmailUpdateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTo

`func (o *EmailUpdateRequest) GetTo() PublicEmailToDetails`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailUpdateRequest) GetToOk() (*PublicEmailToDetails, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailUpdateRequest) SetTo(v PublicEmailToDetails)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailUpdateRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetSubcategory

`func (o *EmailUpdateRequest) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *EmailUpdateRequest) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *EmailUpdateRequest) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.

### HasSubcategory

`func (o *EmailUpdateRequest) HasSubcategory() bool`

HasSubcategory returns a boolean if a field has been set.

### GetSendOnPublish

`func (o *EmailUpdateRequest) GetSendOnPublish() bool`

GetSendOnPublish returns the SendOnPublish field if non-nil, zero value otherwise.

### GetSendOnPublishOk

`func (o *EmailUpdateRequest) GetSendOnPublishOk() (*bool, bool)`

GetSendOnPublishOk returns a tuple with the SendOnPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnPublish

`func (o *EmailUpdateRequest) SetSendOnPublish(v bool)`

SetSendOnPublish sets SendOnPublish field to given value.

### HasSendOnPublish

`func (o *EmailUpdateRequest) HasSendOnPublish() bool`

HasSendOnPublish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


