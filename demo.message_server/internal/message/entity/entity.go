package entity

type Entity struct {
	ID      MsgID
	Content string `db:"content"`
}

type MsgID string
