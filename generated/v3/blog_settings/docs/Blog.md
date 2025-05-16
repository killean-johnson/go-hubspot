# Blog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**AbsoluteUrl** | **string** |  | 
**Description** | **string** | The Description of this Blog. | 
**Language** | **string** | The explicitly defined language of the Blog. If null, the Blog will default to the language of the Domain. | 
**TranslatedFromId** | **string** | ID of the primary Blog this object was translated from. | 
**PublicAccessRules** | **[]map[string]interface{}** | Rules for require member registration to access private content. | 
**PublicTitle** | **string** | The public title of this Blog. | 
**AllowComments** | **bool** | Boolean determining whether or not this blog allows public comments. | 
**DeletedAt** | **time.Time** | The timestamp (ISO8601 format) when this Blog was deleted. | 
**HtmlTitle** | **string** | The html title of this Blog. | 
**PublicAccessRulesEnabled** | **bool** | Boolean to determine whether or not to respect publicAccessRules. | 
**Name** | **string** | The internal name of the blog. | 
**Id** | **string** | The unique ID of the Blog. | 
**Updated** | **time.Time** |  | 
**Slug** | **string** | The path of the this blog. This field is appended to the domain to construct the url of this blog. | 

## Methods

### NewBlog

`func NewBlog(created time.Time, absoluteUrl string, description string, language string, translatedFromId string, publicAccessRules []map[string]interface{}, publicTitle string, allowComments bool, deletedAt time.Time, htmlTitle string, publicAccessRulesEnabled bool, name string, id string, updated time.Time, slug string, ) *Blog`

NewBlog instantiates a new Blog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogWithDefaults

`func NewBlogWithDefaults() *Blog`

NewBlogWithDefaults instantiates a new Blog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Blog) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Blog) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Blog) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAbsoluteUrl

`func (o *Blog) GetAbsoluteUrl() string`

GetAbsoluteUrl returns the AbsoluteUrl field if non-nil, zero value otherwise.

### GetAbsoluteUrlOk

`func (o *Blog) GetAbsoluteUrlOk() (*string, bool)`

GetAbsoluteUrlOk returns a tuple with the AbsoluteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteUrl

`func (o *Blog) SetAbsoluteUrl(v string)`

SetAbsoluteUrl sets AbsoluteUrl field to given value.


### GetDescription

`func (o *Blog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Blog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Blog) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLanguage

`func (o *Blog) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Blog) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Blog) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetTranslatedFromId

`func (o *Blog) GetTranslatedFromId() string`

GetTranslatedFromId returns the TranslatedFromId field if non-nil, zero value otherwise.

### GetTranslatedFromIdOk

`func (o *Blog) GetTranslatedFromIdOk() (*string, bool)`

GetTranslatedFromIdOk returns a tuple with the TranslatedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedFromId

`func (o *Blog) SetTranslatedFromId(v string)`

SetTranslatedFromId sets TranslatedFromId field to given value.


### GetPublicAccessRules

`func (o *Blog) GetPublicAccessRules() []map[string]interface{}`

GetPublicAccessRules returns the PublicAccessRules field if non-nil, zero value otherwise.

### GetPublicAccessRulesOk

`func (o *Blog) GetPublicAccessRulesOk() (*[]map[string]interface{}, bool)`

GetPublicAccessRulesOk returns a tuple with the PublicAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessRules

`func (o *Blog) SetPublicAccessRules(v []map[string]interface{})`

SetPublicAccessRules sets PublicAccessRules field to given value.


### GetPublicTitle

`func (o *Blog) GetPublicTitle() string`

GetPublicTitle returns the PublicTitle field if non-nil, zero value otherwise.

### GetPublicTitleOk

`func (o *Blog) GetPublicTitleOk() (*string, bool)`

GetPublicTitleOk returns a tuple with the PublicTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicTitle

`func (o *Blog) SetPublicTitle(v string)`

SetPublicTitle sets PublicTitle field to given value.


### GetAllowComments

`func (o *Blog) GetAllowComments() bool`

GetAllowComments returns the AllowComments field if non-nil, zero value otherwise.

### GetAllowCommentsOk

`func (o *Blog) GetAllowCommentsOk() (*bool, bool)`

GetAllowCommentsOk returns a tuple with the AllowComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowComments

`func (o *Blog) SetAllowComments(v bool)`

SetAllowComments sets AllowComments field to given value.


### GetDeletedAt

`func (o *Blog) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Blog) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Blog) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetHtmlTitle

`func (o *Blog) GetHtmlTitle() string`

GetHtmlTitle returns the HtmlTitle field if non-nil, zero value otherwise.

### GetHtmlTitleOk

`func (o *Blog) GetHtmlTitleOk() (*string, bool)`

GetHtmlTitleOk returns a tuple with the HtmlTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlTitle

`func (o *Blog) SetHtmlTitle(v string)`

SetHtmlTitle sets HtmlTitle field to given value.


### GetPublicAccessRulesEnabled

`func (o *Blog) GetPublicAccessRulesEnabled() bool`

GetPublicAccessRulesEnabled returns the PublicAccessRulesEnabled field if non-nil, zero value otherwise.

### GetPublicAccessRulesEnabledOk

`func (o *Blog) GetPublicAccessRulesEnabledOk() (*bool, bool)`

GetPublicAccessRulesEnabledOk returns a tuple with the PublicAccessRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessRulesEnabled

`func (o *Blog) SetPublicAccessRulesEnabled(v bool)`

SetPublicAccessRulesEnabled sets PublicAccessRulesEnabled field to given value.


### GetName

`func (o *Blog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Blog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Blog) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *Blog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Blog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Blog) SetId(v string)`

SetId sets Id field to given value.


### GetUpdated

`func (o *Blog) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Blog) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Blog) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetSlug

`func (o *Blog) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Blog) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Blog) SetSlug(v string)`

SetSlug sets Slug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


