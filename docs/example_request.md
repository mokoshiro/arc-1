### Register

```
curl -X POST 'http://localhost:8000/api/member' -d \
'{
  "addr": "127.0.0.1:9000",
  "latitude": 35.11981010,
  "longitude": 135.100000,
  "peer_id": "aaaa"
}'
```

`latitude` and `longitude` are expressed by float64 value.

*Response*

```
{} // empty
```

### Get Member By Radius

```
curl -X GET 'http://localhost:8000/api/member' -d \
'{
  "peer_id": "aaaa",
  "longitude": 134.0000,
  "latitude": 34.0000,
  "radius": 100,
  "unit": "km"
}'
```

`unit` takes a some case such as `m | km | ft | mi`


### Update Member location

```
curl -X PUT 'http://localhost:8000/api/member' -d \
'{
  "peer_id": "aaaa",
  "longitude": 134.0000,
  "latitude": 34.0000
}'
```

### Delete Member

```
curl -X DELETE 'http://localhost:8080/api/member' -d '{"id": "popo"}'
```

### Signaling Request

```
curl -X POST 'http://localhost:8000/api/room/notification' -d \
'{
  "peers": [
    "aaaa",
    "bbbb",
    "cccc",
    "dddd"
  ]
}
'
```

