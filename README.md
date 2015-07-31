# 	bconv

Bytes Convert

----------------------------------

##	Get

go get github.com/everfore/bconv

##	Define

json represents for interface{}, for instance struct, map, slice.

bson represents for bson.M, actually which is a map[string]interface{}.

## 	Convert Conclusion

###		json <----> bson

json to bson, using bson.Marshal

bson to json, using json.Marshal

###		jsonbytes <--X--> bsonbytes

Well, jsonbytes do not equal bsonbytes. They are different []byte.

We can not use string(bsonbytes) as string.

##		Usage

###		Init

		var testBson = &bson.M{
			"Name": "Golang",
		}

		b := NewBson(testBson)


		type Lang struct {
			Name string
		}

		var testJson = &Lang{
			Name: "Golang ?",
		}

		j := NewJson(testJson)


###		Bson

		b := NewBson(testBson)

		bbbs := b.BBytes()
		t.Log(bbbs)

		bjbs := b.JBytes()
		t.Log(bjbs)

		bj := b.Json()
		t.Log(bj)

###		Json

		j := NewJson(testJson)

		jbbs := j.BBytes()
		t.Log(jbbs)

		jjbs := j.JBytes()
		t.Log(jjbs)

		jb := j.Bson()
		t.Log(jb)

##		Dependency pkg from 3rd party

		go get labix.org/v2/mgo/bson