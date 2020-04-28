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

package deployments_traffic_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTrafficFilterRulesetDeploymentAssociationsParams creates a new GetTrafficFilterRulesetDeploymentAssociationsParams object
// with the default values initialized.
func NewGetTrafficFilterRulesetDeploymentAssociationsParams() *GetTrafficFilterRulesetDeploymentAssociationsParams {
	var ()
	return &GetTrafficFilterRulesetDeploymentAssociationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTrafficFilterRulesetDeploymentAssociationsParamsWithTimeout creates a new GetTrafficFilterRulesetDeploymentAssociationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTrafficFilterRulesetDeploymentAssociationsParamsWithTimeout(timeout time.Duration) *GetTrafficFilterRulesetDeploymentAssociationsParams {
	var ()
	return &GetTrafficFilterRulesetDeploymentAssociationsParams{

		timeout: timeout,
	}
}

// NewGetTrafficFilterRulesetDeploymentAssociationsParamsWithContext creates a new GetTrafficFilterRulesetDeploymentAssociationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTrafficFilterRulesetDeploymentAssociationsParamsWithContext(ctx context.Context) *GetTrafficFilterRulesetDeploymentAssociationsParams {
	var ()
	return &GetTrafficFilterRulesetDeploymentAssociationsParams{

		Context: ctx,
	}
}

// NewGetTrafficFilterRulesetDeploymentAssociationsParamsWithHTTPClient creates a new GetTrafficFilterRulesetDeploymentAssociationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTrafficFilterRulesetDeploymentAssociationsParamsWithHTTPClient(client *http.Client) *GetTrafficFilterRulesetDeploymentAssociationsParams {
	var ()
	return &GetTrafficFilterRulesetDeploymentAssociationsParams{
		HTTPClient: client,
	}
}

/*GetTrafficFilterRulesetDeploymentAssociationsParams contains all the parameters to send to the API endpoint
for the get traffic filter ruleset deployment associations operation typically these are written to a http.Request
*/
type GetTrafficFilterRulesetDeploymentAssociationsParams struct {

	/*RulesetID
	  The mandatory ruleset ID.

	*/
	RulesetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get traffic filter ruleset deployment associations params
func (o *GetTrafficFilterRulesetDeploymentAssociationsParams) WithTimeout(timeout time.Duration) *GetTrafficFilterRulesetDeploymentAssociationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get traffic filter ruleset deployment associations params
func (o *GetTrafficFilterRulesetDeploymentAssociationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get traffic filter ruleset deployment associations params
func (o *GetTrafficFilterRulesetDeploymentAssociationsParams) WithContext(ctx context.Context) *GetTrafficFilterRulesetDeploymentAssociationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get traffic filter ruleset deployment associations params
func (o *GetTrafficFilterRulesetDeploymentAssociationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get traffic filter ruleset deployment associations params
func (o *GetTrafficFilterRulesetDeploymentAssociationsParams) WithHTTPClient(client *http.Client) *GetTrafficFilterRulesetDeploymentAssociationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get traffic filter ruleset deployment associations params
func (o *GetTrafficFilterRulesetDeploymentAssociationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRulesetID adds the rulesetID to the get traffic filter ruleset deployment associations params
func (o *GetTrafficFilterRulesetDeploymentAssociationsParams) WithRulesetID(rulesetID string) *GetTrafficFilterRulesetDeploymentAssociationsParams {
	o.SetRulesetID(rulesetID)
	return o
}

// SetRulesetID adds the rulesetId to the get traffic filter ruleset deployment associations params
func (o *GetTrafficFilterRulesetDeploymentAssociationsParams) SetRulesetID(rulesetID string) {
	o.RulesetID = rulesetID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTrafficFilterRulesetDeploymentAssociationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ruleset_id
	if err := r.SetPathParam("ruleset_id", o.RulesetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}