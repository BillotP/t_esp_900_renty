package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BillotP/gorenty"
)

// TGBASEURL is the base endpoint to get and set bot updates
const TGBASEURL = "https://api.telegram.org/bot"

// Message is telegram mesage API format
type Message struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

// Chat is telegram bot chat client
type Chat struct {
	client     *http.Client
	chatID     string
	tgbotToken string
	botURL     string
}

// NewChat return a authenticated client to send telegram message with
func NewChat(chatID, botToken string) *Chat {
	return &Chat{
		client:     http.DefaultClient,
		chatID:     chatID,
		tgbotToken: botToken,
		botURL:     TGBASEURL + botToken,
	}
}

// SendMessage send a message your client chat id
func (c *Chat) SendMessage(format string, args ...interface{}) (err error) {
	var reqbody []byte
	var res *http.Response
	var nreq *http.Request
	nmsg := Message{
		ChatID: c.chatID,
		Text:   fmt.Sprintf(format, args...),
	}
	if reqbody, err = json.Marshal(nmsg); err != nil {
		return err
	}
	if nreq, err = http.NewRequest(http.MethodPost, c.botURL+"/sendMessage", bytes.NewReader(reqbody)); err != nil {
		return err
	}
	nreq.Header.Add("Content-Type", "application/json")
	if res, err = c.client.Do(nreq); err != nil {
		fmt.Printf("Error(SendMessage): %s\n", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		if goscrappy.Debug {
			fmt.Printf("Error(SendMessage): Invalid req : %s\n", string(bodyBytes))
		}
	}
	defer res.Body.Close()
	return err
}
