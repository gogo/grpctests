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

package grpc

import (
	protoio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/test"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"testing"
	"time"
)

var msg = test.NewPopulatedNinOptStruct(rand.New(rand.NewSource(time.Now().UnixNano())), true)

type myBenchServer struct{}

func (this *myBenchServer) Down(m *MyRequest, s MyBench_DownServer) error {
	for {
		if err := s.Send(msg); err != nil {
			return err
		}
	}
}

type benchClient struct {
	writer protoio.WriteCloser
	reader protoio.ReadCloser
	conn   net.Conn
}

func newBenchClient(conn net.Conn) *benchClient {
	return &benchClient{
		writer: protoio.NewDelimitedWriter(conn),
		reader: protoio.NewDelimitedReader(conn, 1024*1024),
		conn:   conn,
	}
}

func (this *benchClient) Down(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (MyBench_DownClient, error) {
	if err := this.writer.WriteMsg(in); err != nil {
		return nil, err
	}
	return &benchDownClient{nil, this.reader}, nil
}

type benchDownClient struct {
	grpc.ClientStream
	reader protoio.ReadCloser
}

func (this *benchDownClient) Recv() (*test.NinOptStruct, error) {
	m := &test.NinOptStruct{}
	err := this.reader.ReadMsg(m)
	return m, err
}

func setupBenchTCP(t testing.TB, s MyBenchServer) (*benchServer, *benchClient) {
	serverConn, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	_, port, err := net.SplitHostPort(serverConn.Addr().String())
	if err != nil {
		log.Fatalf("Failed to parse listener address: %v", err)
	}
	addr := "localhost:" + port
	server := newBenchServer(serverConn, s)
	go server.Serve()
	clientConn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	client := newBenchClient(clientConn)
	return server, client
}

func setupBenchGRPC(t testing.TB, mybench MyBenchServer) (*grpc.Server, MyBenchClient) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	_, port, err := net.SplitHostPort(lis.Addr().String())
	if err != nil {
		log.Fatalf("Failed to parse listener address: %v", err)
	}
	s := grpc.NewServer(grpc.MaxConcurrentStreams(10))
	RegisterMyBenchServer(s, mybench)
	go s.Serve(lis)
	addr := "localhost:" + port
	conn, err := grpc.Dial(addr)
	if err != nil {
		log.Fatalf("Dial(%q) = %v", addr, err)
	}
	tc := NewMyBenchClient(conn)
	return s, tc
}

type benchServer struct {
	lis  net.Listener
	s    MyBenchServer
	conn net.Conn
}

func newBenchServer(lis net.Listener, s MyBenchServer) *benchServer {
	return &benchServer{
		lis: lis,
		s:   s,
	}
}

func (this *benchServer) Serve() {
	conn, err := this.lis.Accept()
	if err != nil {
		panic(err)
	}
	this.conn = conn
	reader := protoio.NewDelimitedReader(conn, 1024*1024)
	msg := &MyRequest{}
	err = reader.ReadMsg(msg)
	if err != nil {
		panic(err)
	}
	this.s.Down(msg, newBenchDownServer(conn))
	this.conn.Close()
}

func (this *benchServer) Stop() {
	this.conn.Close()
}

type benchDownServer struct {
	grpc.ServerStream
	conn protoio.WriteCloser
}

func newBenchDownServer(conn io.Writer) *benchDownServer {
	return &benchDownServer{nil, protoio.NewDelimitedWriter(conn)}
}

func (this *benchDownServer) Send(m *test.NinOptStruct) error {
	if err := this.conn.WriteMsg(m); err != nil {
		return err
	}
	return nil
}

func BenchmarkDownGRPC(b *testing.B) {
	server, client := setupBenchGRPC(b, &myBenchServer{})
	defer server.Stop()
	down, err := client.Down(context.Background(), &MyRequest{1})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = down.Recv()
	}
}

func BenchmarkDownTCP(b *testing.B) {
	server, client := setupBenchTCP(b, &myBenchServer{})
	defer server.Stop()
	down, err := client.Down(context.Background(), &MyRequest{1})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = down.Recv()
	}
}
