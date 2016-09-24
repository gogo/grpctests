// Copyright (c) 2015, Vastech SA (PTY) LTD. All rights reserved.
// http://github.com/gogo/protobuf/gogoproto
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
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"testing"
)

type aServer struct{}

func (this *aServer) UnaryCall(c context.Context, s *MyRequest) (*MyResponse, error) {
	return &MyResponse{s.Value}, nil
}
func (this *aServer) Downstream(m *MyRequest, s MyTest_DownstreamServer) error {
	for i := 0; i < int(m.Value); i++ {
		err := s.Send(&MyMsg{int64(i)})
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *aServer) Upstream(s MyTest_UpstreamServer) error {
	rec, err := s.Recv()
	sum := int64(0)
	for err == nil {
		sum += rec.Value
		rec, err = s.Recv()
	}
	return s.SendAndClose(&MyResponse{sum})
}
func (this *aServer) Bidi(b MyTest_BidiServer) error {
	var err error
	msg := &MyMsg{}
	for {
		msg, err = b.Recv()
		if err != nil {
			break
		}
		err = b.Send(&MyMsg2{msg.Value})
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return nil
	}
	return err
}

type bServer struct{}

func (this *bServer) UnaryCall(c context.Context, s *MyRequest) (*MyResponse, error) {
	return nil, nil
}
func (this *bServer) Downstream(m *MyRequest, s MyTest_DownstreamServer) error {
	mymsg := &MyMsg{1}
	for {
		if err := s.Send(mymsg); err != nil {
			return err
		}
	}
}
func (this *bServer) Upstream(s MyTest_UpstreamServer) error {
	return nil
}
func (this *bServer) Bidi(b MyTest_BidiServer) error {
	return nil
}

func setup(t testing.TB, mytest MyTestServer) (*grpc.Server, MyTestClient) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}
	_, port, err := net.SplitHostPort(lis.Addr().String())
	if err != nil {
		t.Fatalf("Failed to parse listener address: %v", err)
	}
	s := grpc.NewServer(grpc.MaxConcurrentStreams(1))
	RegisterMyTestServer(s, mytest)
	go s.Serve(lis)
	addr := "localhost:" + port
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithDecompressor(grpc.NewGZIPDecompressor()))
	if err != nil {
		t.Fatalf("Dial(%q) = %v", addr, err)
	}
	tc := NewMyTestClient(conn)
	return s, tc
}

func TestUnary(t *testing.T) {
	server, client := setup(t, &aServer{})
	defer server.Stop()
	want := int64(5)
	reply, err := client.UnaryCall(context.Background(), &MyRequest{want})
	if err != nil || reply.Value != want {
		t.Fatalf("UnaryCall(_, _) = %d, %v, want %d, <nil>", reply.Value, err, want)
	}
}

func TestDownstream(t *testing.T) {
	server, client := setup(t, &aServer{})
	defer server.Stop()
	num := int64(10)
	down, err := client.Downstream(context.Background(), &MyRequest{num})
	if err != nil {
		t.Fatal(err)
	}
	m := &MyMsg{}
	last := -1
	for err == nil {
		m, err = down.Recv()
		if err == nil {
			if int(m.Value) == (last + 1) {
				last++
				if last == int(num)-1 {
					break
				}
			} else {
				t.Errorf("out of order last = %d this = %d", last, m.Value)
			}
		}
	}
	if last != int(num)-1 {
		t.Fatalf("wrong last %d expected %d", last, num-1)
	}
}

func total(n int) int {
	t := 0
	for i := 0; i < n; i++ {
		t += i
	}
	return t
}

func TestUpstream(t *testing.T) {
	server, client := setup(t, &aServer{})
	defer server.Stop()
	num := 10
	up, err := client.Upstream(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < num; i++ {
		err := up.Send(&MyMsg{int64(i)})
		if err != nil {
			t.Fatal(err)
		}
	}
	res, err := up.CloseAndRecv()
	if err != nil {
		t.Fatal(err)
	}
	sum := total(num)
	if res.Value != int64(sum) {
		t.Fatalf("expected %d got %d", sum, res.Value)
	}
}

func TestBidi(t *testing.T) {
	server, client := setup(t, &aServer{})
	defer server.Stop()
	num := 10
	bidi, err := client.Bidi(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		for i := 0; i < num; i++ {
			err := bidi.Send(&MyMsg{int64(i)})
			if err != nil {
				t.Fatal(err)
			}
		}
	}()
	for i := 0; i < num; i++ {
		msg, err := bidi.Recv()
		if err != nil {
			t.Fatal(err)
		}
		if msg.Value != int64(i) {
			t.Fatalf("expected %d got %d", i, msg.Value)
		}
	}
}

func setupTCP(t testing.TB, s MyTestServer) (*server, *client) {
	serverConn, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	_, port, err := net.SplitHostPort(serverConn.Addr().String())
	if err != nil {
		log.Fatalf("Failed to parse listener address: %v", err)
	}
	addr := "localhost:" + port
	server := newServer(serverConn, s)
	go server.Serve()
	clientConn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	client := NewClient(clientConn)
	return server, client
}

type server struct {
	lis  net.Listener
	s    MyTestServer
	conn net.Conn
}

func newServer(lis net.Listener, s MyTestServer) *server {
	return &server{
		lis: lis,
		s:   s,
	}
}

func (this *server) Serve() {
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
	this.s.Downstream(msg, newDownServer(conn))
	this.conn.Close()
}

func (this *server) Stop() {
	this.conn.Close()
}

type downServer struct {
	grpc.ServerStream
	conn protoio.WriteCloser
}

func newDownServer(conn io.Writer) *downServer {
	return &downServer{nil, protoio.NewDelimitedWriter(conn)}
}

func (this *downServer) Send(m *MyMsg) error {
	if err := this.conn.WriteMsg(m); err != nil {
		return err
	}
	return nil
}

type client struct {
	writer protoio.WriteCloser
	reader protoio.ReadCloser
	conn   net.Conn
}

func NewClient(conn net.Conn) *client {
	return &client{
		writer: protoio.NewDelimitedWriter(conn),
		reader: protoio.NewDelimitedReader(conn, 1024*1024),
		conn:   conn,
	}
}

func (this *client) UnaryCall(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	panic("not implemented")
}

func (this *client) Downstream(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (MyTest_DownstreamClient, error) {
	if err := this.writer.WriteMsg(in); err != nil {
		return nil, err
	}
	return &downClient{nil, this.reader}, nil
}

func (this *client) Upstream(ctx context.Context, opts ...grpc.CallOption) (MyTest_UpstreamClient, error) {
	panic("not implemented")
}
func (this *client) Bidi(ctx context.Context, opts ...grpc.CallOption) (MyTest_BidiClient, error) {
	panic("not implemented")
}

type downClient struct {
	grpc.ClientStream
	reader protoio.ReadCloser
}

func (this *downClient) Recv() (*MyMsg, error) {
	m := &MyMsg{}
	err := this.reader.ReadMsg(m)
	return m, err
}

func TestDownstreamTCP(t *testing.T) {
	server, client := setupTCP(t, &aServer{})
	defer server.Stop()
	num := int64(10)
	down, err := client.Downstream(context.Background(), &MyRequest{num})
	if err != nil {
		t.Fatal(err)
	}
	m := &MyMsg{}
	last := -1
	for err == nil {
		m, err = down.Recv()
		if err == nil {
			if int(m.Value) == (last + 1) {
				last++
			} else {
				t.Errorf("out of order last = %d this = %d", last, m.Value)
			}
		}
	}
	if last != int(num)-1 {
		t.Fatalf("wrong last %d expected %d", last, num-1)
	}
}

func BenchmarkDownstreamTCP(b *testing.B) {
	server, client := setupTCP(b, &bServer{})
	defer server.Stop()
	down, err := client.Downstream(context.Background(), &MyRequest{1})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = down.Recv()
	}
}

func BenchmarkDownstreamGRPC(b *testing.B) {
	server, client := setup(b, &bServer{})
	defer server.Stop()
	down, err := client.Downstream(context.Background(), &MyRequest{1})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = down.Recv()
	}
}
