package main

import (
	"net/http"
	_ "net/http/pprof"

	"sync"

	"./potato"
	"github.com/robfig/cron"
)

func startProfile() {
	http.ListenAndServe(":6060", nil)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		cronVolume := cron.New()
		//cronVolume.AddFunc("* */8 * * * *", func() { potato.BdbCompaction() })
		cronVolume.AddFunc("*/5 * * * * *", func() { potato.Heartbeat() })
		cronVolume.AddFunc("*/3 * * * * *", func() { potato.RunReplicateParallel() })
		cronVolume.Start()
	}()

	go func() {
		potato.StartVolumeServer()
	}()

	go func() {
		potato.StartHttpServer()
	}()

	go func() {
		startProfile()
	}()

	potato.OnReady()

	wg.Wait()

}
