package main

import (
	"bytes"
	"encoding/json"
	"fingbot/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

const fingReportFileName = "fing_report.json"

func sendFileToSever(fingReport model.FingReport) {

	serverUrl := os.Getenv("server_url")

	requestBody, err := json.Marshal(fingReport)

	if err != nil {
		fmt.Printf("%s", err)
	} else {

		fmt.Printf("%s: Sending file to server\n", time.Now().String())

		resp, err := http.Post(fmt.Sprintf("%s/send", serverUrl), "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from error: %s ", err)
			}
		}

		defer resp.Body.Close()
	}

}

func parseFile(wg *sync.WaitGroup) {
	for {
		fileContent, err := ioutil.ReadFile(fingReportFileName)

		if err == nil {
			var fingReport model.FingReport

			err = json.Unmarshal(fileContent, &fingReport)

			if err == nil {
				fingReport.BotDateTime = time.Now().String()

				sendFileToSever(fingReport)

			} else {
				fmt.Printf("%s", err)
			}
		} else {
			fmt.Printf("%s", err)
		}

		time.Sleep(20 * time.Second)
	}
}

func main() {

	commandStr := fmt.Sprintf("fing")
	commandArgs := fmt.Sprintf("-o table,json,%s", fingReportFileName)

	cmd := exec.Command(commandStr, commandArgs)

	err := cmd.Start()

	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go parseFile(&wg)

	wg.Wait()

	os.Exit(0)
}
