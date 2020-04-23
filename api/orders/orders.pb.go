// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orders.proto

package orders

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

type Order struct {
	AcceptLossOfReturnRight bool             `protobuf:"varint,1,opt,name=AcceptLossOfReturnRight,proto3" json:"AcceptLossOfReturnRight,omitempty"`
	Id                      string           `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	OrderNumber             string           `protobuf:"bytes,3,opt,name=OrderNumber,proto3" json:"OrderNumber,omitempty"`
	State                   int32            `protobuf:"varint,4,opt,name=State,proto3" json:"State,omitempty"`
	VatMode                 int32            `protobuf:"varint,5,opt,name=VatMode,proto3" json:"VatMode,omitempty"`
	CreatedAt               string           `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	ConfirmedAt             string           `protobuf:"bytes,7,opt,name=ConfirmedAt,proto3" json:"ConfirmedAt,omitempty"`
	ShippedAt               string           `protobuf:"bytes,8,opt,name=ShippedAt,proto3" json:"ShippedAt,omitempty"`
	PayedAt                 string           `protobuf:"bytes,9,opt,name=PayedAt,proto3" json:"PayedAt,omitempty"`
	SellerComment           string           `protobuf:"bytes,10,opt,name=SellerComment,proto3" json:"SellerComment,omitempty"`
	InvoiceAddress          *InvoiceAddress  `protobuf:"bytes,11,opt,name=invoiceAddress,proto3" json:"invoiceAddress,omitempty"`
	ShippingAddress         *ShippingAddress `protobuf:"bytes,12,opt,name=shippingAddress,proto3" json:"shippingAddress,omitempty"`
	ShippingCost            float32          `protobuf:"fixed32,13,opt,name=ShippingCost,proto3" json:"ShippingCost,omitempty"`
	PaymentMethod           int32            `protobuf:"varint,14,opt,name=PaymentMethod,proto3" json:"PaymentMethod,omitempty"`
	OrderItems              []*OrderItem     `protobuf:"bytes,15,rep,name=orderItems,proto3" json:"orderItems,omitempty"`
	Customer                *Customer        `protobuf:"bytes,16,opt,name=Customer,proto3" json:"Customer,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}         `json:"-"`
	XXX_unrecognized        []byte           `json:"-"`
	XXX_sizecache           int32            `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetAcceptLossOfReturnRight() bool {
	if m != nil {
		return m.AcceptLossOfReturnRight
	}
	return false
}

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetOrderNumber() string {
	if m != nil {
		return m.OrderNumber
	}
	return ""
}

func (m *Order) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *Order) GetVatMode() int32 {
	if m != nil {
		return m.VatMode
	}
	return 0
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Order) GetConfirmedAt() string {
	if m != nil {
		return m.ConfirmedAt
	}
	return ""
}

func (m *Order) GetShippedAt() string {
	if m != nil {
		return m.ShippedAt
	}
	return ""
}

func (m *Order) GetPayedAt() string {
	if m != nil {
		return m.PayedAt
	}
	return ""
}

func (m *Order) GetSellerComment() string {
	if m != nil {
		return m.SellerComment
	}
	return ""
}

func (m *Order) GetInvoiceAddress() *InvoiceAddress {
	if m != nil {
		return m.InvoiceAddress
	}
	return nil
}

func (m *Order) GetShippingAddress() *ShippingAddress {
	if m != nil {
		return m.ShippingAddress
	}
	return nil
}

func (m *Order) GetShippingCost() float32 {
	if m != nil {
		return m.ShippingCost
	}
	return 0
}

func (m *Order) GetPaymentMethod() int32 {
	if m != nil {
		return m.PaymentMethod
	}
	return 0
}

func (m *Order) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

func (m *Order) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type ShippingAddress struct {
	Company              string   `protobuf:"bytes,1,opt,name=Company,proto3" json:"Company,omitempty"`
	Street               string   `protobuf:"bytes,2,opt,name=Street,proto3" json:"Street,omitempty"`
	HouseNumber          string   `protobuf:"bytes,3,opt,name=HouseNumber,proto3" json:"HouseNumber,omitempty"`
	Line2                string   `protobuf:"bytes,4,opt,name=Line2,proto3" json:"Line2,omitempty"`
	Line3                string   `protobuf:"bytes,5,opt,name=Line3,proto3" json:"Line3,omitempty"`
	City                 string   `protobuf:"bytes,6,opt,name=City,proto3" json:"City,omitempty"`
	Zip                  string   `protobuf:"bytes,7,opt,name=Zip,proto3" json:"Zip,omitempty"`
	State                string   `protobuf:"bytes,8,opt,name=State,proto3" json:"State,omitempty"`
	Country              string   `protobuf:"bytes,9,opt,name=Country,proto3" json:"Country,omitempty"`
	CountryISO2          string   `protobuf:"bytes,10,opt,name=CountryISO2,proto3" json:"CountryISO2,omitempty"`
	FirstName            string   `protobuf:"bytes,11,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,12,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Phone                string   `protobuf:"bytes,13,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Email                string   `protobuf:"bytes,14,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShippingAddress) Reset()         { *m = ShippingAddress{} }
func (m *ShippingAddress) String() string { return proto.CompactTextString(m) }
func (*ShippingAddress) ProtoMessage()    {}
func (*ShippingAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{1}
}

func (m *ShippingAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShippingAddress.Unmarshal(m, b)
}
func (m *ShippingAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShippingAddress.Marshal(b, m, deterministic)
}
func (m *ShippingAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShippingAddress.Merge(m, src)
}
func (m *ShippingAddress) XXX_Size() int {
	return xxx_messageInfo_ShippingAddress.Size(m)
}
func (m *ShippingAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ShippingAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ShippingAddress proto.InternalMessageInfo

func (m *ShippingAddress) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *ShippingAddress) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *ShippingAddress) GetHouseNumber() string {
	if m != nil {
		return m.HouseNumber
	}
	return ""
}

func (m *ShippingAddress) GetLine2() string {
	if m != nil {
		return m.Line2
	}
	return ""
}

func (m *ShippingAddress) GetLine3() string {
	if m != nil {
		return m.Line3
	}
	return ""
}

func (m *ShippingAddress) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *ShippingAddress) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *ShippingAddress) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ShippingAddress) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *ShippingAddress) GetCountryISO2() string {
	if m != nil {
		return m.CountryISO2
	}
	return ""
}

func (m *ShippingAddress) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *ShippingAddress) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *ShippingAddress) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ShippingAddress) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type InvoiceAddress struct {
	Company              string   `protobuf:"bytes,1,opt,name=Company,proto3" json:"Company,omitempty"`
	Street               string   `protobuf:"bytes,2,opt,name=Street,proto3" json:"Street,omitempty"`
	HouseNumber          string   `protobuf:"bytes,3,opt,name=HouseNumber,proto3" json:"HouseNumber,omitempty"`
	Line2                string   `protobuf:"bytes,4,opt,name=Line2,proto3" json:"Line2,omitempty"`
	Line3                string   `protobuf:"bytes,5,opt,name=Line3,proto3" json:"Line3,omitempty"`
	City                 string   `protobuf:"bytes,6,opt,name=City,proto3" json:"City,omitempty"`
	Zip                  string   `protobuf:"bytes,7,opt,name=Zip,proto3" json:"Zip,omitempty"`
	State                string   `protobuf:"bytes,8,opt,name=State,proto3" json:"State,omitempty"`
	Country              string   `protobuf:"bytes,9,opt,name=Country,proto3" json:"Country,omitempty"`
	CountryISO2          string   `protobuf:"bytes,10,opt,name=CountryISO2,proto3" json:"CountryISO2,omitempty"`
	FirstName            string   `protobuf:"bytes,11,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,12,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Phone                string   `protobuf:"bytes,13,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Email                string   `protobuf:"bytes,14,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvoiceAddress) Reset()         { *m = InvoiceAddress{} }
func (m *InvoiceAddress) String() string { return proto.CompactTextString(m) }
func (*InvoiceAddress) ProtoMessage()    {}
func (*InvoiceAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{2}
}

func (m *InvoiceAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceAddress.Unmarshal(m, b)
}
func (m *InvoiceAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceAddress.Marshal(b, m, deterministic)
}
func (m *InvoiceAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceAddress.Merge(m, src)
}
func (m *InvoiceAddress) XXX_Size() int {
	return xxx_messageInfo_InvoiceAddress.Size(m)
}
func (m *InvoiceAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceAddress.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceAddress proto.InternalMessageInfo

func (m *InvoiceAddress) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *InvoiceAddress) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *InvoiceAddress) GetHouseNumber() string {
	if m != nil {
		return m.HouseNumber
	}
	return ""
}

func (m *InvoiceAddress) GetLine2() string {
	if m != nil {
		return m.Line2
	}
	return ""
}

func (m *InvoiceAddress) GetLine3() string {
	if m != nil {
		return m.Line3
	}
	return ""
}

func (m *InvoiceAddress) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *InvoiceAddress) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *InvoiceAddress) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *InvoiceAddress) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *InvoiceAddress) GetCountryISO2() string {
	if m != nil {
		return m.CountryISO2
	}
	return ""
}

func (m *InvoiceAddress) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *InvoiceAddress) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *InvoiceAddress) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *InvoiceAddress) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Shipment struct {
	ShippingId                string   `protobuf:"bytes,1,opt,name=ShippingId,proto3" json:"ShippingId,omitempty"`
	OrderId                   string   `protobuf:"bytes,2,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	Comment                   string   `protobuf:"bytes,3,opt,name=Comment,proto3" json:"Comment,omitempty"`
	ShippingProviderId        int32    `protobuf:"varint,4,opt,name=ShippingProviderId,proto3" json:"ShippingProviderId,omitempty"`
	ShippingProviderProductId int32    `protobuf:"varint,5,opt,name=ShippingProviderProductId,proto3" json:"ShippingProviderProductId,omitempty"`
	CarrierId                 int32    `protobuf:"varint,6,opt,name=CarrierId,proto3" json:"CarrierId,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Shipment) Reset()         { *m = Shipment{} }
func (m *Shipment) String() string { return proto.CompactTextString(m) }
func (*Shipment) ProtoMessage()    {}
func (*Shipment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{3}
}

func (m *Shipment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shipment.Unmarshal(m, b)
}
func (m *Shipment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shipment.Marshal(b, m, deterministic)
}
func (m *Shipment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shipment.Merge(m, src)
}
func (m *Shipment) XXX_Size() int {
	return xxx_messageInfo_Shipment.Size(m)
}
func (m *Shipment) XXX_DiscardUnknown() {
	xxx_messageInfo_Shipment.DiscardUnknown(m)
}

var xxx_messageInfo_Shipment proto.InternalMessageInfo

func (m *Shipment) GetShippingId() string {
	if m != nil {
		return m.ShippingId
	}
	return ""
}

func (m *Shipment) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Shipment) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Shipment) GetShippingProviderId() int32 {
	if m != nil {
		return m.ShippingProviderId
	}
	return 0
}

func (m *Shipment) GetShippingProviderProductId() int32 {
	if m != nil {
		return m.ShippingProviderProductId
	}
	return 0
}

func (m *Shipment) GetCarrierId() int32 {
	if m != nil {
		return m.CarrierId
	}
	return 0
}

type OrderItem struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Quantity             float32  `protobuf:"fixed32,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	TotalPrice           float32  `protobuf:"fixed32,3,opt,name=TotalPrice,proto3" json:"TotalPrice,omitempty"`
	Discount             float32  `protobuf:"fixed32,4,opt,name=Discount,proto3" json:"Discount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderItem) Reset()         { *m = OrderItem{} }
func (m *OrderItem) String() string { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()    {}
func (*OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{4}
}

func (m *OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderItem.Unmarshal(m, b)
}
func (m *OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderItem.Marshal(b, m, deterministic)
}
func (m *OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderItem.Merge(m, src)
}
func (m *OrderItem) XXX_Size() int {
	return xxx_messageInfo_OrderItem.Size(m)
}
func (m *OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderItem proto.InternalMessageInfo

func (m *OrderItem) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *OrderItem) GetQuantity() float32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *OrderItem) GetTotalPrice() float32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *OrderItem) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

type Customer struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Number               int32    `protobuf:"varint,2,opt,name=Number,proto3" json:"Number,omitempty"`
	Type                 int32    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Tel1                 string   `protobuf:"bytes,6,opt,name=Tel1,proto3" json:"Tel1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{5}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Customer) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Customer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Customer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Customer) GetTel1() string {
	if m != nil {
		return m.Tel1
	}
	return ""
}

type Product struct {
	SKU                  string   `protobuf:"bytes,1,opt,name=SKU,proto3" json:"SKU,omitempty"`
	EAN                  string   `protobuf:"bytes,2,opt,name=EAN,proto3" json:"EAN,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{6}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetSKU() string {
	if m != nil {
		return m.SKU
	}
	return ""
}

func (m *Product) GetEAN() string {
	if m != nil {
		return m.EAN
	}
	return ""
}

func (m *Product) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*Order)(nil), "orders.Order")
	proto.RegisterType((*ShippingAddress)(nil), "orders.ShippingAddress")
	proto.RegisterType((*InvoiceAddress)(nil), "orders.InvoiceAddress")
	proto.RegisterType((*Shipment)(nil), "orders.Shipment")
	proto.RegisterType((*OrderItem)(nil), "orders.OrderItem")
	proto.RegisterType((*Customer)(nil), "orders.Customer")
	proto.RegisterType((*Product)(nil), "orders.Product")
}

func init() {
	proto.RegisterFile("orders.proto", fileDescriptor_e0f5d4cf0fc9e41b)
}

var fileDescriptor_e0f5d4cf0fc9e41b = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9d, 0x4f, 0x6f, 0xd2, 0xa4, 0xac, 0x50, 0x6b, 0x10, 0x42, 0x51, 0xc4, 0x21, 0x48,
	0xa8, 0x52, 0xd3, 0x0b, 0x07, 0x84, 0x14, 0x99, 0x22, 0x22, 0xda, 0x26, 0x6c, 0x0a, 0x07, 0x6e,
	0x6e, 0xbc, 0x6d, 0x56, 0x8a, 0xbd, 0xd6, 0x7a, 0x53, 0x29, 0x77, 0xce, 0xfc, 0x50, 0x0e, 0xf4,
	0x2f, 0xa0, 0x99, 0xdd, 0x4d, 0x9c, 0x88, 0x1e, 0xb9, 0x71, 0x9b, 0xf7, 0x66, 0x76, 0x76, 0x67,
	0xe6, 0x79, 0x4c, 0xda, 0x52, 0x25, 0x5c, 0x15, 0x27, 0xb9, 0x92, 0x5a, 0xd2, 0xba, 0x41, 0xfd,
	0x87, 0x2a, 0xa9, 0x4d, 0xc0, 0xa4, 0x6f, 0xc9, 0xf1, 0x68, 0x3e, 0xe7, 0xb9, 0xbe, 0x90, 0x45,
	0x31, 0xb9, 0x65, 0x5c, 0xaf, 0x54, 0xc6, 0xc4, 0xdd, 0x42, 0x87, 0x5e, 0xcf, 0x1b, 0x34, 0xd9,
	0x63, 0x6e, 0xda, 0x21, 0xfe, 0x38, 0x09, 0xfd, 0x9e, 0x37, 0x08, 0x98, 0x3f, 0x4e, 0x68, 0x8f,
	0xb4, 0x30, 0xe5, 0xd5, 0x2a, 0xbd, 0xe1, 0x2a, 0xac, 0xa0, 0xa3, 0x4c, 0xd1, 0xa7, 0xa4, 0x36,
	0xd3, 0xb1, 0xe6, 0x61, 0xb5, 0xe7, 0x0d, 0x6a, 0xcc, 0x00, 0x1a, 0x92, 0xc6, 0xb7, 0x58, 0x5f,
	0xca, 0x84, 0x87, 0x35, 0xe4, 0x1d, 0xa4, 0x2f, 0x48, 0x10, 0x29, 0x1e, 0x6b, 0x9e, 0x8c, 0x74,
	0x58, 0xc7, 0x7c, 0x5b, 0x02, 0xee, 0x8b, 0x64, 0x76, 0x2b, 0x54, 0x8a, 0xfe, 0x86, 0xb9, 0xaf,
	0x44, 0xc1, 0xf9, 0xd9, 0x42, 0xe4, 0x39, 0xfa, 0x9b, 0xe6, 0xfc, 0x86, 0x80, 0x7b, 0xa7, 0xf1,
	0x1a, 0x7d, 0x01, 0xfa, 0x1c, 0xa4, 0xaf, 0xc8, 0xc1, 0x8c, 0x2f, 0x97, 0x5c, 0x45, 0x32, 0x4d,
	0x79, 0xa6, 0x43, 0x82, 0xfe, 0x5d, 0x92, 0xbe, 0x27, 0x1d, 0x91, 0xdd, 0x4b, 0x31, 0xe7, 0xa3,
	0x24, 0x51, 0xbc, 0x28, 0xc2, 0x56, 0xcf, 0x1b, 0xb4, 0x86, 0x47, 0x27, 0xb6, 0xe5, 0xe3, 0x1d,
	0x2f, 0xdb, 0x8b, 0xa6, 0x23, 0xd2, 0x2d, 0xe0, 0x31, 0x22, 0xbb, 0x73, 0x09, 0xda, 0x98, 0xe0,
	0xd8, 0x25, 0x98, 0xed, 0xba, 0xd9, 0x7e, 0x3c, 0xed, 0x93, 0xb6, 0x8b, 0x89, 0x64, 0xa1, 0xc3,
	0x83, 0x9e, 0x37, 0xf0, 0xd9, 0x0e, 0x07, 0xc5, 0x4c, 0xe3, 0x35, 0xbc, 0xf8, 0x92, 0xeb, 0x85,
	0x4c, 0xc2, 0x0e, 0x36, 0x79, 0x97, 0xa4, 0xa7, 0x84, 0xe0, 0xa5, 0x63, 0xcd, 0xd3, 0x22, 0xec,
	0xf6, 0x2a, 0x83, 0xd6, 0xf0, 0x89, 0x7b, 0xc7, 0xc4, 0x79, 0x58, 0x29, 0x88, 0xbe, 0x21, 0xcd,
	0x68, 0x55, 0x68, 0x99, 0x72, 0x15, 0x1e, 0xe2, 0xc3, 0x0f, 0xdd, 0x01, 0xc7, 0xb3, 0x4d, 0x44,
	0xff, 0xc1, 0x27, 0xdd, 0xbd, 0x7a, 0x60, 0x02, 0x91, 0x4c, 0xf3, 0x38, 0x5b, 0xa3, 0xd6, 0x02,
	0xe6, 0x20, 0x3d, 0x22, 0xf5, 0x99, 0x56, 0x9c, 0x6b, 0xab, 0x2f, 0x8b, 0x60, 0xe6, 0x9f, 0xe4,
	0xaa, 0xe0, 0xbb, 0x1a, 0x2b, 0x51, 0xa0, 0xb1, 0x0b, 0x91, 0xf1, 0x21, 0x6a, 0x2c, 0x60, 0x06,
	0x38, 0xf6, 0x0c, 0x15, 0x66, 0xd9, 0x33, 0x4a, 0x49, 0x35, 0x12, 0x7a, 0x6d, 0xa5, 0x85, 0x36,
	0x3d, 0x24, 0x95, 0xef, 0x22, 0xb7, 0x6a, 0x02, 0x73, 0xab, 0x5a, 0xa3, 0xa0, 0xad, 0x6a, 0x23,
	0xb9, 0xca, 0xb4, 0x5a, 0x3b, 0xf5, 0x58, 0x68, 0x74, 0x89, 0xe6, 0x78, 0x36, 0x19, 0x5a, 0xed,
	0x94, 0x29, 0xd0, 0xe5, 0x47, 0xa1, 0x0a, 0x7d, 0x15, 0xa7, 0x1c, 0x45, 0x13, 0xb0, 0x2d, 0x41,
	0x9f, 0x93, 0xe6, 0x45, 0x6c, 0x9d, 0x6d, 0x74, 0x6e, 0x30, 0xbc, 0x65, 0xba, 0x90, 0x19, 0xc7,
	0x49, 0x07, 0xcc, 0x00, 0x60, 0xcf, 0xd3, 0x58, 0x2c, 0x71, 0xb4, 0x01, 0x33, 0xa0, 0xff, 0xdb,
	0x27, 0x9d, 0x5d, 0x09, 0xfe, 0x6f, 0xf8, 0x3f, 0x6d, 0xf8, 0x2f, 0x8f, 0x34, 0x41, 0xe2, 0xb8,
	0x1d, 0x5e, 0x12, 0xe2, 0xe4, 0x3e, 0x4e, 0x6c, 0xb7, 0x4b, 0x0c, 0x94, 0x63, 0x3e, 0x2b, 0xb7,
	0x42, 0x1d, 0xb4, 0x43, 0xc2, 0xbd, 0x53, 0xd9, 0x0c, 0x09, 0x73, 0x9e, 0x10, 0xea, 0x32, 0x4c,
	0x95, 0xbc, 0x17, 0xe6, 0xb8, 0x59, 0xa6, 0x7f, 0xf1, 0xd0, 0x77, 0xe4, 0xd9, 0x3e, 0x3b, 0x55,
	0x32, 0x59, 0xcd, 0xf5, 0x38, 0xb1, 0xbb, 0xf6, 0xf1, 0x00, 0xdc, 0xbe, 0xb1, 0x52, 0x02, 0x2f,
	0xa9, 0x63, 0xf4, 0x96, 0xe8, 0xff, 0xf4, 0x48, 0xb0, 0xd9, 0x0b, 0xf4, 0x35, 0x69, 0xe4, 0xe6,
	0x20, 0x96, 0xda, 0x1a, 0x76, 0xdd, 0x2a, 0xb0, 0xf9, 0x98, 0xf3, 0x43, 0xb7, 0xbf, 0xac, 0xe2,
	0x4c, 0x83, 0x0e, 0x7c, 0xdc, 0x57, 0x1b, 0x0c, 0x4d, 0xbb, 0x96, 0x3a, 0x5e, 0x4e, 0x95, 0x98,
	0x73, 0xac, 0xde, 0x67, 0x25, 0x06, 0xce, 0x7e, 0x10, 0xc5, 0x1c, 0x26, 0x8b, 0x65, 0xfb, 0x6c,
	0x83, 0xfb, 0x3f, 0xbc, 0xed, 0x3e, 0xb2, 0xff, 0x26, 0x78, 0x4a, 0x05, 0xff, 0x4d, 0x47, 0xa4,
	0x6e, 0x15, 0xec, 0x63, 0x21, 0x16, 0x81, 0x20, 0xf5, 0x3a, 0x37, 0x57, 0xd5, 0x18, 0xda, 0xc0,
	0xa1, 0x14, 0x8c, 0x9e, 0xab, 0x4e, 0x06, 0x66, 0xe0, 0xb5, 0xd2, 0xc0, 0x21, 0xf2, 0x9a, 0x2f,
	0x4f, 0x9d, 0x9c, 0xc1, 0xee, 0x47, 0xa4, 0x61, 0x4b, 0x06, 0x65, 0xcf, 0x3e, 0x7f, 0xb5, 0xb3,
	0x07, 0x13, 0x98, 0xf3, 0xd1, 0x95, 0x1d, 0x38, 0x98, 0x90, 0xf8, 0x5a, 0xe8, 0x25, 0xb7, 0xa3,
	0x36, 0xe0, 0xa6, 0x8e, 0x7f, 0xeb, 0xb3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xbd, 0x59,
	0x8a, 0xbd, 0x07, 0x00, 0x00,
}
