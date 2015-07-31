# 	bconv

Bytes Convert

----------------------------------

##	Define

json represents for interface{}, for instance struct, map, slice.

bson represents for bson.M, actually which is a map[string]interface{}.

## 	Convert Conclusion

###		json <----> bson

json to bson, using bson.Marshal

bson to json, using json.Marshal

###		jsonbytes <--X--> bsonbytes

Well, jsonbytes is not bsonbytes. They are different []byte.

We can not use string(bsonbytes) as string.