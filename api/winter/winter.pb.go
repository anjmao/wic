// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/winter/winter.proto

package winter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PlayRequest struct {
	Start                *Start   `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	ShootAt              *ShootAt `protobuf:"bytes,2,opt,name=shootAt,proto3" json:"shootAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayRequest) Reset()         { *m = PlayRequest{} }
func (m *PlayRequest) String() string { return proto.CompactTextString(m) }
func (*PlayRequest) ProtoMessage()    {}
func (*PlayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_winter_b75287c3f643ddb2, []int{0}
}
func (m *PlayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayRequest.Unmarshal(m, b)
}
func (m *PlayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayRequest.Marshal(b, m, deterministic)
}
func (dst *PlayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayRequest.Merge(dst, src)
}
func (m *PlayRequest) XXX_Size() int {
	return xxx_messageInfo_PlayRequest.Size(m)
}
func (m *PlayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlayRequest proto.InternalMessageInfo

func (m *PlayRequest) GetStart() *Start {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *PlayRequest) GetShootAt() *ShootAt {
	if m != nil {
		return m.ShootAt
	}
	return nil
}

type PlayReply struct {
	Ready                *Ready       `protobuf:"bytes,1,opt,name=ready,proto3" json:"ready,omitempty"`
	Zombie               *Zombie      `protobuf:"bytes,2,opt,name=zombie,proto3" json:"zombie,omitempty"`
	ShootResult          *ShootResult `protobuf:"bytes,3,opt,name=shootResult,proto3" json:"shootResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PlayReply) Reset()         { *m = PlayReply{} }
func (m *PlayReply) String() string { return proto.CompactTextString(m) }
func (*PlayReply) ProtoMessage()    {}
func (*PlayReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_winter_b75287c3f643ddb2, []int{1}
}
func (m *PlayReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayReply.Unmarshal(m, b)
}
func (m *PlayReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayReply.Marshal(b, m, deterministic)
}
func (dst *PlayReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayReply.Merge(dst, src)
}
func (m *PlayReply) XXX_Size() int {
	return xxx_messageInfo_PlayReply.Size(m)
}
func (m *PlayReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayReply.DiscardUnknown(m)
}

var xxx_messageInfo_PlayReply proto.InternalMessageInfo

func (m *PlayReply) GetReady() *Ready {
	if m != nil {
		return m.Ready
	}
	return nil
}

func (m *PlayReply) GetZombie() *Zombie {
	if m != nil {
		return m.Zombie
	}
	return nil
}

func (m *PlayReply) GetShootResult() *ShootResult {
	if m != nil {
		return m.ShootResult
	}
	return nil
}

type Start struct {
	PlayerName           string   `protobuf:"bytes,1,opt,name=playerName,proto3" json:"playerName,omitempty"`
	GameID               string   `protobuf:"bytes,2,opt,name=gameID,proto3" json:"gameID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Start) Reset()         { *m = Start{} }
func (m *Start) String() string { return proto.CompactTextString(m) }
func (*Start) ProtoMessage()    {}
func (*Start) Descriptor() ([]byte, []int) {
	return fileDescriptor_winter_b75287c3f643ddb2, []int{2}
}
func (m *Start) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Start.Unmarshal(m, b)
}
func (m *Start) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Start.Marshal(b, m, deterministic)
}
func (dst *Start) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Start.Merge(dst, src)
}
func (m *Start) XXX_Size() int {
	return xxx_messageInfo_Start.Size(m)
}
func (m *Start) XXX_DiscardUnknown() {
	xxx_messageInfo_Start.DiscardUnknown(m)
}

var xxx_messageInfo_Start proto.InternalMessageInfo

func (m *Start) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *Start) GetGameID() string {
	if m != nil {
		return m.GameID
	}
	return ""
}

type ShootAt struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShootAt) Reset()         { *m = ShootAt{} }
func (m *ShootAt) String() string { return proto.CompactTextString(m) }
func (*ShootAt) ProtoMessage()    {}
func (*ShootAt) Descriptor() ([]byte, []int) {
	return fileDescriptor_winter_b75287c3f643ddb2, []int{3}
}
func (m *ShootAt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShootAt.Unmarshal(m, b)
}
func (m *ShootAt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShootAt.Marshal(b, m, deterministic)
}
func (dst *ShootAt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShootAt.Merge(dst, src)
}
func (m *ShootAt) XXX_Size() int {
	return xxx_messageInfo_ShootAt.Size(m)
}
func (m *ShootAt) XXX_DiscardUnknown() {
	xxx_messageInfo_ShootAt.DiscardUnknown(m)
}

var xxx_messageInfo_ShootAt proto.InternalMessageInfo

func (m *ShootAt) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ShootAt) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type ShootResult struct {
	PlayerName           string   `protobuf:"bytes,1,opt,name=playerName,proto3" json:"playerName,omitempty"`
	ZombieName           string   `protobuf:"bytes,2,opt,name=zombieName,proto3" json:"zombieName,omitempty"`
	Points               int32    `protobuf:"varint,3,opt,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShootResult) Reset()         { *m = ShootResult{} }
func (m *ShootResult) String() string { return proto.CompactTextString(m) }
func (*ShootResult) ProtoMessage()    {}
func (*ShootResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_winter_b75287c3f643ddb2, []int{4}
}
func (m *ShootResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShootResult.Unmarshal(m, b)
}
func (m *ShootResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShootResult.Marshal(b, m, deterministic)
}
func (dst *ShootResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShootResult.Merge(dst, src)
}
func (m *ShootResult) XXX_Size() int {
	return xxx_messageInfo_ShootResult.Size(m)
}
func (m *ShootResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ShootResult.DiscardUnknown(m)
}

var xxx_messageInfo_ShootResult proto.InternalMessageInfo

func (m *ShootResult) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *ShootResult) GetZombieName() string {
	if m != nil {
		return m.ZombieName
	}
	return ""
}

func (m *ShootResult) GetPoints() int32 {
	if m != nil {
		return m.Points
	}
	return 0
}

type Ready struct {
	GameID               string   `protobuf:"bytes,1,opt,name=gameID,proto3" json:"gameID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ready) Reset()         { *m = Ready{} }
func (m *Ready) String() string { return proto.CompactTextString(m) }
func (*Ready) ProtoMessage()    {}
func (*Ready) Descriptor() ([]byte, []int) {
	return fileDescriptor_winter_b75287c3f643ddb2, []int{5}
}
func (m *Ready) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ready.Unmarshal(m, b)
}
func (m *Ready) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ready.Marshal(b, m, deterministic)
}
func (dst *Ready) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ready.Merge(dst, src)
}
func (m *Ready) XXX_Size() int {
	return xxx_messageInfo_Ready.Size(m)
}
func (m *Ready) XXX_DiscardUnknown() {
	xxx_messageInfo_Ready.DiscardUnknown(m)
}

var xxx_messageInfo_Ready proto.InternalMessageInfo

func (m *Ready) GetGameID() string {
	if m != nil {
		return m.GameID
	}
	return ""
}

type Zombie struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ReachedWall          bool     `protobuf:"varint,4,opt,name=reachedWall,proto3" json:"reachedWall,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Zombie) Reset()         { *m = Zombie{} }
func (m *Zombie) String() string { return proto.CompactTextString(m) }
func (*Zombie) ProtoMessage()    {}
func (*Zombie) Descriptor() ([]byte, []int) {
	return fileDescriptor_winter_b75287c3f643ddb2, []int{6}
}
func (m *Zombie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Zombie.Unmarshal(m, b)
}
func (m *Zombie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Zombie.Marshal(b, m, deterministic)
}
func (dst *Zombie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Zombie.Merge(dst, src)
}
func (m *Zombie) XXX_Size() int {
	return xxx_messageInfo_Zombie.Size(m)
}
func (m *Zombie) XXX_DiscardUnknown() {
	xxx_messageInfo_Zombie.DiscardUnknown(m)
}

var xxx_messageInfo_Zombie proto.InternalMessageInfo

func (m *Zombie) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Zombie) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Zombie) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Zombie) GetReachedWall() bool {
	if m != nil {
		return m.ReachedWall
	}
	return false
}

type WonPlayer struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WonPlayer) Reset()         { *m = WonPlayer{} }
func (m *WonPlayer) String() string { return proto.CompactTextString(m) }
func (*WonPlayer) ProtoMessage()    {}
func (*WonPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_winter_b75287c3f643ddb2, []int{7}
}
func (m *WonPlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WonPlayer.Unmarshal(m, b)
}
func (m *WonPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WonPlayer.Marshal(b, m, deterministic)
}
func (dst *WonPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WonPlayer.Merge(dst, src)
}
func (m *WonPlayer) XXX_Size() int {
	return xxx_messageInfo_WonPlayer.Size(m)
}
func (m *WonPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_WonPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_WonPlayer proto.InternalMessageInfo

func (m *WonPlayer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListGamesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGamesRequest) Reset()         { *m = ListGamesRequest{} }
func (m *ListGamesRequest) String() string { return proto.CompactTextString(m) }
func (*ListGamesRequest) ProtoMessage()    {}
func (*ListGamesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_winter_b75287c3f643ddb2, []int{8}
}
func (m *ListGamesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGamesRequest.Unmarshal(m, b)
}
func (m *ListGamesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGamesRequest.Marshal(b, m, deterministic)
}
func (dst *ListGamesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGamesRequest.Merge(dst, src)
}
func (m *ListGamesRequest) XXX_Size() int {
	return xxx_messageInfo_ListGamesRequest.Size(m)
}
func (m *ListGamesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGamesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGamesRequest proto.InternalMessageInfo

type ListGamesReply struct {
	Games                []string `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGamesReply) Reset()         { *m = ListGamesReply{} }
func (m *ListGamesReply) String() string { return proto.CompactTextString(m) }
func (*ListGamesReply) ProtoMessage()    {}
func (*ListGamesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_winter_b75287c3f643ddb2, []int{9}
}
func (m *ListGamesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGamesReply.Unmarshal(m, b)
}
func (m *ListGamesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGamesReply.Marshal(b, m, deterministic)
}
func (dst *ListGamesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGamesReply.Merge(dst, src)
}
func (m *ListGamesReply) XXX_Size() int {
	return xxx_messageInfo_ListGamesReply.Size(m)
}
func (m *ListGamesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGamesReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListGamesReply proto.InternalMessageInfo

func (m *ListGamesReply) GetGames() []string {
	if m != nil {
		return m.Games
	}
	return nil
}

func init() {
	proto.RegisterType((*PlayRequest)(nil), "winter.PlayRequest")
	proto.RegisterType((*PlayReply)(nil), "winter.PlayReply")
	proto.RegisterType((*Start)(nil), "winter.Start")
	proto.RegisterType((*ShootAt)(nil), "winter.ShootAt")
	proto.RegisterType((*ShootResult)(nil), "winter.ShootResult")
	proto.RegisterType((*Ready)(nil), "winter.Ready")
	proto.RegisterType((*Zombie)(nil), "winter.Zombie")
	proto.RegisterType((*WonPlayer)(nil), "winter.WonPlayer")
	proto.RegisterType((*ListGamesRequest)(nil), "winter.ListGamesRequest")
	proto.RegisterType((*ListGamesReply)(nil), "winter.ListGamesReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WinterGameClient is the client API for WinterGame service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WinterGameClient interface {
	Play(ctx context.Context, opts ...grpc.CallOption) (WinterGame_PlayClient, error)
	ListGames(ctx context.Context, in *ListGamesRequest, opts ...grpc.CallOption) (*ListGamesReply, error)
}

type winterGameClient struct {
	cc *grpc.ClientConn
}

func NewWinterGameClient(cc *grpc.ClientConn) WinterGameClient {
	return &winterGameClient{cc}
}

func (c *winterGameClient) Play(ctx context.Context, opts ...grpc.CallOption) (WinterGame_PlayClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WinterGame_serviceDesc.Streams[0], "/winter.WinterGame/Play", opts...)
	if err != nil {
		return nil, err
	}
	x := &winterGamePlayClient{stream}
	return x, nil
}

type WinterGame_PlayClient interface {
	Send(*PlayRequest) error
	Recv() (*PlayReply, error)
	grpc.ClientStream
}

type winterGamePlayClient struct {
	grpc.ClientStream
}

func (x *winterGamePlayClient) Send(m *PlayRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *winterGamePlayClient) Recv() (*PlayReply, error) {
	m := new(PlayReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *winterGameClient) ListGames(ctx context.Context, in *ListGamesRequest, opts ...grpc.CallOption) (*ListGamesReply, error) {
	out := new(ListGamesReply)
	err := c.cc.Invoke(ctx, "/winter.WinterGame/ListGames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WinterGameServer is the server API for WinterGame service.
type WinterGameServer interface {
	Play(WinterGame_PlayServer) error
	ListGames(context.Context, *ListGamesRequest) (*ListGamesReply, error)
}

func RegisterWinterGameServer(s *grpc.Server, srv WinterGameServer) {
	s.RegisterService(&_WinterGame_serviceDesc, srv)
}

func _WinterGame_Play_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WinterGameServer).Play(&winterGamePlayServer{stream})
}

type WinterGame_PlayServer interface {
	Send(*PlayReply) error
	Recv() (*PlayRequest, error)
	grpc.ServerStream
}

type winterGamePlayServer struct {
	grpc.ServerStream
}

func (x *winterGamePlayServer) Send(m *PlayReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *winterGamePlayServer) Recv() (*PlayRequest, error) {
	m := new(PlayRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WinterGame_ListGames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WinterGameServer).ListGames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/winter.WinterGame/ListGames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WinterGameServer).ListGames(ctx, req.(*ListGamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WinterGame_serviceDesc = grpc.ServiceDesc{
	ServiceName: "winter.WinterGame",
	HandlerType: (*WinterGameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGames",
			Handler:    _WinterGame_ListGames_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Play",
			Handler:       _WinterGame_Play_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/winter/winter.proto",
}

func init() { proto.RegisterFile("api/winter/winter.proto", fileDescriptor_winter_b75287c3f643ddb2) }

var fileDescriptor_winter_b75287c3f643ddb2 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x8b, 0x13, 0x31,
	0x14, 0x35, 0xb6, 0x33, 0xeb, 0xdc, 0xd1, 0x55, 0xaf, 0xb2, 0x0e, 0xfb, 0xb0, 0x3b, 0x44, 0x5c,
	0xc6, 0x97, 0x55, 0x56, 0x7d, 0x5e, 0x04, 0x41, 0x04, 0x91, 0x12, 0x1f, 0x0a, 0x05, 0x1f, 0x52,
	0x1b, 0xec, 0xc0, 0x7c, 0x39, 0x49, 0xb1, 0xe3, 0xab, 0x3f, 0xc0, 0xbf, 0x2c, 0xb9, 0x49, 0xc7,
	0xb4, 0x88, 0xfb, 0xd4, 0xde, 0x73, 0x4e, 0xce, 0x3d, 0x37, 0xb9, 0x03, 0x4f, 0x64, 0x57, 0xbe,
	0xf8, 0x51, 0x36, 0x46, 0xf5, 0xfe, 0xe7, 0xb2, 0xeb, 0x5b, 0xd3, 0x62, 0xec, 0x2a, 0xfe, 0x05,
	0xd2, 0x59, 0x25, 0x07, 0xa1, 0xbe, 0x6f, 0x94, 0x36, 0xf8, 0x14, 0x22, 0x6d, 0x64, 0x6f, 0x32,
	0x96, 0xb3, 0x22, 0xbd, 0xba, 0x77, 0xe9, 0x0f, 0x7d, 0xb6, 0xa0, 0x70, 0x1c, 0x3e, 0x87, 0x23,
	0xbd, 0x6e, 0x5b, 0xf3, 0xd6, 0x64, 0xb7, 0x49, 0x76, 0x7f, 0x94, 0x39, 0x58, 0xec, 0x78, 0xfe,
	0x9b, 0x41, 0xe2, 0xfc, 0xbb, 0x6a, 0xb0, 0xee, 0xbd, 0x92, 0xab, 0xe1, 0xd0, 0x5d, 0x58, 0x50,
	0x38, 0x0e, 0x2f, 0x20, 0xfe, 0xd9, 0xd6, 0xcb, 0x52, 0x79, 0xf3, 0xe3, 0x9d, 0x6a, 0x41, 0xa8,
	0xf0, 0x2c, 0xbe, 0x81, 0x94, 0xba, 0x08, 0xa5, 0x37, 0x95, 0xc9, 0x26, 0x24, 0x7e, 0xb4, 0x97,
	0xc4, 0x51, 0x22, 0xd4, 0xf1, 0x6b, 0x88, 0x68, 0x18, 0x3c, 0x03, 0xe8, 0x2a, 0x39, 0xa8, 0xfe,
	0x93, 0xac, 0x15, 0x25, 0x4a, 0x44, 0x80, 0xe0, 0x09, 0xc4, 0xdf, 0x64, 0xad, 0x3e, 0xbc, 0xa3,
	0x1c, 0x89, 0xf0, 0x15, 0x7f, 0x06, 0x47, 0x7e, 0x4c, 0xbc, 0x0b, 0x6c, 0x4b, 0x27, 0x23, 0xc1,
	0xb6, 0xb6, 0x1a, 0x48, 0x1b, 0x09, 0x36, 0x70, 0x05, 0x69, 0x90, 0xe1, 0xc6, 0x6e, 0x67, 0x00,
	0x6e, 0x2e, 0xe2, 0x5d, 0xc7, 0x00, 0xb1, 0x69, 0xba, 0xb6, 0x6c, 0x8c, 0xa6, 0x41, 0x23, 0xe1,
	0x2b, 0x7e, 0x0e, 0x11, 0xdd, 0x5e, 0x10, 0x97, 0xed, 0xc5, 0x5d, 0x40, 0xec, 0x2e, 0xee, 0x7f,
	0x69, 0x11, 0x61, 0xda, 0xd8, 0xc6, 0x13, 0x3a, 0x4b, 0xff, 0x31, 0x87, 0xb4, 0x57, 0xf2, 0xeb,
	0x5a, 0xad, 0xe6, 0xb2, 0xaa, 0xb2, 0x69, 0xce, 0x8a, 0x3b, 0x22, 0x84, 0xf8, 0x39, 0x24, 0xf3,
	0xb6, 0x99, 0xd1, 0x14, 0xa3, 0x05, 0xfb, 0x6b, 0xc1, 0x11, 0x1e, 0x7c, 0x2c, 0xb5, 0x79, 0x2f,
	0x6b, 0xa5, 0xfd, 0x8a, 0xf1, 0x0b, 0x38, 0x0e, 0x30, 0xbb, 0x16, 0x8f, 0x21, 0xb2, 0x61, 0x75,
	0xc6, 0xf2, 0x49, 0x91, 0x08, 0x57, 0x5c, 0xfd, 0x62, 0x00, 0x73, 0x7a, 0x4c, 0x2b, 0xc5, 0xd7,
	0x30, 0xb5, 0x8d, 0x70, 0x7c, 0xe1, 0x60, 0x6d, 0x4f, 0x1f, 0xee, 0x83, 0x5d, 0x35, 0xf0, 0x5b,
	0x05, 0x7b, 0xc9, 0xf0, 0x1a, 0x92, 0xb1, 0x19, 0x66, 0x3b, 0xd5, 0x61, 0xa6, 0xd3, 0x93, 0x7f,
	0x30, 0x64, 0xb2, 0x8c, 0xe9, 0x73, 0x79, 0xf5, 0x27, 0x00, 0x00, 0xff, 0xff, 0x62, 0x7d, 0x83,
	0xd3, 0x49, 0x03, 0x00, 0x00,
}