/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Position type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Position{}

// Position struct for Position
type Position struct {
	Character int32 `json:"character"`
	Line      int32 `json:"line"`
}

type _Position Position

// NewPosition instantiates a new Position object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPosition(character int32, line int32) *Position {
	this := Position{}
	this.Character = character
	this.Line = line
	return &this
}

// NewPositionWithDefaults instantiates a new Position object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionWithDefaults() *Position {
	this := Position{}
	return &this
}

// GetCharacter returns the Character field value
func (o *Position) GetCharacter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Character
}

// GetCharacterOk returns a tuple with the Character field value
// and a boolean to check if the value has been set.
func (o *Position) GetCharacterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Character, true
}

// SetCharacter sets field value
func (o *Position) SetCharacter(v int32) {
	o.Character = v
}

// GetLine returns the Line field value
func (o *Position) GetLine() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Line
}

// GetLineOk returns a tuple with the Line field value
// and a boolean to check if the value has been set.
func (o *Position) GetLineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line, true
}

// SetLine sets field value
func (o *Position) SetLine(v int32) {
	o.Line = v
}

func (o Position) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Position) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["character"] = o.Character
	toSerialize["line"] = o.Line
	return toSerialize, nil
}

func (o *Position) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"character",
		"line",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPosition := _Position{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPosition)

	if err != nil {
		return err
	}

	*o = Position(varPosition)

	return err
}

type NullablePosition struct {
	value *Position
	isSet bool
}

func (v NullablePosition) Get() *Position {
	return v.value
}

func (v *NullablePosition) Set(val *Position) {
	v.value = val
	v.isSet = true
}

func (v NullablePosition) IsSet() bool {
	return v.isSet
}

func (v *NullablePosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePosition(val *Position) *NullablePosition {
	return &NullablePosition{value: val, isSet: true}
}

func (v NullablePosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}