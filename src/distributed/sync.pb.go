// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: distributed/sync.proto

package distributed

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

type Flags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchSize int32 `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
}

func (x *Flags) Reset() {
	*x = Flags{}
	mi := &file_distributed_sync_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Flags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flags) ProtoMessage() {}

func (x *Flags) ProtoReflect() protoreflect.Message {
	mi := &file_distributed_sync_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flags.ProtoReflect.Descriptor instead.
func (*Flags) Descriptor() ([]byte, []int) {
	return file_distributed_sync_proto_rawDescGZIP(), []int{0}
}

func (x *Flags) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  string            `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Url     string            `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Body    string            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Headers map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Params  map[string]string `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_distributed_sync_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_distributed_sync_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_distributed_sync_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Request) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Request) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Request) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	From int32    `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	To   int32    `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
	Dict []string `protobuf:"bytes,4,rep,name=dict,proto3" json:"dict,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	mi := &file_distributed_sync_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_distributed_sync_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_distributed_sync_proto_rawDescGZIP(), []int{2}
}

func (x *Params) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Params) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *Params) GetTo() int32 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *Params) GetDict() []string {
	if x != nil {
		return x.Dict
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        int32             `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	StatusTxt     string            `protobuf:"bytes,2,opt,name=statusTxt,proto3" json:"statusTxt,omitempty"`
	Body          string            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	ProtoVer      string            `protobuf:"bytes,4,opt,name=protoVer,proto3" json:"protoVer,omitempty"`
	ContentLength int64             `protobuf:"varint,5,opt,name=contentLength,proto3" json:"contentLength,omitempty"`
	Headers       map[string]string `protobuf:"bytes,6,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_distributed_sync_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_distributed_sync_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_distributed_sync_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetStatusTxt() string {
	if x != nil {
		return x.StatusTxt
	}
	return ""
}

func (x *Response) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Response) GetProtoVer() string {
	if x != nil {
		return x.ProtoVer
	}
	return ""
}

func (x *Response) GetContentLength() int64 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

func (x *Response) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int32       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Reponses []*Response `protobuf:"bytes,2,rep,name=reponses,proto3" json:"reponses,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	mi := &file_distributed_sync_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_distributed_sync_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_distributed_sync_proto_rawDescGZIP(), []int{4}
}

func (x *Log) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Log) GetReponses() []*Response {
	if x != nil {
		return x.Reponses
	}
	return nil
}

type Criteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Response *Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Criteria) Reset() {
	*x = Criteria{}
	mi := &file_distributed_sync_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Criteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Criteria) ProtoMessage() {}

func (x *Criteria) ProtoReflect() protoreflect.Message {
	mi := &file_distributed_sync_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Criteria.ProtoReflect.Descriptor instead.
func (*Criteria) Descriptor() ([]byte, []int) {
	return file_distributed_sync_proto_rawDescGZIP(), []int{5}
}

func (x *Criteria) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Criteria) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *Request           `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Criteria *Criteria          `protobuf:"bytes,2,opt,name=criteria,proto3" json:"criteria,omitempty"`
	Params   map[string]*Params `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_distributed_sync_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_distributed_sync_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_distributed_sync_proto_rawDescGZIP(), []int{6}
}

func (x *Config) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Config) GetCriteria() *Criteria {
	if x != nil {
		return x.Criteria
	}
	return nil
}

func (x *Config) GetParams() map[string]*Params {
	if x != nil {
		return x.Params
	}
	return nil
}

type SyncMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HELLO, ACK, START, STOP, SUCCESS, LOG
	Action      string  `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Config      *Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Flags       *Flags  `protobuf:"bytes,3,opt,name=flags,proto3" json:"flags,omitempty"` // TODO: Remove, useless
	ResponseLog *Log    `protobuf:"bytes,4,opt,name=responseLog,proto3" json:"responseLog,omitempty"`
}

func (x *SyncMessage) Reset() {
	*x = SyncMessage{}
	mi := &file_distributed_sync_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SyncMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMessage) ProtoMessage() {}

func (x *SyncMessage) ProtoReflect() protoreflect.Message {
	mi := &file_distributed_sync_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMessage.ProtoReflect.Descriptor instead.
func (*SyncMessage) Descriptor() ([]byte, []int) {
	return file_distributed_sync_proto_rawDescGZIP(), []int{7}
}

func (x *SyncMessage) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SyncMessage) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *SyncMessage) GetFlags() *Flags {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *SyncMessage) GetResponseLog() *Log {
	if x != nil {
		return x.ResponseLog
	}
	return nil
}

var File_distributed_sync_proto protoreflect.FileDescriptor

var file_distributed_sync_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2f, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x26, 0x0a, 0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xa9, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x35, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x3a, 0x0a, 0x0c,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x54, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x63, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x69, 0x63, 0x74, 0x22, 0x8a, 0x02, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x36, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x48, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x22, 0x4b, 0x0a, 0x08, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xdc, 0x01,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12,
	0x31, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x48, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9e, 0x01, 0x0a,
	0x0b, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x05, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12,
	0x2c, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67,
	0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_distributed_sync_proto_rawDescOnce sync.Once
	file_distributed_sync_proto_rawDescData = file_distributed_sync_proto_rawDesc
)

func file_distributed_sync_proto_rawDescGZIP() []byte {
	file_distributed_sync_proto_rawDescOnce.Do(func() {
		file_distributed_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_distributed_sync_proto_rawDescData)
	})
	return file_distributed_sync_proto_rawDescData
}

var file_distributed_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_distributed_sync_proto_goTypes = []any{
	(*Flags)(nil),       // 0: proto.Flags
	(*Request)(nil),     // 1: proto.Request
	(*Params)(nil),      // 2: proto.Params
	(*Response)(nil),    // 3: proto.Response
	(*Log)(nil),         // 4: proto.Log
	(*Criteria)(nil),    // 5: proto.Criteria
	(*Config)(nil),      // 6: proto.Config
	(*SyncMessage)(nil), // 7: proto.SyncMessage
	nil,                 // 8: proto.Request.HeadersEntry
	nil,                 // 9: proto.Request.ParamsEntry
	nil,                 // 10: proto.Response.HeadersEntry
	nil,                 // 11: proto.Config.ParamsEntry
}
var file_distributed_sync_proto_depIdxs = []int32{
	8,  // 0: proto.Request.headers:type_name -> proto.Request.HeadersEntry
	9,  // 1: proto.Request.params:type_name -> proto.Request.ParamsEntry
	10, // 2: proto.Response.headers:type_name -> proto.Response.HeadersEntry
	3,  // 3: proto.Log.reponses:type_name -> proto.Response
	3,  // 4: proto.Criteria.response:type_name -> proto.Response
	1,  // 5: proto.Config.request:type_name -> proto.Request
	5,  // 6: proto.Config.criteria:type_name -> proto.Criteria
	11, // 7: proto.Config.params:type_name -> proto.Config.ParamsEntry
	6,  // 8: proto.SyncMessage.config:type_name -> proto.Config
	0,  // 9: proto.SyncMessage.flags:type_name -> proto.Flags
	4,  // 10: proto.SyncMessage.responseLog:type_name -> proto.Log
	2,  // 11: proto.Config.ParamsEntry.value:type_name -> proto.Params
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_distributed_sync_proto_init() }
func file_distributed_sync_proto_init() {
	if File_distributed_sync_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_distributed_sync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_distributed_sync_proto_goTypes,
		DependencyIndexes: file_distributed_sync_proto_depIdxs,
		MessageInfos:      file_distributed_sync_proto_msgTypes,
	}.Build()
	File_distributed_sync_proto = out.File
	file_distributed_sync_proto_rawDesc = nil
	file_distributed_sync_proto_goTypes = nil
	file_distributed_sync_proto_depIdxs = nil
}
