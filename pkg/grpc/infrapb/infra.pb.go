// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra.proto

package infra

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

type MachineClassMeta struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Revision             int32    `protobuf:"varint,2,opt,name=revision" json:"revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineClassMeta) Reset()         { *m = MachineClassMeta{} }
func (m *MachineClassMeta) String() string { return proto.CompactTextString(m) }
func (*MachineClassMeta) ProtoMessage()    {}
func (*MachineClassMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{0}
}
func (m *MachineClassMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MachineClassMeta.Unmarshal(m, b)
}
func (m *MachineClassMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MachineClassMeta.Marshal(b, m, deterministic)
}
func (dst *MachineClassMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineClassMeta.Merge(dst, src)
}
func (m *MachineClassMeta) XXX_Size() int {
	return xxx_messageInfo_MachineClassMeta.Size(m)
}
func (m *MachineClassMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineClassMeta.DiscardUnknown(m)
}

var xxx_messageInfo_MachineClassMeta proto.InternalMessageInfo

func (m *MachineClassMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MachineClassMeta) GetRevision() int32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

type CloudConfigMeta struct {
	MachineClassMeta     *MachineClassMeta `protobuf:"bytes,1,opt,name=machineClassMeta" json:"machineClassMeta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CloudConfigMeta) Reset()         { *m = CloudConfigMeta{} }
func (m *CloudConfigMeta) String() string { return proto.CompactTextString(m) }
func (*CloudConfigMeta) ProtoMessage()    {}
func (*CloudConfigMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{1}
}
func (m *CloudConfigMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudConfigMeta.Unmarshal(m, b)
}
func (m *CloudConfigMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudConfigMeta.Marshal(b, m, deterministic)
}
func (dst *CloudConfigMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudConfigMeta.Merge(dst, src)
}
func (m *CloudConfigMeta) XXX_Size() int {
	return xxx_messageInfo_CloudConfigMeta.Size(m)
}
func (m *CloudConfigMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudConfigMeta.DiscardUnknown(m)
}

var xxx_messageInfo_CloudConfigMeta proto.InternalMessageInfo

func (m *CloudConfigMeta) GetMachineClassMeta() *MachineClassMeta {
	if m != nil {
		return m.MachineClassMeta
	}
	return nil
}

type MachineClass struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineClass) Reset()         { *m = MachineClass{} }
func (m *MachineClass) String() string { return proto.CompactTextString(m) }
func (*MachineClass) ProtoMessage()    {}
func (*MachineClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{2}
}
func (m *MachineClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MachineClass.Unmarshal(m, b)
}
func (m *MachineClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MachineClass.Marshal(b, m, deterministic)
}
func (dst *MachineClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineClass.Merge(dst, src)
}
func (m *MachineClass) XXX_Size() int {
	return xxx_messageInfo_MachineClass.Size(m)
}
func (m *MachineClass) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineClass.DiscardUnknown(m)
}

var xxx_messageInfo_MachineClass proto.InternalMessageInfo

func (m *MachineClass) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *MachineClass) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CloudConfig struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudConfig) Reset()         { *m = CloudConfig{} }
func (m *CloudConfig) String() string { return proto.CompactTextString(m) }
func (*CloudConfig) ProtoMessage()    {}
func (*CloudConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{3}
}
func (m *CloudConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudConfig.Unmarshal(m, b)
}
func (m *CloudConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudConfig.Marshal(b, m, deterministic)
}
func (dst *CloudConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudConfig.Merge(dst, src)
}
func (m *CloudConfig) XXX_Size() int {
	return xxx_messageInfo_CloudConfig.Size(m)
}
func (m *CloudConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CloudConfig proto.InternalMessageInfo

func (m *CloudConfig) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *CloudConfig) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// The message containing driver side structures
type DriverSide struct {
	// operationID to match response with asynchronous requests
	OperationID int32 `protobuf:"varint,1,opt,name=operationID" json:"operationID,omitempty"`
	// operationType can be register, create, delete, getVMs
	OperationType string `protobuf:"bytes,2,opt,name=operationType" json:"operationType,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*DriverSide_RegisterResp
	//	*DriverSide_Createresponse
	//	*DriverSide_Deleteresponse
	//	*DriverSide_Listresponse
	Response             isDriverSide_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DriverSide) Reset()         { *m = DriverSide{} }
func (m *DriverSide) String() string { return proto.CompactTextString(m) }
func (*DriverSide) ProtoMessage()    {}
func (*DriverSide) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{4}
}
func (m *DriverSide) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverSide.Unmarshal(m, b)
}
func (m *DriverSide) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverSide.Marshal(b, m, deterministic)
}
func (dst *DriverSide) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverSide.Merge(dst, src)
}
func (m *DriverSide) XXX_Size() int {
	return xxx_messageInfo_DriverSide.Size(m)
}
func (m *DriverSide) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverSide.DiscardUnknown(m)
}

var xxx_messageInfo_DriverSide proto.InternalMessageInfo

type isDriverSide_Response interface {
	isDriverSide_Response()
}

type DriverSide_RegisterResp struct {
	RegisterResp *DriverSideRegisterationResp `protobuf:"bytes,3,opt,name=registerResp,oneof"`
}
type DriverSide_Createresponse struct {
	Createresponse *DriverSideCreateResp `protobuf:"bytes,4,opt,name=createresponse,oneof"`
}
type DriverSide_Deleteresponse struct {
	Deleteresponse *DriverSideDeleteResp `protobuf:"bytes,5,opt,name=deleteresponse,oneof"`
}
type DriverSide_Listresponse struct {
	Listresponse *DriverSideListResp `protobuf:"bytes,6,opt,name=listresponse,oneof"`
}

func (*DriverSide_RegisterResp) isDriverSide_Response()   {}
func (*DriverSide_Createresponse) isDriverSide_Response() {}
func (*DriverSide_Deleteresponse) isDriverSide_Response() {}
func (*DriverSide_Listresponse) isDriverSide_Response()   {}

func (m *DriverSide) GetResponse() isDriverSide_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *DriverSide) GetOperationID() int32 {
	if m != nil {
		return m.OperationID
	}
	return 0
}

func (m *DriverSide) GetOperationType() string {
	if m != nil {
		return m.OperationType
	}
	return ""
}

func (m *DriverSide) GetRegisterResp() *DriverSideRegisterationResp {
	if x, ok := m.GetResponse().(*DriverSide_RegisterResp); ok {
		return x.RegisterResp
	}
	return nil
}

func (m *DriverSide) GetCreateresponse() *DriverSideCreateResp {
	if x, ok := m.GetResponse().(*DriverSide_Createresponse); ok {
		return x.Createresponse
	}
	return nil
}

func (m *DriverSide) GetDeleteresponse() *DriverSideDeleteResp {
	if x, ok := m.GetResponse().(*DriverSide_Deleteresponse); ok {
		return x.Deleteresponse
	}
	return nil
}

func (m *DriverSide) GetListresponse() *DriverSideListResp {
	if x, ok := m.GetResponse().(*DriverSide_Listresponse); ok {
		return x.Listresponse
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DriverSide) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DriverSide_OneofMarshaler, _DriverSide_OneofUnmarshaler, _DriverSide_OneofSizer, []interface{}{
		(*DriverSide_RegisterResp)(nil),
		(*DriverSide_Createresponse)(nil),
		(*DriverSide_Deleteresponse)(nil),
		(*DriverSide_Listresponse)(nil),
	}
}

func _DriverSide_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DriverSide)
	// response
	switch x := m.Response.(type) {
	case *DriverSide_RegisterResp:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RegisterResp); err != nil {
			return err
		}
	case *DriverSide_Createresponse:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Createresponse); err != nil {
			return err
		}
	case *DriverSide_Deleteresponse:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Deleteresponse); err != nil {
			return err
		}
	case *DriverSide_Listresponse:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Listresponse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DriverSide.Response has unexpected type %T", x)
	}
	return nil
}

func _DriverSide_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DriverSide)
	switch tag {
	case 3: // response.registerResp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DriverSideRegisterationResp)
		err := b.DecodeMessage(msg)
		m.Response = &DriverSide_RegisterResp{msg}
		return true, err
	case 4: // response.createresponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DriverSideCreateResp)
		err := b.DecodeMessage(msg)
		m.Response = &DriverSide_Createresponse{msg}
		return true, err
	case 5: // response.deleteresponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DriverSideDeleteResp)
		err := b.DecodeMessage(msg)
		m.Response = &DriverSide_Deleteresponse{msg}
		return true, err
	case 6: // response.listresponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DriverSideListResp)
		err := b.DecodeMessage(msg)
		m.Response = &DriverSide_Listresponse{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DriverSide_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DriverSide)
	// response
	switch x := m.Response.(type) {
	case *DriverSide_RegisterResp:
		s := proto.Size(x.RegisterResp)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DriverSide_Createresponse:
		s := proto.Size(x.Createresponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DriverSide_Deleteresponse:
		s := proto.Size(x.Deleteresponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DriverSide_Listresponse:
		s := proto.Size(x.Listresponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DriverSideRegisterationResp struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Group                string   `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	Kind                 string   `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverSideRegisterationResp) Reset()         { *m = DriverSideRegisterationResp{} }
func (m *DriverSideRegisterationResp) String() string { return proto.CompactTextString(m) }
func (*DriverSideRegisterationResp) ProtoMessage()    {}
func (*DriverSideRegisterationResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{4, 0}
}
func (m *DriverSideRegisterationResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverSideRegisterationResp.Unmarshal(m, b)
}
func (m *DriverSideRegisterationResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverSideRegisterationResp.Marshal(b, m, deterministic)
}
func (dst *DriverSideRegisterationResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverSideRegisterationResp.Merge(dst, src)
}
func (m *DriverSideRegisterationResp) XXX_Size() int {
	return xxx_messageInfo_DriverSideRegisterationResp.Size(m)
}
func (m *DriverSideRegisterationResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverSideRegisterationResp.DiscardUnknown(m)
}

var xxx_messageInfo_DriverSideRegisterationResp proto.InternalMessageInfo

func (m *DriverSideRegisterationResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DriverSideRegisterationResp) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *DriverSideRegisterationResp) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *DriverSideRegisterationResp) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type DriverSideCreateResp struct {
	ProviderID           string   `protobuf:"bytes,1,opt,name=providerID" json:"providerID,omitempty"`
	Nodename             string   `protobuf:"bytes,2,opt,name=nodename" json:"nodename,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverSideCreateResp) Reset()         { *m = DriverSideCreateResp{} }
func (m *DriverSideCreateResp) String() string { return proto.CompactTextString(m) }
func (*DriverSideCreateResp) ProtoMessage()    {}
func (*DriverSideCreateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{4, 1}
}
func (m *DriverSideCreateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverSideCreateResp.Unmarshal(m, b)
}
func (m *DriverSideCreateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverSideCreateResp.Marshal(b, m, deterministic)
}
func (dst *DriverSideCreateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverSideCreateResp.Merge(dst, src)
}
func (m *DriverSideCreateResp) XXX_Size() int {
	return xxx_messageInfo_DriverSideCreateResp.Size(m)
}
func (m *DriverSideCreateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverSideCreateResp.DiscardUnknown(m)
}

var xxx_messageInfo_DriverSideCreateResp proto.InternalMessageInfo

func (m *DriverSideCreateResp) GetProviderID() string {
	if m != nil {
		return m.ProviderID
	}
	return ""
}

func (m *DriverSideCreateResp) GetNodename() string {
	if m != nil {
		return m.Nodename
	}
	return ""
}

func (m *DriverSideCreateResp) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DriverSideDeleteResp struct {
	Error                string   `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverSideDeleteResp) Reset()         { *m = DriverSideDeleteResp{} }
func (m *DriverSideDeleteResp) String() string { return proto.CompactTextString(m) }
func (*DriverSideDeleteResp) ProtoMessage()    {}
func (*DriverSideDeleteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{4, 2}
}
func (m *DriverSideDeleteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverSideDeleteResp.Unmarshal(m, b)
}
func (m *DriverSideDeleteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverSideDeleteResp.Marshal(b, m, deterministic)
}
func (dst *DriverSideDeleteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverSideDeleteResp.Merge(dst, src)
}
func (m *DriverSideDeleteResp) XXX_Size() int {
	return xxx_messageInfo_DriverSideDeleteResp.Size(m)
}
func (m *DriverSideDeleteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverSideDeleteResp.DiscardUnknown(m)
}

var xxx_messageInfo_DriverSideDeleteResp proto.InternalMessageInfo

func (m *DriverSideDeleteResp) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DriverSideListResp struct {
	List                 []string `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverSideListResp) Reset()         { *m = DriverSideListResp{} }
func (m *DriverSideListResp) String() string { return proto.CompactTextString(m) }
func (*DriverSideListResp) ProtoMessage()    {}
func (*DriverSideListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{4, 3}
}
func (m *DriverSideListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverSideListResp.Unmarshal(m, b)
}
func (m *DriverSideListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverSideListResp.Marshal(b, m, deterministic)
}
func (dst *DriverSideListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverSideListResp.Merge(dst, src)
}
func (m *DriverSideListResp) XXX_Size() int {
	return xxx_messageInfo_DriverSideListResp.Size(m)
}
func (m *DriverSideListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverSideListResp.DiscardUnknown(m)
}

var xxx_messageInfo_DriverSideListResp proto.InternalMessageInfo

func (m *DriverSideListResp) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *DriverSideListResp) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// The message containing the MCM side structures
type MCMside struct {
	// operationID to match response with asynchronous requests
	OperationID int32 `protobuf:"varint,1,opt,name=operationID" json:"operationID,omitempty"`
	// operationType can be register, create, delete, getVMs
	OperationType        string                  `protobuf:"bytes,2,opt,name=operationType" json:"operationType,omitempty"`
	Operationparams      *MCMsideOperationParams `protobuf:"bytes,3,opt,name=operationparams" json:"operationparams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MCMside) Reset()         { *m = MCMside{} }
func (m *MCMside) String() string { return proto.CompactTextString(m) }
func (*MCMside) ProtoMessage()    {}
func (*MCMside) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{5}
}
func (m *MCMside) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MCMside.Unmarshal(m, b)
}
func (m *MCMside) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MCMside.Marshal(b, m, deterministic)
}
func (dst *MCMside) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MCMside.Merge(dst, src)
}
func (m *MCMside) XXX_Size() int {
	return xxx_messageInfo_MCMside.Size(m)
}
func (m *MCMside) XXX_DiscardUnknown() {
	xxx_messageInfo_MCMside.DiscardUnknown(m)
}

var xxx_messageInfo_MCMside proto.InternalMessageInfo

func (m *MCMside) GetOperationID() int32 {
	if m != nil {
		return m.OperationID
	}
	return 0
}

func (m *MCMside) GetOperationType() string {
	if m != nil {
		return m.OperationType
	}
	return ""
}

func (m *MCMside) GetOperationparams() *MCMsideOperationParams {
	if m != nil {
		return m.Operationparams
	}
	return nil
}

type MCMsideOperationParams struct {
	MachineClassMetaData *MachineClassMeta `protobuf:"bytes,1,opt,name=machineClassMetaData" json:"machineClassMetaData,omitempty"`
	// cloudConfig has the secrets (username/password info to connect to provider account)
	Credentials          string   `protobuf:"bytes,2,opt,name=credentials" json:"credentials,omitempty"`
	MachineID            string   `protobuf:"bytes,3,opt,name=machineID" json:"machineID,omitempty"`
	MachineName          string   `protobuf:"bytes,4,opt,name=machineName" json:"machineName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MCMsideOperationParams) Reset()         { *m = MCMsideOperationParams{} }
func (m *MCMsideOperationParams) String() string { return proto.CompactTextString(m) }
func (*MCMsideOperationParams) ProtoMessage()    {}
func (*MCMsideOperationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_infra_835ce489110120ef, []int{5, 0}
}
func (m *MCMsideOperationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MCMsideOperationParams.Unmarshal(m, b)
}
func (m *MCMsideOperationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MCMsideOperationParams.Marshal(b, m, deterministic)
}
func (dst *MCMsideOperationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MCMsideOperationParams.Merge(dst, src)
}
func (m *MCMsideOperationParams) XXX_Size() int {
	return xxx_messageInfo_MCMsideOperationParams.Size(m)
}
func (m *MCMsideOperationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MCMsideOperationParams.DiscardUnknown(m)
}

var xxx_messageInfo_MCMsideOperationParams proto.InternalMessageInfo

func (m *MCMsideOperationParams) GetMachineClassMetaData() *MachineClassMeta {
	if m != nil {
		return m.MachineClassMetaData
	}
	return nil
}

func (m *MCMsideOperationParams) GetCredentials() string {
	if m != nil {
		return m.Credentials
	}
	return ""
}

func (m *MCMsideOperationParams) GetMachineID() string {
	if m != nil {
		return m.MachineID
	}
	return ""
}

func (m *MCMsideOperationParams) GetMachineName() string {
	if m != nil {
		return m.MachineName
	}
	return ""
}

func init() {
	proto.RegisterType((*MachineClassMeta)(nil), "infra.MachineClassMeta")
	proto.RegisterType((*CloudConfigMeta)(nil), "infra.CloudConfigMeta")
	proto.RegisterType((*MachineClass)(nil), "infra.MachineClass")
	proto.RegisterType((*CloudConfig)(nil), "infra.CloudConfig")
	proto.RegisterType((*DriverSide)(nil), "infra.DriverSide")
	proto.RegisterType((*DriverSideRegisterationResp)(nil), "infra.DriverSide.registerationResp")
	proto.RegisterType((*DriverSideCreateResp)(nil), "infra.DriverSide.createResp")
	proto.RegisterType((*DriverSideDeleteResp)(nil), "infra.DriverSide.deleteResp")
	proto.RegisterType((*DriverSideListResp)(nil), "infra.DriverSide.listResp")
	proto.RegisterType((*MCMside)(nil), "infra.MCMside")
	proto.RegisterType((*MCMsideOperationParams)(nil), "infra.MCMside.operationParams")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InfragrpcClient is the client API for Infragrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfragrpcClient interface {
	// infra operations
	Register(ctx context.Context, opts ...grpc.CallOption) (Infragrpc_RegisterClient, error)
	// Get Machine Class
	GetMachineClass(ctx context.Context, in *MachineClassMeta, opts ...grpc.CallOption) (*MachineClass, error)
	// Get Cloud Config
	GetCloudConfig(ctx context.Context, in *CloudConfigMeta, opts ...grpc.CallOption) (*CloudConfig, error)
}

type infragrpcClient struct {
	cc *grpc.ClientConn
}

func NewInfragrpcClient(cc *grpc.ClientConn) InfragrpcClient {
	return &infragrpcClient{cc}
}

func (c *infragrpcClient) Register(ctx context.Context, opts ...grpc.CallOption) (Infragrpc_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Infragrpc_serviceDesc.Streams[0], "/infra.Infragrpc/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &infragrpcRegisterClient{stream}
	return x, nil
}

type Infragrpc_RegisterClient interface {
	Send(*DriverSide) error
	Recv() (*MCMside, error)
	grpc.ClientStream
}

type infragrpcRegisterClient struct {
	grpc.ClientStream
}

func (x *infragrpcRegisterClient) Send(m *DriverSide) error {
	return x.ClientStream.SendMsg(m)
}

func (x *infragrpcRegisterClient) Recv() (*MCMside, error) {
	m := new(MCMside)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *infragrpcClient) GetMachineClass(ctx context.Context, in *MachineClassMeta, opts ...grpc.CallOption) (*MachineClass, error) {
	out := new(MachineClass)
	err := c.cc.Invoke(ctx, "/infra.Infragrpc/GetMachineClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infragrpcClient) GetCloudConfig(ctx context.Context, in *CloudConfigMeta, opts ...grpc.CallOption) (*CloudConfig, error) {
	out := new(CloudConfig)
	err := c.cc.Invoke(ctx, "/infra.Infragrpc/GetCloudConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfragrpcServer is the server API for Infragrpc service.
type InfragrpcServer interface {
	// infra operations
	Register(Infragrpc_RegisterServer) error
	// Get Machine Class
	GetMachineClass(context.Context, *MachineClassMeta) (*MachineClass, error)
	// Get Cloud Config
	GetCloudConfig(context.Context, *CloudConfigMeta) (*CloudConfig, error)
}

func RegisterInfragrpcServer(s *grpc.Server, srv InfragrpcServer) {
	s.RegisterService(&_Infragrpc_serviceDesc, srv)
}

func _Infragrpc_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InfragrpcServer).Register(&infragrpcRegisterServer{stream})
}

type Infragrpc_RegisterServer interface {
	Send(*MCMside) error
	Recv() (*DriverSide, error)
	grpc.ServerStream
}

type infragrpcRegisterServer struct {
	grpc.ServerStream
}

func (x *infragrpcRegisterServer) Send(m *MCMside) error {
	return x.ServerStream.SendMsg(m)
}

func (x *infragrpcRegisterServer) Recv() (*DriverSide, error) {
	m := new(DriverSide)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Infragrpc_GetMachineClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineClassMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfragrpcServer).GetMachineClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.Infragrpc/GetMachineClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfragrpcServer).GetMachineClass(ctx, req.(*MachineClassMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Infragrpc_GetCloudConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudConfigMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfragrpcServer).GetCloudConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.Infragrpc/GetCloudConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfragrpcServer).GetCloudConfig(ctx, req.(*CloudConfigMeta))
	}
	return interceptor(ctx, in, info, handler)
}

var _Infragrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infra.Infragrpc",
	HandlerType: (*InfragrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMachineClass",
			Handler:    _Infragrpc_GetMachineClass_Handler,
		},
		{
			MethodName: "GetCloudConfig",
			Handler:    _Infragrpc_GetCloudConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _Infragrpc_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "infra.proto",
}

func init() { proto.RegisterFile("infra.proto", fileDescriptor_infra_835ce489110120ef) }

var fileDescriptor_infra_835ce489110120ef = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x6e, 0xb6, 0x65, 0x6b, 0xae, 0xa5, 0xdd, 0x4c, 0x05, 0x51, 0x34, 0x4d, 0x55, 0xe0, 0xa1,
	0x4f, 0x15, 0xda, 0x90, 0xe0, 0x09, 0xc1, 0x5a, 0xb1, 0x55, 0xa8, 0x08, 0x19, 0xc4, 0x23, 0x92,
	0x69, 0x6e, 0xc5, 0x5a, 0x1b, 0x47, 0x4e, 0x56, 0x89, 0xff, 0xc3, 0x0f, 0x41, 0xe2, 0x7f, 0xf0,
	0x5b, 0x90, 0x1d, 0x27, 0x71, 0x93, 0x4e, 0x08, 0x89, 0x37, 0xdf, 0xe5, 0xfb, 0x3e, 0xdf, 0x9d,
	0xbf, 0x1c, 0x74, 0x78, 0x7c, 0x23, 0xd9, 0x38, 0x91, 0x22, 0x13, 0xc4, 0xd5, 0x41, 0x78, 0x09,
	0xc7, 0x73, 0xb6, 0xf8, 0xc6, 0x63, 0x9c, 0xac, 0x58, 0x9a, 0xce, 0x31, 0x63, 0x84, 0xc0, 0x41,
	0xcc, 0xd6, 0xe8, 0x3b, 0x43, 0x67, 0xe4, 0x51, 0x7d, 0x26, 0x01, 0xb4, 0x25, 0x6e, 0x78, 0xca,
	0x45, 0xec, 0xef, 0x0d, 0x9d, 0x91, 0x4b, 0xcb, 0x38, 0xfc, 0x0c, 0xfd, 0xc9, 0x4a, 0xdc, 0x45,
	0x13, 0x11, 0xdf, 0xf0, 0xa5, 0x96, 0x98, 0xc0, 0xf1, 0xba, 0x26, 0xab, 0xe5, 0x3a, 0xe7, 0x8f,
	0xc7, 0x79, 0x15, 0xf5, 0x5b, 0x69, 0x83, 0x10, 0xbe, 0x84, 0xae, 0x8d, 0x52, 0x75, 0x45, 0xcc,
	0x08, 0x79, 0x54, 0x9f, 0xc9, 0x00, 0x5c, 0x94, 0x52, 0x48, 0x5d, 0x94, 0x47, 0xf3, 0x20, 0x7c,
	0x01, 0x1d, 0xab, 0xa2, 0x7f, 0x20, 0xfe, 0x70, 0x01, 0xa6, 0x92, 0x6f, 0x50, 0x7e, 0xe4, 0x11,
	0x92, 0x21, 0x74, 0x44, 0x82, 0x92, 0x65, 0x5c, 0xc4, 0xb3, 0xa9, 0xe6, 0xbb, 0xd4, 0x4e, 0x91,
	0xa7, 0xf0, 0xa0, 0x0c, 0x3f, 0x7d, 0x4f, 0xd0, 0xc8, 0x6d, 0x27, 0xc9, 0x0c, 0xba, 0x12, 0x97,
	0x3c, 0xcd, 0x50, 0x52, 0x4c, 0x13, 0x7f, 0x5f, 0x8f, 0xe2, 0x89, 0x19, 0x45, 0x75, 0xe1, 0xb8,
	0x40, 0x69, 0xaa, 0x82, 0x5e, 0xb7, 0xe8, 0x16, 0x95, 0xbc, 0x85, 0xde, 0x42, 0x22, 0xcb, 0x50,
	0x62, 0x9a, 0x88, 0x38, 0x45, 0xff, 0x40, 0x8b, 0x9d, 0x36, 0xc5, 0x72, 0x9c, 0x51, 0xa9, 0xb1,
	0x94, 0x4e, 0x84, 0x2b, 0xb4, 0x74, 0xdc, 0xfb, 0x74, 0x72, 0x5c, 0xa1, 0xb3, 0xcd, 0x22, 0xaf,
	0xa1, 0xbb, 0xe2, 0x69, 0x56, 0xaa, 0x1c, 0x6a, 0x95, 0xa0, 0xa9, 0xa2, 0x50, 0x45, 0x47, 0x36,
	0x23, 0xb8, 0x85, 0x93, 0x46, 0xdb, 0x3b, 0x3d, 0x38, 0x00, 0x77, 0x29, 0xc5, 0x5d, 0x52, 0x3c,
	0x99, 0x0e, 0x14, 0xf2, 0x96, 0xc7, 0x91, 0x9e, 0xa9, 0x47, 0xf5, 0x99, 0xf8, 0x70, 0xb4, 0x41,
	0xa9, 0xcd, 0x7a, 0xa0, 0xd3, 0x45, 0x18, 0x7c, 0x01, 0xa8, 0xc6, 0x42, 0xce, 0x00, 0x12, 0x29,
	0x36, 0x3c, 0x42, 0x69, 0x9e, 0xd7, 0xa3, 0x56, 0x46, 0xb9, 0x3e, 0x16, 0x11, 0xea, 0x4a, 0xf2,
	0x4b, 0xcb, 0xb8, 0x32, 0xd0, 0xbe, 0x65, 0xa0, 0x20, 0x04, 0xa8, 0xc6, 0x55, 0x61, 0x1c, 0x1b,
	0xf3, 0x1c, 0xda, 0xc5, 0x30, 0x54, 0xf5, 0xea, 0xec, 0x3b, 0xc3, 0x7d, 0x55, 0xbd, 0x3a, 0xef,
	0xb6, 0xe6, 0x25, 0xa8, 0x3f, 0x30, 0x1f, 0x59, 0xf8, 0x7b, 0x0f, 0x8e, 0xe6, 0x93, 0x79, 0xfa,
	0x3f, 0x3d, 0x7a, 0x0d, 0xfd, 0x32, 0x91, 0x30, 0xc9, 0xd6, 0xa9, 0xb1, 0xe9, 0x59, 0xf1, 0xc7,
	0xe6, 0x17, 0x8e, 0x4b, 0xd4, 0x07, 0x8d, 0xa2, 0x75, 0x5a, 0xf0, 0xd3, 0xb1, 0xa4, 0x72, 0x10,
	0x79, 0x07, 0x83, 0xfa, 0xff, 0x3d, 0x65, 0x7f, 0x5f, 0x0a, 0x3b, 0x49, 0xaa, 0xe5, 0x85, 0xc4,
	0x08, 0xe3, 0x8c, 0xb3, 0x55, 0x6a, 0xda, 0xb1, 0x53, 0xe4, 0x14, 0x3c, 0xc3, 0x9c, 0x4d, 0xcd,
	0x03, 0x55, 0x09, 0xc5, 0x37, 0xc1, 0x7b, 0xf5, 0xb2, 0xb9, 0x45, 0xec, 0xd4, 0xf9, 0x2f, 0x07,
	0xbc, 0x99, 0x2a, 0x69, 0x29, 0x93, 0x05, 0xb9, 0x80, 0x36, 0x35, 0x0e, 0x25, 0x27, 0x0d, 0x67,
	0x07, 0xbd, 0xed, 0x01, 0x85, 0xad, 0x91, 0xf3, 0xcc, 0x21, 0x6f, 0xa0, 0x7f, 0x85, 0xd9, 0xd6,
	0x02, 0xbb, 0xaf, 0xcd, 0xe0, 0xe1, 0x8e, 0x0f, 0x61, 0x8b, 0xbc, 0x82, 0xde, 0x15, 0x66, 0xf6,
	0x26, 0x7b, 0x64, 0x80, 0xb5, 0x7d, 0x1b, 0x90, 0x66, 0x3e, 0x6c, 0x7d, 0x3d, 0xd4, 0xab, 0xfe,
	0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x88, 0xde, 0x0f, 0xf9, 0x05, 0x00, 0x00,
}
