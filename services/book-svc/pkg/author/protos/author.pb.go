// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/author.proto

package protos

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type GetAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int64 `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *GetAuthorRequest) Reset() {
	*x = GetAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorRequest) ProtoMessage() {}

func (x *GetAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorRequest) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{0}
}

func (x *GetAuthorRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type GetAllAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search *wrappers.StringValue `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Limit  int32                 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32                 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetAllAuthorRequest) Reset() {
	*x = GetAllAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAuthorRequest) ProtoMessage() {}

func (x *GetAllAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAuthorRequest.ProtoReflect.Descriptor instead.
func (*GetAllAuthorRequest) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllAuthorRequest) GetSearch() *wrappers.StringValue {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *GetAllAuthorRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllAuthorRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetAllAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
}

func (x *GetAllAuthorResponse) Reset() {
	*x = GetAllAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAuthorResponse) ProtoMessage() {}

func (x *GetAllAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAuthorResponse.ProtoReflect.Descriptor instead.
func (*GetAllAuthorResponse) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllAuthorResponse) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

type CreateAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string               `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string               `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Biography   string               `protobuf:"bytes,4,opt,name=biography,proto3" json:"biography,omitempty"`
	DateOfBirth *timestamp.Timestamp `protobuf:"bytes,5,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
}

func (x *CreateAuthorRequest) Reset() {
	*x = CreateAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorRequest) ProtoMessage() {}

func (x *CreateAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthorRequest) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAuthorRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateAuthorRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateAuthorRequest) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *CreateAuthorRequest) GetDateOfBirth() *timestamp.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

type UpdateAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int64   `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Authors  *Author `protobuf:"bytes,2,opt,name=authors,proto3" json:"authors,omitempty"`
}

func (x *UpdateAuthorRequest) Reset() {
	*x = UpdateAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthorRequest) ProtoMessage() {}

func (x *UpdateAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthorRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuthorRequest) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAuthorRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *UpdateAuthorRequest) GetAuthors() *Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

type DeleteAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int64 `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *DeleteAuthorRequest) Reset() {
	*x = DeleteAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthorRequest) ProtoMessage() {}

func (x *DeleteAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthorRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthorRequest) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAuthorRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId    int64                `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	FirstName   string               `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string               `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Biography   string               `protobuf:"bytes,4,opt,name=biography,proto3" json:"biography,omitempty"`
	DateOfBirth *timestamp.Timestamp `protobuf:"bytes,5,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{6}
}

func (x *Author) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *Author) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Author) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Author) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *Author) GetDateOfBirth() *timestamp.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *Author) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Author) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[7]
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
	return file_protos_author_proto_rawDescGZIP(), []int{7}
}

var File_protos_author_proto protoreflect.FileDescriptor

var file_protos_author_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22,
	0x79, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x40, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x22, 0xaf, 0x01, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x3e,
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x22, 0x5c,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x22, 0x32, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x22, 0xb5, 0x02, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x32, 0xc7, 0x02, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x3a, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x31, 0x5a, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x49, 0x50, 0x4b, 0x49, 0x4e, 0x39, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2d,
	0x73, 0x76, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_author_proto_rawDescOnce sync.Once
	file_protos_author_proto_rawDescData = file_protos_author_proto_rawDesc
)

func file_protos_author_proto_rawDescGZIP() []byte {
	file_protos_author_proto_rawDescOnce.Do(func() {
		file_protos_author_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_author_proto_rawDescData)
	})
	return file_protos_author_proto_rawDescData
}

var file_protos_author_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_author_proto_goTypes = []any{
	(*GetAuthorRequest)(nil),     // 0: author.GetAuthorRequest
	(*GetAllAuthorRequest)(nil),  // 1: author.GetAllAuthorRequest
	(*GetAllAuthorResponse)(nil), // 2: author.GetAllAuthorResponse
	(*CreateAuthorRequest)(nil),  // 3: author.CreateAuthorRequest
	(*UpdateAuthorRequest)(nil),  // 4: author.UpdateAuthorRequest
	(*DeleteAuthorRequest)(nil),  // 5: author.DeleteAuthorRequest
	(*Author)(nil),               // 6: author.Author
	(*Empty)(nil),                // 7: author.Empty
	(*wrappers.StringValue)(nil), // 8: google.protobuf.StringValue
	(*timestamp.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_protos_author_proto_depIdxs = []int32{
	8,  // 0: author.GetAllAuthorRequest.search:type_name -> google.protobuf.StringValue
	6,  // 1: author.GetAllAuthorResponse.authors:type_name -> author.Author
	9,  // 2: author.CreateAuthorRequest.date_of_birth:type_name -> google.protobuf.Timestamp
	6,  // 3: author.UpdateAuthorRequest.authors:type_name -> author.Author
	9,  // 4: author.Author.date_of_birth:type_name -> google.protobuf.Timestamp
	9,  // 5: author.Author.created_at:type_name -> google.protobuf.Timestamp
	9,  // 6: author.Author.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 7: author.AuthorService.GetAuthor:input_type -> author.GetAuthorRequest
	1,  // 8: author.AuthorService.GetAllAuthor:input_type -> author.GetAllAuthorRequest
	3,  // 9: author.AuthorService.CreateAuthor:input_type -> author.CreateAuthorRequest
	4,  // 10: author.AuthorService.UpdateAuthor:input_type -> author.UpdateAuthorRequest
	5,  // 11: author.AuthorService.DeleteAuthor:input_type -> author.DeleteAuthorRequest
	6,  // 12: author.AuthorService.GetAuthor:output_type -> author.Author
	2,  // 13: author.AuthorService.GetAllAuthor:output_type -> author.GetAllAuthorResponse
	6,  // 14: author.AuthorService.CreateAuthor:output_type -> author.Author
	6,  // 15: author.AuthorService.UpdateAuthor:output_type -> author.Author
	7,  // 16: author.AuthorService.DeleteAuthor:output_type -> author.Empty
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_protos_author_proto_init() }
func file_protos_author_proto_init() {
	if File_protos_author_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_author_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAuthorRequest); i {
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
		file_protos_author_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllAuthorRequest); i {
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
		file_protos_author_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllAuthorResponse); i {
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
		file_protos_author_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAuthorRequest); i {
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
		file_protos_author_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateAuthorRequest); i {
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
		file_protos_author_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteAuthorRequest); i {
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
		file_protos_author_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Author); i {
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
		file_protos_author_proto_msgTypes[7].Exporter = func(v any, i int) any {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_author_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_author_proto_goTypes,
		DependencyIndexes: file_protos_author_proto_depIdxs,
		MessageInfos:      file_protos_author_proto_msgTypes,
	}.Build()
	File_protos_author_proto = out.File
	file_protos_author_proto_rawDesc = nil
	file_protos_author_proto_goTypes = nil
	file_protos_author_proto_depIdxs = nil
}
