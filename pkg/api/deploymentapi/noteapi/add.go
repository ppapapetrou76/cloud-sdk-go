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

package noteapi

import (
	"context"
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments_notes"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// AddParams is consumed by Add.
type AddParams struct {
	Params
	Message string
	UserID  string
}

// Validate ensures the parameters are valid
func (params AddParams) Validate() error {
	var merr = multierror.NewPrefixed("deployment note add")
	if params.UserID == "" {
		merr = merr.Append(errors.New(errEmptyUserID))
	}

	if params.Message == "" {
		merr = merr.Append(errors.New(errEmptyNoteMessage))
	}

	merr = merr.Append(params.Params.Validate())

	return merr.ErrorOrNil()
}

// Add posts a new message to the specified deployment
func Add(params AddParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(params.V1API.DeploymentsNotes.CreateDeploymentNote(
		deployments_notes.NewCreateDeploymentNoteParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithDeploymentID(params.ID).
			WithBody(&models.Note{
				Message: ec.String(params.Message),
				UserID:  params.UserID,
			}),
		params.AuthWriter,
	))
}
