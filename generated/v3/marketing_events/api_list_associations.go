/*
Marketing Events

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events

import (
	"bytes"
	"context"
	"io"
	"net/http"
	
	"github.com/killean-johnson/go-hubspot"
"net/url"
	"strings"
)


// ListAssociationsAPIService ListAssociationsAPI service
type ListAssociationsAPIService service

type ApiDeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIdsRequest struct {
	ctx context.Context
	ApiService *ListAssociationsAPIService
	externalAccountId string
	externalEventId string
	listId string
}

func (r ApiDeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIdsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIdsExecute(r)
}

/*
DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIds Disassociate a list from a marketing event

Disassociates a list from a marketing event by external account id, external event id, and ILS list id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalAccountId The accountId that is associated with this marketing event in the external event application.
 @param externalEventId The id of the marketing event in the external event application.
 @param listId The ILS ID of the list.
 @return ApiDeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIdsRequest
*/
func (a *ListAssociationsAPIService) DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIds(ctx context.Context, externalAccountId string, externalEventId string, listId string) ApiDeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIdsRequest {
	return ApiDeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIdsRequest{
		ApiService: a,
		ctx: ctx,
		externalAccountId: externalAccountId,
		externalEventId: externalEventId,
		listId: listId,
	}
}

// Execute executes the request
func (a *ListAssociationsAPIService) DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIdsExecute(r ApiDeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIdsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListAssociationsAPIService.DeleteMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdDisassociateByExternalAccountAndEventIds")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/associations/{externalAccountId}/{externalEventId}/lists/{listId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalAccountId"+"}", url.PathEscape(parameterValueToString(r.externalAccountId, "externalAccountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"externalEventId"+"}", url.PathEscape(parameterValueToString(r.externalEventId, "externalEventId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

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

type ApiDeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventIdRequest struct {
	ctx context.Context
	ApiService *ListAssociationsAPIService
	marketingEventId string
	listId string
}

func (r ApiDeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventIdExecute(r)
}

/*
DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventId Disassociate a list from a marketing event

Disassociates a list from a marketing event by marketing event id and ILS list id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marketingEventId The internal id of the marketing event in HubSpot.
 @param listId The ILS ID of the list.
 @return ApiDeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventIdRequest
*/
func (a *ListAssociationsAPIService) DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventId(ctx context.Context, marketingEventId string, listId string) ApiDeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventIdRequest {
	return ApiDeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventIdRequest{
		ApiService: a,
		ctx: ctx,
		marketingEventId: marketingEventId,
		listId: listId,
	}
}

// Execute executes the request
func (a *ListAssociationsAPIService) DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventIdExecute(r ApiDeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListAssociationsAPIService.DeleteMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdDisassociateByMarketingEventId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/associations/{marketingEventId}/lists/{listId}"
	localVarPath = strings.Replace(localVarPath, "{"+"marketingEventId"+"}", url.PathEscape(parameterValueToString(r.marketingEventId, "marketingEventId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

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

type ApiGetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIdsRequest struct {
	ctx context.Context
	ApiService *ListAssociationsAPIService
	externalAccountId string
	externalEventId string
}

func (r ApiGetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIdsRequest) Execute() (*CollectionResponseWithTotalPublicListNoPaging, *http.Response, error) {
	return r.ApiService.GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIdsExecute(r)
}

/*
GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds Get lists associated with a marketing event

Gets lists associated with a marketing event by external account id and external event id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalAccountId The accountId that is associated with this marketing event in the external event application
 @param externalEventId The id of the marketing event in the external event application.
 @return ApiGetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIdsRequest
*/
func (a *ListAssociationsAPIService) GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds(ctx context.Context, externalAccountId string, externalEventId string) ApiGetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIdsRequest {
	return ApiGetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIdsRequest{
		ApiService: a,
		ctx: ctx,
		externalAccountId: externalAccountId,
		externalEventId: externalEventId,
	}
}

// Execute executes the request
//  @return CollectionResponseWithTotalPublicListNoPaging
func (a *ListAssociationsAPIService) GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIdsExecute(r ApiGetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIdsRequest) (*CollectionResponseWithTotalPublicListNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponseWithTotalPublicListNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListAssociationsAPIService.GetMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsGetAllByExternalAccountAndEventIds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/associations/{externalAccountId}/{externalEventId}/lists"
	localVarPath = strings.Replace(localVarPath, "{"+"externalAccountId"+"}", url.PathEscape(parameterValueToString(r.externalAccountId, "externalAccountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"externalEventId"+"}", url.PathEscape(parameterValueToString(r.externalEventId, "externalEventId")), -1)

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

type ApiGetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventIdRequest struct {
	ctx context.Context
	ApiService *ListAssociationsAPIService
	marketingEventId string
}

func (r ApiGetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventIdRequest) Execute() (*CollectionResponseWithTotalPublicListNoPaging, *http.Response, error) {
	return r.ApiService.GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventIdExecute(r)
}

/*
GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId Get lists associated with a marketing event

Gets lists associated with a marketing event by marketing event id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marketingEventId The internal id of the marketing event in HubSpot.
 @return ApiGetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventIdRequest
*/
func (a *ListAssociationsAPIService) GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId(ctx context.Context, marketingEventId string) ApiGetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventIdRequest {
	return ApiGetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventIdRequest{
		ApiService: a,
		ctx: ctx,
		marketingEventId: marketingEventId,
	}
}

// Execute executes the request
//  @return CollectionResponseWithTotalPublicListNoPaging
func (a *ListAssociationsAPIService) GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventIdExecute(r ApiGetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventIdRequest) (*CollectionResponseWithTotalPublicListNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponseWithTotalPublicListNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListAssociationsAPIService.GetMarketingV3MarketingEventsAssociationsMarketingEventIdListsGetAllByMarketingEventId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/associations/{marketingEventId}/lists"
	localVarPath = strings.Replace(localVarPath, "{"+"marketingEventId"+"}", url.PathEscape(parameterValueToString(r.marketingEventId, "marketingEventId")), -1)

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

type ApiPutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIdsRequest struct {
	ctx context.Context
	ApiService *ListAssociationsAPIService
	externalAccountId string
	externalEventId string
	listId string
}

func (r ApiPutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIdsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIdsExecute(r)
}

/*
PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIds Associate a list with a marketing event

Associates a list with a marketing event by external account id, external event id, and ILS list id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalAccountId The accountId that is associated with this marketing event in the external event application.
 @param externalEventId The id of the marketing event in the external event application.
 @param listId The ILS ID of the list.
 @return ApiPutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIdsRequest
*/
func (a *ListAssociationsAPIService) PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIds(ctx context.Context, externalAccountId string, externalEventId string, listId string) ApiPutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIdsRequest {
	return ApiPutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIdsRequest{
		ApiService: a,
		ctx: ctx,
		externalAccountId: externalAccountId,
		externalEventId: externalEventId,
		listId: listId,
	}
}

// Execute executes the request
func (a *ListAssociationsAPIService) PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIdsExecute(r ApiPutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIdsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListAssociationsAPIService.PutMarketingV3MarketingEventsAssociationsExternalAccountIdExternalEventIdListsListIdAssociateByExternalAccountAndEventIds")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/associations/{externalAccountId}/{externalEventId}/lists/{listId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalAccountId"+"}", url.PathEscape(parameterValueToString(r.externalAccountId, "externalAccountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"externalEventId"+"}", url.PathEscape(parameterValueToString(r.externalEventId, "externalEventId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

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

type ApiPutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventIdRequest struct {
	ctx context.Context
	ApiService *ListAssociationsAPIService
	marketingEventId string
	listId string
}

func (r ApiPutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventIdExecute(r)
}

/*
PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventId Associate a list with a marketing event

Associates a list with a marketing event by marketing event id and ILS list id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marketingEventId The internal id of the marketing event in HubSpot.
 @param listId The ILS ID of the list.
 @return ApiPutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventIdRequest
*/
func (a *ListAssociationsAPIService) PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventId(ctx context.Context, marketingEventId string, listId string) ApiPutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventIdRequest {
	return ApiPutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventIdRequest{
		ApiService: a,
		ctx: ctx,
		marketingEventId: marketingEventId,
		listId: listId,
	}
}

// Execute executes the request
func (a *ListAssociationsAPIService) PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventIdExecute(r ApiPutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListAssociationsAPIService.PutMarketingV3MarketingEventsAssociationsMarketingEventIdListsListIdAssociateByMarketingEventId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/associations/{marketingEventId}/lists/{listId}"
	localVarPath = strings.Replace(localVarPath, "{"+"marketingEventId"+"}", url.PathEscape(parameterValueToString(r.marketingEventId, "marketingEventId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

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
