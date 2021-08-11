package data

import (
	"encoding/json"
	"io"
)	

func (d *Response_wrapper_structure) Response_wrapper_structureToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *Strike_Meta_Request_Structure) FromJSONToStrike_Meta_Request_Structure (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}