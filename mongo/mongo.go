package mongo

import (
	"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func mgoExec(collectionName string, f func(*mgo.Collection) error) error {
	session := GetSession()
	defer session.Close()
	collection := session.DB("").C(collectionName)

	return f(collection)
}

func Insert(collectionName string, docs ...interface{}) error {

	exec := func(collection *mgo.Collection) error {
		return collection.Insert(docs...)
	}

	return mgoExec(collectionName, exec)
}

func Update(collectionName string, selector bson.M, update interface{}) error {
	exec := func(collection *mgo.Collection) error {
		return collection.Update(selector, update)
	}

	return mgoExec(collectionName, exec)
}

func Find(collectionName string, query interface{}) (interface{}, error) {
	result := make([]interface{}, 0)

	exec := func(collection *mgo.Collection) error {
		return collection.Find(query).All(&result)
	}
	if err := mgoExec(collectionName, exec); err != nil {
		return nil, err
	}

	return &result, nil
}

func First(collectionName string, id interface{}) (interface{}, error) {
	var result interface{}

	exec := func(collection *mgo.Collection) error {
		return collection.FindId(id).One(&result)
	}

	if err := mgoExec(collectionName, exec); err != nil {
		return nil, err
	}

	return &result, nil
}

func Delete(collectionName string, id interface{}) error {
	switch idType := id.(type) {
	case string:
		id = bson.ObjectIdHex(id.(string))
	case bson.ObjectId:
	default:
		return errors.New("mgo delete expect type is string or bson.ObjectId。 not" + idType.(string))
	}

	exec := func(collection *mgo.Collection) error {
		return collection.RemoveId(id)
	}

	return mgoExec(collectionName, exec)
}
