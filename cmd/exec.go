// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "check health of endpoints",
	Long:  `check health of endpoints`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("exec called")
		exec()
	},
}

type Result struct {
	Responses []Response `json:"responses"`
}

// Response is return value to stdout
type Response struct {
	URL        string `json:"url"`
	StatusCode int    `json:"statusCode"`
}

func exec() {

	list, err := readLines(viper.GetString("recordFileName"))
	if err != nil {
		log.Fatal(err)
	}

	datas := make([]Response, 0)
	for _, url := range list {
		resp, err := http.Get(url)

		statusCode := 0

		if err != nil {
			log.Println(err)
		} else {
			statusCode = resp.StatusCode
			defer resp.Body.Close()
		}
		var data = Response{}
		data.StatusCode = statusCode
		data.URL = url
		datas = append(datas, data)
	}

	var res = Result{}
	res.Responses = datas
	outputJSON, err := json.Marshal(&res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stdout, string(outputJSON))
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
