package tasks

import (
	"github.com/albrow/jobs"

	log "github.com/Sirupsen/logrus"
)

type Pool struct {
	instance *jobs.Pool
}

func NewPool(conf *jobs.PoolConfig) *Pool {
	log.WithField("conf", conf).Debug("Creating a new Pool")

	p, err := jobs.NewPool(conf)
	if err != nil {
		log.WithError(err).Fatal("Error creating Pool's instance.")
	}
	return &Pool{instance:p}
}

func (p *Pool) Start() {
	log.Debug("Starting Pool")

	if err := p.instance.Start(); err != nil {
		log.WithError(err).Fatal("Error starting the Pool.")
	}
}

func (p *Pool) Close() {
	log.Debug("Closing the Pool")

	p.instance.Close()
	if err := p.instance.Wait(); err != nil {
		log.WithError(err).Fatal("Error closing the Pool.")
	}
}