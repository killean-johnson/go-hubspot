# Go API client for automation_v4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v4
- Package version: 1.0.0
- Generator version: 7.12.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import automation_v4 "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `automation_v4.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), automation_v4.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `automation_v4.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), automation_v4.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `automation_v4.ContextOperationServerIndices` and `automation_v4.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), automation_v4.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), automation_v4.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EmailCampaignsAPI* | [**GetAutomationV4FlowsEmailCampaigns**](docs/EmailCampaignsAPI.md#getautomationv4flowsemailcampaigns) | **Get** /automation/v4/flows/email-campaigns | Retrieve workflow emails
*WorkflowIDMappingsAPI* | [**PostAutomationV4WorkflowIdMappingsBatchRead**](docs/WorkflowIDMappingsAPI.md#postautomationv4workflowidmappingsbatchread) | **Post** /automation/v4/workflow-id-mappings/batch/read | Retrieve migrated workflow mappings
*WorkflowsAPI* | [**DeleteAutomationV4FlowsFlowId**](docs/WorkflowsAPI.md#deleteautomationv4flowsflowid) | **Delete** /automation/v4/flows/{flowId} | Delete a workflow
*WorkflowsAPI* | [**GetAutomationV4Flows**](docs/WorkflowsAPI.md#getautomationv4flows) | **Get** /automation/v4/flows | Retrieve workflows
*WorkflowsAPI* | [**GetAutomationV4FlowsFlowId**](docs/WorkflowsAPI.md#getautomationv4flowsflowid) | **Get** /automation/v4/flows/{flowId} | Retrieve workflow details
*WorkflowsAPI* | [**PostAutomationV4Flows**](docs/WorkflowsAPI.md#postautomationv4flows) | **Post** /automation/v4/flows | Create a new workflow
*WorkflowsAPI* | [**PostAutomationV4FlowsBatchRead**](docs/WorkflowsAPI.md#postautomationv4flowsbatchread) | **Post** /automation/v4/flows/batch/read | Retrieve a batch of workflows
*WorkflowsAPI* | [**PutAutomationV4FlowsFlowId**](docs/WorkflowsAPI.md#putautomationv4flowsflowid) | **Put** /automation/v4/flows/{flowId} | Update a workflow


## Documentation For Models

 - [ApiABTestBranchAction](docs/ApiABTestBranchAction.md)
 - [ApiActionDataValue](docs/ApiActionDataValue.md)
 - [ApiAppendObjectPropertyValue](docs/ApiAppendObjectPropertyValue.md)
 - [ApiAssociationDataSource](docs/ApiAssociationDataSource.md)
 - [ApiAssociationTimestampDataSource](docs/ApiAssociationTimestampDataSource.md)
 - [ApiAuthKeyWebhookAuthSettings](docs/ApiAuthKeyWebhookAuthSettings.md)
 - [ApiBlockedDate](docs/ApiBlockedDate.md)
 - [ApiConnection](docs/ApiConnection.md)
 - [ApiContactFlow](docs/ApiContactFlow.md)
 - [ApiContactFlowAllOfActions](docs/ApiContactFlowAllOfActions.md)
 - [ApiContactFlowAllOfDataSources](docs/ApiContactFlowAllOfDataSources.md)
 - [ApiContactFlowAllOfEnrollmentCriteria](docs/ApiContactFlowAllOfEnrollmentCriteria.md)
 - [ApiContactFlowAllOfEnrollmentSchedule](docs/ApiContactFlowAllOfEnrollmentSchedule.md)
 - [ApiContactFlowAllOfEventAnchor](docs/ApiContactFlowAllOfEventAnchor.md)
 - [ApiContactFlowAllOfGoalFilterBranch](docs/ApiContactFlowAllOfGoalFilterBranch.md)
 - [ApiContactFlowCreateRequest](docs/ApiContactFlowCreateRequest.md)
 - [ApiContactFlowCreateRequestAllOfActions](docs/ApiContactFlowCreateRequestAllOfActions.md)
 - [ApiContactFlowCreateRequestAllOfDataSources](docs/ApiContactFlowCreateRequestAllOfDataSources.md)
 - [ApiContactFlowCreateRequestAllOfEnrollmentCriteria](docs/ApiContactFlowCreateRequestAllOfEnrollmentCriteria.md)
 - [ApiContactFlowCreateRequestAllOfEnrollmentSchedule](docs/ApiContactFlowCreateRequestAllOfEnrollmentSchedule.md)
 - [ApiContactFlowCreateRequestAllOfEventAnchor](docs/ApiContactFlowCreateRequestAllOfEventAnchor.md)
 - [ApiContactFlowCreateRequestAllOfGoalFilterBranch](docs/ApiContactFlowCreateRequestAllOfGoalFilterBranch.md)
 - [ApiContactFlowPutRequest](docs/ApiContactFlowPutRequest.md)
 - [ApiContactFlowPutRequestAllOfActions](docs/ApiContactFlowPutRequestAllOfActions.md)
 - [ApiContactFlowPutRequestAllOfEnrollmentCriteria](docs/ApiContactFlowPutRequestAllOfEnrollmentCriteria.md)
 - [ApiContactFlowPutRequestAllOfEnrollmentSchedule](docs/ApiContactFlowPutRequestAllOfEnrollmentSchedule.md)
 - [ApiContactFlowPutRequestAllOfEventAnchor](docs/ApiContactFlowPutRequestAllOfEventAnchor.md)
 - [ApiContactFlowPutRequestAllOfGoalFilterBranch](docs/ApiContactFlowPutRequestAllOfGoalFilterBranch.md)
 - [ApiContactPropertyAnchor](docs/ApiContactPropertyAnchor.md)
 - [ApiCustomCodeAction](docs/ApiCustomCodeAction.md)
 - [ApiCustomCodeActionOutputFieldsInner](docs/ApiCustomCodeActionOutputFieldsInner.md)
 - [ApiDailyEnrollmentSchedule](docs/ApiDailyEnrollmentSchedule.md)
 - [ApiEnrollmentEventPropertyValue](docs/ApiEnrollmentEventPropertyValue.md)
 - [ApiEnumerationOutputField](docs/ApiEnumerationOutputField.md)
 - [ApiEventBasedEnrollmentCriteria](docs/ApiEventBasedEnrollmentCriteria.md)
 - [ApiEventBasedEnrollmentCriteriaRefinementCriteria](docs/ApiEventBasedEnrollmentCriteriaRefinementCriteria.md)
 - [ApiFetchedObjectPropertyValue](docs/ApiFetchedObjectPropertyValue.md)
 - [ApiFlow](docs/ApiFlow.md)
 - [ApiFlowBatchFetchFlowIdCoordinate](docs/ApiFlowBatchFetchFlowIdCoordinate.md)
 - [ApiFlowBatchFetchMigrationFlowIdCoordinate](docs/ApiFlowBatchFetchMigrationFlowIdCoordinate.md)
 - [ApiFlowBatchFetchMigrationWorkflowIdCoordinate](docs/ApiFlowBatchFetchMigrationWorkflowIdCoordinate.md)
 - [ApiFlowBatchInput](docs/ApiFlowBatchInput.md)
 - [ApiFlowBatchInputInputsInner](docs/ApiFlowBatchInputInputsInner.md)
 - [ApiFlowBatchMigrationInput](docs/ApiFlowBatchMigrationInput.md)
 - [ApiFlowBatchMigrationInputInputsInner](docs/ApiFlowBatchMigrationInputInputsInner.md)
 - [ApiFlowCreateRequest](docs/ApiFlowCreateRequest.md)
 - [ApiFlowEmailCampaign](docs/ApiFlowEmailCampaign.md)
 - [ApiFlowListing](docs/ApiFlowListing.md)
 - [ApiFlowPutRequest](docs/ApiFlowPutRequest.md)
 - [ApiIncrementValue](docs/ApiIncrementValue.md)
 - [ApiInputVariable](docs/ApiInputVariable.md)
 - [ApiInputVariableValue](docs/ApiInputVariableValue.md)
 - [ApiListBasedEnrollmentCriteria](docs/ApiListBasedEnrollmentCriteria.md)
 - [ApiListBasedEnrollmentCriteriaListFilterBranch](docs/ApiListBasedEnrollmentCriteriaListFilterBranch.md)
 - [ApiListBranch](docs/ApiListBranch.md)
 - [ApiListBranchAction](docs/ApiListBranchAction.md)
 - [ApiListBranchFilterBranch](docs/ApiListBranchFilterBranch.md)
 - [ApiManualEnrollmentCriteria](docs/ApiManualEnrollmentCriteria.md)
 - [ApiMonthlyRelativeDaysEnrollmentSchedule](docs/ApiMonthlyRelativeDaysEnrollmentSchedule.md)
 - [ApiMonthlySpecificDaysEnrollmentSchedule](docs/ApiMonthlySpecificDaysEnrollmentSchedule.md)
 - [ApiObjectPropertyValue](docs/ApiObjectPropertyValue.md)
 - [ApiPlatformFlow](docs/ApiPlatformFlow.md)
 - [ApiPlatformFlowAllOfActions](docs/ApiPlatformFlowAllOfActions.md)
 - [ApiPlatformFlowAllOfDataSources](docs/ApiPlatformFlowAllOfDataSources.md)
 - [ApiPlatformFlowAllOfEnrollmentCriteria](docs/ApiPlatformFlowAllOfEnrollmentCriteria.md)
 - [ApiPlatformFlowAllOfEnrollmentSchedule](docs/ApiPlatformFlowAllOfEnrollmentSchedule.md)
 - [ApiPlatformFlowAllOfSuppressionFilterBranch](docs/ApiPlatformFlowAllOfSuppressionFilterBranch.md)
 - [ApiPlatformFlowCreateRequest](docs/ApiPlatformFlowCreateRequest.md)
 - [ApiPlatformFlowCreateRequestAllOfActions](docs/ApiPlatformFlowCreateRequestAllOfActions.md)
 - [ApiPlatformFlowCreateRequestAllOfDataSources](docs/ApiPlatformFlowCreateRequestAllOfDataSources.md)
 - [ApiPlatformFlowCreateRequestAllOfEnrollmentCriteria](docs/ApiPlatformFlowCreateRequestAllOfEnrollmentCriteria.md)
 - [ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule](docs/ApiPlatformFlowCreateRequestAllOfEnrollmentSchedule.md)
 - [ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch](docs/ApiPlatformFlowCreateRequestAllOfSuppressionFilterBranch.md)
 - [ApiPlatformFlowPutRequest](docs/ApiPlatformFlowPutRequest.md)
 - [ApiPlatformFlowPutRequestAllOfActions](docs/ApiPlatformFlowPutRequestAllOfActions.md)
 - [ApiPlatformFlowPutRequestAllOfEnrollmentCriteria](docs/ApiPlatformFlowPutRequestAllOfEnrollmentCriteria.md)
 - [ApiPlatformFlowPutRequestAllOfEnrollmentSchedule](docs/ApiPlatformFlowPutRequestAllOfEnrollmentSchedule.md)
 - [ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch](docs/ApiPlatformFlowPutRequestAllOfSuppressionFilterBranch.md)
 - [ApiPropertyBasedEnrollmentSchedule](docs/ApiPropertyBasedEnrollmentSchedule.md)
 - [ApiRelativeDateTimeValue](docs/ApiRelativeDateTimeValue.md)
 - [ApiSignatureWebhookAuthSettings](docs/ApiSignatureWebhookAuthSettings.md)
 - [ApiSingleConnectionAction](docs/ApiSingleConnectionAction.md)
 - [ApiSort](docs/ApiSort.md)
 - [ApiStaticAppendValue](docs/ApiStaticAppendValue.md)
 - [ApiStaticBranch](docs/ApiStaticBranch.md)
 - [ApiStaticBranchAction](docs/ApiStaticBranchAction.md)
 - [ApiStaticBranchActionInputValue](docs/ApiStaticBranchActionInputValue.md)
 - [ApiStaticDateAnchor](docs/ApiStaticDateAnchor.md)
 - [ApiStaticTimeZoneStrategy](docs/ApiStaticTimeZoneStrategy.md)
 - [ApiStaticValue](docs/ApiStaticValue.md)
 - [ApiTimeDelay](docs/ApiTimeDelay.md)
 - [ApiTimeDelayTimeZoneStrategy](docs/ApiTimeDelayTimeZoneStrategy.md)
 - [ApiTimeOfDay](docs/ApiTimeOfDay.md)
 - [ApiTimeWindow](docs/ApiTimeWindow.md)
 - [ApiTimestampValue](docs/ApiTimestampValue.md)
 - [ApiUnEnrollmentSetting](docs/ApiUnEnrollmentSetting.md)
 - [ApiWebhookAction](docs/ApiWebhookAction.md)
 - [ApiWebhookActionAuthSettings](docs/ApiWebhookActionAuthSettings.md)
 - [ApiWeeklyEnrollmentSchedule](docs/ApiWeeklyEnrollmentSchedule.md)
 - [ApiYearlyEnrollmentSchedule](docs/ApiYearlyEnrollmentSchedule.md)
 - [BatchResponseApiFlow](docs/BatchResponseApiFlow.md)
 - [BatchResponseApiFlowWithErrors](docs/BatchResponseApiFlowWithErrors.md)
 - [BatchResponseFlowIdWorkflowIdMappingResponse](docs/BatchResponseFlowIdWorkflowIdMappingResponse.md)
 - [BatchResponseFlowIdWorkflowIdMappingResponseWithErrors](docs/BatchResponseFlowIdWorkflowIdMappingResponseWithErrors.md)
 - [CollectionResponseApiFlowEmailCampaign](docs/CollectionResponseApiFlowEmailCampaign.md)
 - [CollectionResponseApiFlowListingForwardPaging](docs/CollectionResponseApiFlowListingForwardPaging.md)
 - [Error](docs/Error.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [FlowIdWorkflowIdMappingResponse](docs/FlowIdWorkflowIdMappingResponse.md)
 - [ForwardPaging](docs/ForwardPaging.md)
 - [NextPage](docs/NextPage.md)
 - [Paging](docs/Paging.md)
 - [PreviousPage](docs/PreviousPage.md)
 - [PublicAbsoluteComparativeTimestampRefineBy](docs/PublicAbsoluteComparativeTimestampRefineBy.md)
 - [PublicAbsoluteRangedTimestampRefineBy](docs/PublicAbsoluteRangedTimestampRefineBy.md)
 - [PublicAdsSearchFilter](docs/PublicAdsSearchFilter.md)
 - [PublicAdsTimeFilter](docs/PublicAdsTimeFilter.md)
 - [PublicAllHistoryRefineBy](docs/PublicAllHistoryRefineBy.md)
 - [PublicAllPropertyTypesOperation](docs/PublicAllPropertyTypesOperation.md)
 - [PublicAndFilterBranch](docs/PublicAndFilterBranch.md)
 - [PublicAssociationFilterBranch](docs/PublicAssociationFilterBranch.md)
 - [PublicAssociationInListFilter](docs/PublicAssociationInListFilter.md)
 - [PublicBoolPropertyOperation](docs/PublicBoolPropertyOperation.md)
 - [PublicCalendarDatePropertyOperation](docs/PublicCalendarDatePropertyOperation.md)
 - [PublicCampaignInfluencedFilter](docs/PublicCampaignInfluencedFilter.md)
 - [PublicCommunicationSubscriptionFilter](docs/PublicCommunicationSubscriptionFilter.md)
 - [PublicComparativeDatePropertyOperation](docs/PublicComparativeDatePropertyOperation.md)
 - [PublicComparativePropertyUpdatedOperation](docs/PublicComparativePropertyUpdatedOperation.md)
 - [PublicConstantFilter](docs/PublicConstantFilter.md)
 - [PublicCtaAnalyticsFilter](docs/PublicCtaAnalyticsFilter.md)
 - [PublicDatePoint](docs/PublicDatePoint.md)
 - [PublicDatePropertyOperation](docs/PublicDatePropertyOperation.md)
 - [PublicDateTimePropertyOperation](docs/PublicDateTimePropertyOperation.md)
 - [PublicEmailEventFilter](docs/PublicEmailEventFilter.md)
 - [PublicEmailSubscriptionFilter](docs/PublicEmailSubscriptionFilter.md)
 - [PublicEnumerationPropertyOperation](docs/PublicEnumerationPropertyOperation.md)
 - [PublicEventAnalyticsFilter](docs/PublicEventAnalyticsFilter.md)
 - [PublicEventFilterMetadata](docs/PublicEventFilterMetadata.md)
 - [PublicFiscalQuarterReference](docs/PublicFiscalQuarterReference.md)
 - [PublicFiscalYearReference](docs/PublicFiscalYearReference.md)
 - [PublicFormSubmissionFilter](docs/PublicFormSubmissionFilter.md)
 - [PublicFormSubmissionFilterCoalescingRefineBy](docs/PublicFormSubmissionFilterCoalescingRefineBy.md)
 - [PublicFormSubmissionOnPageFilter](docs/PublicFormSubmissionOnPageFilter.md)
 - [PublicInListFilter](docs/PublicInListFilter.md)
 - [PublicInListFilterMetadata](docs/PublicInListFilterMetadata.md)
 - [PublicIndexOffset](docs/PublicIndexOffset.md)
 - [PublicIndexedTimePoint](docs/PublicIndexedTimePoint.md)
 - [PublicIndexedTimePointIndexReference](docs/PublicIndexedTimePointIndexReference.md)
 - [PublicIntegrationEventFilter](docs/PublicIntegrationEventFilter.md)
 - [PublicMonthReference](docs/PublicMonthReference.md)
 - [PublicMultiStringPropertyOperation](docs/PublicMultiStringPropertyOperation.md)
 - [PublicNotAllFilterBranch](docs/PublicNotAllFilterBranch.md)
 - [PublicNotAnyFilterBranch](docs/PublicNotAnyFilterBranch.md)
 - [PublicNowReference](docs/PublicNowReference.md)
 - [PublicNumAssociationsFilter](docs/PublicNumAssociationsFilter.md)
 - [PublicNumOccurrencesRefineBy](docs/PublicNumOccurrencesRefineBy.md)
 - [PublicNumberPropertyOperation](docs/PublicNumberPropertyOperation.md)
 - [PublicOrFilterBranch](docs/PublicOrFilterBranch.md)
 - [PublicPageViewAnalyticsFilter](docs/PublicPageViewAnalyticsFilter.md)
 - [PublicPrivacyAnalyticsFilter](docs/PublicPrivacyAnalyticsFilter.md)
 - [PublicPropertyAssociationFilterBranch](docs/PublicPropertyAssociationFilterBranch.md)
 - [PublicPropertyAssociationFilterBranchFiltersInner](docs/PublicPropertyAssociationFilterBranchFiltersInner.md)
 - [PublicPropertyAssociationInListFilter](docs/PublicPropertyAssociationInListFilter.md)
 - [PublicPropertyFilter](docs/PublicPropertyFilter.md)
 - [PublicPropertyReferencedTime](docs/PublicPropertyReferencedTime.md)
 - [PublicQuarterReference](docs/PublicQuarterReference.md)
 - [PublicRangedDatePropertyOperation](docs/PublicRangedDatePropertyOperation.md)
 - [PublicRangedNumberPropertyOperation](docs/PublicRangedNumberPropertyOperation.md)
 - [PublicRangedTimeOperation](docs/PublicRangedTimeOperation.md)
 - [PublicRelativeComparativeTimestampRefineBy](docs/PublicRelativeComparativeTimestampRefineBy.md)
 - [PublicRelativeRangedTimestampRefineBy](docs/PublicRelativeRangedTimestampRefineBy.md)
 - [PublicRestrictedFilterBranch](docs/PublicRestrictedFilterBranch.md)
 - [PublicRollingDateRangePropertyOperation](docs/PublicRollingDateRangePropertyOperation.md)
 - [PublicRollingPropertyUpdatedOperation](docs/PublicRollingPropertyUpdatedOperation.md)
 - [PublicSetOccurrencesRefineBy](docs/PublicSetOccurrencesRefineBy.md)
 - [PublicStringPropertyOperation](docs/PublicStringPropertyOperation.md)
 - [PublicSurveyMonkeyFilter](docs/PublicSurveyMonkeyFilter.md)
 - [PublicSurveyMonkeyValueFilter](docs/PublicSurveyMonkeyValueFilter.md)
 - [PublicSurveyMonkeyValueFilterValueComparison](docs/PublicSurveyMonkeyValueFilterValueComparison.md)
 - [PublicTimeOffset](docs/PublicTimeOffset.md)
 - [PublicTimePointOperation](docs/PublicTimePointOperation.md)
 - [PublicTimePointOperationTimePoint](docs/PublicTimePointOperationTimePoint.md)
 - [PublicTodayReference](docs/PublicTodayReference.md)
 - [PublicUnifiedEventsFilter](docs/PublicUnifiedEventsFilter.md)
 - [PublicUnifiedEventsFilterBranch](docs/PublicUnifiedEventsFilterBranch.md)
 - [PublicWebinarFilter](docs/PublicWebinarFilter.md)
 - [PublicWeekReference](docs/PublicWeekReference.md)
 - [PublicYearReference](docs/PublicYearReference.md)
 - [StandardError](docs/StandardError.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **automation**: Read from and write to my Workflows

Example

```go
auth := context.WithValue(context.Background(), automation_v4.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, automation_v4.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### private_apps

- **Type**: API key
- **API key parameter name**: private-app
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: private_apps and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		automation_v4.ContextAPIKeys,
		map[string]automation_v4.APIKey{
			"private_apps": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



