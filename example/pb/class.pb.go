// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/pb/class.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ClassType int32

const (
	ClassType_INVALID ClassType = 0
	ClassType_REGULAR ClassType = 1
	ClassType_SERIES  ClassType = 2
)

var ClassType_name = map[int32]string{
	0: "INVALID",
	1: "REGULAR",
	2: "SERIES",
}

var ClassType_value = map[string]int32{
	"INVALID": 0,
	"REGULAR": 1,
	"SERIES":  2,
}

func (x ClassType) String() string {
	return proto.EnumName(ClassType_name, int32(x))
}

func (ClassType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5cf53c799ece0d8b, []int{0}
}

type Class struct {
	Id          string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Area        float32            `protobuf:"fixed32,2,opt,name=area,proto3" json:"area,omitempty"`
	Strength    int32              `protobuf:"varint,3,opt,name=strength,proto3" json:"strength,omitempty"`
	IsDeleted   bool               `protobuf:"varint,4,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	Type        ClassType          `protobuf:"varint,5,opt,name=type,proto3,enum=appointy.class.v1.ClassType" json:"type,omitempty"`
	Instructors []*ServiceProvider `protobuf:"bytes,6,rep,name=instructors,proto3" json:"instructors,omitempty"`
	// Types that are valid to be assigned to Charge:
	//	*Class_PerInstance
	//	*Class_Lumpsum
	Charge               isClass_Charge    `protobuf_oneof:"charge"`
	Metadata             map[string]string `protobuf:"bytes,9,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Parent               string            `protobuf:"bytes,10,opt,name=parent,proto3" json:"parent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Class) Reset()         { *m = Class{} }
func (m *Class) String() string { return proto.CompactTextString(m) }
func (*Class) ProtoMessage()    {}
func (*Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf53c799ece0d8b, []int{0}
}

func (m *Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Class.Unmarshal(m, b)
}
func (m *Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Class.Marshal(b, m, deterministic)
}
func (m *Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Class.Merge(m, src)
}
func (m *Class) XXX_Size() int {
	return xxx_messageInfo_Class.Size(m)
}
func (m *Class) XXX_DiscardUnknown() {
	xxx_messageInfo_Class.DiscardUnknown(m)
}

var xxx_messageInfo_Class proto.InternalMessageInfo

func (m *Class) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Class) GetArea() float32 {
	if m != nil {
		return m.Area
	}
	return 0
}

func (m *Class) GetStrength() int32 {
	if m != nil {
		return m.Strength
	}
	return 0
}

func (m *Class) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *Class) GetType() ClassType {
	if m != nil {
		return m.Type
	}
	return ClassType_INVALID
}

func (m *Class) GetInstructors() []*ServiceProvider {
	if m != nil {
		return m.Instructors
	}
	return nil
}

type isClass_Charge interface {
	isClass_Charge()
}

type Class_PerInstance struct {
	PerInstance string `protobuf:"bytes,7,opt,name=per_instance,json=perInstance,proto3,oneof"`
}

type Class_Lumpsum struct {
	Lumpsum int32 `protobuf:"varint,8,opt,name=Lumpsum,proto3,oneof"`
}

func (*Class_PerInstance) isClass_Charge() {}

func (*Class_Lumpsum) isClass_Charge() {}

func (m *Class) GetCharge() isClass_Charge {
	if m != nil {
		return m.Charge
	}
	return nil
}

func (m *Class) GetPerInstance() string {
	if x, ok := m.GetCharge().(*Class_PerInstance); ok {
		return x.PerInstance
	}
	return ""
}

func (m *Class) GetLumpsum() int32 {
	if x, ok := m.GetCharge().(*Class_Lumpsum); ok {
		return x.Lumpsum
	}
	return 0
}

func (m *Class) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Class) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Class) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Class_PerInstance)(nil),
		(*Class_Lumpsum)(nil),
	}
}

type CreateClassReq struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Class                *Class   `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateClassReq) Reset()         { *m = CreateClassReq{} }
func (m *CreateClassReq) String() string { return proto.CompactTextString(m) }
func (*CreateClassReq) ProtoMessage()    {}
func (*CreateClassReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf53c799ece0d8b, []int{1}
}

func (m *CreateClassReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClassReq.Unmarshal(m, b)
}
func (m *CreateClassReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClassReq.Marshal(b, m, deterministic)
}
func (m *CreateClassReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClassReq.Merge(m, src)
}
func (m *CreateClassReq) XXX_Size() int {
	return xxx_messageInfo_CreateClassReq.Size(m)
}
func (m *CreateClassReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClassReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClassReq proto.InternalMessageInfo

func (m *CreateClassReq) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateClassReq) GetClass() *Class {
	if m != nil {
		return m.Class
	}
	return nil
}

type GetClassReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClassReq) Reset()         { *m = GetClassReq{} }
func (m *GetClassReq) String() string { return proto.CompactTextString(m) }
func (*GetClassReq) ProtoMessage()    {}
func (*GetClassReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf53c799ece0d8b, []int{2}
}

func (m *GetClassReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClassReq.Unmarshal(m, b)
}
func (m *GetClassReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClassReq.Marshal(b, m, deterministic)
}
func (m *GetClassReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClassReq.Merge(m, src)
}
func (m *GetClassReq) XXX_Size() int {
	return xxx_messageInfo_GetClassReq.Size(m)
}
func (m *GetClassReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClassReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetClassReq proto.InternalMessageInfo

func (m *GetClassReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ServiceProvider struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceProvider) Reset()         { *m = ServiceProvider{} }
func (m *ServiceProvider) String() string { return proto.CompactTextString(m) }
func (*ServiceProvider) ProtoMessage()    {}
func (*ServiceProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf53c799ece0d8b, []int{3}
}

func (m *ServiceProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceProvider.Unmarshal(m, b)
}
func (m *ServiceProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceProvider.Marshal(b, m, deterministic)
}
func (m *ServiceProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceProvider.Merge(m, src)
}
func (m *ServiceProvider) XXX_Size() int {
	return xxx_messageInfo_ServiceProvider.Size(m)
}
func (m *ServiceProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceProvider.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceProvider proto.InternalMessageInfo

func (m *ServiceProvider) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceProvider) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

type PerInstance struct {
	Cost                 *Price   `protobuf:"bytes,2,opt,name=cost,proto3" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PerInstance) Reset()         { *m = PerInstance{} }
func (m *PerInstance) String() string { return proto.CompactTextString(m) }
func (*PerInstance) ProtoMessage()    {}
func (*PerInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf53c799ece0d8b, []int{4}
}

func (m *PerInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerInstance.Unmarshal(m, b)
}
func (m *PerInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerInstance.Marshal(b, m, deterministic)
}
func (m *PerInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerInstance.Merge(m, src)
}
func (m *PerInstance) XXX_Size() int {
	return xxx_messageInfo_PerInstance.Size(m)
}
func (m *PerInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_PerInstance.DiscardUnknown(m)
}

var xxx_messageInfo_PerInstance proto.InternalMessageInfo

func (m *PerInstance) GetCost() *Price {
	if m != nil {
		return m.Cost
	}
	return nil
}

type Lumpsum struct {
	Cost                 *Price   `protobuf:"bytes,1,opt,name=cost,proto3" json:"cost,omitempty"`
	NumberOfSessions     int32    `protobuf:"varint,2,opt,name=number_of_sessions,json=numberOfSessions,proto3" json:"number_of_sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lumpsum) Reset()         { *m = Lumpsum{} }
func (m *Lumpsum) String() string { return proto.CompactTextString(m) }
func (*Lumpsum) ProtoMessage()    {}
func (*Lumpsum) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf53c799ece0d8b, []int{5}
}

func (m *Lumpsum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lumpsum.Unmarshal(m, b)
}
func (m *Lumpsum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lumpsum.Marshal(b, m, deterministic)
}
func (m *Lumpsum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lumpsum.Merge(m, src)
}
func (m *Lumpsum) XXX_Size() int {
	return xxx_messageInfo_Lumpsum.Size(m)
}
func (m *Lumpsum) XXX_DiscardUnknown() {
	xxx_messageInfo_Lumpsum.DiscardUnknown(m)
}

var xxx_messageInfo_Lumpsum proto.InternalMessageInfo

func (m *Lumpsum) GetCost() *Price {
	if m != nil {
		return m.Cost
	}
	return nil
}

func (m *Lumpsum) GetNumberOfSessions() int32 {
	if m != nil {
		return m.NumberOfSessions
	}
	return 0
}

type Price struct {
	Amount               int32    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency             string   `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf53c799ece0d8b, []int{6}
}

func (m *Price) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Price.Unmarshal(m, b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Price.Marshal(b, m, deterministic)
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return xxx_messageInfo_Price.Size(m)
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Price) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func init() {
	proto.RegisterEnum("appointy.class.v1.ClassType", ClassType_name, ClassType_value)
	proto.RegisterType((*Class)(nil), "appointy.class.v1.Class")
	proto.RegisterMapType((map[string]string)(nil), "appointy.class.v1.Class.MetadataEntry")
	proto.RegisterType((*CreateClassReq)(nil), "appointy.class.v1.CreateClassReq")
	proto.RegisterType((*GetClassReq)(nil), "appointy.class.v1.GetClassReq")
	proto.RegisterType((*ServiceProvider)(nil), "appointy.class.v1.ServiceProvider")
	proto.RegisterType((*PerInstance)(nil), "appointy.class.v1.PerInstance")
	proto.RegisterType((*Lumpsum)(nil), "appointy.class.v1.Lumpsum")
	proto.RegisterType((*Price)(nil), "appointy.class.v1.Price")
}

func init() { proto.RegisterFile("example/pb/class.proto", fileDescriptor_5cf53c799ece0d8b) }

var fileDescriptor_5cf53c799ece0d8b = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x61, 0x6f, 0xd3, 0x30,
	0x10, 0x5d, 0xd2, 0xa6, 0x6b, 0x2e, 0x30, 0x8a, 0x85, 0xa6, 0x68, 0x62, 0x52, 0x14, 0x24, 0x14,
	0xa1, 0x29, 0x63, 0xe3, 0x0b, 0x62, 0x5f, 0xd8, 0xd6, 0x6a, 0xab, 0x54, 0x46, 0xe5, 0x02, 0x42,
	0x7c, 0xa9, 0xdc, 0xe4, 0xba, 0x45, 0x34, 0x89, 0xb1, 0x9d, 0x8a, 0xfc, 0x3d, 0x7e, 0x19, 0x8a,
	0x93, 0x45, 0x65, 0x63, 0x88, 0x6f, 0x79, 0x2f, 0xf7, 0xec, 0x77, 0xef, 0xce, 0xb0, 0x8b, 0x3f,
	0x59, 0xca, 0x57, 0x78, 0xc8, 0x17, 0x87, 0xd1, 0x8a, 0x49, 0x19, 0x72, 0x91, 0xab, 0x9c, 0x3c,
	0x65, 0x9c, 0xe7, 0x49, 0xa6, 0xca, 0xb0, 0x66, 0xd7, 0x47, 0xfe, 0xaf, 0x0e, 0x58, 0xe7, 0x15,
	0x20, 0x3b, 0x60, 0x26, 0xb1, 0x6b, 0x78, 0x46, 0x60, 0x53, 0x33, 0x89, 0x09, 0x81, 0x2e, 0x13,
	0xc8, 0x5c, 0xd3, 0x33, 0x02, 0x93, 0xea, 0x6f, 0xb2, 0x07, 0x7d, 0xa9, 0x04, 0x66, 0xd7, 0xea,
	0xc6, 0xed, 0x78, 0x46, 0x60, 0xd1, 0x16, 0x93, 0x7d, 0x80, 0x44, 0xce, 0x63, 0x5c, 0xa1, 0xc2,
	0xd8, 0xed, 0x7a, 0x46, 0xd0, 0xa7, 0x76, 0x22, 0x87, 0x35, 0x41, 0x5e, 0x43, 0x57, 0x95, 0x1c,
	0x5d, 0xcb, 0x33, 0x82, 0x9d, 0xe3, 0xe7, 0xe1, 0x3d, 0x2b, 0xa1, 0xb6, 0xf1, 0xa9, 0xe4, 0x48,
	0x75, 0x25, 0x19, 0x82, 0x93, 0x64, 0x52, 0x89, 0x22, 0x52, 0xb9, 0x90, 0x6e, 0xcf, 0xeb, 0x04,
	0xce, 0xb1, 0xff, 0x17, 0xe1, 0x0c, 0xc5, 0x3a, 0x89, 0x70, 0x2a, 0xf2, 0x75, 0x12, 0xa3, 0xa0,
	0x9b, 0x32, 0xf2, 0x02, 0x1e, 0x71, 0x14, 0xf3, 0x8a, 0x62, 0x59, 0x84, 0xee, 0x76, 0xd5, 0xe0,
	0xe5, 0x16, 0x75, 0x38, 0x8a, 0x71, 0x43, 0x92, 0x3d, 0xd8, 0x9e, 0x14, 0x29, 0x97, 0x45, 0xea,
	0xf6, 0xab, 0xb6, 0x2e, 0xb7, 0xe8, 0x2d, 0x41, 0xce, 0xa0, 0x9f, 0xa2, 0x62, 0x31, 0x53, 0xcc,
	0xb5, 0xb5, 0x87, 0x97, 0x0f, 0x99, 0x0f, 0x3f, 0x34, 0x85, 0xa3, 0x4c, 0x89, 0x92, 0xb6, 0x3a,
	0xb2, 0x0b, 0x3d, 0xce, 0x04, 0x66, 0xca, 0x05, 0x9d, 0x6f, 0x83, 0xf6, 0x4e, 0xe0, 0xf1, 0x1f,
	0x12, 0x32, 0x80, 0xce, 0x77, 0x2c, 0x9b, 0x29, 0x54, 0x9f, 0xe4, 0x19, 0x58, 0x6b, 0xb6, 0x2a,
	0x50, 0xcf, 0xc1, 0xa6, 0x35, 0x78, 0x67, 0xbe, 0x35, 0xce, 0xfa, 0xd0, 0x8b, 0x6e, 0x98, 0xb8,
	0x46, 0xff, 0x2b, 0xec, 0x9c, 0x0b, 0x64, 0x0a, 0xb5, 0x0b, 0x8a, 0x3f, 0x36, 0x2e, 0x34, 0x36,
	0x2f, 0x24, 0x21, 0x58, 0xda, 0xb2, 0x3e, 0xcd, 0x39, 0x76, 0x1f, 0xea, 0x84, 0xd6, 0x65, 0xfe,
	0x3e, 0x38, 0x17, 0xa8, 0xda, 0x63, 0xef, 0xec, 0x88, 0xff, 0x1e, 0x9e, 0xdc, 0x09, 0xff, 0xde,
	0x1a, 0xed, 0x03, 0x2c, 0x13, 0x21, 0xd5, 0x3c, 0x63, 0xe9, 0x6d, 0x13, 0xb6, 0x66, 0xae, 0x58,
	0x8a, 0xfe, 0x09, 0x38, 0xd3, 0x8d, 0x41, 0x1c, 0x40, 0x37, 0xca, 0xa5, 0xfa, 0x87, 0xbd, 0xa9,
	0x48, 0x22, 0xa4, 0xba, 0xca, 0xc7, 0x76, 0x6c, 0xad, 0xd0, 0xf8, 0x1f, 0x21, 0x39, 0x00, 0x92,
	0x15, 0xe9, 0x02, 0xc5, 0x3c, 0x5f, 0xce, 0x25, 0x4a, 0x99, 0xe4, 0x59, 0x9d, 0x89, 0x45, 0x07,
	0xf5, 0x9f, 0x8f, 0xcb, 0x59, 0xc3, 0xfb, 0x27, 0x60, 0x69, 0x71, 0x95, 0x2a, 0x4b, 0xf3, 0xa2,
	0x49, 0xd5, 0xa2, 0x0d, 0xaa, 0x9e, 0x45, 0x54, 0x08, 0x81, 0x59, 0x54, 0x36, 0x1d, 0xb6, 0xf8,
	0xd5, 0x11, 0xd8, 0xed, 0x62, 0x13, 0x07, 0xb6, 0xc7, 0x57, 0x5f, 0x4e, 0x27, 0xe3, 0xe1, 0x60,
	0xab, 0x02, 0x74, 0x74, 0xf1, 0x79, 0x72, 0x4a, 0x07, 0x06, 0x01, 0xe8, 0xcd, 0x46, 0x74, 0x3c,
	0x9a, 0x0d, 0xcc, 0xb3, 0xee, 0x37, 0x93, 0x2f, 0x16, 0x3d, 0xfd, 0x66, 0xdf, 0xfc, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x4f, 0xcc, 0x69, 0x71, 0xcd, 0x03, 0x00, 0x00,
}
