package email

import (
	"log"
	"testing"
)

func TestEmail_Send(t *testing.T) {
	newEmail := NewQQExmail("邮箱账号", "授权码或者邮箱密码")
	text := newEmail.Info("标题", "作者", []string{"收件人列表"}).
		SendText("文本内容")
	log.Println(text)
}
