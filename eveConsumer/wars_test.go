package eveConsumer

import (
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestWarsTrigger(t *testing.T) {
	_, err := warsTrigger(eC)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestWarsConsumer(t *testing.T) {
	r := ctx.Cache.Get()
	defer r.Close()
	eC.warAddToQueue(1)
	eC.warAddToQueue(2)
	eC.warAddToQueue(3)
	for {
		_, err := warConsumer(eC, r)
		if err != nil {
			t.Error(err)
			return
		}
		if i, _ := redis.Int(r.Do("SCARD", "EVEDATA_warQueue")); i == 0 {
			break
		}
	}
}
