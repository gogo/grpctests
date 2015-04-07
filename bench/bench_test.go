// Copyright (c) 2015, Vastech SA (PTY) LTD. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package bench

import (
	generated "github.com/gogo/grpctests/bench/combos/both"
	reflected "github.com/gogo/grpctests/bench/combos/neither"
	"github.com/gogo/protobuf/codec"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"testing"
	"time"
)

var seed = time.Now().UnixNano()

type genServer struct{}

func (this *genServer) BigStream(m *generated.Request, s generated.Bencher_BigStreamServer) error {
	msg := generated.NewPopulatedBig(rand.New(rand.NewSource(seed)), true)
	for {
		if err := s.Send(msg); err != nil {
			return err
		}
	}
}

func (this *genServer) MediumStream(m *generated.Request, s generated.Bencher_MediumStreamServer) error {
	msg := generated.NewPopulatedMedium(rand.New(rand.NewSource(seed)), true)
	for {
		if err := s.Send(msg); err != nil {
			return err
		}
	}
}

func (this *genServer) SmallStream(m *generated.Request, s generated.Bencher_SmallStreamServer) error {
	msg := generated.NewPopulatedSmall(rand.New(rand.NewSource(seed)), true)
	for {
		if err := s.Send(msg); err != nil {
			return err
		}
	}
}

func setupGen() (*grpc.Server, generated.BencherClient) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	_, port, err := net.SplitHostPort(lis.Addr().String())
	if err != nil {
		log.Fatalf("Failed to parse listener address: %v", err)
	}
	s := grpc.NewServer(grpc.CustomCodec(codec.New(1024 * 1024)))
	generated.RegisterBencherServer(s, &genServer{})
	go s.Serve(lis)
	addr := "localhost:" + port
	conn, err := grpc.Dial(addr, grpc.WithCodec(codec.New(1024*1024)))
	if err != nil {
		log.Fatalf("Dial(%q) = %v", addr, err)
	}
	tc := generated.NewBencherClient(conn)
	return s, tc
}

type refServer struct{}

func (this *refServer) BigStream(m *reflected.Request, s reflected.Bencher_BigStreamServer) error {
	msg := reflected.NewPopulatedBig(rand.New(rand.NewSource(seed)), true)
	for {
		if err := s.Send(msg); err != nil {
			return err
		}
	}
}

func (this *refServer) MediumStream(m *reflected.Request, s reflected.Bencher_MediumStreamServer) error {
	msg := reflected.NewPopulatedMedium(rand.New(rand.NewSource(seed)), true)
	for {
		if err := s.Send(msg); err != nil {
			return err
		}
	}
}

func (this *refServer) SmallStream(m *reflected.Request, s reflected.Bencher_SmallStreamServer) error {
	msg := reflected.NewPopulatedSmall(rand.New(rand.NewSource(seed)), true)
	for {
		if err := s.Send(msg); err != nil {
			return err
		}
	}
}

func setupRef() (*grpc.Server, reflected.BencherClient) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	_, port, err := net.SplitHostPort(lis.Addr().String())
	if err != nil {
		log.Fatalf("Failed to parse listener address: %v", err)
	}
	s := grpc.NewServer(grpc.MaxConcurrentStreams(1))
	reflected.RegisterBencherServer(s, &refServer{})
	go s.Serve(lis)
	addr := "localhost:" + port
	conn, err := grpc.Dial(addr)
	if err != nil {
		log.Fatalf("Dial(%q) = %v", addr, err)
	}
	tc := reflected.NewBencherClient(conn)
	return s, tc
}

func BenchmarkDownBigGen(b *testing.B) {
	server, client := setupGen()
	defer server.Stop()
	down, err := client.BigStream(context.Background(), &generated.Request{1})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = down.Recv()
	}
}

func BenchmarkDownBigRef(b *testing.B) {
	server, client := setupRef()
	defer server.Stop()
	down, err := client.BigStream(context.Background(), &reflected.Request{1})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = down.Recv()
	}
}

func BenchmarkDownMediumGen(b *testing.B) {
	server, client := setupGen()
	defer server.Stop()
	down, err := client.MediumStream(context.Background(), &generated.Request{1})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = down.Recv()
	}
}

func BenchmarkDownMediumRef(b *testing.B) {
	server, client := setupRef()
	defer server.Stop()
	down, err := client.MediumStream(context.Background(), &reflected.Request{1})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = down.Recv()
	}
}

func BenchmarkDownSmallGen(b *testing.B) {
	server, client := setupGen()
	defer server.Stop()
	down, err := client.SmallStream(context.Background(), &generated.Request{1})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = down.Recv()
	}
}

func BenchmarkDownSmallRef(b *testing.B) {
	server, client := setupRef()
	defer server.Stop()
	down, err := client.SmallStream(context.Background(), &reflected.Request{1})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = down.Recv()
	}
}
