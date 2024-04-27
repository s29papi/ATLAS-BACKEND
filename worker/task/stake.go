package task

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/s29papi/wag3r-bot/worker/client"
)

const (
	stakeTickerDur = 12
)

type Stake struct {
	FC            *firestore.Client
	Client        *client.HTTPService
	wg            *sync.WaitGroup
	ticker        *time.Ticker
	fireReq       chan struct{}
	firePauseTick chan struct{}
	fireStop      chan struct{}
}

func (s *Stake) Duration() time.Duration {
	return time.Duration(stakeTickerDur) * time.Second
}

func (s *Stake) StartProcess(intervals time.Duration) {
	var wg sync.WaitGroup
	s.wg = &wg
	s.ticker = time.NewTicker(intervals)
	s.fireReq = make(chan struct{})
	s.firePauseTick = make(chan struct{})
	s.fireStop = make(chan struct{})
	go s.tick(s.ticker.C)

	s.wg.Add(1)
	go s.process()
	s.wg.Wait()
	log.Println("Hello Task is done")
}

func (s *Stake) process() {
	defer s.wg.Done()

	for {
		select {
		case <-s.fireReq:

			s.internalTestProcess()

			s.firePauseTick <- struct{}{}
		case <-s.fireStop:
			log.Println("Hello Task is stopped")
			return
		}
	}
}

func (s *Stake) tick(t <-chan time.Time) {
	for {
		<-t
		log.Println("Tick go-routine:, new request initiated")
		s.fireReq <- struct{}{}
		log.Println("Tick go-routine: paused.")
		<-s.firePauseTick
	}
}

func (s *Stake) StopProcess() {
	s.ticker.Stop()
	s.fireStop <- struct{}{}
}

func (s *Stake) internalProcess() {
	// fetch last request notifications time stamp
	// notifications := fetchToshiPayBotNotifications(s.Client)
	// lastnotificationsupdatetime := int64(0)
	// notifsByLastUpdate := filterNotificationsByLastUpdate(notifications, lastnotificationsupdatetime)
	// notifsByMention := filterNotificationsByMentions(notifsByLastUpdate)
	// txs := notifs2Tx(notifsByMention.Notifications)
	// payloads := getValidStakePayload(txs)
	// //  _=  castStakeReplies(s.Client, payloads)
	// fmt.Println(payloads)
	// update last request notifications timestamp
	// fmt.Println(notifications)
}

func (s *Stake) internalTestProcess() {

	notifications := fetchToshiPayBotNotifications(s.Client)
	lastnotificationsupdatetime := getCurrentReqTimeStamp(s.FC)
	fmt.Println(lastnotificationsupdatetime)
	if lastnotificationsupdatetime == 0 {
		return
	}
	notifsByLastUpdate := filterToshiPayBotNotificationsByLastUpdate(notifications, lastnotificationsupdatetime)
	notifsByNetwork := filterNotificationsByNetwork(notifsByLastUpdate)
	txs := notifs2Tx(s.Client, notifsByNetwork.Messages)
	if len(txs) == 0 {
		return
	}
	timestamp := sortHighestTime(txs)
	if lastnotificationsupdatetime == timestamp {
		return
	}
	payloads := getValidToshiPayPayload(txs)
	castToshiPayReplies(s.Client, payloads)

	updatelastnotificationsupdatetime(context.Background(), s.FC, timestamp)

}
