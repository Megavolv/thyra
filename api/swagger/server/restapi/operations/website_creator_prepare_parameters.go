// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// WebsiteCreatorPrepareMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var WebsiteCreatorPrepareMaxParseMemory int64 = 32 << 20

// NewWebsiteCreatorPrepareParams creates a new WebsiteCreatorPrepareParams object
//
// There are no default values defined in the spec.
func NewWebsiteCreatorPrepareParams() WebsiteCreatorPrepareParams {

	return WebsiteCreatorPrepareParams{}
}

// WebsiteCreatorPrepareParams contains all the bound params for the website creator prepare operation
// typically these are obtained from a http.Request
//
// swagger:parameters websiteCreatorPrepare
type WebsiteCreatorPrepareParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Name of the Wallet in which the website will be deployed.
	  Required: true
	  In: formData
	*/
	Nickname string
	/*URL without dot (.), upper case and special characters
	  Required: true
	  Pattern: ^[a-z0-9]+$
	  In: formData
	*/
	URL string
	/*Website contents in a ZIP file.
	  Required: true
	  In: formData
	*/
	Zipfile io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewWebsiteCreatorPrepareParams() beforehand.
func (o *WebsiteCreatorPrepareParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(WebsiteCreatorPrepareMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	fdNickname, fdhkNickname, _ := fds.GetOK("nickname")
	if err := o.bindNickname(fdNickname, fdhkNickname, route.Formats); err != nil {
		res = append(res, err)
	}

	fdURL, fdhkURL, _ := fds.GetOK("url")
	if err := o.bindURL(fdURL, fdhkURL, route.Formats); err != nil {
		res = append(res, err)
	}

	zipfile, zipfileHeader, err := r.FormFile("zipfile")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "zipfile", err))
	} else if err := o.bindZipfile(zipfile, zipfileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.Zipfile = &runtime.File{Data: zipfile, Header: zipfileHeader}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNickname binds and validates parameter Nickname from formData.
func (o *WebsiteCreatorPrepareParams) bindNickname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("nickname", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("nickname", "formData", raw); err != nil {
		return err
	}
	o.Nickname = raw

	return nil
}

// bindURL binds and validates parameter URL from formData.
func (o *WebsiteCreatorPrepareParams) bindURL(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("url", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("url", "formData", raw); err != nil {
		return err
	}
	o.URL = raw

	if err := o.validateURL(formats); err != nil {
		return err
	}

	return nil
}

// validateURL carries on validations for parameter URL
func (o *WebsiteCreatorPrepareParams) validateURL(formats strfmt.Registry) error {

	if err := validate.Pattern("url", "formData", o.URL, `^[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// bindZipfile binds file parameter Zipfile.
//
// The only supported validations on files are MinLength and MaxLength
func (o *WebsiteCreatorPrepareParams) bindZipfile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}
