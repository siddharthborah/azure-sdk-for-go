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
)

// OperationStatusBackupVaultContextServer is a fake server for instances of the armdataprotection.OperationStatusBackupVaultContextClient type.
type OperationStatusBackupVaultContextServer struct {
	// Get is the fake for method OperationStatusBackupVaultContextClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *armdataprotection.OperationStatusBackupVaultContextClientGetOptions) (resp azfake.Responder[armdataprotection.OperationStatusBackupVaultContextClientGetResponse], errResp azfake.ErrorResponder)
}

// NewOperationStatusBackupVaultContextServerTransport creates a new instance of OperationStatusBackupVaultContextServerTransport with the provided implementation.
// The returned OperationStatusBackupVaultContextServerTransport instance is connected to an instance of armdataprotection.OperationStatusBackupVaultContextClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationStatusBackupVaultContextServerTransport(srv *OperationStatusBackupVaultContextServer) *OperationStatusBackupVaultContextServerTransport {
	return &OperationStatusBackupVaultContextServerTransport{srv: srv}
}

// OperationStatusBackupVaultContextServerTransport connects instances of armdataprotection.OperationStatusBackupVaultContextClient to instances of OperationStatusBackupVaultContextServer.
// Don't use this type directly, use NewOperationStatusBackupVaultContextServerTransport instead.
type OperationStatusBackupVaultContextServerTransport struct {
	srv *OperationStatusBackupVaultContextServer
}

// Do implements the policy.Transporter interface for OperationStatusBackupVaultContextServerTransport.
func (o *OperationStatusBackupVaultContextServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OperationStatusBackupVaultContextClient.Get":
		resp, err = o.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OperationStatusBackupVaultContextServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operationStatus/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := o.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
