package endpoint

import (
	"context"
	"fmt"

	scv "MiniProject/git.bluebird.id/mini/Agama/server"

	pb "MiniProject/git.bluebird.id/mini/Agama/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcAgamaServer struct {
	addAgama              grpctransport.Handler
	readAgamaByKeterangan grpctransport.Handler
	readAgama             grpctransport.Handler
	updateAgama           grpctransport.Handler
}

func NewGRPCAgamaServer(endpoints AgamaEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.AgamaServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcAgamaServer{
		addAgama: grpctransport.NewServer(endpoints.AddAgamaEndpoint,
			decodeAddAgamaRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddAgama", logger)))...),
		readAgamaByKeterangan: grpctransport.NewServer(endpoints.ReadAgamaByKeteranganEndpoint,
			decodeReadAgamaByKeteranganRequest,
			encodeReadAgamaByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadAgamaByKeterangan", logger)))...),
		readAgama: grpctransport.NewServer(endpoints.ReadAgamaEndpoint,
			decodeReadAgamaRequest,
			encodeReadAgamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadAgama", logger)))...),
		updateAgama: grpctransport.NewServer(endpoints.UpdateAgamaEndpoint,
			decodeUpdateAgamaRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateAgama", logger)))...),
	}
}

func decodeAddAgamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddAgamaReq)
	return scv.Agama{IDAgama: req.GetIDAgama(), NamaAgama: req.GetNamaAgama(),
		Status: req.GetStatus(), Keterangan: req.GetKeterangan(), CreateBy: req.GetCreateBy()}, nil
}

func decodeReadAgamaByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadAgamaByKeteranganReq)
	return scv.Agama{Keterangan: req.Keterangan}, nil
}

func decodeReadAgamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateAgamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateAgamaReq)
	return scv.Agama{
		IDAgama:    req.IDAgama,
		NamaAgama:  req.NamaAgama,
		Status:     req.Status,
		Keterangan: req.Keterangan,
		UpdateBy:   req.UpdateBy}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadAgamaByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Agama)
	fmt.Println("server")
	return &pb.ReadAgamaByKeteranganResp{IDAgama: resp.IDAgama, NamaAgama: resp.NamaAgama, Status: resp.Status,
		Keterangan: resp.Keterangan, CreateBy: resp.CreateBy}, nil
	//return &pb.ReadAgamaByKeteranganResp{}, nil

	//rsp := &pb.ReadAgamaByKeteranganResp{}
	//fmt.Println("server", rsp)
	// for _, v := range resp {
	// itm := &pb.ReadAgamaByKeteranganResp{
	// 	IDAgama:    v.IDAgama,
	// 	NamaAgama:  v.NamaAgama,
	// 	Status:     v.Status,
	// 	Keterangan: v.Keterangan,
	// 	CreateBy:   v.CreateBy,
	// }
	//rsp.AllAgama = append(rsp.AllAgama, itm)
	// }
	//return resp, nil

}

func encodeReadAgamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Agamas)

	rsp := &pb.ReadAgamaResp{}

	for _, v := range resp {
		itm := &pb.ReadAgamaByKeteranganResp{
			IDAgama:    v.IDAgama,
			NamaAgama:  v.NamaAgama,
			Status:     v.Status,
			Keterangan: v.Keterangan,
			CreateBy:   v.CreateBy,
		}
		rsp.AllAgama = append(rsp.AllAgama, itm)
	}
	return rsp, nil
}

func (s *grpcAgamaServer) AddAgama(ctx oldcontext.Context, agama *pb.AddAgamaReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addAgama.ServeGRPC(ctx, agama)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcAgamaServer) ReadAgamaByKeterangan(ctx oldcontext.Context, keterangan *pb.ReadAgamaByKeteranganReq) (*pb.ReadAgamaByKeteranganResp, error) {
	_, resp, err := s.readAgamaByKeterangan.ServeGRPC(ctx, keterangan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadAgamaByKeteranganResp), nil
}

func (s *grpcAgamaServer) ReadAgama(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadAgamaResp, error) {
	_, resp, err := s.readAgama.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadAgamaResp), nil
}

func (s *grpcAgamaServer) UpdateAgama(ctx oldcontext.Context, cus *pb.UpdateAgamaReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateAgama.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
