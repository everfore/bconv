package bconv

import (
	"labix.org/v2/mgo/bson"
	"testing"
)

type Person struct {
	Name string
}

var (
	testBson2 = &bson.M{
		"Name": "SHAALX",
	}
	testJson2 = &Person{
		Name: "SHAALX ?",
	}
)

func log_err(err error, t *testing.T) {
	if isError(err) {
		t.Error(err)
	}
}

func TestBson2BsonBytes(t *testing.T) {
	b, err := Bson2BsonBytes(testBson2)
	log_err(err, t)
	t.Log(b)
	t.Log(string(b)) // error string

	retBson, err := BsonBytes2Bson(b)
	log_err(err, t)
	t.Log(retBson)
}

func TestBson2JsonBytes(t *testing.T) {
	b, err := Bson2JsonBytes(testBson2)
	log_err(err, t)
	t.Log(b)
	t.Log(string(b)) //ok

	retBson, err := JsonBytes2Bson(b)
	log_err(err, t)
	t.Log(retBson)
}

func TestJson2JsonBytes(t *testing.T) {
	p := Person{Name: "IS SHAALX ?"}
	b, err := Json2JsonBytes(p)
	log_err(err, t)
	t.Log(b)
	t.Log(string(b))

	var j interface{}
	JsonBytes2Json(b, &j)
	log_err(err, t)
	t.Logf("%#v", j)
}

// // failed
// func TestJson2BsonBytes(t *testing.T) {
// 	p := Person{Name: "IS SHAALX."}
// 	b, err := _Json2BsonBytes(p)
// 	log_err(err, t)
// 	t.Log(b)
// 	t.Log(string(b))

// 	var j interface{}
// 	_BsonBytes2Json(b, &j)
// 	log_err(err, t)
// 	t.Logf("%#v", j)
// }

func TestBson2Json(t *testing.T) {
	var j Person
	err := Bson2Json(testBson2, &j)
	log_err(err, t)
	t.Logf("%#v", j)

	bson_, err := Json2Bson(j)
	log_err(err, t)
	t.Logf("%#v", bson_)
}

// // failed
// func TestBson2Json_(t *testing.T) {
// 	var j Person
// 	err := _Bson2Json(testBson, &j)
// 	log_err(err, t)
// 	t.Logf("%#v", j)

// 	bson_, err := Json2Bson(j)
// 	log_err(err, t)
// 	t.Logf("%#v", bson_)
// }

func TestJson2Bson(t *testing.T) {
	bson_, err := Json2Bson(testJson2)
	log_err(err, t)
	t.Logf("%#v", bson_)

	var p Person
	Bson2Json(bson_, &p)
	log_err(err, t)
	t.Logf("%#v", p)
}

// // failed
// func TestJson2Bson_(t *testing.T) {
// 	bson_, err := _Json2Bson(testJson)
// 	log_err(err, t)
// 	t.Logf("%#v", bson_)

// 	var p Person
// 	Bson2Json(bson_, &p)
// 	log_err(err, t)
// 	t.Logf("%#v", p)
// }

func TestI2I(t *testing.T) {
	var p Person
	err := I2I(testJson2, &p)
	log_err(err, t)
	t.Logf("%#v", p)

	err = I2I(testBson2, &p)
	log_err(err, t)
	t.Logf("%#v", p)
}
