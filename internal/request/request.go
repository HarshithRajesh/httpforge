package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

func (r *RequestLine) ValidHttp() bool {
	return r.HttpVersion == "HTTP/1.1"
}

var ERROR_MALFORMED_REQUEST_LINE = fmt.Errorf("malformed request-line")
var ERROR_UNSUPPORTED_HTTP_VERSION = fmt.Errorf("error unsupported http version")

var SEPERATOR = "\r\n"

func parseRequestLine(b string) (*RequestLine, string, error) {
	idx := strings.Index(b, SEPERATOR)
	if idx == -1 {
		return nil, b, nil
	}
	startLine := b[:idx]
	restOfMessage := b[idx+len(SEPERATOR):]
	parts := strings.Split(startLine, " ")
	if len(parts) != 3 {
		return nil, restOfMessage, ERROR_MALFORMED_REQUEST_LINE
	}
	rl := &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   parts[2],
	}

	if !rl.ValidHttp() {
		return nil, restOfMessage, ERROR_UNSUPPORTED_HTTP_VERSION
	}
	return rl, restOfMessage, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Join(
			fmt.Errorf("unable to io.ReadAll: "),
			err,
		)
	}

	str := string(data)

	rl, _, err := parseRequestLine(str)

	return &Request{
		RequestLine: *rl,
	}, err

}
