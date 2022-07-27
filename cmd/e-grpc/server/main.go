package main

import (
	"fmt"
	"github.com/dk-sirius/example/cmd/e-grpc/api"
	"github.com/dk-sirius/example/cmd/e-grpc/server/echo"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const port = ":8080"

func main() {
	p := port
	if len(os.Args) > 1 {
		p = ":" + os.Args[1]
	}
	fmt.Println(p)
	lis, err := net.Listen("tcp", p)
	if err != nil {
		log.Fatalf("listen port %v falt", p)
	}
	s := grpc.NewServer()
	api.RegisterEchoServer(s, &echo.EchoServie{Port: p})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("serve fatal %v", err.Error())
	}
}
