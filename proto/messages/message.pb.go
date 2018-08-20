// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package messages

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

type BackupType int32

const (
	BackupType_LOGICAL   BackupType = 0
	BackupType_HOTBACKUP BackupType = 1
)

var BackupType_name = map[int32]string{
	0: "LOGICAL",
	1: "HOTBACKUP",
}
var BackupType_value = map[string]int32{
	"LOGICAL":   0,
	"HOTBACKUP": 1,
}

func (x BackupType) String() string {
	return proto.EnumName(BackupType_name, int32(x))
}
func (BackupType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{0}
}

type DestinationType int32

const (
	DestinationType_FILE DestinationType = 0
	DestinationType_AWS  DestinationType = 1
)

var DestinationType_name = map[int32]string{
	0: "FILE",
	1: "AWS",
}
var DestinationType_value = map[string]int32{
	"FILE": 0,
	"AWS":  1,
}

func (x DestinationType) String() string {
	return proto.EnumName(DestinationType_name, int32(x))
}
func (DestinationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{1}
}

type CompressionType int32

const (
	CompressionType_NO_COMPRESSION CompressionType = 0
	CompressionType_GZIP           CompressionType = 1
	CompressionType_SNAPPY         CompressionType = 2
	CompressionType_LZ4            CompressionType = 3
)

var CompressionType_name = map[int32]string{
	0: "NO_COMPRESSION",
	1: "GZIP",
	2: "SNAPPY",
	3: "LZ4",
}
var CompressionType_value = map[string]int32{
	"NO_COMPRESSION": 0,
	"GZIP":           1,
	"SNAPPY":         2,
	"LZ4":            3,
}

func (x CompressionType) String() string {
	return proto.EnumName(CompressionType_name, int32(x))
}
func (CompressionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{2}
}

type Cypher int32

const (
	Cypher_NO_CYPHER Cypher = 0
	Cypher_AES       Cypher = 1
	Cypher_DES       Cypher = 2
	Cypher_RC4       Cypher = 3
	Cypher_RSA       Cypher = 4
)

var Cypher_name = map[int32]string{
	0: "NO_CYPHER",
	1: "AES",
	2: "DES",
	3: "RC4",
	4: "RSA",
}
var Cypher_value = map[string]int32{
	"NO_CYPHER": 0,
	"AES":       1,
	"DES":       2,
	"RC4":       3,
	"RSA":       4,
}

func (x Cypher) String() string {
	return proto.EnumName(Cypher_name, int32(x))
}
func (Cypher) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{3}
}

type NodeType int32

const (
	NodeType_UNDEFINED        NodeType = 0
	NodeType_MONGOD           NodeType = 1
	NodeType_MONGOD_REPLSET   NodeType = 2
	NodeType_MONGOD_SHARDSVR  NodeType = 3
	NodeType_MONGOD_CONFIGSVR NodeType = 4
	NodeType_MONGOS           NodeType = 5
)

var NodeType_name = map[int32]string{
	0: "UNDEFINED",
	1: "MONGOD",
	2: "MONGOD_REPLSET",
	3: "MONGOD_SHARDSVR",
	4: "MONGOD_CONFIGSVR",
	5: "MONGOS",
}
var NodeType_value = map[string]int32{
	"UNDEFINED":        0,
	"MONGOD":           1,
	"MONGOD_REPLSET":   2,
	"MONGOD_SHARDSVR":  3,
	"MONGOD_CONFIGSVR": 4,
	"MONGOS":           5,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{4}
}

type ServerMessage_MessageType int32

const (
	ServerMessage_ERROR           ServerMessage_MessageType = 0
	ServerMessage_PING            ServerMessage_MessageType = 1
	ServerMessage_REGISTRATION_OK ServerMessage_MessageType = 2
	ServerMessage_START_BACKUP    ServerMessage_MessageType = 3
	ServerMessage_STOP_BACKUP     ServerMessage_MessageType = 4
	ServerMessage_GET_STATUS      ServerMessage_MessageType = 5
)

var ServerMessage_MessageType_name = map[int32]string{
	0: "ERROR",
	1: "PING",
	2: "REGISTRATION_OK",
	3: "START_BACKUP",
	4: "STOP_BACKUP",
	5: "GET_STATUS",
}
var ServerMessage_MessageType_value = map[string]int32{
	"ERROR":           0,
	"PING":            1,
	"REGISTRATION_OK": 2,
	"START_BACKUP":    3,
	"STOP_BACKUP":     4,
	"GET_STATUS":      5,
}

func (x ServerMessage_MessageType) String() string {
	return proto.EnumName(ServerMessage_MessageType_name, int32(x))
}
func (ServerMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{0, 0}
}

type ClientMessage_MessageType int32

const (
	ClientMessage_ERROR          ClientMessage_MessageType = 0
	ClientMessage_PONG           ClientMessage_MessageType = 1
	ClientMessage_REGISTER       ClientMessage_MessageType = 2
	ClientMessage_BACKUP_STARTED ClientMessage_MessageType = 3
	ClientMessage_BACKUP_FINISH  ClientMessage_MessageType = 4
	ClientMessage_STATUS         ClientMessage_MessageType = 5
)

var ClientMessage_MessageType_name = map[int32]string{
	0: "ERROR",
	1: "PONG",
	2: "REGISTER",
	3: "BACKUP_STARTED",
	4: "BACKUP_FINISH",
	5: "STATUS",
}
var ClientMessage_MessageType_value = map[string]int32{
	"ERROR":          0,
	"PONG":           1,
	"REGISTER":       2,
	"BACKUP_STARTED": 3,
	"BACKUP_FINISH":  4,
	"STATUS":         5,
}

func (x ClientMessage_MessageType) String() string {
	return proto.EnumName(ClientMessage_MessageType_name, int32(x))
}
func (ClientMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{1, 0}
}

type ServerMessage struct {
	Version int32                     `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Type    ServerMessage_MessageType `protobuf:"varint,2,opt,name=type,proto3,enum=messages.ServerMessage_MessageType" json:"type,omitempty"`
	Message []byte                    `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*ServerMessage_StartBackupMsg
	//	*ServerMessage_StopBackupMsg
	Payload              isServerMessage_Payload `protobuf_oneof:"Payload"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ServerMessage) Reset()         { *m = ServerMessage{} }
func (m *ServerMessage) String() string { return proto.CompactTextString(m) }
func (*ServerMessage) ProtoMessage()    {}
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{0}
}
func (m *ServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerMessage.Unmarshal(m, b)
}
func (m *ServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerMessage.Marshal(b, m, deterministic)
}
func (dst *ServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerMessage.Merge(dst, src)
}
func (m *ServerMessage) XXX_Size() int {
	return xxx_messageInfo_ServerMessage.Size(m)
}
func (m *ServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ServerMessage proto.InternalMessageInfo

func (m *ServerMessage) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ServerMessage) GetType() ServerMessage_MessageType {
	if m != nil {
		return m.Type
	}
	return ServerMessage_ERROR
}

func (m *ServerMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type isServerMessage_Payload interface {
	isServerMessage_Payload()
}

type ServerMessage_StartBackupMsg struct {
	StartBackupMsg *StartBackup `protobuf:"bytes,4,opt,name=StartBackupMsg,proto3,oneof"`
}

type ServerMessage_StopBackupMsg struct {
	StopBackupMsg *StopBackup `protobuf:"bytes,5,opt,name=StopBackupMsg,proto3,oneof"`
}

func (*ServerMessage_StartBackupMsg) isServerMessage_Payload() {}

func (*ServerMessage_StopBackupMsg) isServerMessage_Payload() {}

func (m *ServerMessage) GetPayload() isServerMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ServerMessage) GetStartBackupMsg() *StartBackup {
	if x, ok := m.GetPayload().(*ServerMessage_StartBackupMsg); ok {
		return x.StartBackupMsg
	}
	return nil
}

func (m *ServerMessage) GetStopBackupMsg() *StopBackup {
	if x, ok := m.GetPayload().(*ServerMessage_StopBackupMsg); ok {
		return x.StopBackupMsg
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServerMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServerMessage_OneofMarshaler, _ServerMessage_OneofUnmarshaler, _ServerMessage_OneofSizer, []interface{}{
		(*ServerMessage_StartBackupMsg)(nil),
		(*ServerMessage_StopBackupMsg)(nil),
	}
}

func _ServerMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServerMessage)
	// Payload
	switch x := m.Payload.(type) {
	case *ServerMessage_StartBackupMsg:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StartBackupMsg); err != nil {
			return err
		}
	case *ServerMessage_StopBackupMsg:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StopBackupMsg); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ServerMessage.Payload has unexpected type %T", x)
	}
	return nil
}

func _ServerMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServerMessage)
	switch tag {
	case 4: // Payload.StartBackupMsg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StartBackup)
		err := b.DecodeMessage(msg)
		m.Payload = &ServerMessage_StartBackupMsg{msg}
		return true, err
	case 5: // Payload.StopBackupMsg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StopBackup)
		err := b.DecodeMessage(msg)
		m.Payload = &ServerMessage_StopBackupMsg{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ServerMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServerMessage)
	// Payload
	switch x := m.Payload.(type) {
	case *ServerMessage_StartBackupMsg:
		s := proto.Size(x.StartBackupMsg)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerMessage_StopBackupMsg:
		s := proto.Size(x.StopBackupMsg)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ClientMessage struct {
	Version  int32                     `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Type     ClientMessage_MessageType `protobuf:"varint,2,opt,name=type,proto3,enum=messages.ClientMessage_MessageType" json:"type,omitempty"`
	ClientID string                    `protobuf:"bytes,3,opt,name=clientID,proto3" json:"clientID,omitempty"`
	Message  []byte                    `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*ClientMessage_RegisterMsg
	//	*ClientMessage_PingMsg
	//	*ClientMessage_BackupFinishedMsg
	Payload              isClientMessage_Payload `protobuf_oneof:"Payload"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientMessage) Reset()         { *m = ClientMessage{} }
func (m *ClientMessage) String() string { return proto.CompactTextString(m) }
func (*ClientMessage) ProtoMessage()    {}
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{1}
}
func (m *ClientMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientMessage.Unmarshal(m, b)
}
func (m *ClientMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientMessage.Marshal(b, m, deterministic)
}
func (dst *ClientMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientMessage.Merge(dst, src)
}
func (m *ClientMessage) XXX_Size() int {
	return xxx_messageInfo_ClientMessage.Size(m)
}
func (m *ClientMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ClientMessage proto.InternalMessageInfo

func (m *ClientMessage) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ClientMessage) GetType() ClientMessage_MessageType {
	if m != nil {
		return m.Type
	}
	return ClientMessage_ERROR
}

func (m *ClientMessage) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *ClientMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type isClientMessage_Payload interface {
	isClientMessage_Payload()
}

type ClientMessage_RegisterMsg struct {
	RegisterMsg *RegisterPayload `protobuf:"bytes,5,opt,name=RegisterMsg,proto3,oneof"`
}

type ClientMessage_PingMsg struct {
	PingMsg *PongPayload `protobuf:"bytes,6,opt,name=PingMsg,proto3,oneof"`
}

type ClientMessage_BackupFinishedMsg struct {
	BackupFinishedMsg *BackupFinished `protobuf:"bytes,7,opt,name=BackupFinishedMsg,proto3,oneof"`
}

func (*ClientMessage_RegisterMsg) isClientMessage_Payload() {}

func (*ClientMessage_PingMsg) isClientMessage_Payload() {}

func (*ClientMessage_BackupFinishedMsg) isClientMessage_Payload() {}

func (m *ClientMessage) GetPayload() isClientMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ClientMessage) GetRegisterMsg() *RegisterPayload {
	if x, ok := m.GetPayload().(*ClientMessage_RegisterMsg); ok {
		return x.RegisterMsg
	}
	return nil
}

func (m *ClientMessage) GetPingMsg() *PongPayload {
	if x, ok := m.GetPayload().(*ClientMessage_PingMsg); ok {
		return x.PingMsg
	}
	return nil
}

func (m *ClientMessage) GetBackupFinishedMsg() *BackupFinished {
	if x, ok := m.GetPayload().(*ClientMessage_BackupFinishedMsg); ok {
		return x.BackupFinishedMsg
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ClientMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ClientMessage_OneofMarshaler, _ClientMessage_OneofUnmarshaler, _ClientMessage_OneofSizer, []interface{}{
		(*ClientMessage_RegisterMsg)(nil),
		(*ClientMessage_PingMsg)(nil),
		(*ClientMessage_BackupFinishedMsg)(nil),
	}
}

func _ClientMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ClientMessage)
	// Payload
	switch x := m.Payload.(type) {
	case *ClientMessage_RegisterMsg:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RegisterMsg); err != nil {
			return err
		}
	case *ClientMessage_PingMsg:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PingMsg); err != nil {
			return err
		}
	case *ClientMessage_BackupFinishedMsg:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BackupFinishedMsg); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ClientMessage.Payload has unexpected type %T", x)
	}
	return nil
}

func _ClientMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ClientMessage)
	switch tag {
	case 5: // Payload.RegisterMsg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RegisterPayload)
		err := b.DecodeMessage(msg)
		m.Payload = &ClientMessage_RegisterMsg{msg}
		return true, err
	case 6: // Payload.PingMsg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PongPayload)
		err := b.DecodeMessage(msg)
		m.Payload = &ClientMessage_PingMsg{msg}
		return true, err
	case 7: // Payload.BackupFinishedMsg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BackupFinished)
		err := b.DecodeMessage(msg)
		m.Payload = &ClientMessage_BackupFinishedMsg{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ClientMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ClientMessage)
	// Payload
	switch x := m.Payload.(type) {
	case *ClientMessage_RegisterMsg:
		s := proto.Size(x.RegisterMsg)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientMessage_PingMsg:
		s := proto.Size(x.PingMsg)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientMessage_BackupFinishedMsg:
		s := proto.Size(x.BackupFinishedMsg)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type RegisterPayload struct {
	NodeType             NodeType `protobuf:"varint,1,opt,name=NodeType,proto3,enum=messages.NodeType" json:"NodeType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPayload) Reset()         { *m = RegisterPayload{} }
func (m *RegisterPayload) String() string { return proto.CompactTextString(m) }
func (*RegisterPayload) ProtoMessage()    {}
func (*RegisterPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{2}
}
func (m *RegisterPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPayload.Unmarshal(m, b)
}
func (m *RegisterPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPayload.Marshal(b, m, deterministic)
}
func (dst *RegisterPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPayload.Merge(dst, src)
}
func (m *RegisterPayload) XXX_Size() int {
	return xxx_messageInfo_RegisterPayload.Size(m)
}
func (m *RegisterPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPayload.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPayload proto.InternalMessageInfo

func (m *RegisterPayload) GetNodeType() NodeType {
	if m != nil {
		return m.NodeType
	}
	return NodeType_UNDEFINED
}

type PongPayload struct {
	Timestamp            float32  `protobuf:"fixed32,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongPayload) Reset()         { *m = PongPayload{} }
func (m *PongPayload) String() string { return proto.CompactTextString(m) }
func (*PongPayload) ProtoMessage()    {}
func (*PongPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{3}
}
func (m *PongPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongPayload.Unmarshal(m, b)
}
func (m *PongPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongPayload.Marshal(b, m, deterministic)
}
func (dst *PongPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongPayload.Merge(dst, src)
}
func (m *PongPayload) XXX_Size() int {
	return xxx_messageInfo_PongPayload.Size(m)
}
func (m *PongPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PongPayload.DiscardUnknown(m)
}

var xxx_messageInfo_PongPayload proto.InternalMessageInfo

func (m *PongPayload) GetTimestamp() float32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type StartBackup struct {
	BackupType           BackupType      `protobuf:"varint,1,opt,name=BackupType,proto3,enum=messages.BackupType" json:"BackupType,omitempty"`
	DestinationType      DestinationType `protobuf:"varint,2,opt,name=DestinationType,proto3,enum=messages.DestinationType" json:"DestinationType,omitempty"`
	DestinationName      string          `protobuf:"bytes,3,opt,name=DestinationName,proto3" json:"DestinationName,omitempty"`
	DestinationDir       string          `protobuf:"bytes,4,opt,name=DestinationDir,proto3" json:"DestinationDir,omitempty"`
	CompressionType      CompressionType `protobuf:"varint,5,opt,name=CompressionType,proto3,enum=messages.CompressionType" json:"CompressionType,omitempty"`
	Cypher               Cypher          `protobuf:"varint,6,opt,name=Cypher,proto3,enum=messages.Cypher" json:"Cypher,omitempty"`
	OplogStartTime       float32         `protobuf:"fixed32,7,opt,name=OplogStartTime,proto3" json:"OplogStartTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StartBackup) Reset()         { *m = StartBackup{} }
func (m *StartBackup) String() string { return proto.CompactTextString(m) }
func (*StartBackup) ProtoMessage()    {}
func (*StartBackup) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{4}
}
func (m *StartBackup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartBackup.Unmarshal(m, b)
}
func (m *StartBackup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartBackup.Marshal(b, m, deterministic)
}
func (dst *StartBackup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBackup.Merge(dst, src)
}
func (m *StartBackup) XXX_Size() int {
	return xxx_messageInfo_StartBackup.Size(m)
}
func (m *StartBackup) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBackup.DiscardUnknown(m)
}

var xxx_messageInfo_StartBackup proto.InternalMessageInfo

func (m *StartBackup) GetBackupType() BackupType {
	if m != nil {
		return m.BackupType
	}
	return BackupType_LOGICAL
}

func (m *StartBackup) GetDestinationType() DestinationType {
	if m != nil {
		return m.DestinationType
	}
	return DestinationType_FILE
}

func (m *StartBackup) GetDestinationName() string {
	if m != nil {
		return m.DestinationName
	}
	return ""
}

func (m *StartBackup) GetDestinationDir() string {
	if m != nil {
		return m.DestinationDir
	}
	return ""
}

func (m *StartBackup) GetCompressionType() CompressionType {
	if m != nil {
		return m.CompressionType
	}
	return CompressionType_NO_COMPRESSION
}

func (m *StartBackup) GetCypher() Cypher {
	if m != nil {
		return m.Cypher
	}
	return Cypher_NO_CYPHER
}

func (m *StartBackup) GetOplogStartTime() float32 {
	if m != nil {
		return m.OplogStartTime
	}
	return 0
}

type StopBackup struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopBackup) Reset()         { *m = StopBackup{} }
func (m *StopBackup) String() string { return proto.CompactTextString(m) }
func (*StopBackup) ProtoMessage()    {}
func (*StopBackup) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{5}
}
func (m *StopBackup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopBackup.Unmarshal(m, b)
}
func (m *StopBackup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopBackup.Marshal(b, m, deterministic)
}
func (dst *StopBackup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopBackup.Merge(dst, src)
}
func (m *StopBackup) XXX_Size() int {
	return xxx_messageInfo_StopBackup.Size(m)
}
func (m *StopBackup) XXX_DiscardUnknown() {
	xxx_messageInfo_StopBackup.DiscardUnknown(m)
}

var xxx_messageInfo_StopBackup proto.InternalMessageInfo

type Status struct {
	DBBackUpRunning    bool       `protobuf:"varint,1,opt,name=DBBackUpRunning,proto3" json:"DBBackUpRunning,omitempty"`
	OplogBackupRunning bool       `protobuf:"varint,2,opt,name=OplogBackupRunning,proto3" json:"OplogBackupRunning,omitempty"`
	BackupType         BackupType `protobuf:"varint,3,opt,name=BackupType,proto3,enum=messages.BackupType" json:"BackupType,omitempty"`
	BytesSent          uint64     `protobuf:"varint,4,opt,name=BytesSent,proto3" json:"BytesSent,omitempty"`
	LastOplogTS        float32    `protobuf:"fixed32,5,opt,name=LastOplogTS,proto3" json:"LastOplogTS,omitempty"`
	BackupCompleted    float32    `protobuf:"fixed32,6,opt,name=BackupCompleted,proto3" json:"BackupCompleted,omitempty"`
	LastError          string     `protobuf:"bytes,7,opt,name=LastError,proto3" json:"LastError,omitempty"`
	//
	DestinationType      DestinationType `protobuf:"varint,8,opt,name=DestinationType,proto3,enum=messages.DestinationType" json:"DestinationType,omitempty"`
	DestinationName      string          `protobuf:"bytes,9,opt,name=DestinationName,proto3" json:"DestinationName,omitempty"`
	DestinationDir       string          `protobuf:"bytes,10,opt,name=DestinationDir,proto3" json:"DestinationDir,omitempty"`
	CompressionType      CompressionType `protobuf:"varint,11,opt,name=CompressionType,proto3,enum=messages.CompressionType" json:"CompressionType,omitempty"`
	Cypher               Cypher          `protobuf:"varint,12,opt,name=Cypher,proto3,enum=messages.Cypher" json:"Cypher,omitempty"`
	OplogStartTime       float32         `protobuf:"fixed32,13,opt,name=OplogStartTime,proto3" json:"OplogStartTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{6}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetDBBackUpRunning() bool {
	if m != nil {
		return m.DBBackUpRunning
	}
	return false
}

func (m *Status) GetOplogBackupRunning() bool {
	if m != nil {
		return m.OplogBackupRunning
	}
	return false
}

func (m *Status) GetBackupType() BackupType {
	if m != nil {
		return m.BackupType
	}
	return BackupType_LOGICAL
}

func (m *Status) GetBytesSent() uint64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *Status) GetLastOplogTS() float32 {
	if m != nil {
		return m.LastOplogTS
	}
	return 0
}

func (m *Status) GetBackupCompleted() float32 {
	if m != nil {
		return m.BackupCompleted
	}
	return 0
}

func (m *Status) GetLastError() string {
	if m != nil {
		return m.LastError
	}
	return ""
}

func (m *Status) GetDestinationType() DestinationType {
	if m != nil {
		return m.DestinationType
	}
	return DestinationType_FILE
}

func (m *Status) GetDestinationName() string {
	if m != nil {
		return m.DestinationName
	}
	return ""
}

func (m *Status) GetDestinationDir() string {
	if m != nil {
		return m.DestinationDir
	}
	return ""
}

func (m *Status) GetCompressionType() CompressionType {
	if m != nil {
		return m.CompressionType
	}
	return CompressionType_NO_COMPRESSION
}

func (m *Status) GetCypher() Cypher {
	if m != nil {
		return m.Cypher
	}
	return Cypher_NO_CYPHER
}

func (m *Status) GetOplogStartTime() float32 {
	if m != nil {
		return m.OplogStartTime
	}
	return 0
}

type BackupFinished struct {
	LastOplogTS          float32  `protobuf:"fixed32,1,opt,name=LastOplogTS,proto3" json:"LastOplogTS,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackupFinished) Reset()         { *m = BackupFinished{} }
func (m *BackupFinished) String() string { return proto.CompactTextString(m) }
func (*BackupFinished) ProtoMessage()    {}
func (*BackupFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_06f594d554e9be0b, []int{7}
}
func (m *BackupFinished) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupFinished.Unmarshal(m, b)
}
func (m *BackupFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupFinished.Marshal(b, m, deterministic)
}
func (dst *BackupFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupFinished.Merge(dst, src)
}
func (m *BackupFinished) XXX_Size() int {
	return xxx_messageInfo_BackupFinished.Size(m)
}
func (m *BackupFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupFinished.DiscardUnknown(m)
}

var xxx_messageInfo_BackupFinished proto.InternalMessageInfo

func (m *BackupFinished) GetLastOplogTS() float32 {
	if m != nil {
		return m.LastOplogTS
	}
	return 0
}

func init() {
	proto.RegisterType((*ServerMessage)(nil), "messages.ServerMessage")
	proto.RegisterType((*ClientMessage)(nil), "messages.ClientMessage")
	proto.RegisterType((*RegisterPayload)(nil), "messages.RegisterPayload")
	proto.RegisterType((*PongPayload)(nil), "messages.PongPayload")
	proto.RegisterType((*StartBackup)(nil), "messages.StartBackup")
	proto.RegisterType((*StopBackup)(nil), "messages.StopBackup")
	proto.RegisterType((*Status)(nil), "messages.Status")
	proto.RegisterType((*BackupFinished)(nil), "messages.BackupFinished")
	proto.RegisterEnum("messages.BackupType", BackupType_name, BackupType_value)
	proto.RegisterEnum("messages.DestinationType", DestinationType_name, DestinationType_value)
	proto.RegisterEnum("messages.CompressionType", CompressionType_name, CompressionType_value)
	proto.RegisterEnum("messages.Cypher", Cypher_name, Cypher_value)
	proto.RegisterEnum("messages.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("messages.ServerMessage_MessageType", ServerMessage_MessageType_name, ServerMessage_MessageType_value)
	proto.RegisterEnum("messages.ClientMessage_MessageType", ClientMessage_MessageType_name, ClientMessage_MessageType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessagesClient is the client API for Messages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessagesClient interface {
	MessagesChat(ctx context.Context, opts ...grpc.CallOption) (Messages_MessagesChatClient, error)
}

type messagesClient struct {
	cc *grpc.ClientConn
}

func NewMessagesClient(cc *grpc.ClientConn) MessagesClient {
	return &messagesClient{cc}
}

func (c *messagesClient) MessagesChat(ctx context.Context, opts ...grpc.CallOption) (Messages_MessagesChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Messages_serviceDesc.Streams[0], "/messages.Messages/MessagesChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagesMessagesChatClient{stream}
	return x, nil
}

type Messages_MessagesChatClient interface {
	Send(*ClientMessage) error
	Recv() (*ServerMessage, error)
	grpc.ClientStream
}

type messagesMessagesChatClient struct {
	grpc.ClientStream
}

func (x *messagesMessagesChatClient) Send(m *ClientMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messagesMessagesChatClient) Recv() (*ServerMessage, error) {
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessagesServer is the server API for Messages service.
type MessagesServer interface {
	MessagesChat(Messages_MessagesChatServer) error
}

func RegisterMessagesServer(s *grpc.Server, srv MessagesServer) {
	s.RegisterService(&_Messages_serviceDesc, srv)
}

func _Messages_MessagesChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagesServer).MessagesChat(&messagesMessagesChatServer{stream})
}

type Messages_MessagesChatServer interface {
	Send(*ServerMessage) error
	Recv() (*ClientMessage, error)
	grpc.ServerStream
}

type messagesMessagesChatServer struct {
	grpc.ServerStream
}

func (x *messagesMessagesChatServer) Send(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messagesMessagesChatServer) Recv() (*ClientMessage, error) {
	m := new(ClientMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Messages_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.Messages",
	HandlerType: (*MessagesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MessagesChat",
			Handler:       _Messages_MessagesChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_06f594d554e9be0b) }

var fileDescriptor_message_06f594d554e9be0b = []byte{
	// 979 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5d, 0x6f, 0xa3, 0x46,
	0x14, 0x35, 0x18, 0x7f, 0x5d, 0x7f, 0x64, 0x76, 0x9a, 0x6a, 0xbd, 0xab, 0x7d, 0x88, 0xdc, 0xaa,
	0xb2, 0xd2, 0xca, 0x6a, 0xdd, 0x48, 0x7d, 0x68, 0xab, 0x0a, 0x03, 0xb6, 0xd1, 0x3a, 0x80, 0x66,
	0x48, 0xab, 0xdd, 0x17, 0x8b, 0x26, 0x23, 0x07, 0x35, 0x01, 0x04, 0x64, 0x55, 0xff, 0xc5, 0x3e,
	0xf4, 0xbf, 0xf4, 0xad, 0x8f, 0xd5, 0x0c, 0x10, 0x30, 0x49, 0xaa, 0x54, 0xdb, 0x27, 0x0f, 0x67,
	0xee, 0xb9, 0xf7, 0x72, 0xce, 0xf5, 0x0c, 0x30, 0xbc, 0x65, 0x49, 0xe2, 0xed, 0xd8, 0x2c, 0x8a,
	0xc3, 0x34, 0xc4, 0xdd, 0xfc, 0x31, 0x99, 0xfc, 0x2d, 0xc3, 0x90, 0xb2, 0xf8, 0x03, 0x8b, 0xcf,
	0x33, 0x08, 0x8f, 0xa1, 0xf3, 0x81, 0xc5, 0x89, 0x1f, 0x06, 0x63, 0xe9, 0x44, 0x9a, 0xb6, 0x48,
	0xf1, 0x88, 0xbf, 0x03, 0x25, 0xdd, 0x47, 0x6c, 0x2c, 0x9f, 0x48, 0xd3, 0xd1, 0xfc, 0xb3, 0x59,
	0x91, 0x64, 0x76, 0x90, 0x60, 0x96, 0xff, 0xba, 0xfb, 0x88, 0x11, 0x41, 0xe0, 0x29, 0xf3, 0xd8,
	0x71, 0xf3, 0x44, 0x9a, 0x0e, 0x48, 0xf1, 0x88, 0x7f, 0x82, 0x11, 0x4d, 0xbd, 0x38, 0x5d, 0x78,
	0x97, 0xbf, 0xdd, 0x45, 0xe7, 0xc9, 0x6e, 0xac, 0x9c, 0x48, 0xd3, 0xfe, 0xfc, 0xd3, 0x4a, 0xf2,
	0x72, 0x7f, 0xdd, 0x20, 0xb5, 0x70, 0xfc, 0x03, 0x0c, 0x69, 0x1a, 0x46, 0x25, 0xbf, 0x25, 0xf8,
	0xc7, 0x55, 0x7e, 0xb1, 0xbd, 0x6e, 0x90, 0xc3, 0xe0, 0x49, 0x04, 0xfd, 0x4a, 0xb7, 0xb8, 0x07,
	0x2d, 0x83, 0x10, 0x9b, 0xa0, 0x06, 0xee, 0x82, 0xe2, 0x98, 0xd6, 0x0a, 0x49, 0xf8, 0x13, 0x38,
	0x22, 0xc6, 0xca, 0xa4, 0x2e, 0x51, 0x5d, 0xd3, 0xb6, 0xb6, 0xf6, 0x5b, 0x24, 0x63, 0x04, 0x03,
	0xea, 0xaa, 0xc4, 0xdd, 0x2e, 0x54, 0xed, 0xed, 0x85, 0x83, 0x9a, 0xf8, 0x08, 0xfa, 0xd4, 0xb5,
	0x9d, 0x02, 0x50, 0xf0, 0x08, 0x60, 0x65, 0xb8, 0x5b, 0xea, 0xaa, 0xee, 0x05, 0x45, 0xad, 0x89,
	0xd2, 0x6d, 0xa3, 0xcb, 0x45, 0x0f, 0x3a, 0x8e, 0xb7, 0xbf, 0x09, 0xbd, 0xab, 0xc9, 0x1f, 0x4d,
	0x18, 0x6a, 0x37, 0x3e, 0x0b, 0xd2, 0x8f, 0x90, 0xfe, 0x20, 0xc1, 0x23, 0xd2, 0xbf, 0x86, 0xee,
	0xa5, 0x08, 0x31, 0x75, 0xa1, 0x7d, 0x8f, 0xdc, 0x3f, 0x57, 0x6d, 0x51, 0x0e, 0x6d, 0xf9, 0x11,
	0xfa, 0x84, 0xed, 0xfc, 0x24, 0x65, 0x71, 0xa9, 0xe9, 0xab, 0xb2, 0x6a, 0xb1, 0x99, 0xbf, 0xca,
	0xba, 0x41, 0xaa, 0xf1, 0xf8, 0x1b, 0xe8, 0x38, 0x7e, 0xb0, 0xe3, 0xd4, 0x76, 0xdd, 0x4e, 0x27,
	0x0c, 0x76, 0x25, 0xad, 0x88, 0xc3, 0x6b, 0x78, 0x91, 0xd9, 0xb2, 0xf4, 0x03, 0x3f, 0xb9, 0x66,
	0x57, 0x9c, 0xdc, 0x11, 0xe4, 0x71, 0x49, 0x3e, 0x0c, 0x59, 0x37, 0xc8, 0x43, 0xd2, 0xc4, 0xff,
	0x57, 0x4f, 0x6d, 0xe1, 0xe9, 0x00, 0xba, 0x99, 0xa7, 0x06, 0x41, 0x32, 0xc6, 0x30, 0xca, 0x5c,
	0xdb, 0x0a, 0x4f, 0x0d, 0x1d, 0x35, 0xf1, 0x0b, 0x18, 0xe6, 0xd8, 0xd2, 0xb4, 0x4c, 0xba, 0x46,
	0x0a, 0x06, 0x68, 0x3f, 0x6d, 0xa6, 0x0a, 0x47, 0x35, 0x51, 0xf0, 0x0c, 0xba, 0x56, 0x78, 0x25,
	0xba, 0x10, 0x76, 0x8e, 0xe6, 0xb8, 0x7c, 0x93, 0x62, 0x87, 0xdc, 0xc7, 0x4c, 0xbe, 0x84, 0x7e,
	0x45, 0x1c, 0xfc, 0x06, 0x7a, 0xae, 0x7f, 0xcb, 0x92, 0xd4, 0xbb, 0x8d, 0x04, 0x5f, 0x26, 0x25,
	0x30, 0xf9, 0x4b, 0x86, 0x7e, 0xe5, 0xaf, 0x80, 0xcf, 0x00, 0xb2, 0x55, 0xa5, 0xdc, 0x71, 0x5d,
	0x38, 0x51, 0xb0, 0x12, 0x87, 0x35, 0x38, 0xd2, 0x59, 0x92, 0xfa, 0x81, 0x97, 0xfa, 0x61, 0xe0,
	0x96, 0x13, 0x56, 0xf1, 0xba, 0x16, 0x40, 0xea, 0x0c, 0x3c, 0x3d, 0x48, 0x62, 0x79, 0xb7, 0x2c,
	0x9f, 0xb4, 0x3a, 0x8c, 0xbf, 0x80, 0x51, 0x05, 0xd2, 0xfd, 0x58, 0xcc, 0x5d, 0x8f, 0xd4, 0x50,
	0xde, 0x96, 0x16, 0xde, 0x46, 0x31, 0x4b, 0x92, 0xa2, 0xad, 0x56, 0xbd, 0xad, 0x5a, 0x00, 0xa9,
	0x33, 0xf0, 0x14, 0xda, 0xda, 0x3e, 0xba, 0x66, 0xb1, 0x98, 0xc1, 0xd1, 0x1c, 0x55, 0xb8, 0x02,
	0x27, 0xf9, 0x3e, 0x6f, 0xcb, 0x8e, 0x6e, 0xc2, 0x9d, 0xd0, 0x93, 0x4b, 0x2c, 0x06, 0x4f, 0x26,
	0x35, 0x74, 0x32, 0x00, 0x28, 0x8f, 0x8f, 0xc9, 0x9f, 0x0a, 0xb4, 0x69, 0xea, 0xa5, 0x77, 0x89,
	0x50, 0x60, 0xc1, 0xe1, 0x8b, 0x88, 0xdc, 0x05, 0x81, 0x1f, 0xec, 0x84, 0x03, 0x5d, 0x52, 0x87,
	0xf1, 0x0c, 0xb0, 0x48, 0x9a, 0xe5, 0x28, 0x82, 0x65, 0x11, 0xfc, 0xc8, 0x4e, 0xcd, 0xd6, 0xe6,
	0x33, 0x6d, 0x7d, 0x03, 0xbd, 0xc5, 0x3e, 0x65, 0x09, 0x65, 0x41, 0x2a, 0x24, 0x56, 0x48, 0x09,
	0xe0, 0x13, 0xe8, 0x6f, 0xbc, 0x24, 0x15, 0xd5, 0x5c, 0x2a, 0x94, 0x95, 0x49, 0x15, 0xe2, 0xef,
	0x93, 0x65, 0xe3, 0x9a, 0xde, 0xb0, 0x94, 0x5d, 0x09, 0x0d, 0x65, 0x52, 0x87, 0x79, 0x25, 0x4e,
	0x34, 0xe2, 0x38, 0x8c, 0x85, 0x6a, 0x3d, 0x52, 0x02, 0x8f, 0x8d, 0x57, 0xf7, 0xff, 0x18, 0xaf,
	0xde, 0x73, 0xc7, 0x0b, 0x9e, 0x3b, 0x5e, 0xfd, 0x8f, 0x18, 0xaf, 0xc1, 0x7f, 0x1e, 0xaf, 0xe1,
	0xa3, 0xe3, 0x35, 0x87, 0xd1, 0xe1, 0x69, 0x56, 0x77, 0x4a, 0x7a, 0xe0, 0xd4, 0xe9, 0xb4, 0x3a,
	0x1f, 0xb8, 0x0f, 0x9d, 0x8d, 0xbd, 0x32, 0x35, 0x75, 0x83, 0x1a, 0x78, 0x08, 0xbd, 0xb5, 0xed,
	0xe6, 0xd7, 0x91, 0x74, 0xfa, 0xf9, 0x03, 0x2f, 0xf8, 0x79, 0xb8, 0x34, 0x37, 0x06, 0x6a, 0xe0,
	0x0e, 0x34, 0xd5, 0x5f, 0x28, 0x92, 0x4e, 0xf5, 0x07, 0xd2, 0xf0, 0xd3, 0xd1, 0xb2, 0xb7, 0x9a,
	0x7d, 0xee, 0x10, 0x83, 0x52, 0xd3, 0xb6, 0xb2, 0x93, 0x74, 0xf5, 0xde, 0x74, 0x90, 0x24, 0x0e,
	0x45, 0x4b, 0x75, 0x9c, 0x77, 0x48, 0xe6, 0x59, 0x36, 0xef, 0xcf, 0x50, 0xf3, 0xf4, 0xfb, 0x42,
	0x1b, 0xde, 0x04, 0x27, 0xbf, 0x73, 0xd6, 0x06, 0xc9, 0xeb, 0x18, 0x14, 0x49, 0x7c, 0xa1, 0x1b,
	0x34, 0xe3, 0x10, 0xed, 0x0c, 0x35, 0xc5, 0x82, 0xaa, 0x48, 0x39, 0x8d, 0xca, 0x63, 0x93, 0xd3,
	0x2f, 0x2c, 0xdd, 0x58, 0x9a, 0x96, 0xa1, 0xa3, 0x06, 0x2f, 0x76, 0x6e, 0x5b, 0x2b, 0x5b, 0x47,
	0x12, 0x6f, 0x2b, 0x5b, 0x6f, 0x89, 0xe1, 0x6c, 0xa8, 0xe1, 0x22, 0x99, 0x5f, 0xd5, 0x39, 0x46,
	0xd7, 0x2a, 0xd1, 0xe9, 0xcf, 0x04, 0x35, 0xf1, 0x31, 0xa0, 0x1c, 0xd4, 0x6c, 0x6b, 0x69, 0xae,
	0x38, 0xaa, 0xdc, 0xa7, 0xa2, 0xa8, 0x35, 0x27, 0xd0, 0xcd, 0x6f, 0x8c, 0x04, 0x2f, 0x61, 0x50,
	0xac, 0xb5, 0x6b, 0x2f, 0xc5, 0x2f, 0x9f, 0xb8, 0x6a, 0x5f, 0xbf, 0x7c, 0xe2, 0xf3, 0x67, 0xd2,
	0x98, 0x4a, 0x5f, 0x4b, 0x8b, 0xaf, 0xe0, 0x95, 0x1f, 0xce, 0x76, 0x71, 0x74, 0x39, 0x63, 0xbf,
	0x7b, 0xfc, 0xef, 0x92, 0xdc, 0x13, 0x16, 0xc3, 0xa2, 0x84, 0xc3, 0xbf, 0xc6, 0x1c, 0xe9, 0xd7,
	0xb6, 0xf8, 0x2c, 0xfb, 0xf6, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x92, 0x92, 0x54, 0xa7,
	0x09, 0x00, 0x00,
}