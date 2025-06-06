/*
Campaigns Public Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package campaigns_public_api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	
	"github.com/killean-johnson/go-hubspot"
"net/url"
	"strings"
	"reflect"
)


// BasicAPIService BasicAPI service
type BasicAPIService service

type ApiDeleteMarketingV3CampaignsCampaignGuidRequest struct {
	ctx context.Context
	ApiService *BasicAPIService
	campaignGuid string
}

func (r ApiDeleteMarketingV3CampaignsCampaignGuidRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMarketingV3CampaignsCampaignGuidExecute(r)
}

/*
DeleteMarketingV3CampaignsCampaignGuid Delete campaign 

Delete a specified campaign from the system.
This call will return a 204 No Content response regardless of whether the campaignGuid provided corresponds to an existing campaign or not.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaignGuid Unique identifier for the campaign, formatted as a UUID.
 @return ApiDeleteMarketingV3CampaignsCampaignGuidRequest
*/
func (a *BasicAPIService) DeleteMarketingV3CampaignsCampaignGuid(ctx context.Context, campaignGuid string) ApiDeleteMarketingV3CampaignsCampaignGuidRequest {
	return ApiDeleteMarketingV3CampaignsCampaignGuidRequest{
		ApiService: a,
		ctx: ctx,
		campaignGuid: campaignGuid,
	}
}

// Execute executes the request
func (a *BasicAPIService) DeleteMarketingV3CampaignsCampaignGuidExecute(r ApiDeleteMarketingV3CampaignsCampaignGuidRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicAPIService.DeleteMarketingV3CampaignsCampaignGuid")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/campaigns/{campaignGuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignGuid"+"}", url.PathEscape(parameterValueToString(r.campaignGuid, "campaignGuid")), -1)

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

type ApiGetMarketingV3CampaignsCampaignGuidRequest struct {
	ctx context.Context
	ApiService *BasicAPIService
	campaignGuid string
	startDate *string
	endDate *string
	properties *[]string
}

// Start date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched.
func (r ApiGetMarketingV3CampaignsCampaignGuidRequest) StartDate(startDate string) ApiGetMarketingV3CampaignsCampaignGuidRequest {
	r.startDate = &startDate
	return r
}

//  End date to fetch asset metrics, formatted as YYYY-MM-DD. This date is used to fetch the metrics associated with the assets for a specified period. If not provided, no asset metrics will be fetched.
func (r ApiGetMarketingV3CampaignsCampaignGuidRequest) EndDate(endDate string) ApiGetMarketingV3CampaignsCampaignGuidRequest {
	r.endDate = &endDate
	return r
}

// A comma-separated list of the properties to be returned in the response. If any of the specified properties has empty value on the requested object, they will be ignored and not returned in response. If this parameter is empty, the response will include an empty properties map.
func (r ApiGetMarketingV3CampaignsCampaignGuidRequest) Properties(properties []string) ApiGetMarketingV3CampaignsCampaignGuidRequest {
	r.properties = &properties
	return r
}

func (r ApiGetMarketingV3CampaignsCampaignGuidRequest) Execute() (*PublicCampaignWithAssets, *http.Response, error) {
	return r.ApiService.GetMarketingV3CampaignsCampaignGuidExecute(r)
}

/*
GetMarketingV3CampaignsCampaignGuid Read a campaign

Get a campaign identified by a specific campaignGuid with the given properties. Along with the campaign information, it also returns information about assets. Depending on the query parameters used, this can also be used to return information about the corresponding assets' metrics. Metrics are available only if startDate and endDate are provided.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaignGuid Unique identifier for the campaign, formatted as a UUID.
 @return ApiGetMarketingV3CampaignsCampaignGuidRequest
*/
func (a *BasicAPIService) GetMarketingV3CampaignsCampaignGuid(ctx context.Context, campaignGuid string) ApiGetMarketingV3CampaignsCampaignGuidRequest {
	return ApiGetMarketingV3CampaignsCampaignGuidRequest{
		ApiService: a,
		ctx: ctx,
		campaignGuid: campaignGuid,
	}
}

// Execute executes the request
//  @return PublicCampaignWithAssets
func (a *BasicAPIService) GetMarketingV3CampaignsCampaignGuidExecute(r ApiGetMarketingV3CampaignsCampaignGuidRequest) (*PublicCampaignWithAssets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PublicCampaignWithAssets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicAPIService.GetMarketingV3CampaignsCampaignGuid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/campaigns/{campaignGuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignGuid"+"}", url.PathEscape(parameterValueToString(r.campaignGuid, "campaignGuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	}
	if r.properties != nil {
		t := *r.properties
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "properties", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "properties", t, "form", "multi")
		}
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

type ApiPatchMarketingV3CampaignsCampaignGuidRequest struct {
	ctx context.Context
	ApiService *BasicAPIService
	campaignGuid string
	publicCampaignInput *PublicCampaignInput
}

func (r ApiPatchMarketingV3CampaignsCampaignGuidRequest) PublicCampaignInput(publicCampaignInput PublicCampaignInput) ApiPatchMarketingV3CampaignsCampaignGuidRequest {
	r.publicCampaignInput = &publicCampaignInput
	return r
}

func (r ApiPatchMarketingV3CampaignsCampaignGuidRequest) Execute() (*PublicCampaign, *http.Response, error) {
	return r.ApiService.PatchMarketingV3CampaignsCampaignGuidExecute(r)
}

/*
PatchMarketingV3CampaignsCampaignGuid Update campaign

Perform a partial update of a campaign identified by the specified campaignGuid. Provided property values will be overwritten. Read-only and non-existent properties will cause 400 error.
If an empty string is passed for any property in the Batch Update, it will reset that property's value.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaignGuid Unique identifier for the campaign, formatted as a UUID. 
 @return ApiPatchMarketingV3CampaignsCampaignGuidRequest
*/
func (a *BasicAPIService) PatchMarketingV3CampaignsCampaignGuid(ctx context.Context, campaignGuid string) ApiPatchMarketingV3CampaignsCampaignGuidRequest {
	return ApiPatchMarketingV3CampaignsCampaignGuidRequest{
		ApiService: a,
		ctx: ctx,
		campaignGuid: campaignGuid,
	}
}

// Execute executes the request
//  @return PublicCampaign
func (a *BasicAPIService) PatchMarketingV3CampaignsCampaignGuidExecute(r ApiPatchMarketingV3CampaignsCampaignGuidRequest) (*PublicCampaign, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PublicCampaign
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicAPIService.PatchMarketingV3CampaignsCampaignGuid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/campaigns/{campaignGuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignGuid"+"}", url.PathEscape(parameterValueToString(r.campaignGuid, "campaignGuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicCampaignInput == nil {
		return localVarReturnValue, nil, reportError("publicCampaignInput is required and must be specified")
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
	localVarPostBody = r.publicCampaignInput
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

type ApiPostMarketingV3CampaignsRequest struct {
	ctx context.Context
	ApiService *BasicAPIService
	publicCampaignInput *PublicCampaignInput
}

func (r ApiPostMarketingV3CampaignsRequest) PublicCampaignInput(publicCampaignInput PublicCampaignInput) ApiPostMarketingV3CampaignsRequest {
	r.publicCampaignInput = &publicCampaignInput
	return r
}

func (r ApiPostMarketingV3CampaignsRequest) Execute() (*PublicCampaign, *http.Response, error) {
	return r.ApiService.PostMarketingV3CampaignsExecute(r)
}

/*
PostMarketingV3Campaigns Create a campaign

Create a campaign with the given properties and return the campaign object, including the campaignGuid and created properties. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostMarketingV3CampaignsRequest
*/
func (a *BasicAPIService) PostMarketingV3Campaigns(ctx context.Context) ApiPostMarketingV3CampaignsRequest {
	return ApiPostMarketingV3CampaignsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PublicCampaign
func (a *BasicAPIService) PostMarketingV3CampaignsExecute(r ApiPostMarketingV3CampaignsRequest) (*PublicCampaign, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PublicCampaign
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicAPIService.PostMarketingV3Campaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/campaigns/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicCampaignInput == nil {
		return localVarReturnValue, nil, reportError("publicCampaignInput is required and must be specified")
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
	localVarPostBody = r.publicCampaignInput
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
