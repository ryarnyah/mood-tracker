// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: mood.proto

package mood

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
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

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record  uint32 `protobuf:"varint,1,opt,name=record,proto3" json:"record,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetRecord() uint32 {
	if x != nil {
		return x.Record
	}
	return 0
}

func (x *Entry) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type EntryWithDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record      uint32               `protobuf:"varint,1,opt,name=record,proto3" json:"record,omitempty"`
	Comment     string               `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	RecordEntry *timestamp.Timestamp `protobuf:"bytes,3,opt,name=record_entry,json=recordEntry,proto3" json:"record_entry,omitempty"`
}

func (x *EntryWithDate) Reset() {
	*x = EntryWithDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryWithDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryWithDate) ProtoMessage() {}

func (x *EntryWithDate) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryWithDate.ProtoReflect.Descriptor instead.
func (*EntryWithDate) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{1}
}

func (x *EntryWithDate) GetRecord() uint32 {
	if x != nil {
		return x.Record
	}
	return 0
}

func (x *EntryWithDate) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *EntryWithDate) GetRecordEntry() *timestamp.Timestamp {
	if x != nil {
		return x.RecordEntry
	}
	return nil
}

type AddEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry           *Entry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	MoodId          int64  `protobuf:"varint,2,opt,name=mood_id,json=moodId,proto3" json:"mood_id,omitempty"`
	EntryAccessCode string `protobuf:"bytes,3,opt,name=entry_access_code,json=entryAccessCode,proto3" json:"entry_access_code,omitempty"`
}

func (x *AddEntryRequest) Reset() {
	*x = AddEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEntryRequest) ProtoMessage() {}

func (x *AddEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEntryRequest.ProtoReflect.Descriptor instead.
func (*AddEntryRequest) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{2}
}

func (x *AddEntryRequest) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *AddEntryRequest) GetMoodId() int64 {
	if x != nil {
		return x.MoodId
	}
	return 0
}

func (x *AddEntryRequest) GetEntryAccessCode() string {
	if x != nil {
		return x.EntryAccessCode
	}
	return ""
}

type AddEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddEntryResponse) Reset() {
	*x = AddEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEntryResponse) ProtoMessage() {}

func (x *AddEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEntryResponse.ProtoReflect.Descriptor instead.
func (*AddEntryResponse) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{3}
}

type GetMoodFromEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoodId          int64  `protobuf:"varint,2,opt,name=mood_id,json=moodId,proto3" json:"mood_id,omitempty"`
	EntryAccessCode string `protobuf:"bytes,3,opt,name=entry_access_code,json=entryAccessCode,proto3" json:"entry_access_code,omitempty"`
}

func (x *GetMoodFromEntryRequest) Reset() {
	*x = GetMoodFromEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoodFromEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoodFromEntryRequest) ProtoMessage() {}

func (x *GetMoodFromEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoodFromEntryRequest.ProtoReflect.Descriptor instead.
func (*GetMoodFromEntryRequest) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{4}
}

func (x *GetMoodFromEntryRequest) GetMoodId() int64 {
	if x != nil {
		return x.MoodId
	}
	return 0
}

func (x *GetMoodFromEntryRequest) GetEntryAccessCode() string {
	if x != nil {
		return x.EntryAccessCode
	}
	return ""
}

type GetMoodFromEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetMoodFromEntryResponse) Reset() {
	*x = GetMoodFromEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoodFromEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoodFromEntryResponse) ProtoMessage() {}

func (x *GetMoodFromEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoodFromEntryResponse.ProtoReflect.Descriptor instead.
func (*GetMoodFromEntryResponse) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{5}
}

func (x *GetMoodFromEntryResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetMoodFromEntryResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type GetMoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoodId         int64  `protobuf:"varint,1,opt,name=mood_id,json=moodId,proto3" json:"mood_id,omitempty"`
	MoodAccessCode string `protobuf:"bytes,2,opt,name=mood_access_code,json=moodAccessCode,proto3" json:"mood_access_code,omitempty"`
}

func (x *GetMoodRequest) Reset() {
	*x = GetMoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoodRequest) ProtoMessage() {}

func (x *GetMoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoodRequest.ProtoReflect.Descriptor instead.
func (*GetMoodRequest) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{6}
}

func (x *GetMoodRequest) GetMoodId() int64 {
	if x != nil {
		return x.MoodId
	}
	return 0
}

func (x *GetMoodRequest) GetMoodAccessCode() string {
	if x != nil {
		return x.MoodAccessCode
	}
	return ""
}

type RecordStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordEntry *timestamp.Timestamp `protobuf:"bytes,1,opt,name=record_entry,json=recordEntry,proto3" json:"record_entry,omitempty"`
	Count       int64                `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *RecordStat) Reset() {
	*x = RecordStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordStat) ProtoMessage() {}

func (x *RecordStat) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordStat.ProtoReflect.Descriptor instead.
func (*RecordStat) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{7}
}

func (x *RecordStat) GetRecordEntry() *timestamp.Timestamp {
	if x != nil {
		return x.RecordEntry
	}
	return nil
}

func (x *RecordStat) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type MoodStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record      uint32        `protobuf:"varint,1,opt,name=record,proto3" json:"record,omitempty"`
	RecordStats []*RecordStat `protobuf:"bytes,2,rep,name=record_stats,json=recordStats,proto3" json:"record_stats,omitempty"`
}

func (x *MoodStat) Reset() {
	*x = MoodStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoodStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoodStat) ProtoMessage() {}

func (x *MoodStat) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoodStat.ProtoReflect.Descriptor instead.
func (*MoodStat) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{8}
}

func (x *MoodStat) GetRecord() uint32 {
	if x != nil {
		return x.Record
	}
	return 0
}

func (x *MoodStat) GetRecordStats() []*RecordStat {
	if x != nil {
		return x.RecordStats
	}
	return nil
}

type GetMoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string           `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string           `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Entries []*EntryWithDate `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
	Stats   []*MoodStat      `protobuf:"bytes,4,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetMoodResponse) Reset() {
	*x = GetMoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoodResponse) ProtoMessage() {}

func (x *GetMoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoodResponse.ProtoReflect.Descriptor instead.
func (*GetMoodResponse) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{9}
}

func (x *GetMoodResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetMoodResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetMoodResponse) GetEntries() []*EntryWithDate {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *GetMoodResponse) GetStats() []*MoodStat {
	if x != nil {
		return x.Stats
	}
	return nil
}

type CreateMoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title                 string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content               string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	NumberOfRecordsNeeded uint32   `protobuf:"varint,3,opt,name=number_of_records_needed,json=numberOfRecordsNeeded,proto3" json:"number_of_records_needed,omitempty"`
	Emails                []string `protobuf:"bytes,4,rep,name=emails,proto3" json:"emails,omitempty"`
}

func (x *CreateMoodRequest) Reset() {
	*x = CreateMoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMoodRequest) ProtoMessage() {}

func (x *CreateMoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMoodRequest.ProtoReflect.Descriptor instead.
func (*CreateMoodRequest) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{10}
}

func (x *CreateMoodRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateMoodRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateMoodRequest) GetNumberOfRecordsNeeded() uint32 {
	if x != nil {
		return x.NumberOfRecordsNeeded
	}
	return 0
}

func (x *CreateMoodRequest) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

type CreateMoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoodId             int64    `protobuf:"varint,1,opt,name=mood_id,json=moodId,proto3" json:"mood_id,omitempty"`
	MoodAccessCode     string   `protobuf:"bytes,2,opt,name=mood_access_code,json=moodAccessCode,proto3" json:"mood_access_code,omitempty"`
	EntriesAccessCodes []string `protobuf:"bytes,3,rep,name=entries_access_codes,json=entriesAccessCodes,proto3" json:"entries_access_codes,omitempty"`
}

func (x *CreateMoodResponse) Reset() {
	*x = CreateMoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMoodResponse) ProtoMessage() {}

func (x *CreateMoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMoodResponse.ProtoReflect.Descriptor instead.
func (*CreateMoodResponse) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{11}
}

func (x *CreateMoodResponse) GetMoodId() int64 {
	if x != nil {
		return x.MoodId
	}
	return 0
}

func (x *CreateMoodResponse) GetMoodAccessCode() string {
	if x != nil {
		return x.MoodAccessCode
	}
	return ""
}

func (x *CreateMoodResponse) GetEntriesAccessCodes() []string {
	if x != nil {
		return x.EntriesAccessCodes
	}
	return nil
}

var File_mood_proto protoreflect.FileDescriptor

var file_mood_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77,
	0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x20,
	0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08,
	0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00, 0x18, 0x04, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x57, 0x69, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00, 0x18, 0x04, 0x52,
	0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81,
	0x01, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x05, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x07, 0x6d, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x06, 0x6d, 0x6f,
	0x6f, 0x64, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x11, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58, 0x01, 0x90, 0x01, 0x04, 0x52, 0x0f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x41,
	0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x71, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6f, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x07, 0x6d, 0x6f,
	0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x10, 0x00, 0x52, 0x06, 0x6d, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x11, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58, 0x01, 0x90, 0x01,
	0x04, 0x52, 0x0f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x5e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6f, 0x64, 0x46, 0x72, 0x6f,
	0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2,
	0xdf, 0x1f, 0x05, 0x58, 0x01, 0x78, 0x81, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x21, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x66, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x07, 0x6d, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x06, 0x6d,
	0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x10, 0x6d, 0x6f, 0x6f, 0x64, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58, 0x01, 0x90, 0x01, 0x04, 0x52, 0x0e, 0x6d, 0x6f, 0x6f, 0x64,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x61, 0x0a, 0x0a, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x52, 0x0a,
	0x08, 0x4d, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x2e, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x22, 0xa0, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58, 0x01, 0x78, 0x81, 0x01, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81, 0x04,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x22, 0xc4, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58,
	0x01, 0x78, 0x81, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf,
	0x1f, 0x03, 0x78, 0x81, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3f,
	0x0a, 0x18, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x18, 0x15, 0x52, 0x15, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x12,
	0xa9, 0x01, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x90, 0x01, 0xe2, 0xdf, 0x1f, 0x8b, 0x01, 0x0a, 0x84, 0x01, 0x5e, 0x5b, 0x61, 0x2d, 0x7a,
	0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2e, 0x21, 0x23, 0x24, 0x25, 0x25, 0x26, 0x27, 0x2a, 0x2b,
	0x2f, 0x3d, 0x3f, 0x5e, 0x5f, 0x7b, 0x7c, 0x7d, 0x7e, 0x2d, 0x5d, 0x2b, 0x40, 0x5b, 0x61, 0x2d,
	0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x41,
	0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d,
	0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x28, 0x3f, 0x3a, 0x5c, 0x2e, 0x5b,
	0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x61, 0x2d,
	0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b,
	0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x29, 0x2a, 0x24, 0x58,
	0x01, 0x68, 0x14, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x07, 0x6d, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x06, 0x6d, 0x6f, 0x6f,
	0x64, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x10, 0x6d, 0x6f, 0x6f, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2,
	0xdf, 0x1f, 0x05, 0x58, 0x01, 0x90, 0x01, 0x04, 0x52, 0x0e, 0x6d, 0x6f, 0x6f, 0x64, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x32, 0xed, 0x01, 0x0a, 0x04, 0x4d,
	0x6f, 0x6f, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6f,
	0x64, 0x46, 0x72, 0x6f, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x6f, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6f, 0x64, 0x46, 0x72,
	0x6f, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6f, 0x64, 0x12, 0x0f, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6f, 0x64, 0x12,
	0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x3b,
	0x6d, 0x6f, 0x6f, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mood_proto_rawDescOnce sync.Once
	file_mood_proto_rawDescData = file_mood_proto_rawDesc
)

func file_mood_proto_rawDescGZIP() []byte {
	file_mood_proto_rawDescOnce.Do(func() {
		file_mood_proto_rawDescData = protoimpl.X.CompressGZIP(file_mood_proto_rawDescData)
	})
	return file_mood_proto_rawDescData
}

var file_mood_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_mood_proto_goTypes = []interface{}{
	(*Entry)(nil),                    // 0: Entry
	(*EntryWithDate)(nil),            // 1: EntryWithDate
	(*AddEntryRequest)(nil),          // 2: AddEntryRequest
	(*AddEntryResponse)(nil),         // 3: AddEntryResponse
	(*GetMoodFromEntryRequest)(nil),  // 4: GetMoodFromEntryRequest
	(*GetMoodFromEntryResponse)(nil), // 5: GetMoodFromEntryResponse
	(*GetMoodRequest)(nil),           // 6: GetMoodRequest
	(*RecordStat)(nil),               // 7: RecordStat
	(*MoodStat)(nil),                 // 8: MoodStat
	(*GetMoodResponse)(nil),          // 9: GetMoodResponse
	(*CreateMoodRequest)(nil),        // 10: CreateMoodRequest
	(*CreateMoodResponse)(nil),       // 11: CreateMoodResponse
	(*timestamp.Timestamp)(nil),      // 12: google.protobuf.Timestamp
}
var file_mood_proto_depIdxs = []int32{
	12, // 0: EntryWithDate.record_entry:type_name -> google.protobuf.Timestamp
	0,  // 1: AddEntryRequest.entry:type_name -> Entry
	12, // 2: RecordStat.record_entry:type_name -> google.protobuf.Timestamp
	7,  // 3: MoodStat.record_stats:type_name -> RecordStat
	1,  // 4: GetMoodResponse.entries:type_name -> EntryWithDate
	8,  // 5: GetMoodResponse.stats:type_name -> MoodStat
	2,  // 6: Mood.AddEntry:input_type -> AddEntryRequest
	4,  // 7: Mood.GetMoodFromEntry:input_type -> GetMoodFromEntryRequest
	6,  // 8: Mood.GetMood:input_type -> GetMoodRequest
	10, // 9: Mood.CreateMood:input_type -> CreateMoodRequest
	3,  // 10: Mood.AddEntry:output_type -> AddEntryResponse
	5,  // 11: Mood.GetMoodFromEntry:output_type -> GetMoodFromEntryResponse
	9,  // 12: Mood.GetMood:output_type -> GetMoodResponse
	11, // 13: Mood.CreateMood:output_type -> CreateMoodResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_mood_proto_init() }
func file_mood_proto_init() {
	if File_mood_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mood_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_mood_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryWithDate); i {
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
		file_mood_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEntryRequest); i {
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
		file_mood_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEntryResponse); i {
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
		file_mood_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoodFromEntryRequest); i {
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
		file_mood_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoodFromEntryResponse); i {
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
		file_mood_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoodRequest); i {
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
		file_mood_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordStat); i {
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
		file_mood_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoodStat); i {
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
		file_mood_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoodResponse); i {
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
		file_mood_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMoodRequest); i {
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
		file_mood_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMoodResponse); i {
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
			RawDescriptor: file_mood_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mood_proto_goTypes,
		DependencyIndexes: file_mood_proto_depIdxs,
		MessageInfos:      file_mood_proto_msgTypes,
	}.Build()
	File_mood_proto = out.File
	file_mood_proto_rawDesc = nil
	file_mood_proto_goTypes = nil
	file_mood_proto_depIdxs = nil
}
