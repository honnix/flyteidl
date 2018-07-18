// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/tasks.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskCategory int32

const (
	// Task category that identifies if the system can use default structures in UI, etc to drive the task
	// TODO should we add Container type of tasks as a special Class?
	TaskCategory_SingleStepTask TaskCategory = 0
	TaskCategory_MultiStepTask  TaskCategory = 1
)

var TaskCategory_name = map[int32]string{
	0: "SingleStepTask",
	1: "MultiStepTask",
}
var TaskCategory_value = map[string]int32{
	"SingleStepTask": 0,
	"MultiStepTask":  1,
}

func (x TaskCategory) String() string {
	return proto.EnumName(TaskCategory_name, int32(x))
}
func (TaskCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tasks_1630009bae58f45a, []int{0}
}

type RuntimeMetadata_RuntimeType int32

const (
	RuntimeMetadata_FlyteSDK RuntimeMetadata_RuntimeType = 0
	RuntimeMetadata_Other    RuntimeMetadata_RuntimeType = 1
)

var RuntimeMetadata_RuntimeType_name = map[int32]string{
	0: "FlyteSDK",
	1: "Other",
}
var RuntimeMetadata_RuntimeType_value = map[string]int32{
	"FlyteSDK": 0,
	"Other":    1,
}

func (x RuntimeMetadata_RuntimeType) String() string {
	return proto.EnumName(RuntimeMetadata_RuntimeType_name, int32(x))
}
func (RuntimeMetadata_RuntimeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tasks_1630009bae58f45a, []int{1, 0}
}

// A customizable interface to convey resources requested for a container. This can be interpretted differently for different
// container engines.
type Resources struct {
	// The desired set of resources requested.
	Requests map[string]string `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Defines a set of bounds (e.g. min/max) within which the task can reliably run.
	Limits map[string]string `protobuf:"bytes,2,rep,name=limits" json:"limits,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Defines a set of bounds (e.g. min/max) within which the task can reliably run.
	EnvDictionary        map[string]string `protobuf:"bytes,3,rep,name=env_dictionary,json=envDictionary" json:"env_dictionary,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_1630009bae58f45a, []int{0}
}
func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (dst *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(dst, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetRequests() map[string]string {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *Resources) GetLimits() map[string]string {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *Resources) GetEnvDictionary() map[string]string {
	if m != nil {
		return m.EnvDictionary
	}
	return nil
}

// Runtime information. This is losely defined to allow for extensibility.
type RuntimeMetadata struct {
	// Type of runtime.
	Type RuntimeMetadata_RuntimeType `protobuf:"varint,1,opt,name=type,enum=core.RuntimeMetadata_RuntimeType" json:"type,omitempty"`
	// Version of the runtime. All versions should be backward compatible. However, certain cases call for version
	// checks to ensure tighter validation or setting expectations.
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	// +optional It can be used to provide extra information about the runtime (e.g. python, golang... etc.).
	Flavor               string   `protobuf:"bytes,3,opt,name=flavor" json:"flavor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuntimeMetadata) Reset()         { *m = RuntimeMetadata{} }
func (m *RuntimeMetadata) String() string { return proto.CompactTextString(m) }
func (*RuntimeMetadata) ProtoMessage()    {}
func (*RuntimeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_1630009bae58f45a, []int{1}
}
func (m *RuntimeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeMetadata.Unmarshal(m, b)
}
func (m *RuntimeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeMetadata.Marshal(b, m, deterministic)
}
func (dst *RuntimeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeMetadata.Merge(dst, src)
}
func (m *RuntimeMetadata) XXX_Size() int {
	return xxx_messageInfo_RuntimeMetadata.Size(m)
}
func (m *RuntimeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeMetadata proto.InternalMessageInfo

func (m *RuntimeMetadata) GetType() RuntimeMetadata_RuntimeType {
	if m != nil {
		return m.Type
	}
	return RuntimeMetadata_FlyteSDK
}

func (m *RuntimeMetadata) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RuntimeMetadata) GetFlavor() string {
	if m != nil {
		return m.Flavor
	}
	return ""
}

// Task Metadata
type TaskMetadata struct {
	// Indicates whether the system should attempt to lookup this task's output to avoid duplication of work.
	Discoverable bool `protobuf:"varint,1,opt,name=discoverable" json:"discoverable,omitempty"`
	// Runtime information about the task.
	Runtime *RuntimeMetadata `protobuf:"bytes,2,opt,name=runtime" json:"runtime,omitempty"`
	// The overall timeout of a task.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout" json:"timeout,omitempty"`
	// Number of retries per task.
	Retries              *RetryStrategy `protobuf:"bytes,5,opt,name=retries" json:"retries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TaskMetadata) Reset()         { *m = TaskMetadata{} }
func (m *TaskMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskMetadata) ProtoMessage()    {}
func (*TaskMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_1630009bae58f45a, []int{2}
}
func (m *TaskMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskMetadata.Unmarshal(m, b)
}
func (m *TaskMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskMetadata.Marshal(b, m, deterministic)
}
func (dst *TaskMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskMetadata.Merge(dst, src)
}
func (m *TaskMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskMetadata.Size(m)
}
func (m *TaskMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskMetadata proto.InternalMessageInfo

func (m *TaskMetadata) GetDiscoverable() bool {
	if m != nil {
		return m.Discoverable
	}
	return false
}

func (m *TaskMetadata) GetRuntime() *RuntimeMetadata {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *TaskMetadata) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *TaskMetadata) GetRetries() *RetryStrategy {
	if m != nil {
		return m.Retries
	}
	return nil
}

// A Task structure that uniquely identifies a task in the system
// Tasks are registered as a first step in the system.
type TaskTemplate struct {
	// Auto generated taskId by the system. Task Id uniquely identifies this task globally.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Category of the task. These are predefined and help provide defaults
	Category TaskCategory `protobuf:"varint,2,opt,name=category,enum=core.TaskCategory" json:"category,omitempty"`
	// A predefined yet extensible Task type identifier. This can be used to customize any of the components. If no
	// extensions are provided in the system, Flyte will resolve the this task to its TaskCategory and default the
	// implementation registered for the TaskCategory.
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	// Extra metadata about the task.
	Metadata *TaskMetadata `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
	// A strongly typed interface for the task. This enables others to use this task within a workflow and gauarantees
	// compile-time validation of the workflow to avoid costly runtime failures.
	Interface *TypedInterface `protobuf:"bytes,5,opt,name=interface" json:"interface,omitempty"`
	// Custom data about the task. This is extensible to allow various plugins in the system.
	Custom []byte `protobuf:"bytes,6,opt,name=custom,proto3" json:"custom,omitempty"`
	// Known task types that the system will guarantee plugins for. Custom SDK plugins are allowed to set these if needed.
	// If no corresponding execution-layer plugins are found, the system will default to handling these using built-in
	// handlers.
	//
	// Types that are valid to be assigned to Task:
	//	*TaskTemplate_Container
	Task                 isTaskTemplate_Task `protobuf_oneof:"task"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TaskTemplate) Reset()         { *m = TaskTemplate{} }
func (m *TaskTemplate) String() string { return proto.CompactTextString(m) }
func (*TaskTemplate) ProtoMessage()    {}
func (*TaskTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_1630009bae58f45a, []int{3}
}
func (m *TaskTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskTemplate.Unmarshal(m, b)
}
func (m *TaskTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskTemplate.Marshal(b, m, deterministic)
}
func (dst *TaskTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskTemplate.Merge(dst, src)
}
func (m *TaskTemplate) XXX_Size() int {
	return xxx_messageInfo_TaskTemplate.Size(m)
}
func (m *TaskTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_TaskTemplate proto.InternalMessageInfo

type isTaskTemplate_Task interface {
	isTaskTemplate_Task()
}

type TaskTemplate_Container struct {
	Container *Container `protobuf:"bytes,7,opt,name=container,oneof"`
}

func (*TaskTemplate_Container) isTaskTemplate_Task() {}

func (m *TaskTemplate) GetTask() isTaskTemplate_Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *TaskTemplate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskTemplate) GetCategory() TaskCategory {
	if m != nil {
		return m.Category
	}
	return TaskCategory_SingleStepTask
}

func (m *TaskTemplate) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TaskTemplate) GetMetadata() *TaskMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TaskTemplate) GetInterface() *TypedInterface {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *TaskTemplate) GetCustom() []byte {
	if m != nil {
		return m.Custom
	}
	return nil
}

func (m *TaskTemplate) GetContainer() *Container {
	if x, ok := m.GetTask().(*TaskTemplate_Container); ok {
		return x.Container
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TaskTemplate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TaskTemplate_OneofMarshaler, _TaskTemplate_OneofUnmarshaler, _TaskTemplate_OneofSizer, []interface{}{
		(*TaskTemplate_Container)(nil),
	}
}

func _TaskTemplate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TaskTemplate)
	// task
	switch x := m.Task.(type) {
	case *TaskTemplate_Container:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Container); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TaskTemplate.Task has unexpected type %T", x)
	}
	return nil
}

func _TaskTemplate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TaskTemplate)
	switch tag {
	case 7: // task.container
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Container)
		err := b.DecodeMessage(msg)
		m.Task = &TaskTemplate_Container{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TaskTemplate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TaskTemplate)
	// task
	switch x := m.Task.(type) {
	case *TaskTemplate_Container:
		s := proto.Size(x.Container)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Container struct {
	// Container image url. Eg: docker/redis:latest
	Image string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	// Command to be executed, if not provided, the default entrypoint in the container image will be used.
	Command []string `protobuf:"bytes,2,rep,name=command" json:"command,omitempty"`
	// These will default to Flyte given paths. If provided, the system will not append known paths. If the task still
	// needs flyte's inputs and outputs path, add $(FLYTE_INPUT_FILE), $(FLYTE_OUTPUT_FILE) wherever makes sense and the
	// system will populate these before executing the container.
	Args []string `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
	// Container resources requirement as specified by the container engine.
	Resources *Resources `protobuf:"bytes,4,opt,name=resources" json:"resources,omitempty"`
	// Environment variables will be set as the container is starting up.
	Env []*KeyValuePair `protobuf:"bytes,5,rep,name=env" json:"env,omitempty"`
	// Allows extra configs to be available for the container.
	// TODO: elaborate on how configs will become available.
	Config               []*KeyValuePair `protobuf:"bytes,6,rep,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_1630009bae58f45a, []int{4}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Container) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Container) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Container) GetEnv() []*KeyValuePair {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Container) GetConfig() []*KeyValuePair {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Resources)(nil), "core.Resources")
	proto.RegisterMapType((map[string]string)(nil), "core.Resources.EnvDictionaryEntry")
	proto.RegisterMapType((map[string]string)(nil), "core.Resources.LimitsEntry")
	proto.RegisterMapType((map[string]string)(nil), "core.Resources.RequestsEntry")
	proto.RegisterType((*RuntimeMetadata)(nil), "core.RuntimeMetadata")
	proto.RegisterType((*TaskMetadata)(nil), "core.TaskMetadata")
	proto.RegisterType((*TaskTemplate)(nil), "core.TaskTemplate")
	proto.RegisterType((*Container)(nil), "core.Container")
	proto.RegisterEnum("core.TaskCategory", TaskCategory_name, TaskCategory_value)
	proto.RegisterEnum("core.RuntimeMetadata_RuntimeType", RuntimeMetadata_RuntimeType_name, RuntimeMetadata_RuntimeType_value)
}

func init() { proto.RegisterFile("flyteidl/core/tasks.proto", fileDescriptor_tasks_1630009bae58f45a) }

var fileDescriptor_tasks_1630009bae58f45a = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xf3, 0x44,
	0x10, 0xae, 0xf3, 0xe1, 0xc4, 0x93, 0x34, 0x0d, 0x4b, 0x41, 0x6e, 0xa0, 0xa8, 0x58, 0x08, 0x55,
	0x95, 0xea, 0x48, 0xa9, 0x2a, 0x51, 0xb8, 0xa0, 0x7e, 0x20, 0xaa, 0x52, 0x81, 0x36, 0x15, 0x07,
	0x2e, 0x68, 0x6b, 0x4f, 0xc2, 0xaa, 0xb6, 0x37, 0xac, 0xd7, 0x91, 0x7c, 0xe3, 0x5f, 0x70, 0xe7,
	0xcf, 0x70, 0xe5, 0x27, 0xa1, 0x5d, 0xaf, 0x9d, 0xa4, 0xef, 0xfb, 0x1e, 0x7a, 0x8a, 0x67, 0xe6,
	0x79, 0x9e, 0xcc, 0xe7, 0xc2, 0xd1, 0x22, 0x29, 0x15, 0xf2, 0x38, 0x99, 0x46, 0x42, 0xe2, 0x54,
	0xb1, 0xfc, 0x25, 0x0f, 0x57, 0x52, 0x28, 0x41, 0x3a, 0xda, 0x33, 0x39, 0xde, 0x05, 0xf0, 0x4c,
	0xa1, 0x5c, 0xb0, 0x08, 0x2b, 0xd0, 0xe4, 0xf3, 0xdd, 0x70, 0xc2, 0x15, 0x4a, 0x96, 0x58, 0x89,
	0xc9, 0x17, 0x4b, 0x21, 0x96, 0x09, 0x4e, 0x8d, 0xf5, 0x5c, 0x2c, 0xa6, 0x71, 0x21, 0x99, 0xe2,
	0x22, 0xab, 0xe2, 0xc1, 0x5f, 0x6d, 0xf0, 0x28, 0xe6, 0xa2, 0x90, 0x11, 0xe6, 0xe4, 0x0a, 0xfa,
	0x12, 0xff, 0x2c, 0x30, 0x57, 0xb9, 0xef, 0x9c, 0xb4, 0x4f, 0x07, 0xb3, 0xe3, 0x50, 0xab, 0x86,
	0x0d, 0x24, 0xa4, 0x36, 0x7e, 0x97, 0x29, 0x59, 0xd2, 0x06, 0x4e, 0x2e, 0xc0, 0x4d, 0x78, 0xca,
	0x55, 0xee, 0xb7, 0x0c, 0xf1, 0xb3, 0xd7, 0xc4, 0x9f, 0x4c, 0xb4, 0xa2, 0x59, 0x28, 0xb9, 0x87,
	0x11, 0x66, 0xeb, 0xdf, 0x63, 0x1e, 0xe9, 0x94, 0x98, 0x2c, 0xfd, 0xb6, 0x21, 0x07, 0xaf, 0xc9,
	0x77, 0xd9, 0xfa, 0xb6, 0x01, 0x55, 0x1a, 0xfb, 0xb8, 0xed, 0x9b, 0x7c, 0x07, 0xfb, 0x3b, 0xa9,
	0x91, 0x31, 0xb4, 0x5f, 0xb0, 0xf4, 0x9d, 0x13, 0xe7, 0xd4, 0xa3, 0xfa, 0x93, 0x1c, 0x42, 0x77,
	0xcd, 0x92, 0x02, 0xfd, 0x96, 0xf1, 0x55, 0xc6, 0xb7, 0xad, 0x6f, 0x9c, 0xc9, 0x15, 0x0c, 0xb6,
	0xd2, 0x7b, 0x13, 0xf5, 0x7b, 0x20, 0xef, 0x26, 0xf7, 0x16, 0x85, 0xe0, 0x1f, 0x07, 0x0e, 0x68,
	0x91, 0x29, 0x9e, 0xe2, 0x23, 0x2a, 0x16, 0x33, 0xc5, 0xc8, 0x25, 0x74, 0x54, 0xb9, 0x42, 0x23,
	0x30, 0x9a, 0x7d, 0x69, 0xdb, 0xb1, 0x0b, 0xaa, 0xed, 0xa7, 0x72, 0x85, 0xd4, 0xc0, 0x89, 0x0f,
	0xbd, 0x35, 0xca, 0x9c, 0x8b, 0xcc, 0xfe, 0x4d, 0x6d, 0x92, 0x4f, 0xc1, 0x5d, 0x24, 0x6c, 0x2d,
	0xa4, 0xdf, 0x36, 0x01, 0x6b, 0x05, 0x5f, 0xc3, 0x60, 0x4b, 0x86, 0x0c, 0xa1, 0xff, 0x83, 0x5e,
	0xa7, 0xf9, 0xed, 0xc3, 0x78, 0x8f, 0x78, 0xd0, 0xfd, 0x59, 0xfd, 0x81, 0x72, 0xec, 0x04, 0xff,
	0x3a, 0x30, 0x7c, 0x62, 0xf9, 0x4b, 0x93, 0x61, 0x00, 0xc3, 0x98, 0xe7, 0x91, 0x58, 0xa3, 0x64,
	0xcf, 0x49, 0x95, 0x69, 0x9f, 0xee, 0xf8, 0xc8, 0x14, 0x7a, 0xb2, 0x12, 0x37, 0xe9, 0x0c, 0x66,
	0x9f, 0xbc, 0xb7, 0x10, 0x5a, 0xa3, 0xc8, 0x05, 0xf4, 0xf4, 0xaf, 0x28, 0x94, 0xdf, 0x31, 0x84,
	0xa3, 0xb0, 0xda, 0xdf, 0xb0, 0xde, 0xdf, 0xf0, 0xd6, 0xee, 0x2f, 0xad, 0x91, 0xe4, 0x1c, 0x7a,
	0x12, 0x95, 0xe4, 0x98, 0xfb, 0x5d, 0x43, 0xfa, 0xb8, 0xde, 0x1e, 0x25, 0xcb, 0xb9, 0x92, 0x4c,
	0xe1, 0xb2, 0xa4, 0x35, 0x26, 0xf8, 0xbb, 0x55, 0x55, 0xf2, 0x84, 0xe9, 0x2a, 0x61, 0x0a, 0xc9,
	0x08, 0x5a, 0x3c, 0xb6, 0xa3, 0x6a, 0xf1, 0x98, 0x84, 0xd0, 0x8f, 0x34, 0x47, 0xc8, 0xd2, 0xa4,
	0x3d, 0x9a, 0x91, 0x4a, 0x50, 0xb3, 0x6e, 0x6c, 0x84, 0x36, 0x18, 0x42, 0xec, 0xac, 0xaa, 0xc6,
	0x56, 0x83, 0x08, 0xa1, 0x9f, 0xda, 0xea, 0x6c, 0x25, 0x5b, 0x1a, 0x4d, 0xdd, 0x0d, 0x86, 0xcc,
	0xc0, 0x6b, 0xee, 0xda, 0x56, 0x71, 0x68, 0x09, 0xe5, 0x0a, 0xe3, 0xfb, 0x3a, 0x46, 0x37, 0x30,
	0x3d, 0xd2, 0xa8, 0xc8, 0x95, 0x48, 0x7d, 0xf7, 0xc4, 0x39, 0x1d, 0x52, 0x6b, 0x91, 0x29, 0x78,
	0x91, 0xc8, 0x14, 0xe3, 0x19, 0x4a, 0xbf, 0x67, 0xb4, 0x0e, 0x2a, 0xad, 0x9b, 0xda, 0xfd, 0xe3,
	0x1e, 0xdd, 0x60, 0xae, 0x5d, 0xe8, 0xe8, 0x57, 0x27, 0xf8, 0xcf, 0x01, 0xaf, 0x81, 0xe8, 0x85,
	0xe5, 0x29, 0x5b, 0xa2, 0xed, 0x4c, 0x65, 0xe8, 0x0d, 0x8b, 0x44, 0x9a, 0xb2, 0x2c, 0x36, 0x77,
	0xee, 0xd1, 0xda, 0xd4, 0x6d, 0x60, 0x72, 0x99, 0x9b, 0x0b, 0xf6, 0xa8, 0xf9, 0x26, 0xe7, 0xe0,
	0xc9, 0xfa, 0x86, 0x6d, 0x1f, 0x0e, 0x5e, 0x9d, 0x36, 0xdd, 0x20, 0xc8, 0x57, 0xd0, 0xc6, 0x6c,
	0xed, 0x77, 0xcd, 0x1b, 0x60, 0x1b, 0xf6, 0x80, 0xe5, 0xaf, 0xfa, 0x54, 0x7e, 0x61, 0x5c, 0x52,
	0x1d, 0x26, 0x67, 0xe0, 0x46, 0x22, 0x5b, 0xf0, 0xa5, 0xef, 0x7e, 0x10, 0x68, 0x11, 0x67, 0x97,
	0xd5, 0xac, 0x6f, 0x36, 0xb3, 0x1a, 0xcd, 0x79, 0xb6, 0x4c, 0x70, 0xae, 0x70, 0xa5, 0x23, 0xe3,
	0x3d, 0xf2, 0x11, 0xec, 0x3f, 0x16, 0x89, 0xe2, 0x8d, 0xcb, 0xb9, 0x76, 0x7f, 0x33, 0x4f, 0xef,
	0xb3, 0x6b, 0xd6, 0xee, 0xe2, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x5e, 0x28, 0x9e, 0xa4,
	0x05, 0x00, 0x00,
}