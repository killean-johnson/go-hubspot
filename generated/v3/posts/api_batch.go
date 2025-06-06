/*
Posts

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posts

import (
	"bytes"
	"context"
	"io"
	"net/http"
	
	"github.com/killean-johnson/go-hubspot"
"net/url"
)


// BatchAPIService BatchAPI service
type BatchAPIService service

type ApiPostCmsV3BlogsPostsBatchArchiveArchiveRequest struct {
	ctx context.Context
	ApiService *BatchAPIService
	batchInputString *BatchInputString
}

// The JSON array of Blog Post ids.
func (r ApiPostCmsV3BlogsPostsBatchArchiveArchiveRequest) BatchInputString(batchInputString BatchInputString) ApiPostCmsV3BlogsPostsBatchArchiveArchiveRequest {
	r.batchInputString = &batchInputString
	return r
}

func (r ApiPostCmsV3BlogsPostsBatchArchiveArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCmsV3BlogsPostsBatchArchiveArchiveExecute(r)
}

/*
PostCmsV3BlogsPostsBatchArchiveArchive Delete a batch of blog posts

Delete a blog post by ID. 
Note: This is not the same as the in-app `archive` function. To perform a dashboard `archive` send an normal update with the `archivedInDashboard` field set to `true`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCmsV3BlogsPostsBatchArchiveArchiveRequest
*/
func (a *BatchAPIService) PostCmsV3BlogsPostsBatchArchiveArchive(ctx context.Context) ApiPostCmsV3BlogsPostsBatchArchiveArchiveRequest {
	return ApiPostCmsV3BlogsPostsBatchArchiveArchiveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *BatchAPIService) PostCmsV3BlogsPostsBatchArchiveArchiveExecute(r ApiPostCmsV3BlogsPostsBatchArchiveArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCmsV3BlogsPostsBatchArchiveArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/blogs/posts/batch/archive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputString == nil {
		return nil, reportError("batchInputString is required and must be specified")
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
	localVarPostBody = r.batchInputString
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

type ApiPostCmsV3BlogsPostsBatchCreateCreateRequest struct {
	ctx context.Context
	ApiService *BatchAPIService
	batchInputBlogPost *BatchInputBlogPost
}

// The JSON array of new Blog Posts to create.
func (r ApiPostCmsV3BlogsPostsBatchCreateCreateRequest) BatchInputBlogPost(batchInputBlogPost BatchInputBlogPost) ApiPostCmsV3BlogsPostsBatchCreateCreateRequest {
	r.batchInputBlogPost = &batchInputBlogPost
	return r
}

func (r ApiPostCmsV3BlogsPostsBatchCreateCreateRequest) Execute() (*BatchResponseBlogPost, *http.Response, error) {
	return r.ApiService.PostCmsV3BlogsPostsBatchCreateCreateExecute(r)
}

/*
PostCmsV3BlogsPostsBatchCreateCreate Create a batch of blog posts

Create a batch of blog posts, specifying their content in the request body.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCmsV3BlogsPostsBatchCreateCreateRequest
*/
func (a *BatchAPIService) PostCmsV3BlogsPostsBatchCreateCreate(ctx context.Context) ApiPostCmsV3BlogsPostsBatchCreateCreateRequest {
	return ApiPostCmsV3BlogsPostsBatchCreateCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BatchResponseBlogPost
func (a *BatchAPIService) PostCmsV3BlogsPostsBatchCreateCreateExecute(r ApiPostCmsV3BlogsPostsBatchCreateCreateRequest) (*BatchResponseBlogPost, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchResponseBlogPost
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCmsV3BlogsPostsBatchCreateCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/blogs/posts/batch/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputBlogPost == nil {
		return localVarReturnValue, nil, reportError("batchInputBlogPost is required and must be specified")
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
	localVarPostBody = r.batchInputBlogPost
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

type ApiPostCmsV3BlogsPostsBatchReadReadRequest struct {
	ctx context.Context
	ApiService *BatchAPIService
	batchInputString *BatchInputString
	archived *bool
}

// The JSON array of Blog Post ids.
func (r ApiPostCmsV3BlogsPostsBatchReadReadRequest) BatchInputString(batchInputString BatchInputString) ApiPostCmsV3BlogsPostsBatchReadReadRequest {
	r.batchInputString = &batchInputString
	return r
}

// Specifies whether to return deleted blog posts Defaults to &#x60;false&#x60;.
func (r ApiPostCmsV3BlogsPostsBatchReadReadRequest) Archived(archived bool) ApiPostCmsV3BlogsPostsBatchReadReadRequest {
	r.archived = &archived
	return r
}

func (r ApiPostCmsV3BlogsPostsBatchReadReadRequest) Execute() (*BatchResponseBlogPost, *http.Response, error) {
	return r.ApiService.PostCmsV3BlogsPostsBatchReadReadExecute(r)
}

/*
PostCmsV3BlogsPostsBatchReadRead Retrieve a batch of Blog Posts

Retrieve a batch of blog posts by ID. identified in the request body.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCmsV3BlogsPostsBatchReadReadRequest
*/
func (a *BatchAPIService) PostCmsV3BlogsPostsBatchReadRead(ctx context.Context) ApiPostCmsV3BlogsPostsBatchReadReadRequest {
	return ApiPostCmsV3BlogsPostsBatchReadReadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BatchResponseBlogPost
func (a *BatchAPIService) PostCmsV3BlogsPostsBatchReadReadExecute(r ApiPostCmsV3BlogsPostsBatchReadReadRequest) (*BatchResponseBlogPost, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchResponseBlogPost
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCmsV3BlogsPostsBatchReadRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/blogs/posts/batch/read"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputString == nil {
		return localVarReturnValue, nil, reportError("batchInputString is required and must be specified")
	}

	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
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
	localVarPostBody = r.batchInputString
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

type ApiPostCmsV3BlogsPostsBatchUpdateUpdateRequest struct {
	ctx context.Context
	ApiService *BatchAPIService
	batchInputJsonNode *BatchInputJsonNode
	archived *bool
}

// A JSON array of the JSON representations of the updated Blog Posts.
func (r ApiPostCmsV3BlogsPostsBatchUpdateUpdateRequest) BatchInputJsonNode(batchInputJsonNode BatchInputJsonNode) ApiPostCmsV3BlogsPostsBatchUpdateUpdateRequest {
	r.batchInputJsonNode = &batchInputJsonNode
	return r
}

// Specifies whether to update deleted Blog Posts. Defaults to &#x60;false&#x60;.
func (r ApiPostCmsV3BlogsPostsBatchUpdateUpdateRequest) Archived(archived bool) ApiPostCmsV3BlogsPostsBatchUpdateUpdateRequest {
	r.archived = &archived
	return r
}

func (r ApiPostCmsV3BlogsPostsBatchUpdateUpdateRequest) Execute() (*BatchResponseBlogPost, *http.Response, error) {
	return r.ApiService.PostCmsV3BlogsPostsBatchUpdateUpdateExecute(r)
}

/*
PostCmsV3BlogsPostsBatchUpdateUpdate Update a batch of Blog Posts

Update a batch of blog posts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCmsV3BlogsPostsBatchUpdateUpdateRequest
*/
func (a *BatchAPIService) PostCmsV3BlogsPostsBatchUpdateUpdate(ctx context.Context) ApiPostCmsV3BlogsPostsBatchUpdateUpdateRequest {
	return ApiPostCmsV3BlogsPostsBatchUpdateUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BatchResponseBlogPost
func (a *BatchAPIService) PostCmsV3BlogsPostsBatchUpdateUpdateExecute(r ApiPostCmsV3BlogsPostsBatchUpdateUpdateRequest) (*BatchResponseBlogPost, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchResponseBlogPost
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCmsV3BlogsPostsBatchUpdateUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/blogs/posts/batch/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputJsonNode == nil {
		return localVarReturnValue, nil, reportError("batchInputJsonNode is required and must be specified")
	}

	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
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
	localVarPostBody = r.batchInputJsonNode
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
