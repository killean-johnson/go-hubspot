# Go API client for actions_v4

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
import actions_v4 "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `actions_v4.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), actions_v4.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `actions_v4.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), actions_v4.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `actions_v4.ContextOperationServerIndices` and `actions_v4.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), actions_v4.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), actions_v4.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CallbacksAPI* | [**PostAutomationV4ActionsCallbacksCallbackIdCompleteComplete**](docs/CallbacksAPI.md#postautomationv4actionscallbackscallbackidcompletecomplete) | **Post** /automation/v4/actions/callbacks/{callbackId}/complete | Completes a callback
*CallbacksAPI* | [**PostAutomationV4ActionsCallbacksCompleteCompleteBatch**](docs/CallbacksAPI.md#postautomationv4actionscallbackscompletecompletebatch) | **Post** /automation/v4/actions/callbacks/complete | Complete a batch of callbacks
*DefinitionsAPI* | [**DeleteAutomationV4ActionsAppIdDefinitionIdArchive**](docs/DefinitionsAPI.md#deleteautomationv4actionsappiddefinitionidarchive) | **Delete** /automation/v4/actions/{appId}/{definitionId} | Delete an action definition
*DefinitionsAPI* | [**GetAutomationV4ActionsAppIdDefinitionIdGetById**](docs/DefinitionsAPI.md#getautomationv4actionsappiddefinitionidgetbyid) | **Get** /automation/v4/actions/{appId}/{definitionId} | Retrieve a custom action definition
*DefinitionsAPI* | [**GetAutomationV4ActionsAppIdGetPage**](docs/DefinitionsAPI.md#getautomationv4actionsappidgetpage) | **Get** /automation/v4/actions/{appId} | Retrieve custom action definitions
*DefinitionsAPI* | [**PatchAutomationV4ActionsAppIdDefinitionIdUpdate**](docs/DefinitionsAPI.md#patchautomationv4actionsappiddefinitionidupdate) | **Patch** /automation/v4/actions/{appId}/{definitionId} | Update an existing action definition
*DefinitionsAPI* | [**PostAutomationV4ActionsAppIdCreate**](docs/DefinitionsAPI.md#postautomationv4actionsappidcreate) | **Post** /automation/v4/actions/{appId} | Create a new custom action definition
*FunctionsAPI* | [**DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeArchiveByFunctionType**](docs/FunctionsAPI.md#deleteautomationv4actionsappiddefinitionidfunctionsfunctiontypearchivebyfunctiontype) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Delete a function for a definition
*FunctionsAPI* | [**DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdArchive**](docs/FunctionsAPI.md#deleteautomationv4actionsappiddefinitionidfunctionsfunctiontypefunctionidarchive) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Archive a function for a definition
*FunctionsAPI* | [**GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdGetById**](docs/FunctionsAPI.md#getautomationv4actionsappiddefinitionidfunctionsfunctiontypefunctionidgetbyid) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Retrieve a function from a given definition
*FunctionsAPI* | [**GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeGetByFunctionType**](docs/FunctionsAPI.md#getautomationv4actionsappiddefinitionidfunctionsfunctiontypegetbyfunctiontype) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Retrieve functions by a type for a given definition
*FunctionsAPI* | [**GetAutomationV4ActionsAppIdDefinitionIdFunctionsGetPage**](docs/FunctionsAPI.md#getautomationv4actionsappiddefinitionidfunctionsgetpage) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions | Retrieve functions for a given definition
*FunctionsAPI* | [**PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeCreateOrReplaceByFunctionType**](docs/FunctionsAPI.md#putautomationv4actionsappiddefinitionidfunctionsfunctiontypecreateorreplacebyfunctiontype) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Insert a function for a definition
*FunctionsAPI* | [**PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdCreateOrReplace**](docs/FunctionsAPI.md#putautomationv4actionsappiddefinitionidfunctionsfunctiontypefunctionidcreateorreplace) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Update a function for a definition
*RevisionsAPI* | [**GetAutomationV4ActionsAppIdDefinitionIdRevisionsGetPage**](docs/RevisionsAPI.md#getautomationv4actionsappiddefinitionidrevisionsgetpage) | **Get** /automation/v4/actions/{appId}/{definitionId}/revisions | Retrieve revisions for a given definition
*RevisionsAPI* | [**GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdGetById**](docs/RevisionsAPI.md#getautomationv4actionsappiddefinitionidrevisionsrevisionidgetbyid) | **Get** /automation/v4/actions/{appId}/{definitionId}/revisions/{revisionId} | Retrieve a specific revision of a definition


## Documentation For Models

 - [BatchInputCallbackCompletionBatchRequest](docs/BatchInputCallbackCompletionBatchRequest.md)
 - [CallbackCompletionBatchRequest](docs/CallbackCompletionBatchRequest.md)
 - [CallbackCompletionRequest](docs/CallbackCompletionRequest.md)
 - [CollectionResponsePublicActionDefinitionForwardPaging](docs/CollectionResponsePublicActionDefinitionForwardPaging.md)
 - [CollectionResponsePublicActionFunctionIdentifierNoPaging](docs/CollectionResponsePublicActionFunctionIdentifierNoPaging.md)
 - [CollectionResponsePublicActionRevisionForwardPaging](docs/CollectionResponsePublicActionRevisionForwardPaging.md)
 - [Error](docs/Error.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [FieldTypeDefinition](docs/FieldTypeDefinition.md)
 - [ForwardPaging](docs/ForwardPaging.md)
 - [InputFieldDefinition](docs/InputFieldDefinition.md)
 - [NextPage](docs/NextPage.md)
 - [Option](docs/Option.md)
 - [OutputFieldDefinition](docs/OutputFieldDefinition.md)
 - [PublicActionDefinition](docs/PublicActionDefinition.md)
 - [PublicActionDefinitionEgg](docs/PublicActionDefinitionEgg.md)
 - [PublicActionDefinitionInputFieldDependenciesInner](docs/PublicActionDefinitionInputFieldDependenciesInner.md)
 - [PublicActionDefinitionPatch](docs/PublicActionDefinitionPatch.md)
 - [PublicActionFunction](docs/PublicActionFunction.md)
 - [PublicActionFunctionIdentifier](docs/PublicActionFunctionIdentifier.md)
 - [PublicActionLabels](docs/PublicActionLabels.md)
 - [PublicActionRevision](docs/PublicActionRevision.md)
 - [PublicConditionalSingleFieldDependency](docs/PublicConditionalSingleFieldDependency.md)
 - [PublicExecutionTranslationRule](docs/PublicExecutionTranslationRule.md)
 - [PublicObjectRequestOptions](docs/PublicObjectRequestOptions.md)
 - [PublicSingleFieldDependency](docs/PublicSingleFieldDependency.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### developer_hapikey

- **Type**: API key
- **API key parameter name**: hapikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: developer_hapikey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		actions_v4.ContextAPIKeys,
		map[string]actions_v4.APIKey{
			"developer_hapikey": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **automation**: Read from and write to my Workflows

Example

```go
auth := context.WithValue(context.Background(), actions_v4.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, actions_v4.ContextOAuth2, tokenSource)
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
		actions_v4.ContextAPIKeys,
		map[string]actions_v4.APIKey{
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



