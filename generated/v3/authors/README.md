# Go API client for authors

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

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
import authors "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `authors.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), authors.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `authors.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), authors.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `authors.ContextOperationServerIndices` and `authors.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), authors.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), authors.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BlogAuthorsAPI* | [**DeleteCmsV3BlogsAuthorsObjectIdArchive**](docs/BlogAuthorsAPI.md#deletecmsv3blogsauthorsobjectidarchive) | **Delete** /cms/v3/blogs/authors/{objectId} | Delete a Blog Author
*BlogAuthorsAPI* | [**GetCmsV3BlogsAuthorsGetPage**](docs/BlogAuthorsAPI.md#getcmsv3blogsauthorsgetpage) | **Get** /cms/v3/blogs/authors | Get all Blog Authors
*BlogAuthorsAPI* | [**GetCmsV3BlogsAuthorsObjectIdGetById**](docs/BlogAuthorsAPI.md#getcmsv3blogsauthorsobjectidgetbyid) | **Get** /cms/v3/blogs/authors/{objectId} | Retrieve a Blog Author
*BlogAuthorsAPI* | [**PatchCmsV3BlogsAuthorsObjectIdUpdate**](docs/BlogAuthorsAPI.md#patchcmsv3blogsauthorsobjectidupdate) | **Patch** /cms/v3/blogs/authors/{objectId} | Update a Blog Author
*BlogAuthorsAPI* | [**PostCmsV3BlogsAuthorsBatchArchiveArchiveBatch**](docs/BlogAuthorsAPI.md#postcmsv3blogsauthorsbatcharchivearchivebatch) | **Post** /cms/v3/blogs/authors/batch/archive | Delete a batch of Blog Authors
*BlogAuthorsAPI* | [**PostCmsV3BlogsAuthorsBatchCreateCreateBatch**](docs/BlogAuthorsAPI.md#postcmsv3blogsauthorsbatchcreatecreatebatch) | **Post** /cms/v3/blogs/authors/batch/create | Create a batch of Blog Authors
*BlogAuthorsAPI* | [**PostCmsV3BlogsAuthorsBatchReadReadBatch**](docs/BlogAuthorsAPI.md#postcmsv3blogsauthorsbatchreadreadbatch) | **Post** /cms/v3/blogs/authors/batch/read | Retrieve a batch of Blog Authors
*BlogAuthorsAPI* | [**PostCmsV3BlogsAuthorsBatchUpdateUpdateBatch**](docs/BlogAuthorsAPI.md#postcmsv3blogsauthorsbatchupdateupdatebatch) | **Post** /cms/v3/blogs/authors/batch/update | Update a batch of Blog Authors
*BlogAuthorsAPI* | [**PostCmsV3BlogsAuthorsCreate**](docs/BlogAuthorsAPI.md#postcmsv3blogsauthorscreate) | **Post** /cms/v3/blogs/authors | Create a new Blog Author
*BlogAuthorsAPI* | [**PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupAttachToLangGroup**](docs/BlogAuthorsAPI.md#postcmsv3blogsauthorsmultilanguageattachtolanggroupattachtolanggroup) | **Post** /cms/v3/blogs/authors/multi-language/attach-to-lang-group | Attach a Blog Author to a multi-language group
*BlogAuthorsAPI* | [**PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationCreateLangVariation**](docs/BlogAuthorsAPI.md#postcmsv3blogsauthorsmultilanguagecreatelanguagevariationcreatelangvariation) | **Post** /cms/v3/blogs/authors/multi-language/create-language-variation | Create a new language variation
*BlogAuthorsAPI* | [**PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupDetachFromLangGroup**](docs/BlogAuthorsAPI.md#postcmsv3blogsauthorsmultilanguagedetachfromlanggroupdetachfromlanggroup) | **Post** /cms/v3/blogs/authors/multi-language/detach-from-lang-group | Detach a Blog Author from a multi-language group
*BlogAuthorsAPI* | [**PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesUpdateLangs**](docs/BlogAuthorsAPI.md#postcmsv3blogsauthorsmultilanguageupdatelanguagesupdatelangs) | **Post** /cms/v3/blogs/authors/multi-language/update-languages | Update languages of multi-language group
*BlogAuthorsAPI* | [**PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimarySetLangPrimary**](docs/BlogAuthorsAPI.md#putcmsv3blogsauthorsmultilanguagesetnewlangprimarysetlangprimary) | **Put** /cms/v3/blogs/authors/multi-language/set-new-lang-primary | Set a new primary language


## Documentation For Models

 - [AttachToLangPrimaryRequestVNext](docs/AttachToLangPrimaryRequestVNext.md)
 - [BatchInputBlogAuthor](docs/BatchInputBlogAuthor.md)
 - [BatchInputJsonNode](docs/BatchInputJsonNode.md)
 - [BatchInputString](docs/BatchInputString.md)
 - [BatchResponseBlogAuthor](docs/BatchResponseBlogAuthor.md)
 - [BatchResponseBlogAuthorWithErrors](docs/BatchResponseBlogAuthorWithErrors.md)
 - [BlogAuthor](docs/BlogAuthor.md)
 - [BlogAuthorCloneRequestVNext](docs/BlogAuthorCloneRequestVNext.md)
 - [CollectionResponseWithTotalBlogAuthorForwardPaging](docs/CollectionResponseWithTotalBlogAuthorForwardPaging.md)
 - [DetachFromLangGroupRequestVNext](docs/DetachFromLangGroupRequestVNext.md)
 - [Error](docs/Error.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [ForwardPaging](docs/ForwardPaging.md)
 - [NextPage](docs/NextPage.md)
 - [SetNewLanguagePrimaryRequestVNext](docs/SetNewLanguagePrimaryRequestVNext.md)
 - [StandardError](docs/StandardError.md)
 - [UpdateLanguagesRequestVNext](docs/UpdateLanguagesRequestVNext.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **content**: Read from and write to my Content

Example

```go
auth := context.WithValue(context.Background(), authors.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, authors.ContextOAuth2, tokenSource)
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
		authors.ContextAPIKeys,
		map[string]authors.APIKey{
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



