package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	// _ "github.com/s29papi/wag3r-bot/bot/env"
)

func main() {
	startBot()
	startServer()
}

func startBot() {
	req_dur := os.Getenv("REQUEST_DURATION")
	// how can you create an alternate dev environment
	// if len(req_dur) == 0 {
	// 	req_dur = env.DURATION_STR
	// }
	val, err := strconv.Atoi(req_dur)
	if err != nil {
		log.Fatal("Error: conversion of DURATION_STR to int")
	}
	dur := time.Duration(val) * time.Second
	t := time.NewTicker(dur)

	go func(t <-chan time.Time) {
		for {
			select {
			case <-t:
				log.Println("Hello world")
			}
		}
	}(t.C)
}

func startServer() {
	http.ListenAndServe(":8090", nil)
}

// i would need to get the logs of this service
/**
* we have a single go routine
* bot.Start() --> 3 go routines ---> a channel and an os.Exit()
* 1st go routine is called tick. has a for loop and select statement which blocks
* pending when its channels recieve a value
*
*
*
 */
//  curl http://arena.wield.co:2281/v1/castsByFid?fid=2

// this bot is on render
// we have a free hub
// a frontend vercel
//
