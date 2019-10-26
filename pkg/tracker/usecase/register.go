package usecase

import (
	"errors"
	"os/exec"

	"fmt"

	"strings"

	"context"

	"github.com/Bo0km4n/arc/pkg/tracker/api/proto"
	"github.com/Bo0km4n/arc/pkg/tracker/cmd/option"
	"github.com/Bo0km4n/arc/pkg/tracker/domain/repository"
	"github.com/Bo0km4n/arc/pkg/tracker/logger"
)

type MemberUsecase interface {
	Register(req *proto.RegisterRequest) error
	GetMemberByRadius(ctx context.Context, req *proto.GetMemberByRadiusRequest) (*proto.GetMemberByRadiusResponse, error)
	Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error)
}

type memberUsecase struct {
	mr repository.MemberRepository
}

func (ru *memberUsecase) Register(req *proto.RegisterRequest) error {
	// preprocess request
	hashStr, err := ru.invokeH3(option.Opt.GeoResolution, req.Longitude, req.Latitude)
	if err != nil {
		logger.L.Error(err.Error())
		return err
	}
	logger.L.Debug(hashStr)
	if err := ru.mr.Register(hashStr, req.Id, req.Longitude, req.Latitude); err != nil {
		logger.L.Error(err.Error())
		return err
	}
	return nil
}

func (mu *memberUsecase) GetMemberByRadius(ctx context.Context, req *proto.GetMemberByRadiusRequest) (*proto.GetMemberByRadiusResponse, error) {
	hashStr, err := mu.invokeH3(option.Opt.GeoResolution, req.Longitude, req.Latitude)
	if err != nil {
		logger.L.Error(err.Error())
		return nil, err
	}
	var unit string
	switch req.Unit {
	case proto.GetMemberByRadiusRequest_KM:
		unit = "km"
	case proto.GetMemberByRadiusRequest_M:
		unit = "m"
	case proto.GetMemberByRadiusRequest_MI:
		unit = "mi"
	case proto.GetMemberByRadiusRequest_FT:
		unit = "ft"
	default:
		return nil, errors.New("Not matched unit name")
	}

	res, err := mu.mr.GetMemberByRadius(
		hashStr, req.Longitude,
		req.Latitude, req.Radius, unit)
	if err != nil {
		logger.L.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (ru *memberUsecase) invokeH3(resolution int, longitude, latitude float64) (string, error) {
	args := []string{
		"--resolution", fmt.Sprintf("%d", resolution),
		"--longitude", fmt.Sprintf("%f", longitude),
		"--latitude", fmt.Sprintf("%f", latitude),
	}
	command := exec.Command("geoToH3", args...)
	b, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	result := strings.TrimRight(string(b), "\n")
	return result, err
}

func (mu *memberUsecase) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	hashStr, err := mu.invokeH3(option.Opt.GeoResolution, req.Longitude, req.Latitude)
	if err != nil {
		logger.L.Error(err.Error())
		return nil, err
	}
	if err := mu.mr.Update(hashStr, req.PeerId, req.Longitude, req.Latitude); err != nil {
		return nil, err
	}
	return &proto.UpdateResponse{}, nil
}

func NewMemberUsecase(mr repository.MemberRepository) MemberUsecase {
	return &memberUsecase{
		mr: mr,
	}
}
