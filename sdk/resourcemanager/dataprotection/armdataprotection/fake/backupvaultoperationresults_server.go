//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// BackupVaultOperationResultsServer is a fake server for instances of the armdataprotection.BackupVaultOperationResultsClient type.
type BackupVaultOperationResultsServer struct {
	// Get is the fake for method BackupVaultOperationResultsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	Get func(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *armdataprotection.BackupVaultOperationResultsClientGetOptions) (resp azfake.Responder[armdataprotection.BackupVaultOperationResultsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewBackupVaultOperationResultsServerTransport creates a new instance of BackupVaultOperationResultsServerTransport with the provided implementation.
// The returned BackupVaultOperationResultsServerTransport instance is connected to an instance of armdataprotection.BackupVaultOperationResultsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupVaultOperationResultsServerTransport(srv *BackupVaultOperationResultsServer) *BackupVaultOperationResultsServerTransport {
	return &BackupVaultOperationResultsServerTransport{srv: srv}
}

// BackupVaultOperationResultsServerTransport connects instances of armdataprotection.BackupVaultOperationResultsClient to instances of BackupVaultOperationResultsServer.
// Don't use this type directly, use NewBackupVaultOperationResultsServerTransport instead.
type BackupVaultOperationResultsServerTransport struct {
	srv *BackupVaultOperationResultsServer
}

// Do implements the policy.Transporter interface for BackupVaultOperationResultsServerTransport.
func (b *BackupVaultOperationResultsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BackupVaultOperationResultsClient.Get":
		resp, err = b.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BackupVaultOperationResultsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BackupVaultResource, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).AzureAsyncOperation; val != nil {
		resp.Header.Set("Azure-AsyncOperation", *val)
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}
