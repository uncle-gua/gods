// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashmap

import (
	"encoding/json"
	"github.com/uncle-gua/gods/containers"
	"github.com/uncle-gua/gods/utils"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Map)(nil)
var _ containers.JSONDeserializer = (*Map)(nil)

// ToJSON outputs the JSON representation of the map.
func (m *Map) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	for key, value := range m.m {
		elements[utils.ToString(key)] = value
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *Map) FromJSON(data []byte) error {
	elements := make(map[string]interface{})
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.m[key] = value
		}
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *Map) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}
