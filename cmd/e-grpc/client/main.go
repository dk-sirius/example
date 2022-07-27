package main

import (
	"context"
	"fmt"
	"github.com/dk-sirius/example/cmd/e-grpc/api"
	resolver2 "github.com/dk-sirius/example/cmd/e-grpc/client/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"os"
	"time"
)

func main() {
	resolver.Register(resolver2.NewRobinBuilder(map[string][]string{"echo": {
		"localhost:8080", "localhost:8081",
	}}, "srv-echo"))
	srvHost := "srv-echo://echo"
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for i := 10; i > 0; i-- {
				go request(srvHost)
			}
		case <-time.After(2 * time.Second):
			goto LOOP
		}

	}
LOOP:
	fmt.Println("finish")
}

func request(host string) {
	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	conn, err := grpc.Dial(host, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connecting ...%v", err)
	}
	defer conn.Close()
	c := api.NewEchoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := c.Say(ctx, &api.RequestEcho{
		Name: name,
	})
	if err != nil {
		log.Fatalf("fatal at %v", err.Error())
	}
	log.Printf("Echo \t%v", resp.GetName())
}
