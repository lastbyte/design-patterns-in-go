package singleton

import "testing"

func TestGetMessageQueue(t *testing.T) {

	mq := GetMessageQueue()

	if mq == nil {
		t.Fatal("Error occured in test")
	}
}

func TestSendMessage(t *testing.T) {
	mq := GetMessageQueue()

	mq.SendMessage("message")

	if mq.msgCount != 1 {
		t.Fatal("error in sending message")
	}
}

func TestSendEmptyMessage(t *testing.T) {
	mq := GetMessageQueue()

	mq.SendMessage("")

	if mq.msgCount == 0 {
		t.Fatal("error in sending message")
	}

}
