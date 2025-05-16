# FieldGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupType** | **string** |  | 
**RichTextType** | **string** | The type of rich text included. The default value is text. | 
**RichText** | Pointer to **string** | A block of rich text or an image. Those can be used to add extra information for the customers filling in the form. If the field group includes fields, the rich text will be displayed before the fields. | [optional] 
**Fields** | [**[]DependentFieldDependentField**](DependentFieldDependentField.md) | The form fields included in the group | 

## Methods

### NewFieldGroup

`func NewFieldGroup(groupType string, richTextType string, fields []DependentFieldDependentField, ) *FieldGroup`

NewFieldGroup instantiates a new FieldGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldGroupWithDefaults

`func NewFieldGroupWithDefaults() *FieldGroup`

NewFieldGroupWithDefaults instantiates a new FieldGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupType

`func (o *FieldGroup) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *FieldGroup) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *FieldGroup) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.


### GetRichTextType

`func (o *FieldGroup) GetRichTextType() string`

GetRichTextType returns the RichTextType field if non-nil, zero value otherwise.

### GetRichTextTypeOk

`func (o *FieldGroup) GetRichTextTypeOk() (*string, bool)`

GetRichTextTypeOk returns a tuple with the RichTextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichTextType

`func (o *FieldGroup) SetRichTextType(v string)`

SetRichTextType sets RichTextType field to given value.


### GetRichText

`func (o *FieldGroup) GetRichText() string`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *FieldGroup) GetRichTextOk() (*string, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *FieldGroup) SetRichText(v string)`

SetRichText sets RichText field to given value.

### HasRichText

`func (o *FieldGroup) HasRichText() bool`

HasRichText returns a boolean if a field has been set.

### GetFields

`func (o *FieldGroup) GetFields() []DependentFieldDependentField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FieldGroup) GetFieldsOk() (*[]DependentFieldDependentField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FieldGroup) SetFields(v []DependentFieldDependentField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


