package IoTDM

import (
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type DatabaseHandle struct {
	session *mgo.Session
}

func GetDatabaseConnection(url string) (*DatabaseHandle, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &DatabaseHandle{session: session}, nil
}

func (handle *DatabaseHandle) GetProfile(identifier string) (Profile, error) {
	c := handle.session.DB("IoTDM").C("profiles")

	result := Profile{}
	err := c.Find(bson.M{"identifier": identifier}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
