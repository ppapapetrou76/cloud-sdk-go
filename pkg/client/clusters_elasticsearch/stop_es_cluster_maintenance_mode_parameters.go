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
	"github.com/go-openapi/swag"
)

// NewStopEsClusterMaintenanceModeParams creates a new StopEsClusterMaintenanceModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopEsClusterMaintenanceModeParams() *StopEsClusterMaintenanceModeParams {
	return &StopEsClusterMaintenanceModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopEsClusterMaintenanceModeParamsWithTimeout creates a new StopEsClusterMaintenanceModeParams object
// with the ability to set a timeout on a request.
func NewStopEsClusterMaintenanceModeParamsWithTimeout(timeout time.Duration) *StopEsClusterMaintenanceModeParams {
	return &StopEsClusterMaintenanceModeParams{
		timeout: timeout,
	}
}

// NewStopEsClusterMaintenanceModeParamsWithContext creates a new StopEsClusterMaintenanceModeParams object
// with the ability to set a context for a request.
func NewStopEsClusterMaintenanceModeParamsWithContext(ctx context.Context) *StopEsClusterMaintenanceModeParams {
	return &StopEsClusterMaintenanceModeParams{
		Context: ctx,
	}
}

// NewStopEsClusterMaintenanceModeParamsWithHTTPClient creates a new StopEsClusterMaintenanceModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopEsClusterMaintenanceModeParamsWithHTTPClient(client *http.Client) *StopEsClusterMaintenanceModeParams {
	return &StopEsClusterMaintenanceModeParams{
		HTTPClient: client,
	}
}

/* StopEsClusterMaintenanceModeParams contains all the parameters to send to the API endpoint
   for the stop es cluster maintenance mode operation.

   Typically these are written to a http.Request.
*/
type StopEsClusterMaintenanceModeParams struct {

	/* ClusterID.

	   The Elasticsearch cluster identifier.
	*/
	ClusterID string

	/* IgnoreMissing.

	   When `true` and the instance does not exist, proceeds to the next instance, or treats the instance as an error.
	*/
	IgnoreMissing *bool

	/* InstanceIds.

	   A comma-separated list of instance identifiers.
	*/
	InstanceIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop es cluster maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopEsClusterMaintenanceModeParams) WithDefaults() *StopEsClusterMaintenanceModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop es cluster maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopEsClusterMaintenanceModeParams) SetDefaults() {
	var (
		ignoreMissingDefault = bool(false)
	)

	val := StopEsClusterMaintenanceModeParams{
		IgnoreMissing: &ignoreMissingDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) WithTimeout(timeout time.Duration) *StopEsClusterMaintenanceModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) WithContext(ctx context.Context) *StopEsClusterMaintenanceModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) WithHTTPClient(client *http.Client) *StopEsClusterMaintenanceModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) WithClusterID(clusterID string) *StopEsClusterMaintenanceModeParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithIgnoreMissing adds the ignoreMissing to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) WithIgnoreMissing(ignoreMissing *bool) *StopEsClusterMaintenanceModeParams {
	o.SetIgnoreMissing(ignoreMissing)
	return o
}

// SetIgnoreMissing adds the ignoreMissing to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) SetIgnoreMissing(ignoreMissing *bool) {
	o.IgnoreMissing = ignoreMissing
}

// WithInstanceIds adds the instanceIds to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) WithInstanceIds(instanceIds []string) *StopEsClusterMaintenanceModeParams {
	o.SetInstanceIds(instanceIds)
	return o
}

// SetInstanceIds adds the instanceIds to the stop es cluster maintenance mode params
func (o *StopEsClusterMaintenanceModeParams) SetInstanceIds(instanceIds []string) {
	o.InstanceIds = instanceIds
}

// WriteToRequest writes these params to a swagger request
func (o *StopEsClusterMaintenanceModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.IgnoreMissing != nil {

		// query param ignore_missing
		var qrIgnoreMissing bool

		if o.IgnoreMissing != nil {
			qrIgnoreMissing = *o.IgnoreMissing
		}
		qIgnoreMissing := swag.FormatBool(qrIgnoreMissing)
		if qIgnoreMissing != "" {

			if err := r.SetQueryParam("ignore_missing", qIgnoreMissing); err != nil {
				return err
			}
		}
	}

	if o.InstanceIds != nil {

		// binding items for instance_ids
		joinedInstanceIds := o.bindParamInstanceIds(reg)

		// path array param instance_ids
		// SetPathParam does not support variadic arguments, since we used JoinByFormat
		// we can send the first item in the array as it's all the items of the previous
		// array joined together
		if len(joinedInstanceIds) > 0 {
			if err := r.SetPathParam("instance_ids", joinedInstanceIds[0]); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamStopEsClusterMaintenanceMode binds the parameter instance_ids
func (o *StopEsClusterMaintenanceModeParams) bindParamInstanceIds(formats strfmt.Registry) []string {
	instanceIdsIR := o.InstanceIds

	var instanceIdsIC []string
	for _, instanceIdsIIR := range instanceIdsIR { // explode []string

		instanceIdsIIV := instanceIdsIIR // string as string
		instanceIdsIC = append(instanceIdsIC, instanceIdsIIV)
	}

	// items.CollectionFormat: "csv"
	instanceIdsIS := swag.JoinByFormat(instanceIdsIC, "csv")

	return instanceIdsIS
}
