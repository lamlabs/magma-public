// Code generated by go-swagger; DO NOT EDIT.

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutNetworksNetworkIDTiersTierIDNameReader is a Reader for the PutNetworksNetworkIDTiersTierIDName structure.
type PutNetworksNetworkIDTiersTierIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDTiersTierIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDTiersTierIDNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDTiersTierIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDTiersTierIDNameNoContent creates a PutNetworksNetworkIDTiersTierIDNameNoContent with default headers values
func NewPutNetworksNetworkIDTiersTierIDNameNoContent() *PutNetworksNetworkIDTiersTierIDNameNoContent {
	return &PutNetworksNetworkIDTiersTierIDNameNoContent{}
}

/*PutNetworksNetworkIDTiersTierIDNameNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDTiersTierIDNameNoContent struct {
}

func (o *PutNetworksNetworkIDTiersTierIDNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/tiers/{tier_id}/name][%d] putNetworksNetworkIdTiersTierIdNameNoContent ", 204)
}

func (o *PutNetworksNetworkIDTiersTierIDNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDTiersTierIDNameDefault creates a PutNetworksNetworkIDTiersTierIDNameDefault with default headers values
func NewPutNetworksNetworkIDTiersTierIDNameDefault(code int) *PutNetworksNetworkIDTiersTierIDNameDefault {
	return &PutNetworksNetworkIDTiersTierIDNameDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDTiersTierIDNameDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDTiersTierIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID tiers tier ID name default response
func (o *PutNetworksNetworkIDTiersTierIDNameDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDTiersTierIDNameDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/tiers/{tier_id}/name][%d] PutNetworksNetworkIDTiersTierIDName default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDTiersTierIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDTiersTierIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}