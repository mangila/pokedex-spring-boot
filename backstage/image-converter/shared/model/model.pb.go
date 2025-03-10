// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: model.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type StreamRecord struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StreamKey     string                 `protobuf:"bytes,1,opt,name=stream_key,json=streamKey,proto3" json:"stream_key,omitempty"`
	RecordId      string                 `protobuf:"bytes,2,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	Data          map[string]string      `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamRecord) Reset() {
	*x = StreamRecord{}
	mi := &file_model_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRecord) ProtoMessage() {}

func (x *StreamRecord) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRecord.ProtoReflect.Descriptor instead.
func (*StreamRecord) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *StreamRecord) GetStreamKey() string {
	if x != nil {
		return x.StreamKey
	}
	return ""
}

func (x *StreamRecord) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *StreamRecord) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type EntryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         *anypb.Any             `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntryRequest) Reset() {
	*x = EntryRequest{}
	mi := &file_model_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryRequest) ProtoMessage() {}

func (x *EntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryRequest.ProtoReflect.Descriptor instead.
func (*EntryRequest) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *EntryRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *EntryRequest) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type GenerationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Generation    string                 `protobuf:"bytes,1,opt,name=generation,proto3" json:"generation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerationRequest) Reset() {
	*x = GenerationRequest{}
	mi := &file_model_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerationRequest) ProtoMessage() {}

func (x *GenerationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerationRequest.ProtoReflect.Descriptor instead.
func (*GenerationRequest) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{2}
}

func (x *GenerationRequest) GetGeneration() string {
	if x != nil {
		return x.Generation
	}
	return ""
}

type GenerationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PokemonName   []string               `protobuf:"bytes,1,rep,name=pokemon_name,json=pokemonName,proto3" json:"pokemon_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerationResponse) Reset() {
	*x = GenerationResponse{}
	mi := &file_model_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerationResponse) ProtoMessage() {}

func (x *GenerationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerationResponse.ProtoReflect.Descriptor instead.
func (*GenerationResponse) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3}
}

func (x *GenerationResponse) GetPokemonName() []string {
	if x != nil {
		return x.PokemonName
	}
	return nil
}

type PokemonSpeciesRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	PokemonSpeciesName string                 `protobuf:"bytes,1,opt,name=pokemon_species_name,json=pokemonSpeciesName,proto3" json:"pokemon_species_name,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *PokemonSpeciesRequest) Reset() {
	*x = PokemonSpeciesRequest{}
	mi := &file_model_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PokemonSpeciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonSpeciesRequest) ProtoMessage() {}

func (x *PokemonSpeciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonSpeciesRequest.ProtoReflect.Descriptor instead.
func (*PokemonSpeciesRequest) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{4}
}

func (x *PokemonSpeciesRequest) GetPokemonSpeciesName() string {
	if x != nil {
		return x.PokemonSpeciesName
	}
	return ""
}

type PokemonSpecies struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SpeciesId     int32                  `protobuf:"varint,1,opt,name=species_id,json=speciesId,proto3" json:"species_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Generation    string                 `protobuf:"bytes,3,opt,name=generation,proto3" json:"generation,omitempty"`
	Names         []*PokemonName         `protobuf:"bytes,4,rep,name=names,proto3" json:"names,omitempty"`
	Descriptions  []*PokemonDescription  `protobuf:"bytes,5,rep,name=descriptions,proto3" json:"descriptions,omitempty"`
	Genera        []*PokemonGenera       `protobuf:"bytes,6,rep,name=genera,proto3" json:"genera,omitempty"`
	Evolutions    []*PokemonEvolution    `protobuf:"bytes,7,rep,name=evolutions,proto3" json:"evolutions,omitempty"`
	Varieties     []*Pokemon             `protobuf:"bytes,8,rep,name=varieties,proto3" json:"varieties,omitempty"`
	Special       *PokemonSpecial        `protobuf:"bytes,9,opt,name=special,proto3" json:"special,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PokemonSpecies) Reset() {
	*x = PokemonSpecies{}
	mi := &file_model_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PokemonSpecies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonSpecies) ProtoMessage() {}

func (x *PokemonSpecies) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonSpecies.ProtoReflect.Descriptor instead.
func (*PokemonSpecies) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{5}
}

func (x *PokemonSpecies) GetSpeciesId() int32 {
	if x != nil {
		return x.SpeciesId
	}
	return 0
}

func (x *PokemonSpecies) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PokemonSpecies) GetGeneration() string {
	if x != nil {
		return x.Generation
	}
	return ""
}

func (x *PokemonSpecies) GetNames() []*PokemonName {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *PokemonSpecies) GetDescriptions() []*PokemonDescription {
	if x != nil {
		return x.Descriptions
	}
	return nil
}

func (x *PokemonSpecies) GetGenera() []*PokemonGenera {
	if x != nil {
		return x.Genera
	}
	return nil
}

func (x *PokemonSpecies) GetEvolutions() []*PokemonEvolution {
	if x != nil {
		return x.Evolutions
	}
	return nil
}

func (x *PokemonSpecies) GetVarieties() []*Pokemon {
	if x != nil {
		return x.Varieties
	}
	return nil
}

func (x *PokemonSpecies) GetSpecial() *PokemonSpecial {
	if x != nil {
		return x.Special
	}
	return nil
}

type PokemonDescription struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Description   string                 `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Language      string                 `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PokemonDescription) Reset() {
	*x = PokemonDescription{}
	mi := &file_model_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PokemonDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonDescription) ProtoMessage() {}

func (x *PokemonDescription) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonDescription.ProtoReflect.Descriptor instead.
func (*PokemonDescription) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{6}
}

func (x *PokemonDescription) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PokemonDescription) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type Pokemon struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PokemonId     int32                  `protobuf:"varint,1,opt,name=pokemon_id,json=pokemonId,proto3" json:"pokemon_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsDefault     bool                   `protobuf:"varint,3,opt,name=isDefault,proto3" json:"isDefault,omitempty"`
	Height        string                 `protobuf:"bytes,4,opt,name=height,proto3" json:"height,omitempty"`
	Weight        string                 `protobuf:"bytes,5,opt,name=weight,proto3" json:"weight,omitempty"`
	Types         []string               `protobuf:"bytes,6,rep,name=types,proto3" json:"types,omitempty"`
	Stats         []*PokemonStat         `protobuf:"bytes,7,rep,name=stats,proto3" json:"stats,omitempty"`
	Media         []*PokemonMedia        `protobuf:"bytes,8,rep,name=media,proto3" json:"media,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pokemon) Reset() {
	*x = Pokemon{}
	mi := &file_model_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pokemon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pokemon) ProtoMessage() {}

func (x *Pokemon) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pokemon.ProtoReflect.Descriptor instead.
func (*Pokemon) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{7}
}

func (x *Pokemon) GetPokemonId() int32 {
	if x != nil {
		return x.PokemonId
	}
	return 0
}

func (x *Pokemon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pokemon) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

func (x *Pokemon) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *Pokemon) GetWeight() string {
	if x != nil {
		return x.Weight
	}
	return ""
}

func (x *Pokemon) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *Pokemon) GetStats() []*PokemonStat {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Pokemon) GetMedia() []*PokemonMedia {
	if x != nil {
		return x.Media
	}
	return nil
}

type PokemonEvolution struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         int32                  `protobuf:"varint,1,opt,name=order,proto3" json:"order,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PokemonEvolution) Reset() {
	*x = PokemonEvolution{}
	mi := &file_model_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PokemonEvolution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonEvolution) ProtoMessage() {}

func (x *PokemonEvolution) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonEvolution.ProtoReflect.Descriptor instead.
func (*PokemonEvolution) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{8}
}

func (x *PokemonEvolution) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *PokemonEvolution) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PokemonGenera struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Genera        string                 `protobuf:"bytes,1,opt,name=genera,proto3" json:"genera,omitempty"`
	Language      string                 `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PokemonGenera) Reset() {
	*x = PokemonGenera{}
	mi := &file_model_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PokemonGenera) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonGenera) ProtoMessage() {}

func (x *PokemonGenera) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonGenera.ProtoReflect.Descriptor instead.
func (*PokemonGenera) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{9}
}

func (x *PokemonGenera) GetGenera() string {
	if x != nil {
		return x.Genera
	}
	return ""
}

func (x *PokemonGenera) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type PokemonName struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Language      string                 `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PokemonName) Reset() {
	*x = PokemonName{}
	mi := &file_model_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PokemonName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonName) ProtoMessage() {}

func (x *PokemonName) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonName.ProtoReflect.Descriptor instead.
func (*PokemonName) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{10}
}

func (x *PokemonName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PokemonName) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type PokemonSpecial struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSpecial     bool                   `protobuf:"varint,1,opt,name=isSpecial,proto3" json:"isSpecial,omitempty"`
	Legendary     bool                   `protobuf:"varint,2,opt,name=legendary,proto3" json:"legendary,omitempty"`
	Mythical      bool                   `protobuf:"varint,3,opt,name=mythical,proto3" json:"mythical,omitempty"`
	Baby          bool                   `protobuf:"varint,4,opt,name=baby,proto3" json:"baby,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PokemonSpecial) Reset() {
	*x = PokemonSpecial{}
	mi := &file_model_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PokemonSpecial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonSpecial) ProtoMessage() {}

func (x *PokemonSpecial) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonSpecial.ProtoReflect.Descriptor instead.
func (*PokemonSpecial) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{11}
}

func (x *PokemonSpecial) GetIsSpecial() bool {
	if x != nil {
		return x.IsSpecial
	}
	return false
}

func (x *PokemonSpecial) GetLegendary() bool {
	if x != nil {
		return x.Legendary
	}
	return false
}

func (x *PokemonSpecial) GetMythical() bool {
	if x != nil {
		return x.Mythical
	}
	return false
}

func (x *PokemonSpecial) GetBaby() bool {
	if x != nil {
		return x.Baby
	}
	return false
}

type PokemonStat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value         int32                  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PokemonStat) Reset() {
	*x = PokemonStat{}
	mi := &file_model_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PokemonStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonStat) ProtoMessage() {}

func (x *PokemonStat) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonStat.ProtoReflect.Descriptor instead.
func (*PokemonStat) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{12}
}

func (x *PokemonStat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PokemonStat) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type PokemonMedia struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SpeciesId     int32                  `protobuf:"varint,3,opt,name=species_id,json=speciesId,proto3" json:"species_id,omitempty"`
	PokemonId     int32                  `protobuf:"varint,4,opt,name=pokemon_id,json=pokemonId,proto3" json:"pokemon_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PokemonMedia) Reset() {
	*x = PokemonMedia{}
	mi := &file_model_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PokemonMedia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonMedia) ProtoMessage() {}

func (x *PokemonMedia) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonMedia.ProtoReflect.Descriptor instead.
func (*PokemonMedia) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{13}
}

func (x *PokemonMedia) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PokemonMedia) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PokemonMedia) GetSpeciesId() int32 {
	if x != nil {
		return x.SpeciesId
	}
	return 0
}

func (x *PokemonMedia) GetPokemonId() int32 {
	if x != nil {
		return x.PokemonId
	}
	return 0
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4c, 0x0a, 0x0c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x33, 0x0a, 0x11, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37,
	0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x6b, 0x65,
	0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x15, 0x50, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x14, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x65, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0xee, 0x02, 0x0a, 0x0e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x65, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0c,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x12, 0x31, 0x0a,
	0x0a, 0x65, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x45, 0x76, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x26, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x09, 0x76,
	0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x6f, 0x6b, 0x65,
	0x6d, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x22, 0x52, 0x0a, 0x12, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0xe9, 0x01, 0x0a, 0x07, 0x50, 0x6f, 0x6b, 0x65,
	0x6d, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x6f, 0x6b, 0x65,
	0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x23,
	0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x05, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x22, 0x3c, 0x0a, 0x10, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x45, 0x76,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x43, 0x0a, 0x0d, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x3d, 0x0a, 0x0b, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x7c, 0x0a, 0x0e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x79, 0x74, 0x68, 0x69, 0x63, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6d, 0x79, 0x74, 0x68, 0x69, 0x63, 0x61, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x61, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62,
	0x61, 0x62, 0x79, 0x22, 0x37, 0x0a, 0x0b, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x80, 0x01, 0x0a,
	0x0c, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x49, 0x64, 0x42,
	0x47, 0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x6d, 0x61,
	0x6e, 0x67, 0x69, 0x6c, 0x61, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x64, 0x65, 0x78, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x01, 0x5a, 0x0c, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData []byte
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_proto_rawDesc), len(file_model_proto_rawDesc)))
	})
	return file_model_proto_rawDescData
}

var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_model_proto_goTypes = []any{
	(*StreamRecord)(nil),          // 0: StreamRecord
	(*EntryRequest)(nil),          // 1: EntryRequest
	(*GenerationRequest)(nil),     // 2: GenerationRequest
	(*GenerationResponse)(nil),    // 3: GenerationResponse
	(*PokemonSpeciesRequest)(nil), // 4: PokemonSpeciesRequest
	(*PokemonSpecies)(nil),        // 5: PokemonSpecies
	(*PokemonDescription)(nil),    // 6: PokemonDescription
	(*Pokemon)(nil),               // 7: Pokemon
	(*PokemonEvolution)(nil),      // 8: PokemonEvolution
	(*PokemonGenera)(nil),         // 9: PokemonGenera
	(*PokemonName)(nil),           // 10: PokemonName
	(*PokemonSpecial)(nil),        // 11: PokemonSpecial
	(*PokemonStat)(nil),           // 12: PokemonStat
	(*PokemonMedia)(nil),          // 13: PokemonMedia
	nil,                           // 14: StreamRecord.DataEntry
	(*anypb.Any)(nil),             // 15: google.protobuf.Any
}
var file_model_proto_depIdxs = []int32{
	14, // 0: StreamRecord.data:type_name -> StreamRecord.DataEntry
	15, // 1: EntryRequest.value:type_name -> google.protobuf.Any
	10, // 2: PokemonSpecies.names:type_name -> PokemonName
	6,  // 3: PokemonSpecies.descriptions:type_name -> PokemonDescription
	9,  // 4: PokemonSpecies.genera:type_name -> PokemonGenera
	8,  // 5: PokemonSpecies.evolutions:type_name -> PokemonEvolution
	7,  // 6: PokemonSpecies.varieties:type_name -> Pokemon
	11, // 7: PokemonSpecies.special:type_name -> PokemonSpecial
	12, // 8: Pokemon.stats:type_name -> PokemonStat
	13, // 9: Pokemon.media:type_name -> PokemonMedia
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_proto_rawDesc), len(file_model_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}
