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

// SearchDeploymentsReader is a Reader for the SearchDeployments structure.
type SearchDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSearchDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchDeploymentsOK creates a SearchDeploymentsOK with default headers values
func NewSearchDeploymentsOK() *SearchDeploymentsOK {
	return &SearchDeploymentsOK{}
}

/*SearchDeploymentsOK handles this case with default header values.

The list of deployments that match the specified query and belong to the authenticated user
*/
type SearchDeploymentsOK struct {
	Payload *models.DeploymentsSearchResponse
}

func (o *SearchDeploymentsOK) Error() string {
	return fmt.Sprintf("[POST /deployments/_search][%d] searchDeploymentsOK  %+v", 200, o.Payload)
}

func (o *SearchDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentsSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchDeploymentsUnauthorized creates a SearchDeploymentsUnauthorized with default headers values
func NewSearchDeploymentsUnauthorized() *SearchDeploymentsUnauthorized {
	return &SearchDeploymentsUnauthorized{}
}

/*SearchDeploymentsUnauthorized handles this case with default header values.

You are not authorized to perform this action
*/
type SearchDeploymentsUnauthorized struct {
	Payload *models.BasicFailedReply
}

func (o *SearchDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /deployments/_search][%d] searchDeploymentsUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
