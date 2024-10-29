# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToUser** | Pointer to **string** | The user ID to whom the message is sent | [optional] 
**ToGroup** | Pointer to **string** | The group ID to whom the message is sent | [optional] 
**Text** | Pointer to **string** | The message text | [optional] 
**DataId** | Pointer to **string** | The data ID | [optional] 

## Methods

### NewMessage

`func NewMessage() *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToUser

`func (o *Message) GetToUser() string`

GetToUser returns the ToUser field if non-nil, zero value otherwise.

### GetToUserOk

`func (o *Message) GetToUserOk() (*string, bool)`

GetToUserOk returns a tuple with the ToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUser

`func (o *Message) SetToUser(v string)`

SetToUser sets ToUser field to given value.

### HasToUser

`func (o *Message) HasToUser() bool`

HasToUser returns a boolean if a field has been set.

### GetToGroup

`func (o *Message) GetToGroup() string`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *Message) GetToGroupOk() (*string, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *Message) SetToGroup(v string)`

SetToGroup sets ToGroup field to given value.

### HasToGroup

`func (o *Message) HasToGroup() bool`

HasToGroup returns a boolean if a field has been set.

### GetText

`func (o *Message) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Message) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Message) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Message) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDataId

`func (o *Message) GetDataId() string`

GetDataId returns the DataId field if non-nil, zero value otherwise.

### GetDataIdOk

`func (o *Message) GetDataIdOk() (*string, bool)`

GetDataIdOk returns a tuple with the DataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataId

`func (o *Message) SetDataId(v string)`

SetDataId sets DataId field to given value.

### HasDataId

`func (o *Message) HasDataId() bool`

HasDataId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


