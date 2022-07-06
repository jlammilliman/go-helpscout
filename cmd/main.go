package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	helpscout "github.com/jlammilliman/go-helpscout/pkg/api"
)

func main() {

	convos, currPage, pages, count, err := hs.Conversations.ListMailboxFolderConversations(123, 321, 0, "", nil, "")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(convos)
	log.Println(currPage)
	log.Println(pages)
	log.Println(count)
	log.Println(err)

	// Construct the application logger.

	hs := helpscout.New()
	r := api.ServiceRouter(hs)

	fmt.Println("Starting VPS Processor")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9069"
	}

	go func() {
		if err := http.ListenAndServe((":" + port), r); err != nil {
			log.Fatal("[ERROR] Error starting http listener. %s", err)
		}
	}()

	// Start the server and the processing go routines.
	hs.Start()

	// Create a channel to block routine and when a signal is received shutdown the server cleanly.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// We received an interrupt signal, shut down gracefully.
	if err := hs.Shutdown(ctx); err != nil {
		fmt.Printf("Processor server shutdown error: %v\n", err)
	}
}
