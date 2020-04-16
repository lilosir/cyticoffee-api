package main

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/lilosir/cyticoffee-api/routes"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   "https://e5dac4453b1c4561ab501d2ec66569ad@o304131.ingest.sentry.io/5202155",
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	r := routes.SetupRoutes()
	r.Run(":8090") // listen and serve on 0.0.0.0:8090
}
