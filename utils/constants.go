package utils

// HttpMethod https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods

const (
	HttpGet     HttpMethod = "GET"
	HttpPost    HttpMethod = "POST"
	HttpPut     HttpMethod = "PUT"
	HttpPatch   HttpMethod = "PATCH"
	HttpDelete  HttpMethod = "DELETE"
	HttpHead    HttpMethod = "HEAD"
	HttpConnect HttpMethod = "CONNECT"
	HttpOptions HttpMethod = "OPTIONS"
	HttpTrace   HttpMethod = "TRACE"
)

const (
	ActionAllow Action = "ALLOW"
	ActionDeny  Action = "DENY"
)
