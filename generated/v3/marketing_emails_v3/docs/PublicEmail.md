# PublicEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackSurveyId** | Pointer to **string** | The ID of the feedback survey linked to the email. | [optional] 
**Subject** | Pointer to **string** | The subject of the email. | [optional] 
**PublishDate** | Pointer to **time.Time** | The date and time the email is scheduled for, in ISO8601 representation. This is only used in local time or scheduled emails. | [optional] 
**IsTransactional** | Pointer to **bool** | Returns whether the email is a transactional email or not. This is read only. | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The email type, this is derived from other properties on the email such as subcategory. | [optional] 
**Content** | Pointer to [**PublicEmailContent**](PublicEmailContent.md) |  | [optional] 
**BusinessUnitId** | Pointer to **string** |  | [optional] 
**Webversion** | Pointer to [**PublicWebversionDetails**](PublicWebversionDetails.md) |  | [optional] 
**Archived** | Pointer to **bool** | Determines if the email is archived or not. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time of the email&#39;s creation, in ISO8601 representation. | [optional] 
**Stats** | Pointer to [**EmailStatisticsData**](EmailStatisticsData.md) |  | [optional] 
**From** | Pointer to [**PublicEmailFromDetails**](PublicEmailFromDetails.md) |  | [optional] 
**Id** | Pointer to **string** | The email ID. | [optional] 
**State** | Pointer to **string** | The email state. | [optional] 
**CreatedById** | Pointer to **string** | The id of the user who created the email. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time of the last update to the email, in ISO8601 representation. | [optional] 
**RssData** | Pointer to [**PublicRssEmailDetails**](PublicRssEmailDetails.md) |  | [optional] 
**PublishedAt** | Pointer to **time.Time** | The date and time the email was published at, in ISO8601 representation. | [optional] 
**PublishedById** | Pointer to **string** |  | [optional] 
**IsPublished** | Pointer to **bool** | Returns the published status of the email. This is read only. | [optional] 
**Testing** | Pointer to [**PublicEmailTestingDetails**](PublicEmailTestingDetails.md) |  | [optional] 
**UpdatedById** | Pointer to **string** | The id of the user who last updated the email. | [optional] 
**FolderId** | Pointer to **int64** |  | [optional] 
**SubscriptionDetails** | Pointer to [**PublicEmailSubscriptionDetails**](PublicEmailSubscriptionDetails.md) |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** | The name of the email, as displayed on the email dashboard. | [optional] 
**ActiveDomain** | Pointer to **string** | The active domain of the email. | [optional] 
**Campaign** | Pointer to **string** | The ID of the campaign this email is associated to. | [optional] 
**To** | Pointer to [**PublicEmailToDetails**](PublicEmailToDetails.md) |  | [optional] 
**Subcategory** | Pointer to **string** | The email subcategory. | [optional] 
**SendOnPublish** | Pointer to **bool** | Determines whether the email will be sent immediately on publish. | [optional] 

## Methods

### NewPublicEmail

`func NewPublicEmail() *PublicEmail`

NewPublicEmail instantiates a new PublicEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEmailWithDefaults

`func NewPublicEmailWithDefaults() *PublicEmail`

NewPublicEmailWithDefaults instantiates a new PublicEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackSurveyId

`func (o *PublicEmail) GetFeedbackSurveyId() string`

GetFeedbackSurveyId returns the FeedbackSurveyId field if non-nil, zero value otherwise.

### GetFeedbackSurveyIdOk

`func (o *PublicEmail) GetFeedbackSurveyIdOk() (*string, bool)`

GetFeedbackSurveyIdOk returns a tuple with the FeedbackSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackSurveyId

`func (o *PublicEmail) SetFeedbackSurveyId(v string)`

SetFeedbackSurveyId sets FeedbackSurveyId field to given value.

### HasFeedbackSurveyId

`func (o *PublicEmail) HasFeedbackSurveyId() bool`

HasFeedbackSurveyId returns a boolean if a field has been set.

### GetSubject

`func (o *PublicEmail) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PublicEmail) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PublicEmail) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PublicEmail) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetPublishDate

`func (o *PublicEmail) GetPublishDate() time.Time`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *PublicEmail) GetPublishDateOk() (*time.Time, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *PublicEmail) SetPublishDate(v time.Time)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *PublicEmail) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### GetIsTransactional

`func (o *PublicEmail) GetIsTransactional() bool`

GetIsTransactional returns the IsTransactional field if non-nil, zero value otherwise.

### GetIsTransactionalOk

`func (o *PublicEmail) GetIsTransactionalOk() (*bool, bool)`

GetIsTransactionalOk returns a tuple with the IsTransactional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTransactional

`func (o *PublicEmail) SetIsTransactional(v bool)`

SetIsTransactional sets IsTransactional field to given value.

### HasIsTransactional

`func (o *PublicEmail) HasIsTransactional() bool`

HasIsTransactional returns a boolean if a field has been set.

### GetLanguage

`func (o *PublicEmail) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PublicEmail) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PublicEmail) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PublicEmail) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetType

`func (o *PublicEmail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicEmail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicEmail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PublicEmail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *PublicEmail) GetContent() PublicEmailContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PublicEmail) GetContentOk() (*PublicEmailContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PublicEmail) SetContent(v PublicEmailContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *PublicEmail) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *PublicEmail) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *PublicEmail) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *PublicEmail) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *PublicEmail) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetWebversion

`func (o *PublicEmail) GetWebversion() PublicWebversionDetails`

GetWebversion returns the Webversion field if non-nil, zero value otherwise.

### GetWebversionOk

`func (o *PublicEmail) GetWebversionOk() (*PublicWebversionDetails, bool)`

GetWebversionOk returns a tuple with the Webversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebversion

`func (o *PublicEmail) SetWebversion(v PublicWebversionDetails)`

SetWebversion sets Webversion field to given value.

### HasWebversion

`func (o *PublicEmail) HasWebversion() bool`

HasWebversion returns a boolean if a field has been set.

### GetArchived

`func (o *PublicEmail) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PublicEmail) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PublicEmail) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PublicEmail) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicEmail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicEmail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicEmail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicEmail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStats

`func (o *PublicEmail) GetStats() EmailStatisticsData`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *PublicEmail) GetStatsOk() (*EmailStatisticsData, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *PublicEmail) SetStats(v EmailStatisticsData)`

SetStats sets Stats field to given value.

### HasStats

`func (o *PublicEmail) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetFrom

`func (o *PublicEmail) GetFrom() PublicEmailFromDetails`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PublicEmail) GetFromOk() (*PublicEmailFromDetails, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PublicEmail) SetFrom(v PublicEmailFromDetails)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PublicEmail) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *PublicEmail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicEmail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicEmail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicEmail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *PublicEmail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PublicEmail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PublicEmail) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PublicEmail) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreatedById

`func (o *PublicEmail) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *PublicEmail) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *PublicEmail) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *PublicEmail) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicEmail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicEmail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicEmail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicEmail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRssData

`func (o *PublicEmail) GetRssData() PublicRssEmailDetails`

GetRssData returns the RssData field if non-nil, zero value otherwise.

### GetRssDataOk

`func (o *PublicEmail) GetRssDataOk() (*PublicRssEmailDetails, bool)`

GetRssDataOk returns a tuple with the RssData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssData

`func (o *PublicEmail) SetRssData(v PublicRssEmailDetails)`

SetRssData sets RssData field to given value.

### HasRssData

`func (o *PublicEmail) HasRssData() bool`

HasRssData returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PublicEmail) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PublicEmail) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PublicEmail) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *PublicEmail) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetPublishedById

`func (o *PublicEmail) GetPublishedById() string`

GetPublishedById returns the PublishedById field if non-nil, zero value otherwise.

### GetPublishedByIdOk

`func (o *PublicEmail) GetPublishedByIdOk() (*string, bool)`

GetPublishedByIdOk returns a tuple with the PublishedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedById

`func (o *PublicEmail) SetPublishedById(v string)`

SetPublishedById sets PublishedById field to given value.

### HasPublishedById

`func (o *PublicEmail) HasPublishedById() bool`

HasPublishedById returns a boolean if a field has been set.

### GetIsPublished

`func (o *PublicEmail) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *PublicEmail) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *PublicEmail) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *PublicEmail) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetTesting

`func (o *PublicEmail) GetTesting() PublicEmailTestingDetails`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *PublicEmail) GetTestingOk() (*PublicEmailTestingDetails, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *PublicEmail) SetTesting(v PublicEmailTestingDetails)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *PublicEmail) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetUpdatedById

`func (o *PublicEmail) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *PublicEmail) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *PublicEmail) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.

### HasUpdatedById

`func (o *PublicEmail) HasUpdatedById() bool`

HasUpdatedById returns a boolean if a field has been set.

### GetFolderId

`func (o *PublicEmail) GetFolderId() int64`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *PublicEmail) GetFolderIdOk() (*int64, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *PublicEmail) SetFolderId(v int64)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *PublicEmail) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetSubscriptionDetails

`func (o *PublicEmail) GetSubscriptionDetails() PublicEmailSubscriptionDetails`

GetSubscriptionDetails returns the SubscriptionDetails field if non-nil, zero value otherwise.

### GetSubscriptionDetailsOk

`func (o *PublicEmail) GetSubscriptionDetailsOk() (*PublicEmailSubscriptionDetails, bool)`

GetSubscriptionDetailsOk returns a tuple with the SubscriptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDetails

`func (o *PublicEmail) SetSubscriptionDetails(v PublicEmailSubscriptionDetails)`

SetSubscriptionDetails sets SubscriptionDetails field to given value.

### HasSubscriptionDetails

`func (o *PublicEmail) HasSubscriptionDetails() bool`

HasSubscriptionDetails returns a boolean if a field has been set.

### GetDeletedAt

`func (o *PublicEmail) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PublicEmail) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PublicEmail) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PublicEmail) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *PublicEmail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicEmail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicEmail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicEmail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActiveDomain

`func (o *PublicEmail) GetActiveDomain() string`

GetActiveDomain returns the ActiveDomain field if non-nil, zero value otherwise.

### GetActiveDomainOk

`func (o *PublicEmail) GetActiveDomainOk() (*string, bool)`

GetActiveDomainOk returns a tuple with the ActiveDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDomain

`func (o *PublicEmail) SetActiveDomain(v string)`

SetActiveDomain sets ActiveDomain field to given value.

### HasActiveDomain

`func (o *PublicEmail) HasActiveDomain() bool`

HasActiveDomain returns a boolean if a field has been set.

### GetCampaign

`func (o *PublicEmail) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *PublicEmail) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *PublicEmail) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *PublicEmail) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetTo

`func (o *PublicEmail) GetTo() PublicEmailToDetails`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PublicEmail) GetToOk() (*PublicEmailToDetails, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PublicEmail) SetTo(v PublicEmailToDetails)`

SetTo sets To field to given value.

### HasTo

`func (o *PublicEmail) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetSubcategory

`func (o *PublicEmail) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *PublicEmail) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *PublicEmail) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.

### HasSubcategory

`func (o *PublicEmail) HasSubcategory() bool`

HasSubcategory returns a boolean if a field has been set.

### GetSendOnPublish

`func (o *PublicEmail) GetSendOnPublish() bool`

GetSendOnPublish returns the SendOnPublish field if non-nil, zero value otherwise.

### GetSendOnPublishOk

`func (o *PublicEmail) GetSendOnPublishOk() (*bool, bool)`

GetSendOnPublishOk returns a tuple with the SendOnPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnPublish

`func (o *PublicEmail) SetSendOnPublish(v bool)`

SetSendOnPublish sets SendOnPublish field to given value.

### HasSendOnPublish

`func (o *PublicEmail) HasSendOnPublish() bool`

HasSendOnPublish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


