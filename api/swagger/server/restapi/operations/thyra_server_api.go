// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewThyraServerAPI creates a new ThyraServer instance
func NewThyraServerAPI(spec *loads.Document) *ThyraServerAPI {
	return &ThyraServerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		BinProducer: runtime.ByteStreamProducer(),
		CSSProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("css producer has not yet been implemented")
		}),
		HTMLProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("html producer has not yet been implemented")
		}),
		JSONProducer: runtime.JSONProducer(),
		TextWebpProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("textWebp producer has not yet been implemented")
		}),
		TxtProducer: runtime.TextProducer(),

		BrowseHandler: BrowseHandlerFunc(func(params BrowseParams) middleware.Responder {
			return middleware.NotImplemented("operation Browse has not yet been implemented")
		}),
		CmdExecuteFunctionHandler: CmdExecuteFunctionHandlerFunc(func(params CmdExecuteFunctionParams) middleware.Responder {
			return middleware.NotImplemented("operation CmdExecuteFunction has not yet been implemented")
		}),
		KpiHandler: KpiHandlerFunc(func(params KpiParams) middleware.Responder {
			return middleware.NotImplemented("operation Kpi has not yet been implemented")
		}),
		MgmtWalletCreateHandler: MgmtWalletCreateHandlerFunc(func(params MgmtWalletCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation MgmtWalletCreate has not yet been implemented")
		}),
		MgmtWalletDeleteHandler: MgmtWalletDeleteHandlerFunc(func(params MgmtWalletDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation MgmtWalletDelete has not yet been implemented")
		}),
		MgmtWalletGetHandler: MgmtWalletGetHandlerFunc(func(params MgmtWalletGetParams) middleware.Responder {
			return middleware.NotImplemented("operation MgmtWalletGet has not yet been implemented")
		}),
		MgmtWalletGetOneHandler: MgmtWalletGetOneHandlerFunc(func(params MgmtWalletGetOneParams) middleware.Responder {
			return middleware.NotImplemented("operation MgmtWalletGetOne has not yet been implemented")
		}),
		MgmtWalletImportHandler: MgmtWalletImportHandlerFunc(func(params MgmtWalletImportParams) middleware.Responder {
			return middleware.NotImplemented("operation MgmtWalletImport has not yet been implemented")
		}),
		MyDomainsHandler: MyDomainsHandlerFunc(func(params MyDomainsParams) middleware.Responder {
			return middleware.NotImplemented("operation MyDomains has not yet been implemented")
		}),
		ThyraWalletHandler: ThyraWalletHandlerFunc(func(params ThyraWalletParams) middleware.Responder {
			return middleware.NotImplemented("operation ThyraWallet has not yet been implemented")
		}),
		ThyraWebsiteCreatorHandler: ThyraWebsiteCreatorHandlerFunc(func(params ThyraWebsiteCreatorParams) middleware.Responder {
			return middleware.NotImplemented("operation ThyraWebsiteCreator has not yet been implemented")
		}),
		WebsiteCreatorPrepareHandler: WebsiteCreatorPrepareHandlerFunc(func(params WebsiteCreatorPrepareParams) middleware.Responder {
			return middleware.NotImplemented("operation WebsiteCreatorPrepare has not yet been implemented")
		}),
		WebsiteCreatorUploadHandler: WebsiteCreatorUploadHandlerFunc(func(params WebsiteCreatorUploadParams) middleware.Responder {
			return middleware.NotImplemented("operation WebsiteCreatorUpload has not yet been implemented")
		}),
	}
}

/*ThyraServerAPI Thyra HTTP server API. */
type ThyraServerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// BinProducer registers a producer for the following mime types:
	//   - image/png
	BinProducer runtime.Producer
	// CSSProducer registers a producer for the following mime types:
	//   - text/css
	CSSProducer runtime.Producer
	// HTMLProducer registers a producer for the following mime types:
	//   - text/html
	HTMLProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer
	// TextWebpProducer registers a producer for the following mime types:
	//   - text/webp
	TextWebpProducer runtime.Producer
	// TxtProducer registers a producer for the following mime types:
	//   - application/javascript
	TxtProducer runtime.Producer

	// BrowseHandler sets the operation handler for the browse operation
	BrowseHandler BrowseHandler
	// CmdExecuteFunctionHandler sets the operation handler for the cmd execute function operation
	CmdExecuteFunctionHandler CmdExecuteFunctionHandler
	// KpiHandler sets the operation handler for the kpi operation
	KpiHandler KpiHandler
	// MgmtWalletCreateHandler sets the operation handler for the mgmt wallet create operation
	MgmtWalletCreateHandler MgmtWalletCreateHandler
	// MgmtWalletDeleteHandler sets the operation handler for the mgmt wallet delete operation
	MgmtWalletDeleteHandler MgmtWalletDeleteHandler
	// MgmtWalletGetHandler sets the operation handler for the mgmt wallet get operation
	MgmtWalletGetHandler MgmtWalletGetHandler
	// MgmtWalletGetOneHandler sets the operation handler for the mgmt wallet get one operation
	MgmtWalletGetOneHandler MgmtWalletGetOneHandler
	// MgmtWalletImportHandler sets the operation handler for the mgmt wallet import operation
	MgmtWalletImportHandler MgmtWalletImportHandler
	// MyDomainsHandler sets the operation handler for the my domains operation
	MyDomainsHandler MyDomainsHandler
	// ThyraWalletHandler sets the operation handler for the thyra wallet operation
	ThyraWalletHandler ThyraWalletHandler
	// ThyraWebsiteCreatorHandler sets the operation handler for the thyra website creator operation
	ThyraWebsiteCreatorHandler ThyraWebsiteCreatorHandler
	// WebsiteCreatorPrepareHandler sets the operation handler for the website creator prepare operation
	WebsiteCreatorPrepareHandler WebsiteCreatorPrepareHandler
	// WebsiteCreatorUploadHandler sets the operation handler for the website creator upload operation
	WebsiteCreatorUploadHandler WebsiteCreatorUploadHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *ThyraServerAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *ThyraServerAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *ThyraServerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ThyraServerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ThyraServerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ThyraServerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ThyraServerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ThyraServerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ThyraServerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ThyraServerAPI
func (o *ThyraServerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}
	if o.CSSProducer == nil {
		unregistered = append(unregistered, "CSSProducer")
	}
	if o.HTMLProducer == nil {
		unregistered = append(unregistered, "HTMLProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.TextWebpProducer == nil {
		unregistered = append(unregistered, "TextWebpProducer")
	}
	if o.TxtProducer == nil {
		unregistered = append(unregistered, "TxtProducer")
	}

	if o.BrowseHandler == nil {
		unregistered = append(unregistered, "BrowseHandler")
	}
	if o.CmdExecuteFunctionHandler == nil {
		unregistered = append(unregistered, "CmdExecuteFunctionHandler")
	}
	if o.KpiHandler == nil {
		unregistered = append(unregistered, "KpiHandler")
	}
	if o.MgmtWalletCreateHandler == nil {
		unregistered = append(unregistered, "MgmtWalletCreateHandler")
	}
	if o.MgmtWalletDeleteHandler == nil {
		unregistered = append(unregistered, "MgmtWalletDeleteHandler")
	}
	if o.MgmtWalletGetHandler == nil {
		unregistered = append(unregistered, "MgmtWalletGetHandler")
	}
	if o.MgmtWalletGetOneHandler == nil {
		unregistered = append(unregistered, "MgmtWalletGetOneHandler")
	}
	if o.MgmtWalletImportHandler == nil {
		unregistered = append(unregistered, "MgmtWalletImportHandler")
	}
	if o.MyDomainsHandler == nil {
		unregistered = append(unregistered, "MyDomainsHandler")
	}
	if o.ThyraWalletHandler == nil {
		unregistered = append(unregistered, "ThyraWalletHandler")
	}
	if o.ThyraWebsiteCreatorHandler == nil {
		unregistered = append(unregistered, "ThyraWebsiteCreatorHandler")
	}
	if o.WebsiteCreatorPrepareHandler == nil {
		unregistered = append(unregistered, "WebsiteCreatorPrepareHandler")
	}
	if o.WebsiteCreatorUploadHandler == nil {
		unregistered = append(unregistered, "WebsiteCreatorUploadHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ThyraServerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ThyraServerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *ThyraServerAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *ThyraServerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *ThyraServerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "image/png":
			result["image/png"] = o.BinProducer
		case "text/css":
			result["text/css"] = o.CSSProducer
		case "text/html":
			result["text/html"] = o.HTMLProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		case "text/webp":
			result["text/webp"] = o.TextWebpProducer
		case "application/javascript":
			result["application/javascript"] = o.TxtProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ThyraServerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the thyra server API
func (o *ThyraServerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ThyraServerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/browse/{address}/{resource}"] = NewBrowse(o.context, o.BrowseHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/cmd/executeFunction"] = NewCmdExecuteFunction(o.context, o.CmdExecuteFunctionHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/kpi"] = NewKpi(o.context, o.KpiHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mgmt/wallet"] = NewMgmtWalletCreate(o.context, o.MgmtWalletCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/mgmt/wallet/{nickname}"] = NewMgmtWalletDelete(o.context, o.MgmtWalletDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/mgmt/wallet"] = NewMgmtWalletGet(o.context, o.MgmtWalletGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/mgmt/wallet/{nickname}"] = NewMgmtWalletGetOne(o.context, o.MgmtWalletGetOneHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/mgmt/wallet"] = NewMgmtWalletImport(o.context, o.MgmtWalletImportHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/my/domains/{nickname}"] = NewMyDomains(o.context, o.MyDomainsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/thyra/wallet/{resource}"] = NewThyraWallet(o.context, o.ThyraWalletHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/thyra/websiteCreator/{resource}"] = NewThyraWebsiteCreator(o.context, o.ThyraWebsiteCreatorHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/websiteCreator/prepare"] = NewWebsiteCreatorPrepare(o.context, o.WebsiteCreatorPrepareHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/websiteCreator/upload"] = NewWebsiteCreatorUpload(o.context, o.WebsiteCreatorUploadHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ThyraServerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ThyraServerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ThyraServerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ThyraServerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *ThyraServerAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
