package models

import "testing"

func TestMaintKillMails(t *testing.T) {
	err := MaintKillMails()
	if err != nil {
		t.Error(err)
	}
}

func TestMaintMarket(t *testing.T) {
	err := MaintMarket()
	if err != nil {
		t.Error(err)
	}
}
