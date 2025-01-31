/*
msgsvc

msgsvc api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgsvcclientgo

import (
	"encoding/json"
)

// checks if the CreateMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMessage{}

// CreateMessage struct for CreateMessage
type CreateMessage struct {
	// The user ID to whom the message is sent
	ToUser *string `json:"toUser,omitempty"`
	// The group ID to whom the message is sent
	ToGroup *string `json:"toGroup,omitempty"`
	// The message text
	Text *string `json:"text,omitempty"`
	// The data ID
	DataId *string `json:"dataId,omitempty"`
	// The file name
	FileName *string `json:"fileName,omitempty"`
}

// NewCreateMessage instantiates a new CreateMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessage() *CreateMessage {
	this := CreateMessage{}
	return &this
}

// NewCreateMessageWithDefaults instantiates a new CreateMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessageWithDefaults() *CreateMessage {
	this := CreateMessage{}
	return &this
}

// GetToUser returns the ToUser field value if set, zero value otherwise.
func (o *CreateMessage) GetToUser() string {
	if o == nil || IsNil(o.ToUser) {
		var ret string
		return ret
	}
	return *o.ToUser
}

// GetToUserOk returns a tuple with the ToUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessage) GetToUserOk() (*string, bool) {
	if o == nil || IsNil(o.ToUser) {
		return nil, false
	}
	return o.ToUser, true
}

// HasToUser returns a boolean if a field has been set.
func (o *CreateMessage) HasToUser() bool {
	if o != nil && !IsNil(o.ToUser) {
		return true
	}

	return false
}

// SetToUser gets a reference to the given string and assigns it to the ToUser field.
func (o *CreateMessage) SetToUser(v string) {
	o.ToUser = &v
}

// GetToGroup returns the ToGroup field value if set, zero value otherwise.
func (o *CreateMessage) GetToGroup() string {
	if o == nil || IsNil(o.ToGroup) {
		var ret string
		return ret
	}
	return *o.ToGroup
}

// GetToGroupOk returns a tuple with the ToGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessage) GetToGroupOk() (*string, bool) {
	if o == nil || IsNil(o.ToGroup) {
		return nil, false
	}
	return o.ToGroup, true
}

// HasToGroup returns a boolean if a field has been set.
func (o *CreateMessage) HasToGroup() bool {
	if o != nil && !IsNil(o.ToGroup) {
		return true
	}

	return false
}

// SetToGroup gets a reference to the given string and assigns it to the ToGroup field.
func (o *CreateMessage) SetToGroup(v string) {
	o.ToGroup = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CreateMessage) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessage) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CreateMessage) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *CreateMessage) SetText(v string) {
	o.Text = &v
}

// GetDataId returns the DataId field value if set, zero value otherwise.
func (o *CreateMessage) GetDataId() string {
	if o == nil || IsNil(o.DataId) {
		var ret string
		return ret
	}
	return *o.DataId
}

// GetDataIdOk returns a tuple with the DataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessage) GetDataIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataId) {
		return nil, false
	}
	return o.DataId, true
}

// HasDataId returns a boolean if a field has been set.
func (o *CreateMessage) HasDataId() bool {
	if o != nil && !IsNil(o.DataId) {
		return true
	}

	return false
}

// SetDataId gets a reference to the given string and assigns it to the DataId field.
func (o *CreateMessage) SetDataId(v string) {
	o.DataId = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *CreateMessage) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessage) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *CreateMessage) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *CreateMessage) SetFileName(v string) {
	o.FileName = &v
}

func (o CreateMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ToUser) {
		toSerialize["toUser"] = o.ToUser
	}
	if !IsNil(o.ToGroup) {
		toSerialize["toGroup"] = o.ToGroup
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.DataId) {
		toSerialize["dataId"] = o.DataId
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	return toSerialize, nil
}

type NullableCreateMessage struct {
	value *CreateMessage
	isSet bool
}

func (v NullableCreateMessage) Get() *CreateMessage {
	return v.value
}

func (v *NullableCreateMessage) Set(val *CreateMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessage(val *CreateMessage) *NullableCreateMessage {
	return &NullableCreateMessage{value: val, isSet: true}
}

func (v NullableCreateMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


