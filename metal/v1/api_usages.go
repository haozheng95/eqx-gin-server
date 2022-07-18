/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// UsagesApiService UsagesApi service
type UsagesApiService service

type ApiFindDeviceUsagesRequest struct {
	ctx           context.Context
	ApiService    *UsagesApiService
	id            string
	createdAfter  *string
	createdBefore *string
}

// Filter usages created after this date
func (r ApiFindDeviceUsagesRequest) CreatedAfter(createdAfter string) ApiFindDeviceUsagesRequest {
	r.createdAfter = &createdAfter
	return r
}

// Filter usages created before this date
func (r ApiFindDeviceUsagesRequest) CreatedBefore(createdBefore string) ApiFindDeviceUsagesRequest {
	r.createdBefore = &createdBefore
	return r
}

func (r ApiFindDeviceUsagesRequest) Execute() (*FindDeviceUsages200Response, *http.Response, error) {
	return r.ApiService.FindDeviceUsagesExecute(r)
}

/*
FindDeviceUsages Retrieve all usages for device

Returns all usages for a device.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Device UUID
 @return ApiFindDeviceUsagesRequest
*/
func (a *UsagesApiService) FindDeviceUsages(ctx context.Context, id string) ApiFindDeviceUsagesRequest {
	return ApiFindDeviceUsagesRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return FindDeviceUsages200Response
func (a *UsagesApiService) FindDeviceUsagesExecute(r ApiFindDeviceUsagesRequest) (*FindDeviceUsages200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FindDeviceUsages200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsagesApiService.FindDeviceUsages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{id}/usages"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createdAfter != nil {
		localVarQueryParams.Add("created[after]", parameterToString(*r.createdAfter, ""))
	}
	if r.createdBefore != nil {
		localVarQueryParams.Add("created[before]", parameterToString(*r.createdBefore, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v DeleteAPIKey401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v DeleteAPIKey401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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

type ApiFindProjectUsageRequest struct {
	ctx           context.Context
	ApiService    *UsagesApiService
	id            string
	createdAfter  *string
	createdBefore *string
}

// Filter usages created after this date
func (r ApiFindProjectUsageRequest) CreatedAfter(createdAfter string) ApiFindProjectUsageRequest {
	r.createdAfter = &createdAfter
	return r
}

// Filter usages created before this date
func (r ApiFindProjectUsageRequest) CreatedBefore(createdBefore string) ApiFindProjectUsageRequest {
	r.createdBefore = &createdBefore
	return r
}

func (r ApiFindProjectUsageRequest) Execute() (*FindProjectUsage200Response, *http.Response, error) {
	return r.ApiService.FindProjectUsageExecute(r)
}

/*
FindProjectUsage Retrieve all usages for project

Returns all usages for a project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Project UUID
 @return ApiFindProjectUsageRequest
*/
func (a *UsagesApiService) FindProjectUsage(ctx context.Context, id string) ApiFindProjectUsageRequest {
	return ApiFindProjectUsageRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return FindProjectUsage200Response
func (a *UsagesApiService) FindProjectUsageExecute(r ApiFindProjectUsageRequest) (*FindProjectUsage200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FindProjectUsage200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsagesApiService.FindProjectUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{id}/usages"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createdAfter != nil {
		localVarQueryParams.Add("created[after]", parameterToString(*r.createdAfter, ""))
	}
	if r.createdBefore != nil {
		localVarQueryParams.Add("created[before]", parameterToString(*r.createdBefore, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v DeleteAPIKey401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v DeleteAPIKey401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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