/*
Discounts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discounts

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

type ApiDeleteCrmV3ObjectsDiscountsDiscountIdRequest struct {
	ctx context.Context
	ApiService *BasicAPIService
	discountId string
}

func (r ApiDeleteCrmV3ObjectsDiscountsDiscountIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrmV3ObjectsDiscountsDiscountIdExecute(r)
}

/*
DeleteCrmV3ObjectsDiscountsDiscountId Archive

Move an Object identified by `{discountId}` to the recycling bin.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param discountId
 @return ApiDeleteCrmV3ObjectsDiscountsDiscountIdRequest
*/
func (a *BasicAPIService) DeleteCrmV3ObjectsDiscountsDiscountId(ctx context.Context, discountId string) ApiDeleteCrmV3ObjectsDiscountsDiscountIdRequest {
	return ApiDeleteCrmV3ObjectsDiscountsDiscountIdRequest{
		ApiService: a,
		ctx: ctx,
		discountId: discountId,
	}
}

// Execute executes the request
func (a *BasicAPIService) DeleteCrmV3ObjectsDiscountsDiscountIdExecute(r ApiDeleteCrmV3ObjectsDiscountsDiscountIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicAPIService.DeleteCrmV3ObjectsDiscountsDiscountId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/discounts/{discountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"discountId"+"}", url.PathEscape(parameterValueToString(r.discountId, "discountId")), -1)

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

type ApiGetCrmV3ObjectsDiscountsRequest struct {
	ctx context.Context
	ApiService *BasicAPIService
	limit *int32
	after *string
	properties *[]string
	propertiesWithHistory *[]string
	associations *[]string
	archived *bool
}

// The maximum number of results to display per page.
func (r ApiGetCrmV3ObjectsDiscountsRequest) Limit(limit int32) ApiGetCrmV3ObjectsDiscountsRequest {
	r.limit = &limit
	return r
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiGetCrmV3ObjectsDiscountsRequest) After(after string) ApiGetCrmV3ObjectsDiscountsRequest {
	r.after = &after
	return r
}

// A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetCrmV3ObjectsDiscountsRequest) Properties(properties []string) ApiGetCrmV3ObjectsDiscountsRequest {
	r.properties = &properties
	return r
}

// A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored. Usage of this parameter will reduce the maximum number of objects that can be read by a single request.
func (r ApiGetCrmV3ObjectsDiscountsRequest) PropertiesWithHistory(propertiesWithHistory []string) ApiGetCrmV3ObjectsDiscountsRequest {
	r.propertiesWithHistory = &propertiesWithHistory
	return r
}

// A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored.
func (r ApiGetCrmV3ObjectsDiscountsRequest) Associations(associations []string) ApiGetCrmV3ObjectsDiscountsRequest {
	r.associations = &associations
	return r
}

// Whether to return only results that have been archived.
func (r ApiGetCrmV3ObjectsDiscountsRequest) Archived(archived bool) ApiGetCrmV3ObjectsDiscountsRequest {
	r.archived = &archived
	return r
}

func (r ApiGetCrmV3ObjectsDiscountsRequest) Execute() (*CollectionResponseSimplePublicObjectWithAssociationsForwardPaging, *http.Response, error) {
	return r.ApiService.GetCrmV3ObjectsDiscountsExecute(r)
}

/*
GetCrmV3ObjectsDiscounts List

Read a page of discounts. Control what is returned via the `properties` query param.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCrmV3ObjectsDiscountsRequest
*/
func (a *BasicAPIService) GetCrmV3ObjectsDiscounts(ctx context.Context) ApiGetCrmV3ObjectsDiscountsRequest {
	return ApiGetCrmV3ObjectsDiscountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
func (a *BasicAPIService) GetCrmV3ObjectsDiscountsExecute(r ApiGetCrmV3ObjectsDiscountsRequest) (*CollectionResponseSimplePublicObjectWithAssociationsForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicAPIService.GetCrmV3ObjectsDiscounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/discounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 10
		r.limit = &defaultValue
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "form", "")
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
	if r.propertiesWithHistory != nil {
		t := *r.propertiesWithHistory
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "propertiesWithHistory", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "propertiesWithHistory", t, "form", "multi")
		}
	}
	if r.associations != nil {
		t := *r.associations
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "associations", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "associations", t, "form", "multi")
		}
	}
	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
	} else {
		var defaultValue bool = false
		r.archived = &defaultValue
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

type ApiGetCrmV3ObjectsDiscountsDiscountIdRequest struct {
	ctx context.Context
	ApiService *BasicAPIService
	discountId string
	properties *[]string
	propertiesWithHistory *[]string
	associations *[]string
	archived *bool
	idProperty *string
}

// A comma separated list of the properties to be returned in the response. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetCrmV3ObjectsDiscountsDiscountIdRequest) Properties(properties []string) ApiGetCrmV3ObjectsDiscountsDiscountIdRequest {
	r.properties = &properties
	return r
}

// A comma separated list of the properties to be returned along with their history of previous values. If any of the specified properties are not present on the requested object(s), they will be ignored.
func (r ApiGetCrmV3ObjectsDiscountsDiscountIdRequest) PropertiesWithHistory(propertiesWithHistory []string) ApiGetCrmV3ObjectsDiscountsDiscountIdRequest {
	r.propertiesWithHistory = &propertiesWithHistory
	return r
}

// A comma separated list of object types to retrieve associated IDs for. If any of the specified associations do not exist, they will be ignored.
func (r ApiGetCrmV3ObjectsDiscountsDiscountIdRequest) Associations(associations []string) ApiGetCrmV3ObjectsDiscountsDiscountIdRequest {
	r.associations = &associations
	return r
}

// Whether to return only results that have been archived.
func (r ApiGetCrmV3ObjectsDiscountsDiscountIdRequest) Archived(archived bool) ApiGetCrmV3ObjectsDiscountsDiscountIdRequest {
	r.archived = &archived
	return r
}

// The name of a property whose values are unique for this object
func (r ApiGetCrmV3ObjectsDiscountsDiscountIdRequest) IdProperty(idProperty string) ApiGetCrmV3ObjectsDiscountsDiscountIdRequest {
	r.idProperty = &idProperty
	return r
}

func (r ApiGetCrmV3ObjectsDiscountsDiscountIdRequest) Execute() (*SimplePublicObjectWithAssociations, *http.Response, error) {
	return r.ApiService.GetCrmV3ObjectsDiscountsDiscountIdExecute(r)
}

/*
GetCrmV3ObjectsDiscountsDiscountId Read

Read an Object identified by `{discountId}`. `{discountId}` refers to the internal object ID by default, or optionally any unique property value as specified by the `idProperty` query param.  Control what is returned via the `properties` query param.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param discountId
 @return ApiGetCrmV3ObjectsDiscountsDiscountIdRequest
*/
func (a *BasicAPIService) GetCrmV3ObjectsDiscountsDiscountId(ctx context.Context, discountId string) ApiGetCrmV3ObjectsDiscountsDiscountIdRequest {
	return ApiGetCrmV3ObjectsDiscountsDiscountIdRequest{
		ApiService: a,
		ctx: ctx,
		discountId: discountId,
	}
}

// Execute executes the request
//  @return SimplePublicObjectWithAssociations
func (a *BasicAPIService) GetCrmV3ObjectsDiscountsDiscountIdExecute(r ApiGetCrmV3ObjectsDiscountsDiscountIdRequest) (*SimplePublicObjectWithAssociations, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SimplePublicObjectWithAssociations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicAPIService.GetCrmV3ObjectsDiscountsDiscountId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/discounts/{discountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"discountId"+"}", url.PathEscape(parameterValueToString(r.discountId, "discountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.propertiesWithHistory != nil {
		t := *r.propertiesWithHistory
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "propertiesWithHistory", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "propertiesWithHistory", t, "form", "multi")
		}
	}
	if r.associations != nil {
		t := *r.associations
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "associations", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "associations", t, "form", "multi")
		}
	}
	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
	} else {
		var defaultValue bool = false
		r.archived = &defaultValue
	}
	if r.idProperty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "idProperty", r.idProperty, "form", "")
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

type ApiPatchCrmV3ObjectsDiscountsDiscountIdRequest struct {
	ctx context.Context
	ApiService *BasicAPIService
	discountId string
	simplePublicObjectInput *SimplePublicObjectInput
	idProperty *string
}

func (r ApiPatchCrmV3ObjectsDiscountsDiscountIdRequest) SimplePublicObjectInput(simplePublicObjectInput SimplePublicObjectInput) ApiPatchCrmV3ObjectsDiscountsDiscountIdRequest {
	r.simplePublicObjectInput = &simplePublicObjectInput
	return r
}

// The name of a property whose values are unique for this object
func (r ApiPatchCrmV3ObjectsDiscountsDiscountIdRequest) IdProperty(idProperty string) ApiPatchCrmV3ObjectsDiscountsDiscountIdRequest {
	r.idProperty = &idProperty
	return r
}

func (r ApiPatchCrmV3ObjectsDiscountsDiscountIdRequest) Execute() (*SimplePublicObject, *http.Response, error) {
	return r.ApiService.PatchCrmV3ObjectsDiscountsDiscountIdExecute(r)
}

/*
PatchCrmV3ObjectsDiscountsDiscountId Update

Perform a partial update of an Object identified by `{discountId}`or optionally a unique property value as specified by the `idProperty` query param. `{discountId}` refers to the internal object ID by default, and the `idProperty` query param refers to a property whose values are unique for the object. Provided property values will be overwritten. Read-only and non-existent properties will result in an error. Properties values can be cleared by passing an empty string.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param discountId
 @return ApiPatchCrmV3ObjectsDiscountsDiscountIdRequest
*/
func (a *BasicAPIService) PatchCrmV3ObjectsDiscountsDiscountId(ctx context.Context, discountId string) ApiPatchCrmV3ObjectsDiscountsDiscountIdRequest {
	return ApiPatchCrmV3ObjectsDiscountsDiscountIdRequest{
		ApiService: a,
		ctx: ctx,
		discountId: discountId,
	}
}

// Execute executes the request
//  @return SimplePublicObject
func (a *BasicAPIService) PatchCrmV3ObjectsDiscountsDiscountIdExecute(r ApiPatchCrmV3ObjectsDiscountsDiscountIdRequest) (*SimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicAPIService.PatchCrmV3ObjectsDiscountsDiscountId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/discounts/{discountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"discountId"+"}", url.PathEscape(parameterValueToString(r.discountId, "discountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.simplePublicObjectInput == nil {
		return localVarReturnValue, nil, reportError("simplePublicObjectInput is required and must be specified")
	}

	if r.idProperty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "idProperty", r.idProperty, "form", "")
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
	localVarPostBody = r.simplePublicObjectInput
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

type ApiPostCrmV3ObjectsDiscountsRequest struct {
	ctx context.Context
	ApiService *BasicAPIService
	simplePublicObjectInputForCreate *SimplePublicObjectInputForCreate
}

func (r ApiPostCrmV3ObjectsDiscountsRequest) SimplePublicObjectInputForCreate(simplePublicObjectInputForCreate SimplePublicObjectInputForCreate) ApiPostCrmV3ObjectsDiscountsRequest {
	r.simplePublicObjectInputForCreate = &simplePublicObjectInputForCreate
	return r
}

func (r ApiPostCrmV3ObjectsDiscountsRequest) Execute() (*SimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsDiscountsExecute(r)
}

/*
PostCrmV3ObjectsDiscounts Create

Create a discount with the given properties and return a copy of the object, including the ID. Documentation and examples for creating standard discounts is provided.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsDiscountsRequest
*/
func (a *BasicAPIService) PostCrmV3ObjectsDiscounts(ctx context.Context) ApiPostCrmV3ObjectsDiscountsRequest {
	return ApiPostCrmV3ObjectsDiscountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SimplePublicObject
func (a *BasicAPIService) PostCrmV3ObjectsDiscountsExecute(r ApiPostCrmV3ObjectsDiscountsRequest) (*SimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicAPIService.PostCrmV3ObjectsDiscounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/discounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.simplePublicObjectInputForCreate == nil {
		return localVarReturnValue, nil, reportError("simplePublicObjectInputForCreate is required and must be specified")
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
	localVarPostBody = r.simplePublicObjectInputForCreate
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
