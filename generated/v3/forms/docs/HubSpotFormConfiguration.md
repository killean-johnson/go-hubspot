# HubSpotFormConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateNewContactForNewEmail** | **bool** | Whether to create a new contact when a form is submitted with an email address that doesn’t match any in your existing contacts records. | 
**Editable** | **bool** | Whether the form can be edited. | 
**AllowLinkToResetKnownValues** | **bool** | Whether to add a reset link to the form. This removes any pre-populated content on the form and creates a new contact on submission. | 
**LifecycleStages** | Pointer to [**[]LifecycleStage**](LifecycleStage.md) |  | [optional] 
**PostSubmitAction** | [**FormPostSubmitAction**](FormPostSubmitAction.md) |  | 
**Language** | **string** | The language of the form. | 
**PrePopulateKnownValues** | **bool** | Whether contact fields should pre-populate with known information when a contact returns to your site. | 
**Cloneable** | **bool** | Whether the form can be cloned. | 
**NotifyContactOwner** | **bool** | Whether to send a notification email to the contact owner when a submission is received. | 
**RecaptchaEnabled** | **bool** | Whether CAPTCHA (spam prevention) is enabled. | 
**Archivable** | **bool** | Whether the form can be archived. | 
**NotifyRecipients** | **[]string** | The list of user IDs to receive a notification email when a submission is received. | 

## Methods

### NewHubSpotFormConfiguration

`func NewHubSpotFormConfiguration(createNewContactForNewEmail bool, editable bool, allowLinkToResetKnownValues bool, postSubmitAction FormPostSubmitAction, language string, prePopulateKnownValues bool, cloneable bool, notifyContactOwner bool, recaptchaEnabled bool, archivable bool, notifyRecipients []string, ) *HubSpotFormConfiguration`

NewHubSpotFormConfiguration instantiates a new HubSpotFormConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubSpotFormConfigurationWithDefaults

`func NewHubSpotFormConfigurationWithDefaults() *HubSpotFormConfiguration`

NewHubSpotFormConfigurationWithDefaults instantiates a new HubSpotFormConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateNewContactForNewEmail

`func (o *HubSpotFormConfiguration) GetCreateNewContactForNewEmail() bool`

GetCreateNewContactForNewEmail returns the CreateNewContactForNewEmail field if non-nil, zero value otherwise.

### GetCreateNewContactForNewEmailOk

`func (o *HubSpotFormConfiguration) GetCreateNewContactForNewEmailOk() (*bool, bool)`

GetCreateNewContactForNewEmailOk returns a tuple with the CreateNewContactForNewEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewContactForNewEmail

`func (o *HubSpotFormConfiguration) SetCreateNewContactForNewEmail(v bool)`

SetCreateNewContactForNewEmail sets CreateNewContactForNewEmail field to given value.


### GetEditable

`func (o *HubSpotFormConfiguration) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *HubSpotFormConfiguration) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *HubSpotFormConfiguration) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetAllowLinkToResetKnownValues

`func (o *HubSpotFormConfiguration) GetAllowLinkToResetKnownValues() bool`

GetAllowLinkToResetKnownValues returns the AllowLinkToResetKnownValues field if non-nil, zero value otherwise.

### GetAllowLinkToResetKnownValuesOk

`func (o *HubSpotFormConfiguration) GetAllowLinkToResetKnownValuesOk() (*bool, bool)`

GetAllowLinkToResetKnownValuesOk returns a tuple with the AllowLinkToResetKnownValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLinkToResetKnownValues

`func (o *HubSpotFormConfiguration) SetAllowLinkToResetKnownValues(v bool)`

SetAllowLinkToResetKnownValues sets AllowLinkToResetKnownValues field to given value.


### GetLifecycleStages

`func (o *HubSpotFormConfiguration) GetLifecycleStages() []LifecycleStage`

GetLifecycleStages returns the LifecycleStages field if non-nil, zero value otherwise.

### GetLifecycleStagesOk

`func (o *HubSpotFormConfiguration) GetLifecycleStagesOk() (*[]LifecycleStage, bool)`

GetLifecycleStagesOk returns a tuple with the LifecycleStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStages

`func (o *HubSpotFormConfiguration) SetLifecycleStages(v []LifecycleStage)`

SetLifecycleStages sets LifecycleStages field to given value.

### HasLifecycleStages

`func (o *HubSpotFormConfiguration) HasLifecycleStages() bool`

HasLifecycleStages returns a boolean if a field has been set.

### GetPostSubmitAction

`func (o *HubSpotFormConfiguration) GetPostSubmitAction() FormPostSubmitAction`

GetPostSubmitAction returns the PostSubmitAction field if non-nil, zero value otherwise.

### GetPostSubmitActionOk

`func (o *HubSpotFormConfiguration) GetPostSubmitActionOk() (*FormPostSubmitAction, bool)`

GetPostSubmitActionOk returns a tuple with the PostSubmitAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSubmitAction

`func (o *HubSpotFormConfiguration) SetPostSubmitAction(v FormPostSubmitAction)`

SetPostSubmitAction sets PostSubmitAction field to given value.


### GetLanguage

`func (o *HubSpotFormConfiguration) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *HubSpotFormConfiguration) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *HubSpotFormConfiguration) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetPrePopulateKnownValues

`func (o *HubSpotFormConfiguration) GetPrePopulateKnownValues() bool`

GetPrePopulateKnownValues returns the PrePopulateKnownValues field if non-nil, zero value otherwise.

### GetPrePopulateKnownValuesOk

`func (o *HubSpotFormConfiguration) GetPrePopulateKnownValuesOk() (*bool, bool)`

GetPrePopulateKnownValuesOk returns a tuple with the PrePopulateKnownValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePopulateKnownValues

`func (o *HubSpotFormConfiguration) SetPrePopulateKnownValues(v bool)`

SetPrePopulateKnownValues sets PrePopulateKnownValues field to given value.


### GetCloneable

`func (o *HubSpotFormConfiguration) GetCloneable() bool`

GetCloneable returns the Cloneable field if non-nil, zero value otherwise.

### GetCloneableOk

`func (o *HubSpotFormConfiguration) GetCloneableOk() (*bool, bool)`

GetCloneableOk returns a tuple with the Cloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneable

`func (o *HubSpotFormConfiguration) SetCloneable(v bool)`

SetCloneable sets Cloneable field to given value.


### GetNotifyContactOwner

`func (o *HubSpotFormConfiguration) GetNotifyContactOwner() bool`

GetNotifyContactOwner returns the NotifyContactOwner field if non-nil, zero value otherwise.

### GetNotifyContactOwnerOk

`func (o *HubSpotFormConfiguration) GetNotifyContactOwnerOk() (*bool, bool)`

GetNotifyContactOwnerOk returns a tuple with the NotifyContactOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyContactOwner

`func (o *HubSpotFormConfiguration) SetNotifyContactOwner(v bool)`

SetNotifyContactOwner sets NotifyContactOwner field to given value.


### GetRecaptchaEnabled

`func (o *HubSpotFormConfiguration) GetRecaptchaEnabled() bool`

GetRecaptchaEnabled returns the RecaptchaEnabled field if non-nil, zero value otherwise.

### GetRecaptchaEnabledOk

`func (o *HubSpotFormConfiguration) GetRecaptchaEnabledOk() (*bool, bool)`

GetRecaptchaEnabledOk returns a tuple with the RecaptchaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaEnabled

`func (o *HubSpotFormConfiguration) SetRecaptchaEnabled(v bool)`

SetRecaptchaEnabled sets RecaptchaEnabled field to given value.


### GetArchivable

`func (o *HubSpotFormConfiguration) GetArchivable() bool`

GetArchivable returns the Archivable field if non-nil, zero value otherwise.

### GetArchivableOk

`func (o *HubSpotFormConfiguration) GetArchivableOk() (*bool, bool)`

GetArchivableOk returns a tuple with the Archivable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivable

`func (o *HubSpotFormConfiguration) SetArchivable(v bool)`

SetArchivable sets Archivable field to given value.


### GetNotifyRecipients

`func (o *HubSpotFormConfiguration) GetNotifyRecipients() []string`

GetNotifyRecipients returns the NotifyRecipients field if non-nil, zero value otherwise.

### GetNotifyRecipientsOk

`func (o *HubSpotFormConfiguration) GetNotifyRecipientsOk() (*[]string, bool)`

GetNotifyRecipientsOk returns a tuple with the NotifyRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyRecipients

`func (o *HubSpotFormConfiguration) SetNotifyRecipients(v []string)`

SetNotifyRecipients sets NotifyRecipients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


