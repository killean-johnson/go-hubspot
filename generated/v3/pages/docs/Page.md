# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublishDate** | **time.Time** | The date (ISO8601 format) the page is to be published at. | 
**Language** | **string** | The explicitly defined ISO 639 language code of the page. If null, the page will default to the language of the Domain. | 
**EnableLayoutStylesheets** | **bool** | Boolean to determine whether or not the styles from the template should be applied. | 
**MetaDescription** | **string** | A description that goes in &lt;meta&gt; tag on the page. | 
**AttachedStylesheets** | **[]map[string]map[string]interface{}** | List of stylesheets to attach to this page. These stylesheets are attached to just this page. Order of precedence is bottom to top, just like in the HTML. | 
**Password** | **string** | Set this to create a password protected page. Entering the password will be required to view the page. | 
**PublishImmediately** | **bool** | Set this to true if you want to be published immediately when the schedule publish endpoint is called, and to ignore the publish_date setting. | 
**HtmlTitle** | **string** | The html title of this page. | 
**Translations** | [**map[string]ContentLanguageVariation**](ContentLanguageVariation.md) |  | 
**Id** | **string** | The unique ID of the page. | 
**State** | **string** | An ENUM descibing the current state of this page. | 
**Slug** | **string** | The path of the this page. This field is appended to the domain to construct the url of this page. | 
**CreatedById** | **string** | The ID of the user that created this page. | 
**CurrentlyPublished** | **bool** |  | 
**ArchivedInDashboard** | **bool** | If True, the page will not show up in your dashboard, although the page could still be live. | 
**Created** | **time.Time** |  | 
**ContentTypeCategory** | **string** | An ENUM descibing the type of this object. Should be either LANDING_PAGE or SITE_PAGE. | 
**MabExperimentId** | **string** | The ID of the MAB test (or dynamic test) associated with this page, if applicable | 
**UpdatedById** | **string** | The ID of the user that updated this page. | 
**TranslatedFromId** | **string** | ID of the primary page this object was translated from. | 
**FolderId** | **string** | The ID of the associated folder this landing page is organized under in the app dashboard. | 
**WidgetContainers** | **map[string]map[string]interface{}** | A data structure containing the data for all the modules inside the containers for this page. This will only be populated if the page has widget containers. | 
**PageExpiryRedirectId** | **int64** | The ID of another page this page&#39;s url should redirect to once this page expires. Should only set this or pageExpiryRedirectUrl. | 
**DynamicPageDataSourceType** | **int32** |  | 
**FeaturedImage** | **string** | The featuredImage of this page. | 
**AuthorName** | **string** | The name of the user that updated this page. | 
**Domain** | **string** | The domain this page will resolve to. If null, the page will default to the primary domain for this content type. | 
**Name** | **string** | The internal name of the page. | 
**DynamicPageHubDbTableId** | **string** | The ID of the HubDB table this page references, if applicable | 
**Campaign** | **string** | The GUID of the marketing campaign this page is a part of. | 
**DynamicPageDataSourceId** | **string** |  | 
**EnableDomainStylesheets** | **bool** | Boolean to determine whether or not the styles from the template should be applied. | 
**IncludeDefaultCustomCss** | **bool** | Boolean to determine whether or not the Primary CSS Files should be applied. | 
**Subcategory** | **string** | Details the type of page this is. Should always be landing_page or site_page | 
**LayoutSections** | [**map[string]LayoutSection**](LayoutSection.md) |  | 
**Updated** | **time.Time** |  | 
**FooterHtml** | **string** | Custom HTML for embed codes, javascript that should be placed before the &lt;/body&gt; tag of the page. | 
**Widgets** | **map[string]map[string]interface{}** | A data structure containing the data for all the modules for this page. | 
**HeadHtml** | **string** | Custom HTML for embed codes, javascript, etc. that goes in the &lt;head&gt; tag of the page. | 
**PageExpiryRedirectUrl** | **string** | The URL this page&#39;s url should redirect to once this page expires. Should only set this or pageExpiryRedirectId. | 
**AbStatus** | **string** | The status of the AB test associated with this page, if applicable | 
**UseFeaturedImage** | **bool** | Boolean to determine if this page should use a featuredImage. | 
**AbTestId** | **string** | The ID of the AB test associated with this page, if applicable | 
**FeaturedImageAltText** | **string** | Alt Text of the featuredImage. | 
**ContentGroupId** | **string** |  | 
**PageExpiryEnabled** | **bool** | Boolean describing if the page expiration feature is enabled for this page | 
**TemplatePath** | **string** | String detailing the path of the template used for this page. | 
**Url** | **string** | A generated field representing the URL of this page. | 
**PublicAccessRules** | **[]map[string]interface{}** | Rules for require member registration to access private content. | 
**ArchivedAt** | **time.Time** | The timestamp (ISO8601 format) when this page was deleted. | 
**ThemeSettingsValues** | **map[string]map[string]interface{}** |  | 
**PageExpiryDate** | **int64** | The date at which this page should expire and begin redirecting to another url or page. | 
**PublicAccessRulesEnabled** | **bool** | Boolean to determine whether or not to respect publicAccessRules. | 
**PageRedirected** | **bool** | A generated Boolean describing whether or not this page is currently expired and being redirected. | 
**CurrentState** | **string** | A generated ENUM descibing the current state of this page. | 
**CategoryId** | **int32** | ID of the type of object this is. Should always . | 
**LinkRelCanonicalUrl** | **string** | Optional override to set the URL to be used in the rel&#x3D;canonical link tag on the page. | 

## Methods

### NewPage

`func NewPage(publishDate time.Time, language string, enableLayoutStylesheets bool, metaDescription string, attachedStylesheets []map[string]map[string]interface{}, password string, publishImmediately bool, htmlTitle string, translations map[string]ContentLanguageVariation, id string, state string, slug string, createdById string, currentlyPublished bool, archivedInDashboard bool, created time.Time, contentTypeCategory string, mabExperimentId string, updatedById string, translatedFromId string, folderId string, widgetContainers map[string]map[string]interface{}, pageExpiryRedirectId int64, dynamicPageDataSourceType int32, featuredImage string, authorName string, domain string, name string, dynamicPageHubDbTableId string, campaign string, dynamicPageDataSourceId string, enableDomainStylesheets bool, includeDefaultCustomCss bool, subcategory string, layoutSections map[string]LayoutSection, updated time.Time, footerHtml string, widgets map[string]map[string]interface{}, headHtml string, pageExpiryRedirectUrl string, abStatus string, useFeaturedImage bool, abTestId string, featuredImageAltText string, contentGroupId string, pageExpiryEnabled bool, templatePath string, url string, publicAccessRules []map[string]interface{}, archivedAt time.Time, themeSettingsValues map[string]map[string]interface{}, pageExpiryDate int64, publicAccessRulesEnabled bool, pageRedirected bool, currentState string, categoryId int32, linkRelCanonicalUrl string, ) *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublishDate

`func (o *Page) GetPublishDate() time.Time`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *Page) GetPublishDateOk() (*time.Time, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *Page) SetPublishDate(v time.Time)`

SetPublishDate sets PublishDate field to given value.


### GetLanguage

`func (o *Page) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Page) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Page) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetEnableLayoutStylesheets

`func (o *Page) GetEnableLayoutStylesheets() bool`

GetEnableLayoutStylesheets returns the EnableLayoutStylesheets field if non-nil, zero value otherwise.

### GetEnableLayoutStylesheetsOk

`func (o *Page) GetEnableLayoutStylesheetsOk() (*bool, bool)`

GetEnableLayoutStylesheetsOk returns a tuple with the EnableLayoutStylesheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLayoutStylesheets

`func (o *Page) SetEnableLayoutStylesheets(v bool)`

SetEnableLayoutStylesheets sets EnableLayoutStylesheets field to given value.


### GetMetaDescription

`func (o *Page) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *Page) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *Page) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.


### GetAttachedStylesheets

`func (o *Page) GetAttachedStylesheets() []map[string]map[string]interface{}`

GetAttachedStylesheets returns the AttachedStylesheets field if non-nil, zero value otherwise.

### GetAttachedStylesheetsOk

`func (o *Page) GetAttachedStylesheetsOk() (*[]map[string]map[string]interface{}, bool)`

GetAttachedStylesheetsOk returns a tuple with the AttachedStylesheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedStylesheets

`func (o *Page) SetAttachedStylesheets(v []map[string]map[string]interface{})`

SetAttachedStylesheets sets AttachedStylesheets field to given value.


### GetPassword

`func (o *Page) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Page) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Page) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPublishImmediately

`func (o *Page) GetPublishImmediately() bool`

GetPublishImmediately returns the PublishImmediately field if non-nil, zero value otherwise.

### GetPublishImmediatelyOk

`func (o *Page) GetPublishImmediatelyOk() (*bool, bool)`

GetPublishImmediatelyOk returns a tuple with the PublishImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishImmediately

`func (o *Page) SetPublishImmediately(v bool)`

SetPublishImmediately sets PublishImmediately field to given value.


### GetHtmlTitle

`func (o *Page) GetHtmlTitle() string`

GetHtmlTitle returns the HtmlTitle field if non-nil, zero value otherwise.

### GetHtmlTitleOk

`func (o *Page) GetHtmlTitleOk() (*string, bool)`

GetHtmlTitleOk returns a tuple with the HtmlTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlTitle

`func (o *Page) SetHtmlTitle(v string)`

SetHtmlTitle sets HtmlTitle field to given value.


### GetTranslations

`func (o *Page) GetTranslations() map[string]ContentLanguageVariation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *Page) GetTranslationsOk() (*map[string]ContentLanguageVariation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *Page) SetTranslations(v map[string]ContentLanguageVariation)`

SetTranslations sets Translations field to given value.


### GetId

`func (o *Page) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Page) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Page) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *Page) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Page) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Page) SetState(v string)`

SetState sets State field to given value.


### GetSlug

`func (o *Page) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Page) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Page) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetCreatedById

`func (o *Page) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Page) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Page) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCurrentlyPublished

`func (o *Page) GetCurrentlyPublished() bool`

GetCurrentlyPublished returns the CurrentlyPublished field if non-nil, zero value otherwise.

### GetCurrentlyPublishedOk

`func (o *Page) GetCurrentlyPublishedOk() (*bool, bool)`

GetCurrentlyPublishedOk returns a tuple with the CurrentlyPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyPublished

`func (o *Page) SetCurrentlyPublished(v bool)`

SetCurrentlyPublished sets CurrentlyPublished field to given value.


### GetArchivedInDashboard

`func (o *Page) GetArchivedInDashboard() bool`

GetArchivedInDashboard returns the ArchivedInDashboard field if non-nil, zero value otherwise.

### GetArchivedInDashboardOk

`func (o *Page) GetArchivedInDashboardOk() (*bool, bool)`

GetArchivedInDashboardOk returns a tuple with the ArchivedInDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedInDashboard

`func (o *Page) SetArchivedInDashboard(v bool)`

SetArchivedInDashboard sets ArchivedInDashboard field to given value.


### GetCreated

`func (o *Page) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Page) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Page) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetContentTypeCategory

`func (o *Page) GetContentTypeCategory() string`

GetContentTypeCategory returns the ContentTypeCategory field if non-nil, zero value otherwise.

### GetContentTypeCategoryOk

`func (o *Page) GetContentTypeCategoryOk() (*string, bool)`

GetContentTypeCategoryOk returns a tuple with the ContentTypeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypeCategory

`func (o *Page) SetContentTypeCategory(v string)`

SetContentTypeCategory sets ContentTypeCategory field to given value.


### GetMabExperimentId

`func (o *Page) GetMabExperimentId() string`

GetMabExperimentId returns the MabExperimentId field if non-nil, zero value otherwise.

### GetMabExperimentIdOk

`func (o *Page) GetMabExperimentIdOk() (*string, bool)`

GetMabExperimentIdOk returns a tuple with the MabExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMabExperimentId

`func (o *Page) SetMabExperimentId(v string)`

SetMabExperimentId sets MabExperimentId field to given value.


### GetUpdatedById

`func (o *Page) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Page) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Page) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetTranslatedFromId

`func (o *Page) GetTranslatedFromId() string`

GetTranslatedFromId returns the TranslatedFromId field if non-nil, zero value otherwise.

### GetTranslatedFromIdOk

`func (o *Page) GetTranslatedFromIdOk() (*string, bool)`

GetTranslatedFromIdOk returns a tuple with the TranslatedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedFromId

`func (o *Page) SetTranslatedFromId(v string)`

SetTranslatedFromId sets TranslatedFromId field to given value.


### GetFolderId

`func (o *Page) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Page) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Page) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### GetWidgetContainers

`func (o *Page) GetWidgetContainers() map[string]map[string]interface{}`

GetWidgetContainers returns the WidgetContainers field if non-nil, zero value otherwise.

### GetWidgetContainersOk

`func (o *Page) GetWidgetContainersOk() (*map[string]map[string]interface{}, bool)`

GetWidgetContainersOk returns a tuple with the WidgetContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetContainers

`func (o *Page) SetWidgetContainers(v map[string]map[string]interface{})`

SetWidgetContainers sets WidgetContainers field to given value.


### GetPageExpiryRedirectId

`func (o *Page) GetPageExpiryRedirectId() int64`

GetPageExpiryRedirectId returns the PageExpiryRedirectId field if non-nil, zero value otherwise.

### GetPageExpiryRedirectIdOk

`func (o *Page) GetPageExpiryRedirectIdOk() (*int64, bool)`

GetPageExpiryRedirectIdOk returns a tuple with the PageExpiryRedirectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiryRedirectId

`func (o *Page) SetPageExpiryRedirectId(v int64)`

SetPageExpiryRedirectId sets PageExpiryRedirectId field to given value.


### GetDynamicPageDataSourceType

`func (o *Page) GetDynamicPageDataSourceType() int32`

GetDynamicPageDataSourceType returns the DynamicPageDataSourceType field if non-nil, zero value otherwise.

### GetDynamicPageDataSourceTypeOk

`func (o *Page) GetDynamicPageDataSourceTypeOk() (*int32, bool)`

GetDynamicPageDataSourceTypeOk returns a tuple with the DynamicPageDataSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPageDataSourceType

`func (o *Page) SetDynamicPageDataSourceType(v int32)`

SetDynamicPageDataSourceType sets DynamicPageDataSourceType field to given value.


### GetFeaturedImage

`func (o *Page) GetFeaturedImage() string`

GetFeaturedImage returns the FeaturedImage field if non-nil, zero value otherwise.

### GetFeaturedImageOk

`func (o *Page) GetFeaturedImageOk() (*string, bool)`

GetFeaturedImageOk returns a tuple with the FeaturedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImage

`func (o *Page) SetFeaturedImage(v string)`

SetFeaturedImage sets FeaturedImage field to given value.


### GetAuthorName

`func (o *Page) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *Page) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *Page) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetDomain

`func (o *Page) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Page) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Page) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *Page) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Page) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Page) SetName(v string)`

SetName sets Name field to given value.


### GetDynamicPageHubDbTableId

`func (o *Page) GetDynamicPageHubDbTableId() string`

GetDynamicPageHubDbTableId returns the DynamicPageHubDbTableId field if non-nil, zero value otherwise.

### GetDynamicPageHubDbTableIdOk

`func (o *Page) GetDynamicPageHubDbTableIdOk() (*string, bool)`

GetDynamicPageHubDbTableIdOk returns a tuple with the DynamicPageHubDbTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPageHubDbTableId

`func (o *Page) SetDynamicPageHubDbTableId(v string)`

SetDynamicPageHubDbTableId sets DynamicPageHubDbTableId field to given value.


### GetCampaign

`func (o *Page) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *Page) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *Page) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.


### GetDynamicPageDataSourceId

`func (o *Page) GetDynamicPageDataSourceId() string`

GetDynamicPageDataSourceId returns the DynamicPageDataSourceId field if non-nil, zero value otherwise.

### GetDynamicPageDataSourceIdOk

`func (o *Page) GetDynamicPageDataSourceIdOk() (*string, bool)`

GetDynamicPageDataSourceIdOk returns a tuple with the DynamicPageDataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPageDataSourceId

`func (o *Page) SetDynamicPageDataSourceId(v string)`

SetDynamicPageDataSourceId sets DynamicPageDataSourceId field to given value.


### GetEnableDomainStylesheets

`func (o *Page) GetEnableDomainStylesheets() bool`

GetEnableDomainStylesheets returns the EnableDomainStylesheets field if non-nil, zero value otherwise.

### GetEnableDomainStylesheetsOk

`func (o *Page) GetEnableDomainStylesheetsOk() (*bool, bool)`

GetEnableDomainStylesheetsOk returns a tuple with the EnableDomainStylesheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDomainStylesheets

`func (o *Page) SetEnableDomainStylesheets(v bool)`

SetEnableDomainStylesheets sets EnableDomainStylesheets field to given value.


### GetIncludeDefaultCustomCss

`func (o *Page) GetIncludeDefaultCustomCss() bool`

GetIncludeDefaultCustomCss returns the IncludeDefaultCustomCss field if non-nil, zero value otherwise.

### GetIncludeDefaultCustomCssOk

`func (o *Page) GetIncludeDefaultCustomCssOk() (*bool, bool)`

GetIncludeDefaultCustomCssOk returns a tuple with the IncludeDefaultCustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDefaultCustomCss

`func (o *Page) SetIncludeDefaultCustomCss(v bool)`

SetIncludeDefaultCustomCss sets IncludeDefaultCustomCss field to given value.


### GetSubcategory

`func (o *Page) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *Page) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *Page) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.


### GetLayoutSections

`func (o *Page) GetLayoutSections() map[string]LayoutSection`

GetLayoutSections returns the LayoutSections field if non-nil, zero value otherwise.

### GetLayoutSectionsOk

`func (o *Page) GetLayoutSectionsOk() (*map[string]LayoutSection, bool)`

GetLayoutSectionsOk returns a tuple with the LayoutSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSections

`func (o *Page) SetLayoutSections(v map[string]LayoutSection)`

SetLayoutSections sets LayoutSections field to given value.


### GetUpdated

`func (o *Page) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Page) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Page) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetFooterHtml

`func (o *Page) GetFooterHtml() string`

GetFooterHtml returns the FooterHtml field if non-nil, zero value otherwise.

### GetFooterHtmlOk

`func (o *Page) GetFooterHtmlOk() (*string, bool)`

GetFooterHtmlOk returns a tuple with the FooterHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterHtml

`func (o *Page) SetFooterHtml(v string)`

SetFooterHtml sets FooterHtml field to given value.


### GetWidgets

`func (o *Page) GetWidgets() map[string]map[string]interface{}`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *Page) GetWidgetsOk() (*map[string]map[string]interface{}, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *Page) SetWidgets(v map[string]map[string]interface{})`

SetWidgets sets Widgets field to given value.


### GetHeadHtml

`func (o *Page) GetHeadHtml() string`

GetHeadHtml returns the HeadHtml field if non-nil, zero value otherwise.

### GetHeadHtmlOk

`func (o *Page) GetHeadHtmlOk() (*string, bool)`

GetHeadHtmlOk returns a tuple with the HeadHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadHtml

`func (o *Page) SetHeadHtml(v string)`

SetHeadHtml sets HeadHtml field to given value.


### GetPageExpiryRedirectUrl

`func (o *Page) GetPageExpiryRedirectUrl() string`

GetPageExpiryRedirectUrl returns the PageExpiryRedirectUrl field if non-nil, zero value otherwise.

### GetPageExpiryRedirectUrlOk

`func (o *Page) GetPageExpiryRedirectUrlOk() (*string, bool)`

GetPageExpiryRedirectUrlOk returns a tuple with the PageExpiryRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiryRedirectUrl

`func (o *Page) SetPageExpiryRedirectUrl(v string)`

SetPageExpiryRedirectUrl sets PageExpiryRedirectUrl field to given value.


### GetAbStatus

`func (o *Page) GetAbStatus() string`

GetAbStatus returns the AbStatus field if non-nil, zero value otherwise.

### GetAbStatusOk

`func (o *Page) GetAbStatusOk() (*string, bool)`

GetAbStatusOk returns a tuple with the AbStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbStatus

`func (o *Page) SetAbStatus(v string)`

SetAbStatus sets AbStatus field to given value.


### GetUseFeaturedImage

`func (o *Page) GetUseFeaturedImage() bool`

GetUseFeaturedImage returns the UseFeaturedImage field if non-nil, zero value otherwise.

### GetUseFeaturedImageOk

`func (o *Page) GetUseFeaturedImageOk() (*bool, bool)`

GetUseFeaturedImageOk returns a tuple with the UseFeaturedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFeaturedImage

`func (o *Page) SetUseFeaturedImage(v bool)`

SetUseFeaturedImage sets UseFeaturedImage field to given value.


### GetAbTestId

`func (o *Page) GetAbTestId() string`

GetAbTestId returns the AbTestId field if non-nil, zero value otherwise.

### GetAbTestIdOk

`func (o *Page) GetAbTestIdOk() (*string, bool)`

GetAbTestIdOk returns a tuple with the AbTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbTestId

`func (o *Page) SetAbTestId(v string)`

SetAbTestId sets AbTestId field to given value.


### GetFeaturedImageAltText

`func (o *Page) GetFeaturedImageAltText() string`

GetFeaturedImageAltText returns the FeaturedImageAltText field if non-nil, zero value otherwise.

### GetFeaturedImageAltTextOk

`func (o *Page) GetFeaturedImageAltTextOk() (*string, bool)`

GetFeaturedImageAltTextOk returns a tuple with the FeaturedImageAltText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImageAltText

`func (o *Page) SetFeaturedImageAltText(v string)`

SetFeaturedImageAltText sets FeaturedImageAltText field to given value.


### GetContentGroupId

`func (o *Page) GetContentGroupId() string`

GetContentGroupId returns the ContentGroupId field if non-nil, zero value otherwise.

### GetContentGroupIdOk

`func (o *Page) GetContentGroupIdOk() (*string, bool)`

GetContentGroupIdOk returns a tuple with the ContentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGroupId

`func (o *Page) SetContentGroupId(v string)`

SetContentGroupId sets ContentGroupId field to given value.


### GetPageExpiryEnabled

`func (o *Page) GetPageExpiryEnabled() bool`

GetPageExpiryEnabled returns the PageExpiryEnabled field if non-nil, zero value otherwise.

### GetPageExpiryEnabledOk

`func (o *Page) GetPageExpiryEnabledOk() (*bool, bool)`

GetPageExpiryEnabledOk returns a tuple with the PageExpiryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiryEnabled

`func (o *Page) SetPageExpiryEnabled(v bool)`

SetPageExpiryEnabled sets PageExpiryEnabled field to given value.


### GetTemplatePath

`func (o *Page) GetTemplatePath() string`

GetTemplatePath returns the TemplatePath field if non-nil, zero value otherwise.

### GetTemplatePathOk

`func (o *Page) GetTemplatePathOk() (*string, bool)`

GetTemplatePathOk returns a tuple with the TemplatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePath

`func (o *Page) SetTemplatePath(v string)`

SetTemplatePath sets TemplatePath field to given value.


### GetUrl

`func (o *Page) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Page) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Page) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPublicAccessRules

`func (o *Page) GetPublicAccessRules() []map[string]interface{}`

GetPublicAccessRules returns the PublicAccessRules field if non-nil, zero value otherwise.

### GetPublicAccessRulesOk

`func (o *Page) GetPublicAccessRulesOk() (*[]map[string]interface{}, bool)`

GetPublicAccessRulesOk returns a tuple with the PublicAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessRules

`func (o *Page) SetPublicAccessRules(v []map[string]interface{})`

SetPublicAccessRules sets PublicAccessRules field to given value.


### GetArchivedAt

`func (o *Page) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Page) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Page) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.


### GetThemeSettingsValues

`func (o *Page) GetThemeSettingsValues() map[string]map[string]interface{}`

GetThemeSettingsValues returns the ThemeSettingsValues field if non-nil, zero value otherwise.

### GetThemeSettingsValuesOk

`func (o *Page) GetThemeSettingsValuesOk() (*map[string]map[string]interface{}, bool)`

GetThemeSettingsValuesOk returns a tuple with the ThemeSettingsValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeSettingsValues

`func (o *Page) SetThemeSettingsValues(v map[string]map[string]interface{})`

SetThemeSettingsValues sets ThemeSettingsValues field to given value.


### GetPageExpiryDate

`func (o *Page) GetPageExpiryDate() int64`

GetPageExpiryDate returns the PageExpiryDate field if non-nil, zero value otherwise.

### GetPageExpiryDateOk

`func (o *Page) GetPageExpiryDateOk() (*int64, bool)`

GetPageExpiryDateOk returns a tuple with the PageExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageExpiryDate

`func (o *Page) SetPageExpiryDate(v int64)`

SetPageExpiryDate sets PageExpiryDate field to given value.


### GetPublicAccessRulesEnabled

`func (o *Page) GetPublicAccessRulesEnabled() bool`

GetPublicAccessRulesEnabled returns the PublicAccessRulesEnabled field if non-nil, zero value otherwise.

### GetPublicAccessRulesEnabledOk

`func (o *Page) GetPublicAccessRulesEnabledOk() (*bool, bool)`

GetPublicAccessRulesEnabledOk returns a tuple with the PublicAccessRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessRulesEnabled

`func (o *Page) SetPublicAccessRulesEnabled(v bool)`

SetPublicAccessRulesEnabled sets PublicAccessRulesEnabled field to given value.


### GetPageRedirected

`func (o *Page) GetPageRedirected() bool`

GetPageRedirected returns the PageRedirected field if non-nil, zero value otherwise.

### GetPageRedirectedOk

`func (o *Page) GetPageRedirectedOk() (*bool, bool)`

GetPageRedirectedOk returns a tuple with the PageRedirected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageRedirected

`func (o *Page) SetPageRedirected(v bool)`

SetPageRedirected sets PageRedirected field to given value.


### GetCurrentState

`func (o *Page) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *Page) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *Page) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.


### GetCategoryId

`func (o *Page) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Page) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Page) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetLinkRelCanonicalUrl

`func (o *Page) GetLinkRelCanonicalUrl() string`

GetLinkRelCanonicalUrl returns the LinkRelCanonicalUrl field if non-nil, zero value otherwise.

### GetLinkRelCanonicalUrlOk

`func (o *Page) GetLinkRelCanonicalUrlOk() (*string, bool)`

GetLinkRelCanonicalUrlOk returns a tuple with the LinkRelCanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRelCanonicalUrl

`func (o *Page) SetLinkRelCanonicalUrl(v string)`

SetLinkRelCanonicalUrl sets LinkRelCanonicalUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


