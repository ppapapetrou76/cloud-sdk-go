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

// NewResyncEsClusterParams creates a new ResyncEsClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResyncEsClusterParams() *ResyncEsClusterParams {
	return &ResyncEsClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResyncEsClusterParamsWithTimeout creates a new ResyncEsClusterParams object
// with the ability to set a timeout on a request.
func NewResyncEsClusterParamsWithTimeout(timeout time.Duration) *ResyncEsClusterParams {
	return &ResyncEsClusterParams{
		timeout: timeout,
	}
}

// NewResyncEsClusterParamsWithContext creates a new ResyncEsClusterParams object
// with the ability to set a context for a request.
func NewResyncEsClusterParamsWithContext(ctx context.Context) *ResyncEsClusterParams {
	return &ResyncEsClusterParams{
		Context: ctx,
	}
}

// NewResyncEsClusterParamsWithHTTPClient creates a new ResyncEsClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewResyncEsClusterParamsWithHTTPClient(client *http.Client) *ResyncEsClusterParams {
	return &ResyncEsClusterParams{
		HTTPClient: client,
	}
}

/* ResyncEsClusterParams contains all the parameters to send to the API endpoint
   for the resync es cluster operation.

   Typically these are written to a http.Request.
*/
type ResyncEsClusterParams struct {

	/* ClusterID.

	   The Elasticsearch cluster identifier.
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resync es cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResyncEsClusterParams) WithDefaults() *ResyncEsClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resync es cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResyncEsClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resync es cluster params
func (o *ResyncEsClusterParams) WithTimeout(timeout time.Duration) *ResyncEsClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resync es cluster params
func (o *ResyncEsClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resync es cluster params
func (o *ResyncEsClusterParams) WithContext(ctx context.Context) *ResyncEsClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resync es cluster params
func (o *ResyncEsClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resync es cluster params
func (o *ResyncEsClusterParams) WithHTTPClient(client *http.Client) *ResyncEsClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resync es cluster params
func (o *ResyncEsClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the resync es cluster params
func (o *ResyncEsClusterParams) WithClusterID(clusterID string) *ResyncEsClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the resync es cluster params
func (o *ResyncEsClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *ResyncEsClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
