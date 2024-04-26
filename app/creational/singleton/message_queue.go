package singleton

import (
	"sync"
)

type messageQueue struct {
	msgCount int
}

type MessageQueue interface {
	SendMessage(msg string) int
}

var queue *messageQueue
var qLock = &sync.Mutex{}

func GetMessageQueue() *messageQueue {
	if queue == nil {
		queue = new(messageQueue)
	}
	return queue
}

func (mq *messageQueue) SendMessage(msg string) int {
	if msg != "" {
		mq.msgCount = mq.msgCount + 1
	}
	return mq.msgCount
}
