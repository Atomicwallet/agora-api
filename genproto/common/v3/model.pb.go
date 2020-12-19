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

type InvoiceError_Reason int32

const (
	InvoiceError_UNKNOWN InvoiceError_Reason = 0
	// The provided invoice has already been paid for.
	//
	// This is only applicable when the memo transaction type
	// is SPEND.
	InvoiceError_ALREADY_PAID InvoiceError_Reason = 1
	// The destination in the operation corresponding to this invoice
	// is incorrect.
	InvoiceError_WRONG_DESTINATION InvoiceError_Reason = 2
	// One or more SKUs in the invoice was not found.
	InvoiceError_SKU_NOT_FOUND InvoiceError_Reason = 3
)

var InvoiceError_Reason_name = map[int32]string{
	0: "UNKNOWN",
	1: "ALREADY_PAID",
	2: "WRONG_DESTINATION",
	3: "SKU_NOT_FOUND",
}

var InvoiceError_Reason_value = map[string]int32{
	"UNKNOWN":           0,
	"ALREADY_PAID":      1,
	"WRONG_DESTINATION": 2,
	"SKU_NOT_FOUND":     3,
}

func (x InvoiceError_Reason) String() string {
	return proto.EnumName(InvoiceError_Reason_name, int32(x))
}

func (InvoiceError_Reason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{5, 0}
}

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
	return fileDescriptor_ab58ade4bb87a1ff, []int{2}
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
	return fileDescriptor_ab58ade4bb87a1ff, []int{3}
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
	return fileDescriptor_ab58ade4bb87a1ff, []int{3, 0}
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
	return fileDescriptor_ab58ade4bb87a1ff, []int{4}
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

type InvoiceError struct {
	// The operation index the invoice corresponds to.
	OpIndex uint32 `protobuf:"varint,1,opt,name=op_index,json=opIndex,proto3" json:"op_index,omitempty"`
	// The invoice that was submitted.
	Invoice              *Invoice            `protobuf:"bytes,2,opt,name=invoice,proto3" json:"invoice,omitempty"`
	Reason               InvoiceError_Reason `protobuf:"varint,3,opt,name=reason,proto3,enum=kin.agora.common.v3.InvoiceError_Reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *InvoiceError) Reset()         { *m = InvoiceError{} }
func (m *InvoiceError) String() string { return proto.CompactTextString(m) }
func (*InvoiceError) ProtoMessage()    {}
func (*InvoiceError) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{5}
}

func (m *InvoiceError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceError.Unmarshal(m, b)
}
func (m *InvoiceError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceError.Marshal(b, m, deterministic)
}
func (m *InvoiceError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceError.Merge(m, src)
}
func (m *InvoiceError) XXX_Size() int {
	return xxx_messageInfo_InvoiceError.Size(m)
}
func (m *InvoiceError) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceError.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceError proto.InternalMessageInfo

func (m *InvoiceError) GetOpIndex() uint32 {
	if m != nil {
		return m.OpIndex
	}
	return 0
}

func (m *InvoiceError) GetInvoice() *Invoice {
	if m != nil {
		return m.Invoice
	}
	return nil
}

func (m *InvoiceError) GetReason() InvoiceError_Reason {
	if m != nil {
		return m.Reason
	}
	return InvoiceError_UNKNOWN
}

func init() {
	proto.RegisterEnum("kin.agora.common.v3.InvoiceError_Reason", InvoiceError_Reason_name, InvoiceError_Reason_value)
	proto.RegisterType((*StellarAccountId)(nil), "kin.agora.common.v3.StellarAccountId")
	proto.RegisterType((*TransactionHash)(nil), "kin.agora.common.v3.TransactionHash")
	proto.RegisterType((*InvoiceHash)(nil), "kin.agora.common.v3.InvoiceHash")
	proto.RegisterType((*Invoice)(nil), "kin.agora.common.v3.Invoice")
	proto.RegisterType((*Invoice_LineItem)(nil), "kin.agora.common.v3.Invoice.LineItem")
	proto.RegisterType((*InvoiceList)(nil), "kin.agora.common.v3.InvoiceList")
	proto.RegisterType((*InvoiceError)(nil), "kin.agora.common.v3.InvoiceError")
}

func init() { proto.RegisterFile("common/v3/model.proto", fileDescriptor_ab58ade4bb87a1ff) }

var fileDescriptor_ab58ade4bb87a1ff = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x4f, 0xdb, 0x30,
	0x18, 0xc6, 0x71, 0xfa, 0x2f, 0xb8, 0x30, 0x82, 0x27, 0x44, 0xd4, 0x21, 0xad, 0xaa, 0x36, 0xa9,
	0x42, 0x6b, 0x82, 0xda, 0x0b, 0xd3, 0x2e, 0x24, 0x0b, 0x83, 0x08, 0x94, 0xb2, 0xb4, 0x0c, 0x8d,
	0x69, 0xab, 0x4c, 0x62, 0x15, 0x8b, 0xc4, 0xee, 0x1c, 0xb7, 0xda, 0x38, 0x71, 0xdd, 0x8e, 0x7c,
	0x84, 0x7d, 0xba, 0x7d, 0x86, 0x9e, 0xa6, 0x24, 0x0d, 0xeb, 0x61, 0x62, 0xb7, 0xd7, 0x8f, 0x7f,
	0xcf, 0xeb, 0xf7, 0x79, 0x25, 0xc3, 0xad, 0x80, 0xc7, 0x31, 0x67, 0xe6, 0xac, 0x67, 0xc6, 0x3c,
	0x24, 0x91, 0x31, 0x11, 0x5c, 0x72, 0xf4, 0xf4, 0x86, 0x32, 0x03, 0x8f, 0xb9, 0xc0, 0x46, 0x0e,
	0x18, 0xb3, 0x5e, 0x63, 0x7b, 0x86, 0x23, 0x1a, 0x62, 0x49, 0xcc, 0xa2, 0xc8, 0xe9, 0x96, 0x03,
	0xb5, 0x81, 0x24, 0x51, 0x84, 0x85, 0x15, 0x04, 0x7c, 0xca, 0xa4, 0x1b, 0xa2, 0x3d, 0x58, 0x99,
	0xe1, 0x68, 0x4a, 0x74, 0xd0, 0x04, 0xed, 0x55, 0xbb, 0x31, 0xb7, 0xb7, 0xc5, 0x56, 0x73, 0xbf,
	0xbd, 0xdf, 0xdd, 0xf8, 0x72, 0xf4, 0x09, 0x77, 0x6e, 0xad, 0xce, 0xe5, 0x5e, 0xe7, 0xf5, 0xe7,
	0xdd, 0x17, 0x7e, 0x0e, 0xb6, 0xba, 0x70, 0x63, 0x28, 0x30, 0x4b, 0x70, 0x20, 0x29, 0x67, 0xc7,
	0x38, 0xb9, 0x46, 0xcf, 0x97, 0x9b, 0xac, 0xd9, 0xab, 0x73, 0xbb, 0x7a, 0x5b, 0xd6, 0x9a, 0x7a,
	0xb3, 0xf0, 0x18, 0xb0, 0xee, 0xb2, 0x19, 0xa7, 0x01, 0x79, 0x84, 0xdf, 0xd1, 0x77, 0x0a, 0xfe,
	0x37, 0x80, 0xb5, 0x85, 0x01, 0x1d, 0xc3, 0x0a, 0x95, 0x24, 0x4e, 0x74, 0xd0, 0x2c, 0xb5, 0xeb,
	0xdd, 0x97, 0xc6, 0x3f, 0x32, 0x1b, 0x0b, 0xd8, 0x38, 0xa5, 0x8c, 0xb8, 0x92, 0xc4, 0x76, 0x7d,
	0x6e, 0xab, 0xf7, 0xa0, 0xa2, 0x02, 0xed, 0x4e, 0xf5, 0xf3, 0x06, 0x8d, 0x7b, 0x00, 0xd5, 0x02,
	0x40, 0x4d, 0x58, 0x91, 0x54, 0x46, 0x45, 0x70, 0x38, 0xb7, 0x6b, 0xa2, 0xa2, 0x01, 0xfd, 0x0e,
	0xf8, 0xf9, 0x05, 0x7a, 0x05, 0xeb, 0x21, 0x49, 0x02, 0x41, 0x27, 0x69, 0x50, 0x5d, 0x59, 0xe6,
	0x56, 0xf4, 0x3b, 0xc5, 0x5f, 0xbe, 0x46, 0x0d, 0x58, 0xc5, 0x71, 0xba, 0x54, 0xbd, 0xd4, 0x04,
	0xed, 0x92, 0xad, 0xec, 0x01, 0x7f, 0xa1, 0xa0, 0x06, 0x2c, 0x25, 0x37, 0x53, 0xbd, 0x9c, 0xa5,
	0x55, 0xe7, 0x76, 0xe5, 0xb6, 0x94, 0xbe, 0x93, 0x8a, 0xad, 0xc1, 0xc3, 0x6a, 0x4e, 0x69, 0x22,
	0x91, 0x03, 0x55, 0x9a, 0x1f, 0x8b, 0xc0, 0x3b, 0x8f, 0x05, 0xce, 0xe6, 0xb9, 0x07, 0x65, 0x15,
	0x68, 0xa1, 0xff, 0xe0, 0x6c, 0xfd, 0x50, 0xe0, 0xda, 0x82, 0x38, 0x14, 0x82, 0x0b, 0xd4, 0x82,
	0x2a, 0x9f, 0x8c, 0x28, 0x0b, 0xc9, 0xb7, 0x2c, 0xf0, 0xba, 0x5d, 0x9b, 0xdb, 0xe5, 0x5d, 0x45,
	0x0f, 0xfd, 0x1a, 0x9f, 0xb8, 0xa9, 0x8e, 0x0e, 0x60, 0x6d, 0xd1, 0x20, 0xcb, 0xfa, 0xbf, 0x97,
	0xd3, 0x1c, 0x3f, 0x81, 0xa2, 0x01, 0xbf, 0xb0, 0xa1, 0x03, 0x58, 0x15, 0x04, 0x27, 0x9c, 0x65,
	0x3b, 0x78, 0xd2, 0x6d, 0x3f, 0xd6, 0x20, 0x1b, 0xcc, 0xf0, 0x33, 0xde, 0x5f, 0xf8, 0x5a, 0xef,
	0x61, 0x35, 0x57, 0x50, 0x1d, 0xd6, 0xce, 0xbd, 0x13, 0xaf, 0x7f, 0xe1, 0x69, 0x2b, 0x48, 0x83,
	0x6b, 0xd6, 0xa9, 0x7f, 0x68, 0x39, 0x1f, 0x47, 0x67, 0x96, 0xeb, 0x68, 0x00, 0x6d, 0xc1, 0xcd,
	0x0b, 0xbf, 0xef, 0x1d, 0x8d, 0x9c, 0xc3, 0xc1, 0xd0, 0xf5, 0xac, 0xa1, 0xdb, 0xf7, 0x34, 0x05,
	0x6d, 0xc2, 0xf5, 0xc1, 0xc9, 0xf9, 0xc8, 0xeb, 0x0f, 0x47, 0xef, 0xfa, 0xe7, 0x9e, 0xa3, 0x95,
	0xec, 0xaf, 0xf0, 0x19, 0x17, 0xe3, 0xa5, 0x49, 0xc6, 0x84, 0xfd, 0x9d, 0xe6, 0xd2, 0x19, 0x53,
	0x79, 0x3d, 0xbd, 0x4a, 0x15, 0xf3, 0x86, 0x32, 0x12, 0xf0, 0xe4, 0x7b, 0x22, 0x49, 0x6c, 0x66,
	0x70, 0x07, 0x4f, 0x68, 0x87, 0x32, 0x49, 0x04, 0xc3, 0x91, 0x39, 0x26, 0x2c, 0xfb, 0x49, 0xe6,
	0xc3, 0x6f, 0x7c, 0x93, 0x57, 0xbf, 0x94, 0xba, 0x75, 0x66, 0xbf, 0xcd, 0xea, 0x0f, 0xbd, 0xab,
	0x6a, 0x46, 0xf5, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xaa, 0x05, 0x0d, 0xb6, 0x03, 0x00,
	0x00,
}
