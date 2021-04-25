package main

import (
	"database/sql"
	"github.com/SubashMourougayane/grpc-todo-server/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type server struct { }

func (s server) CreateTodo(ctx context.Context, req *proto.CreateTodoReq) (*proto.CreateTodoRes, error) {
	panic("implement me")
}

func (s server) ReadTodo(ctx context.Context, req *proto.ReadTodoReq) (*proto.ReadTodoRes, error) {
	panic("implement me")
}

func (s server) UpdateTodo(ctx context.Context, req *proto.UpdateTodoReq) (*proto.UpdateTodoRes, error) {
	panic("implement me")
}

func (s server) DeleteTodo(ctx context.Context, req *proto.DeleteTodoReq) (*proto.DeleteTodoRes, error) {
	panic("implement me")
}

func (s server) ListTodo(ctx context.Context, req *proto.ListTodoReq) (*proto.ListTodoRes, error) {
	panic("implement me")
}

func main() {
	listener, err := net.Listen("tcp",":8080")

	if err != nil {
		panic("Server could not start in port 8080")
	}

	srv := grpc.NewServer()

	proto.RegisterTodoServiceServer(srv, &server{})
	reflection.Register(srv)

	db,err := sql.Open("mysql", "root:password@/todo")
	if err != nil{
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
