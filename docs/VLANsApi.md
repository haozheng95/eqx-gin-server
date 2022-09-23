# \VLANsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVirtualNetwork**](VLANsApi.md#CreateVirtualNetwork) | **Post** /projects/{id}/virtual-networks | Create a virtual network
[**DeleteVirtualNetwork**](VLANsApi.md#DeleteVirtualNetwork) | **Delete** /virtual-networks/{id} | Delete a virtual network
[**FindVirtualNetworks**](VLANsApi.md#FindVirtualNetworks) | **Get** /projects/{id}/virtual-networks | Retrieve all virtual networks
[**GetVirtualNetwork**](VLANsApi.md#GetVirtualNetwork) | **Get** /virtual-networks/{id} | Get a virtual network



## CreateVirtualNetwork

> FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork CreateVirtualNetwork(ctx, id).CreateVirtualNetworkRequest(createVirtualNetworkRequest).Execute()

Create a virtual network



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    createVirtualNetworkRequest := *openapiclient.NewCreateVirtualNetworkRequest("ProjectId_example") // CreateVirtualNetworkRequest | Virtual Network to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VLANsApi.CreateVirtualNetwork(context.Background(), id).CreateVirtualNetworkRequest(createVirtualNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.CreateVirtualNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVirtualNetwork`: FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.CreateVirtualNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVirtualNetworkRequest** | [**CreateVirtualNetworkRequest**](CreateVirtualNetworkRequest.md) | Virtual Network to create | 

### Return type

[**FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork**](FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualNetwork

> FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork DeleteVirtualNetwork(ctx, id).Execute()

Delete a virtual network



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Virtual Network UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VLANsApi.DeleteVirtualNetwork(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.DeleteVirtualNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualNetwork`: FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.DeleteVirtualNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Virtual Network UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork**](FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindVirtualNetworks

> FindVirtualNetworks200Response FindVirtualNetworks(ctx, id).Include(include).Exclude(exclude).Facility(facility).Metro(metro).Execute()

Retrieve all virtual networks



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project UUID
    include := []string{"Inner_example"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
    exclude := []string{"Inner_example"} // []string | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. (optional)
    facility := "facility_example" // string | Filter by Facility ID (uuid) or Facility Code (optional)
    metro := "metro_example" // string | Filter by Metro ID (uuid) or Metro Code (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VLANsApi.FindVirtualNetworks(context.Background(), id).Include(include).Exclude(exclude).Facility(facility).Metro(metro).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.FindVirtualNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindVirtualNetworks`: FindVirtualNetworks200Response
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.FindVirtualNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindVirtualNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. | 
 **exclude** | **[]string** | Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects. | 
 **facility** | **string** | Filter by Facility ID (uuid) or Facility Code | 
 **metro** | **string** | Filter by Metro ID (uuid) or Metro Code | 

### Return type

[**FindVirtualNetworks200Response**](FindVirtualNetworks200Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualNetwork

> FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork GetVirtualNetwork(ctx, id).Execute()

Get a virtual network



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Virtual Network UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VLANsApi.GetVirtualNetwork(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.GetVirtualNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualNetwork`: FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork
    fmt.Fprintf(os.Stdout, "Response from `VLANsApi.GetVirtualNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Virtual Network UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork**](FindDeviceById200ResponseNetworkPortsInnerNativeVirtualNetwork.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

