package tasks

import (
	"time"
	"errors"
	"github.com/albrow/jobs"

	log "github.com/Sirupsen/logrus"
	"proof/plugs"
)

type Plug struct {
	Name string
	Type *jobs.Type
	Job  *jobs.Job
}

func (p *Plug) Register(retries uint) *jobs.Type {
	log.WithField("retries", retries).Debug("Registering `Plug` type Job")

	t, err := jobs.RegisterType(p.Name, retries, func(p *plugs.Payload) error {
		log.WithField("payload", p).Debug("Handler got called for Plug")

		v, k := plugs.Handlers[p.Name]

		if !k {
			err := errors.New("Couldn't find handler for Provided Plug.")
			log.WithError(err).Fatalf("Plug: `%s` not found.", p.Name)
			return err
		}

		v.(func(*plugs.Payload))(p)
		return nil
	})
	if err != nil {
		log.WithError(err).Fatal("Error registering job")
	}

	log.WithField("type", t).Debug("Type of the registered job")
	p.Type = t
	return t
}

func (p *Plug) Schedule(priority int, time time.Time, data interface{}) *jobs.Job {
	log.WithFields(log.Fields{
		"priority": priority,
		"time": time,
		"data": data,
	}).Debug("Scheduling a new Plug job")

	j, err := p.Type.Schedule(priority, time, data)
	if err != nil {
		log.WithError(err).Fatal("Error scheduling job")
	}
	p.Job = j
	return j
}

