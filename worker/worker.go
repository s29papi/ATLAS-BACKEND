package worker

import (
	"log"
	"strconv"
	"time"

	"github.com/s29papi/wag3r-bot/worker/client"
	"github.com/s29papi/wag3r-bot/worker/env"
)

type Worker struct {
	T               *time.Ticker
	stopped         bool
	lastProcReqTime *int64
	s               *client.HTTPService
	Req             chan struct{}
	pauseTickerFn   chan struct{}

	done chan struct{}
}

func NewWorker() *Worker {
	val, err := strconv.Atoi(env.DURATION_STR)
	if err != nil {
		log.Fatalln("Error: conversion of DURATION_STR to int")
	}
	dur := time.Duration(val) * time.Second
	t := time.NewTicker(dur)
	s := client.NewHTTPService()
	return &Worker{
		T:               t,
		s:               s,
		Req:             make(chan struct{}),
		lastProcReqTime: new(int64),
		pauseTickerFn:   make(chan struct{}),

		done: make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go w.tick(w.T.C)
	go w.workloop()
	<-w.done

	log.Println("Bot Stopping")
	w.stopped = true

}

// update tick to start after processing is don
func (w *Worker) tick(t <-chan time.Time) {
	for {
		<-t
		log.Println("Tick go-routine:, new request initiated")
		w.Req <- struct{}{}
		log.Println("Tick go-routine: paused.")
		<-w.pauseTickerFn
	}
}

// stopping workloop mean stop processing requests
func (w *Worker) workloop() {
	for {
		<-w.Req
		log.Println("Initiating a new request")
		mentions := GetMentions(w.s)
		w.process(mentions)
		w.pauseTickerFn <- struct{}{}
		log.Println("Tick go-routine: un-paused.")
	}
}

func (w *Worker) Stop() {
	if w.stopped {
		log.Println("Can't stop Bot, Bot was just stopped")
		return
	}
	w.T.Stop()

	log.Println("Tick go-routine: stopped.")
	log.Println("Exiting...")
	w.done <- struct{}{}
}
