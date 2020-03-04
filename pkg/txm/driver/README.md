## Example Requests

```
curl -XPOST "http://localhost:8001/api/peer" -d \
'{
    "peer_id": "aaaa",
	"addr": "127.0.0.1:8080",
	"credential": "xxxx",
    "longitude": 127.00000,
    "latitude": 35.0000000
}'
```

```
curl -XPUT "http://localhost:8001/api/peer/location" -d \
'{
    "peer_id": "aaaa",
    "longitude": 127.00000,
    "latitude": 35.0000000
}'
```

```
curl -XDELETE "http://localhost:8001/api/peer" -d \
'{
    "peer_id": "aaaa"
}'
```