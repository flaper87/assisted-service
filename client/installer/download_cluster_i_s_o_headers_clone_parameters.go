// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDownloadClusterISOHeadersCloneParams creates a new DownloadClusterISOHeadersCloneParams object
// with the default values initialized.
func NewDownloadClusterISOHeadersCloneParams() *DownloadClusterISOHeadersCloneParams {
	var ()
	return &DownloadClusterISOHeadersCloneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadClusterISOHeadersCloneParamsWithTimeout creates a new DownloadClusterISOHeadersCloneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadClusterISOHeadersCloneParamsWithTimeout(timeout time.Duration) *DownloadClusterISOHeadersCloneParams {
	var ()
	return &DownloadClusterISOHeadersCloneParams{

		timeout: timeout,
	}
}

// NewDownloadClusterISOHeadersCloneParamsWithContext creates a new DownloadClusterISOHeadersCloneParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadClusterISOHeadersCloneParamsWithContext(ctx context.Context) *DownloadClusterISOHeadersCloneParams {
	var ()
	return &DownloadClusterISOHeadersCloneParams{

		Context: ctx,
	}
}

// NewDownloadClusterISOHeadersCloneParamsWithHTTPClient creates a new DownloadClusterISOHeadersCloneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadClusterISOHeadersCloneParamsWithHTTPClient(client *http.Client) *DownloadClusterISOHeadersCloneParams {
	var ()
	return &DownloadClusterISOHeadersCloneParams{
		HTTPClient: client,
	}
}

/*DownloadClusterISOHeadersCloneParams contains all the parameters to send to the API endpoint
for the download cluster i s o headers clone operation typically these are written to a http.Request
*/
type DownloadClusterISOHeadersCloneParams struct {

	/*ClusterID
	  The cluster whose ISO headers should be downloaded.

	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download cluster i s o headers clone params
func (o *DownloadClusterISOHeadersCloneParams) WithTimeout(timeout time.Duration) *DownloadClusterISOHeadersCloneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download cluster i s o headers clone params
func (o *DownloadClusterISOHeadersCloneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download cluster i s o headers clone params
func (o *DownloadClusterISOHeadersCloneParams) WithContext(ctx context.Context) *DownloadClusterISOHeadersCloneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download cluster i s o headers clone params
func (o *DownloadClusterISOHeadersCloneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download cluster i s o headers clone params
func (o *DownloadClusterISOHeadersCloneParams) WithHTTPClient(client *http.Client) *DownloadClusterISOHeadersCloneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download cluster i s o headers clone params
func (o *DownloadClusterISOHeadersCloneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the download cluster i s o headers clone params
func (o *DownloadClusterISOHeadersCloneParams) WithClusterID(clusterID strfmt.UUID) *DownloadClusterISOHeadersCloneParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the download cluster i s o headers clone params
func (o *DownloadClusterISOHeadersCloneParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadClusterISOHeadersCloneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
