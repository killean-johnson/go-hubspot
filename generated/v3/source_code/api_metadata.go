/*
Source Code

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package source_code

import (
	"bytes"
	"context"
	"io"
	"net/http"
	
	"github.com/killean-johnson/go-hubspot"
"net/url"
	"strings"
)


// MetadataAPIService MetadataAPI service
type MetadataAPIService service

type ApiGetCmsV3SourceCodeEnvironmentMetadataPathGetRequest struct {
	ctx context.Context
	ApiService *MetadataAPIService
	environment string
	path string
	properties *string
}

func (r ApiGetCmsV3SourceCodeEnvironmentMetadataPathGetRequest) Properties(properties string) ApiGetCmsV3SourceCodeEnvironmentMetadataPathGetRequest {
	r.properties = &properties
	return r
}

func (r ApiGetCmsV3SourceCodeEnvironmentMetadataPathGetRequest) Execute() (*AssetFileMetadata, *http.Response, error) {
	return r.ApiService.GetCmsV3SourceCodeEnvironmentMetadataPathGetExecute(r)
}

/*
GetCmsV3SourceCodeEnvironmentMetadataPathGet Get the metadata for a file

Gets the metadata object for the file at the specified path in the specified environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environment The environment of the file (\"draft\" or \"published\").
 @param path The file system location of the file.
 @return ApiGetCmsV3SourceCodeEnvironmentMetadataPathGetRequest
*/
func (a *MetadataAPIService) GetCmsV3SourceCodeEnvironmentMetadataPathGet(ctx context.Context, environment string, path string) ApiGetCmsV3SourceCodeEnvironmentMetadataPathGetRequest {
	return ApiGetCmsV3SourceCodeEnvironmentMetadataPathGetRequest{
		ApiService: a,
		ctx: ctx,
		environment: environment,
		path: path,
	}
}

// Execute executes the request
//  @return AssetFileMetadata
func (a *MetadataAPIService) GetCmsV3SourceCodeEnvironmentMetadataPathGetExecute(r ApiGetCmsV3SourceCodeEnvironmentMetadataPathGetRequest) (*AssetFileMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssetFileMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataAPIService.GetCmsV3SourceCodeEnvironmentMetadataPathGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/{environment}/metadata/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", url.PathEscape(parameterValueToString(r.environment, "environment")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
