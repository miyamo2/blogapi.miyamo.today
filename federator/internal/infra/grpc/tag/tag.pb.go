// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: tag.proto

package tag

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTagByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTagByIdRequest) Reset() {
	*x = GetTagByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagByIdRequest) ProtoMessage() {}

func (x *GetTagByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagByIdRequest.ProtoReflect.Descriptor instead.
func (*GetTagByIdRequest) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{0}
}

func (x *GetTagByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetNextTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First int32   `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	After *string `protobuf:"bytes,2,opt,name=after,proto3,oneof" json:"after,omitempty"`
}

func (x *GetNextTagsRequest) Reset() {
	*x = GetNextTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNextTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNextTagsRequest) ProtoMessage() {}

func (x *GetNextTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNextTagsRequest.ProtoReflect.Descriptor instead.
func (*GetNextTagsRequest) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{1}
}

func (x *GetNextTagsRequest) GetFirst() int32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *GetNextTagsRequest) GetAfter() string {
	if x != nil && x.After != nil {
		return *x.After
	}
	return ""
}

type GetPrevTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Last   int32   `protobuf:"varint,1,opt,name=last,proto3" json:"last,omitempty"`
	Before *string `protobuf:"bytes,2,opt,name=before,proto3,oneof" json:"before,omitempty"`
}

func (x *GetPrevTagsRequest) Reset() {
	*x = GetPrevTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrevTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrevTagsRequest) ProtoMessage() {}

func (x *GetPrevTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrevTagsRequest.ProtoReflect.Descriptor instead.
func (*GetPrevTagsRequest) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{2}
}

func (x *GetPrevTagsRequest) GetLast() int32 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *GetPrevTagsRequest) GetBefore() string {
	if x != nil && x.Before != nil {
		return *x.Before
	}
	return ""
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Articles []*Article `protobuf:"bytes,3,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{3}
}

func (x *Tag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetArticles() []*Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body         string                 `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	ThumbnailUrl string                 `protobuf:"bytes,4,opt,name=thumbnailUrl,proto3" json:"thumbnailUrl,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{4}
}

func (x *Article) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Article) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *Article) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Article) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetTagByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *GetTagByIdResponse) Reset() {
	*x = GetTagByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagByIdResponse) ProtoMessage() {}

func (x *GetTagByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagByIdResponse.ProtoReflect.Descriptor instead.
func (*GetTagByIdResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{5}
}

func (x *GetTagByIdResponse) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type GetAllTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *GetAllTagsResponse) Reset() {
	*x = GetAllTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTagsResponse) ProtoMessage() {}

func (x *GetAllTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTagsResponse.ProtoReflect.Descriptor instead.
func (*GetAllTagsResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllTagsResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type GetNextTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags        []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	StillExists bool   `protobuf:"varint,2,opt,name=stillExists,proto3" json:"stillExists,omitempty"`
}

func (x *GetNextTagResponse) Reset() {
	*x = GetNextTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNextTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNextTagResponse) ProtoMessage() {}

func (x *GetNextTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNextTagResponse.ProtoReflect.Descriptor instead.
func (*GetNextTagResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{7}
}

func (x *GetNextTagResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetNextTagResponse) GetStillExists() bool {
	if x != nil {
		return x.StillExists
	}
	return false
}

type GetPrevTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags        []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	StillExists bool   `protobuf:"varint,2,opt,name=stillExists,proto3" json:"stillExists,omitempty"`
}

func (x *GetPrevTagResponse) Reset() {
	*x = GetPrevTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrevTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrevTagResponse) ProtoMessage() {}

func (x *GetPrevTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrevTagResponse.ProtoReflect.Descriptor instead.
func (*GetPrevTagResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{8}
}

func (x *GetPrevTagResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetPrevTagResponse) GetStillExists() bool {
	if x != nil {
		return x.StillExists
	}
	return false
}

var File_tag_proto protoreflect.FileDescriptor

var file_tag_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x74, 0x61, 0x67,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x61,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x54,
	0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0x53, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x07,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x30, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x67, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74,
	0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x32, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22,
	0x54, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x54, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x76,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x61, 0x67, 0x2e,
	0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x69,
	0x6c, 0x6c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x73, 0x74, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0x8c, 0x02, 0x0a, 0x0a,
	0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x67, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x17, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x78, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x65, 0x76, 0x54, 0x61, 0x67, 0x73, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tag_proto_rawDescOnce sync.Once
	file_tag_proto_rawDescData = file_tag_proto_rawDesc
)

func file_tag_proto_rawDescGZIP() []byte {
	file_tag_proto_rawDescOnce.Do(func() {
		file_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_tag_proto_rawDescData)
	})
	return file_tag_proto_rawDescData
}

var file_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tag_proto_goTypes = []interface{}{
	(*GetTagByIdRequest)(nil),     // 0: tag.GetTagByIdRequest
	(*GetNextTagsRequest)(nil),    // 1: tag.GetNextTagsRequest
	(*GetPrevTagsRequest)(nil),    // 2: tag.GetPrevTagsRequest
	(*Tag)(nil),                   // 3: tag.Tag
	(*Article)(nil),               // 4: tag.Article
	(*GetTagByIdResponse)(nil),    // 5: tag.GetTagByIdResponse
	(*GetAllTagsResponse)(nil),    // 6: tag.GetAllTagsResponse
	(*GetNextTagResponse)(nil),    // 7: tag.GetNextTagResponse
	(*GetPrevTagResponse)(nil),    // 8: tag.GetPrevTagResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_tag_proto_depIdxs = []int32{
	4,  // 0: tag.Tag.articles:type_name -> tag.Article
	9,  // 1: tag.Article.createdAt:type_name -> google.protobuf.Timestamp
	9,  // 2: tag.Article.updatedAt:type_name -> google.protobuf.Timestamp
	3,  // 3: tag.GetTagByIdResponse.tag:type_name -> tag.Tag
	3,  // 4: tag.GetAllTagsResponse.tags:type_name -> tag.Tag
	3,  // 5: tag.GetNextTagResponse.tags:type_name -> tag.Tag
	3,  // 6: tag.GetPrevTagResponse.tags:type_name -> tag.Tag
	0,  // 7: tag.TagService.GetTagById:input_type -> tag.GetTagByIdRequest
	10, // 8: tag.TagService.GetAllTags:input_type -> google.protobuf.Empty
	1,  // 9: tag.TagService.GetNextTags:input_type -> tag.GetNextTagsRequest
	2,  // 10: tag.TagService.GetPrevTags:input_type -> tag.GetPrevTagsRequest
	5,  // 11: tag.TagService.GetTagById:output_type -> tag.GetTagByIdResponse
	6,  // 12: tag.TagService.GetAllTags:output_type -> tag.GetAllTagsResponse
	7,  // 13: tag.TagService.GetNextTags:output_type -> tag.GetNextTagResponse
	8,  // 14: tag.TagService.GetPrevTags:output_type -> tag.GetPrevTagResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_tag_proto_init() }
func file_tag_proto_init() {
	if File_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagByIdRequest); i {
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
		file_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNextTagsRequest); i {
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
		file_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrevTagsRequest); i {
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
		file_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagByIdResponse); i {
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
		file_tag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTagsResponse); i {
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
		file_tag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNextTagResponse); i {
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
		file_tag_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrevTagResponse); i {
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
	file_tag_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_tag_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tag_proto_goTypes,
		DependencyIndexes: file_tag_proto_depIdxs,
		MessageInfos:      file_tag_proto_msgTypes,
	}.Build()
	File_tag_proto = out.File
	file_tag_proto_rawDesc = nil
	file_tag_proto_goTypes = nil
	file_tag_proto_depIdxs = nil
}
