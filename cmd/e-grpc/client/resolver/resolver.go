package resolver

import (
	"google.golang.org/grpc/resolver"
	"log"
)

type RobinBuilder struct {
	AddrStore map[string][]string
	scheme    string
}

func NewRobinBuilder(addrs map[string][]string, scheme string) *RobinBuilder {
	return &RobinBuilder{
		AddrStore: addrs,
		scheme:    scheme,
	}
}

func (r *RobinBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	rb := &RobinResolver{
		target:  target,
		cc:      cc,
		address: r.AddrStore,
	}
	rb.start()
	return rb, nil
}

func (r *RobinBuilder) Scheme() string {
	return r.scheme
}

type RobinResolver struct {
	target  resolver.Target
	cc      resolver.ClientConn
	address map[string][]string
}

func (rb *RobinResolver) start() {
	if v, ok := rb.address[rb.target.URL.Host]; ok {
		store := make([]resolver.Address, len(v))
		for i, _ := range v {
			store[i] = resolver.Address{
				Addr: v[i],
			}
		}
		if len(store) > 0 {
			err := rb.cc.UpdateState(resolver.State{Addresses: store})
			if err != nil {
				rb.cc.ReportError(err)
				log.Fatalf("update state error %v", err)
			}
		}
	}
}

func (rb *RobinResolver) ResolveNow(opt resolver.ResolveNowOptions) {
	log.Printf("reslveNow %v", opt)
}

func (rb *RobinResolver) Close() {

}
