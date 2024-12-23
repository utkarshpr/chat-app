package models

import (
	"errors"
	"time"
)

// ContactRequest represents a request to add a new contact.
type ContactRequest struct {
	FromUserID string `json:"from_user_id" bson:"from_user_id"`
	ToUserID   string `json:"to_user_id" bson:"to_user_id"`
	Status     string `json:"status" bson:"status"` // pending, accepted, blocked
}

// ContactActionRequest represents an action (block or remove) on a contact.
type ContactActionRequest struct {
	UserID    string `json:"user_id" bson:"user_id"`
	ContactID string `json:"contact_id" bson:"contact_id"`
	Action    string `json:"action" bson:"action"` // remove, block
}

// Contact represents a user's contact information and status.
type Contact struct {
	FromUserID string    `json:"from_user_id" bson:"from_user_id"`
	ToUserID   string    `json:"to_user_id" bson:"to_user_id"`
	Status     string    `json:"status" bson:"status"` // pending, accepted, blocked
	LastOnline time.Time `json:"last_online" bson:"last_online"`
}

const (
	StatusPending  = "pending"
	StatusAccepted = "accepted"
	StatusRejected = "rejected"
)

// IsValidStatus validates the status value.
func (car *ContactRequest) IsValidStatus() bool {
	switch car.Status {
	case StatusPending, StatusAccepted, StatusRejected:
		return true
	}
	return false
}

const (
	ActionRemove = "remove"
	ActionBlock  = "block"
)

// IsValid checks if the action is valid.
func (car *ContactActionRequest) IsValid() error {
	switch car.Action {
	case ActionRemove, ActionBlock:
		return nil
	default:
		return errors.New("invalid action value")
	}
}
