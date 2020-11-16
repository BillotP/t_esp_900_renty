// Package models contains all db schemes
package models

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/BillotP/gorenty"
	"gorm.io/gorm"
)

var (
	maxLimit = 200
	// ErrInvalidText is returned when message Text content is empty
	ErrInvalidText = errors.New("invalid message, can't be empty")
	// ErrInvalidLimit is returned when param limit is invalid
	ErrInvalidLimit = errors.New("invalid limit param (should be an int from 1 to 200)")
)

// Message is the sms content to send
type Message struct {
	Text                   string    `json:"text" gorm:"index:message_index"`
	Active                 bool      `json:"active"`
	Templated              bool      `json:"templated"`
	ScrappingRequestItemID uuid.UUID `json:"scrappingRequestId" gorm:"index:message_index"`
}

// MessageItem is a message in db
type MessageItem struct {
	Base
	Message
}

// Valid check message content (should be improved)
func (m Message) Valid() bool {
	return m.Text != ""
}

// Save a new message in database (encode text in base64 first)
func (m Message) Save(db *gorm.DB) (*MessageItem, error) {
	if !m.Valid() { // Check for empty text
		return nil, ErrInvalidText
	}
	enc := base64.StdEncoding.EncodeToString([]byte(m.Text))
	m.Text = enc // Encode text to base64
	msgitem := &MessageItem{
		Message: m,
	}
	if err := db.Create(msgitem).Error; err != nil {
		return nil, err
	}
	return msgitem, nil
}

// Update an existing message in database
func (m Message) Update(db *gorm.DB, uuid uuid.UUID) (*MessageItem, error) {
	if m.Text != "" {
		enc := base64.StdEncoding.EncodeToString([]byte(m.Text))
		m.Text = enc // Encode text to base64
	}
	msgitem := &MessageItem{
		Base: Base{
			UUID: uuid,
		},
		Message: m,
	}
	if err := db.Model(msgitem).Updates(msgitem).Error; err != nil {
		return nil, err
	}
	return msgitem, nil
}

// Delete a message in database
func (m *Message) Delete(db *gorm.DB, uuid uuid.UUID) (*MessageItem, error) {
	msgitem := &MessageItem{
		Base: Base{
			UUID: uuid,
		},
	}
	if err := db.Delete(msgitem).Error; err != nil {
		return nil, err
	}
	return msgitem, nil
}

// DecodeText from base64 storage format
func (m *Message) DecodeText() {
	dec, err := base64.StdEncoding.DecodeString(m.Text)
	if err != nil {
		fmt.Printf("Error(DecodeText): message text [%s] is not valid base64 encoded : %s\n",
			m.Text,
			err.Error())
	}
	m.Text = string(dec)
	if goscrappy.Debug {
		fmt.Println("Info(DecodeText): Got text ", m.Text)
	}
}

// GetMessageByID return a messageitem by its id
func GetMessageByID(db *gorm.DB, id string) (*MessageItem, error) {
	var err error
	var message MessageItem
	var itemuuid uuid.UUID
	if itemuuid, err = uuid.Parse(id); err != nil {
		return nil, err
	}
	if err := db.Where("id = ?", itemuuid).Find(&message).Error; err != nil {
		fmt.Printf("Error(GetMessageByID): %s\n", err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	message.DecodeText()
	return &message, nil
}

// GetMessages return all message in db
func GetMessages(db *gorm.DB, limit *string) (*[]MessageItem, error) {
	var err error
	var messages []MessageItem
	if limit != nil {
		var limitval int
		if limitval, err = strconv.Atoi(*limit); err != nil {
			fmt.Printf("Error(GetMessages): Invalid limit param : %s\n", err.Error())
			return nil, ErrInvalidLimit
		} else if limitval < 1 || limitval > maxLimit {
			return nil, ErrInvalidLimit
		}
		if goscrappy.Debug {
			fmt.Printf("Info(GetMessages): Will looking for %v first messages\n", limitval)
		}
		if err = db.Limit(limitval).Find(&messages).Error; err != nil {
			fmt.Printf("Error(GetMessages): %s\n", err.Error())
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
	} else if err = db.Find(&messages).Error; err != nil {
		fmt.Printf("Error(): %s\n", err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	for el := range messages {
		messages[el].DecodeText()
		if goscrappy.Debug {
			fmt.Printf("Info(GetMessages): Found message [%+v]\n", messages[el])
		}
	}
	return &messages, nil
}
