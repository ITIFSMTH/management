package cron

import "github.com/robfig/cron/v3"

func InitCronTasks() {
	// New cron
	c := cron.New()

	// Close all shifts in 23:00
	c.AddFunc("0 23 * * *", CloseShifts)

	// Send operator's captcha's
	c.AddFunc("* * * * *", CheckShift)

	c.Start()
}
