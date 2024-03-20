package iconst

// MIME types
const (
	MIMEApplicationJSON            = "application/json"
	MIMEApplicationJSONCharsetUTF8 = "application/json; charset=utf-8"
	MIMEApplicationXML             = "application/xml"
	MIMEApplicationXMLCharsetUTF8  = "application/xml; charset=utf-8"
	MIMETextHTML                   = "text/html"
	MIMETextHTMLCharsetUTF8        = "text/html; charset=utf-8"
	MIMETextPlain                  = "text/plain"
	MIMETextPlainCharsetUTF8       = "text/plain; charset=utf-8"
	MIMEOctetStream                = "application/octet-stream"
	MIMEApplicationForm            = "application/x-www-form-urlencoded"
)

// HTTP Header Fields
const (
	HeaderAuthorization  = "Authorization"    // Requests
	HeaderContentType    = "Content-Type"     // Requests, Responses
	HeaderUserAgent      = "User-Agent"       // Requests
	HeaderXRequestID     = "X-Request-Id"     // Requests
	HeaderXForwardedFor  = "X-Forwarded-For"  // Requests
	HeaderXForwardedHost = "X-Forwarded-Host" // Requests
)
