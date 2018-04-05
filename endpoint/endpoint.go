package endpoint

import (
	"context"

	svc "MiniProject/git.bluebird.id/mini/Agama/server"
	kit "github.com/go-kit/kit/endpoint"
)

type AgamaEndpoint struct {
	AddAgamaEndpoint              kit.Endpoint
	ReadAgamaByKeteranganEndpoint kit.Endpoint
	ReadAgamaEndpoint             kit.Endpoint
	UpdateAgamaEndpoint           kit.Endpoint
}

func NewAgamaEndpoint(service svc.AgamaService) AgamaEndpoint {
	addAgamaEp := makeAddAgamaEndpoint(service)
	readAgamaEp := makeReadAgamaEndpoint(service)
	updateAgamaEp := makeUpdateAgamaEndpoint(service)
	readAgamaByKeteranganEp := makeReadAgamaByKeteranganEndpoint(service)
	return AgamaEndpoint{AddAgamaEndpoint: addAgamaEp,
		ReadAgamaEndpoint:             readAgamaEp,
		UpdateAgamaEndpoint:           updateAgamaEp,
		ReadAgamaByKeteranganEndpoint: readAgamaByKeteranganEp,
	}
}

func makeAddAgamaEndpoint(service svc.AgamaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Agama)
		err := service.AddAgamaService(ctx, req)
		return nil, err
	}
}

func makeReadAgamaEndpoint(service svc.AgamaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadAgamaService(ctx)
		return result, err
	}
}

func makeUpdateAgamaEndpoint(service svc.AgamaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Agama)
		err := service.UpdateAgamaService(ctx, req)
		return nil, err
	}
}

func makeReadAgamaByKeteranganEndpoint(service svc.AgamaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Agama)
		result, err := service.ReadAgamaByKeteranganService(ctx, req.Keterangan)
		return result, err
	}
}
