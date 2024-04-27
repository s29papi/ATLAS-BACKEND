package renderDoNotWindDown

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/s29papi/wag3r-bot/worker/env"
)

type RenderDoNotWindDown struct {
	req                     *http.Request
	c                       *http.Client
	wg                      *sync.WaitGroup
	stopRenderDoNotWindDown chan struct{}
	done                    chan struct{}
}

var (
	dev_endpoint_renderDoNotWindDown  = "http://localhost:8181/api/worker/dummy-rdWD"
	prod_endpoint_renderDoNotWindDown = "https://toshi-pay-bot.onrender.com/api/worker/dummy-rdWD"
)

func NewRenderDoNotWindDown(dev bool, client *http.Client) *RenderDoNotWindDown {
	var wg sync.WaitGroup
	var req *http.Request
	var err error
	s := make(chan struct{})
	d := make(chan struct{})
	if dev {
		req, err = http.NewRequest("GET", dev_endpoint_renderDoNotWindDown, nil)
		if err != nil {
			log.Fatalln("Error creating request:", err)
		}
	} else {
		req, err = http.NewRequest("GET", prod_endpoint_renderDoNotWindDown, nil)
		if err != nil {
			log.Fatalln("Error creating request:", err)
		}
	}
	authValue := "Bearer " + env.AUTH_TOKEN
	req.Header.Add("Authorization", authValue)

	return &RenderDoNotWindDown{req: req, c: client, wg: &wg, stopRenderDoNotWindDown: s, done: d}
}

func (rDWD *RenderDoNotWindDown) Start() {
	dur := time.Duration(30) * time.Second
	t := time.NewTicker(dur)
	rDWD.wg.Add(1)
	go renderDoNotWindDownLoop(rDWD.c, rDWD.req, t.C, rDWD.stopRenderDoNotWindDown, rDWD.wg)
	rDWD.wg.Wait()
	t.Stop()
	rDWD.done <- struct{}{}
}

func (rDWD *RenderDoNotWindDown) Stop() {
	rDWD.stopRenderDoNotWindDown <- struct{}{}
	<-rDWD.done
}

func renderDoNotWindDownLoop(client *http.Client, req *http.Request, t <-chan time.Time, c <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-t:
			log.Println("Tick go-routine:, new request render do not wind down initiated")
			resp, err := client.Do(req)
			if err != nil {
				log.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()
			log.Printf("Render Do Not Wind Down Status: Asking is render_do_not_wind_down active, Responce: %v\n", resp.Status)
		case <-c:
			return
		}
	}
}
