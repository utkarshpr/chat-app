package models

type Message struct {
	ID          string `json:"message_id" bson:"message_id"`
	ChatID      string `json:"chat_id" bson:"chat_id"`
	SenderID    string `json:"sender_id" bson:"sender_id"`
	RecipientID string `json:"recipient_id" bson:"recipient_id"`
	Content     string `json:"content" bson:"content"`
	MediaURL    string `json:"media_url,omitempty" bson:"media_url,omitempty"`
	Timestamp   string `json:"timestamp" bson:"timestamp"`
	Status      string `json:"status" bson:"status"`
}

type MessageStatusUpdate struct {
	MessageID string `json:"message_id"`
	Status    string `json:"status"`
}

type GetMessage struct {
	Content   string `json:"content" bson:"content"`
	MediaURL  string `json:"media_url,omitempty" bson:"media_url,omitempty"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}
