package msgs

import (
	"bytes"
	"testing"
)

func TestMsgs(t *testing.T) {

	wantFrom := "testFrom"
	wantData := []byte("testData")

	mes := NewMessage(wantFrom, wantData)
	_ = mes.ID()
	if mes.From() != wantFrom {
		t.Errorf("want from %s got from %s", wantFrom, mes.From())
	}

	mq := NewMessageQueue()
	mq.Push(mes)
	mq.Push(NewMessage(wantFrom, nil))
	if mq.Len() != 2 {
		t.Error("length message queue must be 2")
	}
	mes = mq.Front()
	gotData := mes.Data()
	if !bytes.Equal(wantData, gotData) {
		t.Error("want data != got data")
	}
	mq.RemoveFront()
	mes = mq.Front()
	gotData = mes.Data()
	if nil != gotData {
		t.Error("want data != got data")
	}
	mq.RemoveFront()
	if mq.Len() != 0 {
		t.Error("length message queue must be 0")
	}

}
