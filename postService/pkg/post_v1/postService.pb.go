// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: post_v1/postService.proto

package post_v1

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

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId string `protobuf:"bytes,3,opt,name=authorId,proto3" json:"authorId,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_v1_postService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_postService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_postService_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreatePostRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type UpdatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   int64  `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId string `protobuf:"bytes,4,opt,name=authorId,proto3" json:"authorId,omitempty"`
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_v1_postService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_postService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_postService_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePostRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *UpdatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdatePostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdatePostRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   int64  `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=authorId,proto3" json:"authorId,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_v1_postService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_postService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_postService_proto_rawDescGZIP(), []int{2}
}

func (x *DeletePostRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *DeletePostRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type GetPostByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   int64  `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=authorId,proto3" json:"authorId,omitempty"`
}

func (x *GetPostByIdRequest) Reset() {
	*x = GetPostByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_v1_postService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByIdRequest) ProtoMessage() {}

func (x *GetPostByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_postService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostByIdRequest.ProtoReflect.Descriptor instead.
func (*GetPostByIdRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_postService_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostByIdRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *GetPostByIdRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type GetListPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	AuthorId string `protobuf:"bytes,3,opt,name=authorId,proto3" json:"authorId,omitempty"`
}

func (x *GetListPostsRequest) Reset() {
	*x = GetListPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_v1_postService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListPostsRequest) ProtoMessage() {}

func (x *GetListPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_postService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListPostsRequest.ProtoReflect.Descriptor instead.
func (*GetListPostsRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_postService_proto_rawDescGZIP(), []int{4}
}

func (x *GetListPostsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListPostsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetListPostsRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type PostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   int64  `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId string `protobuf:"bytes,4,opt,name=authorId,proto3" json:"authorId,omitempty"`
}

func (x *PostResponse) Reset() {
	*x = PostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_v1_postService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponse) ProtoMessage() {}

func (x *PostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_postService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostResponse.ProtoReflect.Descriptor instead.
func (*PostResponse) Descriptor() ([]byte, []int) {
	return file_post_v1_postService_proto_rawDescGZIP(), []int{5}
}

func (x *PostResponse) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *PostResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PostResponse) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type GetListPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts    []*PostResponse `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	NextPage int32           `protobuf:"varint,2,opt,name=nextPage,proto3" json:"nextPage,omitempty"`
}

func (x *GetListPostsResponse) Reset() {
	*x = GetListPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_v1_postService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListPostsResponse) ProtoMessage() {}

func (x *GetListPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_postService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListPostsResponse.ProtoReflect.Descriptor instead.
func (*GetListPostsResponse) Descriptor() ([]byte, []int) {
	return file_post_v1_postService_proto_rawDescGZIP(), []int{6}
}

func (x *GetListPostsResponse) GetPosts() []*PostResponse {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *GetListPostsResponse) GetNextPage() int32 {
	if x != nil {
		return x.NextPage
	}
	return 0
}

var File_post_v1_postService_proto protoreflect.FileDescriptor

var file_post_v1_postService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x76, 0x31, 0x22, 0x5f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x77, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x47,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x22, 0x61, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x32, 0xe0, 0x02, 0x0a, 0x0b, 0x50, 0x6f,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12,
	0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x43, 0x5a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x72, 0x79, 0x2f, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_v1_postService_proto_rawDescOnce sync.Once
	file_post_v1_postService_proto_rawDescData = file_post_v1_postService_proto_rawDesc
)

func file_post_v1_postService_proto_rawDescGZIP() []byte {
	file_post_v1_postService_proto_rawDescOnce.Do(func() {
		file_post_v1_postService_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_v1_postService_proto_rawDescData)
	})
	return file_post_v1_postService_proto_rawDescData
}

var file_post_v1_postService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_post_v1_postService_proto_goTypes = []interface{}{
	(*CreatePostRequest)(nil),    // 0: post_v1.CreatePostRequest
	(*UpdatePostRequest)(nil),    // 1: post_v1.UpdatePostRequest
	(*DeletePostRequest)(nil),    // 2: post_v1.DeletePostRequest
	(*GetPostByIdRequest)(nil),   // 3: post_v1.GetPostByIdRequest
	(*GetListPostsRequest)(nil),  // 4: post_v1.GetListPostsRequest
	(*PostResponse)(nil),         // 5: post_v1.PostResponse
	(*GetListPostsResponse)(nil), // 6: post_v1.GetListPostsResponse
}
var file_post_v1_postService_proto_depIdxs = []int32{
	5, // 0: post_v1.GetListPostsResponse.posts:type_name -> post_v1.PostResponse
	0, // 1: post_v1.PostService.CreatePost:input_type -> post_v1.CreatePostRequest
	1, // 2: post_v1.PostService.UpdatePost:input_type -> post_v1.UpdatePostRequest
	2, // 3: post_v1.PostService.DeletePost:input_type -> post_v1.DeletePostRequest
	3, // 4: post_v1.PostService.GetPostById:input_type -> post_v1.GetPostByIdRequest
	4, // 5: post_v1.PostService.GetListPosts:input_type -> post_v1.GetListPostsRequest
	5, // 6: post_v1.PostService.CreatePost:output_type -> post_v1.PostResponse
	5, // 7: post_v1.PostService.UpdatePost:output_type -> post_v1.PostResponse
	5, // 8: post_v1.PostService.DeletePost:output_type -> post_v1.PostResponse
	5, // 9: post_v1.PostService.GetPostById:output_type -> post_v1.PostResponse
	6, // 10: post_v1.PostService.GetListPosts:output_type -> post_v1.GetListPostsResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_post_v1_postService_proto_init() }
func file_post_v1_postService_proto_init() {
	if File_post_v1_postService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_v1_postService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
		file_post_v1_postService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostRequest); i {
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
		file_post_v1_postService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
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
		file_post_v1_postService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostByIdRequest); i {
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
		file_post_v1_postService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListPostsRequest); i {
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
		file_post_v1_postService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostResponse); i {
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
		file_post_v1_postService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListPostsResponse); i {
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
			RawDescriptor: file_post_v1_postService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_v1_postService_proto_goTypes,
		DependencyIndexes: file_post_v1_postService_proto_depIdxs,
		MessageInfos:      file_post_v1_postService_proto_msgTypes,
	}.Build()
	File_post_v1_postService_proto = out.File
	file_post_v1_postService_proto_rawDesc = nil
	file_post_v1_postService_proto_goTypes = nil
	file_post_v1_postService_proto_depIdxs = nil
}
