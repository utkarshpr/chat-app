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

type EditMessage struct {
	ID         string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FromUserID string `json:"from_user_id" gorm:"type:uuid;not null"`
	ToUserID   string `json:"to_user_id" gorm:"type:uuid;not null"`
	NewText    string `json:"new_text" gorm:"type:text"`
}
