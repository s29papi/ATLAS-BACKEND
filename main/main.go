package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	// _ "github.com/s29papi/wag3r-bot/bot/env"
)

func main() {
	val, err := strconv.Atoi(os.Getenv("REQUEST_DURATION"))
	if err != nil {
		log.Fatal("Error: conversion of DURATION_STR to int")
	}
	dur := time.Duration(val) * time.Second
	t := time.NewTicker(dur)
	fmt.Println(t)
	fmt.Println(dur)
	go func(t <-chan time.Time) {
		for {
			select {
			case <-t:
				log.Println("Hello world")
			}
		}
	}(t.C)
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
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
