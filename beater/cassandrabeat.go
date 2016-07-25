package beater

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/goomzee/cassandrabeat/config"
)

type Cassandrabeat struct {
	done       chan struct{}
	config     config.Config
	client     publisher.Client

	table      []string
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Cassandrabeat{
		done: make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Cassandrabeat) Run(b *beat.Beat) error {
	logp.Info("cassandrabeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	bt.table = bt.config.Table[:]
	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		for _, table := range bt.table {
			logp.Info("Getting latency for table: %s", table)
			bt.getLatency(table)
		}
		logp.Info("Event sent")
	}
}

func (bt *Cassandrabeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func (bt *Cassandrabeat) getLatency(table string) {
	cmdName := "awkscript.sh"
	cmdArgs := []string{table}
	cmdOut := exec.Command(cmdName, cmdArgs...).Output

	output, _ := cmdOut()
	latency := strings.Split(string(output), "\n")

	var read_latency, write_latency float64
	if strings.Compare(latency[0], "NaN") == 0 {
		read_latency = 0.0
	} else {
		read_latency, _ = strconv.ParseFloat(latency[0], 64)
	}
	if strings.Compare(latency[1], "NaN") == 0 {
		write_latency = 0.0
	} else {
		write_latency, _ = strconv.ParseFloat(latency[1], 64)
	}

	event := common.MapStr {
		"@timestamp":	 common.Time(time.Now()),
		"type":		 "stats",
		"count":	 1,
		"table_name":	 table,
		"write_latency": write_latency,
		"read_latency":	 read_latency,
	}

	bt.client.PublishEvent(event)
}
