// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: coffee.proto

package coffee

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

type AddCoffeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *AddCoffeeRequest) Reset() {
	*x = AddCoffeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCoffeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCoffeeRequest) ProtoMessage() {}

func (x *AddCoffeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCoffeeRequest.ProtoReflect.Descriptor instead.
func (*AddCoffeeRequest) Descriptor() ([]byte, []int) {
	return file_coffee_proto_rawDescGZIP(), []int{0}
}

func (x *AddCoffeeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddCoffeeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCoffeeRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type AddCoffeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *AddCoffeeResponse) Reset() {
	*x = AddCoffeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCoffeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCoffeeResponse) ProtoMessage() {}

func (x *AddCoffeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCoffeeResponse.ProtoReflect.Descriptor instead.
func (*AddCoffeeResponse) Descriptor() ([]byte, []int) {
	return file_coffee_proto_rawDescGZIP(), []int{1}
}

func (x *AddCoffeeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCoffeeResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateCoffeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	CoffeeName string `protobuf:"bytes,2,opt,name=coffeeName,proto3" json:"coffeeName,omitempty"`
}

func (x *CreateCoffeeResponse) Reset() {
	*x = CreateCoffeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoffeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoffeeResponse) ProtoMessage() {}

func (x *CreateCoffeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoffeeResponse.ProtoReflect.Descriptor instead.
func (*CreateCoffeeResponse) Descriptor() ([]byte, []int) {
	return file_coffee_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCoffeeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateCoffeeResponse) GetCoffeeName() string {
	if x != nil {
		return x.CoffeeName
	}
	return ""
}

type FindCoffeeById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindCoffeeById) Reset() {
	*x = FindCoffeeById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCoffeeById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCoffeeById) ProtoMessage() {}

func (x *FindCoffeeById) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCoffeeById.ProtoReflect.Descriptor instead.
func (*FindCoffeeById) Descriptor() ([]byte, []int) {
	return file_coffee_proto_rawDescGZIP(), []int{3}
}

func (x *FindCoffeeById) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCoffeeOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCoffeeOrder) Reset() {
	*x = DeleteCoffeeOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCoffeeOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCoffeeOrder) ProtoMessage() {}

func (x *DeleteCoffeeOrder) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCoffeeOrder.ProtoReflect.Descriptor instead.
func (*DeleteCoffeeOrder) Descriptor() ([]byte, []int) {
	return file_coffee_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCoffeeOrder) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteOrderResponse) Reset() {
	*x = DeleteOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderResponse) ProtoMessage() {}

func (x *DeleteOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderResponse.ProtoReflect.Descriptor instead.
func (*DeleteOrderResponse) Descriptor() ([]byte, []int) {
	return file_coffee_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateCoffee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *AddCoffeeRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *UpdateCoffee) Reset() {
	*x = UpdateCoffee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoffee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoffee) ProtoMessage() {}

func (x *UpdateCoffee) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoffee.ProtoReflect.Descriptor instead.
func (*UpdateCoffee) Descriptor() ([]byte, []int) {
	return file_coffee_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCoffee) GetRequest() *AddCoffeeRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetAllCoffee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allcoffees []*AddCoffeeRequest `protobuf:"bytes,1,rep,name=Allcoffees,proto3" json:"Allcoffees,omitempty"`
}

func (x *GetAllCoffee) Reset() {
	*x = GetAllCoffee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCoffee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCoffee) ProtoMessage() {}

func (x *GetAllCoffee) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCoffee.ProtoReflect.Descriptor instead.
func (*GetAllCoffee) Descriptor() ([]byte, []int) {
	return file_coffee_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllCoffee) GetAllcoffees() []*AddCoffeeRequest {
	if x != nil {
		return x.Allcoffees
	}
	return nil
}

var File_coffee_proto protoreflect.FileDescriptor

var file_coffee_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x3d, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x50, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a,
	0x0e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x66, 0x66, 0x65, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x66, 0x66,
	0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x66, 0x66,
	0x65, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x66, 0x66,
	0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x41, 0x6c, 0x6c, 0x63, 0x6f,
	0x66, 0x66, 0x65, 0x65, 0x73, 0x32, 0xa3, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x66, 0x66, 0x65, 0x65, 0x12, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x66,
	0x66, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x46,
	0x69, 0x6e, 0x64, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0f, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x12,
	0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x66, 0x66,
	0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x66, 0x66, 0x65, 0x65, 0x1a, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x66, 0x66, 0x65,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x1a, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x12, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x1a, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x66,
	0x66, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coffee_proto_rawDescOnce sync.Once
	file_coffee_proto_rawDescData = file_coffee_proto_rawDesc
)

func file_coffee_proto_rawDescGZIP() []byte {
	file_coffee_proto_rawDescOnce.Do(func() {
		file_coffee_proto_rawDescData = protoimpl.X.CompressGZIP(file_coffee_proto_rawDescData)
	})
	return file_coffee_proto_rawDescData
}

var file_coffee_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_coffee_proto_goTypes = []interface{}{
	(*AddCoffeeRequest)(nil),     // 0: AddCoffeeRequest
	(*AddCoffeeResponse)(nil),    // 1: AddCoffeeResponse
	(*CreateCoffeeResponse)(nil), // 2: CreateCoffeeResponse
	(*FindCoffeeById)(nil),       // 3: FindCoffeeById
	(*DeleteCoffeeOrder)(nil),    // 4: DeleteCoffeeOrder
	(*DeleteOrderResponse)(nil),  // 5: DeleteOrderResponse
	(*UpdateCoffee)(nil),         // 6: UpdateCoffee
	(*GetAllCoffee)(nil),         // 7: GetAllCoffee
}
var file_coffee_proto_depIdxs = []int32{
	0, // 0: UpdateCoffee.request:type_name -> AddCoffeeRequest
	0, // 1: GetAllCoffee.Allcoffees:type_name -> AddCoffeeRequest
	0, // 2: CoffeeService.AddCoffee:input_type -> AddCoffeeRequest
	3, // 3: CoffeeService.FindCoffeeByID:input_type -> FindCoffeeById
	6, // 4: CoffeeService.UpdateCoffeeInfo:input_type -> UpdateCoffee
	4, // 5: CoffeeService.DeleteCoffeeInfo:input_type -> DeleteCoffeeOrder
	7, // 6: CoffeeService.FindAllCoffee:input_type -> GetAllCoffee
	1, // 7: CoffeeService.AddCoffee:output_type -> AddCoffeeResponse
	1, // 8: CoffeeService.FindCoffeeByID:output_type -> AddCoffeeResponse
	1, // 9: CoffeeService.UpdateCoffeeInfo:output_type -> AddCoffeeResponse
	5, // 10: CoffeeService.DeleteCoffeeInfo:output_type -> DeleteOrderResponse
	1, // 11: CoffeeService.FindAllCoffee:output_type -> AddCoffeeResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coffee_proto_init() }
func file_coffee_proto_init() {
	if File_coffee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coffee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCoffeeRequest); i {
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
		file_coffee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCoffeeResponse); i {
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
		file_coffee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoffeeResponse); i {
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
		file_coffee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCoffeeById); i {
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
		file_coffee_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCoffeeOrder); i {
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
		file_coffee_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrderResponse); i {
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
		file_coffee_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCoffee); i {
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
		file_coffee_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCoffee); i {
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
			RawDescriptor: file_coffee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coffee_proto_goTypes,
		DependencyIndexes: file_coffee_proto_depIdxs,
		MessageInfos:      file_coffee_proto_msgTypes,
	}.Build()
	File_coffee_proto = out.File
	file_coffee_proto_rawDesc = nil
	file_coffee_proto_goTypes = nil
	file_coffee_proto_depIdxs = nil
}
