package bot

type getUpdatesResponse struct {
	Result []update
}

type update struct {
	UpdateId      int64   `json:"update_id"`
	Message       message `json:",omitempty"`
	EditedMessage message `json:"edited_message,omitempty"`
}

type message struct {
	MessageId int64    `json:"message_id"`
	From      user     `json:",omitempty"`
	Document  document `json:",omitempty"`
}

type user struct {
}

type document struct {
}
