package eveConsumer

import (
	"log"
	"time"

	"github.com/antihax/evedata/appContext"
)

// EVEConsumer provides the microservice which conducts backend
// polling of EVE Crest and XML resources as needed.
type EVEConsumer struct {
	ctx                 *appContext.AppContext
	consumerStopChannel chan bool
	triggersStopChannel chan bool
}

// NewEVEConsumer creates a new EveConsumer
func NewEVEConsumer(ctx *appContext.AppContext) *EVEConsumer {
	e := &EVEConsumer{ctx, make(chan bool), make(chan bool)}
	return e
}

func (c *EVEConsumer) goConsumer() {
	r := c.ctx.Cache.Get()
	defer r.Close()

	// Run Phase
	for {
		workDone := false
		select {
		case <-c.consumerStopChannel:
			return
		default:
			if v, err := c.killmailCheckQueue(r); v != "" && err == nil {
				err := c.killmailConsume(v, r)
				workDone = true
				if err != nil {
					log.Printf("EVEConsumer: %v\n", err)
				}
			} else if err != nil {
				log.Printf("EVEConsumer: %v\n", err)
			}

			if v, err := c.entityCheckQueue(r); v > 0 && err == nil {
				err := c.entityConsume(v, r)
				workDone = true
				if err != nil {
					log.Printf("EVEConsumer: %v\n", err)
				}
			} else if err != nil {
				log.Printf("EVEConsumer: %v\n", err)
			}

			if v, err := c.marketOrderCheckQueue(r); v > 0 && err == nil {
				err := c.marketOrderConsume(v, r)
				workDone = true
				if err != nil {
					log.Printf("EVEConsumer: %v\n", err)
				}
			} else if err != nil {
				log.Printf("EVEConsumer: %v\n", err)
			}

			if v, err := c.marketHistoryCheckQueue(r); v != "" && err == nil {
				err := c.marketHistoryConsume(v, r)
				workDone = true
				if err != nil {
					log.Printf("EVEConsumer: %v\n", err)
				}
			} else if err != nil {
				log.Printf("EVEConsumer: %v\n", err)
			}

			// This really isnt much work.
			if err := c.marketRegionCheckQueue(r); err != nil {
				log.Printf("EVEConsumer: %v\n", err)
			}
		}

		// Sleep a brief bit if we didnt do anything
		if workDone == false {
			time.Sleep(time.Second * 5)
		}
	}
}

func (c *EVEConsumer) goTriggers() {
	log.Printf("EVEConsumer: Running Triggers\n")
	rate := time.Second * 60 * 30
	throttle := time.Tick(rate)
	for {
		select {
		case <-c.triggersStopChannel:
			log.Printf("EVEConsumer: Shutting Down\n")
			return
		default:
			c.checkWars()
			c.marketHistoryUpdateTrigger()
			c.contactSync()
			c.checkPublicStructures()
			c.checkNPCCorps()
			c.checkEntities()
		}
		<-throttle
	}
}

// Load deferrable data.
func (c *EVEConsumer) initConsumer() {
	r := c.ctx.Cache.Get()
	defer r.Close()
	// Load Phase
	c.initKillConsumer(r)
}

// RunConsumer starts the consumer and returns.
func (c *EVEConsumer) RunConsumer() {
	// Load deferrable data.
	go c.initConsumer()

	for i := 0; i < c.ctx.Conf.EVEConsumer.Consumers; i++ {
		go c.goConsumer() // Run consumers in a loop
	}

	go c.goTriggers() // Time triggered queries
	if c.ctx.Conf.EVEConsumer.ZKillEnabled == true {
		go c.goZKillConsumer()
		go c.goZKillTemporaryConsumer()
	}

	log.Printf("EVEConsumer: Started\n")
}

// StopConsumer shuts down any running go routines and returns.
func (c *EVEConsumer) StopConsumer() {
	log.Printf("EVEConsumer: Stopping Consumer\n")
	for i := 0; i > c.ctx.Conf.EVEConsumer.Consumers; i++ {
		c.consumerStopChannel <- true
	}
	c.triggersStopChannel <- true
	log.Printf("EVEConsumer: Stopped\n")
}
