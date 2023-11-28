package main

import (
	"bytes"
	"context"
	pb "file_loader_proxy/proto"
	"file_loader_proxy/utils"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// weed server -s3  to start local s3 gateway

var s3Client = utils.NewS3Client()

type S3ServiceServer struct {
	pb.UnimplementedS3ServiceServer
}

func (s *S3ServiceServer) StoreFileS3(ctx context.Context, in *pb.S3FileStoreRequest) (*pb.S3FileStoreResponse, error) {
	_, err := s3Client.PutBucketItem(&utils.S3ObjectBucketInput{
		Bucket: "/test",
		Key:    in.Key,
		Body:   bytes.NewReader(in.FileData),
	})
	if err != nil {
		return nil, err
	}
	return &pb.S3FileStoreResponse{
		Success: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:17773")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterS3ServiceServer(grpcServer, &S3ServiceServer{})
	err = grpcServer.Serve(lis)

	fmt.Println("listening: on port 19777", err)
	if err != nil {
		panic(err.Error())
	}
}

//func main() {
//	f, err := os.Open("test_data/photo.jpg")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	allBytes, err := io.ReadAll(f)
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Println(allBytes)
//	//item, err := s3Client.PutBucketItem(utils.S3ObjectBucketInput{
//	//	Bucket: "/test",
//	//	Key:    "test_photo",
//	//	Body:   bytes.NewReader(allBytes),
//	//})
//	//if err != nil {
//	//	panic(err.Error())
//	//}
//	//fmt.Println(*item.ETag)
//}

// does not support v2
