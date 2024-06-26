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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
	"net/http"
	"net/url"
	"regexp"
)

// Server is a fake server for instances of the armrecoveryservicesbackup.Client type.
type Server struct {
	// BeginBMSPrepareDataMove is the fake for method Client.BeginBMSPrepareDataMove
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginBMSPrepareDataMove func(ctx context.Context, vaultName string, resourceGroupName string, parameters armrecoveryservicesbackup.PrepareDataMoveRequest, options *armrecoveryservicesbackup.ClientBeginBMSPrepareDataMoveOptions) (resp azfake.PollerResponder[armrecoveryservicesbackup.ClientBMSPrepareDataMoveResponse], errResp azfake.ErrorResponder)

	// BeginBMSTriggerDataMove is the fake for method Client.BeginBMSTriggerDataMove
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginBMSTriggerDataMove func(ctx context.Context, vaultName string, resourceGroupName string, parameters armrecoveryservicesbackup.TriggerDataMoveRequest, options *armrecoveryservicesbackup.ClientBeginBMSTriggerDataMoveOptions) (resp azfake.PollerResponder[armrecoveryservicesbackup.ClientBMSTriggerDataMoveResponse], errResp azfake.ErrorResponder)

	// GetOperationStatus is the fake for method Client.GetOperationStatus
	// HTTP status codes to indicate success: http.StatusOK
	GetOperationStatus func(ctx context.Context, vaultName string, resourceGroupName string, operationID string, options *armrecoveryservicesbackup.ClientGetOperationStatusOptions) (resp azfake.Responder[armrecoveryservicesbackup.ClientGetOperationStatusResponse], errResp azfake.ErrorResponder)

	// BeginMoveRecoveryPoint is the fake for method Client.BeginMoveRecoveryPoint
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginMoveRecoveryPoint func(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters armrecoveryservicesbackup.MoveRPAcrossTiersRequest, options *armrecoveryservicesbackup.ClientBeginMoveRecoveryPointOptions) (resp azfake.PollerResponder[armrecoveryservicesbackup.ClientMoveRecoveryPointResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of armrecoveryservicesbackup.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{
		srv:                     srv,
		beginBMSPrepareDataMove: newTracker[azfake.PollerResponder[armrecoveryservicesbackup.ClientBMSPrepareDataMoveResponse]](),
		beginBMSTriggerDataMove: newTracker[azfake.PollerResponder[armrecoveryservicesbackup.ClientBMSTriggerDataMoveResponse]](),
		beginMoveRecoveryPoint:  newTracker[azfake.PollerResponder[armrecoveryservicesbackup.ClientMoveRecoveryPointResponse]](),
	}
}

// ServerTransport connects instances of armrecoveryservicesbackup.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv                     *Server
	beginBMSPrepareDataMove *tracker[azfake.PollerResponder[armrecoveryservicesbackup.ClientBMSPrepareDataMoveResponse]]
	beginBMSTriggerDataMove *tracker[azfake.PollerResponder[armrecoveryservicesbackup.ClientBMSTriggerDataMoveResponse]]
	beginMoveRecoveryPoint  *tracker[azfake.PollerResponder[armrecoveryservicesbackup.ClientMoveRecoveryPointResponse]]
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Client.BeginBMSPrepareDataMove":
		resp, err = s.dispatchBeginBMSPrepareDataMove(req)
	case "Client.BeginBMSTriggerDataMove":
		resp, err = s.dispatchBeginBMSTriggerDataMove(req)
	case "Client.GetOperationStatus":
		resp, err = s.dispatchGetOperationStatus(req)
	case "Client.BeginMoveRecoveryPoint":
		resp, err = s.dispatchBeginMoveRecoveryPoint(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerTransport) dispatchBeginBMSPrepareDataMove(req *http.Request) (*http.Response, error) {
	if s.srv.BeginBMSPrepareDataMove == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginBMSPrepareDataMove not implemented")}
	}
	beginBMSPrepareDataMove := s.beginBMSPrepareDataMove.get(req)
	if beginBMSPrepareDataMove == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupstorageconfig/vaultstorageconfig/prepareDataMove`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.PrepareDataMoveRequest](req)
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginBMSPrepareDataMove(req.Context(), vaultNameParam, resourceGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginBMSPrepareDataMove = &respr
		s.beginBMSPrepareDataMove.add(req, beginBMSPrepareDataMove)
	}

	resp, err := server.PollerResponderNext(beginBMSPrepareDataMove, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginBMSPrepareDataMove.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginBMSPrepareDataMove) {
		s.beginBMSPrepareDataMove.remove(req)
	}

	return resp, nil
}

func (s *ServerTransport) dispatchBeginBMSTriggerDataMove(req *http.Request) (*http.Response, error) {
	if s.srv.BeginBMSTriggerDataMove == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginBMSTriggerDataMove not implemented")}
	}
	beginBMSTriggerDataMove := s.beginBMSTriggerDataMove.get(req)
	if beginBMSTriggerDataMove == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupstorageconfig/vaultstorageconfig/triggerDataMove`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.TriggerDataMoveRequest](req)
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginBMSTriggerDataMove(req.Context(), vaultNameParam, resourceGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginBMSTriggerDataMove = &respr
		s.beginBMSTriggerDataMove.add(req, beginBMSTriggerDataMove)
	}

	resp, err := server.PollerResponderNext(beginBMSTriggerDataMove, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginBMSTriggerDataMove.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginBMSTriggerDataMove) {
		s.beginBMSTriggerDataMove.remove(req)
	}

	return resp, nil
}

func (s *ServerTransport) dispatchGetOperationStatus(req *http.Request) (*http.Response, error) {
	if s.srv.GetOperationStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetOperationStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupstorageconfig/vaultstorageconfig/operationStatus/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetOperationStatus(req.Context(), vaultNameParam, resourceGroupNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchBeginMoveRecoveryPoint(req *http.Request) (*http.Response, error) {
	if s.srv.BeginMoveRecoveryPoint == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginMoveRecoveryPoint not implemented")}
	}
	beginMoveRecoveryPoint := s.beginMoveRecoveryPoint.get(req)
	if beginMoveRecoveryPoint == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectionContainers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recoveryPoints/(?P<recoveryPointId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/move`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 7 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.MoveRPAcrossTiersRequest](req)
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
		if err != nil {
			return nil, err
		}
		protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
		if err != nil {
			return nil, err
		}
		recoveryPointIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPointId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginMoveRecoveryPoint(req.Context(), vaultNameParam, resourceGroupNameParam, fabricNameParam, containerNameParam, protectedItemNameParam, recoveryPointIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginMoveRecoveryPoint = &respr
		s.beginMoveRecoveryPoint.add(req, beginMoveRecoveryPoint)
	}

	resp, err := server.PollerResponderNext(beginMoveRecoveryPoint, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		s.beginMoveRecoveryPoint.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginMoveRecoveryPoint) {
		s.beginMoveRecoveryPoint.remove(req)
	}

	return resp, nil
}
