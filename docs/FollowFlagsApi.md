# \FollowFlagsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFlagFollower**](FollowFlagsApi.md#DeleteFlagFollower) | **Delete** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/followers/{memberId} | Remove a member as a follower of a flag in a project and environment
[**GetFlagFollowers**](FollowFlagsApi.md#GetFlagFollowers) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/followers | Get followers of a flag in a project and environment
[**GetFollowersByProjEnv**](FollowFlagsApi.md#GetFollowersByProjEnv) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/followers | Get followers of all flags in a given project and environment
[**PutFlagFollower**](FollowFlagsApi.md#PutFlagFollower) | **Put** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/followers/{memberId} | Add a member as a follower of a flag in a project and environment



## DeleteFlagFollower

> DeleteFlagFollower(ctx, projectKey, featureFlagKey, environmentKey, memberId).Execute()

Remove a member as a follower of a flag in a project and environment



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
    projectKey := "projectKey_example" // string | The project key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key
    memberId := "memberId_example" // string | The memberId of the member to remove as a follower of the flag. Reader roles can only remove themselves.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FollowFlagsApi.DeleteFlagFollower(context.Background(), projectKey, featureFlagKey, environmentKey, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FollowFlagsApi.DeleteFlagFollower``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**memberId** | **string** | The memberId of the member to remove as a follower of the flag. Reader roles can only remove themselves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlagFollowerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagFollowers

> FlagFollowersGetRep GetFlagFollowers(ctx, projectKey, featureFlagKey, environmentKey).Execute()

Get followers of a flag in a project and environment



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
    projectKey := "projectKey_example" // string | The project key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FollowFlagsApi.GetFlagFollowers(context.Background(), projectKey, featureFlagKey, environmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FollowFlagsApi.GetFlagFollowers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlagFollowers`: FlagFollowersGetRep
    fmt.Fprintf(os.Stdout, "Response from `FollowFlagsApi.GetFlagFollowers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlagFollowersGetRep**](FlagFollowersGetRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowersByProjEnv

> FlagFollowersByProjEnvGetRep GetFollowersByProjEnv(ctx, projectKey, environmentKey).Execute()

Get followers of all flags in a given project and environment



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
    projectKey := "projectKey_example" // string | The project key
    environmentKey := "environmentKey_example" // string | The environment key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FollowFlagsApi.GetFollowersByProjEnv(context.Background(), projectKey, environmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FollowFlagsApi.GetFollowersByProjEnv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFollowersByProjEnv`: FlagFollowersByProjEnvGetRep
    fmt.Fprintf(os.Stdout, "Response from `FollowFlagsApi.GetFollowersByProjEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowersByProjEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlagFollowersByProjEnvGetRep**](FlagFollowersByProjEnvGetRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFlagFollower

> PutFlagFollower(ctx, projectKey, featureFlagKey, environmentKey, memberId).Execute()

Add a member as a follower of a flag in a project and environment



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
    projectKey := "projectKey_example" // string | The project key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key
    memberId := "memberId_example" // string | The memberId of the member to add as a follower of the flag. Reader roles can only add themselves.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FollowFlagsApi.PutFlagFollower(context.Background(), projectKey, featureFlagKey, environmentKey, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FollowFlagsApi.PutFlagFollower``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**memberId** | **string** | The memberId of the member to add as a follower of the flag. Reader roles can only add themselves. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFlagFollowerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

