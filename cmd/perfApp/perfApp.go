package main

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rsevilla87/perfapp/internal/perf"
	"github.com/rsevilla87/perfapp/pkg/euler"
	"github.com/rsevilla87/perfapp/pkg/timestamp"
	"github.com/rsevilla87/perfapp/pkg/health"
	"github.com/rsevilla87/perfapp/pkg/utils"
)

var tables []map[string]string

func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	go handleInterrupt(c)
	if os.Getenv("POSTGRESQL_RETRY_INTERVAL") != "" {
		retryInt, err := strconv.Atoi(os.Getenv("POSTGRESQL_RETRY_INTERVAL"))
		if err != nil {
			utils.ErrorHandler(err)
		}
		perf.DB.RetryInt = retryInt
	}
	perf.Connect2Db()
	tables = append(tables, euler.Tables, timestamp.Tables)
	if err := perf.CreateTables(tables); err != nil {
		utils.ErrorHandler(err)
	}
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/euler", euler.Handler)
	http.HandleFunc("/ready", timestamp.Handler)
	http.HandleFunc("/health", health.Handler)
	log.Printf("Listening at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		utils.ErrorHandler(err)
	}
}

func handleInterrupt(c <-chan os.Signal) {
	<-c
	log.Println("Interrupt signal received")
	if err := perf.DropTables(tables); err != nil {
		utils.ErrorHandler(err)
	}
	os.Exit(0)
}
