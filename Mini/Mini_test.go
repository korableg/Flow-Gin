package Mini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/korableg/mini-gin/Mini/Errors"
	"github.com/korableg/mini-gin/Mini/Messages"
	"github.com/korableg/mini-gin/Mini/Node"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestMini(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	m := NewMini()

	hub, err := m.NewHub("testHub")
	if err != nil {
		t.Fatal(err)
	}

	hub = nil
	hub = m.GetHub("testHub")
	if hub == nil {
		t.Fatal("Err get hub")
	}

	node_producer, _ := m.NewNode("node_producer")
	node_consumer, _ := m.NewNode("node_consumer")

	m.AddNodeToHub(hub, node_consumer)
	NameNodeConsumer := node_consumer.Name()
	node_consumer = nil
	node_consumer = m.GetNode(NameNodeConsumer)
	if node_consumer == nil {
		t.Fatal("No finded node_consumer")
	}
	nodes_Count := rand.Intn(100) + 1

	for i := 0; i < nodes_Count; i++ {
		n, _ := m.NewNode("testNode" + strconv.Itoa(i))
		m.AddNodeToHub(hub, n)
	}

	messageCount := rand.Intn(100) + 1
	mSended := make([]*Messages.Message, 0, messageCount*3)
	mSendedChan := make(chan interface{})
	mSendedChan2 := make(chan interface{})
	mSendedChan3 := make(chan interface{})
	mReceivedChan := make(chan []*Messages.Message)

	funcSend := func(out chan interface{}) {
		for i := 0; i < messageCount; i++ {
			data := make([]byte, rand.Intn(1024*1024*10))
			rand.Read(data)
			mSended = append(mSended, m.SendMessage(node_producer, hub, data))
		}
		out <- 1
	}

	funcReceive := func() {
		mReceived := make([]*Messages.Message, messageCount*3, messageCount*3)
		for i := 0; i < messageCount*3; {
			mes := m.GetMessage(node_consumer)
			if mes != nil {
				mReceived[i] = mes
				m.RemoveMessage(node_consumer)
				i++
			} else {
				time.Sleep(time.Millisecond * 500)
			}
		}
		mReceivedChan <- mReceived
	}

	go funcSend(mSendedChan)
	go funcSend(mSendedChan2)
	go funcSend(mSendedChan3)
	go funcReceive()

	<-mSendedChan
	<-mSendedChan2
	<-mSendedChan3

	mReceived := <-mReceivedChan

	for i := 0; i < messageCount*3; i++ {
		if mSended[i] != mReceived[i] {
			t.Fatal("Sended != Received")
		}
	}

	fmt.Printf("Nodes count %d\n", nodes_Count)
	fmt.Printf("Messages received %d\n", len(mReceived))

}

func TestNode_MarshalJSON(t *testing.T) {

	n, _ := Node.NewNode("Test")

	nodeBytes, err := json.Marshal(n)
	if err != nil {
		t.Error(err)
	}
	_ = nodeBytes

}

func TestNode_UnmarshalJSON(t *testing.T) {

	buf := bytes.NewBuffer(nil)
	buf.WriteString("{}")
	node := Node.Node{}
	err := json.Unmarshal(buf.Bytes(), &node)
	if err != Errors.ERR_NODE_NAME_ISEMPTY {
		t.Error("Error must be ERR_NODE_NAME_ISEMPTY")
	}

	buf = bytes.NewBuffer(nil)
	buf.WriteString("{\"name\":\"Test\"}")
	node = Node.Node{}
	json.Unmarshal(buf.Bytes(), &node)
	if len(node.Name()) == 0 {
		t.Error("Node Name is empty")
	}

	buf = bytes.NewBuffer(nil)
	buf.WriteString("{\"name\":\"Test\"}")
	node = Node.Node{}
	json.Unmarshal(buf.Bytes(), &node)
	if len(node.Name()) == 0 {
		t.Error("Node Name is empty")
	}

	buf = bytes.NewBuffer(nil)
	buf.WriteString("{\"name\":\"Test\", \"id\":\"TestId\"}")
	node = Node.Node{}
	json.Unmarshal(buf.Bytes(), &node)

}
