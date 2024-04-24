package s11n

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type jsonEncocder interface {
	Encode(v any) error
	SetIndent(prefix, indent string)
	SetEscapeHTML(on bool)
}

type jsonDecoder interface {
	UseNumber()
	DisallowUnknownFields()
	Decode(v any) error
	Buffered() io.Reader
}

// JSONSerializer is a JSON serializer for Echo framework.
type JSONSerializer[E jsonEncocder, D jsonDecoder] struct {
	Encoder func(w io.Writer) E
	Decoder func(r io.Reader) D
}

// Serialize writes the JSON encoding of i to the response.
func (j *JSONSerializer[E, D]) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := j.Encoder(c.Response())
	return enc.Encode(i)
}

// Deserialize reads the JSON from the request and stores it in i.
func (j *JSONSerializer[E, D]) Deserialize(c echo.Context, i interface{}) error {
	err := json.NewDecoder(c.Request().Body).Decode(i)
	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
	}
	var se *json.SyntaxError
	if errors.As(err, &se) {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
	}
	return err
}
