package server

import (
	"context"
	"io"
	"log"

	models "github.com/devjuank/go-protobuffers-grpc/models"
	"github.com/devjuank/go-protobuffers-grpc/repository"
	"github.com/devjuank/go-protobuffers-grpc/testpb"
)

type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{repo: repo}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := &models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}
	err := s.repo.SetTest(ctx, test)
	if err != nil {
		return nil, err
	}
	return &testpb.SetTestResponse{
		Id: test.Id,
	}, nil
}

func (s *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
			return err
		}

		question := &models.Question{
			Id:       msg.GetId(),
			Question: msg.GetQuestion(),
			Answer:   msg.GetAnswer(),
			TestId:   msg.GetTestId(),
		}

		err = s.repo.SetQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
	}
}
