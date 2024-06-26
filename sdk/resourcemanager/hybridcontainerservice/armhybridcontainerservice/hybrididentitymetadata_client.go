//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcontainerservice

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// HybridIdentityMetadataClient contains the methods for the HybridIdentityMetadata group.
// Don't use this type directly, use NewHybridIdentityMetadataClient() instead.
type HybridIdentityMetadataClient struct {
	internal *arm.Client
}

// NewHybridIdentityMetadataClient creates a new instance of HybridIdentityMetadataClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHybridIdentityMetadataClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*HybridIdentityMetadataClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HybridIdentityMetadataClient{
		internal: cl,
	}
	return client, nil
}

// BeginDelete - Deletes the hybrid identity metadata proxy resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - HybridIdentityMetadataClientBeginDeleteOptions contains the optional parameters for the HybridIdentityMetadataClient.BeginDelete
//     method.
func (client *HybridIdentityMetadataClient) BeginDelete(ctx context.Context, connectedClusterResourceURI string, options *HybridIdentityMetadataClientBeginDeleteOptions) (*runtime.Poller[HybridIdentityMetadataClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, connectedClusterResourceURI, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HybridIdentityMetadataClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HybridIdentityMetadataClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the hybrid identity metadata proxy resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
func (client *HybridIdentityMetadataClient) deleteOperation(ctx context.Context, connectedClusterResourceURI string, options *HybridIdentityMetadataClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "HybridIdentityMetadataClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, connectedClusterResourceURI, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HybridIdentityMetadataClient) deleteCreateRequest(ctx context.Context, connectedClusterResourceURI string, options *HybridIdentityMetadataClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/hybridIdentityMetadata/default"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the hybrid identity metadata proxy resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - HybridIdentityMetadataClientGetOptions contains the optional parameters for the HybridIdentityMetadataClient.Get
//     method.
func (client *HybridIdentityMetadataClient) Get(ctx context.Context, connectedClusterResourceURI string, options *HybridIdentityMetadataClientGetOptions) (HybridIdentityMetadataClientGetResponse, error) {
	var err error
	const operationName = "HybridIdentityMetadataClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, connectedClusterResourceURI, options)
	if err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HybridIdentityMetadataClient) getCreateRequest(ctx context.Context, connectedClusterResourceURI string, options *HybridIdentityMetadataClientGetOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/hybridIdentityMetadata/default"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HybridIdentityMetadataClient) getHandleResponse(resp *http.Response) (HybridIdentityMetadataClientGetResponse, error) {
	result := HybridIdentityMetadataClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadata); err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	return result, nil
}

// NewListByClusterPager - Lists the hybrid identity metadata proxy resource in a provisioned cluster instance.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - HybridIdentityMetadataClientListByClusterOptions contains the optional parameters for the HybridIdentityMetadataClient.NewListByClusterPager
//     method.
func (client *HybridIdentityMetadataClient) NewListByClusterPager(connectedClusterResourceURI string, options *HybridIdentityMetadataClientListByClusterOptions) *runtime.Pager[HybridIdentityMetadataClientListByClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[HybridIdentityMetadataClientListByClusterResponse]{
		More: func(page HybridIdentityMetadataClientListByClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HybridIdentityMetadataClientListByClusterResponse) (HybridIdentityMetadataClientListByClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HybridIdentityMetadataClient.NewListByClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByClusterCreateRequest(ctx, connectedClusterResourceURI, options)
			}, nil)
			if err != nil {
				return HybridIdentityMetadataClientListByClusterResponse{}, err
			}
			return client.listByClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *HybridIdentityMetadataClient) listByClusterCreateRequest(ctx context.Context, connectedClusterResourceURI string, options *HybridIdentityMetadataClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/hybridIdentityMetadata"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *HybridIdentityMetadataClient) listByClusterHandleResponse(resp *http.Response) (HybridIdentityMetadataClientListByClusterResponse, error) {
	result := HybridIdentityMetadataClientListByClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadataList); err != nil {
		return HybridIdentityMetadataClientListByClusterResponse{}, err
	}
	return result, nil
}

// Put - Creates the hybrid identity metadata proxy resource that facilitates the managed identity provisioning.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-15-preview
//   - connectedClusterResourceURI - The fully qualified Azure Resource manager identifier of the connected cluster resource.
//   - options - HybridIdentityMetadataClientPutOptions contains the optional parameters for the HybridIdentityMetadataClient.Put
//     method.
func (client *HybridIdentityMetadataClient) Put(ctx context.Context, connectedClusterResourceURI string, body HybridIdentityMetadata, options *HybridIdentityMetadataClientPutOptions) (HybridIdentityMetadataClientPutResponse, error) {
	var err error
	const operationName = "HybridIdentityMetadataClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, connectedClusterResourceURI, body, options)
	if err != nil {
		return HybridIdentityMetadataClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridIdentityMetadataClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return HybridIdentityMetadataClientPutResponse{}, err
	}
	resp, err := client.putHandleResponse(httpResp)
	return resp, err
}

// putCreateRequest creates the Put request.
func (client *HybridIdentityMetadataClient) putCreateRequest(ctx context.Context, connectedClusterResourceURI string, body HybridIdentityMetadata, options *HybridIdentityMetadataClientPutOptions) (*policy.Request, error) {
	urlPath := "/{connectedClusterResourceUri}/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/hybridIdentityMetadata/default"
	urlPath = strings.ReplaceAll(urlPath, "{connectedClusterResourceUri}", connectedClusterResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *HybridIdentityMetadataClient) putHandleResponse(resp *http.Response) (HybridIdentityMetadataClientPutResponse, error) {
	result := HybridIdentityMetadataClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadata); err != nil {
		return HybridIdentityMetadataClientPutResponse{}, err
	}
	return result, nil
}
