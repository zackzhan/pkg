package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/assert"
)

func TestCron(t *testing.T) {
	cronExpression := "0 0 8 1/1 * ?"
	now := time.Now()

	specParser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	mailExpireCron, err := specParser.Parse(cronExpression)
	assert.NoError(t, err, nil)

	nextTime := mailExpireCron.Next(now)
	interval := nextTime.Sub(now)

	fmt.Println("Next execution time:", nextTime)
	fmt.Println("Interval until next execution:", interval)
}
