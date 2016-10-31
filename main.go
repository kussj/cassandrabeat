package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/kussj/cassandrabeat/beater"
)

func main() {
	err := beat.Run("cassandrabeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
