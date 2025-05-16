# PublicBatchMigrationMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegacyListIdsToIdsMapping** | [**[]PublicMigrationMapping**](PublicMigrationMapping.md) |  | 
**MissingLegacyListIds** | **[]string** | A list of legacy list ids that were passed in but not found. It will be empty if no id&#39;s are missing | 

## Methods

### NewPublicBatchMigrationMapping

`func NewPublicBatchMigrationMapping(legacyListIdsToIdsMapping []PublicMigrationMapping, missingLegacyListIds []string, ) *PublicBatchMigrationMapping`

NewPublicBatchMigrationMapping instantiates a new PublicBatchMigrationMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicBatchMigrationMappingWithDefaults

`func NewPublicBatchMigrationMappingWithDefaults() *PublicBatchMigrationMapping`

NewPublicBatchMigrationMappingWithDefaults instantiates a new PublicBatchMigrationMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegacyListIdsToIdsMapping

`func (o *PublicBatchMigrationMapping) GetLegacyListIdsToIdsMapping() []PublicMigrationMapping`

GetLegacyListIdsToIdsMapping returns the LegacyListIdsToIdsMapping field if non-nil, zero value otherwise.

### GetLegacyListIdsToIdsMappingOk

`func (o *PublicBatchMigrationMapping) GetLegacyListIdsToIdsMappingOk() (*[]PublicMigrationMapping, bool)`

GetLegacyListIdsToIdsMappingOk returns a tuple with the LegacyListIdsToIdsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyListIdsToIdsMapping

`func (o *PublicBatchMigrationMapping) SetLegacyListIdsToIdsMapping(v []PublicMigrationMapping)`

SetLegacyListIdsToIdsMapping sets LegacyListIdsToIdsMapping field to given value.


### GetMissingLegacyListIds

`func (o *PublicBatchMigrationMapping) GetMissingLegacyListIds() []string`

GetMissingLegacyListIds returns the MissingLegacyListIds field if non-nil, zero value otherwise.

### GetMissingLegacyListIdsOk

`func (o *PublicBatchMigrationMapping) GetMissingLegacyListIdsOk() (*[]string, bool)`

GetMissingLegacyListIdsOk returns a tuple with the MissingLegacyListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingLegacyListIds

`func (o *PublicBatchMigrationMapping) SetMissingLegacyListIds(v []string)`

SetMissingLegacyListIds sets MissingLegacyListIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


