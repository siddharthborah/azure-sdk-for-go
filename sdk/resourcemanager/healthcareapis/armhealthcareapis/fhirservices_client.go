//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

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

// FhirServicesClient contains the methods for the FhirServices group.
// Don't use this type directly, use NewFhirServicesClient() instead.
type FhirServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFhirServicesClient creates a new instance of FhirServicesClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFhirServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FhirServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FhirServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a FHIR Service resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - fhirServiceName - The name of FHIR Service resource.
//   - fhirservice - The parameters for creating or updating a Fhir Service resource.
//   - options - FhirServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the FhirServicesClient.BeginCreateOrUpdate
//     method.
func (client *FhirServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, fhirServiceName string, fhirservice FhirService, options *FhirServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[FhirServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, fhirServiceName, fhirservice, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FhirServicesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FhirServicesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a FHIR Service resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
func (client *FhirServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, fhirServiceName string, fhirservice FhirService, options *FhirServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FhirServicesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, fhirServiceName, fhirservice, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FhirServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, fhirServiceName string, fhirservice FhirService, options *FhirServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/fhirservices/{fhirServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if fhirServiceName == "" {
		return nil, errors.New("parameter fhirServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fhirServiceName}", url.PathEscape(fhirServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, fhirservice); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a FHIR Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - fhirServiceName - The name of FHIR Service resource.
//   - workspaceName - The name of workspace resource.
//   - options - FhirServicesClientBeginDeleteOptions contains the optional parameters for the FhirServicesClient.BeginDelete
//     method.
func (client *FhirServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, fhirServiceName string, workspaceName string, options *FhirServicesClientBeginDeleteOptions) (*runtime.Poller[FhirServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, fhirServiceName, workspaceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FhirServicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FhirServicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a FHIR Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
func (client *FhirServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, fhirServiceName string, workspaceName string, options *FhirServicesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "FhirServicesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, fhirServiceName, workspaceName, options)
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
func (client *FhirServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, fhirServiceName string, workspaceName string, options *FhirServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/fhirservices/{fhirServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fhirServiceName == "" {
		return nil, errors.New("parameter fhirServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fhirServiceName}", url.PathEscape(fhirServiceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the specified FHIR Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - fhirServiceName - The name of FHIR Service resource.
//   - options - FhirServicesClientGetOptions contains the optional parameters for the FhirServicesClient.Get method.
func (client *FhirServicesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, fhirServiceName string, options *FhirServicesClientGetOptions) (FhirServicesClientGetResponse, error) {
	var err error
	const operationName = "FhirServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, fhirServiceName, options)
	if err != nil {
		return FhirServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FhirServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FhirServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FhirServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, fhirServiceName string, options *FhirServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/fhirservices/{fhirServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if fhirServiceName == "" {
		return nil, errors.New("parameter fhirServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fhirServiceName}", url.PathEscape(fhirServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FhirServicesClient) getHandleResponse(resp *http.Response) (FhirServicesClientGetResponse, error) {
	result := FhirServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FhirService); err != nil {
		return FhirServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Lists all FHIR Services for the given workspace
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - options - FhirServicesClientListByWorkspaceOptions contains the optional parameters for the FhirServicesClient.NewListByWorkspacePager
//     method.
func (client *FhirServicesClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *FhirServicesClientListByWorkspaceOptions) *runtime.Pager[FhirServicesClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[FhirServicesClientListByWorkspaceResponse]{
		More: func(page FhirServicesClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FhirServicesClientListByWorkspaceResponse) (FhirServicesClientListByWorkspaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FhirServicesClient.NewListByWorkspacePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return FhirServicesClientListByWorkspaceResponse{}, err
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *FhirServicesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *FhirServicesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/fhirservices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *FhirServicesClient) listByWorkspaceHandleResponse(resp *http.Response) (FhirServicesClientListByWorkspaceResponse, error) {
	result := FhirServicesClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FhirServiceCollection); err != nil {
		return FhirServicesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch FHIR Service details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - fhirServiceName - The name of FHIR Service resource.
//   - workspaceName - The name of workspace resource.
//   - fhirservicePatchResource - The parameters for updating a Fhir Service.
//   - options - FhirServicesClientBeginUpdateOptions contains the optional parameters for the FhirServicesClient.BeginUpdate
//     method.
func (client *FhirServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, fhirServiceName string, workspaceName string, fhirservicePatchResource FhirServicePatchResource, options *FhirServicesClientBeginUpdateOptions) (*runtime.Poller[FhirServicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, fhirServiceName, workspaceName, fhirservicePatchResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FhirServicesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FhirServicesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch FHIR Service details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
func (client *FhirServicesClient) update(ctx context.Context, resourceGroupName string, fhirServiceName string, workspaceName string, fhirservicePatchResource FhirServicePatchResource, options *FhirServicesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FhirServicesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, fhirServiceName, workspaceName, fhirservicePatchResource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *FhirServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, fhirServiceName string, workspaceName string, fhirservicePatchResource FhirServicePatchResource, options *FhirServicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/fhirservices/{fhirServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fhirServiceName == "" {
		return nil, errors.New("parameter fhirServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fhirServiceName}", url.PathEscape(fhirServiceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, fhirservicePatchResource); err != nil {
		return nil, err
	}
	return req, nil
}
