package api

import (
	"context"

	"sync"

	"github.com/Bo0km4n/arc/pkg/arclet/api/proto"
	"github.com/Bo0km4n/arc/pkg/arclet/resources"
)

type ArcletServer struct {
	TrackerResources []*resources.Tracker
	DBResources      []*resources.DB
	sync.RWMutex
}

func CreateTracker(c context.Context, in *proto.Empty) (*proto.CreateTrackerResponse, error) {
	return nil, nil
}

func Run() {

}
