package transaction

import (
	"fmt"
	"testing"

	"ddd_think/demo.message_server/entity"
)

func TestSave1(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want entity.MsgID
	}{
		// TODO: Add test cases.
		{"1", args{content: "abc"}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Save(tt.args.content)
			fmt.Println(got)
		})
	}
}
