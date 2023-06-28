package webhandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mahendrarathore1742/keyvaluestorages/package/kafka"
	"github.com/mahendrarathore1742/keyvaluestorages/package/keystore"
	kafkaGo "github.com/segmentio/kafka-go"
)

type IncommingPostReqest struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

func SetValues(kv *keystore.KeyStore, kafkaWriter *kafkaGo.Writer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var Request IncommingPostReqest
		err := json.NewDecoder(r.Body).Decode(&Request)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = kafka.AppendCommandLog(r.Context(), kafkaWriter, []byte(fmt.Sprintf("Client address=%s", r.RemoteAddr)), []byte(fmt.Sprintf("%s: %s", Request.Key, Request.Value)))

		if err != nil {
			w.Write([]byte(err.Error()))
			log.Fatalln(err)
		}

		kv.SetValue(Request.Key, Request.Value)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode("SUCCESS")
	}
}
