// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPutEsProxyRequestsParams creates a new PutEsProxyRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutEsProxyRequestsParams() *PutEsProxyRequestsParams {
	return &PutEsProxyRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutEsProxyRequestsParamsWithTimeout creates a new PutEsProxyRequestsParams object
// with the ability to set a timeout on a request.
func NewPutEsProxyRequestsParamsWithTimeout(timeout time.Duration) *PutEsProxyRequestsParams {
	return &PutEsProxyRequestsParams{
		timeout: timeout,
	}
}

// NewPutEsProxyRequestsParamsWithContext creates a new PutEsProxyRequestsParams object
// with the ability to set a context for a request.
func NewPutEsProxyRequestsParamsWithContext(ctx context.Context) *PutEsProxyRequestsParams {
	return &PutEsProxyRequestsParams{
		Context: ctx,
	}
}

// NewPutEsProxyRequestsParamsWithHTTPClient creates a new PutEsProxyRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutEsProxyRequestsParamsWithHTTPClient(client *http.Client) *PutEsProxyRequestsParams {
	return &PutEsProxyRequestsParams{
		HTTPClient: client,
	}
}

/* PutEsProxyRequestsParams contains all the parameters to send to the API endpoint
   for the put es proxy requests operation.

   Typically these are written to a http.Request.
*/
type PutEsProxyRequestsParams struct {

	/* XManagementRequest.

	   X-Management-Request header value. Needs to be set to true
	*/
	XManagementRequest string

	/* Body.

	   The JSON payload to proxy to the Elasticsearch cluster
	*/
	Body string

	/* ClusterID.

	   Identifier for the Elasticsearch cluster
	*/
	ClusterID string

	/* ElasticsearchPath.

	   The URL part to proxy to the Elasticsearch cluster. Example: _search or _cat/indices?v&h=i,tm&s=tm:desc
	*/
	ElasticsearchPath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put es proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEsProxyRequestsParams) WithDefaults() *PutEsProxyRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put es proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEsProxyRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put es proxy requests params
func (o *PutEsProxyRequestsParams) WithTimeout(timeout time.Duration) *PutEsProxyRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put es proxy requests params
func (o *PutEsProxyRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put es proxy requests params
func (o *PutEsProxyRequestsParams) WithContext(ctx context.Context) *PutEsProxyRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put es proxy requests params
func (o *PutEsProxyRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put es proxy requests params
func (o *PutEsProxyRequestsParams) WithHTTPClient(client *http.Client) *PutEsProxyRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put es proxy requests params
func (o *PutEsProxyRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXManagementRequest adds the xManagementRequest to the put es proxy requests params
func (o *PutEsProxyRequestsParams) WithXManagementRequest(xManagementRequest string) *PutEsProxyRequestsParams {
	o.SetXManagementRequest(xManagementRequest)
	return o
}

// SetXManagementRequest adds the xManagementRequest to the put es proxy requests params
func (o *PutEsProxyRequestsParams) SetXManagementRequest(xManagementRequest string) {
	o.XManagementRequest = xManagementRequest
}

// WithBody adds the body to the put es proxy requests params
func (o *PutEsProxyRequestsParams) WithBody(body string) *PutEsProxyRequestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put es proxy requests params
func (o *PutEsProxyRequestsParams) SetBody(body string) {
	o.Body = body
}

// WithClusterID adds the clusterID to the put es proxy requests params
func (o *PutEsProxyRequestsParams) WithClusterID(clusterID string) *PutEsProxyRequestsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the put es proxy requests params
func (o *PutEsProxyRequestsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithElasticsearchPath adds the elasticsearchPath to the put es proxy requests params
func (o *PutEsProxyRequestsParams) WithElasticsearchPath(elasticsearchPath string) *PutEsProxyRequestsParams {
	o.SetElasticsearchPath(elasticsearchPath)
	return o
}

// SetElasticsearchPath adds the elasticsearchPath to the put es proxy requests params
func (o *PutEsProxyRequestsParams) SetElasticsearchPath(elasticsearchPath string) {
	o.ElasticsearchPath = elasticsearchPath
}

// WriteToRequest writes these params to a swagger request
func (o *PutEsProxyRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Management-Request
	if err := r.SetHeaderParam("X-Management-Request", o.XManagementRequest); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param elasticsearch_path
	if err := r.SetPathParam("elasticsearch_path", o.ElasticsearchPath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
