package plugs

import (
	log "github.com/Sirupsen/logrus"
)

type Independent struct{}

func (i *Independent) Fetch(p *Payload) {
	log.WithField("payload", p).Debug("Fetching articles from Independent")
}