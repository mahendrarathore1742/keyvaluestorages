package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var done = make(chan interface{})
var interrupt = make(chan os.Signal, 1)

func receiveHandler(connection *websocket.Conn) {

	defer close(done)

	for {

		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Fatalln("Error in receive:", err)
			return
		}

		prettyJSON, err := PrettyString(string(msg))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(prettyJSON)
	}

}

var watchCMD = &cobra.Command{

	Use:   "watch",
	Short: "Watch the changes happening to all keys in the store",
	Long: `Watch the server for changes done by other clients.
	The client's address and the edited key value pair is shown in JSON as received.
	`,

	Run: func(cmd *cobra.Command, args []string) {

		signal.Notify(interrupt, os.Interrupt)
		socketUrl := "ws://" + serverAddress + "/subscribe"

		conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)

		if err != nil {
			log.Fatal("Error connecting to Websocket Server:", err)
		}

		defer conn.Close()
		go receiveHandler(conn)

		for {
			select {

			case <-time.After(time.Duration(1) * time.Microsecond * 1000):
				err := conn.WriteMessage(websocket.TextMessage, []byte("Keepalive ping"))
				if err != nil {
					log.Fatalln("Error during writing to websocket:", err)
					return
				}

			case <-interrupt:
				log.Println("Received SIGINT interrupt signal. Closing all pending connections")

				// just close websocket connection.
				err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

				if err != nil {
					log.Fatalln("Error during closing websocket:", err)
					return
				}

				select {
				case <-done:
					log.Println("Receiver Channel Closed! Exiting....")

				case <-time.After(time.Duration(1) * time.Second):
					log.Println("Timeout in closing receiving channel. Exiting....")

				}

				return
			}
		}
	},
}

func init() {
	rootCMD.AddCommand(watchCMD)

}
