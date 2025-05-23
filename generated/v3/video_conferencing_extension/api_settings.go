/*
Video Conferencing Extension

These APIs allow you to specify URLs that can be used to interact with a video conferencing application, to allow HubSpot to add video conference links to meeting requests with contacts.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package video_conferencing_extension

import (
	"bytes"
	"context"
	"io"
	"net/http"
	
	"github.com/killean-johnson/go-hubspot"
"net/url"
	"strings"
)


// SettingsAPIService SettingsAPI service
type SettingsAPIService service

type ApiDeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchiveRequest struct {
	ctx context.Context
	ApiService *SettingsAPIService
	appId int32
}

func (r ApiDeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchiveExecute(r)
}

/*
DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive Delete settings

Deletes the settings for a video conference application with the specified ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal.
 @return ApiDeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchiveRequest
*/
func (a *SettingsAPIService) DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive(ctx context.Context, appId int32) ApiDeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchiveRequest {
	return ApiDeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchiveRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
func (a *SettingsAPIService) DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchiveExecute(r ApiDeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsAPIService.DeleteCrmV3ExtensionsVideoconferencingSettingsAppIdArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/videoconferencing/settings/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

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

type ApiGetCrmV3ExtensionsVideoconferencingSettingsAppIdGetByIdRequest struct {
	ctx context.Context
	ApiService *SettingsAPIService
	appId int32
}

func (r ApiGetCrmV3ExtensionsVideoconferencingSettingsAppIdGetByIdRequest) Execute() (*ExternalSettings, *http.Response, error) {
	return r.ApiService.GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetByIdExecute(r)
}

/*
GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById Get settings

Return the settings for a video conference application with the specified ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal.
 @return ApiGetCrmV3ExtensionsVideoconferencingSettingsAppIdGetByIdRequest
*/
func (a *SettingsAPIService) GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById(ctx context.Context, appId int32) ApiGetCrmV3ExtensionsVideoconferencingSettingsAppIdGetByIdRequest {
	return ApiGetCrmV3ExtensionsVideoconferencingSettingsAppIdGetByIdRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return ExternalSettings
func (a *SettingsAPIService) GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetByIdExecute(r ApiGetCrmV3ExtensionsVideoconferencingSettingsAppIdGetByIdRequest) (*ExternalSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExternalSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsAPIService.GetCrmV3ExtensionsVideoconferencingSettingsAppIdGetById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/videoconferencing/settings/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

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

type ApiPutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceRequest struct {
	ctx context.Context
	ApiService *SettingsAPIService
	appId int32
	externalSettings *ExternalSettings
}

func (r ApiPutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceRequest) ExternalSettings(externalSettings ExternalSettings) ApiPutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceRequest {
	r.externalSettings = &externalSettings
	return r
}

func (r ApiPutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceRequest) Execute() (*ExternalSettings, *http.Response, error) {
	return r.ApiService.PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceExecute(r)
}

/*
PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace Update settings

Updates the settings for a video conference application with the specified ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the video conference application. This is the identifier of the application created in your HubSpot developer portal.
 @return ApiPutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceRequest
*/
func (a *SettingsAPIService) PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace(ctx context.Context, appId int32) ApiPutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceRequest {
	return ApiPutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return ExternalSettings
func (a *SettingsAPIService) PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceExecute(r ApiPutCrmV3ExtensionsVideoconferencingSettingsAppIdReplaceRequest) (*ExternalSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExternalSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsAPIService.PutCrmV3ExtensionsVideoconferencingSettingsAppIdReplace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/videoconferencing/settings/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalSettings == nil {
		return localVarReturnValue, nil, reportError("externalSettings is required and must be specified")
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
	localVarPostBody = r.externalSettings
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
