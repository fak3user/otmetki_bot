package schedule

import (
	cron "github.com/robfig/cron/v3"
)

func InitSchedule() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("0 0 11,16,20 * * *", func() {

	})

	c.Start()
}
