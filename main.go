package main

import (
	"context"
	"flag"
	"fmt"
	golog "log"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/laonix/hopping-race-tracks/dispatcher"
	"github.com/laonix/hopping-race-tracks/input"
	"github.com/laonix/hopping-race-tracks/logger"
)

var (
	file   = flag.String("file", "default.txt", "input file path")
	config = flag.String("config", "default.yaml", "environment configuration file path")
)

func main() {
	flag.Parse()

	loadConfig()

	log := logger.Get()

	testCases, err := input.ParseTestCases(*file)
	if err != nil {
		log.Fatal(err, "failed to parse test cases", "file", *file)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	d := dispatcher.NewTestCaseDispatcher(
		ctx,
		dispatcher.WithDispatcherPipeSize(viper.GetInt("dispatcher.pipe.size")),
		dispatcher.WithDispatcherPoolSize(viper.GetInt("dispatcher.pool.size")),
		dispatcher.WithDispatcherLogger(log),
	)

	log.Debug("start processing test cases", "count", len(testCases))

	wg := &sync.WaitGroup{}

	for _, testCase := range testCases {
		wg.Add(1)
		d.Dispatch(testCase)
	}

	go func() {
		for result := range d.Results() {
			fmt.Println(result)
			log.Info("test case processed", "result", result)

			wg.Done()
		}
	}()

	wg.Wait()

	log.Debug("all test cases processed")

	d.Stop(ctx)
}

func loadConfig() {
	viper.SetConfigFile(*config)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		golog.Fatal(errors.Wrap(err, "read configuration file"))
	}

	viper.AutomaticEnv()
}
