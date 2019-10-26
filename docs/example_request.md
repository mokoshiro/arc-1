### Register

```
curl -X POST 'http://localhost:8080/api/member/register' -d 
'{
  "global_ip_addr": "127.0.0.1",
  "port": "7000",
  "location": {
    "latitude": 63,
    "longitude": 127
  },
  "id": "ffff"
}'
```

`latitude` and `longitude` are expressed by float64 value.

*Response*

```
{} // empty
```

### Get Member By Radius

```
curl -X GET 'http://localhost:8080/api/member?longitude=63.000&latitude=130.0000&radius=200&unit=km&with_coord=true'
```

`unit` takes a some case such as `m | km | ft | mi`


### Update Member location

```
curl -X PUT 'http://localhost:8080/api/member' -d '{"location": {"longitude": 125.77777, "latitude": 35.656565}, "id": "popo"}'
```

### Delete Member

```
curl -X DELETE 'http://localhost:8080/api/member' -d '{"id": "popo"}'
```
