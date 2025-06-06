# Go API client for hubdb

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v3
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
import hubdb "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `hubdb.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), hubdb.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `hubdb.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), hubdb.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `hubdb.ContextOperationServerIndices` and `hubdb.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), hubdb.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), hubdb.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*RowsAPI* | [**DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftPurgeDraftTableRow**](docs/RowsAPI.md#deletecmsv3hubdbtablestableidornamerowsrowiddraftpurgedrafttablerow) | **Delete** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Permanently deletes a row
*RowsAPI* | [**GetCmsV3HubdbTablesTableIdOrNameRowsDraftReadDraftTableRows**](docs/RowsAPI.md#getcmsv3hubdbtablestableidornamerowsdraftreaddrafttablerows) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft | Get rows from draft table
*RowsAPI* | [**GetCmsV3HubdbTablesTableIdOrNameRowsGetTableRows**](docs/RowsAPI.md#getcmsv3hubdbtablestableidornamerowsgettablerows) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows | Get rows for a table
*RowsAPI* | [**GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftGetDraftTableRowById**](docs/RowsAPI.md#getcmsv3hubdbtablestableidornamerowsrowiddraftgetdrafttablerowbyid) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Get a row from the draft table
*RowsAPI* | [**GetCmsV3HubdbTablesTableIdOrNameRowsRowIdGetTableRow**](docs/RowsAPI.md#getcmsv3hubdbtablestableidornamerowsrowidgettablerow) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId} | Get a table row
*RowsAPI* | [**PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftUpdateDraftTableRow**](docs/RowsAPI.md#patchcmsv3hubdbtablestableidornamerowsrowiddraftupdatedrafttablerow) | **Patch** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Updates an existing row
*RowsAPI* | [**PostCmsV3HubdbTablesTableIdOrNameRowsCreateTableRow**](docs/RowsAPI.md#postcmsv3hubdbtablestableidornamerowscreatetablerow) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows | Add a new row to a table
*RowsAPI* | [**PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneCloneDraftTableRow**](docs/RowsAPI.md#postcmsv3hubdbtablestableidornamerowsrowiddraftcloneclonedrafttablerow) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft/clone | Clone a row
*RowsAPI* | [**PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftReplaceDraftTableRow**](docs/RowsAPI.md#putcmsv3hubdbtablestableidornamerowsrowiddraftreplacedrafttablerow) | **Put** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Replaces an existing row
*RowsBatchAPI* | [**PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows**](docs/RowsBatchAPI.md#postcmsv3hubdbtablestableidornamerowsbatchreadreadtablerows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/batch/read | Get a set of rows
*RowsBatchAPI* | [**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows**](docs/RowsBatchAPI.md#postcmsv3hubdbtablestableidornamerowsdraftbatchcloneclonedrafttablerows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/clone | Clone rows in batch
*RowsBatchAPI* | [**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows**](docs/RowsBatchAPI.md#postcmsv3hubdbtablestableidornamerowsdraftbatchcreatecreatedrafttablerows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/create | Create rows in batch
*RowsBatchAPI* | [**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRows**](docs/RowsBatchAPI.md#postcmsv3hubdbtablestableidornamerowsdraftbatchpurgepurgedrafttablerows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/purge | Permanently deletes rows
*RowsBatchAPI* | [**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows**](docs/RowsBatchAPI.md#postcmsv3hubdbtablestableidornamerowsdraftbatchreadreaddrafttablerows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/read | Get a set of rows from draft table
*RowsBatchAPI* | [**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows**](docs/RowsBatchAPI.md#postcmsv3hubdbtablestableidornamerowsdraftbatchreplacereplacedrafttablerows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/replace | Replace rows in batch in draft table
*RowsBatchAPI* | [**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows**](docs/RowsBatchAPI.md#postcmsv3hubdbtablestableidornamerowsdraftbatchupdateupdatedrafttablerows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/update | Update rows in batch in draft table
*TablesAPI* | [**DeleteCmsV3HubdbTablesTableIdOrNameArchiveTable**](docs/TablesAPI.md#deletecmsv3hubdbtablestableidornamearchivetable) | **Delete** /cms/v3/hubdb/tables/{tableIdOrName} | Archive a table
*TablesAPI* | [**GetCmsV3HubdbTablesDraftGetAllDraftTables**](docs/TablesAPI.md#getcmsv3hubdbtablesdraftgetalldrafttables) | **Get** /cms/v3/hubdb/tables/draft | Return all draft tables
*TablesAPI* | [**GetCmsV3HubdbTablesGetAllTables**](docs/TablesAPI.md#getcmsv3hubdbtablesgetalltables) | **Get** /cms/v3/hubdb/tables | Get all published tables
*TablesAPI* | [**GetCmsV3HubdbTablesTableIdOrNameDraftExportExportDraftTable**](docs/TablesAPI.md#getcmsv3hubdbtablestableidornamedraftexportexportdrafttable) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/draft/export | Export a draft table
*TablesAPI* | [**GetCmsV3HubdbTablesTableIdOrNameDraftGetDraftTableDetailsById**](docs/TablesAPI.md#getcmsv3hubdbtablestableidornamedraftgetdrafttabledetailsbyid) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/draft | Get details for a draft table
*TablesAPI* | [**GetCmsV3HubdbTablesTableIdOrNameExportExportTable**](docs/TablesAPI.md#getcmsv3hubdbtablestableidornameexportexporttable) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/export | Export a published version of a table
*TablesAPI* | [**GetCmsV3HubdbTablesTableIdOrNameGetTableDetails**](docs/TablesAPI.md#getcmsv3hubdbtablestableidornamegettabledetails) | **Get** /cms/v3/hubdb/tables/{tableIdOrName} | Get details of a published table
*TablesAPI* | [**PatchCmsV3HubdbTablesTableIdOrNameDraftUpdateDraftTable**](docs/TablesAPI.md#patchcmsv3hubdbtablestableidornamedraftupdatedrafttable) | **Patch** /cms/v3/hubdb/tables/{tableIdOrName}/draft | Update an existing table
*TablesAPI* | [**PostCmsV3HubdbTablesCreateTable**](docs/TablesAPI.md#postcmsv3hubdbtablescreatetable) | **Post** /cms/v3/hubdb/tables | Create a new table
*TablesAPI* | [**PostCmsV3HubdbTablesTableIdOrNameDraftCloneCloneDraftTable**](docs/TablesAPI.md#postcmsv3hubdbtablestableidornamedraftcloneclonedrafttable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/clone | Clone a table
*TablesAPI* | [**PostCmsV3HubdbTablesTableIdOrNameDraftImportImportDraftTable**](docs/TablesAPI.md#postcmsv3hubdbtablestableidornamedraftimportimportdrafttable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/import | Import data into draft table
*TablesAPI* | [**PostCmsV3HubdbTablesTableIdOrNameDraftPublishPublishDraftTable**](docs/TablesAPI.md#postcmsv3hubdbtablestableidornamedraftpublishpublishdrafttable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/publish | Publish a table from draft
*TablesAPI* | [**PostCmsV3HubdbTablesTableIdOrNameDraftResetResetDraftTable**](docs/TablesAPI.md#postcmsv3hubdbtablestableidornamedraftresetresetdrafttable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/reset | Reset a draft table
*TablesAPI* | [**PostCmsV3HubdbTablesTableIdOrNameUnpublishUnpublishTable**](docs/TablesAPI.md#postcmsv3hubdbtablestableidornameunpublishunpublishtable) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/unpublish | Unpublish a table


## Documentation For Models

 - [BatchInputHubDbTableRowBatchCloneRequest](docs/BatchInputHubDbTableRowBatchCloneRequest.md)
 - [BatchInputHubDbTableRowV3BatchUpdateRequest](docs/BatchInputHubDbTableRowV3BatchUpdateRequest.md)
 - [BatchInputHubDbTableRowV3Request](docs/BatchInputHubDbTableRowV3Request.md)
 - [BatchInputString](docs/BatchInputString.md)
 - [BatchResponseHubDbTableRowV3](docs/BatchResponseHubDbTableRowV3.md)
 - [BatchResponseHubDbTableRowV3WithErrors](docs/BatchResponseHubDbTableRowV3WithErrors.md)
 - [BoundedNextPage](docs/BoundedNextPage.md)
 - [BoundedPaging](docs/BoundedPaging.md)
 - [CollectionResponseWithTotalHubDbTableV3ForwardPaging](docs/CollectionResponseWithTotalHubDbTableV3ForwardPaging.md)
 - [Column](docs/Column.md)
 - [ColumnRequest](docs/ColumnRequest.md)
 - [Error](docs/Error.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [ForeignId](docs/ForeignId.md)
 - [ForwardPaging](docs/ForwardPaging.md)
 - [HubDbTableCloneRequest](docs/HubDbTableCloneRequest.md)
 - [HubDbTableRowBatchCloneRequest](docs/HubDbTableRowBatchCloneRequest.md)
 - [HubDbTableRowV3](docs/HubDbTableRowV3.md)
 - [HubDbTableRowV3BatchUpdateRequest](docs/HubDbTableRowV3BatchUpdateRequest.md)
 - [HubDbTableRowV3Request](docs/HubDbTableRowV3Request.md)
 - [HubDbTableV3](docs/HubDbTableV3.md)
 - [HubDbTableV3Request](docs/HubDbTableV3Request.md)
 - [ImportResult](docs/ImportResult.md)
 - [NextPage](docs/NextPage.md)
 - [Option](docs/Option.md)
 - [Paging](docs/Paging.md)
 - [PreviousPage](docs/PreviousPage.md)
 - [RandomAccessCollectionResponseWithTotalHubDbTableRowV3](docs/RandomAccessCollectionResponseWithTotalHubDbTableRowV3.md)
 - [SimpleUser](docs/SimpleUser.md)
 - [StandardError](docs/StandardError.md)
 - [StreamingCollectionResponseWithTotalHubDbTableRowV3](docs/StreamingCollectionResponseWithTotalHubDbTableRowV3.md)
 - [UnifiedCollectionResponseWithTotalBaseHubDbTableRowV3](docs/UnifiedCollectionResponseWithTotalBaseHubDbTableRowV3.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **hubdb**: Read from and write to my HubDB

Example

```go
auth := context.WithValue(context.Background(), hubdb.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, hubdb.ContextOAuth2, tokenSource)
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
		hubdb.ContextAPIKeys,
		map[string]hubdb.APIKey{
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



