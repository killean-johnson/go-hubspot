/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package properties

import (
	"bytes"
	"context"
	"io"
	"net/http"
	
	"github.com/killean-johnson/go-hubspot"
"net/url"
	"strings"
)


// GroupsAPIService GroupsAPI service
type GroupsAPIService service

type ApiDeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchiveRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	objectType string
	groupName string
}

func (r ApiDeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchiveExecute(r)
}

/*
DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchive Archive a property group

Move a property group identified by {groupName} to the recycling bin.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param groupName
 @return ApiDeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchiveRequest
*/
func (a *GroupsAPIService) DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchive(ctx context.Context, objectType string, groupName string) ApiDeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchiveRequest {
	return ApiDeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchiveRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
		groupName: groupName,
	}
}

// Execute executes the request
func (a *GroupsAPIService) DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchiveExecute(r ApiDeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.DeleteCrmV3PropertiesObjectTypeGroupsGroupNameArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/properties/{objectType}/groups/{groupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupName"+"}", url.PathEscape(parameterValueToString(r.groupName, "groupName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCrmV3PropertiesObjectTypeGroupsGetAllRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	objectType string
}

func (r ApiGetCrmV3PropertiesObjectTypeGroupsGetAllRequest) Execute() (*CollectionResponsePropertyGroupNoPaging, *http.Response, error) {
	return r.ApiService.GetCrmV3PropertiesObjectTypeGroupsGetAllExecute(r)
}

/*
GetCrmV3PropertiesObjectTypeGroupsGetAll Read all property groups

Read all existing property groups for the specified object type and HubSpot account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @return ApiGetCrmV3PropertiesObjectTypeGroupsGetAllRequest
*/
func (a *GroupsAPIService) GetCrmV3PropertiesObjectTypeGroupsGetAll(ctx context.Context, objectType string) ApiGetCrmV3PropertiesObjectTypeGroupsGetAllRequest {
	return ApiGetCrmV3PropertiesObjectTypeGroupsGetAllRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//  @return CollectionResponsePropertyGroupNoPaging
func (a *GroupsAPIService) GetCrmV3PropertiesObjectTypeGroupsGetAllExecute(r ApiGetCrmV3PropertiesObjectTypeGroupsGetAllRequest) (*CollectionResponsePropertyGroupNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponsePropertyGroupNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetCrmV3PropertiesObjectTypeGroupsGetAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/properties/{objectType}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCrmV3PropertiesObjectTypeGroupsGroupNameGetByNameRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	objectType string
	groupName string
}

func (r ApiGetCrmV3PropertiesObjectTypeGroupsGroupNameGetByNameRequest) Execute() (*PropertyGroup, *http.Response, error) {
	return r.ApiService.GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByNameExecute(r)
}

/*
GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName Read a property group

Read a property group identified by {groupName}.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param groupName
 @return ApiGetCrmV3PropertiesObjectTypeGroupsGroupNameGetByNameRequest
*/
func (a *GroupsAPIService) GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName(ctx context.Context, objectType string, groupName string) ApiGetCrmV3PropertiesObjectTypeGroupsGroupNameGetByNameRequest {
	return ApiGetCrmV3PropertiesObjectTypeGroupsGroupNameGetByNameRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
		groupName: groupName,
	}
}

// Execute executes the request
//  @return PropertyGroup
func (a *GroupsAPIService) GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByNameExecute(r ApiGetCrmV3PropertiesObjectTypeGroupsGroupNameGetByNameRequest) (*PropertyGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PropertyGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetCrmV3PropertiesObjectTypeGroupsGroupNameGetByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/properties/{objectType}/groups/{groupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupName"+"}", url.PathEscape(parameterValueToString(r.groupName, "groupName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	objectType string
	groupName string
	propertyGroupUpdate *PropertyGroupUpdate
}

func (r ApiPatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateRequest) PropertyGroupUpdate(propertyGroupUpdate PropertyGroupUpdate) ApiPatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateRequest {
	r.propertyGroupUpdate = &propertyGroupUpdate
	return r
}

func (r ApiPatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateRequest) Execute() (*PropertyGroup, *http.Response, error) {
	return r.ApiService.PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateExecute(r)
}

/*
PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate Update a property group

Perform a partial update of a property group identified by {groupName}. Provided fields will be overwritten.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param groupName
 @return ApiPatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateRequest
*/
func (a *GroupsAPIService) PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate(ctx context.Context, objectType string, groupName string) ApiPatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateRequest {
	return ApiPatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
		groupName: groupName,
	}
}

// Execute executes the request
//  @return PropertyGroup
func (a *GroupsAPIService) PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateExecute(r ApiPatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdateRequest) (*PropertyGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PropertyGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.PatchCrmV3PropertiesObjectTypeGroupsGroupNameUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/properties/{objectType}/groups/{groupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupName"+"}", url.PathEscape(parameterValueToString(r.groupName, "groupName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.propertyGroupUpdate == nil {
		return localVarReturnValue, nil, reportError("propertyGroupUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.propertyGroupUpdate
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV3PropertiesObjectTypeGroupsCreateRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	objectType string
	propertyGroupCreate *PropertyGroupCreate
}

func (r ApiPostCrmV3PropertiesObjectTypeGroupsCreateRequest) PropertyGroupCreate(propertyGroupCreate PropertyGroupCreate) ApiPostCrmV3PropertiesObjectTypeGroupsCreateRequest {
	r.propertyGroupCreate = &propertyGroupCreate
	return r
}

func (r ApiPostCrmV3PropertiesObjectTypeGroupsCreateRequest) Execute() (*PropertyGroup, *http.Response, error) {
	return r.ApiService.PostCrmV3PropertiesObjectTypeGroupsCreateExecute(r)
}

/*
PostCrmV3PropertiesObjectTypeGroupsCreate Create a property group

Create and return a copy of a new property group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @return ApiPostCrmV3PropertiesObjectTypeGroupsCreateRequest
*/
func (a *GroupsAPIService) PostCrmV3PropertiesObjectTypeGroupsCreate(ctx context.Context, objectType string) ApiPostCrmV3PropertiesObjectTypeGroupsCreateRequest {
	return ApiPostCrmV3PropertiesObjectTypeGroupsCreateRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//  @return PropertyGroup
func (a *GroupsAPIService) PostCrmV3PropertiesObjectTypeGroupsCreateExecute(r ApiPostCrmV3PropertiesObjectTypeGroupsCreateRequest) (*PropertyGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PropertyGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.PostCrmV3PropertiesObjectTypeGroupsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/properties/{objectType}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.propertyGroupCreate == nil {
		return localVarReturnValue, nil, reportError("propertyGroupCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.propertyGroupCreate
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
