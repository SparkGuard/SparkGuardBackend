// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: orchestrator.proto

package orchestrator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	EventID       uint64                 `protobuf:"varint,2,opt,name=eventID,proto3" json:"eventID,omitempty"`
	WorkID        uint64                 `protobuf:"varint,3,opt,name=workID,proto3" json:"workID,omitempty"`
	Tag           string                 `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_orchestrator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Task) GetEventID() uint64 {
	if x != nil {
		return x.EventID
	}
	return 0
}

func (x *Task) GetWorkID() uint64 {
	if x != nil {
		return x.WorkID
	}
	return 0
}

func (x *Task) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Task) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Runner struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag           string                 `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Runner) Reset() {
	*x = Runner{}
	mi := &file_orchestrator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Runner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Runner) ProtoMessage() {}

func (x *Runner) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Runner.ProtoReflect.Descriptor instead.
func (*Runner) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{1}
}

func (x *Runner) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Runner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Runner) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type GetRunnerInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Runner        *Runner                `protobuf:"bytes,1,opt,name=runner,proto3" json:"runner,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRunnerInfoResponse) Reset() {
	*x = GetRunnerInfoResponse{}
	mi := &file_orchestrator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRunnerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRunnerInfoResponse) ProtoMessage() {}

func (x *GetRunnerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRunnerInfoResponse.ProtoReflect.Descriptor instead.
func (*GetRunnerInfoResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{2}
}

func (x *GetRunnerInfoResponse) GetRunner() *Runner {
	if x != nil {
		return x.Runner
	}
	return nil
}

type GetWorksOfEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventID       uint64                 `protobuf:"varint,1,opt,name=eventID,proto3" json:"eventID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorksOfEventRequest) Reset() {
	*x = GetWorksOfEventRequest{}
	mi := &file_orchestrator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorksOfEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorksOfEventRequest) ProtoMessage() {}

func (x *GetWorksOfEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorksOfEventRequest.ProtoReflect.Descriptor instead.
func (*GetWorksOfEventRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{3}
}

func (x *GetWorksOfEventRequest) GetEventID() uint64 {
	if x != nil {
		return x.EventID
	}
	return 0
}

type GetWorksOfEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkID        []uint64               `protobuf:"varint,1,rep,packed,name=workID,proto3" json:"workID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorksOfEventResponse) Reset() {
	*x = GetWorksOfEventResponse{}
	mi := &file_orchestrator_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorksOfEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorksOfEventResponse) ProtoMessage() {}

func (x *GetWorksOfEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorksOfEventResponse.ProtoReflect.Descriptor instead.
func (*GetWorksOfEventResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{4}
}

func (x *GetWorksOfEventResponse) GetWorkID() []uint64 {
	if x != nil {
		return x.WorkID
	}
	return nil
}

type GetWorksDownloadLinksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkID        []uint64               `protobuf:"varint,1,rep,packed,name=workID,proto3" json:"workID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorksDownloadLinksRequest) Reset() {
	*x = GetWorksDownloadLinksRequest{}
	mi := &file_orchestrator_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorksDownloadLinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorksDownloadLinksRequest) ProtoMessage() {}

func (x *GetWorksDownloadLinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorksDownloadLinksRequest.ProtoReflect.Descriptor instead.
func (*GetWorksDownloadLinksRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{5}
}

func (x *GetWorksDownloadLinksRequest) GetWorkID() []uint64 {
	if x != nil {
		return x.WorkID
	}
	return nil
}

type GetWorksDownloadLinksResponseItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkID        uint64                 `protobuf:"varint,1,opt,name=workID,proto3" json:"workID,omitempty"`
	DownloadLink  string                 `protobuf:"bytes,2,opt,name=downloadLink,proto3" json:"downloadLink,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorksDownloadLinksResponseItem) Reset() {
	*x = GetWorksDownloadLinksResponseItem{}
	mi := &file_orchestrator_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorksDownloadLinksResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorksDownloadLinksResponseItem) ProtoMessage() {}

func (x *GetWorksDownloadLinksResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorksDownloadLinksResponseItem.ProtoReflect.Descriptor instead.
func (*GetWorksDownloadLinksResponseItem) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{6}
}

func (x *GetWorksDownloadLinksResponseItem) GetWorkID() uint64 {
	if x != nil {
		return x.WorkID
	}
	return 0
}

func (x *GetWorksDownloadLinksResponseItem) GetDownloadLink() string {
	if x != nil {
		return x.DownloadLink
	}
	return ""
}

type GetWorksDownloadLinksResponse struct {
	state         protoimpl.MessageState               `protogen:"open.v1"`
	Item          []*GetWorksDownloadLinksResponseItem `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorksDownloadLinksResponse) Reset() {
	*x = GetWorksDownloadLinksResponse{}
	mi := &file_orchestrator_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorksDownloadLinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorksDownloadLinksResponse) ProtoMessage() {}

func (x *GetWorksDownloadLinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorksDownloadLinksResponse.ProtoReflect.Descriptor instead.
func (*GetWorksDownloadLinksResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{7}
}

func (x *GetWorksDownloadLinksResponse) GetItem() []*GetWorksDownloadLinksResponseItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type GetNewTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNewTaskResponse) Reset() {
	*x = GetNewTaskResponse{}
	mi := &file_orchestrator_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNewTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewTaskResponse) ProtoMessage() {}

func (x *GetNewTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewTaskResponse.ProtoReflect.Descriptor instead.
func (*GetNewTaskResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{8}
}

func (x *GetNewTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type CloseTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            []uint64               `protobuf:"varint,1,rep,packed,name=ID,proto3" json:"ID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CloseTaskRequest) Reset() {
	*x = CloseTaskRequest{}
	mi := &file_orchestrator_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloseTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseTaskRequest) ProtoMessage() {}

func (x *CloseTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseTaskRequest.ProtoReflect.Descriptor instead.
func (*CloseTaskRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{9}
}

func (x *CloseTaskRequest) GetID() []uint64 {
	if x != nil {
		return x.ID
	}
	return nil
}

type SendCrossCheckReportMatches struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	FirstWorkPath   string                 `protobuf:"bytes,1,opt,name=firstWorkPath,proto3" json:"firstWorkPath,omitempty"`
	FirstWorkStart  uint64                 `protobuf:"varint,2,opt,name=firstWorkStart,proto3" json:"firstWorkStart,omitempty"`
	FirstWorkSize   uint64                 `protobuf:"varint,3,opt,name=firstWorkSize,proto3" json:"firstWorkSize,omitempty"`
	SecondWorkPath  string                 `protobuf:"bytes,4,opt,name=secondWorkPath,proto3" json:"secondWorkPath,omitempty"`
	SecondWorkStart uint64                 `protobuf:"varint,5,opt,name=secondWorkStart,proto3" json:"secondWorkStart,omitempty"`
	SecondWorkSize  uint64                 `protobuf:"varint,6,opt,name=secondWorkSize,proto3" json:"secondWorkSize,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *SendCrossCheckReportMatches) Reset() {
	*x = SendCrossCheckReportMatches{}
	mi := &file_orchestrator_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendCrossCheckReportMatches) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCrossCheckReportMatches) ProtoMessage() {}

func (x *SendCrossCheckReportMatches) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCrossCheckReportMatches.ProtoReflect.Descriptor instead.
func (*SendCrossCheckReportMatches) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{10}
}

func (x *SendCrossCheckReportMatches) GetFirstWorkPath() string {
	if x != nil {
		return x.FirstWorkPath
	}
	return ""
}

func (x *SendCrossCheckReportMatches) GetFirstWorkStart() uint64 {
	if x != nil {
		return x.FirstWorkStart
	}
	return 0
}

func (x *SendCrossCheckReportMatches) GetFirstWorkSize() uint64 {
	if x != nil {
		return x.FirstWorkSize
	}
	return 0
}

func (x *SendCrossCheckReportMatches) GetSecondWorkPath() string {
	if x != nil {
		return x.SecondWorkPath
	}
	return ""
}

func (x *SendCrossCheckReportMatches) GetSecondWorkStart() uint64 {
	if x != nil {
		return x.SecondWorkStart
	}
	return 0
}

func (x *SendCrossCheckReportMatches) GetSecondWorkSize() uint64 {
	if x != nil {
		return x.SecondWorkSize
	}
	return 0
}

type SendCrossCheckReportRequest struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	FirstWorkID   uint64                         `protobuf:"varint,1,opt,name=firstWorkID,proto3" json:"firstWorkID,omitempty"`
	SecondWorkID  uint64                         `protobuf:"varint,2,opt,name=secondWorkID,proto3" json:"secondWorkID,omitempty"`
	Match         []*SendCrossCheckReportMatches `protobuf:"bytes,3,rep,name=match,proto3" json:"match,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendCrossCheckReportRequest) Reset() {
	*x = SendCrossCheckReportRequest{}
	mi := &file_orchestrator_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendCrossCheckReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCrossCheckReportRequest) ProtoMessage() {}

func (x *SendCrossCheckReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCrossCheckReportRequest.ProtoReflect.Descriptor instead.
func (*SendCrossCheckReportRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{11}
}

func (x *SendCrossCheckReportRequest) GetFirstWorkID() uint64 {
	if x != nil {
		return x.FirstWorkID
	}
	return 0
}

func (x *SendCrossCheckReportRequest) GetSecondWorkID() uint64 {
	if x != nil {
		return x.SecondWorkID
	}
	return 0
}

func (x *SendCrossCheckReportRequest) GetMatch() []*SendCrossCheckReportMatches {
	if x != nil {
		return x.Match
	}
	return nil
}

type SendDefaultReportSegment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkPath      string                 `protobuf:"bytes,1,opt,name=workPath,proto3" json:"workPath,omitempty"`
	WorkStart     uint64                 `protobuf:"varint,2,opt,name=workStart,proto3" json:"workStart,omitempty"`
	WorkSize      uint64                 `protobuf:"varint,3,opt,name=workSize,proto3" json:"workSize,omitempty"`
	Accuracy      float32                `protobuf:"fixed32,4,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendDefaultReportSegment) Reset() {
	*x = SendDefaultReportSegment{}
	mi := &file_orchestrator_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendDefaultReportSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDefaultReportSegment) ProtoMessage() {}

func (x *SendDefaultReportSegment) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDefaultReportSegment.ProtoReflect.Descriptor instead.
func (*SendDefaultReportSegment) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{12}
}

func (x *SendDefaultReportSegment) GetWorkPath() string {
	if x != nil {
		return x.WorkPath
	}
	return ""
}

func (x *SendDefaultReportSegment) GetWorkStart() uint64 {
	if x != nil {
		return x.WorkStart
	}
	return 0
}

func (x *SendDefaultReportSegment) GetWorkSize() uint64 {
	if x != nil {
		return x.WorkSize
	}
	return 0
}

func (x *SendDefaultReportSegment) GetAccuracy() float32 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

type SendDefaultReportRequest struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	WorkID        uint64                      `protobuf:"varint,1,opt,name=workID,proto3" json:"workID,omitempty"`
	Segment       []*SendDefaultReportSegment `protobuf:"bytes,2,rep,name=segment,proto3" json:"segment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendDefaultReportRequest) Reset() {
	*x = SendDefaultReportRequest{}
	mi := &file_orchestrator_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendDefaultReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDefaultReportRequest) ProtoMessage() {}

func (x *SendDefaultReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDefaultReportRequest.ProtoReflect.Descriptor instead.
func (*SendDefaultReportRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{13}
}

func (x *SendDefaultReportRequest) GetWorkID() uint64 {
	if x != nil {
		return x.WorkID
	}
	return 0
}

func (x *SendDefaultReportRequest) GetSegment() []*SendDefaultReportSegment {
	if x != nil {
		return x.Segment
	}
	return nil
}

var File_orchestrator_proto protoreflect.FileDescriptor

var file_orchestrator_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x72, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3e, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x38, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x06, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x06, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x22,
	0x32, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x4f, 0x66, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x22, 0x31, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x4f,
	0x66, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x22, 0x36, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x22, 0x5f,
	0x0a, 0x21, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x22,
	0x57, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x2f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x8b, 0x02,
	0x0a, 0x1b, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x57, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x57, 0x6f, 0x72,
	0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x1b,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x49,
	0x44, 0x12, 0x32, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x05,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x22, 0x8c, 0x01, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c,
	0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x75,
	0x72, 0x61, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x61, 0x63, 0x63, 0x75,
	0x72, 0x61, 0x63, 0x79, 0x22, 0x67, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xb7, 0x04,
	0x0a, 0x0c, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x11, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x3f, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x57,
	0x69, 0x74, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x4f,
	0x66, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x12, 0x1d, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x72, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x46, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x29, 0x5a, 0x27, 0x53, 0x70, 0x61, 0x72, 0x6b,
	0x47, 0x75, 0x61, 0x72, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_orchestrator_proto_rawDescOnce sync.Once
	file_orchestrator_proto_rawDescData []byte
)

func file_orchestrator_proto_rawDescGZIP() []byte {
	file_orchestrator_proto_rawDescOnce.Do(func() {
		file_orchestrator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_orchestrator_proto_rawDesc), len(file_orchestrator_proto_rawDesc)))
	})
	return file_orchestrator_proto_rawDescData
}

var file_orchestrator_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_orchestrator_proto_goTypes = []any{
	(*Task)(nil),                              // 0: Task
	(*Runner)(nil),                            // 1: Runner
	(*GetRunnerInfoResponse)(nil),             // 2: GetRunnerInfoResponse
	(*GetWorksOfEventRequest)(nil),            // 3: GetWorksOfEventRequest
	(*GetWorksOfEventResponse)(nil),           // 4: GetWorksOfEventResponse
	(*GetWorksDownloadLinksRequest)(nil),      // 5: GetWorksDownloadLinksRequest
	(*GetWorksDownloadLinksResponseItem)(nil), // 6: GetWorksDownloadLinksResponseItem
	(*GetWorksDownloadLinksResponse)(nil),     // 7: GetWorksDownloadLinksResponse
	(*GetNewTaskResponse)(nil),                // 8: GetNewTaskResponse
	(*CloseTaskRequest)(nil),                  // 9: CloseTaskRequest
	(*SendCrossCheckReportMatches)(nil),       // 10: SendCrossCheckReportMatches
	(*SendCrossCheckReportRequest)(nil),       // 11: SendCrossCheckReportRequest
	(*SendDefaultReportSegment)(nil),          // 12: SendDefaultReportSegment
	(*SendDefaultReportRequest)(nil),          // 13: SendDefaultReportRequest
	(*emptypb.Empty)(nil),                     // 14: google.protobuf.Empty
}
var file_orchestrator_proto_depIdxs = []int32{
	1,  // 0: GetRunnerInfoResponse.runner:type_name -> Runner
	6,  // 1: GetWorksDownloadLinksResponse.item:type_name -> GetWorksDownloadLinksResponseItem
	0,  // 2: GetNewTaskResponse.task:type_name -> Task
	10, // 3: SendCrossCheckReportRequest.match:type_name -> SendCrossCheckReportMatches
	12, // 4: SendDefaultReportRequest.segment:type_name -> SendDefaultReportSegment
	14, // 5: Orchestrator.GetRunnerInfo:input_type -> google.protobuf.Empty
	14, // 6: Orchestrator.GetNewTask:input_type -> google.protobuf.Empty
	9,  // 7: Orchestrator.CloseTask:input_type -> CloseTaskRequest
	9,  // 8: Orchestrator.CloseTaskWithError:input_type -> CloseTaskRequest
	3,  // 9: Orchestrator.GetWorksOfEvent:input_type -> GetWorksOfEventRequest
	5,  // 10: Orchestrator.GetWorksDownloadLinks:input_type -> GetWorksDownloadLinksRequest
	11, // 11: Orchestrator.SendCrossCheckReport:input_type -> SendCrossCheckReportRequest
	13, // 12: Orchestrator.SendDefaultReport:input_type -> SendDefaultReportRequest
	2,  // 13: Orchestrator.GetRunnerInfo:output_type -> GetRunnerInfoResponse
	8,  // 14: Orchestrator.GetNewTask:output_type -> GetNewTaskResponse
	14, // 15: Orchestrator.CloseTask:output_type -> google.protobuf.Empty
	14, // 16: Orchestrator.CloseTaskWithError:output_type -> google.protobuf.Empty
	4,  // 17: Orchestrator.GetWorksOfEvent:output_type -> GetWorksOfEventResponse
	7,  // 18: Orchestrator.GetWorksDownloadLinks:output_type -> GetWorksDownloadLinksResponse
	14, // 19: Orchestrator.SendCrossCheckReport:output_type -> google.protobuf.Empty
	14, // 20: Orchestrator.SendDefaultReport:output_type -> google.protobuf.Empty
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_orchestrator_proto_init() }
func file_orchestrator_proto_init() {
	if File_orchestrator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_orchestrator_proto_rawDesc), len(file_orchestrator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orchestrator_proto_goTypes,
		DependencyIndexes: file_orchestrator_proto_depIdxs,
		MessageInfos:      file_orchestrator_proto_msgTypes,
	}.Build()
	File_orchestrator_proto = out.File
	file_orchestrator_proto_goTypes = nil
	file_orchestrator_proto_depIdxs = nil
}
