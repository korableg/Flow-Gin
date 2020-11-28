package flow

import (
	"encoding/json"
	"fmt"
	"github.com/korableg/mini-gin/flow/errs"
	hub2 "github.com/korableg/mini-gin/flow/hub"
	"github.com/korableg/mini-gin/flow/msgs"
	"github.com/korableg/mini-gin/flow/node"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestFlow(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	mockDB := new(mockDB)

	m := New(mockDB)

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

	err = m.AddNodeToHub(hub.Name(), nodeConsumer.Name())
	if err != nil {
		t.Error(err)
	}
	NameNodeConsumer := nodeConsumer.Name()
	nodeConsumer = nil
	nodeConsumer = m.GetNode(NameNodeConsumer)
	if nodeConsumer == nil {
		t.Fatal("node_consumer not found")
	}
	if nodeConsumer.Len() != 0 {
		t.Error("message queue len error")
	}
	nodesCount := rand.Intn(100) + 1

	for i := 0; i < nodesCount; i++ {
		n, _ := m.NewNode("testNode" + strconv.Itoa(i))
		m.AddNodeToHub(hub.Name(), n.Name())
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
			mes, _ := m.SendMessage(nodeProducer.Name(), hub.Name(), data)
			mSent = append(mSent, mes)
		}
		out <- 1
	}

	funcReceive := func() {
		mReceived := make([]*msgs.Message, messageCount*3, messageCount*3)
		for i := 0; i < messageCount*3; {
			mes, _ := m.GetMessage(nodeConsumer.Name())
			if mes != nil {
				mReceived[i] = mes
				m.RemoveMessage(nodeConsumer.Name())
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

	err = m.DeleteNode(nodeConsumer.Name())
	if err != nil {
		t.Error(err)
	}
	err = m.DeleteHub(hub.Name())

	err = m.Close()
	if err != nil {
		t.Fatal(err)
	}

}

func TestHub(t *testing.T) {

	nameHub := "TestHub1"
	nameNode := "TestNode1"

	_, err := hub2.New("   ", nil)
	if err != errs.ERR_HUB_NAME_NOT_MATCHED_PATTERN {
		t.Error(err)
	}
	_, err = hub2.New("", nil)
	if err != errs.ERR_HUB_NAME_ISEMPTY {
		t.Error(err)
	}
	_, err = hub2.New("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890", nil)
	if err != errs.ERR_HUB_NAME_OVER100 {
		t.Error(err)
	}
	hub, err := hub2.New(nameHub, new(mockDB))
	if err != nil {
		t.Fatal(err)
	}
	if nameHub != hub.Name() {
		t.Error(nameHub + " != " + hub.Name())
	}

	n, err := node.New(nameNode)
	if err != nil {
		t.Fatal(err)
	}

	err = hub.AddNode(n)
	if err != nil {
		t.Error(err)
	}

	hub.PushMessage(msgs.NewMessage(nameNode, nil))

	_, err = json.Marshal(hub)

	err = hub.DeleteNode(n)
	if err != nil {
		t.Error(err)
	}

}

func TestNode(t *testing.T) {

	nameNode := "TestNode1"

	_, err := node.New("   ")
	if err != errs.ERR_NODE_NAME_NOT_MATCHED_PATTERN {
		t.Error(err)
	}
	_, err = node.New("")
	if err != errs.ERR_NODE_NAME_ISEMPTY {
		t.Error(err)
	}
	_, err = node.New("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	if err != errs.ERR_NODE_NAME_OVER100 {
		t.Error(err)
	}
	n, err := node.New(nameNode)
	if err != nil {
		t.Fatal(err)
	}
	if nameNode != n.Name() {
		t.Error(nameNode + " != " + n.Name())
	}

	n.PushMessage(msgs.NewMessage("TestNode2", nil))

	if n.Len() != 1 {
		t.Error("count messages must be 1")
	}

	m := n.FrontMessage()
	if m == nil {
		t.Error("front message must be not nil")
	}
	n.RemoveFrontMessage()

	if n.Len() != 0 {
		t.Error("count messages must be 0")
	}

	_, err = json.Marshal(n)

}
