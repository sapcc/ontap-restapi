// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/ontap-restapi/example/ontap/models"
)

// VolumeCollectionGetReader is a Reader for the VolumeCollectionGet structure.
type VolumeCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeCollectionGetOK creates a VolumeCollectionGetOK with default headers values
func NewVolumeCollectionGetOK() *VolumeCollectionGetOK {
	return &VolumeCollectionGetOK{}
}

/*
VolumeCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type VolumeCollectionGetOK struct {
	Payload *models.VolumeResponse
}

// IsSuccess returns true when this volume collection get o k response has a 2xx status code
func (o *VolumeCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume collection get o k response has a 3xx status code
func (o *VolumeCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume collection get o k response has a 4xx status code
func (o *VolumeCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume collection get o k response has a 5xx status code
func (o *VolumeCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume collection get o k response a status code equal to that given
func (o *VolumeCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volume collection get o k response
func (o *VolumeCollectionGetOK) Code() int {
	return 200
}

func (o *VolumeCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes][%d] volumeCollectionGetOK %s", 200, payload)
}

func (o *VolumeCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes][%d] volumeCollectionGetOK %s", 200, payload)
}

func (o *VolumeCollectionGetOK) GetPayload() *models.VolumeResponse {
	return o.Payload
}

func (o *VolumeCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeCollectionGetDefault creates a VolumeCollectionGetDefault with default headers values
func NewVolumeCollectionGetDefault(code int) *VolumeCollectionGetDefault {
	return &VolumeCollectionGetDefault{
		_statusCode: code,
	}
}

/*
VolumeCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type VolumeCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this volume collection get default response has a 2xx status code
func (o *VolumeCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume collection get default response has a 3xx status code
func (o *VolumeCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume collection get default response has a 4xx status code
func (o *VolumeCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume collection get default response has a 5xx status code
func (o *VolumeCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume collection get default response a status code equal to that given
func (o *VolumeCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume collection get default response
func (o *VolumeCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *VolumeCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes][%d] volume_collection_get default %s", o._statusCode, payload)
}

func (o *VolumeCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes][%d] volume_collection_get default %s", o._statusCode, payload)
}

func (o *VolumeCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
