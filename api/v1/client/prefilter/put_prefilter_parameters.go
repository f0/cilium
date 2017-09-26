// Code generated by go-swagger; DO NOT EDIT.

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// NewPutPrefilterParams creates a new PutPrefilterParams object
// with the default values initialized.
func NewPutPrefilterParams() *PutPrefilterParams {
	var ()
	return &PutPrefilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPrefilterParamsWithTimeout creates a new PutPrefilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPrefilterParamsWithTimeout(timeout time.Duration) *PutPrefilterParams {
	var ()
	return &PutPrefilterParams{

		timeout: timeout,
	}
}

// NewPutPrefilterParamsWithContext creates a new PutPrefilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPrefilterParamsWithContext(ctx context.Context) *PutPrefilterParams {
	var ()
	return &PutPrefilterParams{

		Context: ctx,
	}
}

// NewPutPrefilterParamsWithHTTPClient creates a new PutPrefilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutPrefilterParamsWithHTTPClient(client *http.Client) *PutPrefilterParams {
	var ()
	return &PutPrefilterParams{
		HTTPClient: client,
	}
}

/*PutPrefilterParams contains all the parameters to send to the API endpoint
for the put prefilter operation typically these are written to a http.Request
*/
type PutPrefilterParams struct {

	/*CidrList
	  List of CIDRs for filter table

	*/
	CidrList *models.CIDRList

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put prefilter params
func (o *PutPrefilterParams) WithTimeout(timeout time.Duration) *PutPrefilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put prefilter params
func (o *PutPrefilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put prefilter params
func (o *PutPrefilterParams) WithContext(ctx context.Context) *PutPrefilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put prefilter params
func (o *PutPrefilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put prefilter params
func (o *PutPrefilterParams) WithHTTPClient(client *http.Client) *PutPrefilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put prefilter params
func (o *PutPrefilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCidrList adds the cidrList to the put prefilter params
func (o *PutPrefilterParams) WithCidrList(cidrList *models.CIDRList) *PutPrefilterParams {
	o.SetCidrList(cidrList)
	return o
}

// SetCidrList adds the cidrList to the put prefilter params
func (o *PutPrefilterParams) SetCidrList(cidrList *models.CIDRList) {
	o.CidrList = cidrList
}

// WriteToRequest writes these params to a swagger request
func (o *PutPrefilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CidrList == nil {
		o.CidrList = new(models.CIDRList)
	}

	if err := r.SetBodyParam(o.CidrList); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
