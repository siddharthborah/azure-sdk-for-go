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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
	"net/http"
	"net/url"
	"regexp"
)

// TransformationsServer is a fake server for instances of the armstreamanalytics.TransformationsClient type.
type TransformationsServer struct {
	// CreateOrReplace is the fake for method TransformationsClient.CreateOrReplace
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrReplace func(ctx context.Context, resourceGroupName string, jobName string, transformationName string, transformation armstreamanalytics.Transformation, options *armstreamanalytics.TransformationsClientCreateOrReplaceOptions) (resp azfake.Responder[armstreamanalytics.TransformationsClientCreateOrReplaceResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TransformationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, jobName string, transformationName string, options *armstreamanalytics.TransformationsClientGetOptions) (resp azfake.Responder[armstreamanalytics.TransformationsClientGetResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method TransformationsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, jobName string, transformationName string, transformation armstreamanalytics.Transformation, options *armstreamanalytics.TransformationsClientUpdateOptions) (resp azfake.Responder[armstreamanalytics.TransformationsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewTransformationsServerTransport creates a new instance of TransformationsServerTransport with the provided implementation.
// The returned TransformationsServerTransport instance is connected to an instance of armstreamanalytics.TransformationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTransformationsServerTransport(srv *TransformationsServer) *TransformationsServerTransport {
	return &TransformationsServerTransport{srv: srv}
}

// TransformationsServerTransport connects instances of armstreamanalytics.TransformationsClient to instances of TransformationsServer.
// Don't use this type directly, use NewTransformationsServerTransport instead.
type TransformationsServerTransport struct {
	srv *TransformationsServer
}

// Do implements the policy.Transporter interface for TransformationsServerTransport.
func (t *TransformationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TransformationsClient.CreateOrReplace":
		resp, err = t.dispatchCreateOrReplace(req)
	case "TransformationsClient.Get":
		resp, err = t.dispatchGet(req)
	case "TransformationsClient.Update":
		resp, err = t.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TransformationsServerTransport) dispatchCreateOrReplace(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrReplace == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrReplace not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StreamAnalytics/streamingjobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transformations/(?P<transformationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstreamanalytics.Transformation](req)
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	transformationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("transformationName")])
	if err != nil {
		return nil, err
	}
	var options *armstreamanalytics.TransformationsClientCreateOrReplaceOptions
	if ifMatchParam != nil || ifNoneMatchParam != nil {
		options = &armstreamanalytics.TransformationsClientCreateOrReplaceOptions{
			IfMatch:     ifMatchParam,
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := t.srv.CreateOrReplace(req.Context(), resourceGroupNameParam, jobNameParam, transformationNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Transformation, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (t *TransformationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StreamAnalytics/streamingjobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transformations/(?P<transformationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	transformationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("transformationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, jobNameParam, transformationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Transformation, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (t *TransformationsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StreamAnalytics/streamingjobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transformations/(?P<transformationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstreamanalytics.Transformation](req)
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	transformationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("transformationName")])
	if err != nil {
		return nil, err
	}
	var options *armstreamanalytics.TransformationsClientUpdateOptions
	if ifMatchParam != nil {
		options = &armstreamanalytics.TransformationsClientUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := t.srv.Update(req.Context(), resourceGroupNameParam, jobNameParam, transformationNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Transformation, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}
