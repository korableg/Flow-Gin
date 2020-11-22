package Mini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/korableg/mini-gin/Mini/errs"
	"github.com/korableg/mini-gin/Mini/msgs"
	"github.com/korableg/mini-gin/Mini/node"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestMini(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	m := New(nil)

	hub, err := m.NewHub("testHub")
	if err != nil {
		t.Fatal(err)
	}

	hub = nil
	hub = m.GetHub("testHub")
	if hub == nil {
		t.Fatal("Err get hub")
	}

	nodeProducer, _ := m.NewNode("node_producer")
	nodeConsumer, _ := m.NewNode("node_consumer")

	m.AddNodeToHub(hub, nodeConsumer)
	NameNodeConsumer := nodeConsumer.Name()
	nodeConsumer = nil
	nodeConsumer = m.GetNode(NameNodeConsumer)
	if nodeConsumer == nil {
		t.Fatal("node_consumer not found")
	}
	nodesCount := rand.Intn(100) + 1

	for i := 0; i < nodesCount; i++ {
		n, _ := m.NewNode("testNode" + strconv.Itoa(i))
		m.AddNodeToHub(hub, n)
	}

	messageCount := rand.Intn(100) + 1
	mSent := make([]*msgs.Message, 0, messageCount*3)
	mSentChan := make(chan interface{})
	mSentChan2 := make(chan interface{})
	mSentChan3 := make(chan interface{})
	mReceivedChan := make(chan []*msgs.Message)

	funcSend := func(out chan interface{}) {
		for i := 0; i < messageCount; i++ {
			data := make([]byte, rand.Intn(1024*1024*10))
			rand.Read(data)
			mSent = append(mSent, m.SendMessage(nodeProducer, hub, data))
		}
		out <- 1
	}

	funcReceive := func() {
		mReceived := make([]*msgs.Message, messageCount*3, messageCount*3)
		for i := 0; i < messageCount*3; {
			mes := m.GetMessage(nodeConsumer)
			if mes != nil {
				mReceived[i] = mes
				m.RemoveMessage(nodeConsumer)
				i++
			} else {
				time.Sleep(time.Millisecond * 500)
			}
		}
		mReceivedChan <- mReceived
	}

	go funcSend(mSentChan)
	go funcSend(mSentChan2)
	go funcSend(mSentChan3)
	go funcReceive()

	<-mSentChan
	<-mSentChan2
	<-mSentChan3

	mReceived := <-mReceivedChan

	for i := 0; i < messageCount*3; i++ {
		if mSent[i] != mReceived[i] {
			t.Fatal("Sent != Received")
		}
	}

	fmt.Printf("Nodes count %d\n", nodesCount)
	fmt.Printf("msgs received %d\n", len(mReceived))

}

func TestNode_MarshalJSON(t *testing.T) {

	n, _ := node.NewNode("Test")

	nodeBytes, err := json.Marshal(n)
	if err != nil {
		t.Error(err)
	}
	_ = nodeBytes

}

func TestNode_UnmarshalJSON(t *testing.T) {

	buf := bytes.NewBuffer(nil)
	buf.WriteString("{}")
	n := node.Node{}
	err := json.Unmarshal(buf.Bytes(), &n)
	if err != errs.ERR_NODE_NAME_ISEMPTY {
		t.Error("Error must be ERR_NODE_NAME_ISEMPTY")
	}

	buf = bytes.NewBuffer(nil)
	buf.WriteString("{\"name\":\"Test\"}")
	n = node.Node{}
	json.Unmarshal(buf.Bytes(), &n)
	if len(n.Name()) == 0 {
		t.Error("node Name is empty")
	}

	buf = bytes.NewBuffer(nil)
	buf.WriteString("{\"name\":\"Test\"}")
	n = node.Node{}
	json.Unmarshal(buf.Bytes(), &n)
	if len(n.Name()) == 0 {
		t.Error("node Name is empty")
	}

	buf = bytes.NewBuffer(nil)
	buf.WriteString("{\"name\":\"Test\", \"id\":\"TestId\"}")
	n = node.Node{}
	json.Unmarshal(buf.Bytes(), &n)

}
