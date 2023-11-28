package utils

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

var ctx = context.Background()

type ClientS3 struct {
	*s3.S3
}

type S3ObjectBucketInput struct {
	Bucket string
	Key    string
	Body   io.ReadSeeker
}

func NewS3Client() ClientS3 {
	sess := session.Must(session.NewSession())

	s3c := s3.New(sess, &aws.Config{
		Credentials: credentials.NewStaticCredentials("98yh5r23", "9o03r2e8uh", ""),
		Endpoint:    aws.String("http://127.0.0.1:8333/"),
		Region:      aws.String("us-east-1"),
	})

	return ClientS3{s3c}
}

func (sC *ClientS3) PutBucketItem(i *S3ObjectBucketInput) (*s3.PutObjectOutput, error) {
	objectPut, err := sC.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(i.Bucket),
		Key:    aws.String(i.Key),
		Body:   i.Body,
	})
	if err != nil {
		return nil, err
	}
	return objectPut, nil
}

func (sC *ClientS3) GetBucketItem(g *S3ObjectBucketInput) (*s3.GetObjectOutput, error) {
	getObjectRes, err := sC.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(g.Bucket),
		Key:    aws.String(g.Key),
	})
	if err != nil {
		return nil, err
	}
	return getObjectRes, nil
}

//func (sC *ClientS3) GetBucketItems() []types.Bucket {
//	buckets, err := sC.PutObject(ctx, &s3.PutObjectInput{
//		Bucket: nil,
//		Key:    nil,
//		Body:   nil,
//	})
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	return buckets.Buckets
//}
