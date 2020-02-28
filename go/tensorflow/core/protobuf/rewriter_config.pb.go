// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/rewriter_config.proto

package tensorflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	framework "github.com/jadekler/tensorflow-serving-client/v2/go/tensorflow/core/framework"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RewriterConfig_Toggle int32

const (
	RewriterConfig_DEFAULT RewriterConfig_Toggle = 0
	RewriterConfig_ON      RewriterConfig_Toggle = 1
	RewriterConfig_OFF     RewriterConfig_Toggle = 2
	// Enable some aggressive optimizations that use assumptions that TF graphs
	// may break. For example, assume the shape of a placeholder matches its
	// actual feed.
	RewriterConfig_AGGRESSIVE RewriterConfig_Toggle = 3
)

var RewriterConfig_Toggle_name = map[int32]string{
	0: "DEFAULT",
	1: "ON",
	2: "OFF",
	3: "AGGRESSIVE",
}

var RewriterConfig_Toggle_value = map[string]int32{
	"DEFAULT":    0,
	"ON":         1,
	"OFF":        2,
	"AGGRESSIVE": 3,
}

func (x RewriterConfig_Toggle) String() string {
	return proto.EnumName(RewriterConfig_Toggle_name, int32(x))
}

func (RewriterConfig_Toggle) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{2, 0}
}

// Enum controlling the number of times to run optimizers. The default is to
// run them twice.
type RewriterConfig_NumIterationsType int32

const (
	RewriterConfig_DEFAULT_NUM_ITERS RewriterConfig_NumIterationsType = 0
	RewriterConfig_ONE               RewriterConfig_NumIterationsType = 1
	RewriterConfig_TWO               RewriterConfig_NumIterationsType = 2
)

var RewriterConfig_NumIterationsType_name = map[int32]string{
	0: "DEFAULT_NUM_ITERS",
	1: "ONE",
	2: "TWO",
}

var RewriterConfig_NumIterationsType_value = map[string]int32{
	"DEFAULT_NUM_ITERS": 0,
	"ONE":               1,
	"TWO":               2,
}

func (x RewriterConfig_NumIterationsType) String() string {
	return proto.EnumName(RewriterConfig_NumIterationsType_name, int32(x))
}

func (RewriterConfig_NumIterationsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{2, 1}
}

type RewriterConfig_MemOptType int32

const (
	// The default setting (SCHEDULING and SWAPPING HEURISTICS only)
	RewriterConfig_DEFAULT_MEM_OPT RewriterConfig_MemOptType = 0
	// Disabled in the meta-optimizer.
	RewriterConfig_NO_MEM_OPT RewriterConfig_MemOptType = 1
	// Driven by manual op-level annotations.
	RewriterConfig_MANUAL RewriterConfig_MemOptType = 2
	// Swapping heuristic will move a tensor from the GPU to the CPU and move
	// it back when needed to reduce peak memory usage.
	RewriterConfig_SWAPPING_HEURISTICS RewriterConfig_MemOptType = 4
	// Recomputation heuristics will recompute ops (such as Relu activation)
	// during backprop instead of storing them, reducing peak memory usage.
	RewriterConfig_RECOMPUTATION_HEURISTICS RewriterConfig_MemOptType = 5
	// Scheduling will split big ops such as AddN and try to enforce a schedule
	// of the new computations that decreases peak memory usage.
	RewriterConfig_SCHEDULING_HEURISTICS RewriterConfig_MemOptType = 6
	// Use any combination of swapping and recomputation heuristics.
	RewriterConfig_HEURISTICS RewriterConfig_MemOptType = 3
)

var RewriterConfig_MemOptType_name = map[int32]string{
	0: "DEFAULT_MEM_OPT",
	1: "NO_MEM_OPT",
	2: "MANUAL",
	4: "SWAPPING_HEURISTICS",
	5: "RECOMPUTATION_HEURISTICS",
	6: "SCHEDULING_HEURISTICS",
	3: "HEURISTICS",
}

var RewriterConfig_MemOptType_value = map[string]int32{
	"DEFAULT_MEM_OPT":          0,
	"NO_MEM_OPT":               1,
	"MANUAL":                   2,
	"SWAPPING_HEURISTICS":      4,
	"RECOMPUTATION_HEURISTICS": 5,
	"SCHEDULING_HEURISTICS":    6,
	"HEURISTICS":               3,
}

func (x RewriterConfig_MemOptType) String() string {
	return proto.EnumName(RewriterConfig_MemOptType_name, int32(x))
}

func (RewriterConfig_MemOptType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{2, 2}
}

type AutoParallelOptions struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	NumReplicas          int32    `protobuf:"varint,2,opt,name=num_replicas,json=numReplicas,proto3" json:"num_replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoParallelOptions) Reset()         { *m = AutoParallelOptions{} }
func (m *AutoParallelOptions) String() string { return proto.CompactTextString(m) }
func (*AutoParallelOptions) ProtoMessage()    {}
func (*AutoParallelOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{0}
}

func (m *AutoParallelOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoParallelOptions.Unmarshal(m, b)
}
func (m *AutoParallelOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoParallelOptions.Marshal(b, m, deterministic)
}
func (m *AutoParallelOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoParallelOptions.Merge(m, src)
}
func (m *AutoParallelOptions) XXX_Size() int {
	return xxx_messageInfo_AutoParallelOptions.Size(m)
}
func (m *AutoParallelOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoParallelOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AutoParallelOptions proto.InternalMessageInfo

func (m *AutoParallelOptions) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *AutoParallelOptions) GetNumReplicas() int32 {
	if m != nil {
		return m.NumReplicas
	}
	return 0
}

type ScopedAllocatorOptions struct {
	// If present, only perform optimization for these ops.
	EnableOp             []string `protobuf:"bytes,1,rep,name=enable_op,json=enableOp,proto3" json:"enable_op,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScopedAllocatorOptions) Reset()         { *m = ScopedAllocatorOptions{} }
func (m *ScopedAllocatorOptions) String() string { return proto.CompactTextString(m) }
func (*ScopedAllocatorOptions) ProtoMessage()    {}
func (*ScopedAllocatorOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{1}
}

func (m *ScopedAllocatorOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScopedAllocatorOptions.Unmarshal(m, b)
}
func (m *ScopedAllocatorOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScopedAllocatorOptions.Marshal(b, m, deterministic)
}
func (m *ScopedAllocatorOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopedAllocatorOptions.Merge(m, src)
}
func (m *ScopedAllocatorOptions) XXX_Size() int {
	return xxx_messageInfo_ScopedAllocatorOptions.Size(m)
}
func (m *ScopedAllocatorOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopedAllocatorOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ScopedAllocatorOptions proto.InternalMessageInfo

func (m *ScopedAllocatorOptions) GetEnableOp() []string {
	if m != nil {
		return m.EnableOp
	}
	return nil
}

type RewriterConfig struct {
	// Optimize tensor layouts (default is ON)
	// e.g. This will try to use NCHW layout on GPU which is faster.
	LayoutOptimizer RewriterConfig_Toggle `protobuf:"varint,1,opt,name=layout_optimizer,json=layoutOptimizer,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"layout_optimizer,omitempty"`
	// Fold constants (default is ON)
	// Statically infer the value of tensors when possible, and materialize the
	// result using constants.
	ConstantFolding RewriterConfig_Toggle `protobuf:"varint,3,opt,name=constant_folding,json=constantFolding,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"constant_folding,omitempty"`
	// Shape optimizations (default is ON)
	// Simplify computations made on shapes.
	ShapeOptimization RewriterConfig_Toggle `protobuf:"varint,13,opt,name=shape_optimization,json=shapeOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"shape_optimization,omitempty"`
	// Remapping (default is ON)
	// Remap subgraphs onto more efficient implementations.
	Remapping RewriterConfig_Toggle `protobuf:"varint,14,opt,name=remapping,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"remapping,omitempty"`
	// Arithmetic optimizations (default is ON)
	// e.g. Simplify arithmetic ops; merge ops with same value (like constants).
	ArithmeticOptimization RewriterConfig_Toggle `protobuf:"varint,7,opt,name=arithmetic_optimization,json=arithmeticOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"arithmetic_optimization,omitempty"`
	// Control dependency optimizations (default is ON).
	// Remove redundant control dependencies, which may enable other optimization.
	DependencyOptimization RewriterConfig_Toggle `protobuf:"varint,8,opt,name=dependency_optimization,json=dependencyOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"dependency_optimization,omitempty"`
	// Loop optimizations (default is ON).
	LoopOptimization RewriterConfig_Toggle `protobuf:"varint,9,opt,name=loop_optimization,json=loopOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"loop_optimization,omitempty"`
	// Function optimizations (default is ON).
	FunctionOptimization RewriterConfig_Toggle `protobuf:"varint,10,opt,name=function_optimization,json=functionOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"function_optimization,omitempty"`
	// Strips debug-related nodes from the graph (off by default).
	DebugStripper RewriterConfig_Toggle `protobuf:"varint,11,opt,name=debug_stripper,json=debugStripper,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"debug_stripper,omitempty"`
	// If true, don't remove unnecessary ops from the graph
	DisableModelPruning bool `protobuf:"varint,2,opt,name=disable_model_pruning,json=disableModelPruning,proto3" json:"disable_model_pruning,omitempty"`
	// Try to allocate some independent Op outputs contiguously in order to
	// merge or eliminate downstream Ops (off by default).
	ScopedAllocatorOptimization RewriterConfig_Toggle `protobuf:"varint,15,opt,name=scoped_allocator_optimization,json=scopedAllocatorOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"scoped_allocator_optimization,omitempty"`
	// Force small ops onto the CPU (default is OFF).
	PinToHostOptimization RewriterConfig_Toggle `protobuf:"varint,18,opt,name=pin_to_host_optimization,json=pinToHostOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"pin_to_host_optimization,omitempty"`
	// Enable the swap of kernel implementations based on the device placement
	// (default is ON).
	ImplementationSelector RewriterConfig_Toggle `protobuf:"varint,22,opt,name=implementation_selector,json=implementationSelector,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"implementation_selector,omitempty"`
	// Optimize data types (default is OFF).
	// e.g., This will try to use float16 on GPU which is faster.
	// Note that this can change the numerical stability of the graph and may
	// require the use of loss scaling to maintain model convergence.
	AutoMixedPrecision RewriterConfig_Toggle `protobuf:"varint,23,opt,name=auto_mixed_precision,json=autoMixedPrecision,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"auto_mixed_precision,omitempty"`
	// Disable the entire meta optimizer (off by default).
	DisableMetaOptimizer bool `protobuf:"varint,19,opt,name=disable_meta_optimizer,json=disableMetaOptimizer,proto3" json:"disable_meta_optimizer,omitempty"`
	// Controls how many times we run the optimizers in meta optimizer (default
	// is once).
	MetaOptimizerIterations RewriterConfig_NumIterationsType `protobuf:"varint,12,opt,name=meta_optimizer_iterations,json=metaOptimizerIterations,proto3,enum=tensorflow.RewriterConfig_NumIterationsType" json:"meta_optimizer_iterations,omitempty"`
	// The minimum number of nodes in a graph to optimizer. For smaller graphs,
	// optimization is skipped.
	// 0 means the system picks an appropriate number.
	// < 0 means do not skip optimization.
	MinGraphNodes int32 `protobuf:"varint,17,opt,name=min_graph_nodes,json=minGraphNodes,proto3" json:"min_graph_nodes,omitempty"`
	// Configures memory optimization passes through the meta-optimizer. Has no
	// effect on manually requested memory optimization passes in the optimizers
	// field.
	MemoryOptimization RewriterConfig_MemOptType `protobuf:"varint,4,opt,name=memory_optimization,json=memoryOptimization,proto3,enum=tensorflow.RewriterConfig_MemOptType" json:"memory_optimization,omitempty"`
	// A node name scope for node names which are valid outputs of recompuations.
	// Inputs to nodes that match this scope may be recomputed (subject either to
	// manual annotation of those input nodes or to manual annotation and
	// heuristics depending on memory_optimization), but the nodes themselves will
	// not be recomputed. This matches any sub-scopes as well, meaning the scope
	// can appear not just as a top-level scope. For example, if the value is
	// "gradients/", the default, it will match node name "gradients/foo",
	// "foo/gradients/bar", but not "foo_gradients/"
	MemoryOptimizerTargetNodeNameScope string `protobuf:"bytes,6,opt,name=memory_optimizer_target_node_name_scope,json=memoryOptimizerTargetNodeNameScope,proto3" json:"memory_optimizer_target_node_name_scope,omitempty"`
	// Maximum number of milliseconds to spend optimizing a single graph before
	// timing out. If equal to 0 the system picks a default (currently 5 minutes).
	// If less than 0 the optimizer will never time out.
	MetaOptimizerTimeoutMs int64 `protobuf:"varint,20,opt,name=meta_optimizer_timeout_ms,json=metaOptimizerTimeoutMs,proto3" json:"meta_optimizer_timeout_ms,omitempty"`
	// Configures AutoParallel optimization passes either through the
	// meta-optimizer or when manually specified through the optimizers field.
	AutoParallel *AutoParallelOptions `protobuf:"bytes,5,opt,name=auto_parallel,json=autoParallel,proto3" json:"auto_parallel,omitempty"`
	// If true, any optimization pass failing will cause the MetaOptimizer to
	// stop with an error. By default - or when set to false, failing passes are
	// skipped silently.
	FailOnOptimizerErrors bool                    `protobuf:"varint,21,opt,name=fail_on_optimizer_errors,json=failOnOptimizerErrors,proto3" json:"fail_on_optimizer_errors,omitempty"`
	ScopedAllocatorOpts   *ScopedAllocatorOptions `protobuf:"bytes,16,opt,name=scoped_allocator_opts,json=scopedAllocatorOpts,proto3" json:"scoped_allocator_opts,omitempty"`
	// If non-empty, will use this as an alternative way to specify a list of
	// optimizations to turn on and the order of the optimizations (replacing the
	// meta-optimizer).
	//
	// Of the RewriterConfig options, only the AutoParallel configuration options
	// (the auto_parallel field) apply to manually requested optimization passes
	// ("autoparallel"). Memory optimization passes ("memory") invoked here are
	// not configurable (in contrast to memory optimization passes through the
	// meta-optimizer) and act only on manual op annotations.
	//
	// Custom optimizers (see custom_optimizers) that are not part of this
	// schedule will be run after - in the order that they were specified.
	Optimizers []string `protobuf:"bytes,100,rep,name=optimizers,proto3" json:"optimizers,omitempty"`
	// list of CustomGraphOptimizers to apply.
	CustomOptimizers []*RewriterConfig_CustomGraphOptimizer `protobuf:"bytes,200,rep,name=custom_optimizers,json=customOptimizers,proto3" json:"custom_optimizers,omitempty"`
	// VerifierConfig specifying the verifiers to be run after every optimizer.
	InterOptimizerVerifierConfig *VerifierConfig `protobuf:"bytes,300,opt,name=inter_optimizer_verifier_config,json=interOptimizerVerifierConfig,proto3" json:"inter_optimizer_verifier_config,omitempty"`
	// VerifierConfig specifying the verifiers to be run at the end, after all
	// optimizers have run.
	PostOptimizationVerifierConfig *VerifierConfig `protobuf:"bytes,301,opt,name=post_optimization_verifier_config,json=postOptimizationVerifierConfig,proto3" json:"post_optimization_verifier_config,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}        `json:"-"`
	XXX_unrecognized               []byte          `json:"-"`
	XXX_sizecache                  int32           `json:"-"`
}

func (m *RewriterConfig) Reset()         { *m = RewriterConfig{} }
func (m *RewriterConfig) String() string { return proto.CompactTextString(m) }
func (*RewriterConfig) ProtoMessage()    {}
func (*RewriterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{2}
}

func (m *RewriterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewriterConfig.Unmarshal(m, b)
}
func (m *RewriterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewriterConfig.Marshal(b, m, deterministic)
}
func (m *RewriterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewriterConfig.Merge(m, src)
}
func (m *RewriterConfig) XXX_Size() int {
	return xxx_messageInfo_RewriterConfig.Size(m)
}
func (m *RewriterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RewriterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RewriterConfig proto.InternalMessageInfo

func (m *RewriterConfig) GetLayoutOptimizer() RewriterConfig_Toggle {
	if m != nil {
		return m.LayoutOptimizer
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetConstantFolding() RewriterConfig_Toggle {
	if m != nil {
		return m.ConstantFolding
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetShapeOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.ShapeOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetRemapping() RewriterConfig_Toggle {
	if m != nil {
		return m.Remapping
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetArithmeticOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.ArithmeticOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetDependencyOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.DependencyOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetLoopOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.LoopOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetFunctionOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.FunctionOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetDebugStripper() RewriterConfig_Toggle {
	if m != nil {
		return m.DebugStripper
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetDisableModelPruning() bool {
	if m != nil {
		return m.DisableModelPruning
	}
	return false
}

func (m *RewriterConfig) GetScopedAllocatorOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.ScopedAllocatorOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetPinToHostOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.PinToHostOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetImplementationSelector() RewriterConfig_Toggle {
	if m != nil {
		return m.ImplementationSelector
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetAutoMixedPrecision() RewriterConfig_Toggle {
	if m != nil {
		return m.AutoMixedPrecision
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetDisableMetaOptimizer() bool {
	if m != nil {
		return m.DisableMetaOptimizer
	}
	return false
}

func (m *RewriterConfig) GetMetaOptimizerIterations() RewriterConfig_NumIterationsType {
	if m != nil {
		return m.MetaOptimizerIterations
	}
	return RewriterConfig_DEFAULT_NUM_ITERS
}

func (m *RewriterConfig) GetMinGraphNodes() int32 {
	if m != nil {
		return m.MinGraphNodes
	}
	return 0
}

func (m *RewriterConfig) GetMemoryOptimization() RewriterConfig_MemOptType {
	if m != nil {
		return m.MemoryOptimization
	}
	return RewriterConfig_DEFAULT_MEM_OPT
}

func (m *RewriterConfig) GetMemoryOptimizerTargetNodeNameScope() string {
	if m != nil {
		return m.MemoryOptimizerTargetNodeNameScope
	}
	return ""
}

func (m *RewriterConfig) GetMetaOptimizerTimeoutMs() int64 {
	if m != nil {
		return m.MetaOptimizerTimeoutMs
	}
	return 0
}

func (m *RewriterConfig) GetAutoParallel() *AutoParallelOptions {
	if m != nil {
		return m.AutoParallel
	}
	return nil
}

func (m *RewriterConfig) GetFailOnOptimizerErrors() bool {
	if m != nil {
		return m.FailOnOptimizerErrors
	}
	return false
}

func (m *RewriterConfig) GetScopedAllocatorOpts() *ScopedAllocatorOptions {
	if m != nil {
		return m.ScopedAllocatorOpts
	}
	return nil
}

func (m *RewriterConfig) GetOptimizers() []string {
	if m != nil {
		return m.Optimizers
	}
	return nil
}

func (m *RewriterConfig) GetCustomOptimizers() []*RewriterConfig_CustomGraphOptimizer {
	if m != nil {
		return m.CustomOptimizers
	}
	return nil
}

func (m *RewriterConfig) GetInterOptimizerVerifierConfig() *VerifierConfig {
	if m != nil {
		return m.InterOptimizerVerifierConfig
	}
	return nil
}

func (m *RewriterConfig) GetPostOptimizationVerifierConfig() *VerifierConfig {
	if m != nil {
		return m.PostOptimizationVerifierConfig
	}
	return nil
}

// Message to describe custom graph optimizer and its parameters
type RewriterConfig_CustomGraphOptimizer struct {
	Name                 string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParameterMap         map[string]*framework.AttrValue `protobuf:"bytes,2,rep,name=parameter_map,json=parameterMap,proto3" json:"parameter_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *RewriterConfig_CustomGraphOptimizer) Reset()         { *m = RewriterConfig_CustomGraphOptimizer{} }
func (m *RewriterConfig_CustomGraphOptimizer) String() string { return proto.CompactTextString(m) }
func (*RewriterConfig_CustomGraphOptimizer) ProtoMessage()    {}
func (*RewriterConfig_CustomGraphOptimizer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{2, 0}
}

func (m *RewriterConfig_CustomGraphOptimizer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewriterConfig_CustomGraphOptimizer.Unmarshal(m, b)
}
func (m *RewriterConfig_CustomGraphOptimizer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewriterConfig_CustomGraphOptimizer.Marshal(b, m, deterministic)
}
func (m *RewriterConfig_CustomGraphOptimizer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewriterConfig_CustomGraphOptimizer.Merge(m, src)
}
func (m *RewriterConfig_CustomGraphOptimizer) XXX_Size() int {
	return xxx_messageInfo_RewriterConfig_CustomGraphOptimizer.Size(m)
}
func (m *RewriterConfig_CustomGraphOptimizer) XXX_DiscardUnknown() {
	xxx_messageInfo_RewriterConfig_CustomGraphOptimizer.DiscardUnknown(m)
}

var xxx_messageInfo_RewriterConfig_CustomGraphOptimizer proto.InternalMessageInfo

func (m *RewriterConfig_CustomGraphOptimizer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RewriterConfig_CustomGraphOptimizer) GetParameterMap() map[string]*framework.AttrValue {
	if m != nil {
		return m.ParameterMap
	}
	return nil
}

func init() {
	proto.RegisterEnum("tensorflow.RewriterConfig_Toggle", RewriterConfig_Toggle_name, RewriterConfig_Toggle_value)
	proto.RegisterEnum("tensorflow.RewriterConfig_NumIterationsType", RewriterConfig_NumIterationsType_name, RewriterConfig_NumIterationsType_value)
	proto.RegisterEnum("tensorflow.RewriterConfig_MemOptType", RewriterConfig_MemOptType_name, RewriterConfig_MemOptType_value)
	proto.RegisterType((*AutoParallelOptions)(nil), "tensorflow.AutoParallelOptions")
	proto.RegisterType((*ScopedAllocatorOptions)(nil), "tensorflow.ScopedAllocatorOptions")
	proto.RegisterType((*RewriterConfig)(nil), "tensorflow.RewriterConfig")
	proto.RegisterType((*RewriterConfig_CustomGraphOptimizer)(nil), "tensorflow.RewriterConfig.CustomGraphOptimizer")
	proto.RegisterMapType((map[string]*framework.AttrValue)(nil), "tensorflow.RewriterConfig.CustomGraphOptimizer.ParameterMapEntry")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/rewriter_config.proto", fileDescriptor_1dd7de60bf190bbb)
}

var fileDescriptor_1dd7de60bf190bbb = []byte{
	// 1206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xeb, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0x2b, 0xbb, 0x71, 0x9b, 0x93, 0x9b, 0x4c, 0xdb, 0x89, 0x9a, 0x76, 0x6d, 0x6a, 0x60,
	0x5b, 0xb0, 0x0d, 0x0e, 0x90, 0xdd, 0x87, 0x02, 0x83, 0x9b, 0x3a, 0x89, 0x81, 0xd8, 0x32, 0x64,
	0x27, 0x05, 0x0a, 0x0c, 0x04, 0x23, 0xd3, 0x8e, 0x50, 0x51, 0x24, 0x28, 0xaa, 0x5d, 0xf6, 0x32,
	0x7b, 0x81, 0xed, 0x3d, 0xf6, 0x28, 0xc3, 0x9e, 0x60, 0x1f, 0x07, 0x52, 0xbe, 0x48, 0x76, 0xd6,
	0x79, 0xdf, 0x24, 0x1e, 0xfe, 0x7f, 0xfc, 0xf3, 0x76, 0x78, 0xa0, 0xa1, 0x68, 0x14, 0x73, 0x39,
	0x0a, 0xf9, 0xfb, 0x23, 0x9f, 0x4b, 0x7a, 0x24, 0x24, 0x57, 0xfc, 0x3a, 0x19, 0x1d, 0x49, 0xfa,
	0x5e, 0x06, 0x8a, 0x4a, 0xec, 0xf3, 0x68, 0x14, 0x8c, 0x1b, 0x26, 0x80, 0x60, 0xde, 0x7f, 0xff,
	0xb3, 0x45, 0xed, 0x48, 0x12, 0x46, 0xdf, 0x73, 0xf9, 0xf6, 0x88, 0x28, 0x25, 0xf1, 0x3b, 0x12,
	0x26, 0x34, 0xd5, 0xed, 0xff, 0xfb, 0x38, 0xef, 0xa8, 0x0c, 0x46, 0xc1, 0xc2, 0x38, 0xf5, 0x1e,
	0x54, 0x9a, 0x89, 0xe2, 0x3d, 0x22, 0x49, 0x18, 0xd2, 0xd0, 0x15, 0x2a, 0xe0, 0x51, 0x8c, 0x76,
	0xa1, 0x44, 0x23, 0x72, 0x1d, 0x52, 0xc7, 0x3a, 0xb0, 0x0e, 0x1f, 0x7a, 0x93, 0x3f, 0xf4, 0x1c,
	0x36, 0xa3, 0x84, 0x61, 0x49, 0x45, 0x18, 0xf8, 0x24, 0x76, 0x0a, 0x07, 0xd6, 0xe1, 0x9a, 0xb7,
	0x11, 0x25, 0xcc, 0x9b, 0x34, 0xd5, 0xbf, 0x86, 0xdd, 0xbe, 0xcf, 0x05, 0x1d, 0x36, 0xc3, 0x90,
	0xfb, 0x44, 0x71, 0x39, 0x85, 0x3e, 0x86, 0xf5, 0x14, 0x83, 0xb9, 0x70, 0xac, 0x83, 0xe2, 0xe1,
	0xba, 0xf7, 0x30, 0x6d, 0x70, 0x45, 0xfd, 0xcf, 0x2a, 0x6c, 0x7b, 0x93, 0xa5, 0x38, 0x31, 0x0e,
	0xd1, 0x05, 0xd8, 0x21, 0xb9, 0xe5, 0x89, 0xc2, 0x5c, 0xa8, 0x80, 0x05, 0xbf, 0x50, 0x69, 0xec,
	0x6c, 0x1f, 0x3f, 0xcf, 0x4c, 0xb3, 0x91, 0x57, 0x35, 0x06, 0x7c, 0x3c, 0x0e, 0xa9, 0xb7, 0x93,
	0x4a, 0xdd, 0xa9, 0x52, 0xd3, 0x7c, 0x1e, 0xc5, 0x8a, 0x44, 0x0a, 0x8f, 0x78, 0x38, 0x0c, 0xa2,
	0xb1, 0x53, 0x5c, 0x99, 0x36, 0x95, 0x9e, 0xa6, 0x4a, 0xd4, 0x03, 0x14, 0xdf, 0x10, 0x41, 0xa7,
	0xd6, 0x88, 0x9e, 0xa2, 0xb3, 0xb5, 0x2a, 0xaf, 0x6c, 0xc4, 0x6e, 0x46, 0x8b, 0x7e, 0x84, 0x75,
	0x49, 0x19, 0x11, 0x42, 0x1b, 0xdb, 0x5e, 0x15, 0x34, 0xd7, 0xa0, 0x37, 0xb0, 0x47, 0x64, 0xa0,
	0x6e, 0x18, 0x55, 0x81, 0x9f, 0xf7, 0xf5, 0x60, 0x55, 0xdc, 0xee, 0x9c, 0x90, 0x33, 0xf7, 0x06,
	0xf6, 0x86, 0x54, 0xd0, 0x68, 0x48, 0x23, 0xff, 0x36, 0xcf, 0x7e, 0xb8, 0x32, 0x7b, 0x4e, 0xc8,
	0xb1, 0xbb, 0x50, 0x0e, 0x39, 0x17, 0x79, 0xea, 0xfa, 0xaa, 0x54, 0x5b, 0x6b, 0x73, 0xbc, 0x2b,
	0xa8, 0x8d, 0x92, 0xc8, 0xd7, 0xdf, 0x79, 0x26, 0xac, 0xca, 0xac, 0x4e, 0xf5, 0x39, 0xee, 0x39,
	0x6c, 0x0f, 0xe9, 0x75, 0x32, 0xc6, 0xb1, 0x92, 0x81, 0x10, 0x54, 0x3a, 0x1b, 0xab, 0x02, 0xb7,
	0x8c, 0xb0, 0x3f, 0xd1, 0xa1, 0x63, 0xa8, 0x0d, 0x83, 0xd8, 0xdc, 0x04, 0xc6, 0x87, 0x34, 0xc4,
	0x42, 0x26, 0x91, 0xde, 0xf6, 0x82, 0xb9, 0x6c, 0x95, 0x49, 0xb0, 0xa3, 0x63, 0xbd, 0x34, 0x84,
	0x28, 0x7c, 0x14, 0x9b, 0x6b, 0x85, 0xc9, 0xf4, 0x5e, 0xe5, 0x67, 0xb7, 0xb3, 0xaa, 0x99, 0xc7,
	0xf1, 0xf2, 0xf5, 0xcc, 0x6c, 0xb4, 0x23, 0x82, 0x08, 0x2b, 0x8e, 0x6f, 0x78, 0xac, 0xf2, 0x23,
	0xa0, 0x55, 0x47, 0xa8, 0x89, 0x20, 0x1a, 0xf0, 0x73, 0x1e, 0xab, 0xc5, 0x43, 0x14, 0x30, 0x11,
	0x52, 0x46, 0x23, 0x65, 0x5a, 0x70, 0x4c, 0x43, 0xea, 0x2b, 0x2e, 0x9d, 0xdd, 0x95, 0x0f, 0x51,
	0x9e, 0xd0, 0x9f, 0x00, 0x50, 0x1f, 0xaa, 0x24, 0x51, 0x1c, 0xb3, 0xe0, 0x67, 0x3a, 0xc4, 0x42,
	0x52, 0x3f, 0x88, 0xb5, 0xe7, 0xbd, 0x55, 0xc1, 0x48, 0xcb, 0x3b, 0x5a, 0xdd, 0x9b, 0x8a, 0xd1,
	0x57, 0xb0, 0x3b, 0xdb, 0x27, 0xaa, 0x48, 0x26, 0x0d, 0x55, 0xcc, 0x46, 0x55, 0xa7, 0x1b, 0x45,
	0x15, 0x99, 0x27, 0x9a, 0x1b, 0x78, 0x94, 0xef, 0x8d, 0xf5, 0x68, 0xc6, 0x6e, 0xec, 0x6c, 0x1a,
	0x3f, 0x5f, 0x7c, 0xc0, 0x4f, 0x37, 0x61, 0xed, 0x59, 0xff, 0xc1, 0xad, 0xa0, 0xde, 0x1e, 0xcb,
	0xf2, 0xe7, 0x41, 0xf4, 0x09, 0xec, 0xb0, 0x20, 0xc2, 0x63, 0x49, 0xc4, 0x0d, 0x8e, 0xf8, 0x90,
	0xc6, 0x4e, 0xd9, 0x24, 0xe4, 0x2d, 0x16, 0x44, 0x67, 0xba, 0xb5, 0xab, 0x1b, 0xd1, 0x15, 0x54,
	0x18, 0x65, 0x5c, 0x2e, 0xdc, 0xdc, 0xfb, 0xc6, 0xcb, 0xc7, 0x1f, 0xf0, 0xd2, 0xa1, 0xcc, 0x15,
	0xca, 0x98, 0x40, 0x29, 0x21, 0xb7, 0xa1, 0x7d, 0xf8, 0x34, 0xcf, 0xa5, 0x12, 0x2b, 0x22, 0xc7,
	0x54, 0x19, 0x37, 0x38, 0x22, 0x8c, 0x62, 0x73, 0xda, 0x9c, 0xd2, 0x81, 0x75, 0xb8, 0xee, 0xd5,
	0x73, 0x10, 0x2a, 0x07, 0xa6, 0xb3, 0x36, 0xd9, 0x25, 0x8c, 0x9a, 0x67, 0x03, 0x7d, 0xbf, 0xb4,
	0x7c, 0x2a, 0x60, 0x54, 0xbf, 0x02, 0x2c, 0x76, 0xaa, 0x07, 0xd6, 0x61, 0xd1, 0xdb, 0xcd, 0x2d,
	0xc8, 0x20, 0x0d, 0x77, 0x62, 0xf4, 0x0a, 0xb6, 0xcc, 0x21, 0x10, 0x93, 0xd7, 0xcc, 0x59, 0x3b,
	0xb0, 0x0e, 0x37, 0x8e, 0x9f, 0x65, 0x67, 0x78, 0xc7, 0x6b, 0xe7, 0x6d, 0x92, 0x4c, 0x23, 0xfa,
	0x16, 0x9c, 0x11, 0x09, 0x42, 0x3c, 0x4f, 0x1f, 0x54, 0x62, 0x2a, 0x25, 0x97, 0xb1, 0x53, 0x33,
	0xfb, 0x5e, 0xd3, 0x71, 0x37, 0x9a, 0x39, 0x68, 0x99, 0xa0, 0x4e, 0x3c, 0x77, 0x5d, 0xd1, 0xd8,
	0xb1, 0x8d, 0x8d, 0x7a, 0xd6, 0xc6, 0xdd, 0x4f, 0xa4, 0x57, 0x59, 0xbe, 0x9b, 0x31, 0x7a, 0x0a,
	0x30, 0x33, 0x12, 0x3b, 0x43, 0xf3, 0x70, 0x66, 0x5a, 0xd0, 0x4f, 0x50, 0xf6, 0x93, 0x58, 0x71,
	0x86, 0x33, 0xdd, 0xfe, 0xd0, 0x0f, 0xec, 0xc6, 0xf1, 0xd1, 0x07, 0x76, 0xf7, 0xc4, 0x88, 0xcc,
	0x39, 0x99, 0x4d, 0xc5, 0xb3, 0x53, 0x94, 0x3b, 0xc7, 0x5f, 0xc3, 0xb3, 0x20, 0xd2, 0x05, 0xca,
	0x7c, 0x35, 0x16, 0x6a, 0x09, 0xe7, 0xb7, 0x82, 0x99, 0x61, 0xb6, 0xfa, 0x68, 0x5c, 0x4d, 0xfa,
	0xa4, 0x83, 0x79, 0x4f, 0x0c, 0x63, 0x86, 0xcd, 0x47, 0xd1, 0x08, 0x9e, 0x8b, 0xc5, 0x7c, 0xb3,
	0x34, 0xca, 0xef, 0xff, 0x3d, 0xca, 0x53, 0xb1, 0x90, 0x74, 0xf2, 0xf1, 0xfd, 0xbf, 0x2c, 0xa8,
	0xde, 0x35, 0x6d, 0x84, 0xe0, 0xbe, 0x3e, 0xad, 0xa6, 0xbe, 0x58, 0xf7, 0xcc, 0x37, 0x1a, 0xc1,
	0x96, 0x3e, 0x49, 0x8c, 0xea, 0xc9, 0x33, 0x22, 0x9c, 0x82, 0x59, 0xd2, 0xe6, 0xff, 0x5c, 0xd2,
	0x46, 0x6f, 0x0a, 0xe9, 0x10, 0xd1, 0x8a, 0x94, 0xbc, 0xf5, 0x36, 0x45, 0xa6, 0x69, 0xff, 0x0a,
	0xca, 0x4b, 0x5d, 0x90, 0x0d, 0xc5, 0xb7, 0xf4, 0x76, 0xe2, 0x47, 0x7f, 0xa2, 0xcf, 0x61, 0xcd,
	0x54, 0x7a, 0x4e, 0xba, 0x0c, 0xb5, 0xdc, 0xa9, 0x56, 0x4a, 0x5e, 0xe9, 0xa0, 0x97, 0xf6, 0xf9,
	0xa1, 0xf0, 0x9d, 0x55, 0xff, 0x06, 0x4a, 0x69, 0x72, 0x43, 0x1b, 0xf0, 0xe0, 0x55, 0xeb, 0xb4,
	0x79, 0x79, 0x31, 0xb0, 0xef, 0xa1, 0x12, 0x14, 0xdc, 0xae, 0x6d, 0xa1, 0x07, 0x50, 0x74, 0x4f,
	0x4f, 0xed, 0x02, 0xda, 0x06, 0x68, 0x9e, 0x9d, 0x79, 0xad, 0x7e, 0xbf, 0x7d, 0xd5, 0xb2, 0x8b,
	0xf5, 0x17, 0x50, 0x5e, 0x4a, 0x42, 0xa8, 0x06, 0xe5, 0x09, 0x02, 0x77, 0x2f, 0x3b, 0xb8, 0x3d,
	0x68, 0x79, 0x7d, 0xfb, 0x9e, 0x81, 0x74, 0x5b, 0x29, 0x6d, 0xf0, 0xda, 0xb5, 0x0b, 0xf5, 0x5f,
	0x2d, 0x80, 0x79, 0xde, 0x40, 0x15, 0xd8, 0x99, 0xea, 0x3a, 0xad, 0x0e, 0x76, 0x7b, 0xda, 0xc2,
	0x36, 0x40, 0xd7, 0x9d, 0xfd, 0x5b, 0x08, 0xa0, 0xd4, 0x69, 0x76, 0x2f, 0x9b, 0x17, 0x76, 0x01,
	0xed, 0x41, 0xa5, 0xff, 0xba, 0xd9, 0xeb, 0xb5, 0xbb, 0x67, 0xf8, 0xbc, 0x75, 0xe9, 0xb5, 0xfb,
	0x83, 0xf6, 0x49, 0xdf, 0xbe, 0x8f, 0x9e, 0x80, 0xe3, 0xb5, 0x4e, 0xdc, 0x4e, 0xef, 0x72, 0xd0,
	0x1c, 0xb4, 0xdd, 0x6e, 0x36, 0xba, 0x86, 0x1e, 0x41, 0xad, 0x7f, 0x72, 0xde, 0x7a, 0x75, 0x79,
	0xb1, 0x20, 0x2c, 0xe9, 0xd1, 0x32, 0xff, 0xc5, 0x97, 0x31, 0x38, 0x5c, 0x8e, 0xb3, 0xcb, 0x37,
	0x2b, 0xa8, 0x5f, 0x56, 0xf3, 0x1b, 0xda, 0xd3, 0x45, 0x72, 0xdc, 0xb3, 0xde, 0xbc, 0x18, 0x07,
	0xea, 0x26, 0xb9, 0x6e, 0xf8, 0x9c, 0x1d, 0x65, 0x4a, 0xec, 0xbb, 0x3f, 0xc7, 0x3c, 0x5f, 0x7b,
	0xff, 0x6d, 0x59, 0xd7, 0x25, 0xf3, 0xf3, 0xe5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39, 0x03,
	0x23, 0x91, 0x09, 0x0c, 0x00, 0x00,
}
