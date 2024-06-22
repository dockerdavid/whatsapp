package jobs

import (
	"compi-whatsapp/pkg/queue"

	"github.com/robfig/cron"
)

var (
	cronJobs *cron.Cron
)

func InitCron() {
	cronJobs = cron.New()
	pendingFiles()
	cronJobs.Start()
}

func pendingFiles() {
	cronJobs.AddFunc("@every 1m", func() {
		queue.SendPendingFiles()
	})
}
