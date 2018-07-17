// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/future.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This document describes a set of tasks to execute and how the final outputs are produced.
type FutureTaskDocument struct {
	// A collection of tasks to execute.
	Tasks []*FutureTaskNode `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
	// An absolute number of the minimum number of successful completions of subtasks. As soon as this criteria is met,
	// the future job will be marked as successful and outputs will be computed.
	MinSuccesses int64 `protobuf:"varint,2,opt,name=min_successes,json=minSuccesses" json:"min_successes,omitempty"`
	// Describes how to bind the final output of the future task from the outputs of executed nodes. The referenced ids
	// in bindings should have the generated id for the subtask.
	Outputs              []*Binding `protobuf:"bytes,3,rep,name=outputs" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FutureTaskDocument) Reset()         { *m = FutureTaskDocument{} }
func (m *FutureTaskDocument) String() string { return proto.CompactTextString(m) }
func (*FutureTaskDocument) ProtoMessage()    {}
func (*FutureTaskDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_future_95dcd39e1aea91af, []int{0}
}
func (m *FutureTaskDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FutureTaskDocument.Unmarshal(m, b)
}
func (m *FutureTaskDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FutureTaskDocument.Marshal(b, m, deterministic)
}
func (dst *FutureTaskDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FutureTaskDocument.Merge(dst, src)
}
func (m *FutureTaskDocument) XXX_Size() int {
	return xxx_messageInfo_FutureTaskDocument.Size(m)
}
func (m *FutureTaskDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_FutureTaskDocument.DiscardUnknown(m)
}

var xxx_messageInfo_FutureTaskDocument proto.InternalMessageInfo

func (m *FutureTaskDocument) GetTasks() []*FutureTaskNode {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *FutureTaskDocument) GetMinSuccesses() int64 {
	if m != nil {
		return m.MinSuccesses
	}
	return 0
}

func (m *FutureTaskDocument) GetOutputs() []*Binding {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type FutureTaskNode struct {
	// A unique identifier prefix within the document. This is used to generate unique ids for each of the tasks in
	// target. The generated ids have this format: <generate_id>:<n> where n starts with 0 and monotonically increases.
	// e.g. if generate_id is set to "my_array_job", the first task in the array job will have the id "my_array_job:0",
	// then "my_array_job:1"... etc.
	GenerateId string `protobuf:"bytes,1,opt,name=generate_id,json=generateId" json:"generate_id,omitempty"`
	// The executable target of the node.
	//
	// Types that are valid to be assigned to Target:
	//	*FutureTaskNode_Array
	//	*FutureTaskNode_HiveQueries
	Target               isFutureTaskNode_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FutureTaskNode) Reset()         { *m = FutureTaskNode{} }
func (m *FutureTaskNode) String() string { return proto.CompactTextString(m) }
func (*FutureTaskNode) ProtoMessage()    {}
func (*FutureTaskNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_future_95dcd39e1aea91af, []int{1}
}
func (m *FutureTaskNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FutureTaskNode.Unmarshal(m, b)
}
func (m *FutureTaskNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FutureTaskNode.Marshal(b, m, deterministic)
}
func (dst *FutureTaskNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FutureTaskNode.Merge(dst, src)
}
func (m *FutureTaskNode) XXX_Size() int {
	return xxx_messageInfo_FutureTaskNode.Size(m)
}
func (m *FutureTaskNode) XXX_DiscardUnknown() {
	xxx_messageInfo_FutureTaskNode.DiscardUnknown(m)
}

var xxx_messageInfo_FutureTaskNode proto.InternalMessageInfo

type isFutureTaskNode_Target interface {
	isFutureTaskNode_Target()
}

type FutureTaskNode_Array struct {
	Array *ArrayJob `protobuf:"bytes,2,opt,name=array,oneof"`
}
type FutureTaskNode_HiveQueries struct {
	HiveQueries *HiveQueryCollection `protobuf:"bytes,3,opt,name=hive_queries,json=hiveQueries,oneof"`
}

func (*FutureTaskNode_Array) isFutureTaskNode_Target()       {}
func (*FutureTaskNode_HiveQueries) isFutureTaskNode_Target() {}

func (m *FutureTaskNode) GetTarget() isFutureTaskNode_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *FutureTaskNode) GetGenerateId() string {
	if m != nil {
		return m.GenerateId
	}
	return ""
}

func (m *FutureTaskNode) GetArray() *ArrayJob {
	if x, ok := m.GetTarget().(*FutureTaskNode_Array); ok {
		return x.Array
	}
	return nil
}

func (m *FutureTaskNode) GetHiveQueries() *HiveQueryCollection {
	if x, ok := m.GetTarget().(*FutureTaskNode_HiveQueries); ok {
		return x.HiveQueries
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FutureTaskNode) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FutureTaskNode_OneofMarshaler, _FutureTaskNode_OneofUnmarshaler, _FutureTaskNode_OneofSizer, []interface{}{
		(*FutureTaskNode_Array)(nil),
		(*FutureTaskNode_HiveQueries)(nil),
	}
}

func _FutureTaskNode_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FutureTaskNode)
	// target
	switch x := m.Target.(type) {
	case *FutureTaskNode_Array:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Array); err != nil {
			return err
		}
	case *FutureTaskNode_HiveQueries:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HiveQueries); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FutureTaskNode.Target has unexpected type %T", x)
	}
	return nil
}

func _FutureTaskNode_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FutureTaskNode)
	switch tag {
	case 2: // target.array
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ArrayJob)
		err := b.DecodeMessage(msg)
		m.Target = &FutureTaskNode_Array{msg}
		return true, err
	case 3: // target.hive_queries
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HiveQueryCollection)
		err := b.DecodeMessage(msg)
		m.Target = &FutureTaskNode_HiveQueries{msg}
		return true, err
	default:
		return false, nil
	}
}

func _FutureTaskNode_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FutureTaskNode)
	// target
	switch x := m.Target.(type) {
	case *FutureTaskNode_Array:
		s := proto.Size(x.Array)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FutureTaskNode_HiveQueries:
		s := proto.Size(x.HiveQueries)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Defines a query to execute on a hive cluster.
type HiveQuery struct {
	// The query string.
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	// Metadata to use when executing the query. Not all fields can be used with all execution engines.
	Metadata             *TaskMetadata `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HiveQuery) Reset()         { *m = HiveQuery{} }
func (m *HiveQuery) String() string { return proto.CompactTextString(m) }
func (*HiveQuery) ProtoMessage()    {}
func (*HiveQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_future_95dcd39e1aea91af, []int{2}
}
func (m *HiveQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveQuery.Unmarshal(m, b)
}
func (m *HiveQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveQuery.Marshal(b, m, deterministic)
}
func (dst *HiveQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveQuery.Merge(dst, src)
}
func (m *HiveQuery) XXX_Size() int {
	return xxx_messageInfo_HiveQuery.Size(m)
}
func (m *HiveQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveQuery.DiscardUnknown(m)
}

var xxx_messageInfo_HiveQuery proto.InternalMessageInfo

func (m *HiveQuery) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *HiveQuery) GetMetadata() *TaskMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Defines a collection of hive queries.
type HiveQueryCollection struct {
	Queries              []*HiveQuery `protobuf:"bytes,2,rep,name=queries" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HiveQueryCollection) Reset()         { *m = HiveQueryCollection{} }
func (m *HiveQueryCollection) String() string { return proto.CompactTextString(m) }
func (*HiveQueryCollection) ProtoMessage()    {}
func (*HiveQueryCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_future_95dcd39e1aea91af, []int{3}
}
func (m *HiveQueryCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveQueryCollection.Unmarshal(m, b)
}
func (m *HiveQueryCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveQueryCollection.Marshal(b, m, deterministic)
}
func (dst *HiveQueryCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveQueryCollection.Merge(dst, src)
}
func (m *HiveQueryCollection) XXX_Size() int {
	return xxx_messageInfo_HiveQueryCollection.Size(m)
}
func (m *HiveQueryCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveQueryCollection.DiscardUnknown(m)
}

var xxx_messageInfo_HiveQueryCollection proto.InternalMessageInfo

func (m *HiveQueryCollection) GetQueries() []*HiveQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

// Defines a collection of containers to execute.
type SwarmDefinition struct {
	// The primary container of a swarm job. This container starts after all the init_containers have successfully
	// completed. The end-state of the primary container determines the overall state of the swarm task.
	PrimaryContainer *Container `protobuf:"bytes,1,opt,name=primary_container,json=primaryContainer" json:"primary_container,omitempty"`
	// List of initialization containers belonging to the pod.
	// Init containers are executed in order prior to containers being started. If any
	// init container fails, the pod is considered to have failed and is handled according
	// to its restartPolicy. The name for an init container or normal container must be
	// unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, or Liveness probes.
	// The resourceRequirements of an init container are taken into account during scheduling
	// by finding the highest request/limit for each resource type, and then using the max of
	// of that value or the sum of the normal containers. Limits are applied to init containers
	// in a similar fashion.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	InitContainers []*Container `protobuf:"bytes,2,rep,name=init_containers,json=initContainers" json:"init_containers,omitempty"`
	// List of containers that get started after all init_containers have successfully completed. The end-state of
	// sidecar_containers is ignored when computing the overall state of the swarm task.
	SidecarContainers    []*Container `protobuf:"bytes,3,rep,name=sidecar_containers,json=sidecarContainers" json:"sidecar_containers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SwarmDefinition) Reset()         { *m = SwarmDefinition{} }
func (m *SwarmDefinition) String() string { return proto.CompactTextString(m) }
func (*SwarmDefinition) ProtoMessage()    {}
func (*SwarmDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_future_95dcd39e1aea91af, []int{4}
}
func (m *SwarmDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwarmDefinition.Unmarshal(m, b)
}
func (m *SwarmDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwarmDefinition.Marshal(b, m, deterministic)
}
func (dst *SwarmDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwarmDefinition.Merge(dst, src)
}
func (m *SwarmDefinition) XXX_Size() int {
	return xxx_messageInfo_SwarmDefinition.Size(m)
}
func (m *SwarmDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_SwarmDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_SwarmDefinition proto.InternalMessageInfo

func (m *SwarmDefinition) GetPrimaryContainer() *Container {
	if m != nil {
		return m.PrimaryContainer
	}
	return nil
}

func (m *SwarmDefinition) GetInitContainers() []*Container {
	if m != nil {
		return m.InitContainers
	}
	return nil
}

func (m *SwarmDefinition) GetSidecarContainers() []*Container {
	if m != nil {
		return m.SidecarContainers
	}
	return nil
}

// Describes a job that can process independent pieces of data concurrently. Multiple copies of the runnable component
// will be executed concurrently.
type ArrayJob struct {
	// Metadata about the task.
	Metadata *TaskMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Defines the maximum number of instances to bring up concurrently at any given point.
	Slots int64 `protobuf:"varint,2,opt,name=slots" json:"slots,omitempty"`
	// Defines the number of successful completions needed to mark the job as success. This number should match
	// the size of the input if the job requires processing of all input data.
	Completions int64 `protobuf:"varint,3,opt,name=completions" json:"completions,omitempty"`
	// Types that are valid to be assigned to Runnable:
	//	*ArrayJob_Container
	//	*ArrayJob_Swarm
	Runnable isArrayJob_Runnable `protobuf_oneof:"runnable"`
	// The location for where the input will be. The usage of this location is engine-dependent.
	// AWS_Batch & K8s_Batch: This location will be passed in to each task in the array job. Each job is responsible for
	// processing only the portion of the input it's meant to based on an environment variable passed into the container
	// . User code should look at BATCH_JOB_ARRAY_INDEX_VAR_NAME to know the environment variable it must look at.
	// For example, in AWS_Batch, BATCH_JOB_ARRAY_INDEX_VAR_NAME will be set to AWS_BATCH_JOB_ARRAY_INDEX. The job can
	// then look at AWS_BATCH_JOB_ARRAY_INDEX to know the index of the job and load the appropriate portion of the input.
	// Azure: The execution engine will have to process the input and slice it for each task. It'll then pass an absolute
	// location to each task for where it can find its input.
	Inputs               *DataLocation `protobuf:"bytes,6,opt,name=inputs" json:"inputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ArrayJob) Reset()         { *m = ArrayJob{} }
func (m *ArrayJob) String() string { return proto.CompactTextString(m) }
func (*ArrayJob) ProtoMessage()    {}
func (*ArrayJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_future_95dcd39e1aea91af, []int{5}
}
func (m *ArrayJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayJob.Unmarshal(m, b)
}
func (m *ArrayJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayJob.Marshal(b, m, deterministic)
}
func (dst *ArrayJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayJob.Merge(dst, src)
}
func (m *ArrayJob) XXX_Size() int {
	return xxx_messageInfo_ArrayJob.Size(m)
}
func (m *ArrayJob) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayJob.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayJob proto.InternalMessageInfo

type isArrayJob_Runnable interface {
	isArrayJob_Runnable()
}

type ArrayJob_Container struct {
	Container *Container `protobuf:"bytes,4,opt,name=container,oneof"`
}
type ArrayJob_Swarm struct {
	Swarm *SwarmDefinition `protobuf:"bytes,5,opt,name=swarm,oneof"`
}

func (*ArrayJob_Container) isArrayJob_Runnable() {}
func (*ArrayJob_Swarm) isArrayJob_Runnable()     {}

func (m *ArrayJob) GetRunnable() isArrayJob_Runnable {
	if m != nil {
		return m.Runnable
	}
	return nil
}

func (m *ArrayJob) GetMetadata() *TaskMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ArrayJob) GetSlots() int64 {
	if m != nil {
		return m.Slots
	}
	return 0
}

func (m *ArrayJob) GetCompletions() int64 {
	if m != nil {
		return m.Completions
	}
	return 0
}

func (m *ArrayJob) GetContainer() *Container {
	if x, ok := m.GetRunnable().(*ArrayJob_Container); ok {
		return x.Container
	}
	return nil
}

func (m *ArrayJob) GetSwarm() *SwarmDefinition {
	if x, ok := m.GetRunnable().(*ArrayJob_Swarm); ok {
		return x.Swarm
	}
	return nil
}

func (m *ArrayJob) GetInputs() *DataLocation {
	if m != nil {
		return m.Inputs
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ArrayJob) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ArrayJob_OneofMarshaler, _ArrayJob_OneofUnmarshaler, _ArrayJob_OneofSizer, []interface{}{
		(*ArrayJob_Container)(nil),
		(*ArrayJob_Swarm)(nil),
	}
}

func _ArrayJob_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ArrayJob)
	// runnable
	switch x := m.Runnable.(type) {
	case *ArrayJob_Container:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Container); err != nil {
			return err
		}
	case *ArrayJob_Swarm:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Swarm); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ArrayJob.Runnable has unexpected type %T", x)
	}
	return nil
}

func _ArrayJob_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ArrayJob)
	switch tag {
	case 4: // runnable.container
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Container)
		err := b.DecodeMessage(msg)
		m.Runnable = &ArrayJob_Container{msg}
		return true, err
	case 5: // runnable.swarm
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SwarmDefinition)
		err := b.DecodeMessage(msg)
		m.Runnable = &ArrayJob_Swarm{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ArrayJob_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ArrayJob)
	// runnable
	switch x := m.Runnable.(type) {
	case *ArrayJob_Container:
		s := proto.Size(x.Container)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ArrayJob_Swarm:
		s := proto.Size(x.Swarm)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DataLocation struct {
	// Refers to the blob store type where data is located.
	BlobStore BlobStore `protobuf:"varint,1,opt,name=blob_store,json=blobStore,enum=core.BlobStore" json:"blob_store,omitempty"`
	// Defines the actual location where data is
	//
	// Types that are valid to be assigned to Location:
	//	*DataLocation_Absolute
	//	*DataLocation_Prefix
	Location             isDataLocation_Location `protobuf_oneof:"location"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DataLocation) Reset()         { *m = DataLocation{} }
func (m *DataLocation) String() string { return proto.CompactTextString(m) }
func (*DataLocation) ProtoMessage()    {}
func (*DataLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_future_95dcd39e1aea91af, []int{6}
}
func (m *DataLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataLocation.Unmarshal(m, b)
}
func (m *DataLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataLocation.Marshal(b, m, deterministic)
}
func (dst *DataLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataLocation.Merge(dst, src)
}
func (m *DataLocation) XXX_Size() int {
	return xxx_messageInfo_DataLocation.Size(m)
}
func (m *DataLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_DataLocation.DiscardUnknown(m)
}

var xxx_messageInfo_DataLocation proto.InternalMessageInfo

type isDataLocation_Location interface {
	isDataLocation_Location()
}

type DataLocation_Absolute struct {
	Absolute string `protobuf:"bytes,2,opt,name=absolute,oneof"`
}
type DataLocation_Prefix struct {
	Prefix string `protobuf:"bytes,3,opt,name=prefix,oneof"`
}

func (*DataLocation_Absolute) isDataLocation_Location() {}
func (*DataLocation_Prefix) isDataLocation_Location()   {}

func (m *DataLocation) GetLocation() isDataLocation_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *DataLocation) GetBlobStore() BlobStore {
	if m != nil {
		return m.BlobStore
	}
	return BlobStore_INVALID
}

func (m *DataLocation) GetAbsolute() string {
	if x, ok := m.GetLocation().(*DataLocation_Absolute); ok {
		return x.Absolute
	}
	return ""
}

func (m *DataLocation) GetPrefix() string {
	if x, ok := m.GetLocation().(*DataLocation_Prefix); ok {
		return x.Prefix
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DataLocation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DataLocation_OneofMarshaler, _DataLocation_OneofUnmarshaler, _DataLocation_OneofSizer, []interface{}{
		(*DataLocation_Absolute)(nil),
		(*DataLocation_Prefix)(nil),
	}
}

func _DataLocation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DataLocation)
	// location
	switch x := m.Location.(type) {
	case *DataLocation_Absolute:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Absolute)
	case *DataLocation_Prefix:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Prefix)
	case nil:
	default:
		return fmt.Errorf("DataLocation.Location has unexpected type %T", x)
	}
	return nil
}

func _DataLocation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DataLocation)
	switch tag {
	case 2: // location.absolute
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Location = &DataLocation_Absolute{x}
		return true, err
	case 3: // location.prefix
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Location = &DataLocation_Prefix{x}
		return true, err
	default:
		return false, nil
	}
}

func _DataLocation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DataLocation)
	// location
	switch x := m.Location.(type) {
	case *DataLocation_Absolute:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Absolute)))
		n += len(x.Absolute)
	case *DataLocation_Prefix:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Prefix)))
		n += len(x.Prefix)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FutureTaskDocument)(nil), "core.FutureTaskDocument")
	proto.RegisterType((*FutureTaskNode)(nil), "core.FutureTaskNode")
	proto.RegisterType((*HiveQuery)(nil), "core.HiveQuery")
	proto.RegisterType((*HiveQueryCollection)(nil), "core.HiveQueryCollection")
	proto.RegisterType((*SwarmDefinition)(nil), "core.SwarmDefinition")
	proto.RegisterType((*ArrayJob)(nil), "core.ArrayJob")
	proto.RegisterType((*DataLocation)(nil), "core.DataLocation")
}

func init() { proto.RegisterFile("core/future.proto", fileDescriptor_future_95dcd39e1aea91af) }

var fileDescriptor_future_95dcd39e1aea91af = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x9b, 0xc6, 0x4d, 0x26, 0x6d, 0xda, 0x6e, 0x8b, 0x64, 0x2a, 0x24, 0x22, 0x23, 0x41,
	0xa9, 0x44, 0x2a, 0x95, 0x0b, 0x07, 0x54, 0x41, 0x5a, 0xa1, 0x80, 0x00, 0xa9, 0x5b, 0x4e, 0x5c,
	0xa2, 0xb5, 0x3d, 0x6d, 0x57, 0xb5, 0x77, 0xc3, 0xee, 0xba, 0x90, 0x2b, 0x47, 0x0e, 0xfc, 0x09,
	0x7e, 0x0d, 0xff, 0x0a, 0xed, 0xfa, 0x2b, 0x8d, 0x22, 0x71, 0xcb, 0xbc, 0x79, 0x6f, 0xe6, 0xf9,
	0x8d, 0x63, 0xd8, 0x8d, 0xa5, 0xc2, 0xe3, 0xab, 0xdc, 0xe4, 0x0a, 0x47, 0x33, 0x25, 0x8d, 0x24,
	0xeb, 0x16, 0x3a, 0xd8, 0x71, 0x0d, 0xc3, 0xf4, 0xad, 0x2e, 0xf0, 0x83, 0x3d, 0x87, 0xa4, 0xdc,
	0xa0, 0x62, 0x69, 0x09, 0x86, 0xbf, 0x3d, 0x20, 0xef, 0x9c, 0xfa, 0x0b, 0xd3, 0xb7, 0xe7, 0x32,
	0xce, 0x33, 0x14, 0x86, 0x1c, 0x41, 0xc7, 0x49, 0x03, 0x6f, 0xd8, 0x3e, 0xec, 0x9f, 0xec, 0x8f,
	0xac, 0x76, 0xd4, 0x10, 0x3f, 0xcb, 0x04, 0x69, 0x41, 0x21, 0x4f, 0x60, 0x2b, 0xe3, 0x62, 0xaa,
	0xf3, 0x38, 0x46, 0xad, 0x51, 0x07, 0x6b, 0x43, 0xef, 0xb0, 0x4d, 0x37, 0x33, 0x2e, 0x2e, 0x2b,
	0x8c, 0x3c, 0x83, 0x0d, 0x99, 0x9b, 0x59, 0x6e, 0x74, 0xd0, 0x76, 0x23, 0xb7, 0x8a, 0x91, 0x63,
	0x2e, 0x12, 0x2e, 0xae, 0x69, 0xd5, 0x0d, 0xff, 0x78, 0x30, 0xb8, 0xbf, 0x87, 0x3c, 0x86, 0xfe,
	0x35, 0x0a, 0x54, 0xcc, 0xe0, 0x94, 0x27, 0x81, 0x37, 0xf4, 0x0e, 0x7b, 0x14, 0x2a, 0xe8, 0x7d,
	0x42, 0x9e, 0x42, 0x87, 0x29, 0xc5, 0xe6, 0x6e, 0x73, 0xff, 0x64, 0x50, 0x8c, 0x7e, 0x6b, 0xa1,
	0x0f, 0x32, 0x9a, 0xb4, 0x68, 0xd1, 0x26, 0xa7, 0xb0, 0x79, 0xc3, 0xef, 0x70, 0xfa, 0x2d, 0x47,
	0xc5, 0xd1, 0x3a, 0xb1, 0xf4, 0x87, 0x05, 0x7d, 0xc2, 0xef, 0xf0, 0x22, 0x47, 0x35, 0x3f, 0x93,
	0x69, 0x8a, 0xb1, 0xe1, 0x52, 0x4c, 0x5a, 0xb4, 0x7f, 0x53, 0xc2, 0x1c, 0xf5, 0xb8, 0x0b, 0xbe,
	0x61, 0xea, 0x1a, 0x4d, 0x78, 0x01, 0xbd, 0x9a, 0x4f, 0xf6, 0xa1, 0x63, 0x27, 0xce, 0x4b, 0x67,
	0x45, 0x41, 0x46, 0xd0, 0xcd, 0xd0, 0xb0, 0x84, 0x19, 0x56, 0xfa, 0x22, 0xc5, 0x22, 0xfb, 0x5c,
	0x9f, 0xca, 0x0e, 0xad, 0x39, 0xe1, 0x1b, 0xd8, 0x5b, 0x61, 0x81, 0x3c, 0x87, 0x8d, 0xca, 0xee,
	0x9a, 0x0b, 0x6e, 0x7b, 0xc9, 0x2e, 0xad, 0xfa, 0xe1, 0x5f, 0x0f, 0xb6, 0x2f, 0xbf, 0x33, 0x95,
	0x9d, 0xe3, 0x15, 0x17, 0xdc, 0xc9, 0x5f, 0xc3, 0xee, 0x4c, 0xf1, 0x8c, 0xa9, 0xf9, 0x34, 0x96,
	0xc2, 0x30, 0x2e, 0x50, 0x39, 0x9f, 0xf5, 0xa0, 0xb3, 0x0a, 0xa6, 0x3b, 0x25, 0xb3, 0x46, 0xc8,
	0x2b, 0xd8, 0xb6, 0x83, 0x1a, 0xe9, 0x92, 0x89, 0x46, 0x3b, 0xb0, 0xbc, 0xba, 0xd4, 0xe4, 0x14,
	0x88, 0xe6, 0x09, 0xc6, 0x4c, 0x2d, 0x8a, 0xdb, 0xab, 0xc5, 0xbb, 0x25, 0xb5, 0xd1, 0x87, 0xbf,
	0xd6, 0xa0, 0x5b, 0x1d, 0xf0, 0x5e, 0x94, 0xde, 0xff, 0xa3, 0xb4, 0x07, 0xd1, 0xa9, 0x34, 0xd5,
	0x9b, 0x58, 0x14, 0x64, 0x08, 0xfd, 0x58, 0x66, 0xb3, 0x14, 0x6d, 0x30, 0xc5, 0xf1, 0xdb, 0x74,
	0x11, 0x22, 0xc7, 0xd0, 0x6b, 0x42, 0x5a, 0x5f, 0x19, 0xd2, 0xa4, 0x45, 0x1b, 0x0e, 0x79, 0x01,
	0x1d, 0x6d, 0x03, 0x0f, 0x3a, 0x8e, 0xfc, 0xa0, 0x20, 0x2f, 0xdd, 0xc0, 0xbe, 0x7f, 0x8e, 0x45,
	0x8e, 0xc0, 0xe7, 0xc2, 0xfd, 0x07, 0xfc, 0xc5, 0xa7, 0x38, 0x67, 0x86, 0x7d, 0x94, 0x31, 0xb3,
	0x64, 0x5a, 0x32, 0xc6, 0x00, 0x5d, 0x95, 0x0b, 0xc1, 0xa2, 0x14, 0xc3, 0x9f, 0x1e, 0x6c, 0x2e,
	0x92, 0xc8, 0x08, 0x20, 0x4a, 0x65, 0x34, 0xd5, 0x46, 0x2a, 0x74, 0x91, 0x0c, 0x2a, 0xa7, 0xe3,
	0x54, 0x46, 0x97, 0x16, 0xa6, 0xbd, 0xa8, 0xfa, 0x49, 0x1e, 0x41, 0x97, 0x45, 0x5a, 0xa6, 0xb9,
	0x41, 0x97, 0x49, 0x6f, 0xd2, 0xa2, 0x35, 0x42, 0x02, 0xf0, 0x67, 0x0a, 0xaf, 0xf8, 0x0f, 0x97,
	0x89, 0xed, 0x95, 0xb5, 0x35, 0x91, 0x96, 0x3b, 0xc7, 0xfe, 0x57, 0xf7, 0x61, 0x89, 0x7c, 0xf7,
	0xe1, 0x78, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x84, 0x0c, 0x95, 0x0e, 0x7a, 0x04, 0x00, 0x00,
}