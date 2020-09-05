// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: controller/storage/session/store/v1/session.proto

package store

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
	_ "github.com/hashicorp/boundary/internal/gen/controller/protooptions"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public is used to access the session via an API
	// @inject_tag: gorm:"primary_key"
	PublicId string `protobuf:"bytes,10,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// user_id for the session
	// @inject_tag: `gorm:"default:null"`
	UserId string `protobuf:"bytes,20,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" gorm:"default:null"`
	// host_id of the session
	// @inject_tag: `gorm:"default:null"`
	HostId string `protobuf:"bytes,30,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty" gorm:"default:null"`
	// server_id that proxied the session
	// @inject_tag: `gorm:"default:null"`
	ServerId string `protobuf:"bytes,40,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" gorm:"default:null"`
	// server_type that proxied the session
	// @inject_tag: `gorm:"default:null"`
	ServerType string `protobuf:"bytes,50,opt,name=server_type,json=serverType,proto3" json:"server_type,omitempty" gorm:"default:null"`
	// @inject_tag: `gorm:"default:null"`
	TargetId string `protobuf:"bytes,60,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty" gorm:"default:null"`
	// set_id for the session
	// @inject_tag: `gorm:"default:null"`
	SetId string `protobuf:"bytes,70,opt,name=set_id,json=setId,proto3" json:"set_id,omitempty" gorm:"default:null"`
	// auth_token_id for the session
	// @inject_tag: `gorm:"default:null"`
	AuthTokenId string `protobuf:"bytes,80,opt,name=auth_token_id,json=authTokenId,proto3" json:"auth_token_id,omitempty" gorm:"default:null"`
	// scope id for the session
	// @inject_tag: `gorm:"default:null"`
	ScopeId string `protobuf:"bytes,90,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" gorm:"default:null"`
	// termination_reason for the session
	// @inject_tag: `gorm:"default:null"`
	TerminationReason string `protobuf:"bytes,100,opt,name=termination_reason,json=terminationReason,proto3" json:"termination_reason,omitempty" gorm:"default:null"`
	// address of the session host
	// @inject_tag: `gorm:"default:null"`
	Address string `protobuf:"bytes,110,opt,name=address,proto3" json:"address,omitempty" gorm:"default:null"`
	// port used by the session
	// @inject_tag: `gorm:"default:null"`
	Port string `protobuf:"bytes,120,opt,name=port,proto3" json:"port,omitempty" gorm:"default:null"`
	// bytes_up total for the session
	// @inject_tag: `gorm:"default:null"`
	BytesUp uint64 `protobuf:"varint,130,opt,name=bytes_up,json=bytesUp,proto3" json:"bytes_up,omitempty" gorm:"default:null"`
	// bytes_down total for the session
	// @inject_tag: `gorm:"default:null"`
	BytesDown uint64 `protobuf:"varint,140,opt,name=bytes_down,json=bytesDown,proto3" json:"bytes_down,omitempty" gorm:"default:null"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,150,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// update_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,160,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// version for the session
	// @inject_tag: `gorm:"default:null"`
	Version uint32 `protobuf:"varint,170,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_session_store_v1_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_session_store_v1_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_controller_storage_session_store_v1_session_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Session) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Session) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *Session) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Session) GetServerType() string {
	if x != nil {
		return x.ServerType
	}
	return ""
}

func (x *Session) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *Session) GetSetId() string {
	if x != nil {
		return x.SetId
	}
	return ""
}

func (x *Session) GetAuthTokenId() string {
	if x != nil {
		return x.AuthTokenId
	}
	return ""
}

func (x *Session) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *Session) GetTerminationReason() string {
	if x != nil {
		return x.TerminationReason
	}
	return ""
}

func (x *Session) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Session) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *Session) GetBytesUp() uint64 {
	if x != nil {
		return x.BytesUp
	}
	return 0
}

func (x *Session) GetBytesDown() uint64 {
	if x != nil {
		return x.BytesDown
	}
	return 0
}

func (x *Session) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Session) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Session) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// session_id references the session public id
	// @inject_tag: gorm:"primary_key"
	SessionId string `protobuf:"bytes,10,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty" gorm:"primary_key"`
	// status of the session
	// @inject_tag: gorm:"column:state"
	Status string `protobuf:"bytes,20,opt,name=status,proto3" json:"status,omitempty" gorm:"column:state"`
	// previous_end_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	PreviousEndTime *timestamp.Timestamp `protobuf:"bytes,30,opt,name=previous_end_time,json=previousEndTime,proto3" json:"previous_end_time,omitempty" gorm:"default:current_timestamp"`
	// start_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp;primary_key"`
	StartTime *timestamp.Timestamp `protobuf:"bytes,40,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" gorm:"default:current_timestamp;primary_key"`
	// end_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	EndTime *timestamp.Timestamp `protobuf:"bytes,50,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" gorm:"default:current_timestamp"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_session_store_v1_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_session_store_v1_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_controller_storage_session_store_v1_session_proto_rawDescGZIP(), []int{1}
}

func (x *State) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *State) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *State) GetPreviousEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.PreviousEndTime
	}
	return nil
}

func (x *State) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *State) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_controller_storage_session_store_v1_session_proto protoreflect.FileDescriptor

var file_controller_storage_session_store_v1_session_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x04, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x2d, 0x0a, 0x12, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x6e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x75, 0x70, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x62, 0x79, 0x74, 0x65, 0x73, 0x55, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x4c, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xa0, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0xaa, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0xa8, 0x02, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x56, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_controller_storage_session_store_v1_session_proto_rawDescOnce sync.Once
	file_controller_storage_session_store_v1_session_proto_rawDescData = file_controller_storage_session_store_v1_session_proto_rawDesc
)

func file_controller_storage_session_store_v1_session_proto_rawDescGZIP() []byte {
	file_controller_storage_session_store_v1_session_proto_rawDescOnce.Do(func() {
		file_controller_storage_session_store_v1_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_session_store_v1_session_proto_rawDescData)
	})
	return file_controller_storage_session_store_v1_session_proto_rawDescData
}

var file_controller_storage_session_store_v1_session_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_storage_session_store_v1_session_proto_goTypes = []interface{}{
	(*Session)(nil),             // 0: controller.storage.session.store.v1.Session
	(*State)(nil),               // 1: controller.storage.session.store.v1.State
	(*timestamp.Timestamp)(nil), // 2: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_session_store_v1_session_proto_depIdxs = []int32{
	2, // 0: controller.storage.session.store.v1.Session.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 1: controller.storage.session.store.v1.Session.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 2: controller.storage.session.store.v1.State.previous_end_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 3: controller.storage.session.store.v1.State.start_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 4: controller.storage.session.store.v1.State.end_time:type_name -> controller.storage.timestamp.v1.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_controller_storage_session_store_v1_session_proto_init() }
func file_controller_storage_session_store_v1_session_proto_init() {
	if File_controller_storage_session_store_v1_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_session_store_v1_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_controller_storage_session_store_v1_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
			RawDescriptor: file_controller_storage_session_store_v1_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_session_store_v1_session_proto_goTypes,
		DependencyIndexes: file_controller_storage_session_store_v1_session_proto_depIdxs,
		MessageInfos:      file_controller_storage_session_store_v1_session_proto_msgTypes,
	}.Build()
	File_controller_storage_session_store_v1_session_proto = out.File
	file_controller_storage_session_store_v1_session_proto_rawDesc = nil
	file_controller_storage_session_store_v1_session_proto_goTypes = nil
	file_controller_storage_session_store_v1_session_proto_depIdxs = nil
}
