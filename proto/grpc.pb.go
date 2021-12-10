// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: grpc.proto

// 定义包名

package proto

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Keys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *Keys) Reset() {
	*x = Keys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keys) ProtoMessage() {}

func (x *Keys) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keys.ProtoReflect.Descriptor instead.
func (*Keys) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Keys) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

// worker发起主任务启动请求
type JobKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerTriggerTime     *timestamp.Timestamp `protobuf:"bytes,1,opt,name=WorkerTriggerTime,proto3" json:"WorkerTriggerTime,omitempty"`
	LeaderTriggerTime     *timestamp.Timestamp `protobuf:"bytes,2,opt,name=LeaderTriggerTime,proto3" json:"LeaderTriggerTime,omitempty"`
	UpdateTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	Name                  string               `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"` // 主任务名
	Mode                  string               `protobuf:"bytes,5,opt,name=mode,proto3" json:"mode,omitempty"` // 指定广播模式或者hash模式
	Keys                  []string             `protobuf:"bytes,6,rep,name=Keys,proto3" json:"Keys,omitempty"` // 待分发的keys
	JobNode               string               `protobuf:"bytes,7,opt,name=JobNode,proto3" json:"JobNode,omitempty"`
	PreDistributeJobs     map[string]*Keys     `protobuf:"bytes,8,rep,name=PreDistributeJobs,proto3" json:"PreDistributeJobs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`          // 预分配任务
	DistributedKeys       map[string]*Keys     `protobuf:"bytes,9,rep,name=DistributedKeys,proto3" json:"DistributedKeys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`              // 已经分配的keys
	DistributedFailedKeys map[string]*Keys     `protobuf:"bytes,10,rep,name=DistributedFailedKeys,proto3" json:"DistributedFailedKeys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 分配失败的节点和keys
	BroadcastIPs          map[string]bool      `protobuf:"bytes,11,rep,name=BroadcastIPs,proto3" json:"BroadcastIPs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`                  // 已经分配的节点
	AllocationResult      bool                 `protobuf:"varint,12,opt,name=AllocationResult,proto3" json:"AllocationResult,omitempty"`                                                                                                  // 任务分配结果
}

func (x *JobKeyReq) Reset() {
	*x = JobKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobKeyReq) ProtoMessage() {}

func (x *JobKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobKeyReq.ProtoReflect.Descriptor instead.
func (*JobKeyReq) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *JobKeyReq) GetWorkerTriggerTime() *timestamp.Timestamp {
	if x != nil {
		return x.WorkerTriggerTime
	}
	return nil
}

func (x *JobKeyReq) GetLeaderTriggerTime() *timestamp.Timestamp {
	if x != nil {
		return x.LeaderTriggerTime
	}
	return nil
}

func (x *JobKeyReq) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *JobKeyReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobKeyReq) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *JobKeyReq) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *JobKeyReq) GetJobNode() string {
	if x != nil {
		return x.JobNode
	}
	return ""
}

func (x *JobKeyReq) GetPreDistributeJobs() map[string]*Keys {
	if x != nil {
		return x.PreDistributeJobs
	}
	return nil
}

func (x *JobKeyReq) GetDistributedKeys() map[string]*Keys {
	if x != nil {
		return x.DistributedKeys
	}
	return nil
}

func (x *JobKeyReq) GetDistributedFailedKeys() map[string]*Keys {
	if x != nil {
		return x.DistributedFailedKeys
	}
	return nil
}

func (x *JobKeyReq) GetBroadcastIPs() map[string]bool {
	if x != nil {
		return x.BroadcastIPs
	}
	return nil
}

func (x *JobKeyReq) GetAllocationResult() bool {
	if x != nil {
		return x.AllocationResult
	}
	return false
}

// leader响应主任务启动请求
type JobKeyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Errno      string `protobuf:"bytes,2,opt,name=errno,proto3" json:"errno,omitempty"`
	ErrMessage string `protobuf:"bytes,3,opt,name=errMessage,proto3" json:"errMessage,omitempty"`
}

func (x *JobKeyResp) Reset() {
	*x = JobKeyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobKeyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobKeyResp) ProtoMessage() {}

func (x *JobKeyResp) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobKeyResp.ProtoReflect.Descriptor instead.
func (*JobKeyResp) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *JobKeyResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *JobKeyResp) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *JobKeyResp) GetErrMessage() string {
	if x != nil {
		return x.ErrMessage
	}
	return ""
}

// leader分配子任务
type ChildJobReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode    string   `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`       // 任务模式
	JobName string   `protobuf:"bytes,2,opt,name=JobName,proto3" json:"JobName,omitempty"` // 主任务名
	JobKeys []string `protobuf:"bytes,3,rep,name=JobKeys,proto3" json:"JobKeys,omitempty"` // 待分发的keys
}

func (x *ChildJobReq) Reset() {
	*x = ChildJobReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChildJobReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChildJobReq) ProtoMessage() {}

func (x *ChildJobReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChildJobReq.ProtoReflect.Descriptor instead.
func (*ChildJobReq) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *ChildJobReq) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *ChildJobReq) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *ChildJobReq) GetJobKeys() []string {
	if x != nil {
		return x.JobKeys
	}
	return nil
}

// worker响应子任务请求
type ChildJobResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Errno      string `protobuf:"bytes,2,opt,name=errno,proto3" json:"errno,omitempty"`
	ErrMessage string `protobuf:"bytes,3,opt,name=errMessage,proto3" json:"errMessage,omitempty"`
}

func (x *ChildJobResp) Reset() {
	*x = ChildJobResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChildJobResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChildJobResp) ProtoMessage() {}

func (x *ChildJobResp) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChildJobResp.ProtoReflect.Descriptor instead.
func (*ChildJobResp) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *ChildJobResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *ChildJobResp) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *ChildJobResp) GetErrMessage() string {
	if x != nil {
		return x.ErrMessage
	}
	return ""
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x22, 0xec, 0x07, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x12, 0x48,
	0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x11, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x11, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4a, 0x6f,
	0x62, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4a, 0x6f, 0x62,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x55, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x2e, 0x50, 0x72, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x50, 0x72, 0x65, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x4f, 0x0a, 0x0f, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x61, 0x0a, 0x15,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x2e, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4b,
	0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x12,
	0x46, 0x0a, 0x0c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x50, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f,
	0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x49, 0x50, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x49, 0x50, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x1a, 0x51, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4f, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x55, 0x0a, 0x1a, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b,
	0x65, 0x79, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f,
	0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x50, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x56, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x72, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x55, 0x0a, 0x0b, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4a, 0x6f,
	0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x4b, 0x65, 0x79, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x4a, 0x6f, 0x62, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x58,
	0x0a, 0x0c, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x72,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x7b, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f,
	0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_grpc_proto_goTypes = []interface{}{
	(*Keys)(nil),                // 0: proto.Keys
	(*JobKeyReq)(nil),           // 1: proto.JobKeyReq
	(*JobKeyResp)(nil),          // 2: proto.JobKeyResp
	(*ChildJobReq)(nil),         // 3: proto.ChildJobReq
	(*ChildJobResp)(nil),        // 4: proto.ChildJobResp
	nil,                         // 5: proto.JobKeyReq.PreDistributeJobsEntry
	nil,                         // 6: proto.JobKeyReq.DistributedKeysEntry
	nil,                         // 7: proto.JobKeyReq.DistributedFailedKeysEntry
	nil,                         // 8: proto.JobKeyReq.BroadcastIPsEntry
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_grpc_proto_depIdxs = []int32{
	9,  // 0: proto.JobKeyReq.WorkerTriggerTime:type_name -> google.protobuf.Timestamp
	9,  // 1: proto.JobKeyReq.LeaderTriggerTime:type_name -> google.protobuf.Timestamp
	9,  // 2: proto.JobKeyReq.UpdateTime:type_name -> google.protobuf.Timestamp
	5,  // 3: proto.JobKeyReq.PreDistributeJobs:type_name -> proto.JobKeyReq.PreDistributeJobsEntry
	6,  // 4: proto.JobKeyReq.DistributedKeys:type_name -> proto.JobKeyReq.DistributedKeysEntry
	7,  // 5: proto.JobKeyReq.DistributedFailedKeys:type_name -> proto.JobKeyReq.DistributedFailedKeysEntry
	8,  // 6: proto.JobKeyReq.BroadcastIPs:type_name -> proto.JobKeyReq.BroadcastIPsEntry
	0,  // 7: proto.JobKeyReq.PreDistributeJobsEntry.value:type_name -> proto.Keys
	0,  // 8: proto.JobKeyReq.DistributedKeysEntry.value:type_name -> proto.Keys
	0,  // 9: proto.JobKeyReq.DistributedFailedKeysEntry.value:type_name -> proto.Keys
	1,  // 10: proto.JobService.TaskStart:input_type -> proto.JobKeyReq
	3,  // 11: proto.JobService.SubTaskStart:input_type -> proto.ChildJobReq
	2,  // 12: proto.JobService.TaskStart:output_type -> proto.JobKeyResp
	4,  // 13: proto.JobService.SubTaskStart:output_type -> proto.ChildJobResp
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keys); i {
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
		file_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobKeyReq); i {
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
		file_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobKeyResp); i {
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
		file_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChildJobReq); i {
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
		file_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChildJobResp); i {
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
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	// TaskStart方法,leader接受worker消息, 启动主任务
	TaskStart(ctx context.Context, in *JobKeyReq, opts ...grpc.CallOption) (*JobKeyResp, error)
	// SubTaskStart,leader分发子任务到worker, 启动子任务
	SubTaskStart(ctx context.Context, in *ChildJobReq, opts ...grpc.CallOption) (*ChildJobResp, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) TaskStart(ctx context.Context, in *JobKeyReq, opts ...grpc.CallOption) (*JobKeyResp, error) {
	out := new(JobKeyResp)
	err := c.cc.Invoke(ctx, "/proto.JobService/TaskStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) SubTaskStart(ctx context.Context, in *ChildJobReq, opts ...grpc.CallOption) (*ChildJobResp, error) {
	out := new(ChildJobResp)
	err := c.cc.Invoke(ctx, "/proto.JobService/SubTaskStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	// TaskStart方法,leader接受worker消息, 启动主任务
	TaskStart(context.Context, *JobKeyReq) (*JobKeyResp, error)
	// SubTaskStart,leader分发子任务到worker, 启动子任务
	SubTaskStart(context.Context, *ChildJobReq) (*ChildJobResp, error)
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) TaskStart(context.Context, *JobKeyReq) (*JobKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskStart not implemented")
}
func (*UnimplementedJobServiceServer) SubTaskStart(context.Context, *ChildJobReq) (*ChildJobResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubTaskStart not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_TaskStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).TaskStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.JobService/TaskStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).TaskStart(ctx, req.(*JobKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_SubTaskStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChildJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SubTaskStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.JobService/SubTaskStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SubTaskStart(ctx, req.(*ChildJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskStart",
			Handler:    _JobService_TaskStart_Handler,
		},
		{
			MethodName: "SubTaskStart",
			Handler:    _JobService_SubTaskStart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}