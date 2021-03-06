// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/openshift/assisted-service/models"
)

// DownloadClusterISOHeadersCloneOKCode is the HTTP code returned for type DownloadClusterISOHeadersCloneOK
const DownloadClusterISOHeadersCloneOKCode int = 200

/*DownloadClusterISOHeadersCloneOK Success.

swagger:response downloadClusterISOHeadersCloneOK
*/
type DownloadClusterISOHeadersCloneOK struct {
	/*Size of the ISO in bytes

	 */
	ContentLength int64 `json:"Content-Length"`
}

// NewDownloadClusterISOHeadersCloneOK creates DownloadClusterISOHeadersCloneOK with default headers values
func NewDownloadClusterISOHeadersCloneOK() *DownloadClusterISOHeadersCloneOK {

	return &DownloadClusterISOHeadersCloneOK{}
}

// WithContentLength adds the contentLength to the download cluster i s o headers clone o k response
func (o *DownloadClusterISOHeadersCloneOK) WithContentLength(contentLength int64) *DownloadClusterISOHeadersCloneOK {
	o.ContentLength = contentLength
	return o
}

// SetContentLength sets the contentLength to the download cluster i s o headers clone o k response
func (o *DownloadClusterISOHeadersCloneOK) SetContentLength(contentLength int64) {
	o.ContentLength = contentLength
}

// WriteResponse to the client
func (o *DownloadClusterISOHeadersCloneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Length

	contentLength := swag.FormatInt64(o.ContentLength)
	if contentLength != "" {
		rw.Header().Set("Content-Length", contentLength)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DownloadClusterISOHeadersCloneBadRequestCode is the HTTP code returned for type DownloadClusterISOHeadersCloneBadRequest
const DownloadClusterISOHeadersCloneBadRequestCode int = 400

/*DownloadClusterISOHeadersCloneBadRequest Error.

swagger:response downloadClusterISOHeadersCloneBadRequest
*/
type DownloadClusterISOHeadersCloneBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterISOHeadersCloneBadRequest creates DownloadClusterISOHeadersCloneBadRequest with default headers values
func NewDownloadClusterISOHeadersCloneBadRequest() *DownloadClusterISOHeadersCloneBadRequest {

	return &DownloadClusterISOHeadersCloneBadRequest{}
}

// WithPayload adds the payload to the download cluster i s o headers clone bad request response
func (o *DownloadClusterISOHeadersCloneBadRequest) WithPayload(payload *models.Error) *DownloadClusterISOHeadersCloneBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o headers clone bad request response
func (o *DownloadClusterISOHeadersCloneBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOHeadersCloneBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterISOHeadersCloneUnauthorizedCode is the HTTP code returned for type DownloadClusterISOHeadersCloneUnauthorized
const DownloadClusterISOHeadersCloneUnauthorizedCode int = 401

/*DownloadClusterISOHeadersCloneUnauthorized Unauthorized.

swagger:response downloadClusterISOHeadersCloneUnauthorized
*/
type DownloadClusterISOHeadersCloneUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadClusterISOHeadersCloneUnauthorized creates DownloadClusterISOHeadersCloneUnauthorized with default headers values
func NewDownloadClusterISOHeadersCloneUnauthorized() *DownloadClusterISOHeadersCloneUnauthorized {

	return &DownloadClusterISOHeadersCloneUnauthorized{}
}

// WithPayload adds the payload to the download cluster i s o headers clone unauthorized response
func (o *DownloadClusterISOHeadersCloneUnauthorized) WithPayload(payload *models.InfraError) *DownloadClusterISOHeadersCloneUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o headers clone unauthorized response
func (o *DownloadClusterISOHeadersCloneUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOHeadersCloneUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterISOHeadersCloneForbiddenCode is the HTTP code returned for type DownloadClusterISOHeadersCloneForbidden
const DownloadClusterISOHeadersCloneForbiddenCode int = 403

/*DownloadClusterISOHeadersCloneForbidden Forbidden.

swagger:response downloadClusterISOHeadersCloneForbidden
*/
type DownloadClusterISOHeadersCloneForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadClusterISOHeadersCloneForbidden creates DownloadClusterISOHeadersCloneForbidden with default headers values
func NewDownloadClusterISOHeadersCloneForbidden() *DownloadClusterISOHeadersCloneForbidden {

	return &DownloadClusterISOHeadersCloneForbidden{}
}

// WithPayload adds the payload to the download cluster i s o headers clone forbidden response
func (o *DownloadClusterISOHeadersCloneForbidden) WithPayload(payload *models.InfraError) *DownloadClusterISOHeadersCloneForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o headers clone forbidden response
func (o *DownloadClusterISOHeadersCloneForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOHeadersCloneForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterISOHeadersCloneNotFoundCode is the HTTP code returned for type DownloadClusterISOHeadersCloneNotFound
const DownloadClusterISOHeadersCloneNotFoundCode int = 404

/*DownloadClusterISOHeadersCloneNotFound Error.

swagger:response downloadClusterISOHeadersCloneNotFound
*/
type DownloadClusterISOHeadersCloneNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterISOHeadersCloneNotFound creates DownloadClusterISOHeadersCloneNotFound with default headers values
func NewDownloadClusterISOHeadersCloneNotFound() *DownloadClusterISOHeadersCloneNotFound {

	return &DownloadClusterISOHeadersCloneNotFound{}
}

// WithPayload adds the payload to the download cluster i s o headers clone not found response
func (o *DownloadClusterISOHeadersCloneNotFound) WithPayload(payload *models.Error) *DownloadClusterISOHeadersCloneNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o headers clone not found response
func (o *DownloadClusterISOHeadersCloneNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOHeadersCloneNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterISOHeadersCloneMethodNotAllowedCode is the HTTP code returned for type DownloadClusterISOHeadersCloneMethodNotAllowed
const DownloadClusterISOHeadersCloneMethodNotAllowedCode int = 405

/*DownloadClusterISOHeadersCloneMethodNotAllowed Method Not Allowed.

swagger:response downloadClusterISOHeadersCloneMethodNotAllowed
*/
type DownloadClusterISOHeadersCloneMethodNotAllowed struct {
}

// NewDownloadClusterISOHeadersCloneMethodNotAllowed creates DownloadClusterISOHeadersCloneMethodNotAllowed with default headers values
func NewDownloadClusterISOHeadersCloneMethodNotAllowed() *DownloadClusterISOHeadersCloneMethodNotAllowed {

	return &DownloadClusterISOHeadersCloneMethodNotAllowed{}
}

// WriteResponse to the client
func (o *DownloadClusterISOHeadersCloneMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

// DownloadClusterISOHeadersCloneConflictCode is the HTTP code returned for type DownloadClusterISOHeadersCloneConflict
const DownloadClusterISOHeadersCloneConflictCode int = 409

/*DownloadClusterISOHeadersCloneConflict Error.

swagger:response downloadClusterISOHeadersCloneConflict
*/
type DownloadClusterISOHeadersCloneConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterISOHeadersCloneConflict creates DownloadClusterISOHeadersCloneConflict with default headers values
func NewDownloadClusterISOHeadersCloneConflict() *DownloadClusterISOHeadersCloneConflict {

	return &DownloadClusterISOHeadersCloneConflict{}
}

// WithPayload adds the payload to the download cluster i s o headers clone conflict response
func (o *DownloadClusterISOHeadersCloneConflict) WithPayload(payload *models.Error) *DownloadClusterISOHeadersCloneConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o headers clone conflict response
func (o *DownloadClusterISOHeadersCloneConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOHeadersCloneConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterISOHeadersCloneInternalServerErrorCode is the HTTP code returned for type DownloadClusterISOHeadersCloneInternalServerError
const DownloadClusterISOHeadersCloneInternalServerErrorCode int = 500

/*DownloadClusterISOHeadersCloneInternalServerError Error.

swagger:response downloadClusterISOHeadersCloneInternalServerError
*/
type DownloadClusterISOHeadersCloneInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterISOHeadersCloneInternalServerError creates DownloadClusterISOHeadersCloneInternalServerError with default headers values
func NewDownloadClusterISOHeadersCloneInternalServerError() *DownloadClusterISOHeadersCloneInternalServerError {

	return &DownloadClusterISOHeadersCloneInternalServerError{}
}

// WithPayload adds the payload to the download cluster i s o headers clone internal server error response
func (o *DownloadClusterISOHeadersCloneInternalServerError) WithPayload(payload *models.Error) *DownloadClusterISOHeadersCloneInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster i s o headers clone internal server error response
func (o *DownloadClusterISOHeadersCloneInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterISOHeadersCloneInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
