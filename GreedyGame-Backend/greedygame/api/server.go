package api

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"encoding/json"

	"github.com/gorilla/mux"
)

type rootChild []*Country
type countryChild []*Devices

type RootNode struct {
	WebReq    int16 `json:"webreq"`
	TimeSpent int16 `json:"timespent"`
	Child     rootChild
}

type Country struct {
	Country   string `json:"country"`
	WebReq    int16  `json:"webreq"`
	TimeSpent int16  `json:"timespent"`
	Child     countryChild
}

type Devices struct {
	Devices   string `json:"devices"`
	WebReq    int16  `json:"webreq"`
	TimeSpent int16  `json:"timespent"`
}

type Res struct {
	Dim []struct {
		Key string `json:"key"`
		Val string `json:"val"`
	} `json:"dim"`
	Metrics []struct {
		Key string `json:"key"`
		Val int    `json:"val"`
	} `json:"metrics"`
}

var rootNode RootNode

func NewServer() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/insert", insert)
	r.HandleFunc("/query", query)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

// route functions below

var res Res

func insert(w http.ResponseWriter, r *http.Request) {
	fmt.Print("request recieved")
	json.NewDecoder(r.Body).Decode(&res)

}

func query(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(res)

	fmt.Print(res.Dim[0])
	fmt.Print(res.Metrics)

}
