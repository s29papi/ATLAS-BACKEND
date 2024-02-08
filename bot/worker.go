package worker

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/s29papi/wag3r-bot/bot/client"
)

type Worker struct {
	T *time.Ticker

	s               *client.HTTPService
	Req             chan struct{}
	interrupt       <-chan os.Signal
	interruptTicker chan struct{}
	pause           chan struct{}
	startstop       chan struct{}
	done            chan struct{}
}

func NewWorker() *Worker {
	req_dur := os.Getenv("REQUEST_DURATION")
	// how can you create an alternate dev environment
	// if len(req_dur) == 0 {
	// 	req_dur = env.DURATION_STR
	// }
	val, err := strconv.Atoi(req_dur)
	if err != nil {
		log.Fatal("Error: conversion of DURATION_STR to int")
	}
	if err != nil {
		log.Fatal("Error: conversion of DURATION_STR to int")
	}
	dur := time.Duration(val) * time.Second
	t := time.NewTicker(dur)
	s := client.NewHTTPService()
	return &Worker{
		T: t,
		s: s,
		// interrupt:       interrupt,
		Req:             make(chan struct{}, 1),
		interruptTicker: make(chan struct{}),
		startstop:       make(chan struct{}),
		pause:           make(chan struct{}),
		done:            make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go w.tick(w.T.C)
	go w.workloop()
	go w.Stop()
	// <-w.done
	os.Exit(0)
}

// update tick to start after processing is don
func (w *Worker) tick(t <-chan time.Time) {
	for {
		select {
		case <-t:
			log.Println("Tick go-routine:, new request initiated")
			w.Req <- struct{}{}
			log.Println("Tick go-routine: paused.")
			<-w.pause

		// potential bug
		case <-w.interruptTicker:
			w.T.Stop()
			log.Println("Tick go-routine: stopped.")
			return
		}
	}

}

func (w *Worker) workloop() {
	for {
		select {
		case buff := <-w.s.Resp:
			w.process(buff)
		case <-w.Req:
			checkNewCast(w.s)
		case signal := <-w.interrupt:
			log.Printf("Received Interrupt signal: %v\n", signal)
			w.startstop <- struct{}{}
		}
	}
}

func (w *Worker) Stop() {
	<-w.startstop

	w.interruptTicker <- struct{}{}
	close(w.Req)
	close(w.startstop)
	log.Println("Exiting...")
	w.done <- struct{}{}
}

// check if a new cast has been added
func (w *Worker) process(d []byte) {
	fmt.Println(string(d))
	w.pause <- struct{}{}
}

func checkNewCast(service *client.HTTPService) {
	req := client.ChannelCastRequest()
	go service.SendRequest(http.MethodGet, req)
}
