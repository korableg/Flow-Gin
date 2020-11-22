package msgs

import (
	"container/list"
	"sync"
	"time"
)

type Message struct {
	id   int64
	from string
	data []byte
}

type MessageQueue struct {
	l     *list.List
	mutex sync.Mutex
}

func NewMessage(from string, data []byte) *Message {
	mes := &Message{
		id:   time.Now().UnixNano(),
		from: from,
		data: data,
	}
	return mes
}

func (m *Message) ID() int64 {
	return m.id
}

func (m *Message) From() string {
	return m.from
}

func (m *Message) Data() []byte {
	return m.data
}

func NewMessageQueue() *MessageQueue {
	m := &MessageQueue{
		l: list.New(),
	}
	return m
}

func (m *MessageQueue) Front() (mes *Message) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	e := m.l.Front()
	if e != nil {
		mes = e.Value.(*Message)
	}

	return mes
}

func (m *MessageQueue) Push(mes *Message) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.l.PushBack(mes)
}

func (m *MessageQueue) RemoveFront() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	e := m.l.Front()
	if e != nil {
		m.l.Remove(e)
	}
}

func (m *MessageQueue) Len() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.l.Len()
}
