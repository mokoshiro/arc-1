package usecase

import (
	"os/exec"

	"fmt"

	"strings"

	"github.com/Bo0km4n/arc/pkg/tracker/api/proto"
	"github.com/Bo0km4n/arc/pkg/tracker/cmd/option"
	"github.com/Bo0km4n/arc/pkg/tracker/domain/repository"
	"github.com/Bo0km4n/arc/pkg/tracker/logger"
)

type RegisterUsecase interface {
	Register(req *proto.RegisterRequest) error
}

type registerUsecase struct {
	rr repository.RegisterRepository
}

func (ru *registerUsecase) Register(req *proto.RegisterRequest) error {
	// preprocess request
	hashStr, err := ru.invokeH3(option.Opt.GeoResolution, req.Longitude, req.Latitude)
	if err != nil {
		logger.L.Error(err.Error())
		return err
	}
	logger.L.Debug(hashStr)
	if err := ru.rr.Register(hashStr, req.Id, req.Longitude, req.Latitude); err != nil {
		logger.L.Error(err.Error())
		return err
	}
	return nil
}

func (ru *registerUsecase) invokeH3(resolution int, longitude, latitude float64) (string, error) {
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

func NewRegisterUsecase(rr repository.RegisterRepository) RegisterUsecase {
	return &registerUsecase{
		rr: rr,
	}
}
