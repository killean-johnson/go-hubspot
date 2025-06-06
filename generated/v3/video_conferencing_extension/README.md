# Go API client for video_conferencing_extension

These APIs allow you to specify URLs that can be used to interact with a video conferencing application, to allow HubSpot to add video conference links to meeting requests with contacts.

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
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import video_conferencing_extension "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `video_conferencing_extension.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), video_conferencing_extension.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `video_conferencing_extension.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), video_conferencing_extension.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `video_conferencing_extension.ContextOperationServerIndices` and `video_conferencing_extension.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), video_conferencing_extension.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), video_conferencing_extension.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*SettingsAPI* | [**DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive**](docs/SettingsAPI.md#deletecrmv3extensionsvideoconferencingsettingsappidarchive) | **Delete** /crm/v3/extensions/videoconferencing/settings/{appId} | Delete settings
*SettingsAPI* | [**GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById**](docs/SettingsAPI.md#getcrmv3extensionsvideoconferencingsettingsappidgetbyid) | **Get** /crm/v3/extensions/videoconferencing/settings/{appId} | Get settings
*SettingsAPI* | [**PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace**](docs/SettingsAPI.md#putcrmv3extensionsvideoconferencingsettingsappidreplace) | **Put** /crm/v3/extensions/videoconferencing/settings/{appId} | Update settings


## Documentation For Models

 - [Error](docs/Error.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [ExternalSettings](docs/ExternalSettings.md)


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
		video_conferencing_extension.ContextAPIKeys,
		map[string]video_conferencing_extension.APIKey{
			"developer_hapikey": {Key: "API_KEY_STRING"},
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



