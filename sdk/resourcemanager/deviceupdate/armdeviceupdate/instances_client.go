//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceupdate

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// InstancesClient contains the methods for the Instances group.
// Don't use this type directly, use NewInstancesClient() instead.
type InstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInstancesClient creates a new instance of InstancesClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates or updates instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The resource group name.
//   - accountName - Account name.
//   - instanceName - Instance name.
//   - instance - Instance details.
//   - options - InstancesClientBeginCreateOptions contains the optional parameters for the InstancesClient.BeginCreate method.
func (client *InstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, instanceName string, instance Instance, options *InstancesClientBeginCreateOptions) (*runtime.Poller[InstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, instanceName, instance, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InstancesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InstancesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates or updates instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
func (client *InstancesClient) create(ctx context.Context, resourceGroupName string, accountName string, instanceName string, instance Instance, options *InstancesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "InstancesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, instanceName, instance, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *InstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, instanceName string, instance Instance, options *InstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, instance); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The resource group name.
//   - accountName - Account name.
//   - instanceName - Instance name.
//   - options - InstancesClientBeginDeleteOptions contains the optional parameters for the InstancesClient.BeginDelete method.
func (client *InstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientBeginDeleteOptions) (*runtime.Poller[InstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, instanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InstancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InstancesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
func (client *InstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "InstancesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, instanceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *InstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns instance details for the given instance and account name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The resource group name.
//   - accountName - Account name.
//   - instanceName - Instance name.
//   - options - InstancesClientGetOptions contains the optional parameters for the InstancesClient.Get method.
func (client *InstancesClient) Get(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientGetOptions) (InstancesClientGetResponse, error) {
	var err error
	const operationName = "InstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, instanceName, options)
	if err != nil {
		return InstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *InstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InstancesClient) getHandleResponse(resp *http.Response) (InstancesClientGetResponse, error) {
	result := InstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Instance); err != nil {
		return InstancesClientGetResponse{}, err
	}
	return result, nil
}

// Head - Checks whether instance exists.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The resource group name.
//   - accountName - Account name.
//   - instanceName - Instance name.
//   - options - InstancesClientHeadOptions contains the optional parameters for the InstancesClient.Head method.
func (client *InstancesClient) Head(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientHeadOptions) (InstancesClientHeadResponse, error) {
	var err error
	const operationName = "InstancesClient.Head"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.headCreateRequest(ctx, resourceGroupName, accountName, instanceName, options)
	if err != nil {
		return InstancesClientHeadResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InstancesClientHeadResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InstancesClientHeadResponse{}, err
	}
	return InstancesClientHeadResponse{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// headCreateRequest creates the Head request.
func (client *InstancesClient) headCreateRequest(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientHeadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByAccountPager - Returns instances for the given account name.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The resource group name.
//   - accountName - Account name.
//   - options - InstancesClientListByAccountOptions contains the optional parameters for the InstancesClient.NewListByAccountPager
//     method.
func (client *InstancesClient) NewListByAccountPager(resourceGroupName string, accountName string, options *InstancesClientListByAccountOptions) *runtime.Pager[InstancesClientListByAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[InstancesClientListByAccountResponse]{
		More: func(page InstancesClientListByAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InstancesClientListByAccountResponse) (InstancesClientListByAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "InstancesClient.NewListByAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return InstancesClientListByAccountResponse{}, err
			}
			return client.listByAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *InstancesClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *InstancesClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *InstancesClient) listByAccountHandleResponse(resp *http.Response) (InstancesClientListByAccountResponse, error) {
	result := InstancesClientListByAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InstanceList); err != nil {
		return InstancesClientListByAccountResponse{}, err
	}
	return result, nil
}

// Update - Updates instance's tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The resource group name.
//   - accountName - Account name.
//   - instanceName - Instance name.
//   - tagUpdatePayload - Updated tags.
//   - options - InstancesClientUpdateOptions contains the optional parameters for the InstancesClient.Update method.
func (client *InstancesClient) Update(ctx context.Context, resourceGroupName string, accountName string, instanceName string, tagUpdatePayload TagUpdate, options *InstancesClientUpdateOptions) (InstancesClientUpdateResponse, error) {
	var err error
	const operationName = "InstancesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, instanceName, tagUpdatePayload, options)
	if err != nil {
		return InstancesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InstancesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *InstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, instanceName string, tagUpdatePayload TagUpdate, options *InstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, tagUpdatePayload); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *InstancesClient) updateHandleResponse(resp *http.Response) (InstancesClientUpdateResponse, error) {
	result := InstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Instance); err != nil {
		return InstancesClientUpdateResponse{}, err
	}
	return result, nil
}
