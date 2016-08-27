package eveConsumer

import "log"

func (c *EVEConsumer) checkWars() {
	c.collectWarsFromCREST()
	c.updateWars()
}

func (c *EVEConsumer) updateWars() {
	rows, err := c.ctx.Db.Query(
		`SELECT id FROM eve.wars 
			WHERE (timeFinished = "0001-01-01 00:00:00" OR timeFinished IS NULL) 
			AND cacheUntil < UTC_TIMESTAMP()`)
	if err != nil {
		log.Printf("EVEConsumer: Failed getting wars: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {

		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Printf("EVEConsumer: Failed translating war ID: %v", err)
			continue
		}
		war, err := c.ctx.EVE.WarByID(id)
		if err != nil {
			log.Printf("EVEConsumer: Failed reading wars: %v", err)
			continue
		}

		_, err = c.ctx.Db.Exec(`UPDATE wars SET
					timeFinished=?, 
					openForAllies=?, 
					mutual=?, 
					cacheUntil=? WHERE id = ?;`,
			war.TimeFinished.String(),
			war.OpenForAllies,
			war.Mutual,
			war.CacheUntil, war.ID)
		if err != nil {
			log.Printf("EVEConsumer: Failed writing wars: %v", err)
			continue
		}

		for _, a := range war.Allies {
			_, err = c.ctx.Db.Exec(`INSERT IGNORE INTO warAllies (id, allyID) VALUES(?,?);`, war.ID, a.ID)
			if err != nil {
				log.Printf("EVEConsumer: Failed writing war allies: %v", err)
				continue
			}

			c.updateEntity(a.HRef, a.ID)
			if err != nil {
				log.Printf("EVEConsumer: Failed writing CREST ref: %v", err)
				continue
			}
		}
	}

}

func (c *EVEConsumer) collectWarsFromCREST() {
	r := struct {
		Value int
		Wait  int
	}{0, 0}

	if err := c.ctx.Db.Get(&r, `
		SELECT value, TIME_TO_SEC(TIMEDIFF(nextCheck, UTC_TIMESTAMP())) AS wait
			FROM states 
			WHERE state = 'wars'
			LIMIT 1;
		`); err != nil {
		log.Printf("EVEConsumer: Error checking state: %v", err)
		return
	}

	if r.Wait >= 0 {
		return
	}

	w, err := c.ctx.EVE.Wars(r.Value)

	if err != nil {
		log.Printf("EVEConsumer: Error checking wars: %v", err)
		return
	}

	// Loop through all of the pages
	for ; w != nil; w, err = w.NextPage() {

		// Update state so we dont have two polling at once.
		_, err = c.ctx.Db.Exec("UPDATE states SET value = ?, nextCheck =? WHERE state = 'wars' LIMIT 1", w.Page, w.CacheUntil)

		if err != nil {
			log.Printf("EVEConsumer: Could not update war state: %v", err)
			continue
		}

		for _, r := range w.Items {
			war, err := c.ctx.EVE.War(r.HRef)
			if err != nil {
				log.Printf("EVEConsumer: Failed reading wars: %v", err)
				continue
			}

			_, err = c.ctx.Db.Exec(`INSERT INTO wars
				(id, timeFinished,timeStarted,timeDeclared,openForAllies,cacheUntil,aggressorID,defenderID,mutual)
				VALUES(?,?,?,?,?,?,?,?,?)
				ON DUPLICATE KEY UPDATE 
					timeFinished=VALUES(timeFinished), 
					openForAllies=VALUES(openForAllies), 
					mutual=VALUES(mutual), 
					cacheUntil=VALUES(cacheUntil);`,
				war.ID, war.TimeFinished.String(), war.TimeStarted.String(), war.TimeDeclared.String(),
				war.OpenForAllies, war.CacheUntil, war.Aggressor.ID,
				war.Defender.ID, war.Mutual)
			if err != nil {
				log.Printf("EVEConsumer: Failed writing wars: %v", err)
				continue
			}

			err = c.updateEntity(war.Aggressor.HRef, war.Aggressor.ID)
			if err != nil {
				log.Printf("EVEConsumer: Failed writing CREST ref: %v", err)
				continue
			}

			err = c.updateEntity(war.Defender.HRef, war.Defender.ID)
			if err != nil {
				log.Printf("EVEConsumer: Failed writing CREST ref: %v", err)
				continue
			}

			for _, a := range war.Allies {
				_, err = c.ctx.Db.Exec(`INSERT IGNORE INTO warAllies (id, allyID) VALUES(?,?);`, war.ID, a.ID)
				if err != nil {
					log.Printf("EVEConsumer: Failed writing war allies: %v", err)
					continue
				}

				err = c.updateEntity(a.HRef, a.ID)
				if err != nil {
					log.Printf("EVEConsumer: Failed writing CREST ref: %v", err)
					continue
				}
			}
		}

	}
}
