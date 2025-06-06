# Go API client for associations

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
import associations "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `associations.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), associations.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `associations.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), associations.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `associations.ContextOperationServerIndices` and `associations.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), associations.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), associations.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BasicAPI* | [**DeleteCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdArchive**](docs/BasicAPI.md#deletecrmv4objectsobjecttypeobjectidassociationstoobjecttypetoobjectidarchive) | **Delete** /crm/v4/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Delete
*BasicAPI* | [**GetCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeGetPage**](docs/BasicAPI.md#getcrmv4objectsobjecttypeobjectidassociationstoobjecttypegetpage) | **Get** /crm/v4/objects/{objectType}/{objectId}/associations/{toObjectType} | List
*BasicAPI* | [**PutCrmV4ObjectsFromObjectTypeFromObjectIdAssociationsDefaultToObjectTypeToObjectIdCreateDefault**](docs/BasicAPI.md#putcrmv4objectsfromobjecttypefromobjectidassociationsdefaulttoobjecttypetoobjectidcreatedefault) | **Put** /crm/v4/objects/{fromObjectType}/{fromObjectId}/associations/default/{toObjectType}/{toObjectId} | Create Default
*BasicAPI* | [**PutCrmV4ObjectsObjectTypeObjectIdAssociationsToObjectTypeToObjectIdCreate**](docs/BasicAPI.md#putcrmv4objectsobjecttypeobjectidassociationstoobjecttypetoobjectidcreate) | **Put** /crm/v4/objects/{objectType}/{objectId}/associations/{toObjectType}/{toObjectId} | Create
*BatchAPI* | [**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive**](docs/BatchAPI.md#postcrmv4associationsfromobjecttypetoobjecttypebatcharchivearchive) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/archive | Delete
*BatchAPI* | [**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault**](docs/BatchAPI.md#postcrmv4associationsfromobjecttypetoobjecttypebatchassociatedefaultcreatedefault) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/associate/default |  Create Default Associations
*BatchAPI* | [**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate**](docs/BatchAPI.md#postcrmv4associationsfromobjecttypetoobjecttypebatchcreatecreate) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/create | Create
*BatchAPI* | [**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels**](docs/BatchAPI.md#postcrmv4associationsfromobjecttypetoobjecttypebatchlabelsarchivearchivelabels) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/labels/archive | Delete Specific Labels
*BatchAPI* | [**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage**](docs/BatchAPI.md#postcrmv4associationsfromobjecttypetoobjecttypebatchreadgetpage) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/read | Read
*ReportAPI* | [**PostCrmV4AssociationsUsageHighUsageReportUserIdRequest**](docs/ReportAPI.md#postcrmv4associationsusagehighusagereportuseridrequest) | **Post** /crm/v4/associations/usage/high-usage-report/{userId} | Report


## Documentation For Models

 - [AssociationSpec](docs/AssociationSpec.md)
 - [AssociationSpecWithLabel](docs/AssociationSpecWithLabel.md)
 - [BatchInputPublicAssociationMultiArchive](docs/BatchInputPublicAssociationMultiArchive.md)
 - [BatchInputPublicAssociationMultiPost](docs/BatchInputPublicAssociationMultiPost.md)
 - [BatchInputPublicDefaultAssociationMultiPost](docs/BatchInputPublicDefaultAssociationMultiPost.md)
 - [BatchInputPublicFetchAssociationsBatchRequest](docs/BatchInputPublicFetchAssociationsBatchRequest.md)
 - [BatchResponseLabelsBetweenObjectPair](docs/BatchResponseLabelsBetweenObjectPair.md)
 - [BatchResponseLabelsBetweenObjectPairWithErrors](docs/BatchResponseLabelsBetweenObjectPairWithErrors.md)
 - [BatchResponsePublicAssociationMultiWithLabel](docs/BatchResponsePublicAssociationMultiWithLabel.md)
 - [BatchResponsePublicAssociationMultiWithLabelWithErrors](docs/BatchResponsePublicAssociationMultiWithLabelWithErrors.md)
 - [BatchResponsePublicDefaultAssociation](docs/BatchResponsePublicDefaultAssociation.md)
 - [CollectionResponseMultiAssociatedObjectWithLabelForwardPaging](docs/CollectionResponseMultiAssociatedObjectWithLabelForwardPaging.md)
 - [DateTime](docs/DateTime.md)
 - [Error](docs/Error.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [ForwardPaging](docs/ForwardPaging.md)
 - [LabelsBetweenObjectPair](docs/LabelsBetweenObjectPair.md)
 - [MultiAssociatedObjectWithLabel](docs/MultiAssociatedObjectWithLabel.md)
 - [NextPage](docs/NextPage.md)
 - [Paging](docs/Paging.md)
 - [PreviousPage](docs/PreviousPage.md)
 - [PublicAssociationMultiArchive](docs/PublicAssociationMultiArchive.md)
 - [PublicAssociationMultiPost](docs/PublicAssociationMultiPost.md)
 - [PublicAssociationMultiWithLabel](docs/PublicAssociationMultiWithLabel.md)
 - [PublicDefaultAssociation](docs/PublicDefaultAssociation.md)
 - [PublicDefaultAssociationMultiPost](docs/PublicDefaultAssociationMultiPost.md)
 - [PublicFetchAssociationsBatchRequest](docs/PublicFetchAssociationsBatchRequest.md)
 - [PublicObjectId](docs/PublicObjectId.md)
 - [ReportCreationResponse](docs/ReportCreationResponse.md)
 - [StandardError](docs/StandardError.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **crm.objects.users.write**: Write User CRM objects
 - **crm.objects.services.write**: Write services
 - **crm.objects.contacts.highly_sensitive.write.v2**: Change contact records (highly-sensitive)
 - **crm.objects.partner-clients.write**: Modify Partner Client CRM objects
 - **crm.objects.custom.write**: Change custom object records
 - **crm.objects.quotes.read**: Quotes
 - **crm.objects.contacts.sensitive.read.v2**:  
 - **crm.objects.custom.highly_sensitive.read.v2**: View custom object records (highly-sensitive)
 - **crm.objects.goals.write**: Write goals
 - **crm.objects.custom.sensitive.read.v2**: View custom object records (sensitive)
 - **tickets.sensitive.v2**: Tickets (sensitive)
 - **crm.objects.appointments.sensitive.read.v2**: Read appointments (sensitive)
 - **crm.objects.contacts.sensitive.write.v2**:  
 - **crm.objects.companies.highly_sensitive.write.v2**: Change company records (highly-sensitive)
 - **crm.objects.partner-clients.read**: View Partner Client CRM objects
 - **crm.objects.deals.write**:  
 - **crm.objects.leads.write**: Modify lead objects
 - **crm.objects.companies.read**:  
 - **crm.objects.listings.read**: Read listings
 - **crm.objects.deals.highly_sensitive.read.v2**: View deal records (highly-sensitive)
 - **crm.objects.contacts.write**:  
 - **crm.objects.carts.write**: Write cart
 - **crm.objects.companies.sensitive.read.v2**:  
 - **crm.objects.invoices.read**: Read invoices objects
 - **crm.objects.subscriptions.read**: Read Commerce Subscriptions
 - **crm.objects.invoices.write**: Write invoices
 - **crm.objects.deals.highly_sensitive.write.v2**: Change deal records (highly-sensitive)
 - **crm.objects.goals.read**: Read goals
 - **crm.objects.custom.highly_sensitive.write.v2**: Change custom object records (highly-sensitive)
 - **crm.objects.commercepayments.read**: Read the COMMERCE_PAYMENT object.
 - **crm.objects.deals.sensitive.write.v2**:  
 - **crm.objects.appointments.write**: Write appointments
 - **media_bridge.read**: Read media and media events
 - **crm.objects.line_items.write**: Line Items
 - **crm.objects.orders.read**: Read Orders
 - **crm.objects.subscriptions.write**: Write to Commerce Subscriptions
 - **crm.objects.contacts.read**:  
 - **crm.objects.courses.read**: Read courses
 - **crm.objects.quotes.write**: Quotes
 - **crm.objects.deals.sensitive.read.v2**:  
 - **crm.objects.line_items.read**: Line Items
 - **crm.objects.users.read**: Read User CRM objects
 - **tickets**: Read and write tickets
 - **crm.objects.companies.sensitive.write.v2**:  
 - **crm.objects.leads.read**: Read lead objects
 - **crm.objects.services.read**: Read services
 - **crm.objects.courses.write**: Write courses
 - **crm.objects.custom.read**: View custom object records
 - **crm.objects.listings.write**: Write listings
 - **crm.objects.companies.write**:  
 - **crm.objects.appointments.sensitive.write.v2**: Write appointments (sensitive)
 - **crm.objects.deals.read**:  
 - **crm.objects.companies.highly_sensitive.read.v2**: View company records (highly-sensitive)
 - **e-commerce**: e-commerce
 - **crm.objects.contacts.highly_sensitive.read.v2**: View contact records (highly-sensitive)
 - **crm.objects.appointments.read**: Read appointments
 - **crm.objects.orders.write**: Write orders
 - **crm.objects.custom.sensitive.write.v2**: Change custom object records (sensitive)
 - **crm.objects.carts.read**: Read carts
 - **tickets.highly_sensitive.v2**: Tickets (highly-sensitive)

Example

```go
auth := context.WithValue(context.Background(), associations.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, associations.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### private_apps_legacy

- **Type**: API key
- **API key parameter name**: private-app-legacy
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: private_apps_legacy and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		associations.ContextAPIKeys,
		map[string]associations.APIKey{
			"private_apps_legacy": {Key: "API_KEY_STRING"},
		},
	)
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
		associations.ContextAPIKeys,
		map[string]associations.APIKey{
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



