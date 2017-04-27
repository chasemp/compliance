package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// BatchInputV2 batch input v2
// swagger:model BatchInputV2
type BatchInputV2 struct {

	// documents
	Documents []*InputV2 `json:"documents"`
}

// Validate validates this batch input v2
func (m *BatchInputV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocuments(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchInputV2) validateDocuments(formats strfmt.Registry) error {

	if swag.IsZero(m.Documents) { // not required
		return nil
	}

	for i := 0; i < len(m.Documents); i++ {

		if swag.IsZero(m.Documents[i]) { // not required
			continue
		}

		if m.Documents[i] != nil {

			if err := m.Documents[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}