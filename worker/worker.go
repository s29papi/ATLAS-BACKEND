package worker

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/s29papi/wag3r-bot/worker/client"
	"github.com/s29papi/wag3r-bot/worker/env"
)

type Worker struct {
	T *time.Ticker

	lastProcReqTime *int64
	s               *client.HTTPService
	Req             chan struct{}
	interrupt       <-chan os.Signal
	interruptTicker chan struct{}
	pauseFn         chan struct{}
	startStopFn     chan struct{}
	done            chan struct{}
}

func NewWorker(signalChan <-chan os.Signal) *Worker {
	val, err := strconv.Atoi(env.DURATION_STR)
	if err != nil {
		log.Fatal("Error: conversion of DURATION_STR to int")
	}
	dur := time.Duration(val) * time.Second
	t := time.NewTicker(dur)
	s := client.NewHTTPService()
	return &Worker{
		T:               t,
		s:               s,
		Req:             make(chan struct{}, 1),
		lastProcReqTime: new(int64),
		interrupt:       signalChan,
		interruptTicker: make(chan struct{}),
		startStopFn:     make(chan struct{}),
		pauseFn:         make(chan struct{}),
		done:            make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go w.tick(w.T.C)
	go w.workloop()
	go w.Stop()
	<-w.done
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
			<-w.pauseFn
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
		case <-w.Req:
			fmt.Println("sent a request")
			mentions := GetMentions(w.s)
			w.process(mentions)
		case <-w.interrupt:
			log.Printf("Received Interrupt signal.")
			w.startStopFn <- struct{}{}
		}
	}
}

func (w *Worker) Stop() {
	<-w.startStopFn

	w.interruptTicker <- struct{}{}
	close(w.Req)
	close(w.startStopFn)
	log.Println("Exiting...")
	w.done <- struct{}{}
}
