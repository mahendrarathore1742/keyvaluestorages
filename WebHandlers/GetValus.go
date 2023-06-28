package webhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/mahendrarathore1742/keyvaluestorages/package/keystore"
)

type IncomingGetRequest struct {
	Key string `json:"key"`
}

type OutgoingPostRequest struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

func GetValues(kv *keystore.KeyStore) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var request IncomingGetRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		value, err := kv.GetValue(request.Key)

		if err != nil {
			http.Error(w, "The given key and key corresponding data is not avaliabe in the DataStore.", http.StatusInternalServerError)
			return
		}

		var response OutgoingPostRequest
		response.Key = request.Key
		response.Value = value

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}

}

func GetAllValues(kv *keystore.KeyStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(kv.Pairs)
	}
}
