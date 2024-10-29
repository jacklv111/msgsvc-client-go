# CreateMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToUser** | Pointer to **string** | The user ID to whom the message is sent | [optional] 
**ToGroup** | Pointer to **string** | The group ID to whom the message is sent | [optional] 
**Text** | Pointer to **string** | The message text | [optional] 
**DataId** | Pointer to **string** | The data ID | [optional] 
**FileName** | Pointer to **string** | The file name | [optional] 

## Methods

### NewCreateMessage

`func NewCreateMessage() *CreateMessage`

NewCreateMessage instantiates a new CreateMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageWithDefaults

`func NewCreateMessageWithDefaults() *CreateMessage`

NewCreateMessageWithDefaults instantiates a new CreateMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToUser

`func (o *CreateMessage) GetToUser() string`

GetToUser returns the ToUser field if non-nil, zero value otherwise.

### GetToUserOk

`func (o *CreateMessage) GetToUserOk() (*string, bool)`

GetToUserOk returns a tuple with the ToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUser

`func (o *CreateMessage) SetToUser(v string)`

SetToUser sets ToUser field to given value.

### HasToUser

`func (o *CreateMessage) HasToUser() bool`

HasToUser returns a boolean if a field has been set.

### GetToGroup

`func (o *CreateMessage) GetToGroup() string`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *CreateMessage) GetToGroupOk() (*string, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *CreateMessage) SetToGroup(v string)`

SetToGroup sets ToGroup field to given value.

### HasToGroup

`func (o *CreateMessage) HasToGroup() bool`

HasToGroup returns a boolean if a field has been set.

### GetText

`func (o *CreateMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDataId

`func (o *CreateMessage) GetDataId() string`

GetDataId returns the DataId field if non-nil, zero value otherwise.

### GetDataIdOk

`func (o *CreateMessage) GetDataIdOk() (*string, bool)`

GetDataIdOk returns a tuple with the DataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataId

`func (o *CreateMessage) SetDataId(v string)`

SetDataId sets DataId field to given value.

### HasDataId

`func (o *CreateMessage) HasDataId() bool`

HasDataId returns a boolean if a field has been set.

### GetFileName

`func (o *CreateMessage) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CreateMessage) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CreateMessage) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CreateMessage) HasFileName() bool`

HasFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


