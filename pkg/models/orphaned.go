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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Orphaned Details about orphaned resources.
// swagger:model Orphaned
type Orphaned struct {

	// List of orphaned APM resource ids
	// Required: true
	Apm []string `json:"apm"`

	// List of orphaned AppSearch resource ids
	// Required: true
	Appsearch []string `json:"appsearch"`

	// List of orphaned Elasticsearch resources
	// Required: true
	Elasticsearch []*OrphanedElasticsearch `json:"elasticsearch"`

	// List of orphaned Kibana resource ids
	// Required: true
	Kibana []string `json:"kibana"`
}

// Validate validates this orphaned
func (m *Orphaned) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppsearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElasticsearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKibana(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Orphaned) validateApm(formats strfmt.Registry) error {

	if err := validate.Required("apm", "body", m.Apm); err != nil {
		return err
	}

	return nil
}

func (m *Orphaned) validateAppsearch(formats strfmt.Registry) error {

	if err := validate.Required("appsearch", "body", m.Appsearch); err != nil {
		return err
	}

	return nil
}

func (m *Orphaned) validateElasticsearch(formats strfmt.Registry) error {

	if err := validate.Required("elasticsearch", "body", m.Elasticsearch); err != nil {
		return err
	}

	for i := 0; i < len(m.Elasticsearch); i++ {
		if swag.IsZero(m.Elasticsearch[i]) { // not required
			continue
		}

		if m.Elasticsearch[i] != nil {
			if err := m.Elasticsearch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elasticsearch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Orphaned) validateKibana(formats strfmt.Registry) error {

	if err := validate.Required("kibana", "body", m.Kibana); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Orphaned) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Orphaned) UnmarshalBinary(b []byte) error {
	var res Orphaned
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
