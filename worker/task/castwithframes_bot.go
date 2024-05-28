package task

import (
	"log"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/s29papi/atlas-backend/worker/client"
)

const (
	tickerDur = 18
)

type CastWithFramesBot struct {
	FC            *firestore.Client
	Client        *client.HTTPService
	wg            *sync.WaitGroup
	ticker        *time.Ticker
	fireReq       chan struct{}
	firePauseTick chan struct{}
	fireStop      chan struct{}
}

func (c *CastWithFramesBot) Duration() time.Duration {
	return time.Duration(tickerDur) * time.Second
}

func (c *CastWithFramesBot) StartProcess(intervals time.Duration) {
	var wg sync.WaitGroup
	c.wg = &wg
	c.ticker = time.NewTicker(intervals)
	c.fireReq = make(chan struct{})
	c.firePauseTick = make(chan struct{})
	c.fireStop = make(chan struct{})
	go c.tick(c.ticker.C)

	c.wg.Add(1)
	go c.process()
	c.wg.Wait()
	log.Println("Hello Task is done")
}

func (c *CastWithFramesBot) process() {
	defer c.wg.Done()

	for {
		select {
		case <-c.fireReq:
			c.internalTestProcess()
			c.firePauseTick <- struct{}{}
		case <-c.fireStop:
			log.Println("Hello Task is stopped")
			return
		}
	}
}

func (c *CastWithFramesBot) tick(t <-chan time.Time) {
	for {
		<-t
		log.Println("Tick go-routine:, new request initiated")
		c.fireReq <- struct{}{}
		log.Println("Tick go-routine: paused.")
		<-c.firePauseTick
	}
}

func (c *CastWithFramesBot) StopProcess() {
	c.ticker.Stop()
	c.fireStop <- struct{}{}
}

func (c *CastWithFramesBot) internalProcess() {

}

func (c *CastWithFramesBot) internalTestProcess() {
	casts := fetchCastWithFrames(c.Client)
	if len(casts.Data) == 0 {
		return
	}
	casts = filterOutRecentFramesCast(casts, c.FC)
	if len(casts.Data) == 0 {
		return
	}
	casts = filterOutNonValidFrames(casts)
	if len(casts.Data) == 0 {
		return
	}
	casts = filterOutNftFramesCast(casts)
	if len(casts.Data) == 0 {
		return
	}
	casts = filterOutNonPowerBadge(casts)
	if len(casts.Data) == 0 {
		return
	}
	casts = filterOutSameHostName(casts)
	if len(casts.Data) == 0 {
		return
	}
	timestamp, hash := sortHighestTimeAndHash(casts.Data)
	dataId := updateRecentFramesCast(casts, c.FC)
	updateCurrentReqInfo(c.FC, timestamp, hash, dataId)
}
