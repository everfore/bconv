package bconv

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"
)

func Bson2BsonBytes(m *bson.M) (b []byte, err error) {
	return bson.Marshal(m)
}

func BsonBytes2Bson(b []byte) (ret *bson.M, err error) {
	err = bson.Unmarshal(b, &ret)
	return
}

func Bson2JsonBytes(m *bson.M) (b []byte, err error) {
	return json.Marshal(m)
}

func JsonBytes2Bson(b []byte) (ret *bson.M, err error) {
	err = json.Unmarshal(b, &ret)
	return
}

// JsonBytes to Json,Interface{}
func JsonBytes2Json(b []byte, i interface{}) (err error) {
	return json.Unmarshal(b, &i)
}

// Json,Interface to JsonBytes
func Json2JsonBytes(i interface{}) (b []byte, err error) {
	return json.Marshal(i)
}

// can not conv BsonBytes to Json, also can not conv Json to BsonBytes.
// failed
func _Json2BsonBytes(i interface{}) (b []byte, err error) {
	return bson.Marshal(i)
}

// failed
func _BsonBytes2Json(b []byte, i interface{}) (err error) {
	return bson.Unmarshal(b, &i)
}

// Bson to Json,InterfaceP{} by JsonBytes
func Bson2Json(m *bson.M, i interface{}) error {
	b, err := json.Marshal(m)
	if isError(err) {
		return err
	}
	err = json.Unmarshal(b, &i)
	if isError(err) {
		return err
	}
	return nil
}

// Bson to Json,InterfaceP{} by BsonBytes
// failed
func _Bson2Json(m *bson.M, i interface{}) error {
	b, err := bson.Marshal(m)
	if isError(err) {
		return err
	}
	err = bson.Unmarshal(b, &i)
	if isError(err) {
		return err
	}
	return nil
}

// Json,Interface{} to Bson by JsonBytes
// failed
// Document is corrupted
func _Json2Bson(i interface{}) (m *bson.M, err error) {
	b, err := json.Marshal(i)
	if isError(err) {
		return
	}
	err = bson.Unmarshal(b, &m)
	return
}

// Json,Interface{} to Bson by BsonBytes
func Json2Bson(i interface{}) (m *bson.M, err error) {
	b, err := bson.Marshal(i)
	if isError(err) {
		return
	}
	err = bson.Unmarshal(b, &m)
	return
}

// Interface{} to Interface{}
func I2I(j interface{}, i interface{}) error {
	b, err := json.Marshal(j)
	if isError(err) {
		return err
	}
	return json.Unmarshal(b, &i)
}

func isError(err error) bool {
	if nil != err {
		return true
	}
	return false
}
