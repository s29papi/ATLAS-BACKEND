package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/s29papi/wag3r-bot/bot/env"
	// _ "github.com/s29papi/wag3r-bot/bot/env"
)

func main() {
	startBot()
	startServer()
}

var id = 318902

func startBot() {
	val, err := strconv.Atoi(env.DURATION_STR)
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
//  curl http://hub.freefarcasterhub.com:2281/v1/castsByMention?fid=318902

// we have a free hub
// a frontend vercel
//
// curl http://arena.wield.co/v1/castsByMention?fid=318902
