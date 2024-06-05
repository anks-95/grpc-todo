package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/anks-95/todo/protogen/todo"
)

type server struct {
	pb.UnimplementedTodoServiceServer
	todos []*pb.Todo
}

func (s *server) PostTodo(ctx context.Context, new *pb.PostTodoRequest) (*pb.PostTodoResponse, error) {
	s.todos = append(s.todos, new.Todo)
	return &pb.PostTodoResponse{Success: true}, nil
}

func (s *server) ListTodos(ctx context.Context, list *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	return &pb.ListTodosResponse{Todos: s.todos}, nil
}

func main() {
	port, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 5002")
	if err := grpcServer.Serve(port); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
