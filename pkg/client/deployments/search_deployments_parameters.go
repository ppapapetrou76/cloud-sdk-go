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

package deployments

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

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewSearchDeploymentsParams creates a new SearchDeploymentsParams object
// with the default values initialized.
func NewSearchDeploymentsParams() *SearchDeploymentsParams {
	var ()
	return &SearchDeploymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchDeploymentsParamsWithTimeout creates a new SearchDeploymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchDeploymentsParamsWithTimeout(timeout time.Duration) *SearchDeploymentsParams {
	var ()
	return &SearchDeploymentsParams{

		timeout: timeout,
	}
}

// NewSearchDeploymentsParamsWithContext creates a new SearchDeploymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchDeploymentsParamsWithContext(ctx context.Context) *SearchDeploymentsParams {
	var ()
	return &SearchDeploymentsParams{

		Context: ctx,
	}
}

// NewSearchDeploymentsParamsWithHTTPClient creates a new SearchDeploymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchDeploymentsParamsWithHTTPClient(client *http.Client) *SearchDeploymentsParams {
	var ()
	return &SearchDeploymentsParams{
		HTTPClient: client,
	}
}

/*SearchDeploymentsParams contains all the parameters to send to the API endpoint
for the search deployments operation typically these are written to a http.Request
*/
type SearchDeploymentsParams struct {

	/*Body
	  (Optional) The search request to execute. NOTE: When not specified, all of the deployments are matched.

	*/
	Body *models.SearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search deployments params
func (o *SearchDeploymentsParams) WithTimeout(timeout time.Duration) *SearchDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search deployments params
func (o *SearchDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search deployments params
func (o *SearchDeploymentsParams) WithContext(ctx context.Context) *SearchDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search deployments params
func (o *SearchDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search deployments params
func (o *SearchDeploymentsParams) WithHTTPClient(client *http.Client) *SearchDeploymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search deployments params
func (o *SearchDeploymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search deployments params
func (o *SearchDeploymentsParams) WithBody(body *models.SearchRequest) *SearchDeploymentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search deployments params
func (o *SearchDeploymentsParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
