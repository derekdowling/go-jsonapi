package jsh

import (
	"encoding/json"
	"fmt"
)

// Links is a top-level document field
type Links struct {
	Self    *Link `json:"self,omitempty"`
	Related *Link `json:"related,omitempty"`
}

// Link is a JSON format type
type Link struct {
	HREF string                 `json:"href,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

func (link *Link) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	switch vv := v.(type) {
	case string:
		link.HREF = vv
	case Link:
		*link = vv
	default:
		return fmt.Errorf("unknown Link type (got %v)", vv)
	}
	return nil
}
