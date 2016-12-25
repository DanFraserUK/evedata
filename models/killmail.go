package models

import (
	"fmt"
	"time"
)

func AddKillmail(id int32, solarSystemID int32, killTime time.Time, victimCharacterID int32, victimCorporationID int32,
	victimAllianceID int32, hash string, attackerCount int, damageTaken int32, x float32, y float32, z float32,
	shipType int32, warID int32) error {
	if _, err := database.Exec(`
		INSERT INTO killmails
			(id,solarSystemID,killTime,victimCharacterID,victimCorporationID,victimAllianceID,hash,
			attackerCount,damageTaken,x,y,z,shipType,warID)
			VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?);
	`, id, solarSystemID, killTime, victimCharacterID, victimCorporationID, victimAllianceID, hash,
		attackerCount, damageTaken, x, y, z, shipType, warID); err != nil {
		return err
	}
	return nil
}

func AddKillmailAttacker(id int32, characterID int32, corporationID int32, allianceID int32, shipType int32,
	finalBlow bool, damageDone int32, weaponType int32, securityStatus float32) error {
	if _, err := database.Exec(`
		INSERT INTO killmailAttackers
			(id,characterID,corporationID,allianceID,shipType,finalBlow,damageDone,weaponType,securityStatus)
			VALUES(?,?,?,?,?,?,?,?,?);
	`, id, characterID, corporationID, allianceID, shipType, finalBlow, damageDone, weaponType, securityStatus); err != nil {
		return err
	}
	return nil
}

func AddKillmailItems(id int32, itemType int32, flag int32, quantityDestroyed int64, quantityDropped int64, singleton int32) error {
	if _, err := database.Exec(`
		INSERT INTO killmailItems
			(id,itemType,flag,quantityDestroyed,quantityDropped,singleton)
			VALUES(?,?,?,?,?,?);	
	`, id, itemType, flag, quantityDestroyed, quantityDropped, singleton); err != nil {
		return err
	}
	return nil
}

func GetKnownKillmails() ([]int64, error) {
	var known []int64
	if err := database.Select(&known, `SELECT id FROM killmails;`); err != nil {
		return nil, err
	}
	return known, nil
}

// [TODO] Break out the CSV into an array
type ConstellationActivity struct {
	Number            int64  `db:"number" json:"number"`
	SolarSystemIDs    string `db:"solarSystemIDs" json:"solarSystemIDs"`
	SolarSystemNames  string `db:"solarSystemNames" json:"solarSystemNames"`
	ConstellationID   int64  `db:"constellationID" json:"constellationID"`
	ConstellationName string `db:"constellationName" json:"constellationName"`
	RegionID          int64  `db:"regionID" json:"regionID"`
	RegionName        string `db:"regionName" json:"regionName"`
}

// [BENCHMARK] 0.015 sec / 0.000 sec
func GetConstellationActivity(id int64, entityType string) ([]ConstellationActivity, error) {
	v := []ConstellationActivity{}

	var victim, entity string

	switch entityType {
	case "corporation":
		victim = fmt.Sprintf("K.victimCorporationID=%d", id)
		entity = fmt.Sprintf("A.corporationID=%d", id)
	case "alliance":
		victim = fmt.Sprintf("K.victimAllianceID=%d", id)
		entity = fmt.Sprintf("A.allianceID=%d", id)
	case "character":
		victim = fmt.Sprintf("K.victimCharacterID=%d", id)
		entity = fmt.Sprintf("A.characterID=%d", id)
	}

	if err := database.Select(&v, `
	SELECT 	
		COUNT(DISTINCT K.id) AS number, 
	    GROUP_CONCAT(DISTINCT S.solarSystemID) AS solarSystemIDs,
	    GROUP_CONCAT(DISTINCT S.solarSystemName) AS solarSystemNames,
	    S.constellationID,
	    constellationName,
	    S.regionID,
	    regionName 
	FROM
	(SELECT K.id AS id, K.solarSystemID AS solarSystemID
		FROM killmails K 
		WHERE 
			K.killTime > DATE_SUB(UTC_TIMESTAMP(), INTERVAL 31 DAY) AND
			`+victim+`
	UNION ALL
		SELECT K.id, K.solarSystemID
			FROM killmails K 
			INNER JOIN killmailAttackers A ON A.id = K.id 
			WHERE 
				K.killTime > DATE_SUB(UTC_TIMESTAMP(), INTERVAL 31 DAY) AND
				`+entity+`
	) K
	INNER JOIN mapSolarSystems S ON K.solarSystemID = S.solarSystemID
	INNER JOIN mapConstellations C ON C.constellationID = S.constellationID
	INNER JOIN mapRegions R ON R.regionID = S.regionID
	GROUP BY S.constellationID`); err != nil {
		return nil, err
	}
	return v, nil
}

type KnownShipTypes struct {
	Number   int64  `db:"number" json:"number"`
	ShipType int64  `db:"shipType" json:"shipType"`
	ShipName string `db:"shipName" json:"shipName"`
}

// [BENCHMARK] 0.016 sec / 0.000 sec
func GetKnownShipTypes(id int64, entityType string) ([]KnownShipTypes, error) {
	v := []KnownShipTypes{}

	var victim, entity string

	switch entityType {
	case "corporation":
		victim = fmt.Sprintf("K.victimCorporationID=%d", id)
		entity = fmt.Sprintf("A.corporationID=%d", id)
	case "alliance":
		victim = fmt.Sprintf("K.victimAllianceID=%d", id)
		entity = fmt.Sprintf("A.allianceID=%d", id)
	case "character":
		victim = fmt.Sprintf("K.victimCharacterID=%d", id)
		entity = fmt.Sprintf("A.characterID=%d", id)
	}

	if err := database.Select(&v, `
	SELECT 	
		COUNT(DISTINCT K.id) AS number, 
	    shipType,
	    typeName AS shipName
	FROM
	(SELECT K.id AS id, K.shipType
		FROM killmails K 
		WHERE 
			K.killTime > DATE_SUB(UTC_TIMESTAMP(), INTERVAL 31 DAY) AND
			`+victim+`
			AND shipType > 0
	UNION ALL
		SELECT K.id, A.shipType
			FROM killmails K 
			INNER JOIN killmailAttackers A ON A.id = K.id 
			WHERE 
				K.killTime > DATE_SUB(UTC_TIMESTAMP(), INTERVAL 31 DAY) AND
				`+entity+`
                AND A.shipType > 0
	) K
	INNER JOIN invTypes T ON K.shipType = T.typeID
	GROUP BY shipType`); err != nil {
		return nil, err
	}
	return v, nil
}
