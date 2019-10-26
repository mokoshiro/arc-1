package repository

import (
	"strconv"

	"github.com/Bo0km4n/arc/pkg/tracker/api/proto"
	"github.com/Bo0km4n/arc/pkg/tracker/cmd/option"
	"github.com/Bo0km4n/arc/pkg/tracker/infra/db"
	"github.com/garyburd/redigo/redis"
)

type MemberRepository interface {
	Register(h3Hash, peerID string, longitude, latitude float64) error
	GetMemberByRadius(
		h3Hash string, longitude,
		latitude float64, radius float64, unit string) (*proto.GetMemberByRadiusResponse, error)
	Update(h3Hash, peerID string, longitude, latitude float64) error
}

type memberKVSRepository struct {
	redisPool *redis.Pool
}

type registerRDBRepository struct {
	// *sql.DB
}

func newRegisterKVSRepository(redisPool *redis.Pool) MemberRepository {
	return &memberKVSRepository{
		redisPool: redisPool,
	}
}

func NewMemberRepository(dbType int) MemberRepository {
	switch dbType {
	case db.DB_REDIS:
		return newRegisterKVSRepository(db.RedisPool)
	default:
		return nil
	}
}

func (mr *memberKVSRepository) Register(h3Hash, peerID string, longitude, latitude float64) error {
	conn := mr.redisPool.Get()
	defer conn.Close()
	if _, err := conn.Do("GEOADD", h3Hash, longitude, latitude, peerID); err != nil {
		return err
	}
	_, err := conn.Do("EXPIRE", h3Hash, option.Opt.RedisKeyExpire)
	return err
}

func (mr *memberKVSRepository) GetMemberByRadius(
	h3Hash string, longitude,
	latitude float64, radius float64,
	unit string) (*proto.GetMemberByRadiusResponse, error) {
	conn := mr.redisPool.Get()
	defer conn.Close()

	res, err := redis.Values(
		conn.Do("GEORADIUS", h3Hash, longitude, latitude, radius, unit, "WITHCOORD"),
	)
	if err != nil {
		return nil, err
	}

	result := make([]*proto.TrackingMember, len(res))

	// result is expressed by follow:
	// 127.0.0.1:16379> GEORADIUS 8915414e4cfffff 127 63 100 km WITHCOORD
	// 1) 1) "aaaa"
	// 2) 1) "126.99999779462814331"
	// 2) "62.99999887738297843"
	// 2) 1) "bbbb"
	// 2) 1) "126.99999779462814331"
	// 2) "62.99999887738297843"
	// 3) 1) "dddd"
	// 2) 1) "126.99999779462814331"
	// 2) "62.99999887738297843"
	// 4) 1) "eeee"
	// 2) 1) "126.99999779462814331"
	// 2) "62.99999887738297843"
	// 5) 1) "ffff"
	// 2) 1) "126.99999779462814331"
	// 2) "62.99999887738297843"
	//
	// Interface value includes some fields as uint8 slice.
	// Longitude and Latitude are passed as string float value converted u8 slice.
	// So, in the below code, it converts interface value to []byte slice,
	// then convert to each type once again.
	for i, v := range res {
		member := &proto.TrackingMember{}
		arrV := v.([]interface{})
		member.PeerId = string(arrV[0].([]byte))
		coord := arrV[1].([]interface{})
		longitudeRaw := string(coord[0].([]byte))
		latitudeRaw := string(coord[1].([]byte))
		member.Longitude, _ = strconv.ParseFloat(longitudeRaw, 64)
		member.Latitude, _ = strconv.ParseFloat(latitudeRaw, 64)
		result[i] = member
	}
	return &proto.GetMemberByRadiusResponse{Members: result}, nil
}

func (mr *memberKVSRepository) Update(h3Hash, peerID string, longitude, latitude float64) error {
	conn := mr.redisPool.Get()
	defer conn.Close()
	if _, err := conn.Do("GEOADD", h3Hash, longitude, latitude, peerID); err != nil {
		return err
	}
	_, err := conn.Do("EXPIRE", h3Hash, option.Opt.RedisKeyExpire)
	return err
}
