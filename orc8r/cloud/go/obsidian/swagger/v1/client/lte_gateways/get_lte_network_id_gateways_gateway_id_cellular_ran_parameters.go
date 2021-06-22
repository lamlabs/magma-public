// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

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

// NewGetLTENetworkIDGatewaysGatewayIDCellularRanParams creates a new GetLTENetworkIDGatewaysGatewayIDCellularRanParams object
// with the default values initialized.
func NewGetLTENetworkIDGatewaysGatewayIDCellularRanParams() *GetLTENetworkIDGatewaysGatewayIDCellularRanParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDCellularRanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularRanParamsWithTimeout creates a new GetLTENetworkIDGatewaysGatewayIDCellularRanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDGatewaysGatewayIDCellularRanParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDCellularRanParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDCellularRanParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularRanParamsWithContext creates a new GetLTENetworkIDGatewaysGatewayIDCellularRanParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDGatewaysGatewayIDCellularRanParamsWithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDCellularRanParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDCellularRanParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularRanParamsWithHTTPClient creates a new GetLTENetworkIDGatewaysGatewayIDCellularRanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDGatewaysGatewayIDCellularRanParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDCellularRanParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDCellularRanParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDCellularRanParams contains all the parameters to send to the API endpoint
for the get LTE network ID gateways gateway ID cellular ran operation typically these are written to a http.Request
*/
type GetLTENetworkIDGatewaysGatewayIDCellularRanParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get LTE network ID gateways gateway ID cellular ran params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDCellularRanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID gateways gateway ID cellular ran params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID gateways gateway ID cellular ran params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) WithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDCellularRanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID gateways gateway ID cellular ran params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID cellular ran params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDCellularRanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID cellular ran params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get LTE network ID gateways gateway ID cellular ran params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) WithGatewayID(gatewayID string) *GetLTENetworkIDGatewaysGatewayIDCellularRanParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get LTE network ID gateways gateway ID cellular ran params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get LTE network ID gateways gateway ID cellular ran params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) WithNetworkID(networkID string) *GetLTENetworkIDGatewaysGatewayIDCellularRanParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID gateways gateway ID cellular ran params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDGatewaysGatewayIDCellularRanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}