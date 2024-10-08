// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: directorytypes/bertydirectory.proto

package directorytypes

import (
	_ "berty.tech/weshnet/v2/pkg/protocoltypes"
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

type Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Register) Reset() {
	*x = Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directorytypes_bertydirectory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register) ProtoMessage() {}

func (x *Register) ProtoReflect() protoreflect.Message {
	mi := &file_directorytypes_bertydirectory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register.ProtoReflect.Descriptor instead.
func (*Register) Descriptor() ([]byte, []int) {
	return file_directorytypes_bertydirectory_proto_rawDescGZIP(), []int{0}
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directorytypes_bertydirectory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_directorytypes_bertydirectory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_directorytypes_bertydirectory_proto_rawDescGZIP(), []int{1}
}

type Unregister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Unregister) Reset() {
	*x = Unregister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directorytypes_bertydirectory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unregister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unregister) ProtoMessage() {}

func (x *Unregister) ProtoReflect() protoreflect.Message {
	mi := &file_directorytypes_bertydirectory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unregister.ProtoReflect.Descriptor instead.
func (*Unregister) Descriptor() ([]byte, []int) {
	return file_directorytypes_bertydirectory_proto_rawDescGZIP(), []int{2}
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoryIdentifier  string `protobuf:"bytes,1,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty" gorm:"index;primaryKey;autoIncrement:false"`
	DirectoryRecordToken string `protobuf:"bytes,2,opt,name=directory_record_token,json=directoryRecordToken,proto3" json:"directory_record_token,omitempty"`
	ExpiresAt            int64  `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	LockedUntil          int64  `protobuf:"varint,4,opt,name=locked_until,json=lockedUntil,proto3" json:"locked_until,omitempty"`
	UnregisterToken      string `protobuf:"bytes,5,opt,name=unregister_token,json=unregisterToken,proto3" json:"unregister_token,omitempty"`
	AccountUri           string `protobuf:"bytes,6,opt,name=account_uri,json=accountUri,proto3" json:"account_uri,omitempty"`
	VerifiedCredential   []byte `protobuf:"bytes,7,opt,name=verified_credential,json=verifiedCredential,proto3" json:"verified_credential,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directorytypes_bertydirectory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_directorytypes_bertydirectory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_directorytypes_bertydirectory_proto_rawDescGZIP(), []int{3}
}

func (x *Record) GetDirectoryIdentifier() string {
	if x != nil {
		return x.DirectoryIdentifier
	}
	return ""
}

func (x *Record) GetDirectoryRecordToken() string {
	if x != nil {
		return x.DirectoryRecordToken
	}
	return ""
}

func (x *Record) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *Record) GetLockedUntil() int64 {
	if x != nil {
		return x.LockedUntil
	}
	return 0
}

func (x *Record) GetUnregisterToken() string {
	if x != nil {
		return x.UnregisterToken
	}
	return ""
}

func (x *Record) GetAccountUri() string {
	if x != nil {
		return x.AccountUri
	}
	return ""
}

func (x *Record) GetVerifiedCredential() []byte {
	if x != nil {
		return x.VerifiedCredential
	}
	return nil
}

type Register_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerifiedCredential      []byte `protobuf:"bytes,1,opt,name=verified_credential,json=verifiedCredential,proto3" json:"verified_credential,omitempty"`
	ExpirationDate          int64  `protobuf:"varint,2,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	LockedUntilDate         int64  `protobuf:"varint,3,opt,name=locked_until_date,json=lockedUntilDate,proto3" json:"locked_until_date,omitempty"`
	AccountUri              string `protobuf:"bytes,4,opt,name=account_uri,json=accountUri,proto3" json:"account_uri,omitempty"`
	OverwriteExistingRecord bool   `protobuf:"varint,5,opt,name=overwrite_existing_record,json=overwriteExistingRecord,proto3" json:"overwrite_existing_record,omitempty"`
}

func (x *Register_Request) Reset() {
	*x = Register_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directorytypes_bertydirectory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register_Request) ProtoMessage() {}

func (x *Register_Request) ProtoReflect() protoreflect.Message {
	mi := &file_directorytypes_bertydirectory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register_Request.ProtoReflect.Descriptor instead.
func (*Register_Request) Descriptor() ([]byte, []int) {
	return file_directorytypes_bertydirectory_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Register_Request) GetVerifiedCredential() []byte {
	if x != nil {
		return x.VerifiedCredential
	}
	return nil
}

func (x *Register_Request) GetExpirationDate() int64 {
	if x != nil {
		return x.ExpirationDate
	}
	return 0
}

func (x *Register_Request) GetLockedUntilDate() int64 {
	if x != nil {
		return x.LockedUntilDate
	}
	return 0
}

func (x *Register_Request) GetAccountUri() string {
	if x != nil {
		return x.AccountUri
	}
	return ""
}

func (x *Register_Request) GetOverwriteExistingRecord() bool {
	if x != nil {
		return x.OverwriteExistingRecord
	}
	return false
}

type Register_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoryRecordToken string `protobuf:"bytes,1,opt,name=directory_record_token,json=directoryRecordToken,proto3" json:"directory_record_token,omitempty"`
	DirectoryIdentifier  string `protobuf:"bytes,2,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty"`
	ExpirationDate       int64  `protobuf:"varint,3,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	UnregisterToken      string `protobuf:"bytes,4,opt,name=unregister_token,json=unregisterToken,proto3" json:"unregister_token,omitempty"`
}

func (x *Register_Reply) Reset() {
	*x = Register_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directorytypes_bertydirectory_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register_Reply) ProtoMessage() {}

func (x *Register_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_directorytypes_bertydirectory_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register_Reply.ProtoReflect.Descriptor instead.
func (*Register_Reply) Descriptor() ([]byte, []int) {
	return file_directorytypes_bertydirectory_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Register_Reply) GetDirectoryRecordToken() string {
	if x != nil {
		return x.DirectoryRecordToken
	}
	return ""
}

func (x *Register_Reply) GetDirectoryIdentifier() string {
	if x != nil {
		return x.DirectoryIdentifier
	}
	return ""
}

func (x *Register_Reply) GetExpirationDate() int64 {
	if x != nil {
		return x.ExpirationDate
	}
	return 0
}

func (x *Register_Reply) GetUnregisterToken() string {
	if x != nil {
		return x.UnregisterToken
	}
	return ""
}

type Query_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoryIdentifiers []string `protobuf:"bytes,1,rep,name=directory_identifiers,json=directoryIdentifiers,proto3" json:"directory_identifiers,omitempty"`
}

func (x *Query_Request) Reset() {
	*x = Query_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directorytypes_bertydirectory_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query_Request) ProtoMessage() {}

func (x *Query_Request) ProtoReflect() protoreflect.Message {
	mi := &file_directorytypes_bertydirectory_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query_Request.ProtoReflect.Descriptor instead.
func (*Query_Request) Descriptor() ([]byte, []int) {
	return file_directorytypes_bertydirectory_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Query_Request) GetDirectoryIdentifiers() []string {
	if x != nil {
		return x.DirectoryIdentifiers
	}
	return nil
}

type Query_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoryIdentifier string `protobuf:"bytes,1,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty"`
	ExpiresAt           int64  `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	AccountUri          string `protobuf:"bytes,3,opt,name=account_uri,json=accountUri,proto3" json:"account_uri,omitempty"`
	VerifiedCredential  []byte `protobuf:"bytes,4,opt,name=verified_credential,json=verifiedCredential,proto3" json:"verified_credential,omitempty"`
}

func (x *Query_Reply) Reset() {
	*x = Query_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directorytypes_bertydirectory_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query_Reply) ProtoMessage() {}

func (x *Query_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_directorytypes_bertydirectory_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query_Reply.ProtoReflect.Descriptor instead.
func (*Query_Reply) Descriptor() ([]byte, []int) {
	return file_directorytypes_bertydirectory_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Query_Reply) GetDirectoryIdentifier() string {
	if x != nil {
		return x.DirectoryIdentifier
	}
	return ""
}

func (x *Query_Reply) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *Query_Reply) GetAccountUri() string {
	if x != nil {
		return x.AccountUri
	}
	return ""
}

func (x *Query_Reply) GetVerifiedCredential() []byte {
	if x != nil {
		return x.VerifiedCredential
	}
	return nil
}

type Unregister_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoryIdentifier  string `protobuf:"bytes,1,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty"`
	DirectoryRecordToken string `protobuf:"bytes,2,opt,name=directory_record_token,json=directoryRecordToken,proto3" json:"directory_record_token,omitempty"`
	UnregisterToken      string `protobuf:"bytes,3,opt,name=unregister_token,json=unregisterToken,proto3" json:"unregister_token,omitempty"` // when not known (ie. device lost) the user can either wait for the record to expire or register again (if the record is still present but unlocked) and then unregister
}

func (x *Unregister_Request) Reset() {
	*x = Unregister_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directorytypes_bertydirectory_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unregister_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unregister_Request) ProtoMessage() {}

func (x *Unregister_Request) ProtoReflect() protoreflect.Message {
	mi := &file_directorytypes_bertydirectory_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unregister_Request.ProtoReflect.Descriptor instead.
func (*Unregister_Request) Descriptor() ([]byte, []int) {
	return file_directorytypes_bertydirectory_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Unregister_Request) GetDirectoryIdentifier() string {
	if x != nil {
		return x.DirectoryIdentifier
	}
	return ""
}

func (x *Unregister_Request) GetDirectoryRecordToken() string {
	if x != nil {
		return x.DirectoryRecordToken
	}
	return ""
}

func (x *Unregister_Request) GetUnregisterToken() string {
	if x != nil {
		return x.UnregisterToken
	}
	return ""
}

type Unregister_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Unregister_Reply) Reset() {
	*x = Unregister_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directorytypes_bertydirectory_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unregister_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unregister_Reply) ProtoMessage() {}

func (x *Unregister_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_directorytypes_bertydirectory_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unregister_Reply.ProtoReflect.Descriptor instead.
func (*Unregister_Reply) Descriptor() ([]byte, []int) {
	return file_directorytypes_bertydirectory_proto_rawDescGZIP(), []int{2, 1}
}

var File_directorytypes_bertydirectory_proto protoreflect.FileDescriptor

var file_directorytypes_bertydirectory_proto_rawDesc = []byte{
	0x0a, 0x23, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x62, 0x65, 0x72, 0x74, 0x79, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x03, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x1a, 0xec, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x27, 0x0a,
	0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x55, 0x72, 0x69, 0x12, 0x3a, 0x0a, 0x19, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a,
	0xc4, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x31, 0x0a, 0x14, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x75,
	0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf5, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x3e, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x15, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73,
	0x1a, 0xab, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x14, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x12, 0x2f, 0x0a,
	0x13, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0xb5,
	0x01, 0x0a, 0x0a, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x9d, 0x01,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x14, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x16,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x75, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x6e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x07, 0x0a,
	0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xe2, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x63, 0x0a, 0x14, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x30, 0x9a, 0x84, 0x9e, 0x03, 0x2b, 0x67, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x3b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x3b, 0x61, 0x75, 0x74,
	0x6f, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x66, 0x61, 0x6c, 0x73, 0x65,
	0x22, 0x52, 0x13, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x29,
	0x0a, 0x10, 0x75, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x6e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x12, 0x2f, 0x0a, 0x13, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x32, 0x93, 0x02, 0x0a, 0x10,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x54, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x62,
	0x65, 0x72, 0x74, 0x79, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4d, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x21, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x30, 0x01, 0x12, 0x5a, 0x0a, 0x0a, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x65,
	0x72, 0x74, 0x79, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x28, 0x5a, 0x26, 0x62, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f,
	0x62, 0x65, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_directorytypes_bertydirectory_proto_rawDescOnce sync.Once
	file_directorytypes_bertydirectory_proto_rawDescData = file_directorytypes_bertydirectory_proto_rawDesc
)

func file_directorytypes_bertydirectory_proto_rawDescGZIP() []byte {
	file_directorytypes_bertydirectory_proto_rawDescOnce.Do(func() {
		file_directorytypes_bertydirectory_proto_rawDescData = protoimpl.X.CompressGZIP(file_directorytypes_bertydirectory_proto_rawDescData)
	})
	return file_directorytypes_bertydirectory_proto_rawDescData
}

var file_directorytypes_bertydirectory_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_directorytypes_bertydirectory_proto_goTypes = []any{
	(*Register)(nil),           // 0: berty.directory.v1.Register
	(*Query)(nil),              // 1: berty.directory.v1.Query
	(*Unregister)(nil),         // 2: berty.directory.v1.Unregister
	(*Record)(nil),             // 3: berty.directory.v1.Record
	(*Register_Request)(nil),   // 4: berty.directory.v1.Register.Request
	(*Register_Reply)(nil),     // 5: berty.directory.v1.Register.Reply
	(*Query_Request)(nil),      // 6: berty.directory.v1.Query.Request
	(*Query_Reply)(nil),        // 7: berty.directory.v1.Query.Reply
	(*Unregister_Request)(nil), // 8: berty.directory.v1.Unregister.Request
	(*Unregister_Reply)(nil),   // 9: berty.directory.v1.Unregister.Reply
}
var file_directorytypes_bertydirectory_proto_depIdxs = []int32{
	4, // 0: berty.directory.v1.DirectoryService.Register:input_type -> berty.directory.v1.Register.Request
	6, // 1: berty.directory.v1.DirectoryService.Query:input_type -> berty.directory.v1.Query.Request
	8, // 2: berty.directory.v1.DirectoryService.Unregister:input_type -> berty.directory.v1.Unregister.Request
	5, // 3: berty.directory.v1.DirectoryService.Register:output_type -> berty.directory.v1.Register.Reply
	7, // 4: berty.directory.v1.DirectoryService.Query:output_type -> berty.directory.v1.Query.Reply
	9, // 5: berty.directory.v1.DirectoryService.Unregister:output_type -> berty.directory.v1.Unregister.Reply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_directorytypes_bertydirectory_proto_init() }
func file_directorytypes_bertydirectory_proto_init() {
	if File_directorytypes_bertydirectory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_directorytypes_bertydirectory_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Register); i {
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
		file_directorytypes_bertydirectory_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Query); i {
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
		file_directorytypes_bertydirectory_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Unregister); i {
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
		file_directorytypes_bertydirectory_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Record); i {
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
		file_directorytypes_bertydirectory_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Register_Request); i {
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
		file_directorytypes_bertydirectory_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Register_Reply); i {
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
		file_directorytypes_bertydirectory_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Query_Request); i {
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
		file_directorytypes_bertydirectory_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Query_Reply); i {
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
		file_directorytypes_bertydirectory_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Unregister_Request); i {
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
		file_directorytypes_bertydirectory_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Unregister_Reply); i {
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
			RawDescriptor: file_directorytypes_bertydirectory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_directorytypes_bertydirectory_proto_goTypes,
		DependencyIndexes: file_directorytypes_bertydirectory_proto_depIdxs,
		MessageInfos:      file_directorytypes_bertydirectory_proto_msgTypes,
	}.Build()
	File_directorytypes_bertydirectory_proto = out.File
	file_directorytypes_bertydirectory_proto_rawDesc = nil
	file_directorytypes_bertydirectory_proto_goTypes = nil
	file_directorytypes_bertydirectory_proto_depIdxs = nil
}
