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

package deployment_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetDeploymentTemplateV2Reader is a Reader for the GetDeploymentTemplateV2 structure.
type GetDeploymentTemplateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentTemplateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentTemplateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeploymentTemplateV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDeploymentTemplateV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentTemplateV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeploymentTemplateV2OK creates a GetDeploymentTemplateV2OK with default headers values
func NewGetDeploymentTemplateV2OK() *GetDeploymentTemplateV2OK {
	return &GetDeploymentTemplateV2OK{}
}

/*GetDeploymentTemplateV2OK handles this case with default header values.

The deployment template was found and returned successfully.
*/
type GetDeploymentTemplateV2OK struct {
	Payload *models.DeploymentTemplateInfoV2
}

func (o *GetDeploymentTemplateV2OK) Error() string {
	return fmt.Sprintf("[GET /deployments/templates/{template_id}][%d] getDeploymentTemplateV2OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentTemplateV2OK) GetPayload() *models.DeploymentTemplateInfoV2 {
	return o.Payload
}

func (o *GetDeploymentTemplateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentTemplateInfoV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentTemplateV2BadRequest creates a GetDeploymentTemplateV2BadRequest with default headers values
func NewGetDeploymentTemplateV2BadRequest() *GetDeploymentTemplateV2BadRequest {
	return &GetDeploymentTemplateV2BadRequest{}
}

/*GetDeploymentTemplateV2BadRequest handles this case with default header values.

The requested region is not supported. (code: `templates.region_not_found`)
*/
type GetDeploymentTemplateV2BadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentTemplateV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /deployments/templates/{template_id}][%d] getDeploymentTemplateV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetDeploymentTemplateV2BadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentTemplateV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentTemplateV2Unauthorized creates a GetDeploymentTemplateV2Unauthorized with default headers values
func NewGetDeploymentTemplateV2Unauthorized() *GetDeploymentTemplateV2Unauthorized {
	return &GetDeploymentTemplateV2Unauthorized{}
}

/*GetDeploymentTemplateV2Unauthorized handles this case with default header values.

The user is not authorized to access requested region. (code: `templates.region_not_allowed`)
*/
type GetDeploymentTemplateV2Unauthorized struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentTemplateV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployments/templates/{template_id}][%d] getDeploymentTemplateV2Unauthorized  %+v", 401, o.Payload)
}

func (o *GetDeploymentTemplateV2Unauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentTemplateV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentTemplateV2NotFound creates a GetDeploymentTemplateV2NotFound with default headers values
func NewGetDeploymentTemplateV2NotFound() *GetDeploymentTemplateV2NotFound {
	return &GetDeploymentTemplateV2NotFound{}
}

/*GetDeploymentTemplateV2NotFound handles this case with default header values.

The deployment template specified by {template_id} cannot be found. (code: `templates.template_not_found`)
*/
type GetDeploymentTemplateV2NotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentTemplateV2NotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/templates/{template_id}][%d] getDeploymentTemplateV2NotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentTemplateV2NotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentTemplateV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}