// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: proto/ticket.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier for the ticket.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the person who has purchased the ticket.
	Purchaser string `protobuf:"bytes,2,opt,name=purchaser,proto3" json:"purchaser,omitempty"`
	// If the ticket is accompanied by a guest.
	IsBringingGuest bool `protobuf:"varint,3,opt,name=isBringingGuest,proto3" json:"isBringingGuest,omitempty"`
	// If the ticket has been physically presented to the guest.
	HasReceivedTicket bool `protobuf:"varint,4,opt,name=hasReceivedTicket,proto3" json:"hasReceivedTicket,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket) GetPurchaser() string {
	if x != nil {
		return x.Purchaser
	}
	return ""
}

func (x *Ticket) GetIsBringingGuest() bool {
	if x != nil {
		return x.IsBringingGuest
	}
	return false
}

func (x *Ticket) GetHasReceivedTicket() bool {
	if x != nil {
		return x.HasReceivedTicket
	}
	return false
}

type BuyTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purchaser       string `protobuf:"bytes,1,opt,name=purchaser,proto3" json:"purchaser,omitempty"`
	IsBringingGuest bool   `protobuf:"varint,2,opt,name=isBringingGuest,proto3" json:"isBringingGuest,omitempty"`
}

func (x *BuyTicketRequest) Reset() {
	*x = BuyTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyTicketRequest) ProtoMessage() {}

func (x *BuyTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyTicketRequest.ProtoReflect.Descriptor instead.
func (*BuyTicketRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *BuyTicketRequest) GetPurchaser() string {
	if x != nil {
		return x.Purchaser
	}
	return ""
}

func (x *BuyTicketRequest) GetIsBringingGuest() bool {
	if x != nil {
		return x.IsBringingGuest
	}
	return false
}

type BuyTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *BuyTicketResponse) Reset() {
	*x = BuyTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyTicketResponse) ProtoMessage() {}

func (x *BuyTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyTicketResponse.ProtoReflect.Descriptor instead.
func (*BuyTicketResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *BuyTicketResponse) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type ListTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ListTicketRequest) Reset() {
	*x = ListTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketRequest) ProtoMessage() {}

func (x *ListTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketRequest.ProtoReflect.Descriptor instead.
func (*ListTicketRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *ListTicketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *ListTicketResponse) Reset() {
	*x = ListTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketResponse) ProtoMessage() {}

func (x *ListTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketResponse.ProtoReflect.Descriptor instead.
func (*ListTicketResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *ListTicketResponse) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type ListAllTicketsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAllTicketsRequest) Reset() {
	*x = ListAllTicketsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllTicketsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllTicketsRequest) ProtoMessage() {}

func (x *ListAllTicketsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllTicketsRequest.ProtoReflect.Descriptor instead.
func (*ListAllTicketsRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{5}
}

type ListAllTicketsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tickets []*Ticket `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
}

func (x *ListAllTicketsResponse) Reset() {
	*x = ListAllTicketsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllTicketsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllTicketsResponse) ProtoMessage() {}

func (x *ListAllTicketsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllTicketsResponse.ProtoReflect.Descriptor instead.
func (*ListAllTicketsResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *ListAllTicketsResponse) GetTickets() []*Ticket {
	if x != nil {
		return x.Tickets
	}
	return nil
}

type UpdateTicketInformationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsBringingGuest   bool   `protobuf:"varint,2,opt,name=isBringingGuest,proto3" json:"isBringingGuest,omitempty"`
	HasReceivedTicket bool   `protobuf:"varint,3,opt,name=hasReceivedTicket,proto3" json:"hasReceivedTicket,omitempty"`
}

func (x *UpdateTicketInformationRequest) Reset() {
	*x = UpdateTicketInformationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTicketInformationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTicketInformationRequest) ProtoMessage() {}

func (x *UpdateTicketInformationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTicketInformationRequest.ProtoReflect.Descriptor instead.
func (*UpdateTicketInformationRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTicketInformationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTicketInformationRequest) GetIsBringingGuest() bool {
	if x != nil {
		return x.IsBringingGuest
	}
	return false
}

func (x *UpdateTicketInformationRequest) GetHasReceivedTicket() bool {
	if x != nil {
		return x.HasReceivedTicket
	}
	return false
}

type UpdateTicketInformationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *UpdateTicketInformationResponse) Reset() {
	*x = UpdateTicketInformationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTicketInformationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTicketInformationResponse) ProtoMessage() {}

func (x *UpdateTicketInformationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTicketInformationResponse.ProtoReflect.Descriptor instead.
func (*UpdateTicketInformationResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTicketInformationResponse) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type DeleteTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTicketRequest) Reset() {
	*x = DeleteTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketRequest) ProtoMessage() {}

func (x *DeleteTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketRequest.ProtoReflect.Descriptor instead.
func (*DeleteTicketRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteTicketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTicketResponse) Reset() {
	*x = DeleteTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketResponse) ProtoMessage() {}

func (x *DeleteTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketResponse.ProtoReflect.Descriptor instead.
func (*DeleteTicketResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{10}
}

var File_proto_ticket_proto protoreflect.FileDescriptor

var file_proto_ticket_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x06,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x73, 0x42, 0x72, 0x69, 0x6e, 0x67, 0x69,
	0x6e, 0x67, 0x47, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69,
	0x73, 0x42, 0x72, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x11, 0x68, 0x61, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x68, 0x61, 0x73, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x5a, 0x0a, 0x10,
	0x42, 0x75, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x0f, 0x69, 0x73, 0x42, 0x72, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x42, 0x72, 0x69, 0x6e, 0x67,
	0x69, 0x6e, 0x67, 0x47, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x11, 0x42, 0x75, 0x79, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x41, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x73, 0x42, 0x72, 0x69, 0x6e, 0x67,
	0x69, 0x6e, 0x67, 0x47, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x69, 0x73, 0x42, 0x72, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x11, 0x68, 0x61, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x68, 0x61, 0x73, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x48, 0x0a,
	0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9e, 0x03, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x09, 0x42, 0x75, 0x79, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75,
	0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4f, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6a, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ticket_proto_rawDescOnce sync.Once
	file_proto_ticket_proto_rawDescData = file_proto_ticket_proto_rawDesc
)

func file_proto_ticket_proto_rawDescGZIP() []byte {
	file_proto_ticket_proto_rawDescOnce.Do(func() {
		file_proto_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ticket_proto_rawDescData)
	})
	return file_proto_ticket_proto_rawDescData
}

var file_proto_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_ticket_proto_goTypes = []interface{}{
	(*Ticket)(nil),                          // 0: proto.Ticket
	(*BuyTicketRequest)(nil),                // 1: proto.BuyTicketRequest
	(*BuyTicketResponse)(nil),               // 2: proto.BuyTicketResponse
	(*ListTicketRequest)(nil),               // 3: proto.ListTicketRequest
	(*ListTicketResponse)(nil),              // 4: proto.ListTicketResponse
	(*ListAllTicketsRequest)(nil),           // 5: proto.ListAllTicketsRequest
	(*ListAllTicketsResponse)(nil),          // 6: proto.ListAllTicketsResponse
	(*UpdateTicketInformationRequest)(nil),  // 7: proto.UpdateTicketInformationRequest
	(*UpdateTicketInformationResponse)(nil), // 8: proto.UpdateTicketInformationResponse
	(*DeleteTicketRequest)(nil),             // 9: proto.DeleteTicketRequest
	(*DeleteTicketResponse)(nil),            // 10: proto.DeleteTicketResponse
}
var file_proto_ticket_proto_depIdxs = []int32{
	0,  // 0: proto.BuyTicketResponse.ticket:type_name -> proto.Ticket
	0,  // 1: proto.ListTicketResponse.ticket:type_name -> proto.Ticket
	0,  // 2: proto.ListAllTicketsResponse.tickets:type_name -> proto.Ticket
	0,  // 3: proto.UpdateTicketInformationResponse.ticket:type_name -> proto.Ticket
	1,  // 4: proto.TicketManager.BuyTicket:input_type -> proto.BuyTicketRequest
	3,  // 5: proto.TicketManager.ListTicket:input_type -> proto.ListTicketRequest
	5,  // 6: proto.TicketManager.ListAllTickets:input_type -> proto.ListAllTicketsRequest
	7,  // 7: proto.TicketManager.UpdateTicketInformation:input_type -> proto.UpdateTicketInformationRequest
	9,  // 8: proto.TicketManager.DeleteTicket:input_type -> proto.DeleteTicketRequest
	2,  // 9: proto.TicketManager.BuyTicket:output_type -> proto.BuyTicketResponse
	4,  // 10: proto.TicketManager.ListTicket:output_type -> proto.ListTicketResponse
	6,  // 11: proto.TicketManager.ListAllTickets:output_type -> proto.ListAllTicketsResponse
	8,  // 12: proto.TicketManager.UpdateTicketInformation:output_type -> proto.UpdateTicketInformationResponse
	10, // 13: proto.TicketManager.DeleteTicket:output_type -> proto.DeleteTicketResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_ticket_proto_init() }
func file_proto_ticket_proto_init() {
	if File_proto_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyTicketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyTicketResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTicketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTicketResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllTicketsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllTicketsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTicketInformationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTicketInformationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTicketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTicketResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ticket_proto_goTypes,
		DependencyIndexes: file_proto_ticket_proto_depIdxs,
		MessageInfos:      file_proto_ticket_proto_msgTypes,
	}.Build()
	File_proto_ticket_proto = out.File
	file_proto_ticket_proto_rawDesc = nil
	file_proto_ticket_proto_goTypes = nil
	file_proto_ticket_proto_depIdxs = nil
}
