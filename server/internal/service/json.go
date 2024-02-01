package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ReadJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(data)
	if err != nil {
		return nil
	}

	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body msut only contain a single json value")
	}

	return nil
}
