package log

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbname      string = "timelog"
	timerecords string = "timerecords"
)

type MongoDB struct {
	serverlist string
}

func NewMongoDB(serverlist string) *MongoDB {
	return &MongoDB{serverlist}
}

func (d *MongoDB) GetRecords() []TimeRecord {
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

func (d *MongoDB) SaveRecords(records []TimeRecord) error {
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
