package beater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/future-architect/vuls/models"
	"github.com/kazuminn/vulsbeat/config"
)

// vulsbeat configuration.
type vulsbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of vulsbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &vulsbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts vulsbeat.
func (bt *vulsbeat) Run(b *beat.Beat) error {
	logp.Info("vulsbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	files := bt.dirwalk(bt.config.Path)

	results := models.ScanResults{}
	for _, file := range files {
		raw, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		var result models.ScanResult
		json.Unmarshal(raw, &result)

		results = append(results, result)
	}

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type":    b.Info.Name,
			"results": results,
		},
	}
	bt.client.Publish(event)
	logp.Info("Event sent")

		select {
		case <-bt.done:
			return nil
		}
}

// Stop stops vulsbeat.
func (bt *vulsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func (bt *vulsbeat) dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(bt.config.Path)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			if file.Name() == filepath.Join(bt.config.Path, "current") {
				continue
			}
			paths = append(paths, bt.dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
