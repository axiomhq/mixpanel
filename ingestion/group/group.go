// Package groups describes data models to interact with groups
// provided by Mixpanel Ingestion API.
package group

// Set describes model of request to set values of groups properties.
// If the groups does not exist, it creates it with these properties.
// If it does exist, it sets the properties to these values, overwriting existing values.
type Set struct {
	Token    string                 `json:"$token"`
	GroupKey string                 `json:"$group_key"`
	GroupID  string                 `json:"$group_id"`
	Set      map[string]interface{} `json:"$set"`
}

// SetOnce describes model of request to set values of groups properties only once, without overwriting.
// If the groups does not exist, it creates it with these properties.
// If it does exist, it sets the properties to these values, overwriting existing values.
type SetOnce struct {
	Token    string                 `json:"$token"`
	GroupKey string                 `json:"$group_key"`
	GroupID  string                 `json:"$group_id"`
	SetOnce  map[string]interface{} `json:"$set_once"`
}

// ListAppend describes model of request to append item from groups list property.
type ListAppend struct {
	Token    string                 `json:"$token"`
	GroupKey string                 `json:"$group_key"`
	GroupID  string                 `json:"$group_id"`
	Append   map[string]interface{} `json:"$append"`
}

// ListRemove describes model of request to remove item from groups list property.
type ListRemove struct {
	Token    string                 `json:"$token"`
	GroupKey string                 `json:"$group_key"`
	GroupID  string                 `json:"$group_id"`
	Remove   map[string]interface{} `json:"$remove"`
}

// Unset describes model of request to unset groups property.
type Unset struct {
	Token    string   `json:"$token"`
	GroupKey string   `json:"$group_key"`
	GroupID  string   `json:"$group_id"`
	Unset    []string `json:"$unset"`
}

// Mutator is internal interface to mark models as groups mutation actions.
type Mutator interface {
	isMutator()
}

func (x Set) isMutator()        {}
func (x SetOnce) isMutator()    {}
func (x NumberAdd) isMutator()  {}
func (x ListAppend) isMutator() {}
func (x ListRemove) isMutator() {}
func (x Unset) isMutator()      {}
