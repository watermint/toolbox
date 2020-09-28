package mo_conversation

import "encoding/json"

type Conversation struct {
	Raw        json.RawMessage `json:"-"`
	Id         string          `json:"id" path:"id"`
	Name       string          `json:"name" path:"name"`
	IsArchived bool            `json:"is_archived" path:"is_archived"`
	IsPrivate  bool            `json:"is_private" path:"is_private"`
	Topic      string          `json:"topic" path:"topic.value"`
	Purpose    string          `json:"purpose" path:"purpose.value"`
	NumMembers int             `json:"num_members" path:"num_members"`
}
