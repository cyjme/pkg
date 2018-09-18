package mongo

import (
	"github.com/globalsign/mgo"
	"log"
)

var Session *mgo.Session

//param dialInfo
//dialInfo := &mgo.DialInfo{
//	Addrs:    []string{"127.0.0.1:27017"},
//	Username: "root",
//	Password: "root",
//	Database: "user-center",
//	Source:   "admin",
//}
func InitMgo(dialInfo *mgo.DialInfo) {
	session, err := mgo.DialWithInfo(dialInfo)
	Session = session
	if err != nil {
		log.Println("mongo dial error", err)
		panic(err)
	}

	log.Println("mongo dial success")
	session.DB(dialInfo.Database)

	session.SetPoolLimit(5000)
}

func GetSession() *mgo.Session {
	return Session.Clone()
}
