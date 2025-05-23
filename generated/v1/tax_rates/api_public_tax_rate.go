/*
Settings Tax Rates

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tax_rates

import (
	"bytes"
	"context"
	"io"
	"net/http"
	
	"github.com/killean-johnson/go-hubspot"
"net/url"
	"strings"
)


// PublicTaxRateAPIService PublicTaxRateAPI service
type PublicTaxRateAPIService service

type ApiGetTaxRatesV1TaxRatesRequest struct {
	ctx context.Context
	ApiService *PublicTaxRateAPIService
	after *string
	limit *int32
	active *bool
}

// 
func (r ApiGetTaxRatesV1TaxRatesRequest) After(after string) ApiGetTaxRatesV1TaxRatesRequest {
	r.after = &after
	return r
}

// 
func (r ApiGetTaxRatesV1TaxRatesRequest) Limit(limit int32) ApiGetTaxRatesV1TaxRatesRequest {
	r.limit = &limit
	return r
}

// 
func (r ApiGetTaxRatesV1TaxRatesRequest) Active(active bool) ApiGetTaxRatesV1TaxRatesRequest {
	r.active = &active
	return r
}

func (r ApiGetTaxRatesV1TaxRatesRequest) Execute() (*CollectionResponsePublicTaxRateGroupForwardPaging, *http.Response, error) {
	return r.ApiService.GetTaxRatesV1TaxRatesExecute(r)
}

/*
GetTaxRatesV1TaxRates Method for GetTaxRatesV1TaxRates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTaxRatesV1TaxRatesRequest
*/
func (a *PublicTaxRateAPIService) GetTaxRatesV1TaxRates(ctx context.Context) ApiGetTaxRatesV1TaxRatesRequest {
	return ApiGetTaxRatesV1TaxRatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionResponsePublicTaxRateGroupForwardPaging
func (a *PublicTaxRateAPIService) GetTaxRatesV1TaxRatesExecute(r ApiGetTaxRatesV1TaxRatesRequest) (*CollectionResponsePublicTaxRateGroupForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponsePublicTaxRateGroupForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicTaxRateAPIService.GetTaxRatesV1TaxRates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax-rates/v1/tax-rates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.active != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "active", r.active, "form", "")
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

type ApiGetTaxRatesV1TaxRatesTaxRateGroupIdRequest struct {
	ctx context.Context
	ApiService *PublicTaxRateAPIService
	taxRateGroupId string
}

func (r ApiGetTaxRatesV1TaxRatesTaxRateGroupIdRequest) Execute() (*PublicTaxRateGroup, *http.Response, error) {
	return r.ApiService.GetTaxRatesV1TaxRatesTaxRateGroupIdExecute(r)
}

/*
GetTaxRatesV1TaxRatesTaxRateGroupId Method for GetTaxRatesV1TaxRatesTaxRateGroupId

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taxRateGroupId 
 @return ApiGetTaxRatesV1TaxRatesTaxRateGroupIdRequest
*/
func (a *PublicTaxRateAPIService) GetTaxRatesV1TaxRatesTaxRateGroupId(ctx context.Context, taxRateGroupId string) ApiGetTaxRatesV1TaxRatesTaxRateGroupIdRequest {
	return ApiGetTaxRatesV1TaxRatesTaxRateGroupIdRequest{
		ApiService: a,
		ctx: ctx,
		taxRateGroupId: taxRateGroupId,
	}
}

// Execute executes the request
//  @return PublicTaxRateGroup
func (a *PublicTaxRateAPIService) GetTaxRatesV1TaxRatesTaxRateGroupIdExecute(r ApiGetTaxRatesV1TaxRatesTaxRateGroupIdRequest) (*PublicTaxRateGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PublicTaxRateGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicTaxRateAPIService.GetTaxRatesV1TaxRatesTaxRateGroupId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax-rates/v1/tax-rates/{taxRateGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxRateGroupId"+"}", url.PathEscape(parameterValueToString(r.taxRateGroupId, "taxRateGroupId")), -1)

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
