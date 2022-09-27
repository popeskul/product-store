// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/products.proto

package domain

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

type ListSortingParams_SortingType int32

const (
	ListSortingParams_name          ListSortingParams_SortingType = 0
	ListSortingParams_price         ListSortingParams_SortingType = 1
	ListSortingParams_changes_count ListSortingParams_SortingType = 2
	ListSortingParams_timestamp     ListSortingParams_SortingType = 3
)

// Enum value maps for ListSortingParams_SortingType.
var (
	ListSortingParams_SortingType_name = map[int32]string{
		0: "name",
		1: "price",
		2: "changes_count",
		3: "timestamp",
	}
	ListSortingParams_SortingType_value = map[string]int32{
		"name":          0,
		"price":         1,
		"changes_count": 2,
		"timestamp":     3,
	}
)

func (x ListSortingParams_SortingType) Enum() *ListSortingParams_SortingType {
	p := new(ListSortingParams_SortingType)
	*p = x
	return p
}

func (x ListSortingParams_SortingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListSortingParams_SortingType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_products_proto_enumTypes[0].Descriptor()
}

func (ListSortingParams_SortingType) Type() protoreflect.EnumType {
	return &file_proto_products_proto_enumTypes[0]
}

func (x ListSortingParams_SortingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListSortingParams_SortingType.Descriptor instead.
func (ListSortingParams_SortingType) EnumDescriptor() ([]byte, []int) {
	return file_proto_products_proto_rawDescGZIP(), []int{4, 0}
}

type ListSortingParams_SortDirection int32

const (
	ListSortingParams_asc  ListSortingParams_SortDirection = 0
	ListSortingParams_desc ListSortingParams_SortDirection = 1
)

// Enum value maps for ListSortingParams_SortDirection.
var (
	ListSortingParams_SortDirection_name = map[int32]string{
		0: "asc",
		1: "desc",
	}
	ListSortingParams_SortDirection_value = map[string]int32{
		"asc":  0,
		"desc": 1,
	}
)

func (x ListSortingParams_SortDirection) Enum() *ListSortingParams_SortDirection {
	p := new(ListSortingParams_SortDirection)
	*p = x
	return p
}

func (x ListSortingParams_SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListSortingParams_SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_products_proto_enumTypes[1].Descriptor()
}

func (ListSortingParams_SortDirection) Type() protoreflect.EnumType {
	return &file_proto_products_proto_enumTypes[1]
}

func (x ListSortingParams_SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListSortingParams_SortDirection.Descriptor instead.
func (ListSortingParams_SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_proto_products_proto_rawDescGZIP(), []int{4, 1}
}

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_proto_products_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_products_proto_rawDescGZIP(), []int{1}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging  *ListPagingParams  `protobuf:"bytes,1,opt,name=paging,proto3" json:"paging,omitempty"`
	Sorting *ListSortingParams `protobuf:"bytes,2,opt,name=sorting,proto3" json:"sorting,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_products_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetPaging() *ListPagingParams {
	if x != nil {
		return x.Paging
	}
	return nil
}

func (x *ListRequest) GetSorting() *ListSortingParams {
	if x != nil {
		return x.Sorting
	}
	return nil
}

type ListPagingParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *ListPagingParams) Reset() {
	*x = ListPagingParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPagingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPagingParams) ProtoMessage() {}

func (x *ListPagingParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPagingParams.ProtoReflect.Descriptor instead.
func (*ListPagingParams) Descriptor() ([]byte, []int) {
	return file_proto_products_proto_rawDescGZIP(), []int{3}
}

func (x *ListPagingParams) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPagingParams) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ListSortingParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortBy        ListSortingParams_SortingType   `protobuf:"varint,1,opt,name=sort_by,json=sortBy,proto3,enum=ListSortingParams_SortingType" json:"sort_by,omitempty"`
	SortDirection ListSortingParams_SortDirection `protobuf:"varint,2,opt,name=sort_direction,json=sortDirection,proto3,enum=ListSortingParams_SortDirection" json:"sort_direction,omitempty"`
}

func (x *ListSortingParams) Reset() {
	*x = ListSortingParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSortingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSortingParams) ProtoMessage() {}

func (x *ListSortingParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSortingParams.ProtoReflect.Descriptor instead.
func (*ListSortingParams) Descriptor() ([]byte, []int) {
	return file_proto_products_proto_rawDescGZIP(), []int{4}
}

func (x *ListSortingParams) GetSortBy() ListSortingParams_SortingType {
	if x != nil {
		return x.SortBy
	}
	return ListSortingParams_name
}

func (x *ListSortingParams) GetSortDirection() ListSortingParams_SortDirection {
	if x != nil {
		return x.SortDirection
	}
	return ListSortingParams_asc
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Total    int32      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_products_proto_rawDescGZIP(), []int{5}
}

func (x *ListResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_products_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_products_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_products_proto_rawDescGZIP(), []int{6}
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_proto_products_proto protoreflect.FileDescriptor

var file_proto_products_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x66, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x41, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0xff, 0x01, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x47, 0x0a, 0x0e, 0x73,
	0x6f, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x0b, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x10, 0x03, 0x22, 0x22, 0x0a, 0x0d, 0x53, 0x6f,
	0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x61,
	0x73, 0x63, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x10, 0x01, 0x22, 0x4a,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x33, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32,
	0x5a, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x0d, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x70,
	0x6b, 0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_products_proto_rawDescOnce sync.Once
	file_proto_products_proto_rawDescData = file_proto_products_proto_rawDesc
)

func file_proto_products_proto_rawDescGZIP() []byte {
	file_proto_products_proto_rawDescOnce.Do(func() {
		file_proto_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_products_proto_rawDescData)
	})
	return file_proto_products_proto_rawDescData
}

var file_proto_products_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_products_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_products_proto_goTypes = []interface{}{
	(ListSortingParams_SortingType)(0),   // 0: ListSortingParams.SortingType
	(ListSortingParams_SortDirection)(0), // 1: ListSortingParams.SortDirection
	(*FetchRequest)(nil),                 // 2: FetchRequest
	(*Empty)(nil),                        // 3: Empty
	(*ListRequest)(nil),                  // 4: ListRequest
	(*ListPagingParams)(nil),             // 5: ListPagingParams
	(*ListSortingParams)(nil),            // 6: ListSortingParams
	(*ListResponse)(nil),                 // 7: ListResponse
	(*Product)(nil),                      // 8: Product
}
var file_proto_products_proto_depIdxs = []int32{
	5, // 0: ListRequest.paging:type_name -> ListPagingParams
	6, // 1: ListRequest.sorting:type_name -> ListSortingParams
	0, // 2: ListSortingParams.sort_by:type_name -> ListSortingParams.SortingType
	1, // 3: ListSortingParams.sort_direction:type_name -> ListSortingParams.SortDirection
	8, // 4: ListResponse.products:type_name -> Product
	2, // 5: ProductsService.Fetch:input_type -> FetchRequest
	4, // 6: ProductsService.List:input_type -> ListRequest
	3, // 7: ProductsService.Fetch:output_type -> Empty
	7, // 8: ProductsService.List:output_type -> ListResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_products_proto_init() }
func file_proto_products_proto_init() {
	if File_proto_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_proto_products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_products_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_proto_products_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPagingParams); i {
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
		file_proto_products_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSortingParams); i {
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
		file_proto_products_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_proto_products_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_proto_products_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_products_proto_goTypes,
		DependencyIndexes: file_proto_products_proto_depIdxs,
		EnumInfos:         file_proto_products_proto_enumTypes,
		MessageInfos:      file_proto_products_proto_msgTypes,
	}.Build()
	File_proto_products_proto = out.File
	file_proto_products_proto_rawDesc = nil
	file_proto_products_proto_goTypes = nil
	file_proto_products_proto_depIdxs = nil
}