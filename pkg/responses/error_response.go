package responses

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error          error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrUnsupportedMediaType() render.Renderer {
	return &ErrorResponse{
		Error:          nil,
		HTTPStatusCode: 415,
		StatusText:     "Unsupported Media Type",
		ErrorText:      "Unsupported Media Type",
	}
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrorResponse{
		Error:          err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
