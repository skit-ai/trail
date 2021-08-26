// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: labels.proto

package dataframes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A simple intent annotation. Intents here are from a fixed set. For marking
//absence of an intent, we use null. We skip label if item is untagged.
type IntentLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Intent string `protobuf:"bytes,2,opt,name=intent,proto3" json:"intent,omitempty"`
}

func (x *IntentLabel) Reset() {
	*x = IntentLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntentLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntentLabel) ProtoMessage() {}

func (x *IntentLabel) ProtoReflect() protoreflect.Message {
	mi := &file_labels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntentLabel.ProtoReflect.Descriptor instead.
func (*IntentLabel) Descriptor() ([]byte, []int) {
	return file_labels_proto_rawDescGZIP(), []int{0}
}

func (x *IntentLabel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IntentLabel) GetIntent() string {
	if x != nil {
		return x.Intent
	}
	return ""
}

// Similar to intent but allows richer labels. Two common patterns are described
//below:
//
//1. Multiple intents: {"op": "AND", "items": [{"intent": "<>": "score": "<>"},
//...]}
//
//2. Ranked intents: [{"intent": "<>": "score": "<>"}, {"intent": "<>":
//"score": "<>"}, ...]
type RichIntentLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RichIntents *anypb.Any `protobuf:"bytes,2,opt,name=rich_intents,json=richIntents,proto3" json:"rich_intents,omitempty"`
}

func (x *RichIntentLabel) Reset() {
	*x = RichIntentLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RichIntentLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RichIntentLabel) ProtoMessage() {}

func (x *RichIntentLabel) ProtoReflect() protoreflect.Message {
	mi := &file_labels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RichIntentLabel.ProtoReflect.Descriptor instead.
func (*RichIntentLabel) Descriptor() ([]byte, []int) {
	return file_labels_proto_rawDescGZIP(), []int{1}
}

func (x *RichIntentLabel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RichIntentLabel) GetRichIntents() *anypb.Any {
	if x != nil {
		return x.RichIntents
	}
	return nil
}

// Entities or slots. Can merge multiple entities in one. If they are
//independent, it's recommended to make new layer.
type EntityLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Entities []*Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *EntityLabel) Reset() {
	*x = EntityLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityLabel) ProtoMessage() {}

func (x *EntityLabel) ProtoReflect() protoreflect.Message {
	mi := &file_labels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityLabel.ProtoReflect.Descriptor instead.
func (*EntityLabel) Descriptor() ([]byte, []int) {
	return file_labels_proto_rawDescGZIP(), []int{2}
}

func (x *EntityLabel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EntityLabel) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

// An entity definition.
type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string     `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Score float32    `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	Value *anypb.Any `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"` // Null here means that this entity is not present
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_labels_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_labels_proto_rawDescGZIP(), []int{3}
}

func (x *Entity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Entity) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Entity) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

// Similar to rich intents. Common patterns are described below:
//
//1. Multiple entities: {"op": "AND", "items": [<Entity>, ...]}
//
//2. Ranked entities: [<Entity>, <Entity> ...]
type RichEntityLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RichEntities *anypb.Any `protobuf:"bytes,3,opt,name=rich_entities,json=richEntities,proto3" json:"rich_entities,omitempty"`
}

func (x *RichEntityLabel) Reset() {
	*x = RichEntityLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labels_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RichEntityLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RichEntityLabel) ProtoMessage() {}

func (x *RichEntityLabel) ProtoReflect() protoreflect.Message {
	mi := &file_labels_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RichEntityLabel.ProtoReflect.Descriptor instead.
func (*RichEntityLabel) Descriptor() ([]byte, []int) {
	return file_labels_proto_rawDescGZIP(), []int{4}
}

func (x *RichEntityLabel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RichEntityLabel) GetRichEntities() *anypb.Any {
	if x != nil {
		return x.RichEntities
	}
	return nil
}

// Plain transcription label
type TranscriptionLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Transcription string `protobuf:"bytes,2,opt,name=transcription,proto3" json:"transcription,omitempty"`
}

func (x *TranscriptionLabel) Reset() {
	*x = TranscriptionLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labels_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscriptionLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscriptionLabel) ProtoMessage() {}

func (x *TranscriptionLabel) ProtoReflect() protoreflect.Message {
	mi := &file_labels_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscriptionLabel.ProtoReflect.Descriptor instead.
func (*TranscriptionLabel) Descriptor() ([]byte, []int) {
	return file_labels_proto_rawDescGZIP(), []int{5}
}

func (x *TranscriptionLabel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TranscriptionLabel) GetTranscription() string {
	if x != nil {
		return x.Transcription
	}
	return ""
}

// Ranked transcription label. This is similar to the ones coming from model
//predictions.
type RichTranscriptionLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Utterances []*Utterance `protobuf:"bytes,2,rep,name=utterances,proto3" json:"utterances,omitempty"`
}

func (x *RichTranscriptionLabel) Reset() {
	*x = RichTranscriptionLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labels_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RichTranscriptionLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RichTranscriptionLabel) ProtoMessage() {}

func (x *RichTranscriptionLabel) ProtoReflect() protoreflect.Message {
	mi := &file_labels_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RichTranscriptionLabel.ProtoReflect.Descriptor instead.
func (*RichTranscriptionLabel) Descriptor() ([]byte, []int) {
	return file_labels_proto_rawDescGZIP(), []int{6}
}

func (x *RichTranscriptionLabel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RichTranscriptionLabel) GetUtterances() []*Utterance {
	if x != nil {
		return x.Utterances
	}
	return nil
}

type LabeledDataFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intents        []*IntentLabel        `protobuf:"bytes,1,rep,name=intents,proto3" json:"intents,omitempty"`
	Entities       []*EntityLabel        `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
	Transcriptions []*TranscriptionLabel `protobuf:"bytes,3,rep,name=transcriptions,proto3" json:"transcriptions,omitempty"`
}

func (x *LabeledDataFrame) Reset() {
	*x = LabeledDataFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labels_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabeledDataFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabeledDataFrame) ProtoMessage() {}

func (x *LabeledDataFrame) ProtoReflect() protoreflect.Message {
	mi := &file_labels_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabeledDataFrame.ProtoReflect.Descriptor instead.
func (*LabeledDataFrame) Descriptor() ([]byte, []int) {
	return file_labels_proto_rawDescGZIP(), []int{7}
}

func (x *LabeledDataFrame) GetIntents() []*IntentLabel {
	if x != nil {
		return x.Intents
	}
	return nil
}

func (x *LabeledDataFrame) GetEntities() []*EntityLabel {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *LabeledDataFrame) GetTranscriptions() []*TranscriptionLabel {
	if x != nil {
		return x.Transcriptions
	}
	return nil
}

var File_labels_proto protoreflect.FileDescriptor

var file_labels_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x5a, 0x0a, 0x0f, 0x52, 0x69, 0x63, 0x68, 0x49,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x72, 0x69,
	0x63, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x72, 0x69, 0x63, 0x68, 0x49, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x49, 0x0a, 0x0b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x5e,
	0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5c,
	0x0a, 0x0f, 0x52, 0x69, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x39, 0x0a, 0x0d, 0x72, 0x69, 0x63, 0x68, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0c,
	0x72, 0x69, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x12,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x16, 0x52, 0x69, 0x63, 0x68,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x32, 0x0a, 0x0a, 0x75, 0x74, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x2e, 0x55, 0x74, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x10, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52,
	0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x0e, 0x5a, 0x0c, 0x2f, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_labels_proto_rawDescOnce sync.Once
	file_labels_proto_rawDescData = file_labels_proto_rawDesc
)

func file_labels_proto_rawDescGZIP() []byte {
	file_labels_proto_rawDescOnce.Do(func() {
		file_labels_proto_rawDescData = protoimpl.X.CompressGZIP(file_labels_proto_rawDescData)
	})
	return file_labels_proto_rawDescData
}

var file_labels_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_labels_proto_goTypes = []interface{}{
	(*IntentLabel)(nil),            // 0: labels.IntentLabel
	(*RichIntentLabel)(nil),        // 1: labels.RichIntentLabel
	(*EntityLabel)(nil),            // 2: labels.EntityLabel
	(*Entity)(nil),                 // 3: labels.Entity
	(*RichEntityLabel)(nil),        // 4: labels.RichEntityLabel
	(*TranscriptionLabel)(nil),     // 5: labels.TranscriptionLabel
	(*RichTranscriptionLabel)(nil), // 6: labels.RichTranscriptionLabel
	(*LabeledDataFrame)(nil),       // 7: labels.LabeledDataFrame
	(*anypb.Any)(nil),              // 8: google.protobuf.Any
	(*Utterance)(nil),              // 9: records.Utterance
}
var file_labels_proto_depIdxs = []int32{
	8, // 0: labels.RichIntentLabel.rich_intents:type_name -> google.protobuf.Any
	3, // 1: labels.EntityLabel.entities:type_name -> labels.Entity
	8, // 2: labels.Entity.value:type_name -> google.protobuf.Any
	8, // 3: labels.RichEntityLabel.rich_entities:type_name -> google.protobuf.Any
	9, // 4: labels.RichTranscriptionLabel.utterances:type_name -> records.Utterance
	0, // 5: labels.LabeledDataFrame.intents:type_name -> labels.IntentLabel
	2, // 6: labels.LabeledDataFrame.entities:type_name -> labels.EntityLabel
	5, // 7: labels.LabeledDataFrame.transcriptions:type_name -> labels.TranscriptionLabel
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_labels_proto_init() }
func file_labels_proto_init() {
	if File_labels_proto != nil {
		return
	}
	file_records_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_labels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntentLabel); i {
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
		file_labels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RichIntentLabel); i {
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
		file_labels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityLabel); i {
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
		file_labels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_labels_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RichEntityLabel); i {
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
		file_labels_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscriptionLabel); i {
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
		file_labels_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RichTranscriptionLabel); i {
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
		file_labels_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabeledDataFrame); i {
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
			RawDescriptor: file_labels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_labels_proto_goTypes,
		DependencyIndexes: file_labels_proto_depIdxs,
		MessageInfos:      file_labels_proto_msgTypes,
	}.Build()
	File_labels_proto = out.File
	file_labels_proto_rawDesc = nil
	file_labels_proto_goTypes = nil
	file_labels_proto_depIdxs = nil
}
