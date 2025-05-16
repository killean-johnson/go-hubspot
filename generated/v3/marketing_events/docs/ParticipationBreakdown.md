# ParticipationBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | [**ParticipationAssociations**](ParticipationAssociations.md) |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Properties** | [**ParticipationProperties**](ParticipationProperties.md) |  | 

## Methods

### NewParticipationBreakdown

`func NewParticipationBreakdown(associations ParticipationAssociations, createdAt time.Time, id string, properties ParticipationProperties, ) *ParticipationBreakdown`

NewParticipationBreakdown instantiates a new ParticipationBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipationBreakdownWithDefaults

`func NewParticipationBreakdownWithDefaults() *ParticipationBreakdown`

NewParticipationBreakdownWithDefaults instantiates a new ParticipationBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *ParticipationBreakdown) GetAssociations() ParticipationAssociations`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *ParticipationBreakdown) GetAssociationsOk() (*ParticipationAssociations, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *ParticipationBreakdown) SetAssociations(v ParticipationAssociations)`

SetAssociations sets Associations field to given value.


### GetCreatedAt

`func (o *ParticipationBreakdown) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ParticipationBreakdown) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ParticipationBreakdown) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *ParticipationBreakdown) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParticipationBreakdown) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParticipationBreakdown) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *ParticipationBreakdown) GetProperties() ParticipationProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ParticipationBreakdown) GetPropertiesOk() (*ParticipationProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ParticipationBreakdown) SetProperties(v ParticipationProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


