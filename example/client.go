package main

import (
	"context"
	"time"

	cli "MiniProject/git.bluebird.id/mini/Agama/endpoint"
	svc "MiniProject/git.bluebird.id/mini/Agama/server"
	opt "MiniProject/git.bluebird.id/mini/util/grpc"
	util "MiniProject/git.bluebird.id/mini/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCAgamaClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add
	client.AddAgamaService(context.Background(), svc.Agama{IDAgama: "mm69", NamaAgama: "Kejawen", CreateBy: "agung", Keterangan: "Kalem"})

	//Get list
	//cusAgama, _ := client.ReadAgamaService(context.Background())
	//fmt.Println(" all agamas:", cusAgama)

	//Get agama by parameter
	//	parameter := "be%"
	//	cuss, _ := client.ReadAgamaByKeteranganService(context.Background(), parameter)
	//	fmt.Println("daftar agama  berdasarkan keterangan:", cuss)

	//Update
	//	client.UpdateAgamaService(context.Background(), svc.Agama{NamaAgama: "Hindu", Status: 1, UpdateBy: "agung", IDAgama: "6ho"})

	/*Get Customer By Email
	cusEmail, _ := client.ReadAgamaByEmailService(context.Background(), "sayaagung58@gmail.com")
	fmt.Println("customer based on email:", cusEmail)*/
}
