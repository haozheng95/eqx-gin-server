# \MetalGatewaysApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetalGateway**](MetalGatewaysApi.md#CreateMetalGateway) | **Post** /projects/{project_id}/metal-gateways | Create a metal gateway
[**DeleteMetalGateway**](MetalGatewaysApi.md#DeleteMetalGateway) | **Delete** /metal-gateways/{id} | Deletes the metal gateway
[**FindMetalGatewayById**](MetalGatewaysApi.md#FindMetalGatewayById) | **Get** /metal-gateways/{id} | Returns the metal gateway
[**FindMetalGatewaysByProject**](MetalGatewaysApi.md#FindMetalGatewaysByProject) | **Get** /projects/{project_id}/metal-gateways | Returns all metal gateways for a project



## CreateMetalGateway

> MetalGateway CreateMetalGateway(ctx, projectId).MetalGateway(metalGateway).Page(page).PerPage(perPage).Execute()

Create a metal gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := TODO // string | Project UUID
    metalGateway := *openapiclient.NewMetalGatewayInput("VirtualNetworkId_example") // MetalGatewayInput | Metal Gateway to create
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetalGatewaysApi.CreateMetalGateway(context.Background(), projectId).MetalGateway(metalGateway).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.CreateMetalGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMetalGateway`: MetalGateway
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.CreateMetalGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetalGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metalGateway** | [**MetalGatewayInput**](MetalGatewayInput.md) | Metal Gateway to create | 
 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**MetalGateway**](MetalGateway.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetalGateway

> DeleteMetalGateway(ctx, id).Execute()

Deletes the metal gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Metal Gateway UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetalGatewaysApi.DeleteMetalGateway(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.DeleteMetalGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Metal Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetalGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindMetalGatewayById

> MetalGateway FindMetalGatewayById(ctx, id).Execute()

Returns the metal gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := TODO // string | Metal Gateway UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetalGatewaysApi.FindMetalGatewayById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.FindMetalGatewayById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindMetalGatewayById`: MetalGateway
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.FindMetalGatewayById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Metal Gateway UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindMetalGatewayByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetalGateway**](MetalGateway.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindMetalGatewaysByProject

> MetalGatewayList FindMetalGatewaysByProject(ctx, projectId).Page(page).PerPage(perPage).Execute()

Returns all metal gateways for a project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := TODO // string | Project UUID
    page := int32(56) // int32 | Page to return (optional) (default to 1)
    perPage := int32(56) // int32 | Items returned per page (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetalGatewaysApi.FindMetalGatewaysByProject(context.Background(), projectId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.FindMetalGatewaysByProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindMetalGatewaysByProject`: MetalGatewayList
    fmt.Fprintf(os.Stdout, "Response from `MetalGatewaysApi.FindMetalGatewaysByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindMetalGatewaysByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page to return | [default to 1]
 **perPage** | **int32** | Items returned per page | [default to 10]

### Return type

[**MetalGatewayList**](MetalGatewayList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
