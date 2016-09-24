// Code generated by protoc-gen-gogo.
// source: combos/unsafemarshaler/bench.proto
// DO NOT EDIT!

/*
Package bench is a generated protocol buffer package.

It is generated from these files:
	combos/unsafemarshaler/bench.proto

It has these top-level messages:
	Request
	Small
	Medium
	Big
*/
package bench

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import unsafe "unsafe"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Num int64 `protobuf:"varint,1,opt,name=Num,json=num,proto3" json:"Num,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorBench, []int{0} }

type Small struct {
	Field3  int32  `protobuf:"varint,3,opt,name=Field3,json=field3,proto3" json:"Field3,omitempty"`
	Field11 uint64 `protobuf:"fixed64,11,opt,name=Field11,json=field11,proto3" json:"Field11,omitempty"`
	Field14 string `protobuf:"bytes,14,opt,name=Field14,json=field14,proto3" json:"Field14,omitempty"`
}

func (m *Small) Reset()                    { *m = Small{} }
func (m *Small) String() string            { return proto.CompactTextString(m) }
func (*Small) ProtoMessage()               {}
func (*Small) Descriptor() ([]byte, []int) { return fileDescriptorBench, []int{1} }

type Medium struct {
	Field1  float64 `protobuf:"fixed64,1,opt,name=Field1,json=field1,proto3" json:"Field1,omitempty"`
	Field2  float32 `protobuf:"fixed32,2,opt,name=Field2,json=field2,proto3" json:"Field2,omitempty"`
	Field3  int32   `protobuf:"varint,3,opt,name=Field3,json=field3,proto3" json:"Field3,omitempty"`
	Field4  int64   `protobuf:"varint,4,opt,name=Field4,json=field4,proto3" json:"Field4,omitempty"`
	Field5  uint32  `protobuf:"varint,5,opt,name=Field5,json=field5,proto3" json:"Field5,omitempty"`
	Field6  uint64  `protobuf:"varint,6,opt,name=Field6,json=field6,proto3" json:"Field6,omitempty"`
	Field7  int32   `protobuf:"zigzag32,7,opt,name=Field7,json=field7,proto3" json:"Field7,omitempty"`
	Field8  int64   `protobuf:"zigzag64,8,opt,name=Field8,json=field8,proto3" json:"Field8,omitempty"`
	Field9  uint32  `protobuf:"fixed32,9,opt,name=Field9,json=field9,proto3" json:"Field9,omitempty"`
	Field10 int32   `protobuf:"fixed32,10,opt,name=Field10,json=field10,proto3" json:"Field10,omitempty"`
	Field11 uint64  `protobuf:"fixed64,11,opt,name=Field11,json=field11,proto3" json:"Field11,omitempty"`
	Field12 int64   `protobuf:"fixed64,12,opt,name=Field12,json=field12,proto3" json:"Field12,omitempty"`
	Field13 bool    `protobuf:"varint,13,opt,name=Field13,json=field13,proto3" json:"Field13,omitempty"`
	Field14 string  `protobuf:"bytes,14,opt,name=Field14,json=field14,proto3" json:"Field14,omitempty"`
	Field15 []byte  `protobuf:"bytes,15,opt,name=Field15,json=field15,proto3" json:"Field15,omitempty"`
}

func (m *Medium) Reset()                    { *m = Medium{} }
func (m *Medium) String() string            { return proto.CompactTextString(m) }
func (*Medium) ProtoMessage()               {}
func (*Medium) Descriptor() ([]byte, []int) { return fileDescriptorBench, []int{2} }

type Big struct {
	Field1  float64  `protobuf:"fixed64,1,opt,name=Field1,json=field1,proto3" json:"Field1,omitempty"`
	Field2  float32  `protobuf:"fixed32,2,opt,name=Field2,json=field2,proto3" json:"Field2,omitempty"`
	Field3  *Medium  `protobuf:"bytes,3,opt,name=Field3,json=field3" json:"Field3,omitempty"`
	Field4  []*Small `protobuf:"bytes,4,rep,name=Field4,json=field4" json:"Field4,omitempty"`
	Field6  uint64   `protobuf:"varint,6,opt,name=Field6,json=field6,proto3" json:"Field6,omitempty"`
	Field7  int32    `protobuf:"zigzag32,7,opt,name=Field7,json=field7,proto3" json:"Field7,omitempty"`
	Field8  *Medium  `protobuf:"bytes,8,opt,name=Field8,json=field8" json:"Field8,omitempty"`
	Field13 bool     `protobuf:"varint,13,opt,name=Field13,json=field13,proto3" json:"Field13,omitempty"`
	Field14 string   `protobuf:"bytes,14,opt,name=Field14,json=field14,proto3" json:"Field14,omitempty"`
	Field15 []byte   `protobuf:"bytes,15,opt,name=Field15,json=field15,proto3" json:"Field15,omitempty"`
}

func (m *Big) Reset()                    { *m = Big{} }
func (m *Big) String() string            { return proto.CompactTextString(m) }
func (*Big) ProtoMessage()               {}
func (*Big) Descriptor() ([]byte, []int) { return fileDescriptorBench, []int{3} }

func (m *Big) GetField3() *Medium {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *Big) GetField4() []*Small {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *Big) GetField8() *Medium {
	if m != nil {
		return m.Field8
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "bench.Request")
	proto.RegisterType((*Small)(nil), "bench.Small")
	proto.RegisterType((*Medium)(nil), "bench.Medium")
	proto.RegisterType((*Big)(nil), "bench.Big")
}
func (this *Request) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Request)
	if !ok {
		that2, ok := that.(Request)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Request")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Request but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Request but is not nil && this == nil")
	}
	if this.Num != that1.Num {
		return fmt.Errorf("Num this(%v) Not Equal that(%v)", this.Num, that1.Num)
	}
	return nil
}
func (this *Request) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Request)
	if !ok {
		that2, ok := that.(Request)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Num != that1.Num {
		return false
	}
	return true
}
func (this *Small) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Small)
	if !ok {
		that2, ok := that.(Small)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Small")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Small but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Small but is not nil && this == nil")
	}
	if this.Field3 != that1.Field3 {
		return fmt.Errorf("Field3 this(%v) Not Equal that(%v)", this.Field3, that1.Field3)
	}
	if this.Field11 != that1.Field11 {
		return fmt.Errorf("Field11 this(%v) Not Equal that(%v)", this.Field11, that1.Field11)
	}
	if this.Field14 != that1.Field14 {
		return fmt.Errorf("Field14 this(%v) Not Equal that(%v)", this.Field14, that1.Field14)
	}
	return nil
}
func (this *Small) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Small)
	if !ok {
		that2, ok := that.(Small)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Field3 != that1.Field3 {
		return false
	}
	if this.Field11 != that1.Field11 {
		return false
	}
	if this.Field14 != that1.Field14 {
		return false
	}
	return true
}
func (this *Medium) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Medium)
	if !ok {
		that2, ok := that.(Medium)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Medium")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Medium but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Medium but is not nil && this == nil")
	}
	if this.Field1 != that1.Field1 {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", this.Field1, that1.Field1)
	}
	if this.Field2 != that1.Field2 {
		return fmt.Errorf("Field2 this(%v) Not Equal that(%v)", this.Field2, that1.Field2)
	}
	if this.Field3 != that1.Field3 {
		return fmt.Errorf("Field3 this(%v) Not Equal that(%v)", this.Field3, that1.Field3)
	}
	if this.Field4 != that1.Field4 {
		return fmt.Errorf("Field4 this(%v) Not Equal that(%v)", this.Field4, that1.Field4)
	}
	if this.Field5 != that1.Field5 {
		return fmt.Errorf("Field5 this(%v) Not Equal that(%v)", this.Field5, that1.Field5)
	}
	if this.Field6 != that1.Field6 {
		return fmt.Errorf("Field6 this(%v) Not Equal that(%v)", this.Field6, that1.Field6)
	}
	if this.Field7 != that1.Field7 {
		return fmt.Errorf("Field7 this(%v) Not Equal that(%v)", this.Field7, that1.Field7)
	}
	if this.Field8 != that1.Field8 {
		return fmt.Errorf("Field8 this(%v) Not Equal that(%v)", this.Field8, that1.Field8)
	}
	if this.Field9 != that1.Field9 {
		return fmt.Errorf("Field9 this(%v) Not Equal that(%v)", this.Field9, that1.Field9)
	}
	if this.Field10 != that1.Field10 {
		return fmt.Errorf("Field10 this(%v) Not Equal that(%v)", this.Field10, that1.Field10)
	}
	if this.Field11 != that1.Field11 {
		return fmt.Errorf("Field11 this(%v) Not Equal that(%v)", this.Field11, that1.Field11)
	}
	if this.Field12 != that1.Field12 {
		return fmt.Errorf("Field12 this(%v) Not Equal that(%v)", this.Field12, that1.Field12)
	}
	if this.Field13 != that1.Field13 {
		return fmt.Errorf("Field13 this(%v) Not Equal that(%v)", this.Field13, that1.Field13)
	}
	if this.Field14 != that1.Field14 {
		return fmt.Errorf("Field14 this(%v) Not Equal that(%v)", this.Field14, that1.Field14)
	}
	if !bytes.Equal(this.Field15, that1.Field15) {
		return fmt.Errorf("Field15 this(%v) Not Equal that(%v)", this.Field15, that1.Field15)
	}
	return nil
}
func (this *Medium) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Medium)
	if !ok {
		that2, ok := that.(Medium)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Field1 != that1.Field1 {
		return false
	}
	if this.Field2 != that1.Field2 {
		return false
	}
	if this.Field3 != that1.Field3 {
		return false
	}
	if this.Field4 != that1.Field4 {
		return false
	}
	if this.Field5 != that1.Field5 {
		return false
	}
	if this.Field6 != that1.Field6 {
		return false
	}
	if this.Field7 != that1.Field7 {
		return false
	}
	if this.Field8 != that1.Field8 {
		return false
	}
	if this.Field9 != that1.Field9 {
		return false
	}
	if this.Field10 != that1.Field10 {
		return false
	}
	if this.Field11 != that1.Field11 {
		return false
	}
	if this.Field12 != that1.Field12 {
		return false
	}
	if this.Field13 != that1.Field13 {
		return false
	}
	if this.Field14 != that1.Field14 {
		return false
	}
	if !bytes.Equal(this.Field15, that1.Field15) {
		return false
	}
	return true
}
func (this *Big) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Big)
	if !ok {
		that2, ok := that.(Big)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Big")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Big but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Big but is not nil && this == nil")
	}
	if this.Field1 != that1.Field1 {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", this.Field1, that1.Field1)
	}
	if this.Field2 != that1.Field2 {
		return fmt.Errorf("Field2 this(%v) Not Equal that(%v)", this.Field2, that1.Field2)
	}
	if !this.Field3.Equal(that1.Field3) {
		return fmt.Errorf("Field3 this(%v) Not Equal that(%v)", this.Field3, that1.Field3)
	}
	if len(this.Field4) != len(that1.Field4) {
		return fmt.Errorf("Field4 this(%v) Not Equal that(%v)", len(this.Field4), len(that1.Field4))
	}
	for i := range this.Field4 {
		if !this.Field4[i].Equal(that1.Field4[i]) {
			return fmt.Errorf("Field4 this[%v](%v) Not Equal that[%v](%v)", i, this.Field4[i], i, that1.Field4[i])
		}
	}
	if this.Field6 != that1.Field6 {
		return fmt.Errorf("Field6 this(%v) Not Equal that(%v)", this.Field6, that1.Field6)
	}
	if this.Field7 != that1.Field7 {
		return fmt.Errorf("Field7 this(%v) Not Equal that(%v)", this.Field7, that1.Field7)
	}
	if !this.Field8.Equal(that1.Field8) {
		return fmt.Errorf("Field8 this(%v) Not Equal that(%v)", this.Field8, that1.Field8)
	}
	if this.Field13 != that1.Field13 {
		return fmt.Errorf("Field13 this(%v) Not Equal that(%v)", this.Field13, that1.Field13)
	}
	if this.Field14 != that1.Field14 {
		return fmt.Errorf("Field14 this(%v) Not Equal that(%v)", this.Field14, that1.Field14)
	}
	if !bytes.Equal(this.Field15, that1.Field15) {
		return fmt.Errorf("Field15 this(%v) Not Equal that(%v)", this.Field15, that1.Field15)
	}
	return nil
}
func (this *Big) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Big)
	if !ok {
		that2, ok := that.(Big)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Field1 != that1.Field1 {
		return false
	}
	if this.Field2 != that1.Field2 {
		return false
	}
	if !this.Field3.Equal(that1.Field3) {
		return false
	}
	if len(this.Field4) != len(that1.Field4) {
		return false
	}
	for i := range this.Field4 {
		if !this.Field4[i].Equal(that1.Field4[i]) {
			return false
		}
	}
	if this.Field6 != that1.Field6 {
		return false
	}
	if this.Field7 != that1.Field7 {
		return false
	}
	if !this.Field8.Equal(that1.Field8) {
		return false
	}
	if this.Field13 != that1.Field13 {
		return false
	}
	if this.Field14 != that1.Field14 {
		return false
	}
	if !bytes.Equal(this.Field15, that1.Field15) {
		return false
	}
	return true
}
func (this *Request) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&bench.Request{")
	s = append(s, "Num: "+fmt.Sprintf("%#v", this.Num)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Small) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&bench.Small{")
	s = append(s, "Field3: "+fmt.Sprintf("%#v", this.Field3)+",\n")
	s = append(s, "Field11: "+fmt.Sprintf("%#v", this.Field11)+",\n")
	s = append(s, "Field14: "+fmt.Sprintf("%#v", this.Field14)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Medium) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 19)
	s = append(s, "&bench.Medium{")
	s = append(s, "Field1: "+fmt.Sprintf("%#v", this.Field1)+",\n")
	s = append(s, "Field2: "+fmt.Sprintf("%#v", this.Field2)+",\n")
	s = append(s, "Field3: "+fmt.Sprintf("%#v", this.Field3)+",\n")
	s = append(s, "Field4: "+fmt.Sprintf("%#v", this.Field4)+",\n")
	s = append(s, "Field5: "+fmt.Sprintf("%#v", this.Field5)+",\n")
	s = append(s, "Field6: "+fmt.Sprintf("%#v", this.Field6)+",\n")
	s = append(s, "Field7: "+fmt.Sprintf("%#v", this.Field7)+",\n")
	s = append(s, "Field8: "+fmt.Sprintf("%#v", this.Field8)+",\n")
	s = append(s, "Field9: "+fmt.Sprintf("%#v", this.Field9)+",\n")
	s = append(s, "Field10: "+fmt.Sprintf("%#v", this.Field10)+",\n")
	s = append(s, "Field11: "+fmt.Sprintf("%#v", this.Field11)+",\n")
	s = append(s, "Field12: "+fmt.Sprintf("%#v", this.Field12)+",\n")
	s = append(s, "Field13: "+fmt.Sprintf("%#v", this.Field13)+",\n")
	s = append(s, "Field14: "+fmt.Sprintf("%#v", this.Field14)+",\n")
	s = append(s, "Field15: "+fmt.Sprintf("%#v", this.Field15)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Big) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 14)
	s = append(s, "&bench.Big{")
	s = append(s, "Field1: "+fmt.Sprintf("%#v", this.Field1)+",\n")
	s = append(s, "Field2: "+fmt.Sprintf("%#v", this.Field2)+",\n")
	if this.Field3 != nil {
		s = append(s, "Field3: "+fmt.Sprintf("%#v", this.Field3)+",\n")
	}
	if this.Field4 != nil {
		s = append(s, "Field4: "+fmt.Sprintf("%#v", this.Field4)+",\n")
	}
	s = append(s, "Field6: "+fmt.Sprintf("%#v", this.Field6)+",\n")
	s = append(s, "Field7: "+fmt.Sprintf("%#v", this.Field7)+",\n")
	if this.Field8 != nil {
		s = append(s, "Field8: "+fmt.Sprintf("%#v", this.Field8)+",\n")
	}
	s = append(s, "Field13: "+fmt.Sprintf("%#v", this.Field13)+",\n")
	s = append(s, "Field14: "+fmt.Sprintf("%#v", this.Field14)+",\n")
	s = append(s, "Field15: "+fmt.Sprintf("%#v", this.Field15)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringBench(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringBench(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Bencher service

type BencherClient interface {
	BigStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Bencher_BigStreamClient, error)
	MediumStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Bencher_MediumStreamClient, error)
	SmallStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Bencher_SmallStreamClient, error)
}

type bencherClient struct {
	cc *grpc.ClientConn
}

func NewBencherClient(cc *grpc.ClientConn) BencherClient {
	return &bencherClient{cc}
}

func (c *bencherClient) BigStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Bencher_BigStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bencher_serviceDesc.Streams[0], c.cc, "/bench.Bencher/BigStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &bencherBigStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bencher_BigStreamClient interface {
	Recv() (*Big, error)
	grpc.ClientStream
}

type bencherBigStreamClient struct {
	grpc.ClientStream
}

func (x *bencherBigStreamClient) Recv() (*Big, error) {
	m := new(Big)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bencherClient) MediumStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Bencher_MediumStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bencher_serviceDesc.Streams[1], c.cc, "/bench.Bencher/MediumStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &bencherMediumStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bencher_MediumStreamClient interface {
	Recv() (*Medium, error)
	grpc.ClientStream
}

type bencherMediumStreamClient struct {
	grpc.ClientStream
}

func (x *bencherMediumStreamClient) Recv() (*Medium, error) {
	m := new(Medium)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bencherClient) SmallStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Bencher_SmallStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bencher_serviceDesc.Streams[2], c.cc, "/bench.Bencher/SmallStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &bencherSmallStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bencher_SmallStreamClient interface {
	Recv() (*Small, error)
	grpc.ClientStream
}

type bencherSmallStreamClient struct {
	grpc.ClientStream
}

func (x *bencherSmallStreamClient) Recv() (*Small, error) {
	m := new(Small)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Bencher service

type BencherServer interface {
	BigStream(*Request, Bencher_BigStreamServer) error
	MediumStream(*Request, Bencher_MediumStreamServer) error
	SmallStream(*Request, Bencher_SmallStreamServer) error
}

func RegisterBencherServer(s *grpc.Server, srv BencherServer) {
	s.RegisterService(&_Bencher_serviceDesc, srv)
}

func _Bencher_BigStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BencherServer).BigStream(m, &bencherBigStreamServer{stream})
}

type Bencher_BigStreamServer interface {
	Send(*Big) error
	grpc.ServerStream
}

type bencherBigStreamServer struct {
	grpc.ServerStream
}

func (x *bencherBigStreamServer) Send(m *Big) error {
	return x.ServerStream.SendMsg(m)
}

func _Bencher_MediumStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BencherServer).MediumStream(m, &bencherMediumStreamServer{stream})
}

type Bencher_MediumStreamServer interface {
	Send(*Medium) error
	grpc.ServerStream
}

type bencherMediumStreamServer struct {
	grpc.ServerStream
}

func (x *bencherMediumStreamServer) Send(m *Medium) error {
	return x.ServerStream.SendMsg(m)
}

func _Bencher_SmallStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BencherServer).SmallStream(m, &bencherSmallStreamServer{stream})
}

type Bencher_SmallStreamServer interface {
	Send(*Small) error
	grpc.ServerStream
}

type bencherSmallStreamServer struct {
	grpc.ServerStream
}

func (x *bencherSmallStreamServer) Send(m *Small) error {
	return x.ServerStream.SendMsg(m)
}

var _Bencher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bench.Bencher",
	HandlerType: (*BencherServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BigStream",
			Handler:       _Bencher_BigStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MediumStream",
			Handler:       _Bencher_MediumStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SmallStream",
			Handler:       _Bencher_SmallStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptorBench,
}

func NewPopulatedRequest(r randyBench, easy bool) *Request {
	this := &Request{}
	this.Num = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Num *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSmall(r randyBench, easy bool) *Small {
	this := &Small{}
	this.Field3 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field3 *= -1
	}
	this.Field11 = uint64(uint64(r.Uint32()))
	this.Field14 = randStringBench(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMedium(r randyBench, easy bool) *Medium {
	this := &Medium{}
	this.Field1 = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Field1 *= -1
	}
	this.Field2 = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.Field2 *= -1
	}
	this.Field3 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field3 *= -1
	}
	this.Field4 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field4 *= -1
	}
	this.Field5 = uint32(r.Uint32())
	this.Field6 = uint64(uint64(r.Uint32()))
	this.Field7 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field7 *= -1
	}
	this.Field8 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field8 *= -1
	}
	this.Field9 = uint32(r.Uint32())
	this.Field10 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field10 *= -1
	}
	this.Field11 = uint64(uint64(r.Uint32()))
	this.Field12 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field12 *= -1
	}
	this.Field13 = bool(bool(r.Intn(2) == 0))
	this.Field14 = randStringBench(r)
	v1 := r.Intn(100)
	this.Field15 = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Field15[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedBig(r randyBench, easy bool) *Big {
	this := &Big{}
	this.Field1 = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Field1 *= -1
	}
	this.Field2 = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.Field2 *= -1
	}
	if r.Intn(10) != 0 {
		this.Field3 = NewPopulatedMedium(r, easy)
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(5)
		this.Field4 = make([]*Small, v2)
		for i := 0; i < v2; i++ {
			this.Field4[i] = NewPopulatedSmall(r, easy)
		}
	}
	this.Field6 = uint64(uint64(r.Uint32()))
	this.Field7 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field7 *= -1
	}
	if r.Intn(10) != 0 {
		this.Field8 = NewPopulatedMedium(r, easy)
	}
	this.Field13 = bool(bool(r.Intn(2) == 0))
	this.Field14 = randStringBench(r)
	v3 := r.Intn(100)
	this.Field15 = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.Field15[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyBench interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneBench(r randyBench) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringBench(r randyBench) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneBench(r)
	}
	return string(tmps)
}
func randUnrecognizedBench(r randyBench, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldBench(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldBench(data []byte, r randyBench, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateBench(data, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		data = encodeVarintPopulateBench(data, uint64(v5))
	case 1:
		data = encodeVarintPopulateBench(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateBench(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateBench(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateBench(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateBench(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Num != 0 {
		n += 1 + sovBench(uint64(m.Num))
	}
	return n
}

func (m *Small) Size() (n int) {
	var l int
	_ = l
	if m.Field3 != 0 {
		n += 1 + sovBench(uint64(m.Field3))
	}
	if m.Field11 != 0 {
		n += 9
	}
	l = len(m.Field14)
	if l > 0 {
		n += 1 + l + sovBench(uint64(l))
	}
	return n
}

func (m *Medium) Size() (n int) {
	var l int
	_ = l
	if m.Field1 != 0 {
		n += 9
	}
	if m.Field2 != 0 {
		n += 5
	}
	if m.Field3 != 0 {
		n += 1 + sovBench(uint64(m.Field3))
	}
	if m.Field4 != 0 {
		n += 1 + sovBench(uint64(m.Field4))
	}
	if m.Field5 != 0 {
		n += 1 + sovBench(uint64(m.Field5))
	}
	if m.Field6 != 0 {
		n += 1 + sovBench(uint64(m.Field6))
	}
	if m.Field7 != 0 {
		n += 1 + sozBench(uint64(m.Field7))
	}
	if m.Field8 != 0 {
		n += 1 + sozBench(uint64(m.Field8))
	}
	if m.Field9 != 0 {
		n += 5
	}
	if m.Field10 != 0 {
		n += 5
	}
	if m.Field11 != 0 {
		n += 9
	}
	if m.Field12 != 0 {
		n += 9
	}
	if m.Field13 {
		n += 2
	}
	l = len(m.Field14)
	if l > 0 {
		n += 1 + l + sovBench(uint64(l))
	}
	l = len(m.Field15)
	if l > 0 {
		n += 1 + l + sovBench(uint64(l))
	}
	return n
}

func (m *Big) Size() (n int) {
	var l int
	_ = l
	if m.Field1 != 0 {
		n += 9
	}
	if m.Field2 != 0 {
		n += 5
	}
	if m.Field3 != nil {
		l = m.Field3.Size()
		n += 1 + l + sovBench(uint64(l))
	}
	if len(m.Field4) > 0 {
		for _, e := range m.Field4 {
			l = e.Size()
			n += 1 + l + sovBench(uint64(l))
		}
	}
	if m.Field6 != 0 {
		n += 1 + sovBench(uint64(m.Field6))
	}
	if m.Field7 != 0 {
		n += 1 + sozBench(uint64(m.Field7))
	}
	if m.Field8 != nil {
		l = m.Field8.Size()
		n += 1 + l + sovBench(uint64(l))
	}
	if m.Field13 {
		n += 2
	}
	l = len(m.Field14)
	if l > 0 {
		n += 1 + l + sovBench(uint64(l))
	}
	l = len(m.Field15)
	if l > 0 {
		n += 1 + l + sovBench(uint64(l))
	}
	return n
}

func sovBench(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBench(x uint64) (n int) {
	return sovBench(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Request) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Num != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintBench(data, i, uint64(m.Num))
	}
	return i, nil
}

func (m *Small) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Small) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Field3 != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintBench(data, i, uint64(m.Field3))
	}
	if m.Field11 != 0 {
		data[i] = 0x59
		i++
		*(*uint64)(unsafe.Pointer(&data[i])) = m.Field11
		i += 8
	}
	if len(m.Field14) > 0 {
		data[i] = 0x72
		i++
		i = encodeVarintBench(data, i, uint64(len(m.Field14)))
		i += copy(data[i:], m.Field14)
	}
	return i, nil
}

func (m *Medium) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Medium) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Field1 != 0 {
		data[i] = 0x9
		i++
		*(*float64)(unsafe.Pointer(&data[i])) = m.Field1
		i += 8
	}
	if m.Field2 != 0 {
		data[i] = 0x15
		i++
		*(*float32)(unsafe.Pointer(&data[i])) = m.Field2
		i += 4
	}
	if m.Field3 != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintBench(data, i, uint64(m.Field3))
	}
	if m.Field4 != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintBench(data, i, uint64(m.Field4))
	}
	if m.Field5 != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintBench(data, i, uint64(m.Field5))
	}
	if m.Field6 != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintBench(data, i, uint64(m.Field6))
	}
	if m.Field7 != 0 {
		data[i] = 0x38
		i++
		i = encodeVarintBench(data, i, uint64((uint32(m.Field7)<<1)^uint32((m.Field7>>31))))
	}
	if m.Field8 != 0 {
		data[i] = 0x40
		i++
		i = encodeVarintBench(data, i, uint64((uint64(m.Field8)<<1)^uint64((m.Field8>>63))))
	}
	if m.Field9 != 0 {
		data[i] = 0x4d
		i++
		*(*uint32)(unsafe.Pointer(&data[i])) = m.Field9
		i += 4
	}
	if m.Field10 != 0 {
		data[i] = 0x55
		i++
		*(*int32)(unsafe.Pointer(&data[i])) = m.Field10
		i += 4
	}
	if m.Field11 != 0 {
		data[i] = 0x59
		i++
		*(*uint64)(unsafe.Pointer(&data[i])) = m.Field11
		i += 8
	}
	if m.Field12 != 0 {
		data[i] = 0x61
		i++
		*(*int64)(unsafe.Pointer(&data[i])) = m.Field12
		i += 8
	}
	if m.Field13 {
		data[i] = 0x68
		i++
		if m.Field13 {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if len(m.Field14) > 0 {
		data[i] = 0x72
		i++
		i = encodeVarintBench(data, i, uint64(len(m.Field14)))
		i += copy(data[i:], m.Field14)
	}
	if len(m.Field15) > 0 {
		data[i] = 0x7a
		i++
		i = encodeVarintBench(data, i, uint64(len(m.Field15)))
		i += copy(data[i:], m.Field15)
	}
	return i, nil
}

func (m *Big) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Big) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Field1 != 0 {
		data[i] = 0x9
		i++
		*(*float64)(unsafe.Pointer(&data[i])) = m.Field1
		i += 8
	}
	if m.Field2 != 0 {
		data[i] = 0x15
		i++
		*(*float32)(unsafe.Pointer(&data[i])) = m.Field2
		i += 4
	}
	if m.Field3 != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintBench(data, i, uint64(m.Field3.Size()))
		n1, err := m.Field3.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Field4) > 0 {
		for _, msg := range m.Field4 {
			data[i] = 0x22
			i++
			i = encodeVarintBench(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Field6 != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintBench(data, i, uint64(m.Field6))
	}
	if m.Field7 != 0 {
		data[i] = 0x38
		i++
		i = encodeVarintBench(data, i, uint64((uint32(m.Field7)<<1)^uint32((m.Field7>>31))))
	}
	if m.Field8 != nil {
		data[i] = 0x42
		i++
		i = encodeVarintBench(data, i, uint64(m.Field8.Size()))
		n2, err := m.Field8.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Field13 {
		data[i] = 0x68
		i++
		if m.Field13 {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if len(m.Field14) > 0 {
		data[i] = 0x72
		i++
		i = encodeVarintBench(data, i, uint64(len(m.Field14)))
		i += copy(data[i:], m.Field14)
	}
	if len(m.Field15) > 0 {
		data[i] = 0x7a
		i++
		i = encodeVarintBench(data, i, uint64(len(m.Field15)))
		i += copy(data[i:], m.Field15)
	}
	return i, nil
}

func encodeFixed64Bench(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Bench(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintBench(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}

func init() { proto.RegisterFile("combos/unsafemarshaler/bench.proto", fileDescriptorBench) }

var fileDescriptorBench = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6e, 0xdb, 0x30,
	0x14, 0x86, 0xfd, 0xac, 0x58, 0x8a, 0x69, 0x3b, 0x49, 0x39, 0x14, 0x44, 0x0a, 0x10, 0x84, 0x90,
	0x00, 0xec, 0x10, 0x4b, 0x96, 0xed, 0xc4, 0x59, 0x35, 0x74, 0x6b, 0x07, 0xfb, 0x04, 0x92, 0x43,
	0xcb, 0x02, 0x2c, 0xab, 0x95, 0xad, 0xbd, 0x87, 0xe8, 0x1d, 0xda, 0xa3, 0x74, 0xec, 0x11, 0x62,
	0xf5, 0x02, 0x1d, 0x33, 0x74, 0x28, 0x4a, 0x51, 0x74, 0x85, 0x42, 0x40, 0x0a, 0x74, 0xd3, 0xfb,
	0x7e, 0x8a, 0xef, 0xf1, 0xff, 0x49, 0x64, 0x2f, 0xd3, 0x24, 0x4c, 0x77, 0x4e, 0xbe, 0xdd, 0x05,
	0x2b, 0x91, 0x04, 0xd9, 0x6e, 0x1d, 0x6c, 0x44, 0xe6, 0x84, 0x62, 0xbb, 0x5c, 0x0f, 0xdf, 0x67,
	0xe9, 0x3e, 0xc5, 0x1d, 0x59, 0x5c, 0xde, 0x44, 0xf1, 0x7e, 0x9d, 0x87, 0xc3, 0x65, 0x9a, 0x38,
	0x51, 0x1a, 0xa5, 0x8e, 0x54, 0xc3, 0x7c, 0x25, 0x2b, 0x59, 0xc8, 0xaf, 0xf2, 0x2f, 0xfb, 0x15,
	0xb2, 0xe6, 0xe2, 0x43, 0x2e, 0x76, 0x7b, 0x7c, 0x81, 0x8c, 0x77, 0x79, 0x42, 0x80, 0x01, 0x37,
	0xe6, 0xc6, 0x36, 0x4f, 0xec, 0x05, 0xea, 0x2c, 0x92, 0x60, 0xb3, 0xc1, 0x2f, 0x91, 0xf9, 0x26,
	0x16, 0x9b, 0x87, 0x31, 0x31, 0x18, 0xf0, 0xce, 0xdc, 0x5c, 0xc9, 0x0a, 0x13, 0x64, 0x49, 0x3e,
	0x1a, 0x91, 0x1e, 0x03, 0x6e, 0xce, 0xad, 0x55, 0x59, 0x1e, 0x95, 0x09, 0x39, 0x63, 0xc0, 0xbb,
	0x95, 0x32, 0xb1, 0x7f, 0xb6, 0x91, 0xf9, 0x56, 0x3c, 0xc4, 0x79, 0xa2, 0xb7, 0x1d, 0xc9, 0xa6,
	0xa0, 0xb6, 0x1d, 0x69, 0xee, 0x91, 0x36, 0x03, 0xde, 0x56, 0xdc, 0x6b, 0x1c, 0xa3, 0xe2, 0x13,
	0x72, 0x22, 0x87, 0x2f, 0xf9, 0x44, 0xf3, 0x29, 0xe9, 0x30, 0xe0, 0x03, 0xc5, 0xa7, 0x9a, 0xdf,
	0x12, 0x93, 0x01, 0x3f, 0x51, 0xfc, 0x56, 0xf3, 0x3b, 0x62, 0x31, 0xe0, 0x2f, 0x14, 0xbf, 0xd3,
	0x7c, 0x46, 0x4e, 0x19, 0x70, 0xac, 0xf8, 0x4c, 0xf3, 0x7b, 0xd2, 0x65, 0xc0, 0x2d, 0xc5, 0xef,
	0x8f, 0x87, 0x77, 0x09, 0x62, 0xc0, 0xcf, 0xab, 0xc3, 0xbb, 0xcf, 0x31, 0xcc, 0x23, 0x7d, 0x06,
	0xfc, 0xa2, 0x52, 0xbc, 0xa3, 0x32, 0x26, 0x03, 0x06, 0xfc, 0xb4, 0x52, 0xc6, 0xcd, 0x26, 0x1f,
	0x95, 0x29, 0x39, 0x67, 0xc0, 0xfb, 0x95, 0x32, 0xb5, 0x3f, 0xb7, 0x91, 0xe1, 0xc7, 0xd1, 0x3f,
	0x7b, 0x7f, 0x5d, 0xf3, 0xbe, 0xe7, 0x0d, 0x86, 0xe5, 0xe5, 0x2b, 0xa3, 0xd4, 0x51, 0x5c, 0xfd,
	0x11, 0x85, 0xc1, 0x7b, 0x5e, 0x5f, 0x2d, 0x93, 0xf7, 0xe8, 0xaf, 0x60, 0x9e, 0x1b, 0xc0, 0x75,
	0x2d, 0x80, 0x86, 0xe6, 0xb3, 0xff, 0xeb, 0x94, 0xf7, 0x09, 0x90, 0xe5, 0xff, 0x6e, 0x23, 0x32,
	0xfc, 0x1a, 0x75, 0xfd, 0x38, 0x5a, 0xec, 0x33, 0x11, 0x24, 0xf8, 0x4c, 0x75, 0x57, 0x0f, 0xe7,
	0x12, 0xa9, 0xda, 0x8f, 0x23, 0x17, 0xb0, 0x83, 0xfa, 0xe5, 0x58, 0x0d, 0xab, 0xeb, 0xb3, 0xbb,
	0x80, 0x6f, 0x50, 0x4f, 0xba, 0xd3, 0xb0, 0xbe, 0xe6, 0xa0, 0x0b, 0xfe, 0xd5, 0xe3, 0x81, 0xc2,
	0x8f, 0x03, 0x85, 0xa7, 0x03, 0x85, 0x2f, 0x05, 0x85, 0xaf, 0x05, 0x6d, 0x7d, 0x2b, 0x68, 0xeb,
	0xb1, 0xa0, 0xf0, 0x54, 0x50, 0xf8, 0xf8, 0x9d, 0xb6, 0x42, 0x53, 0x3e, 0xef, 0xf1, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x29, 0x25, 0x70, 0x63, 0x3a, 0x04, 0x00, 0x00,
}
