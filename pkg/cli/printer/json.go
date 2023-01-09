package printer

import (
	"encoding/json"
	"io"
)

var _ Interface = NewJSON()

type JSON struct {
}

func NewJSON() *JSON {
	return &JSON{}
}

func (printer *JSON) Format() string {
	return "json"
}

func (printer *JSON) Print(obj any, w io.Writer) error {
	data, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return err
	}

	data = append(data, '\n')
	_, err = w.Write(data)
	return err
}
