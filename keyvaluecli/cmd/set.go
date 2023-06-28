package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type SetValuePostRequest struct {
	Key   string
	Value string
}

var value string

// set CMD
var setCMD = &cobra.Command{

	Use:   "set",
	Short: "Performs the set operation and set the value for the given key.",
	Long: `Perform a SET operation.
	When no flag is passed, all key-value pairs are fetched.
	Given a key-value pair, the concerned key-value pair is either added or updated if already exists in datastore.`,

	Run: func(cmd *cobra.Command, args []string) {

		payLoad := SetValuePostRequest{
			Key:   key,
			Value: value,
		}

		p, err := json.Marshal(payLoad)

		if err != nil {
			fmt.Println("Invalid key input")
		}

		request, err := http.Post("http://"+serverAddress+"/set", "application/json", bytes.NewBuffer(p))

		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}

		defer request.Body.Close()
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(body))

	},
}

func init() {
	rootCMD.AddCommand(setCMD)

	setCMD.Flags().StringVarP(&key, "key", "k", "", "The key to search for in the key-value store.")
	setCMD.Flags().StringVarP(&value, "value", "v", "", "The value to set for the given key.")
	_ = setCMD.MarkFlagRequired("key")
	_ = setCMD.MarkFlagRequired("value")

}
