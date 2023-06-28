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

type GetValuePostRequest struct {
	Key string
}

var isAll bool

func PrettyString(str string) (string, error) {
	var PrettyJSON bytes.Buffer

	if err := json.Indent(&PrettyJSON, []byte(str), "", " "); err != nil {
		return "", err
	}

	return PrettyJSON.String(), nil
}

var getCommands = &cobra.Command{

	Use:   "get",
	Short: "Performs the get operation and returns the value for the given key.",
	Long: `Perform a GET operation.
	When no flag is passed, all key-value pairs are fetched.
	Given a key, only the concerned key-value pair is fetched.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if key == "" || isAll {
			request, err := http.Get("http://" + serverAddress + "/getall")

			if err != nil {
				log.Fatalf("Error contacting server: %v", err)
			}

			defer request.Body.Close()
			body, err := ioutil.ReadAll(request.Body)

			if err != nil {
				log.Fatalf("Error reading server response: %v", err)
			}

			prettyJSON, err := PrettyString(string(body))

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(prettyJSON)

		} else {
			payload := GetValuePostRequest{
				Key: key,
			}

			p, err := json.Marshal(payload)
			if err != nil {
				log.Fatalf("Invalid key input. Error: %v", err)
			}

			request, err := http.Post("http://"+serverAddress+"/get", "application/json", bytes.NewBuffer(p))

			if err != nil {
				log.Fatalf("Error contacting server: %v", err)
			}

			defer request.Body.Close()
			body, err := ioutil.ReadAll(request.Body)

			if err != nil {
				log.Fatalf("Error reading server response: %v", err)
			}
			prettyJSON, err := PrettyString(string(body))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(prettyJSON)
		}

	},
}

func init() {
	rootCMD.AddCommand(getCommands)

	getCommands.Flags().StringVarP(&key, "key", "k", "", "The key to search for in the key-value store.")
	getCommands.Flags().BoolVarP(&isAll, "all", "a", false, "Set to true to fetch all keys")
}
