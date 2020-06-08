// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/launch_plan.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
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

// By default any launch plan regardless of state can be used to launch a workflow execution.
// However, at most one version of a launch plan
// (e.g. a NamedEntityIdentifier set of shared project, domain and name values) can be
// active at a time in regards to *schedules*. That is, at most one schedule in a NamedEntityIdentifier
// group will be observed and trigger executions at a defined cadence.
type LaunchPlanState int32

const (
	LaunchPlanState_INACTIVE LaunchPlanState = 0
	LaunchPlanState_ACTIVE   LaunchPlanState = 1
)

var LaunchPlanState_name = map[int32]string{
	0: "INACTIVE",
	1: "ACTIVE",
}

var LaunchPlanState_value = map[string]int32{
	"INACTIVE": 0,
	"ACTIVE":   1,
}

func (x LaunchPlanState) String() string {
	return proto.EnumName(LaunchPlanState_name, int32(x))
}

func (LaunchPlanState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{0}
}

// Request to register a launch plan. A LaunchPlanSpec may include a complete or incomplete set of inputs required
// to launch a workflow execution. By default all launch plans are registered in state INACTIVE. If you wish to
// set the state to ACTIVE, you must submit a LaunchPlanUpdateRequest, after you have created a launch plan.
type LaunchPlanCreateRequest struct {
	// Uniquely identifies a launch plan entity.
	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// User-provided launch plan details, including reference workflow, inputs and other metadata.
	Spec                 *LaunchPlanSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LaunchPlanCreateRequest) Reset()         { *m = LaunchPlanCreateRequest{} }
func (m *LaunchPlanCreateRequest) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanCreateRequest) ProtoMessage()    {}
func (*LaunchPlanCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{0}
}

func (m *LaunchPlanCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanCreateRequest.Unmarshal(m, b)
}
func (m *LaunchPlanCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanCreateRequest.Marshal(b, m, deterministic)
}
func (m *LaunchPlanCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanCreateRequest.Merge(m, src)
}
func (m *LaunchPlanCreateRequest) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanCreateRequest.Size(m)
}
func (m *LaunchPlanCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanCreateRequest proto.InternalMessageInfo

func (m *LaunchPlanCreateRequest) GetId() *core.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *LaunchPlanCreateRequest) GetSpec() *LaunchPlanSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type LaunchPlanCreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaunchPlanCreateResponse) Reset()         { *m = LaunchPlanCreateResponse{} }
func (m *LaunchPlanCreateResponse) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanCreateResponse) ProtoMessage()    {}
func (*LaunchPlanCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{1}
}

func (m *LaunchPlanCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanCreateResponse.Unmarshal(m, b)
}
func (m *LaunchPlanCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanCreateResponse.Marshal(b, m, deterministic)
}
func (m *LaunchPlanCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanCreateResponse.Merge(m, src)
}
func (m *LaunchPlanCreateResponse) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanCreateResponse.Size(m)
}
func (m *LaunchPlanCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanCreateResponse proto.InternalMessageInfo

// A LaunchPlan provides the capability to templatize workflow executions.
// Launch plans simplify associating one or more schedules, inputs and notifications with your workflows.
// Launch plans can be shared and used to trigger executions with predefined inputs even when a workflow
// definition doesn't necessarily have a default value for said input.
type LaunchPlan struct {
	Id                   *core.Identifier   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec                 *LaunchPlanSpec    `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Closure              *LaunchPlanClosure `protobuf:"bytes,3,opt,name=closure,proto3" json:"closure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LaunchPlan) Reset()         { *m = LaunchPlan{} }
func (m *LaunchPlan) String() string { return proto.CompactTextString(m) }
func (*LaunchPlan) ProtoMessage()    {}
func (*LaunchPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{2}
}

func (m *LaunchPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlan.Unmarshal(m, b)
}
func (m *LaunchPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlan.Marshal(b, m, deterministic)
}
func (m *LaunchPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlan.Merge(m, src)
}
func (m *LaunchPlan) XXX_Size() int {
	return xxx_messageInfo_LaunchPlan.Size(m)
}
func (m *LaunchPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlan.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlan proto.InternalMessageInfo

func (m *LaunchPlan) GetId() *core.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *LaunchPlan) GetSpec() *LaunchPlanSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *LaunchPlan) GetClosure() *LaunchPlanClosure {
	if m != nil {
		return m.Closure
	}
	return nil
}

// Response object for list launch plan requests.
type LaunchPlanList struct {
	LaunchPlans []*LaunchPlan `protobuf:"bytes,1,rep,name=launch_plans,json=launchPlans,proto3" json:"launch_plans,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaunchPlanList) Reset()         { *m = LaunchPlanList{} }
func (m *LaunchPlanList) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanList) ProtoMessage()    {}
func (*LaunchPlanList) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{3}
}

func (m *LaunchPlanList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanList.Unmarshal(m, b)
}
func (m *LaunchPlanList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanList.Marshal(b, m, deterministic)
}
func (m *LaunchPlanList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanList.Merge(m, src)
}
func (m *LaunchPlanList) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanList.Size(m)
}
func (m *LaunchPlanList) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanList.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanList proto.InternalMessageInfo

func (m *LaunchPlanList) GetLaunchPlans() []*LaunchPlan {
	if m != nil {
		return m.LaunchPlans
	}
	return nil
}

func (m *LaunchPlanList) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Defines permissions associated with executions created by this launch plan spec.
type Auth struct {
	// Types that are valid to be assigned to Method:
	//	*Auth_AssumableIamRole
	//	*Auth_KubernetesServiceAccount
	Method               isAuth_Method `protobuf_oneof:"method"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{4}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

type isAuth_Method interface {
	isAuth_Method()
}

type Auth_AssumableIamRole struct {
	AssumableIamRole string `protobuf:"bytes,1,opt,name=assumable_iam_role,json=assumableIamRole,proto3,oneof"`
}

type Auth_KubernetesServiceAccount struct {
	KubernetesServiceAccount string `protobuf:"bytes,2,opt,name=kubernetes_service_account,json=kubernetesServiceAccount,proto3,oneof"`
}

func (*Auth_AssumableIamRole) isAuth_Method() {}

func (*Auth_KubernetesServiceAccount) isAuth_Method() {}

func (m *Auth) GetMethod() isAuth_Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *Auth) GetAssumableIamRole() string {
	if x, ok := m.GetMethod().(*Auth_AssumableIamRole); ok {
		return x.AssumableIamRole
	}
	return ""
}

func (m *Auth) GetKubernetesServiceAccount() string {
	if x, ok := m.GetMethod().(*Auth_KubernetesServiceAccount); ok {
		return x.KubernetesServiceAccount
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Auth) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Auth_AssumableIamRole)(nil),
		(*Auth_KubernetesServiceAccount)(nil),
	}
}

// User-provided launch plan definition and configuration values.
type LaunchPlanSpec struct {
	// Reference to the Workflow template that the launch plan references
	WorkflowId *core.Identifier `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	// Metadata for the Launch Plan
	EntityMetadata *LaunchPlanMetadata `protobuf:"bytes,2,opt,name=entity_metadata,json=entityMetadata,proto3" json:"entity_metadata,omitempty"`
	// Input values to be passed for the execution
	DefaultInputs *core.ParameterMap `protobuf:"bytes,3,opt,name=default_inputs,json=defaultInputs,proto3" json:"default_inputs,omitempty"`
	// Fixed, non-overridable inputs for the Launch Plan
	FixedInputs *core.LiteralMap `protobuf:"bytes,4,opt,name=fixed_inputs,json=fixedInputs,proto3" json:"fixed_inputs,omitempty"`
	// String to indicate the role to use to execute the workflow underneath
	Role string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"` // Deprecated: Do not use.
	// Custom labels to be applied to the execution resource.
	Labels *Labels `protobuf:"bytes,6,opt,name=labels,proto3" json:"labels,omitempty"`
	// Custom annotations to be applied to the execution resource.
	Annotations *Annotations `protobuf:"bytes,7,opt,name=annotations,proto3" json:"annotations,omitempty"`
	// Indicates the permission associated with workflow executions triggered with this launch plan.
	Auth                 *Auth     `protobuf:"bytes,8,opt,name=auth,proto3" json:"auth,omitempty"` // Deprecated: Do not use.
	AuthRole             *AuthRole `protobuf:"bytes,9,opt,name=auth_role,json=authRole,proto3" json:"auth_role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LaunchPlanSpec) Reset()         { *m = LaunchPlanSpec{} }
func (m *LaunchPlanSpec) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanSpec) ProtoMessage()    {}
func (*LaunchPlanSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{5}
}

func (m *LaunchPlanSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanSpec.Unmarshal(m, b)
}
func (m *LaunchPlanSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanSpec.Marshal(b, m, deterministic)
}
func (m *LaunchPlanSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanSpec.Merge(m, src)
}
func (m *LaunchPlanSpec) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanSpec.Size(m)
}
func (m *LaunchPlanSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanSpec.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanSpec proto.InternalMessageInfo

func (m *LaunchPlanSpec) GetWorkflowId() *core.Identifier {
	if m != nil {
		return m.WorkflowId
	}
	return nil
}

func (m *LaunchPlanSpec) GetEntityMetadata() *LaunchPlanMetadata {
	if m != nil {
		return m.EntityMetadata
	}
	return nil
}

func (m *LaunchPlanSpec) GetDefaultInputs() *core.ParameterMap {
	if m != nil {
		return m.DefaultInputs
	}
	return nil
}

func (m *LaunchPlanSpec) GetFixedInputs() *core.LiteralMap {
	if m != nil {
		return m.FixedInputs
	}
	return nil
}

// Deprecated: Do not use.
func (m *LaunchPlanSpec) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *LaunchPlanSpec) GetLabels() *Labels {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LaunchPlanSpec) GetAnnotations() *Annotations {
	if m != nil {
		return m.Annotations
	}
	return nil
}

// Deprecated: Do not use.
func (m *LaunchPlanSpec) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *LaunchPlanSpec) GetAuthRole() *AuthRole {
	if m != nil {
		return m.AuthRole
	}
	return nil
}

// Values computed by the flyte platform after launch plan registration.
// These include expected_inputs required to be present in a CreateExecutionRequest
// to launch the reference workflow as well timestamp values associated with the launch plan.
type LaunchPlanClosure struct {
	// Indicate the Launch plan phase
	State LaunchPlanState `protobuf:"varint,1,opt,name=state,proto3,enum=flyteidl.admin.LaunchPlanState" json:"state,omitempty"`
	// Indicates the set of inputs to execute the Launch plan
	ExpectedInputs *core.ParameterMap `protobuf:"bytes,2,opt,name=expected_inputs,json=expectedInputs,proto3" json:"expected_inputs,omitempty"`
	// Indicates the set of outputs from the Launch plan
	ExpectedOutputs *core.VariableMap `protobuf:"bytes,3,opt,name=expected_outputs,json=expectedOutputs,proto3" json:"expected_outputs,omitempty"`
	// Time at which the launch plan was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Time at which the launch plan was last updated.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LaunchPlanClosure) Reset()         { *m = LaunchPlanClosure{} }
func (m *LaunchPlanClosure) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanClosure) ProtoMessage()    {}
func (*LaunchPlanClosure) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{6}
}

func (m *LaunchPlanClosure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanClosure.Unmarshal(m, b)
}
func (m *LaunchPlanClosure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanClosure.Marshal(b, m, deterministic)
}
func (m *LaunchPlanClosure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanClosure.Merge(m, src)
}
func (m *LaunchPlanClosure) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanClosure.Size(m)
}
func (m *LaunchPlanClosure) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanClosure.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanClosure proto.InternalMessageInfo

func (m *LaunchPlanClosure) GetState() LaunchPlanState {
	if m != nil {
		return m.State
	}
	return LaunchPlanState_INACTIVE
}

func (m *LaunchPlanClosure) GetExpectedInputs() *core.ParameterMap {
	if m != nil {
		return m.ExpectedInputs
	}
	return nil
}

func (m *LaunchPlanClosure) GetExpectedOutputs() *core.VariableMap {
	if m != nil {
		return m.ExpectedOutputs
	}
	return nil
}

func (m *LaunchPlanClosure) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *LaunchPlanClosure) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// Additional launch plan attributes included in the LaunchPlanSpec not strictly required to launch
// the reference workflow.
type LaunchPlanMetadata struct {
	// Schedule to execute the Launch Plan
	Schedule *Schedule `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
	// List of notifications based on Execution status transitions
	Notifications        []*Notification `protobuf:"bytes,2,rep,name=notifications,proto3" json:"notifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LaunchPlanMetadata) Reset()         { *m = LaunchPlanMetadata{} }
func (m *LaunchPlanMetadata) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanMetadata) ProtoMessage()    {}
func (*LaunchPlanMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{7}
}

func (m *LaunchPlanMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanMetadata.Unmarshal(m, b)
}
func (m *LaunchPlanMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanMetadata.Marshal(b, m, deterministic)
}
func (m *LaunchPlanMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanMetadata.Merge(m, src)
}
func (m *LaunchPlanMetadata) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanMetadata.Size(m)
}
func (m *LaunchPlanMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanMetadata proto.InternalMessageInfo

func (m *LaunchPlanMetadata) GetSchedule() *Schedule {
	if m != nil {
		return m.Schedule
	}
	return nil
}

func (m *LaunchPlanMetadata) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

// Request to set the referenced launch plan state to the configured value.
type LaunchPlanUpdateRequest struct {
	// Identifier of launch plan for which to change state.
	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Desired state to apply to the launch plan.
	State                LaunchPlanState `protobuf:"varint,2,opt,name=state,proto3,enum=flyteidl.admin.LaunchPlanState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LaunchPlanUpdateRequest) Reset()         { *m = LaunchPlanUpdateRequest{} }
func (m *LaunchPlanUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanUpdateRequest) ProtoMessage()    {}
func (*LaunchPlanUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{8}
}

func (m *LaunchPlanUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanUpdateRequest.Unmarshal(m, b)
}
func (m *LaunchPlanUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanUpdateRequest.Marshal(b, m, deterministic)
}
func (m *LaunchPlanUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanUpdateRequest.Merge(m, src)
}
func (m *LaunchPlanUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanUpdateRequest.Size(m)
}
func (m *LaunchPlanUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanUpdateRequest proto.InternalMessageInfo

func (m *LaunchPlanUpdateRequest) GetId() *core.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *LaunchPlanUpdateRequest) GetState() LaunchPlanState {
	if m != nil {
		return m.State
	}
	return LaunchPlanState_INACTIVE
}

// Purposefully empty, may be populated in the future.
type LaunchPlanUpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaunchPlanUpdateResponse) Reset()         { *m = LaunchPlanUpdateResponse{} }
func (m *LaunchPlanUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*LaunchPlanUpdateResponse) ProtoMessage()    {}
func (*LaunchPlanUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{9}
}

func (m *LaunchPlanUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchPlanUpdateResponse.Unmarshal(m, b)
}
func (m *LaunchPlanUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchPlanUpdateResponse.Marshal(b, m, deterministic)
}
func (m *LaunchPlanUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchPlanUpdateResponse.Merge(m, src)
}
func (m *LaunchPlanUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_LaunchPlanUpdateResponse.Size(m)
}
func (m *LaunchPlanUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchPlanUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchPlanUpdateResponse proto.InternalMessageInfo

// Represents a request struct for finding an active launch plan for a given NamedEntityIdentifier
type ActiveLaunchPlanRequest struct {
	Id                   *NamedEntityIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ActiveLaunchPlanRequest) Reset()         { *m = ActiveLaunchPlanRequest{} }
func (m *ActiveLaunchPlanRequest) String() string { return proto.CompactTextString(m) }
func (*ActiveLaunchPlanRequest) ProtoMessage()    {}
func (*ActiveLaunchPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{10}
}

func (m *ActiveLaunchPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActiveLaunchPlanRequest.Unmarshal(m, b)
}
func (m *ActiveLaunchPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActiveLaunchPlanRequest.Marshal(b, m, deterministic)
}
func (m *ActiveLaunchPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActiveLaunchPlanRequest.Merge(m, src)
}
func (m *ActiveLaunchPlanRequest) XXX_Size() int {
	return xxx_messageInfo_ActiveLaunchPlanRequest.Size(m)
}
func (m *ActiveLaunchPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActiveLaunchPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActiveLaunchPlanRequest proto.InternalMessageInfo

func (m *ActiveLaunchPlanRequest) GetId() *NamedEntityIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Represents a request structure to list active launch plans within a project/domain.
type ActiveLaunchPlanListRequest struct {
	// Name of the project that contains the identifiers.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the domain the identifiers belongs to within the project.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Indicates the number of resources to be returned.
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query.
	// +optional
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	// Sort ordering.
	// +optional
	SortBy               *Sort    `protobuf:"bytes,5,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActiveLaunchPlanListRequest) Reset()         { *m = ActiveLaunchPlanListRequest{} }
func (m *ActiveLaunchPlanListRequest) String() string { return proto.CompactTextString(m) }
func (*ActiveLaunchPlanListRequest) ProtoMessage()    {}
func (*ActiveLaunchPlanListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_368a863574f5e699, []int{11}
}

func (m *ActiveLaunchPlanListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActiveLaunchPlanListRequest.Unmarshal(m, b)
}
func (m *ActiveLaunchPlanListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActiveLaunchPlanListRequest.Marshal(b, m, deterministic)
}
func (m *ActiveLaunchPlanListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActiveLaunchPlanListRequest.Merge(m, src)
}
func (m *ActiveLaunchPlanListRequest) XXX_Size() int {
	return xxx_messageInfo_ActiveLaunchPlanListRequest.Size(m)
}
func (m *ActiveLaunchPlanListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActiveLaunchPlanListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActiveLaunchPlanListRequest proto.InternalMessageInfo

func (m *ActiveLaunchPlanListRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ActiveLaunchPlanListRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ActiveLaunchPlanListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ActiveLaunchPlanListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ActiveLaunchPlanListRequest) GetSortBy() *Sort {
	if m != nil {
		return m.SortBy
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.admin.LaunchPlanState", LaunchPlanState_name, LaunchPlanState_value)
	proto.RegisterType((*LaunchPlanCreateRequest)(nil), "flyteidl.admin.LaunchPlanCreateRequest")
	proto.RegisterType((*LaunchPlanCreateResponse)(nil), "flyteidl.admin.LaunchPlanCreateResponse")
	proto.RegisterType((*LaunchPlan)(nil), "flyteidl.admin.LaunchPlan")
	proto.RegisterType((*LaunchPlanList)(nil), "flyteidl.admin.LaunchPlanList")
	proto.RegisterType((*Auth)(nil), "flyteidl.admin.Auth")
	proto.RegisterType((*LaunchPlanSpec)(nil), "flyteidl.admin.LaunchPlanSpec")
	proto.RegisterType((*LaunchPlanClosure)(nil), "flyteidl.admin.LaunchPlanClosure")
	proto.RegisterType((*LaunchPlanMetadata)(nil), "flyteidl.admin.LaunchPlanMetadata")
	proto.RegisterType((*LaunchPlanUpdateRequest)(nil), "flyteidl.admin.LaunchPlanUpdateRequest")
	proto.RegisterType((*LaunchPlanUpdateResponse)(nil), "flyteidl.admin.LaunchPlanUpdateResponse")
	proto.RegisterType((*ActiveLaunchPlanRequest)(nil), "flyteidl.admin.ActiveLaunchPlanRequest")
	proto.RegisterType((*ActiveLaunchPlanListRequest)(nil), "flyteidl.admin.ActiveLaunchPlanListRequest")
}

func init() { proto.RegisterFile("flyteidl/admin/launch_plan.proto", fileDescriptor_368a863574f5e699) }

var fileDescriptor_368a863574f5e699 = []byte{
	// 943 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x8e, 0x14, 0x59, 0x96, 0x46, 0xb6, 0xec, 0x2e, 0x02, 0x87, 0x95, 0xd3, 0xc4, 0x25, 0x50,
	0x20, 0xfd, 0x09, 0x85, 0x3a, 0xf5, 0xa1, 0x3f, 0x29, 0x20, 0xb9, 0x06, 0x22, 0xd4, 0x49, 0x8d,
	0x75, 0x9a, 0x43, 0x2f, 0xc4, 0x8a, 0x1c, 0x49, 0x5b, 0x2f, 0xb9, 0x2c, 0x77, 0x99, 0x58, 0xe8,
	0xa5, 0xb7, 0xde, 0xfa, 0x1c, 0x45, 0xdf, 0xa5, 0xef, 0x54, 0x70, 0xb9, 0xa4, 0xfe, 0xe2, 0xa6,
	0x2d, 0xd0, 0x93, 0x39, 0x9e, 0xef, 0x9b, 0x9d, 0x99, 0x9d, 0xf9, 0x56, 0x70, 0x34, 0x11, 0x73,
	0x8d, 0x3c, 0x14, 0x7d, 0x16, 0x46, 0x3c, 0xee, 0x0b, 0x96, 0xc5, 0xc1, 0xcc, 0x4f, 0x04, 0x8b,
	0xbd, 0x24, 0x95, 0x5a, 0x92, 0x6e, 0x89, 0xf0, 0x0c, 0xa2, 0x77, 0xaf, 0x62, 0x04, 0x32, 0xc5,
	0xbe, 0xe0, 0x1a, 0x53, 0x26, 0x54, 0x81, 0xee, 0xdd, 0x5f, 0xf5, 0xf2, 0x10, 0x63, 0xcd, 0x27,
	0x1c, 0x53, 0xeb, 0x7f, 0x6f, 0xcd, 0x1f, 0x6b, 0x4c, 0x27, 0x2c, 0xc0, 0x0d, 0x77, 0x91, 0x8e,
	0x0a, 0x66, 0x18, 0x66, 0xa2, 0x74, 0x1f, 0xae, 0xb9, 0x03, 0x19, 0x45, 0xd2, 0x26, 0xda, 0x7b,
	0x30, 0x95, 0x72, 0x2a, 0xb0, 0x6f, 0xac, 0x71, 0x36, 0xe9, 0x6b, 0x1e, 0xa1, 0xd2, 0x2c, 0x4a,
	0x0a, 0x80, 0x7b, 0x0d, 0x77, 0xcf, 0x4d, 0x79, 0x17, 0x82, 0xc5, 0xa7, 0x29, 0x32, 0x8d, 0x14,
	0x7f, 0xca, 0x50, 0x69, 0xf2, 0x21, 0xd4, 0x79, 0xe8, 0xd4, 0x8e, 0x6a, 0x0f, 0x3b, 0xc7, 0xef,
	0x7a, 0x55, 0xc5, 0x79, 0x8e, 0xde, 0xa8, 0xaa, 0x81, 0xd6, 0x79, 0x48, 0x8e, 0xa1, 0xa1, 0x12,
	0x0c, 0x9c, 0xba, 0x01, 0xdf, 0xf7, 0x56, 0xdb, 0xe3, 0x2d, 0x4e, 0xb8, 0x4c, 0x30, 0xa0, 0x06,
	0xeb, 0xf6, 0xc0, 0xd9, 0x3c, 0x59, 0x25, 0x32, 0x56, 0xe8, 0xfe, 0x5e, 0x03, 0x58, 0x38, 0xff,
	0xe7, 0x4c, 0xc8, 0x97, 0xb0, 0x1d, 0x08, 0xa9, 0xb2, 0x14, 0x9d, 0xdb, 0x86, 0xf6, 0xfe, 0xcd,
	0xb4, 0xd3, 0x02, 0x48, 0x4b, 0x86, 0x8b, 0xd0, 0x5d, 0x78, 0xcf, 0xb9, 0xd2, 0xe4, 0x09, 0xec,
	0x2c, 0x4d, 0x8c, 0x72, 0x6a, 0x47, 0xb7, 0x1f, 0x76, 0x8e, 0x7b, 0x37, 0xc7, 0xa4, 0x1d, 0x51,
	0x7d, 0x2b, 0x72, 0x07, 0xb6, 0xb4, 0xbc, 0xc2, 0xd8, 0x94, 0xd0, 0xa6, 0x85, 0xe1, 0xfe, 0x52,
	0x83, 0xc6, 0x20, 0xd3, 0x33, 0xe2, 0x01, 0x61, 0x4a, 0x65, 0x11, 0x1b, 0x0b, 0xf4, 0x39, 0x8b,
	0xfc, 0x54, 0x0a, 0x34, 0xbd, 0x69, 0x3f, 0xbd, 0x45, 0xf7, 0x2b, 0xdf, 0x88, 0x45, 0x54, 0x0a,
	0x24, 0x5f, 0x43, 0xef, 0x2a, 0x1b, 0x63, 0x1a, 0xa3, 0x46, 0xe5, 0x2b, 0x4c, 0x5f, 0xf1, 0x00,
	0x7d, 0x16, 0x04, 0x32, 0x8b, 0x75, 0x71, 0xc6, 0xd3, 0x5b, 0xd4, 0x59, 0x60, 0x2e, 0x0b, 0xc8,
	0xa0, 0x40, 0x0c, 0x5b, 0xd0, 0x8c, 0x50, 0xcf, 0x64, 0xe8, 0xfe, 0xda, 0x58, 0x2e, 0x35, 0xef,
	0x1f, 0xf9, 0x02, 0x3a, 0xaf, 0x65, 0x7a, 0x35, 0x11, 0xf2, 0xb5, 0xff, 0x4f, 0x6e, 0x08, 0x4a,
	0xf4, 0x28, 0x24, 0xdf, 0xc2, 0x5e, 0xfe, 0x7f, 0x3d, 0xf7, 0x23, 0xd4, 0x2c, 0x64, 0x9a, 0xd9,
	0x4b, 0x73, 0x6f, 0xee, 0xd4, 0x33, 0x8b, 0xa4, 0xdd, 0x82, 0x5a, 0xda, 0x64, 0x08, 0xdd, 0x10,
	0x27, 0x2c, 0x13, 0xda, 0xe7, 0x71, 0x92, 0x69, 0x65, 0x6f, 0xf2, 0x70, 0x2d, 0x97, 0x0b, 0x96,
	0xb2, 0x08, 0x35, 0xa6, 0xcf, 0x58, 0x42, 0x77, 0x2d, 0x65, 0x64, 0x18, 0xe4, 0x2b, 0xd8, 0x99,
	0xf0, 0x6b, 0x0c, 0xcb, 0x08, 0x8d, 0x37, 0x56, 0x73, 0x5e, 0xec, 0x76, 0xce, 0xef, 0x18, 0xb8,
	0x65, 0x1f, 0x40, 0xc3, 0xdc, 0xc4, 0x56, 0xde, 0xd1, 0x61, 0xdd, 0xa9, 0x51, 0x63, 0x13, 0x0f,
	0x9a, 0x82, 0x8d, 0x51, 0x28, 0xa7, 0x69, 0xe2, 0x1d, 0x6c, 0x56, 0x97, 0x7b, 0xa9, 0x45, 0x91,
	0x27, 0xd0, 0x61, 0x71, 0x2c, 0x35, 0xd3, 0x5c, 0xc6, 0xca, 0xd9, 0x5e, 0x2f, 0xa3, 0x20, 0x0d,
	0x16, 0x10, 0xba, 0x8c, 0x27, 0x9f, 0x40, 0x83, 0x65, 0x7a, 0xe6, 0xb4, 0x0c, 0xef, 0xce, 0x06,
	0x2f, 0xd3, 0xb3, 0x22, 0xb9, 0x1c, 0x45, 0x4e, 0xa0, 0x9d, 0xff, 0x2d, 0x66, 0xa8, 0x6d, 0x28,
	0xce, 0x9b, 0x28, 0xf9, 0x24, 0xd1, 0x16, 0xb3, 0x5f, 0xee, 0x9f, 0x75, 0x78, 0x67, 0x63, 0x25,
	0xc8, 0x09, 0x6c, 0x29, 0xcd, 0x74, 0x31, 0x8c, 0xdd, 0xe3, 0x07, 0x7f, 0xb3, 0x7b, 0x39, 0x8c,
	0x16, 0x68, 0xf2, 0x0d, 0xec, 0xe1, 0x75, 0x82, 0x81, 0x5e, 0x74, 0xbe, 0xfe, 0xf6, 0xbb, 0xeb,
	0x96, 0x1c, 0xdb, 0xfe, 0x33, 0xd8, 0xaf, 0xa2, 0xc8, 0x4c, 0x2f, 0x8d, 0x40, 0x6f, 0x2d, 0xcc,
	0x4b, 0x96, 0xf2, 0x7c, 0x41, 0xf2, 0x28, 0xd5, 0xc9, 0xdf, 0x15, 0x14, 0xf2, 0x39, 0x40, 0x60,
	0xa4, 0x28, 0xf4, 0x99, 0xb6, 0x13, 0xd0, 0xf3, 0x0a, 0x11, 0xf5, 0x4a, 0x11, 0xf5, 0x5e, 0x94,
	0x22, 0x4a, 0xdb, 0x16, 0x3d, 0xd0, 0x39, 0x35, 0x4b, 0xc2, 0x92, 0xba, 0xf5, 0x76, 0xaa, 0x45,
	0x0f, 0xb4, 0xfb, 0x5b, 0x0d, 0xc8, 0xe6, 0x90, 0x93, 0xcf, 0xa0, 0x55, 0x6a, 0xbd, 0x5d, 0xad,
	0x8d, 0xcb, 0xb9, 0xb4, 0x7e, 0x5a, 0x21, 0xc9, 0x10, 0x76, 0x63, 0x99, 0xef, 0x5b, 0x60, 0x47,
	0xa8, 0x6e, 0xf4, 0xe7, 0xde, 0x3a, 0xf5, 0xf9, 0x12, 0x88, 0xae, 0x52, 0xdc, 0x9f, 0x97, 0x5f,
	0x85, 0xef, 0x4d, 0x9e, 0xff, 0xe1, 0x55, 0xa8, 0x06, 0xa2, 0xfe, 0x6f, 0x06, 0x62, 0xf5, 0x61,
	0x28, 0x0f, 0xb7, 0x0f, 0xc3, 0x05, 0xdc, 0x1d, 0x04, 0x9a, 0xbf, 0xc2, 0x25, 0xf5, 0xb4, 0x89,
	0x9d, 0x2c, 0x25, 0xf6, 0xc1, 0x46, 0xb1, 0x2c, 0xc2, 0xf0, 0xcc, 0x68, 0xc6, 0x6a, 0x92, 0xee,
	0x1f, 0x35, 0x38, 0x5c, 0x0f, 0x99, 0xcb, 0x78, 0x19, 0xd6, 0x81, 0xed, 0x24, 0x95, 0x3f, 0x62,
	0xa0, 0x0b, 0x91, 0xa5, 0xa5, 0x49, 0x0e, 0xa0, 0x19, 0xca, 0x88, 0xf1, 0x52, 0xa9, 0xad, 0x95,
	0x0b, 0xb8, 0xe0, 0x11, 0xd7, 0x66, 0xfe, 0x76, 0x69, 0x61, 0x2c, 0x64, 0xbd, 0xb1, 0x24, 0xeb,
	0xe4, 0x11, 0x6c, 0x2b, 0x99, 0x6a, 0x7f, 0x3c, 0xb7, 0x13, 0xb3, 0xb1, 0xb1, 0x97, 0x32, 0xd5,
	0xb4, 0x99, 0x83, 0x86, 0xf3, 0x8f, 0x3e, 0x86, 0xbd, 0xb5, 0xa6, 0x91, 0x1d, 0x68, 0x8d, 0x9e,
	0x0f, 0x4e, 0x5f, 0x8c, 0x5e, 0x9e, 0xed, 0xdf, 0x22, 0x00, 0x4d, 0xfb, 0x5d, 0x1b, 0x3e, 0xfe,
	0xe1, 0xd3, 0x29, 0xd7, 0xb3, 0x6c, 0xec, 0x05, 0x32, 0xea, 0x8b, 0xf9, 0x44, 0xf7, 0xab, 0x9f,
	0x0a, 0x53, 0x8c, 0xfb, 0xc9, 0xf8, 0xd1, 0x54, 0xf6, 0x57, 0x7f, 0x3d, 0x8c, 0x9b, 0x66, 0x52,
	0x1f, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x2e, 0x32, 0x67, 0x04, 0x09, 0x00, 0x00,
}
