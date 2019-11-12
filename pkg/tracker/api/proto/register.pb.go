// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/tracker/api/proto/register.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetMemberByRadiusRequest_Unit int32

const (
	GetMemberByRadiusRequest_M  GetMemberByRadiusRequest_Unit = 0
	GetMemberByRadiusRequest_KM GetMemberByRadiusRequest_Unit = 1
	GetMemberByRadiusRequest_FT GetMemberByRadiusRequest_Unit = 2
	GetMemberByRadiusRequest_MI GetMemberByRadiusRequest_Unit = 3
)

var GetMemberByRadiusRequest_Unit_name = map[int32]string{
	0: "M",
	1: "KM",
	2: "FT",
	3: "MI",
}

var GetMemberByRadiusRequest_Unit_value = map[string]int32{
	"M":  0,
	"KM": 1,
	"FT": 2,
	"MI": 3,
}

func (x GetMemberByRadiusRequest_Unit) String() string {
	return proto.EnumName(GetMemberByRadiusRequest_Unit_name, int32(x))
}

func (GetMemberByRadiusRequest_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_424c431e6e63db39, []int{2, 0}
}

type RegisterRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Latitude             float64  `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_424c431e6e63db39, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterRequest) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *RegisterRequest) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type RegisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_424c431e6e63db39, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

type GetMemberByRadiusRequest struct {
	Longitude            float64                       `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             float64                       `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Radius               float64                       `protobuf:"fixed64,3,opt,name=radius,proto3" json:"radius,omitempty"`
	Unit                 GetMemberByRadiusRequest_Unit `protobuf:"varint,5,opt,name=unit,proto3,enum=proto.GetMemberByRadiusRequest_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GetMemberByRadiusRequest) Reset()         { *m = GetMemberByRadiusRequest{} }
func (m *GetMemberByRadiusRequest) String() string { return proto.CompactTextString(m) }
func (*GetMemberByRadiusRequest) ProtoMessage()    {}
func (*GetMemberByRadiusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_424c431e6e63db39, []int{2}
}

func (m *GetMemberByRadiusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemberByRadiusRequest.Unmarshal(m, b)
}
func (m *GetMemberByRadiusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemberByRadiusRequest.Marshal(b, m, deterministic)
}
func (m *GetMemberByRadiusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberByRadiusRequest.Merge(m, src)
}
func (m *GetMemberByRadiusRequest) XXX_Size() int {
	return xxx_messageInfo_GetMemberByRadiusRequest.Size(m)
}
func (m *GetMemberByRadiusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberByRadiusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberByRadiusRequest proto.InternalMessageInfo

func (m *GetMemberByRadiusRequest) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *GetMemberByRadiusRequest) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *GetMemberByRadiusRequest) GetRadius() float64 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *GetMemberByRadiusRequest) GetUnit() GetMemberByRadiusRequest_Unit {
	if m != nil {
		return m.Unit
	}
	return GetMemberByRadiusRequest_M
}

type GetMemberByRadiusResponse struct {
	Members              []*TrackingMember `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetMemberByRadiusResponse) Reset()         { *m = GetMemberByRadiusResponse{} }
func (m *GetMemberByRadiusResponse) String() string { return proto.CompactTextString(m) }
func (*GetMemberByRadiusResponse) ProtoMessage()    {}
func (*GetMemberByRadiusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_424c431e6e63db39, []int{3}
}

func (m *GetMemberByRadiusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemberByRadiusResponse.Unmarshal(m, b)
}
func (m *GetMemberByRadiusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemberByRadiusResponse.Marshal(b, m, deterministic)
}
func (m *GetMemberByRadiusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberByRadiusResponse.Merge(m, src)
}
func (m *GetMemberByRadiusResponse) XXX_Size() int {
	return xxx_messageInfo_GetMemberByRadiusResponse.Size(m)
}
func (m *GetMemberByRadiusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberByRadiusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberByRadiusResponse proto.InternalMessageInfo

func (m *GetMemberByRadiusResponse) GetMembers() []*TrackingMember {
	if m != nil {
		return m.Members
	}
	return nil
}

type TrackingMember struct {
	Longitude            float64  `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             float64  `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	PeerId               string   `protobuf:"bytes,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackingMember) Reset()         { *m = TrackingMember{} }
func (m *TrackingMember) String() string { return proto.CompactTextString(m) }
func (*TrackingMember) ProtoMessage()    {}
func (*TrackingMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_424c431e6e63db39, []int{4}
}

func (m *TrackingMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackingMember.Unmarshal(m, b)
}
func (m *TrackingMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackingMember.Marshal(b, m, deterministic)
}
func (m *TrackingMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackingMember.Merge(m, src)
}
func (m *TrackingMember) XXX_Size() int {
	return xxx_messageInfo_TrackingMember.Size(m)
}
func (m *TrackingMember) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackingMember.DiscardUnknown(m)
}

var xxx_messageInfo_TrackingMember proto.InternalMessageInfo

func (m *TrackingMember) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *TrackingMember) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *TrackingMember) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

type UpdateRequest struct {
	PeerId               string   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Longitude            float64  `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             float64  `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_424c431e6e63db39, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *UpdateRequest) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *UpdateRequest) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_424c431e6e63db39, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("proto.GetMemberByRadiusRequest_Unit", GetMemberByRadiusRequest_Unit_name, GetMemberByRadiusRequest_Unit_value)
	proto.RegisterType((*RegisterRequest)(nil), "proto.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "proto.RegisterResponse")
	proto.RegisterType((*GetMemberByRadiusRequest)(nil), "proto.GetMemberByRadiusRequest")
	proto.RegisterType((*GetMemberByRadiusResponse)(nil), "proto.GetMemberByRadiusResponse")
	proto.RegisterType((*TrackingMember)(nil), "proto.TrackingMember")
	proto.RegisterType((*UpdateRequest)(nil), "proto.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "proto.UpdateResponse")
}

func init() {
	proto.RegisterFile("pkg/tracker/api/proto/register.proto", fileDescriptor_424c431e6e63db39)
}

var fileDescriptor_424c431e6e63db39 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0xaf, 0x9a, 0x40,
	0x14, 0x75, 0x40, 0x51, 0x6f, 0x53, 0x4a, 0x27, 0x55, 0x29, 0x69, 0x52, 0x32, 0xb1, 0x09, 0x2b,
	0x49, 0xec, 0xa2, 0xdd, 0x74, 0xd3, 0x45, 0x1b, 0xd3, 0xb2, 0x99, 0x68, 0xd2, 0xa4, 0x8b, 0x06,
	0x65, 0x42, 0x26, 0x2a, 0xd0, 0x61, 0x58, 0xbc, 0x1f, 0xf9, 0x7e, 0xc1, 0xfb, 0x33, 0x2f, 0x0c,
	0xa0, 0xa2, 0x4f, 0xdf, 0xe2, 0xad, 0x6e, 0xee, 0xd7, 0x39, 0x67, 0xce, 0x5c, 0x98, 0x66, 0xdb,
	0xd8, 0x97, 0x22, 0xdc, 0x6c, 0x99, 0xf0, 0xc3, 0x8c, 0xfb, 0x99, 0x48, 0x65, 0xea, 0x0b, 0x16,
	0xf3, 0x5c, 0x32, 0x31, 0x53, 0x29, 0xee, 0xa9, 0x40, 0xfe, 0xc2, 0x1b, 0x5a, 0x37, 0x28, 0xfb,
	0x5f, 0xb0, 0x5c, 0x62, 0x13, 0x34, 0x1e, 0xd9, 0xc8, 0x45, 0xde, 0x90, 0x6a, 0x3c, 0xc2, 0x0e,
	0x0c, 0x76, 0xa1, 0xe4, 0xb2, 0x88, 0x98, 0xad, 0xb9, 0xc8, 0x43, 0xf4, 0x90, 0xe3, 0x0f, 0x30,
	0xdc, 0xa5, 0x49, 0x5c, 0x35, 0x75, 0xd5, 0x3c, 0x16, 0x08, 0x06, 0xeb, 0x08, 0x9e, 0x67, 0x69,
	0x92, 0x33, 0x72, 0x8f, 0xc0, 0xfe, 0xc9, 0x64, 0xc0, 0xf6, 0x6b, 0x26, 0xbe, 0xdf, 0xd1, 0x30,
	0xe2, 0x45, 0xde, 0x50, 0xb7, 0xe0, 0xd0, 0x19, 0xdc, 0x4d, 0x21, 0x63, 0x30, 0x84, 0x82, 0xaa,
	0x55, 0xd4, 0x19, 0xfe, 0x0a, 0xdd, 0x22, 0xe1, 0xd2, 0xee, 0xb9, 0xc8, 0x33, 0xe7, 0xd3, 0xea,
	0xf1, 0xb3, 0x6b, 0x02, 0x66, 0xab, 0x84, 0x4b, 0xaa, 0x36, 0xc8, 0x27, 0xe8, 0x96, 0x19, 0xee,
	0x01, 0x0a, 0xac, 0x0e, 0x36, 0x40, 0xfb, 0x15, 0x58, 0xa8, 0x8c, 0x3f, 0x96, 0x96, 0x56, 0xc6,
	0x60, 0x61, 0xe9, 0xe4, 0x37, 0xbc, 0x7f, 0x02, 0xad, 0x7a, 0x2c, 0xf6, 0xa1, 0xbf, 0x57, 0x9d,
	0xdc, 0x46, 0xae, 0xee, 0xbd, 0x9a, 0x8f, 0x6a, 0x01, 0xcb, 0xf2, 0x7b, 0x78, 0x12, 0x57, 0x7b,
	0xb4, 0x99, 0x22, 0x1b, 0x30, 0xdb, 0xad, 0x17, 0x58, 0x32, 0x81, 0x7e, 0xc6, 0x98, 0xf8, 0xc7,
	0x23, 0xe5, 0xc9, 0x90, 0x1a, 0x65, 0xba, 0x88, 0xc8, 0x1a, 0x5e, 0xaf, 0xb2, 0x28, 0x94, 0xac,
	0xb1, 0xfd, 0x64, 0x12, 0x9d, 0x4e, 0xb6, 0xc9, 0xb5, 0x5b, 0xe4, 0x7a, 0x9b, 0x9c, 0x58, 0x60,
	0x36, 0x1c, 0x95, 0x17, 0xf3, 0x07, 0x04, 0xfd, 0x65, 0x75, 0x95, 0xf8, 0x1b, 0x0c, 0x9a, 0xc3,
	0xc0, 0xe3, 0xda, 0x92, 0xb3, 0x33, 0x74, 0x26, 0x17, 0xf5, 0xfa, 0x82, 0x3a, 0xf8, 0x0f, 0xbc,
	0xbd, 0xf0, 0x1c, 0x7f, 0x7c, 0xe6, 0x6f, 0x1d, 0xf7, 0xfa, 0xc0, 0x01, 0xf9, 0x0b, 0x18, 0x95,
	0x6c, 0xfc, 0xae, 0x9e, 0x6e, 0x39, 0xe5, 0x8c, 0xce, 0xaa, 0xcd, 0xe2, 0xda, 0x50, 0xf5, 0xcf,
	0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x79, 0x16, 0xc2, 0x7d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TrackerClient is the client API for Tracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrackerClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	GetMemberByRadius(ctx context.Context, in *GetMemberByRadiusRequest, opts ...grpc.CallOption) (*GetMemberByRadiusResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type trackerClient struct {
	cc *grpc.ClientConn
}

func NewTrackerClient(cc *grpc.ClientConn) TrackerClient {
	return &trackerClient{cc}
}

func (c *trackerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/proto.Tracker/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) GetMemberByRadius(ctx context.Context, in *GetMemberByRadiusRequest, opts ...grpc.CallOption) (*GetMemberByRadiusResponse, error) {
	out := new(GetMemberByRadiusResponse)
	err := c.cc.Invoke(ctx, "/proto.Tracker/GetMemberByRadius", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/proto.Tracker/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackerServer is the server API for Tracker service.
type TrackerServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	GetMemberByRadius(context.Context, *GetMemberByRadiusRequest) (*GetMemberByRadiusResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
}

func RegisterTrackerServer(s *grpc.Server, srv TrackerServer) {
	s.RegisterService(&_Tracker_serviceDesc, srv)
}

func _Tracker_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Tracker/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_GetMemberByRadius_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberByRadiusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).GetMemberByRadius(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Tracker/GetMemberByRadius",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).GetMemberByRadius(ctx, req.(*GetMemberByRadiusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Tracker/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Tracker",
	HandlerType: (*TrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Tracker_Register_Handler,
		},
		{
			MethodName: "GetMemberByRadius",
			Handler:    _Tracker_GetMemberByRadius_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Tracker_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/tracker/api/proto/register.proto",
}