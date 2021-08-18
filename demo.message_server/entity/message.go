package entity

type MsgID string

type Message struct {
	ID      MsgID
	Content string
}

// New ...
func New(content string) *Message {
	return &Message{
		// ID:      MsgID(uuid.New()),
		Content: content,
	}
}

// Verify 远程验证
func (m Message) Verify() error {
	return nil
}
