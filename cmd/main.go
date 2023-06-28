package main

import (
	"log"
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	webhandlers "github.com/mahendrarathore1742/keyvaluestorages/WebHandlers"
	"github.com/mahendrarathore1742/keyvaluestorages/package/kafka"
	"github.com/mahendrarathore1742/keyvaluestorages/package/keystore"
)

func main() {

	s := keystore.NewValues()

	s.SetValue("Name", "Mahendra")
	s.SetValue("Age", "23")

	kafkaWriter := kafka.GetKafKaWriter()
	defer kafkaWriter.Close()

	kafkaReader := kafka.GetKafKaReader()
	defer kafkaReader.Close()

	r := mux.NewRouter()

	r.Path("/get").Handler(webhandlers.GetValues(&s))
	r.Path("/getall").Handler(webhandlers.GetAllValues(&s))
	r.Path("/set").Handler(webhandlers.SetValues(&s, kafkaWriter))
	r.Path("/subscribe").Handler(webhandlers.Subscribes(&s, kafkaReader))
	log.Fatal(http.ListenAndServe(":3000", gorillaHandlers.CORS()(r)))
}
