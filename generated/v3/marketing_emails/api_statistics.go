/*
Marketing Emails

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_emails

import (
	"bytes"
	"context"
	"io"
	"net/http"
	
	"github.com/killean-johnson/go-hubspot"
"net/url"
	"reflect"
)


// StatisticsAPIService StatisticsAPI service
type StatisticsAPIService service

type ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest struct {
	ctx context.Context
	ApiService *StatisticsAPIService
	interval *string
	startTimestamp *string
	endTimestamp *string
	emailIds *[]int64
}

// The interval to aggregate statistics for.
func (r ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest) Interval(interval string) ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest {
	r.interval = &interval
	return r
}

// The start timestamp of the time span, in ISO8601 representation.
func (r ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest) StartTimestamp(startTimestamp string) ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest {
	r.startTimestamp = &startTimestamp
	return r
}

// The end timestamp of the time span, in ISO8601 representation.
func (r ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest) EndTimestamp(endTimestamp string) ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest {
	r.endTimestamp = &endTimestamp
	return r
}

// Filter by email IDs. Only include statistics of emails with these IDs.
func (r ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest) EmailIds(emailIds []int64) ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest {
	r.emailIds = &emailIds
	return r
}

func (r ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest) Execute() (*CollectionResponseWithTotalEmailStatisticIntervalNoPaging, *http.Response, error) {
	return r.ApiService.GetMarketingV3EmailsStatisticsHistogramGetHistogramExecute(r)
}

/*
GetMarketingV3EmailsStatisticsHistogramGetHistogram Get aggregated statistic intervals.

Get aggregated statistics in intervals for a specified time span. Each interval contains aggregated statistics of the emails that were sent in that time.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest
*/
func (a *StatisticsAPIService) GetMarketingV3EmailsStatisticsHistogramGetHistogram(ctx context.Context) ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest {
	return ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionResponseWithTotalEmailStatisticIntervalNoPaging
func (a *StatisticsAPIService) GetMarketingV3EmailsStatisticsHistogramGetHistogramExecute(r ApiGetMarketingV3EmailsStatisticsHistogramGetHistogramRequest) (*CollectionResponseWithTotalEmailStatisticIntervalNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponseWithTotalEmailStatisticIntervalNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsAPIService.GetMarketingV3EmailsStatisticsHistogramGetHistogram")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/emails/statistics/histogram"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	}
	if r.startTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTimestamp", r.startTimestamp, "form", "")
	}
	if r.endTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTimestamp", r.endTimestamp, "form", "")
	}
	if r.emailIds != nil {
		t := *r.emailIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "emailIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "emailIds", t, "form", "multi")
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

type ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest struct {
	ctx context.Context
	ApiService *StatisticsAPIService
	startTimestamp *string
	endTimestamp *string
	emailIds *[]int64
	property *string
}

// The start timestamp of the time span, in ISO8601 representation.
func (r ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest) StartTimestamp(startTimestamp string) ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest {
	r.startTimestamp = &startTimestamp
	return r
}

// The end timestamp of the time span, in ISO8601 representation.
func (r ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest) EndTimestamp(endTimestamp string) ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest {
	r.endTimestamp = &endTimestamp
	return r
}

// Filter by email IDs. Only include statistics of emails with these IDs.
func (r ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest) EmailIds(emailIds []int64) ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest {
	r.emailIds = &emailIds
	return r
}

// Specifies which email properties should be returned. All properties will be returned by default.
func (r ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest) Property(property string) ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest {
	r.property = &property
	return r
}

func (r ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest) Execute() (*AggregateEmailStatistics, *http.Response, error) {
	return r.ApiService.GetMarketingV3EmailsStatisticsListGetEmailsListExecute(r)
}

/*
GetMarketingV3EmailsStatisticsListGetEmailsList Get aggregated statistics.

Use this endpoint to get aggregated statistics of emails sent in a specified time span. It also returns the list of emails that were sent during the time span.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest
*/
func (a *StatisticsAPIService) GetMarketingV3EmailsStatisticsListGetEmailsList(ctx context.Context) ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest {
	return ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AggregateEmailStatistics
func (a *StatisticsAPIService) GetMarketingV3EmailsStatisticsListGetEmailsListExecute(r ApiGetMarketingV3EmailsStatisticsListGetEmailsListRequest) (*AggregateEmailStatistics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AggregateEmailStatistics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsAPIService.GetMarketingV3EmailsStatisticsListGetEmailsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/emails/statistics/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTimestamp", r.startTimestamp, "form", "")
	}
	if r.endTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTimestamp", r.endTimestamp, "form", "")
	}
	if r.emailIds != nil {
		t := *r.emailIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "emailIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "emailIds", t, "form", "multi")
		}
	}
	if r.property != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "property", r.property, "form", "")
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
