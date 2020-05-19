// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/v3/model.proto

package common

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type StellarAccountId struct {
	// The public key of accounts always starts with a G, so we
	// ensure that the value starts with a G to prevent the secret
	// key from being used.
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StellarAccountId) Reset()         { *m = StellarAccountId{} }
func (m *StellarAccountId) String() string { return proto.CompactTextString(m) }
func (*StellarAccountId) ProtoMessage()    {}
func (*StellarAccountId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{0}
}

func (m *StellarAccountId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StellarAccountId.Unmarshal(m, b)
}
func (m *StellarAccountId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StellarAccountId.Marshal(b, m, deterministic)
}
func (m *StellarAccountId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StellarAccountId.Merge(m, src)
}
func (m *StellarAccountId) XXX_Size() int {
	return xxx_messageInfo_StellarAccountId.Size(m)
}
func (m *StellarAccountId) XXX_DiscardUnknown() {
	xxx_messageInfo_StellarAccountId.DiscardUnknown(m)
}

var xxx_messageInfo_StellarAccountId proto.InternalMessageInfo

func (m *StellarAccountId) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TransactionHash struct {
	// The sha256 hash of a transaction.
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionHash) Reset()         { *m = TransactionHash{} }
func (m *TransactionHash) String() string { return proto.CompactTextString(m) }
func (*TransactionHash) ProtoMessage()    {}
func (*TransactionHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{1}
}

func (m *TransactionHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionHash.Unmarshal(m, b)
}
func (m *TransactionHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionHash.Marshal(b, m, deterministic)
}
func (m *TransactionHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionHash.Merge(m, src)
}
func (m *TransactionHash) XXX_Size() int {
	return xxx_messageInfo_TransactionHash.Size(m)
}
func (m *TransactionHash) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionHash.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionHash proto.InternalMessageInfo

func (m *TransactionHash) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type BigDecimal struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BigDecimal) Reset()         { *m = BigDecimal{} }
func (m *BigDecimal) String() string { return proto.CompactTextString(m) }
func (*BigDecimal) ProtoMessage()    {}
func (*BigDecimal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{2}
}

func (m *BigDecimal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigDecimal.Unmarshal(m, b)
}
func (m *BigDecimal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigDecimal.Marshal(b, m, deterministic)
}
func (m *BigDecimal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigDecimal.Merge(m, src)
}
func (m *BigDecimal) XXX_Size() int {
	return xxx_messageInfo_BigDecimal.Size(m)
}
func (m *BigDecimal) XXX_DiscardUnknown() {
	xxx_messageInfo_BigDecimal.DiscardUnknown(m)
}

var xxx_messageInfo_BigDecimal proto.InternalMessageInfo

func (m *BigDecimal) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type InvoiceHash struct {
	// The SHA-224 hash of the invoice.
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvoiceHash) Reset()         { *m = InvoiceHash{} }
func (m *InvoiceHash) String() string { return proto.CompactTextString(m) }
func (*InvoiceHash) ProtoMessage()    {}
func (*InvoiceHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{3}
}

func (m *InvoiceHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceHash.Unmarshal(m, b)
}
func (m *InvoiceHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceHash.Marshal(b, m, deterministic)
}
func (m *InvoiceHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceHash.Merge(m, src)
}
func (m *InvoiceHash) XXX_Size() int {
	return xxx_messageInfo_InvoiceHash.Size(m)
}
func (m *InvoiceHash) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceHash.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceHash proto.InternalMessageInfo

func (m *InvoiceHash) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Invoice struct {
	Items                []*Invoice_LineItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Invoice) Reset()         { *m = Invoice{} }
func (m *Invoice) String() string { return proto.CompactTextString(m) }
func (*Invoice) ProtoMessage()    {}
func (*Invoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{4}
}

func (m *Invoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoice.Unmarshal(m, b)
}
func (m *Invoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoice.Marshal(b, m, deterministic)
}
func (m *Invoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoice.Merge(m, src)
}
func (m *Invoice) XXX_Size() int {
	return xxx_messageInfo_Invoice.Size(m)
}
func (m *Invoice) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoice.DiscardUnknown(m)
}

var xxx_messageInfo_Invoice proto.InternalMessageInfo

func (m *Invoice) GetItems() []*Invoice_LineItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type Invoice_LineItem struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The amount in quarks.
	Amount int64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// The app SKU related to this line item, if applicable.
	Sku                  []byte   `protobuf:"bytes,4,opt,name=sku,proto3" json:"sku,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invoice_LineItem) Reset()         { *m = Invoice_LineItem{} }
func (m *Invoice_LineItem) String() string { return proto.CompactTextString(m) }
func (*Invoice_LineItem) ProtoMessage()    {}
func (*Invoice_LineItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{4, 0}
}

func (m *Invoice_LineItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoice_LineItem.Unmarshal(m, b)
}
func (m *Invoice_LineItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoice_LineItem.Marshal(b, m, deterministic)
}
func (m *Invoice_LineItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoice_LineItem.Merge(m, src)
}
func (m *Invoice_LineItem) XXX_Size() int {
	return xxx_messageInfo_Invoice_LineItem.Size(m)
}
func (m *Invoice_LineItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoice_LineItem.DiscardUnknown(m)
}

var xxx_messageInfo_Invoice_LineItem proto.InternalMessageInfo

func (m *Invoice_LineItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Invoice_LineItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Invoice_LineItem) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Invoice_LineItem) GetSku() []byte {
	if m != nil {
		return m.Sku
	}
	return nil
}

type InvoiceList struct {
	Invoices             []*Invoice `protobuf:"bytes,1,rep,name=invoices,proto3" json:"invoices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InvoiceList) Reset()         { *m = InvoiceList{} }
func (m *InvoiceList) String() string { return proto.CompactTextString(m) }
func (*InvoiceList) ProtoMessage()    {}
func (*InvoiceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{5}
}

func (m *InvoiceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceList.Unmarshal(m, b)
}
func (m *InvoiceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceList.Marshal(b, m, deterministic)
}
func (m *InvoiceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceList.Merge(m, src)
}
func (m *InvoiceList) XXX_Size() int {
	return xxx_messageInfo_InvoiceList.Size(m)
}
func (m *InvoiceList) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceList.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceList proto.InternalMessageInfo

func (m *InvoiceList) GetInvoices() []*Invoice {
	if m != nil {
		return m.Invoices
	}
	return nil
}

func init() {
	proto.RegisterType((*StellarAccountId)(nil), "kin.common.v3.StellarAccountId")
	proto.RegisterType((*TransactionHash)(nil), "kin.common.v3.TransactionHash")
	proto.RegisterType((*BigDecimal)(nil), "kin.common.v3.BigDecimal")
	proto.RegisterType((*InvoiceHash)(nil), "kin.common.v3.InvoiceHash")
	proto.RegisterType((*Invoice)(nil), "kin.common.v3.Invoice")
	proto.RegisterType((*Invoice_LineItem)(nil), "kin.common.v3.Invoice.LineItem")
	proto.RegisterType((*InvoiceList)(nil), "kin.common.v3.InvoiceList")
}

func init() { proto.RegisterFile("common/v3/model.proto", fileDescriptor_ab58ade4bb87a1ff) }

var fileDescriptor_ab58ade4bb87a1ff = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xed, 0xc6, 0x96, 0xa3, 0xac, 0x5b, 0x22, 0x16, 0x9c, 0x08, 0x13, 0x88, 0x30, 0x3d, 0x28,
	0x25, 0x92, 0xfc, 0x51, 0x4a, 0x92, 0x42, 0x4d, 0x96, 0x40, 0x6b, 0x70, 0x2f, 0x6a, 0x4f, 0x76,
	0x13, 0xd8, 0x48, 0x8b, 0xb2, 0x58, 0xda, 0x35, 0xd2, 0x5a, 0x50, 0xb7, 0x01, 0xff, 0x82, 0x1e,
	0xf2, 0x03, 0xfb, 0x43, 0x74, 0x2a, 0x92, 0x2c, 0xe1, 0x40, 0xc9, 0x6d, 0xe6, 0xcd, 0x7b, 0x3b,
	0xf3, 0x66, 0x07, 0x76, 0x3c, 0x11, 0x45, 0x82, 0x3b, 0xe9, 0xc8, 0x89, 0x84, 0x4f, 0x43, 0x7b,
	0x19, 0x0b, 0x29, 0xd0, 0x9b, 0x05, 0xe3, 0x76, 0x59, 0xb2, 0xd3, 0x51, 0xf7, 0x38, 0x25, 0x21,
	0xf3, 0x89, 0xa4, 0x4e, 0x15, 0x94, 0xbc, 0xde, 0x0d, 0xd4, 0xbe, 0x49, 0x1a, 0x86, 0x24, 0xbe,
	0xf6, 0x3c, 0xb1, 0xe2, 0x72, 0xe2, 0xa3, 0x3e, 0x54, 0x52, 0x12, 0xae, 0xa8, 0x0e, 0x0c, 0x60,
	0x1e, 0xe0, 0x6e, 0x86, 0x8f, 0xe3, 0x8e, 0x71, 0x61, 0x5e, 0x0c, 0x0f, 0xef, 0x3e, 0xcf, 0x89,
	0xb5, 0xbe, 0xb6, 0x66, 0x7d, 0xeb, 0xf2, 0xf6, 0xdd, 0x5b, 0xb7, 0x24, 0xf6, 0x86, 0xf0, 0xf0,
	0x7b, 0x4c, 0x78, 0x42, 0x3c, 0xc9, 0x04, 0xff, 0x42, 0x92, 0x07, 0x74, 0xba, 0xfb, 0xc8, 0x6b,
	0x7c, 0x90, 0xe1, 0xd6, 0xba, 0xa9, 0x19, 0xba, 0x51, 0x69, 0x66, 0x10, 0x62, 0x16, 0xdc, 0x50,
	0x8f, 0x45, 0x24, 0x44, 0xd3, 0xe7, 0x3d, 0x3f, 0x64, 0x78, 0x14, 0x0f, 0x86, 0xce, 0x9d, 0x39,
	0xbe, 0x9a, 0x0f, 0xac, 0xcb, 0xdb, 0x79, 0xde, 0xf0, 0x57, 0xff, 0x7c, 0xf0, 0xfe, 0xf1, 0x77,
	0xff, 0xcc, 0x1c, 0x5f, 0xfd, 0xb0, 0x6b, 0x68, 0xf0, 0x58, 0x50, 0xce, 0xc6, 0xf5, 0x3c, 0x36,
	0x6c, 0x4f, 0x78, 0x2a, 0x98, 0x47, 0x5f, 0x98, 0xe5, 0x44, 0x3f, 0xa9, 0xf8, 0x7f, 0x01, 0xdc,
	0xdf, 0x0a, 0x10, 0x86, 0x0a, 0x93, 0x34, 0x4a, 0x74, 0x60, 0x34, 0xcc, 0xf6, 0xf0, 0xd4, 0x7e,
	0xb6, 0x49, 0x7b, 0x4b, 0xb3, 0xa7, 0x8c, 0xd3, 0x89, 0xa4, 0x11, 0x6e, 0x67, 0x58, 0x7d, 0x02,
	0x8a, 0x0a, 0xb4, 0x8d, 0xea, 0x96, 0xd2, 0xee, 0x1f, 0x00, 0xd5, 0x8a, 0x80, 0x0c, 0xa8, 0x48,
	0x26, 0xc3, 0xca, 0x1a, 0xcc, 0xf0, 0x7e, 0xac, 0x68, 0x40, 0xdf, 0x00, 0xb7, 0x2c, 0xa0, 0x73,
	0xd8, 0xf6, 0x69, 0xe2, 0xc5, 0x6c, 0x99, 0xaf, 0x4f, 0xdf, 0xdb, 0xe5, 0xbd, 0xd2, 0x37, 0x7b,
	0xee, 0x6e, 0x19, 0x1d, 0xc1, 0x16, 0x89, 0xf2, 0xaf, 0xd2, 0x1b, 0x06, 0x30, 0x1b, 0xee, 0x36,
	0x43, 0x5d, 0xd8, 0x48, 0x16, 0x2b, 0xbd, 0x59, 0x78, 0x54, 0x33, 0xac, 0xac, 0x1b, 0x79, 0x8f,
	0x1c, 0xec, 0x7d, 0xad, 0x17, 0x32, 0x65, 0x89, 0x44, 0x9f, 0xa0, 0xca, 0xca, 0xb4, 0xb2, 0x79,
	0xf4, 0x7f, 0x9b, 0xc5, 0x14, 0x4f, 0xa0, 0xa9, 0x02, 0xcd, 0x77, 0x6b, 0x0d, 0x9e, 0xc3, 0x8e,
	0x88, 0x83, 0x42, 0x12, 0xd0, 0x1d, 0xd9, 0x0c, 0x07, 0x4c, 0x3e, 0xac, 0xee, 0x73, 0xc4, 0x59,
	0x30, 0x4e, 0x3d, 0x91, 0xfc, 0x4c, 0x24, 0x2d, 0x12, 0x8b, 0x2c, 0x99, 0xc5, 0xb8, 0xa4, 0x31,
	0x27, 0xa1, 0x13, 0x50, 0x5e, 0x5c, 0xa0, 0x53, 0xdf, 0xef, 0xc7, 0x32, 0xba, 0x6f, 0x15, 0xf8,
	0xe8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x78, 0xa1, 0x4c, 0xda, 0x02, 0x00, 0x00,
}
