// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: blogging_event.proto

package bloggingevent

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

type CreateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body         string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	ThumbnailUrl string   `protobuf:"bytes,3,opt,name=thumbnailUrl,proto3" json:"thumbnailUrl,omitempty"`
	TagNames     []string `protobuf:"bytes,4,rep,name=tagNames,proto3" json:"tagNames,omitempty"`
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_blogging_event_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArticleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateArticleRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CreateArticleRequest) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *CreateArticleRequest) GetTagNames() []string {
	if x != nil {
		return x.TagNames
	}
	return nil
}

type UpdateArticleTitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateArticleTitleRequest) Reset() {
	*x = UpdateArticleTitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleTitleRequest) ProtoMessage() {}

func (x *UpdateArticleTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleTitleRequest.ProtoReflect.Descriptor instead.
func (*UpdateArticleTitleRequest) Descriptor() ([]byte, []int) {
	return file_blogging_event_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateArticleTitleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateArticleTitleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateArticleBodyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *UpdateArticleBodyRequest) Reset() {
	*x = UpdateArticleBodyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleBodyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleBodyRequest) ProtoMessage() {}

func (x *UpdateArticleBodyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleBodyRequest.ProtoReflect.Descriptor instead.
func (*UpdateArticleBodyRequest) Descriptor() ([]byte, []int) {
	return file_blogging_event_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateArticleBodyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateArticleBodyRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type UpdateArticleThumbnailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ThumbnailUrl string `protobuf:"bytes,2,opt,name=thumbnailUrl,proto3" json:"thumbnailUrl,omitempty"`
}

func (x *UpdateArticleThumbnailRequest) Reset() {
	*x = UpdateArticleThumbnailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleThumbnailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleThumbnailRequest) ProtoMessage() {}

func (x *UpdateArticleThumbnailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleThumbnailRequest.ProtoReflect.Descriptor instead.
func (*UpdateArticleThumbnailRequest) Descriptor() ([]byte, []int) {
	return file_blogging_event_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateArticleThumbnailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateArticleThumbnailRequest) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

type AttachTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TagNames []string `protobuf:"bytes,2,rep,name=tagNames,proto3" json:"tagNames,omitempty"`
}

func (x *AttachTagsRequest) Reset() {
	*x = AttachTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachTagsRequest) ProtoMessage() {}

func (x *AttachTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachTagsRequest.ProtoReflect.Descriptor instead.
func (*AttachTagsRequest) Descriptor() ([]byte, []int) {
	return file_blogging_event_proto_rawDescGZIP(), []int{4}
}

func (x *AttachTagsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AttachTagsRequest) GetTagNames() []string {
	if x != nil {
		return x.TagNames
	}
	return nil
}

type DetachTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TagNames []string `protobuf:"bytes,2,rep,name=tagNames,proto3" json:"tagNames,omitempty"`
}

func (x *DetachTagsRequest) Reset() {
	*x = DetachTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetachTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetachTagsRequest) ProtoMessage() {}

func (x *DetachTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetachTagsRequest.ProtoReflect.Descriptor instead.
func (*DetachTagsRequest) Descriptor() ([]byte, []int) {
	return file_blogging_event_proto_rawDescGZIP(), []int{5}
}

func (x *DetachTagsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DetachTagsRequest) GetTagNames() []string {
	if x != nil {
		return x.TagNames
	}
	return nil
}

type BloggingEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId string `protobuf:"bytes,1,opt,name=articleId,proto3" json:"articleId,omitempty"`
	EventId   string `protobuf:"bytes,2,opt,name=eventId,proto3" json:"eventId,omitempty"`
}

func (x *BloggingEventResponse) Reset() {
	*x = BloggingEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloggingEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloggingEventResponse) ProtoMessage() {}

func (x *BloggingEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloggingEventResponse.ProtoReflect.Descriptor instead.
func (*BloggingEventResponse) Descriptor() ([]byte, []int) {
	return file_blogging_event_proto_rawDescGZIP(), []int{6}
}

func (x *BloggingEventResponse) GetArticleId() string {
	if x != nil {
		return x.ArticleId
	}
	return ""
}

func (x *BloggingEventResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*UploadImageRequest_Meta
	//	*UploadImageRequest_Data
	Value isUploadImageRequest_Value `protobuf_oneof:"value"`
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_blogging_event_proto_rawDescGZIP(), []int{7}
}

func (m *UploadImageRequest) GetValue() isUploadImageRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *UploadImageRequest) GetMeta() *Meta {
	if x, ok := x.GetValue().(*UploadImageRequest_Meta); ok {
		return x.Meta
	}
	return nil
}

func (x *UploadImageRequest) GetData() []byte {
	if x, ok := x.GetValue().(*UploadImageRequest_Data); ok {
		return x.Data
	}
	return nil
}

type isUploadImageRequest_Value interface {
	isUploadImageRequest_Value()
}

type UploadImageRequest_Meta struct {
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type UploadImageRequest_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*UploadImageRequest_Meta) isUploadImageRequest_Value() {}

func (*UploadImageRequest_Data) isUploadImageRequest_Value() {}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ContentType string `protobuf:"bytes,2,opt,name=contentType,proto3" json:"contentType,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_blogging_event_proto_rawDescGZIP(), []int{8}
}

func (x *Meta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Meta) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type UploadImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Url     *string `protobuf:"bytes,2,opt,name=url,proto3,oneof" json:"url,omitempty"`
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogging_event_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blogging_event_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_blogging_event_proto_rawDescGZIP(), []int{9}
}

func (x *UploadImageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UploadImageResponse) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

var File_blogging_event_proto protoreflect.FileDescriptor

var file_blogging_event_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3e, 0x0a, 0x18,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x53, 0x0a, 0x1d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72,
	0x6c, 0x22, 0x3f, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x54, 0x61, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x22, 0x3f, 0x0a, 0x11, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x54, 0x61, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x22, 0x4f, 0x0a, 0x15, 0x42, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x07, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3c, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x4e, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f,
	0x75, 0x72, 0x6c, 0x32, 0xbc, 0x05, 0x0a, 0x14, 0x42, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x24, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x29, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x64, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x28, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x12, 0x2d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x54, 0x61, 0x67, 0x73, 0x12, 0x21, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x54, 0x61,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x56, 0x0a, 0x0a, 0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x54, 0x61, 0x67, 0x73, 0x12, 0x21,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x65, 0x74, 0x61, 0x63, 0x68, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blogging_event_proto_rawDescOnce sync.Once
	file_blogging_event_proto_rawDescData = file_blogging_event_proto_rawDesc
)

func file_blogging_event_proto_rawDescGZIP() []byte {
	file_blogging_event_proto_rawDescOnce.Do(func() {
		file_blogging_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_blogging_event_proto_rawDescData)
	})
	return file_blogging_event_proto_rawDescData
}

var file_blogging_event_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_blogging_event_proto_goTypes = []interface{}{
	(*CreateArticleRequest)(nil),          // 0: blogging_event.CreateArticleRequest
	(*UpdateArticleTitleRequest)(nil),     // 1: blogging_event.UpdateArticleTitleRequest
	(*UpdateArticleBodyRequest)(nil),      // 2: blogging_event.UpdateArticleBodyRequest
	(*UpdateArticleThumbnailRequest)(nil), // 3: blogging_event.UpdateArticleThumbnailRequest
	(*AttachTagsRequest)(nil),             // 4: blogging_event.AttachTagsRequest
	(*DetachTagsRequest)(nil),             // 5: blogging_event.DetachTagsRequest
	(*BloggingEventResponse)(nil),         // 6: blogging_event.BloggingEventResponse
	(*UploadImageRequest)(nil),            // 7: blogging_event.UploadImageRequest
	(*Meta)(nil),                          // 8: blogging_event.Meta
	(*UploadImageResponse)(nil),           // 9: blogging_event.UploadImageResponse
}
var file_blogging_event_proto_depIdxs = []int32{
	8, // 0: blogging_event.UploadImageRequest.meta:type_name -> blogging_event.Meta
	0, // 1: blogging_event.BloggingEventService.CreateArticle:input_type -> blogging_event.CreateArticleRequest
	1, // 2: blogging_event.BloggingEventService.UpdateArticleTitle:input_type -> blogging_event.UpdateArticleTitleRequest
	2, // 3: blogging_event.BloggingEventService.UpdateArticleBody:input_type -> blogging_event.UpdateArticleBodyRequest
	3, // 4: blogging_event.BloggingEventService.UpdateArticleThumbnail:input_type -> blogging_event.UpdateArticleThumbnailRequest
	4, // 5: blogging_event.BloggingEventService.AttachTags:input_type -> blogging_event.AttachTagsRequest
	5, // 6: blogging_event.BloggingEventService.DetachTags:input_type -> blogging_event.DetachTagsRequest
	7, // 7: blogging_event.BloggingEventService.UploadImage:input_type -> blogging_event.UploadImageRequest
	6, // 8: blogging_event.BloggingEventService.CreateArticle:output_type -> blogging_event.BloggingEventResponse
	6, // 9: blogging_event.BloggingEventService.UpdateArticleTitle:output_type -> blogging_event.BloggingEventResponse
	6, // 10: blogging_event.BloggingEventService.UpdateArticleBody:output_type -> blogging_event.BloggingEventResponse
	6, // 11: blogging_event.BloggingEventService.UpdateArticleThumbnail:output_type -> blogging_event.BloggingEventResponse
	6, // 12: blogging_event.BloggingEventService.AttachTags:output_type -> blogging_event.BloggingEventResponse
	6, // 13: blogging_event.BloggingEventService.DetachTags:output_type -> blogging_event.BloggingEventResponse
	9, // 14: blogging_event.BloggingEventService.UploadImage:output_type -> blogging_event.UploadImageResponse
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blogging_event_proto_init() }
func file_blogging_event_proto_init() {
	if File_blogging_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blogging_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest); i {
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
		file_blogging_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleTitleRequest); i {
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
		file_blogging_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleBodyRequest); i {
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
		file_blogging_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleThumbnailRequest); i {
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
		file_blogging_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachTagsRequest); i {
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
		file_blogging_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetachTagsRequest); i {
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
		file_blogging_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloggingEventResponse); i {
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
		file_blogging_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageRequest); i {
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
		file_blogging_event_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_blogging_event_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageResponse); i {
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
	file_blogging_event_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*UploadImageRequest_Meta)(nil),
		(*UploadImageRequest_Data)(nil),
	}
	file_blogging_event_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blogging_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blogging_event_proto_goTypes,
		DependencyIndexes: file_blogging_event_proto_depIdxs,
		MessageInfos:      file_blogging_event_proto_msgTypes,
	}.Build()
	File_blogging_event_proto = out.File
	file_blogging_event_proto_rawDesc = nil
	file_blogging_event_proto_goTypes = nil
	file_blogging_event_proto_depIdxs = nil
}
