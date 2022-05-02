package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
)

var count = 1 // count as row number

func main() {

	// read env file
	err := godotenv.Load("creds.env")
	if err != nil {
		log.Panicf("Error loading .env file")

	}

	wg := &sync.WaitGroup{} // wait group so program doesn't exit
	wg.Add(1)

	// define cron
	cronGetWeight := cron.New()

	fmt.Println("Program started, waiting for cronjob..")
	// start new cron
	cronGetWeight.AddFunc("* 23 * * * ", func() {

		// login and fetch token
		login()

		// fetch user ID
		getUserID()

		// Get data
		getData()

		weight := getLastWeight()
		date := getLastDate()

		fmt.Println("last date: " + getLastDate())
		fmt.Println("current weight: " + getLastWeight())

		uploadtosheets(count, weight, date)

		// increment count
		count++

		wg.Wait() // This guarantees this program never exits so cron can keep running as per the cron interval.
	})
}
