package endpoint

import (
	"context"
	"time"

	svc "MiniProject/git.bluebird.id/mini/Agama/server"

	pb "MiniProject/git.bluebird.id/mini/Agama/grpc"

	util "MiniProject/git.bluebird.id/mini/util/grpc"
	disc "MiniProject/git.bluebird.id/mini/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.AgamaService"
)

func NewGRPCAgamaClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.AgamaService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addAgamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddAgamaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addAgamaEp = retry
	}

	var readAgamaByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadAgamaByKeteranganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readAgamaByKeteranganEp = retry
	}

	var readAgamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadAgamaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readAgamaEp = retry
	}

	var updateAgamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateAgama, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateAgamaEp = retry
	}
	return AgamaEndpoint{AddAgamaEndpoint: addAgamaEp,
		ReadAgamaByKeteranganEndpoint: readAgamaByKeteranganEp,
		ReadAgamaEndpoint:             readAgamaEp,
		UpdateAgamaEndpoint:           updateAgamaEp,
	}, nil
}

func encodeAddAgamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Agama)
	return &pb.AddAgamaReq{
		IDAgama:    req.IDAgama,
		NamaAgama:  req.NamaAgama,
		Status:     req.Status,
		Keterangan: req.Keterangan,
		CreateBy:   req.CreateBy,
	}, nil
}

func encodeReadAgamaByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Agama)
	return &pb.ReadAgamaByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func encodeReadAgamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateAgamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Agama)
	return &pb.UpdateAgamaReq{
		IDAgama:    req.IDAgama,
		NamaAgama:  req.NamaAgama,
		Status:     req.Status,
		Keterangan: req.Keterangan,
		UpdateBy:   req.UpdateBy,
	}, nil
}

func decodeAgamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadAgamaByKeteranganRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadAgamaByKeteranganResp)
	return svc.Agama{
		IDAgama:    resp.IDAgama,
		NamaAgama:  resp.NamaAgama,
		Status:     resp.Status,
		Keterangan: resp.Keterangan,
		CreateBy:   resp.CreateBy,
	}, nil
}

func decodeReadAgamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadAgamaResp)
	var rsp svc.Agamas

	for _, v := range resp.AllAgama {
		itm := svc.Agama{
			IDAgama:    v.IDAgama,
			NamaAgama:  v.NamaAgama,
			Status:     v.Status,
			Keterangan: v.Keterangan,
			CreateBy:   v.CreateBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddAgamaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddAgama",
		encodeAddAgamaRequest,
		decodeAgamaResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddAgama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddAgama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadAgamaByKeteranganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadAgamaByKeterangan",
		encodeReadAgamaByKeteranganRequest,
		decodeReadAgamaByKeteranganRespones,
		pb.ReadAgamaByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadAgamaByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadAgamaByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadAgamaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadAgama",
		encodeReadAgamaRequest,
		decodeReadAgamaResponse,
		pb.ReadAgamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadAgama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadAgama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateAgama(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateAgama",
		encodeUpdateAgamaRequest,
		decodeAgamaResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateAgama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateAgama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadAgamaByKeterangan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadAgamaByKeterangan",
		encodeReadAgamaByKeteranganRequest,
		decodeReadAgamaByKeteranganRespones,
		pb.ReadAgamaByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadAgamaByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadAgamaByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
