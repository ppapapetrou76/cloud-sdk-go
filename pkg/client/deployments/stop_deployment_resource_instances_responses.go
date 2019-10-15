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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// StopDeploymentResourceInstancesReader is a Reader for the StopDeploymentResourceInstances structure.
type StopDeploymentResourceInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopDeploymentResourceInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewStopDeploymentResourceInstancesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewStopDeploymentResourceInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStopDeploymentResourceInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewStopDeploymentResourceInstancesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewStopDeploymentResourceInstancesRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopDeploymentResourceInstancesAccepted creates a StopDeploymentResourceInstancesAccepted with default headers values
func NewStopDeploymentResourceInstancesAccepted() *StopDeploymentResourceInstancesAccepted {
	return &StopDeploymentResourceInstancesAccepted{}
}

/*StopDeploymentResourceInstancesAccepted handles this case with default header values.

The stop command was issued successfully
*/
type StopDeploymentResourceInstancesAccepted struct {
	Payload models.DeploymentResourceCommandResponse
}

func (o *StopDeploymentResourceInstancesAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/{instance_ids}/_stop][%d] stopDeploymentResourceInstancesAccepted  %+v", 202, o.Payload)
}

func (o *StopDeploymentResourceInstancesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopDeploymentResourceInstancesForbidden creates a StopDeploymentResourceInstancesForbidden with default headers values
func NewStopDeploymentResourceInstancesForbidden() *StopDeploymentResourceInstancesForbidden {
	return &StopDeploymentResourceInstancesForbidden{}
}

/*StopDeploymentResourceInstancesForbidden handles this case with default header values.

The stop command was prohibited for the given Resource.
*/
type StopDeploymentResourceInstancesForbidden struct {
	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceInstancesForbidden) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/{instance_ids}/_stop][%d] stopDeploymentResourceInstancesForbidden  %+v", 403, o.Payload)
}

func (o *StopDeploymentResourceInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopDeploymentResourceInstancesNotFound creates a StopDeploymentResourceInstancesNotFound with default headers values
func NewStopDeploymentResourceInstancesNotFound() *StopDeploymentResourceInstancesNotFound {
	return &StopDeploymentResourceInstancesNotFound{}
}

/*StopDeploymentResourceInstancesNotFound handles this case with default header values.

The Resource specified by {ref_id} cannot be found
*/
type StopDeploymentResourceInstancesNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceInstancesNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/{instance_ids}/_stop][%d] stopDeploymentResourceInstancesNotFound  %+v", 404, o.Payload)
}

func (o *StopDeploymentResourceInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopDeploymentResourceInstancesUnprocessableEntity creates a StopDeploymentResourceInstancesUnprocessableEntity with default headers values
func NewStopDeploymentResourceInstancesUnprocessableEntity() *StopDeploymentResourceInstancesUnprocessableEntity {
	return &StopDeploymentResourceInstancesUnprocessableEntity{}
}

/*StopDeploymentResourceInstancesUnprocessableEntity handles this case with default header values.

The command sent to a Resource found the Resource in an illegal state, the error message gives more details
*/
type StopDeploymentResourceInstancesUnprocessableEntity struct {
	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceInstancesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/{instance_ids}/_stop][%d] stopDeploymentResourceInstancesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *StopDeploymentResourceInstancesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopDeploymentResourceInstancesRetryWith creates a StopDeploymentResourceInstancesRetryWith with default headers values
func NewStopDeploymentResourceInstancesRetryWith() *StopDeploymentResourceInstancesRetryWith {
	return &StopDeploymentResourceInstancesRetryWith{}
}

/*StopDeploymentResourceInstancesRetryWith handles this case with default header values.

elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type StopDeploymentResourceInstancesRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceInstancesRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/{instance_ids}/_stop][%d] stopDeploymentResourceInstancesRetryWith  %+v", 449, o.Payload)
}

func (o *StopDeploymentResourceInstancesRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
