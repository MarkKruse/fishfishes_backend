package build

import (
	"bytes"
	"fmt"
	"github.com/magefile/mage/target"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/pkg/errors"
)

func ConvertToOpenApi3(inputFilePath, outputFilePath string) error {

	// only convert if the swagger file has changed
	modified, err := target.Path(outputFilePath, inputFilePath)
	if err != nil {
		return err
	}

	if !modified {
		return nil
	}

	fmt.Printf("Converting swagger file '%s' to OpenAPI 3.0 ... \n", inputFilePath)

	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://converter.swagger.io/api/convert", bytes.NewBuffer(input))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/yaml")
	req.Header.Set("Accept", "application/yaml")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	output, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile(outputFilePath, output, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func WithXGoTypes() (string, string) {
	return "client.tmpl", xGoTypeHandler
}

func GenerateSwaggerClient(api, packageName, outputFilePath string, withTemplate ...func() (string, string)) error {

	opts := codegen.Configuration{
		PackageName: packageName,
		Generate: codegen.GenerateOptions{
			Client: true,
			Models: true,
		},
		OutputOptions: codegen.OutputOptions{
			//Needs to be initialized, otherwise an error will occur when templates are added
			UserTemplates: map[string]string{},
		},
		Compatibility: codegen.CompatibilityOptions{
			//NGWDEV-2045 NGSDL needs this flag
			DisableFlattenAdditionalProperties: true,
		},
	}

	for _, template := range withTemplate {
		key, value := template()
		opts.OutputOptions.UserTemplates[key] = value
	}

	swagger, err := util.LoadSwagger(api)
	if err != nil {
		return errors.Wrap(err, "could not load api")
	}

	code, err := codegen.Generate(swagger, opts)
	if err != nil {
		return errors.Wrap(err, "error generating code")
	}

	_ = os.MkdirAll(outputFilePath, os.ModePerm)

	err = os.WriteFile(outputFilePath+"/client.go", []byte(code), 0644)
	if err != nil {
		return errors.Wrap(err, "error writing generating code")
	}

	return nil
}

func GenerateSwaggerTypes(api, packageName, outputFilePath string) error {

	opts := codegen.Configuration{
		PackageName: packageName,
		Generate: codegen.GenerateOptions{
			Client: false,
			Models: true,
		},
		Compatibility: codegen.CompatibilityOptions{
			//NGWDEV-2045 NGSDL needs this flag
			DisableFlattenAdditionalProperties: true,
		},
	}

	swagger, err := util.LoadSwagger(api)
	if err != nil {
		return errors.Wrap(err, "could not load api")
	}

	code, err := codegen.Generate(swagger, opts)
	if err != nil {
		return errors.Wrap(err, "error generating code")
	}

	_ = os.MkdirAll(outputFilePath, os.ModePerm)

	err = os.WriteFile(outputFilePath+"/types.go", []byte(code), 0644)
	if err != nil {
		return errors.Wrap(err, "error writing generating code")
	}

	return nil
}

func GenerateSwaggerServer(api, packageName, outputFilePath string) error {
	opts := codegen.Configuration{
		PackageName: packageName,
		Generate: codegen.GenerateOptions{
			ChiServer:    true,
			Models:       true,
			EmbeddedSpec: true,
		},
		OutputOptions: codegen.OutputOptions{
			SkipPrune: false,
			SkipFmt:   false,
			UserTemplates: map[string]string{
				"chi/chi-interface.tmpl":  chiInterface,
				"chi/chi-handler.tmpl":    chiHandler,
				"chi/chi-middleware.tmpl": chiMiddleware,
			},
		},
		Compatibility: codegen.CompatibilityOptions{
			//NGWDEV-2045 NGSDL needs this flag
			DisableFlattenAdditionalProperties: true,
		},
	}

	swagger, err := util.LoadSwagger(api)
	if err != nil {
		return errors.Wrap(err, "could not load api")
	}

	code, err := codegen.Generate(swagger, opts)
	if err != nil {
		return errors.Wrap(err, "error generating code")
	}

	_ = os.MkdirAll(outputFilePath, os.ModePerm)

	err = os.WriteFile(outputFilePath+"/server.go", []byte(code), 0644)
	if err != nil {
		return errors.Wrap(err, "error writing generating code")
	}

	return nil
}

func GenerateStringFromFile(filePath, outputFilePath, packageName, stringName string) error {

	b, err := os.ReadFile(filePath)
	if err != nil {
		return errors.Wrap(err, "could not read file")
	}

	bs := string(b)
	bs = strings.ReplaceAll(bs, "`", "\\`")

	_ = os.MkdirAll(outputFilePath, os.ModePerm)

	w, err := os.Create(outputFilePath + "/" + stringName + ".go")
	if err != nil {
		return errors.Wrap(err, "error writing generating code")
	}

	_, _ = fmt.Fprintf(w, "package %s\n\nconst %s = `%s`", packageName, stringName, bs)
	return nil
}

var xGoTypeHandler = `
// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
    // create a client with sane default values
    client := Client{
        Server: server,
    }
    // mutate client and add all optional params
    for _, o := range opts {
        if err := o(&client); err != nil {
            return nil, err
        }
    }
    // ensure the server URL always has a trailing slash
    if !strings.HasSuffix(client.Server, "/") {
        client.Server += "/"
    }
    // create httpClient, if not already present
    if client.Client == nil {
        client.Client = &http.Client{}
    }
    return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
{{range . -}}
{{$hasParams := .RequiresParamObject -}}
{{$pathParams := .PathParams -}}
{{$opid := .OperationId -}}
    // {{$opid}} request{{if .HasBody}} with any body{{end}}
    {{$opid}}{{if .HasBody}}WithBody{{end}}(ctx context.Context{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}{{if .HasBody}}, contentType string, body io.Reader{{end}}, reqEditors... RequestEditorFn) (*http.Response, error)
{{range .Bodies}}
    {{$opid}}{{.Suffix}}(ctx context.Context{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}, body {{$opid}}{{.NameTag}}RequestBody, reqEditors... RequestEditorFn) (*http.Response, error)
{{end}}{{/* range .Bodies */}}
{{end}}{{/* range . $opid := .OperationId */}}
}


{{/* Generate client methods */}}
{{range . -}}
{{$hasParams := .RequiresParamObject -}}
{{$pathParams := .PathParams -}}
{{$opid := .OperationId -}}

func (c *Client) {{$opid}}{{if .HasBody}}WithBody{{end}}(ctx context.Context{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}{{if .HasBody}}, contentType string, body io.Reader{{end}}, reqEditors... RequestEditorFn) (*http.Response, error) {
    req, err := New{{$opid}}Request{{if .HasBody}}WithBody{{end}}(c.Server{{genParamNames .PathParams}}{{if $hasParams}}, params{{end}}{{if .HasBody}}, contentType, body{{end}})
    if err != nil {
        return nil, err
    }
    req = req.WithContext(ctx)
    if err := c.applyEditors(ctx, req, reqEditors); err != nil {
        return nil, err
    }
    return c.Client.Do(req)
}

{{range .Bodies}}
func (c *Client) {{$opid}}{{.Suffix}}(ctx context.Context{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}, body {{$opid}}{{.NameTag}}RequestBody, reqEditors... RequestEditorFn) (*http.Response, error) {
    req, err := New{{$opid}}{{.Suffix}}Request(c.Server{{genParamNames $pathParams}}{{if $hasParams}}, params{{end}}, body)
    if err != nil {
        return nil, err
    }
    req = req.WithContext(ctx)
    if err := c.applyEditors(ctx, req, reqEditors); err != nil {
        return nil, err
    }
    return c.Client.Do(req)
}
{{end}}{{/* range .Bodies */}}
{{end}}

{{/* Generate request builders */}}
{{range .}}
{{$hasParams := .RequiresParamObject -}}
{{$pathParams := .PathParams -}}
{{$bodyRequired := .BodyRequired -}}
{{$opid := .OperationId -}}

{{range .Bodies}}
// New{{$opid}}Request{{.Suffix}} calls the generic {{$opid}} builder with {{.ContentType}} body
func New{{$opid}}Request{{.Suffix}}(server string{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}, body {{$opid}}{{.NameTag}}RequestBody) (*http.Request, error) {
    var bodyReader io.Reader
    buf, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    bodyReader = bytes.NewReader(buf)
    return New{{$opid}}RequestWithBody(server{{genParamNames $pathParams}}{{if $hasParams}}, params{{end}}, "{{.ContentType}}", bodyReader)
}
{{end}}

// New{{$opid}}Request{{if .HasBody}}WithBody{{end}} generates requests for {{$opid}}{{if .HasBody}} with any type of body{{end}}
func New{{$opid}}Request{{if .HasBody}}WithBody{{end}}(server string{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}{{if .HasBody}}, contentType string, body io.Reader{{end}}) (*http.Request, error) {
    var err error
{{range $paramIdx, $param := .PathParams}}
    var pathParam{{$paramIdx}} string
    {{if .IsPassThrough}}
    pathParam{{$paramIdx}} = {{.GoVariableName}}
    {{end}}
    {{if .IsJson}}
    var pathParamBuf{{$paramIdx}} []byte
    pathParamBuf{{$paramIdx}}, err = json.Marshal({{.GoVariableName}})
    if err != nil {
        return nil, err
    }
    pathParam{{$paramIdx}} = string(pathParamBuf{{$paramIdx}})
    {{end}}
    {{if .IsStyled}}
    pathParam{{$paramIdx}}, err = runtime.StyleParamWithLocation("{{.Style}}", {{.Explode}}, "{{.ParamName}}", runtime.ParamLocationPath, {{.GoVariableName}})
    if err != nil {
        return nil, err
    }
    {{end}}
{{end}}
    serverURL, err := url.Parse(server)
    if err != nil {
        return nil, err
    }

    operationPath := fmt.Sprintf("{{genParamFmtString .Path}}"{{range $paramIdx, $param := .PathParams}}, pathParam{{$paramIdx}}{{end}})
    if operationPath[0] == '/' {
        operationPath = "." + operationPath
    }

    queryURL, err := serverURL.Parse(operationPath)
    if err != nil {
        return nil, err
    }

{{if .QueryParams}}
    queryValues := queryURL.Query()
{{range $paramIdx, $param := .QueryParams}}
    {{if not .Required}} if params.{{.GoName}} != nil { {{end}}
    {{if .IsPassThrough}}
    queryValues.Add("{{.ParamName}}", {{if not .Required}}*{{end}}params.{{.GoName}})
    {{end}}
    {{if .IsJson}}
    if queryParamBuf, err := json.Marshal({{if not .Required}}*{{end}}params.{{.GoName}}); err != nil {
        return nil, err
    } else {
        queryValues.Add("{{.ParamName}}", string(queryParamBuf))
    }

    {{end}}
    {{if .IsStyled}}
    if queryFrag, err := runtime.StyleParamWithLocation("{{.Style}}", {{.Explode}}, "{{.ParamName}}", runtime.ParamLocationQuery, {{if not .Required}}*{{end}}params.{{.GoName}}); err != nil {
        return nil, err
    } else if parsed, err := url.ParseQuery(queryFrag); err != nil {
       return nil, err
    } else {
       for k, v := range parsed {
           for _, v2 := range v {
               queryValues.Add(k, v2)
           }
       }
    }
    {{end}}
    {{if not .Required}}}{{end}}
{{end}}
    queryURL.RawQuery = queryValues.Encode()
{{end}}{{/* if .QueryParams */}}
    req, err := http.NewRequest("{{.Method}}", queryURL.String(), {{if .HasBody}}body{{else}}nil{{end}})
    if err != nil {
        return nil, err
    }

    {{if .HasBody}}req.Header.Add("Content-Type", contentType){{end}}
{{range $paramIdx, $param := .HeaderParams}}
    {{if not .Required}} if params.{{.GoName}} != nil { {{end}}
    var headerParam{{$paramIdx}} string
    {{if .IsPassThrough}}
    headerParam{{$paramIdx}} = {{if not .Required}}*{{end}}params.{{.GoName}}
    {{end}}
    {{if .IsJson}}
    var headerParamBuf{{$paramIdx}} []byte
    headerParamBuf{{$paramIdx}}, err = json.Marshal({{if not .Required}}*{{end}}params.{{.GoName}})
    if err != nil {
        return nil, err
    }
    headerParam{{$paramIdx}} = string(headerParamBuf{{$paramIdx}})
    {{end}}
    {{if .IsStyled}}
    headerParam{{$paramIdx}}, err = runtime.StyleParamWithLocation("{{.Style}}", {{.Explode}}, "{{.ParamName}}", runtime.ParamLocationHeader, {{if not .Required}}*{{end}}params.{{.GoName}})
    if err != nil {
        return nil, err
    }
    {{end}}
    req.Header.Set("{{.ParamName}}", headerParam{{$paramIdx}})
    {{if not .Required}}}{{end}}
{{end}}

{{range $paramIdx, $param := .CookieParams}}
    {{if not .Required}} if params.{{.GoName}} != nil { {{end}}
    var cookieParam{{$paramIdx}} string
    {{if .IsPassThrough}}
    cookieParam{{$paramIdx}} = {{if not .Required}}*{{end}}params.{{.GoName}}
    {{end}}
    {{if .IsJson}}
    var cookieParamBuf{{$paramIdx}} []byte
    cookieParamBuf{{$paramIdx}}, err = json.Marshal({{if not .Required}}*{{end}}params.{{.GoName}})
    if err != nil {
        return nil, err
    }
    cookieParam{{$paramIdx}} = url.QueryEscape(string(cookieParamBuf{{$paramIdx}}))
    {{end}}
    {{if .IsStyled}}
    cookieParam{{$paramIdx}}, err = runtime.StyleParamWithLocation("simple", {{.Explode}}, "{{.ParamName}}", runtime.ParamLocationCookie, {{if not .Required}}*{{end}}params.{{.GoName}})
    if err != nil {
        return nil, err
    }
    {{end}}
    cookie{{$paramIdx}} := &http.Cookie{
        Name:"{{.ParamName}}",
        Value:cookieParam{{$paramIdx}},
    }
    req.AddCookie(cookie{{$paramIdx}})
    {{if not .Required}}}{{end}}
{{end}}
    return req, nil
}

{{end}}{{/* Range */}}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
    for _, r := range c.RequestEditors {
        if err := r(ctx, req); err != nil {
            return err
        }
    }
    for _, r := range additionalEditors {
        if err := r(ctx, req); err != nil {
            return err
        }
    }
    return nil
}

type BoolwithStringCompatibility bool

func (m *BoolwithStringCompatibility) UnmarshalJSON(b []byte) error {

	value := string(b)
	value = strings.Replace(value, "\"", "", -1)

	if strings.ToLower(value) == "true" {
		*m = true
	}
	if strings.ToLower(value) == "false" {
		*m = false
	}

	return nil
}

type IntegerwithStringCompatibility int

func (m *IntegerwithStringCompatibility) UnmarshalJSON(b []byte) error {

	value := string(b)
	value = strings.Replace(value, "\"", "", -1)

	integer, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*m = IntegerwithStringCompatibility(integer)

	return nil
}

type FloatwithStringCompatibility float32

func (m *FloatwithStringCompatibility) UnmarshalJSON(b []byte) error {

	value := string(b)
	value = strings.Replace(value, "\"", "", -1)

	float, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}
	*m = FloatwithStringCompatibility(float)
	return nil
}
`

var chiInterface = `
	// ServerInterface represents all server handlers.
	type ServerInterface interface {
	{{range .}}{{.SummaryAsComment }}
	// ({{.Method}} {{.Path}})
	{{.OperationId}}(w http.ResponseWriter, r *http.Request{{genParamArgs .PathParams}}{{if .RequiresParamObject}}, params {{.OperationId}}Params{{end}})
	{{end}}
	}

	type ServerWithMiddleware struct {
	ServerInterface
	{{range .}}// {{.Summary | stripNewLines }} ({{.Method}} {{.Path}})
	{{.OperationId}}Middlewares chi.Middlewares
	{{end}}
	}

	func NewServerWithMiddleware(si ServerInterface) ServerWithMiddleware {
		return ServerWithMiddleware{
			ServerInterface: si,
		}
	}
`

var chiHandler = `
// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerWithMiddleware) http.Handler {
  return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerWithMiddleware, r chi.Router) http.Handler {
  {{if .}}wrapper := ServerInterfaceWrapper{
    Handler: si,
    ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
        http.Error(w, err.Error(), http.StatusBadRequest)
    },
  }
  {{end}}
  {{range .}}r.Group(func(r chi.Router) {
    r.With(si.{{.OperationId}}Middlewares...).{{.Method | lower | title }}("{{.Path | swaggerUriToChiUri}}", wrapper.{{.OperationId}})
  })
  {{end}}
  return r
}
`

var chiMiddleware = `
// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
    Handler ServerInterface
    HandlerMiddlewares []MiddlewareFunc
    ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc
{{range .}}{{$opid := .OperationId}}
// {{$opid}} operation middleware
func (siw *ServerInterfaceWrapper) {{$opid}}(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()
  {{if or .RequiresParamObject (gt (len .PathParams) 0) }}
  var err error
  {{end}}
  {{range .PathParams}}// ------------- Path parameter "{{.ParamName}}" -------------
  var {{$varName := .GoVariableName}}{{$varName}} {{.TypeDef}}
  {{if .IsPassThrough}}
  {{$varName}} = chi.URLParam(r, "{{.ParamName}}")
  {{end}}
  {{if .IsJson}}
  err = json.Unmarshal([]byte(chi.URLParam(r, "{{.ParamName}}")), &{{$varName}})
  if err != nil {
    siw.ErrorHandlerFunc(w, r, &UnmarshalingParamError{ParamName: "{{.ParamName}}", Err: err})
    return
  }
  {{end}}
  {{if .IsStyled}}
  err = runtime.BindStyledParameterWithLocation("{{.Style}}",{{.Explode}}, "{{.ParamName}}", runtime.ParamLocationPath, chi.URLParam(r, "{{.ParamName}}"), &{{$varName}})
  if err != nil {
    siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "{{.ParamName}}", Err: err})
    return
  }
  {{end}}
  {{end}}
{{range .SecurityDefinitions}}
  ctx = context.WithValue(ctx, {{.ProviderName | ucFirst}}Scopes, {{toStringArray .Scopes}})
{{end}}
  {{if .RequiresParamObject}}
    // Parameter object where we will unmarshal all parameters from the context
    var params {{.OperationId}}Params
    {{range $paramIdx, $param := .QueryParams}}// ------------- {{if .Required}}Required{{else}}Optional{{end}} query parameter "{{.ParamName}}" -------------
      if paramValue := r.URL.Query().Get("{{.ParamName}}"); paramValue != "" {
      {{if .IsPassThrough}}
        params.{{.GoName}} = {{if not .Required}}&{{end}}paramValue
      {{end}}
      {{if .IsJson}}
        var value {{.TypeDef}}
        err = json.Unmarshal([]byte(paramValue), &value)
        if err != nil {
          siw.ErrorHandlerFunc(w, r, &UnmarshalingParamError{ParamName: "{{.ParamName}}", Err: err})
          return
        }
        params.{{.GoName}} = {{if not .Required}}&{{end}}value
      {{end}}
      }{{if .Required}} else {
          siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "{{.ParamName}}"})
          return
      }{{end}}
      {{if .IsStyled}}
      err = runtime.BindQueryParameter("{{.Style}}", {{.Explode}}, {{.Required}}, "{{.ParamName}}", r.URL.Query(), &params.{{.GoName}})
      if err != nil {
        siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "{{.ParamName}}", Err: err})
        return
      }
      {{end}}
  {{end}}
    {{if .HeaderParams}}
      headers := r.Header
      {{range .HeaderParams}}// ------------- {{if .Required}}Required{{else}}Optional{{end}} header parameter "{{.ParamName}}" -------------
        if valueList, found := headers[http.CanonicalHeaderKey("{{.ParamName}}")]; found {
          var {{.GoName}} {{.TypeDef}}
          n := len(valueList)
          if n != 1 {
            siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "{{.ParamName}}", Count: n})
            return
          }
        {{if .IsPassThrough}}
          params.{{.GoName}} = {{if not .Required}}&{{end}}valueList[0]
        {{end}}
        {{if .IsJson}}
          err = json.Unmarshal([]byte(valueList[0]), &{{.GoName}})
          if err != nil {
            siw.ErrorHandlerFunc(w, r, &UnmarshalingParamError{ParamName: "{{.ParamName}}", Err: err})
            return
          }
        {{end}}
        {{if .IsStyled}}
          err = runtime.BindStyledParameterWithLocation("{{.Style}}",{{.Explode}}, "{{.ParamName}}", runtime.ParamLocationHeader, valueList[0], &{{.GoName}})
          if err != nil {
            siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "{{.ParamName}}", Err: err})
            return
          }
        {{end}}
          params.{{.GoName}} = {{if not .Required}}&{{end}}{{.GoName}}
        } {{if .Required}}else {
            err := fmt.Errorf("Header parameter {{.ParamName}} is required, but not found")
            siw.ErrorHandlerFunc(w, r, &RequiredHeaderError{ParamName: "{{.ParamName}}", Err: err})
            return
        }{{end}}
      {{end}}
    {{end}}
    {{range .CookieParams}}
      var cookie *http.Cookie
      if cookie, err = r.Cookie("{{.ParamName}}"); err == nil {
      {{- if .IsPassThrough}}
        params.{{.GoName}} = {{if not .Required}}&{{end}}cookie.Value
      {{end}}
      {{- if .IsJson}}
        var value {{.TypeDef}}
        var decoded string
        decoded, err := url.QueryUnescape(cookie.Value)
        if err != nil {
          err = fmt.Errorf("Error unescaping cookie parameter '{{.ParamName}}'")
          siw.ErrorHandlerFunc(w, r, &UnescapedCookieParamError{ParamName: "{{.ParamName}}", Err: err})
          return
        }
        err = json.Unmarshal([]byte(decoded), &value)
        if err != nil {
          siw.ErrorHandlerFunc(w, r, &UnmarshalingParamError{ParamName: "{{.ParamName}}", Err: err})
          return
        }
        params.{{.GoName}} = {{if not .Required}}&{{end}}value
      {{end}}
      {{- if .IsStyled}}
        var value {{.TypeDef}}
        err = runtime.BindStyledParameter("simple",{{.Explode}}, "{{.ParamName}}", cookie.Value, &value)
        if err != nil {
          siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "{{.ParamName}}", Err: err})
          return
        }
        params.{{.GoName}} = {{if not .Required}}&{{end}}value
      {{end}}
      }
      {{- if .Required}} else {
        siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "{{.ParamName}}"})
        return
      }
      {{- end}}
    {{end}}
  {{end}}
  var handler = func(w http.ResponseWriter, r *http.Request) {
    siw.Handler.{{.OperationId}}(w, r{{genParamNames .PathParams}}{{if .RequiresParamObject}}, params{{end}})
}
  for _, middleware := range siw.HandlerMiddlewares {
    handler = middleware(handler)
  }
  handler(w, r.WithContext(ctx))
}
{{end}}
type UnescapedCookieParamError struct {
    ParamName string
    Err error
}
func (e *UnescapedCookieParamError) Error() string {
    return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}
func (e *UnescapedCookieParamError) Unwrap() error {
    return e.Err
}
type UnmarshalingParamError struct {
    ParamName string
    Err error
}
func (e *UnmarshalingParamError) Error() string {
    return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}
func (e *UnmarshalingParamError) Unwrap() error {
    return e.Err
}
type RequiredParamError struct {
    ParamName string
}
func (e *RequiredParamError) Error() string {
    return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}
type RequiredHeaderError struct {
    ParamName string
    Err error
}
func (e *RequiredHeaderError) Error() string {
    return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}
func (e *RequiredHeaderError) Unwrap() error {
    return e.Err
}
type InvalidParamFormatError struct {
    ParamName string
    Err error
}
func (e *InvalidParamFormatError) Error() string {
    return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}
func (e *InvalidParamFormatError) Unwrap() error {
    return e.Err
}
type TooManyValuesForParamError struct {
    ParamName string
    Count int
}
func (e *TooManyValuesForParamError) Error() string {
    return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}
`
