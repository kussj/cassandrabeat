package beater

import (
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
)

type Cassandrabeat struct {
	period		time.Duration
	table		[]string
	CbConfig	ConfigSettings
	events		publisher.Client

	done		chan struct{}
}

func New() *Cassandrabeat {
	return &Cassandrabeat{}
}

func (cb *Cassandrabeat) Config(b *beat.Beat) error {
	err := cfgfile.Read(&cb.CbConfig, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	if cb.CbConfig.Input.Period != nil {
		cb.period = time.Duration(*cb.CbConfig.Input.Period) * time.Second
	} else {
		cb.period = 10 * time.Second
	}

	cb.table = cb.CbConfig.Input.Table[:]

	logp.Debug("cassandrabeat", "Init cassandrabeat")
	logp.Debug("cassandrabeat", "Period %v\n", cb.period)

	return nil
}

func (cb *Cassandrabeat) Setup(b *beat.Beat) error {
	cb.events = b.Publisher.Connect()
	cb.done = make(chan struct{})
	return nil
}

func (cb *Cassandrabeat) Run(b *beat.Beat) error {
	ticker := time.NewTicker(cb.period)
	defer ticker.Stop()

	var err error

	for {
		select {
			case <-cb.done:
				return nil
			case <-ticker.C:
		}

		timerStart := time.Now()

		for _, table := range cb.table {
			err = cb.exportTableStats(table)
			if err != nil {
				logp.Err("Error reading table stats: %v", err)
				break
			}
		}

		timerEnd := time.Now()
		duration := timerEnd.Sub(timerStart)
		if duration.Nanoseconds() > cb.period.Nanoseconds() {
			logp.Warn("Ignoring tick(s) due to processing taking longer than one period")
		}
	}

	return err
}

func (cb *Cassandrabeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (cb *Cassandrabeat) Stop() {
	close(cb.done)
}

func (cb *Cassandrabeat) exportTableStats(table string) error {
	read_latency, write_latency := getLatency(table)

	event := common.MapStr {
		"@timestamp":		common.Time(time.Now()),
		"type":			"stats",
		"count":		1,
		"table_name":		table,
		"write_latency":	write_latency,
		"read_latency" :	read_latency,
	}

	cb.events.PublishEvent(event)

	return nil
}

func getLatency(table string) (read, write float64) {
	cmdName := "awkscript.sh"
	cmdArgs := []string{table}
	cmdOut := exec.Command(cmdName, cmdArgs...).Output

	output, _ := cmdOut()
	latency := strings.Split(string(output), "\n")

	if strings.Compare(latency[0], "NaN") == 0 {
		read = 0.0
	} else {
		read, _ = strconv.ParseFloat(latency[0], 64)
	}
	if strings.Compare(latency[1], "NaN") == 0 {
		write = 0.0
	} else {
		write, _ = strconv.ParseFloat(latency[1], 64)
	}

	return
}

