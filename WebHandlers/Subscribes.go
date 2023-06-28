package webhandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mahendrarathore1742/keyvaluestorages/package/keystore"
	kafkaGo "github.com/segmentio/kafka-go"
)

type kafkaRecords struct {
	Topic     string
	Partition int
	Offset    int
	Key       string
	Value     string
}

var upgrader = websocket.Upgrader{}

func Subscribes(kv *keystore.KeyStore, kafkaReader *kafkaGo.Reader) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}

		defer c.Close()

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Error during message reading:", err)
				return
			}

			log.Printf("Received: %s", message)

			m, err := kafkaReader.ReadMessage(context.Background())

			if err != nil {
				log.Fatalln(err)
			}

			var msg kafkaRecords
			msg.Topic = m.Topic
			msg.Partition = m.Partition
			msg.Offset = int(m.Offset)
			msg.Key = string(m.Key)
			msg.Value = string(m.Value)

			err = c.WriteJSON(msg)
			if err != nil {
				log.Println("Error during message writing:", err)
				break
			}
		}

	}

}
