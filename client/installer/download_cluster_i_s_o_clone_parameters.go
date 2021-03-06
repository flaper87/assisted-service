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

// NewDownloadClusterISOCloneParams creates a new DownloadClusterISOCloneParams object
// with the default values initialized.
func NewDownloadClusterISOCloneParams() *DownloadClusterISOCloneParams {
	var ()
	return &DownloadClusterISOCloneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadClusterISOCloneParamsWithTimeout creates a new DownloadClusterISOCloneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadClusterISOCloneParamsWithTimeout(timeout time.Duration) *DownloadClusterISOCloneParams {
	var ()
	return &DownloadClusterISOCloneParams{

		timeout: timeout,
	}
}

// NewDownloadClusterISOCloneParamsWithContext creates a new DownloadClusterISOCloneParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadClusterISOCloneParamsWithContext(ctx context.Context) *DownloadClusterISOCloneParams {
	var ()
	return &DownloadClusterISOCloneParams{

		Context: ctx,
	}
}

// NewDownloadClusterISOCloneParamsWithHTTPClient creates a new DownloadClusterISOCloneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadClusterISOCloneParamsWithHTTPClient(client *http.Client) *DownloadClusterISOCloneParams {
	var ()
	return &DownloadClusterISOCloneParams{
		HTTPClient: client,
	}
}

/*DownloadClusterISOCloneParams contains all the parameters to send to the API endpoint
for the download cluster i s o clone operation typically these are written to a http.Request
*/
type DownloadClusterISOCloneParams struct {

	/*ClusterID
	  The cluster whose ISO should be downloaded.

	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download cluster i s o clone params
func (o *DownloadClusterISOCloneParams) WithTimeout(timeout time.Duration) *DownloadClusterISOCloneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download cluster i s o clone params
func (o *DownloadClusterISOCloneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download cluster i s o clone params
func (o *DownloadClusterISOCloneParams) WithContext(ctx context.Context) *DownloadClusterISOCloneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download cluster i s o clone params
func (o *DownloadClusterISOCloneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download cluster i s o clone params
func (o *DownloadClusterISOCloneParams) WithHTTPClient(client *http.Client) *DownloadClusterISOCloneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download cluster i s o clone params
func (o *DownloadClusterISOCloneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the download cluster i s o clone params
func (o *DownloadClusterISOCloneParams) WithClusterID(clusterID strfmt.UUID) *DownloadClusterISOCloneParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the download cluster i s o clone params
func (o *DownloadClusterISOCloneParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadClusterISOCloneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
