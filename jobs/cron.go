package jobs

import "github.com/robfig/cron/v3"

var c *cron.Cron = cron.New()
var setup bool = false

func StartJobs() {
	if !setup {
		c.AddFunc("*/5 * * * *", CheckNamemc)
		setup = true
	}
	c.Start()
}

func StopJobs() {
	c.Stop()
}
