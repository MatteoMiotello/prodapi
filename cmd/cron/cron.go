package main

import (
	"fmt"
	"github.com/MatteoMiotello/prodapi/internal/bootstrap"
	"github.com/MatteoMiotello/prodapi/internal/jobs"
	"github.com/MatteoMiotello/prodapi/internal/nosql"
	"github.com/robfig/cron/v3"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitMongoDb()
}

func main() {
	defer nosql.Disconnect()

	c := cron.New()
	addJobs(c)

	c.Start()

	fmt.Scanln()
}

func addJobs(c *cron.Cron) {
	_, err := c.AddFunc("@every 1s", jobs.DownloadNextImage)
	if err != nil {
		panic(err)
	}
}
