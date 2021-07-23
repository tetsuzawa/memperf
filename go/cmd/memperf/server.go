package main

import (
	"encoding/json"
	"fmt"
	newrelic "github.com/newrelic/go-agent"
	"github.com/tetsuzawa/memperf"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var (
	frames       = map[int64]memperf.Frame{}
	readFrameNum = 1

	newrelicLicenceKey = ""
)

func Init() {
	rand.Seed(1)
	err := updateFrames()
	if err != nil {
		log.Fatalln(err)
	}

	newrelicLicenceKey = os.Getenv("NEWRELIC_LICENSE_KEY")
	if newrelicLicenceKey == "" {
		log.Fatalln("newrelicLicenceKey is not set")
	}
}

func Run() {
	app, err := newrelic.NewApplication(newrelic.NewConfig("memperf", newrelicLicenceKey))
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/ping", Ping))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/frame", FrameHandler))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/internal/update/frame", UpdateFramesHandler))

	log.Println("server is running at localhost:9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatalln(err)
	}
}

func updateFrames() error {
	log.Printf("reeding frames_%d.json\n", readFrameNum)
	f, err := os.Open(fmt.Sprintf("frames_%d.json", readFrameNum))
	defer f.Close()
	if err != nil {
		return err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	var framesDecorder memperf.Frames
	if err := json.Unmarshal(b, &framesDecorder); err != nil {
		return err
	}
	for _, frame := range framesDecorder.Items {
		frames[frame.Id] = frame
	}

	if readFrameNum == 1 {
		readFrameNum = 2
	} else {
		readFrameNum = 1
	}
	return nil
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func FrameHandler(w http.ResponseWriter, r *http.Request) {
	frameId := rand.Int63n(50000) + 1
	frame, ok := frames[frameId]
	if !ok {
		http.NotFound(w, r)
		return
	}
	b, err := json.Marshal(frame)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func UpdateFramesHandler(w http.ResponseWriter, r *http.Request) {
	if err := updateFrames(); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "updated frames")
}

func HandleWithAgent() {

}
