package log

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbname      string = "timelog"
	timerecords string = "timerecords"
)

type DB struct {
	serverlist string
}

func NewDB(serverlist string) *DB {
	return &DB{serverlist}
}

func (d *DB) GetRecords() []TimeRecord {
	session, err := mgo.Dial(d.serverlist)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	r := session.DB(dbname).C(timerecords)
	result := []TimeRecord{}
	err = r.Find(bson.M{}).All(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func (d *DB) SaveRecords(records []TimeRecord) error {
	session, err := mgo.Dial(d.serverlist)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	r := session.DB(dbname).C(timerecords)

	for _, record := range records {
		err = r.Insert(record)
		if err != nil {
			return err
		}
	}
	return nil
}
