package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"strconv"
	"sync"
)

var count int

func main() {

	// read env file
	err := godotenv.Load("vars.env")
	if err != nil {
		log.Panicf("Error loading .env file")

	}

	// convert count from env to int
	count, _ = strconv.Atoi(os.Getenv("row"))

	wg := &sync.WaitGroup{} // wait group so program doesn't exit
	wg.Add(1)

	// define cron
	cronGetWeight := cron.New()

	fmt.Println("Program started, waiting for cronjob..")
	// start new cron
	cronGetWeight.AddFunc("0 23 * * * ", func() {
		fmt.Println("Starting cronjob...")

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
		fmt.Println("Cronjob finished.")

	})

	cronGetWeight.Start()

	wg.Wait() // This guarantees this program never exits so cron can keep running as per the cron interval.

}
