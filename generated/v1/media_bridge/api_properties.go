/*
CMS Media Bridge

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package media_bridge

import (
	"bytes"
	"context"
	"io"
	"net/http"
	
	"github.com/killean-johnson/go-hubspot"
"net/url"
	"strings"
)


// PropertiesAPIService PropertiesAPI service
type PropertiesAPIService service

type ApiDeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest struct {
	ctx context.Context
	ApiService *PropertiesAPIService
	objectType string
	propertyName string
}

func (r ApiDeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyNameExecute(r)
}

/*
DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyName Method for DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyName

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param propertyName
 @return ApiDeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest
*/
func (a *PropertiesAPIService) DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyName(ctx context.Context, objectType string, propertyName string) ApiDeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest {
	return ApiDeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
		propertyName: propertyName,
	}
}

// Execute executes the request
func (a *PropertiesAPIService) DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyNameExecute(r ApiDeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesAPIService.DeleteMediaBridgeV1AppIdPropertiesObjectTypePropertyName")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media-bridge/v1/{appId}/properties/{objectType}/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", url.PathEscape(parameterValueToString(r.propertyName, "propertyName")), -1)

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

type ApiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest struct {
	ctx context.Context
	ApiService *PropertiesAPIService
	objectType string
	archived *bool
	properties *string
}

// Whether to return only results that have been archived.
func (r ApiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest) Archived(archived bool) ApiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest {
	r.archived = &archived
	return r
}

func (r ApiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest) Properties(properties string) ApiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest {
	r.properties = &properties
	return r
}

func (r ApiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest) Execute() (*CollectionResponsePropertyNoPaging, *http.Response, error) {
	return r.ApiService.GetMediaBridgeV1AppIdPropertiesObjectTypeExecute(r)
}

/*
GetMediaBridgeV1AppIdPropertiesObjectType Method for GetMediaBridgeV1AppIdPropertiesObjectType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @return ApiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest
*/
func (a *PropertiesAPIService) GetMediaBridgeV1AppIdPropertiesObjectType(ctx context.Context, objectType string) ApiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest {
	return ApiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//  @return CollectionResponsePropertyNoPaging
func (a *PropertiesAPIService) GetMediaBridgeV1AppIdPropertiesObjectTypeExecute(r ApiGetMediaBridgeV1AppIdPropertiesObjectTypeRequest) (*CollectionResponsePropertyNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponsePropertyNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesAPIService.GetMediaBridgeV1AppIdPropertiesObjectType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media-bridge/v1/{appId}/properties/{objectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
	} else {
		var defaultValue bool = false
		r.archived = &defaultValue
	}
	if r.properties != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "properties", r.properties, "form", "")
	}
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

type ApiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest struct {
	ctx context.Context
	ApiService *PropertiesAPIService
	objectType string
	propertyName string
	archived *bool
	properties *string
}

// Whether to return only results that have been archived.
func (r ApiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest) Archived(archived bool) ApiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest {
	r.archived = &archived
	return r
}

func (r ApiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest) Properties(properties string) ApiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest {
	r.properties = &properties
	return r
}

func (r ApiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest) Execute() (*Property, *http.Response, error) {
	return r.ApiService.GetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameExecute(r)
}

/*
GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName Method for GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param propertyName
 @return ApiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest
*/
func (a *PropertiesAPIService) GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName(ctx context.Context, objectType string, propertyName string) ApiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest {
	return ApiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
		propertyName: propertyName,
	}
}

// Execute executes the request
//  @return Property
func (a *PropertiesAPIService) GetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameExecute(r ApiGetMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest) (*Property, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Property
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesAPIService.GetMediaBridgeV1AppIdPropertiesObjectTypePropertyName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media-bridge/v1/{appId}/properties/{objectType}/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", url.PathEscape(parameterValueToString(r.propertyName, "propertyName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
	} else {
		var defaultValue bool = false
		r.archived = &defaultValue
	}
	if r.properties != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "properties", r.properties, "form", "")
	}
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

type ApiPatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest struct {
	ctx context.Context
	ApiService *PropertiesAPIService
	objectType string
	propertyName string
	mediaBridgePropertyUpdate *MediaBridgePropertyUpdate
}

func (r ApiPatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest) MediaBridgePropertyUpdate(mediaBridgePropertyUpdate MediaBridgePropertyUpdate) ApiPatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest {
	r.mediaBridgePropertyUpdate = &mediaBridgePropertyUpdate
	return r
}

func (r ApiPatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest) Execute() (*Property, *http.Response, error) {
	return r.ApiService.PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameExecute(r)
}

/*
PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName Method for PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param propertyName
 @return ApiPatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest
*/
func (a *PropertiesAPIService) PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName(ctx context.Context, objectType string, propertyName string) ApiPatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest {
	return ApiPatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
		propertyName: propertyName,
	}
}

// Execute executes the request
//  @return Property
func (a *PropertiesAPIService) PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameExecute(r ApiPatchMediaBridgeV1AppIdPropertiesObjectTypePropertyNameRequest) (*Property, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Property
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesAPIService.PatchMediaBridgeV1AppIdPropertiesObjectTypePropertyName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media-bridge/v1/{appId}/properties/{objectType}/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", url.PathEscape(parameterValueToString(r.propertyName, "propertyName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mediaBridgePropertyUpdate == nil {
		return localVarReturnValue, nil, reportError("mediaBridgePropertyUpdate is required and must be specified")
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
	localVarPostBody = r.mediaBridgePropertyUpdate
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

type ApiPostMediaBridgeV1AppIdPropertiesObjectTypeRequest struct {
	ctx context.Context
	ApiService *PropertiesAPIService
	objectType string
	propertyCreate *PropertyCreate
}

func (r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeRequest) PropertyCreate(propertyCreate PropertyCreate) ApiPostMediaBridgeV1AppIdPropertiesObjectTypeRequest {
	r.propertyCreate = &propertyCreate
	return r
}

func (r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeRequest) Execute() (*Property, *http.Response, error) {
	return r.ApiService.PostMediaBridgeV1AppIdPropertiesObjectTypeExecute(r)
}

/*
PostMediaBridgeV1AppIdPropertiesObjectType Method for PostMediaBridgeV1AppIdPropertiesObjectType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @return ApiPostMediaBridgeV1AppIdPropertiesObjectTypeRequest
*/
func (a *PropertiesAPIService) PostMediaBridgeV1AppIdPropertiesObjectType(ctx context.Context, objectType string) ApiPostMediaBridgeV1AppIdPropertiesObjectTypeRequest {
	return ApiPostMediaBridgeV1AppIdPropertiesObjectTypeRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//  @return Property
func (a *PropertiesAPIService) PostMediaBridgeV1AppIdPropertiesObjectTypeExecute(r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeRequest) (*Property, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Property
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesAPIService.PostMediaBridgeV1AppIdPropertiesObjectType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media-bridge/v1/{appId}/properties/{objectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.propertyCreate == nil {
		return localVarReturnValue, nil, reportError("propertyCreate is required and must be specified")
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
	localVarPostBody = r.propertyCreate
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

type ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveRequest struct {
	ctx context.Context
	ApiService *PropertiesAPIService
	objectType string
	batchInputPropertyName *BatchInputPropertyName
}

func (r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveRequest) BatchInputPropertyName(batchInputPropertyName BatchInputPropertyName) ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveRequest {
	r.batchInputPropertyName = &batchInputPropertyName
	return r
}

func (r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveExecute(r)
}

/*
PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchive Method for PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchive

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @return ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveRequest
*/
func (a *PropertiesAPIService) PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchive(ctx context.Context, objectType string) ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveRequest {
	return ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
	}
}

// Execute executes the request
func (a *PropertiesAPIService) PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveExecute(r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesAPIService.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media-bridge/v1/{appId}/properties/{objectType}/batch/archive"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPropertyName == nil {
		return nil, reportError("batchInputPropertyName is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.batchInputPropertyName
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

type ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateRequest struct {
	ctx context.Context
	ApiService *PropertiesAPIService
	objectType string
	batchInputPropertyCreate *BatchInputPropertyCreate
}

func (r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateRequest) BatchInputPropertyCreate(batchInputPropertyCreate BatchInputPropertyCreate) ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateRequest {
	r.batchInputPropertyCreate = &batchInputPropertyCreate
	return r
}

func (r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateRequest) Execute() (*BatchResponseProperty, *http.Response, error) {
	return r.ApiService.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateExecute(r)
}

/*
PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate Method for PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @return ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateRequest
*/
func (a *PropertiesAPIService) PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate(ctx context.Context, objectType string) ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateRequest {
	return ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//  @return BatchResponseProperty
func (a *PropertiesAPIService) PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateExecute(r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreateRequest) (*BatchResponseProperty, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchResponseProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesAPIService.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media-bridge/v1/{appId}/properties/{objectType}/batch/create"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPropertyCreate == nil {
		return localVarReturnValue, nil, reportError("batchInputPropertyCreate is required and must be specified")
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
	localVarPostBody = r.batchInputPropertyCreate
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

type ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadRequest struct {
	ctx context.Context
	ApiService *PropertiesAPIService
	objectType string
	batchReadInputPropertyName *BatchReadInputPropertyName
}

func (r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadRequest) BatchReadInputPropertyName(batchReadInputPropertyName BatchReadInputPropertyName) ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadRequest {
	r.batchReadInputPropertyName = &batchReadInputPropertyName
	return r
}

func (r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadRequest) Execute() (*BatchResponseProperty, *http.Response, error) {
	return r.ApiService.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadExecute(r)
}

/*
PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead Method for PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @return ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadRequest
*/
func (a *PropertiesAPIService) PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead(ctx context.Context, objectType string) ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadRequest {
	return ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadRequest{
		ApiService: a,
		ctx: ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//  @return BatchResponseProperty
func (a *PropertiesAPIService) PostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadExecute(r ApiPostMediaBridgeV1AppIdPropertiesObjectTypeBatchReadRequest) (*BatchResponseProperty, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchResponseProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesAPIService.PostMediaBridgeV1AppIdPropertiesObjectTypeBatchRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media-bridge/v1/{appId}/properties/{objectType}/batch/read"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchReadInputPropertyName == nil {
		return localVarReturnValue, nil, reportError("batchReadInputPropertyName is required and must be specified")
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
	localVarPostBody = r.batchReadInputPropertyName
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
