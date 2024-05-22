package worker

import (
	"context"
	_ "database/sql"
	"encoding/hex"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/s29papi/atlas-backend/worker/client"
	"github.com/s29papi/atlas-backend/worker/env"
	"github.com/s29papi/atlas-backend/worker/task"
	"google.golang.org/api/option"
)

type Worker struct {
	ctx     context.Context
	fc      *firestore.Client
	s       *client.HTTPService
	tasks   []task.Task
	stopped bool
	done    chan struct{}

	lastProcReqTime *int64
}

func NewWorker() *Worker {
	ctx := context.Background()
	fireStoreCred, err := hex.DecodeString(env.SERVICE_ACCOUNT_JSON_HEX)
	if err != nil {
		log.Fatalln("Error decoding hexadecimal representation:", err)
	}
	opt := option.WithCredentialsJSON(fireStoreCred)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	fc, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	s := client.NewHTTPService()

	return &Worker{
		ctx:             ctx,
		fc:              fc,
		s:               s,
		lastProcReqTime: new(int64),
		done:            make(chan struct{}),
	}
}

func (w *Worker) Start() {
	w.RegisterTasks()

	for i := 0; i < len(w.tasks); i++ {
		intervals := w.tasks[i].Duration()
		go w.tasks[i].StartProcess(intervals)
	}

	<-w.done

	log.Println("Worker Stopping...")
	w.stopped = true
}

func (w *Worker) Stop() {
	if w.stopped {
		log.Println("Can't stop Worker, Worker was just stopped")
		return
	}
	for i := 0; i < len(w.tasks); i++ {
		go w.tasks[i].StopProcess()
	}

	w.fc.Close()

	log.Println("DB closed.")
	log.Println("Tick go-routine: stopped.")
	log.Println("Exiting...")

	w.done <- struct{}{}
}

func (w *Worker) RegisterTasks() {
	stake := task.CastWithFramesBot{
		FC:     w.fc,
		Client: w.s,
	}
	w.tasks = append(w.tasks, &stake)
}
