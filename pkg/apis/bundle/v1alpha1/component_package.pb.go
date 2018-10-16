// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apis/bundle/v1alpha1/component_package.proto

package v1alpha1 // import "github.com/GoogleCloudPlatform/k8s-cluster-bundle/pkg/apis/bundle/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Kubernetes objects grouped into cluster components and versioned together.
// These could be applications or they could be some sort of supporting
// collection of objects.
type ComponentPackage struct {
	// Required. Kubernetes API Version for the ComponentPackage type.
	// Must always be 'gke.io/k8s-cluster-bundle/v1alpha1'
	ApiVersion string `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	// Required. The Kubernetes `kind` for this object. Must be
	// 'ComponentPackage'.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Kubernetes Metadata for the component. The Metadata.name field must be
	// filled out and each component in a bundle must have a unique name. For
	// example you might have a 'kube-apiserver' component or perhaps even a
	// 'kubernetes' component, depending on the granulatarity of the Bundle
	// components.
	Metadata *ObjectMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec for the ComponentPackage.
	Spec                 *ComponentPackageSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ComponentPackage) Reset()         { *m = ComponentPackage{} }
func (m *ComponentPackage) String() string { return proto.CompactTextString(m) }
func (*ComponentPackage) ProtoMessage()    {}
func (*ComponentPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_500109098fa4d83a, []int{0}
}
func (m *ComponentPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentPackage.Unmarshal(m, b)
}
func (m *ComponentPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentPackage.Marshal(b, m, deterministic)
}
func (m *ComponentPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentPackage.Merge(m, src)
}
func (m *ComponentPackage) XXX_Size() int {
	return xxx_messageInfo_ComponentPackage.Size(m)
}
func (m *ComponentPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentPackage.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentPackage proto.InternalMessageInfo

func (m *ComponentPackage) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *ComponentPackage) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ComponentPackage) GetMetadata() *ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ComponentPackage) GetSpec() *ComponentPackageSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// ComponentPackageSpec represents the spec for the component.
type ComponentPackageSpec struct {
	// Version-string for this component. The version should be a SemVer 2 string
	// (see https://semver.org/) of the form X.Y.Z (Major.Minor.Patch).  A
	// major-version changes should indicate breaking changes, minor-versions
	// should indicate backwards compatible features, and patch changes should
	// indicate backwords compatible. If there are any changes to the component,
	// then the version string must be incremented.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Optional. A version-string representing the version of the API this
	// component offers to other components (for the purposes of requirement
	// satisfaction). If no API is offered to other components, then this field
	// may be blank, but then other components may not depend on this component.
	ComponentApiVersion string `protobuf:"bytes,2,opt,name=componentApiVersion,proto3" json:"componentApiVersion,omitempty"`
	// A list of components that must be packaged in a bundle in with this
	// component.
	Requirements []*MinRequirement `protobuf:"bytes,3,rep,name=requirements,proto3" json:"requirements,omitempty"`
	// Structured Kubenetes objects that run as part of this app, whether on the
	// master, on the nodes, or in some other fashio.  These Kubernetes objects
	// are inlined and must be YAML/JSON
	// compatible. Each must have `apiVersion`, `kind`, and `metadata`
	ClusterObjects []*_struct.Struct `protobuf:"bytes,4,rep,name=cluster_objects,json=clusterObjects,proto3" json:"cluster_objects,omitempty"`
	// Cluster objects that are specified via a File-URL. The process of inlining
	// the a component turns cluster object files into cluster objects.
	// During the inline process, if the file is YAML-formatted and contains multiple
	// objects, the objects will be split into separate inline objects. In other
	// words, one cluster object file may result in multiple cluster objects.
	ClusterObjectFiles   []*File  `protobuf:"bytes,5,rep,name=cluster_object_files,json=clusterObjectFiles,proto3" json:"cluster_object_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComponentPackageSpec) Reset()         { *m = ComponentPackageSpec{} }
func (m *ComponentPackageSpec) String() string { return proto.CompactTextString(m) }
func (*ComponentPackageSpec) ProtoMessage()    {}
func (*ComponentPackageSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_500109098fa4d83a, []int{1}
}
func (m *ComponentPackageSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentPackageSpec.Unmarshal(m, b)
}
func (m *ComponentPackageSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentPackageSpec.Marshal(b, m, deterministic)
}
func (m *ComponentPackageSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentPackageSpec.Merge(m, src)
}
func (m *ComponentPackageSpec) XXX_Size() int {
	return xxx_messageInfo_ComponentPackageSpec.Size(m)
}
func (m *ComponentPackageSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentPackageSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentPackageSpec proto.InternalMessageInfo

func (m *ComponentPackageSpec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ComponentPackageSpec) GetComponentApiVersion() string {
	if m != nil {
		return m.ComponentApiVersion
	}
	return ""
}

func (m *ComponentPackageSpec) GetRequirements() []*MinRequirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *ComponentPackageSpec) GetClusterObjects() []*_struct.Struct {
	if m != nil {
		return m.ClusterObjects
	}
	return nil
}

func (m *ComponentPackageSpec) GetClusterObjectFiles() []*File {
	if m != nil {
		return m.ClusterObjectFiles
	}
	return nil
}

// MinRequirement represents a component that this component must be packaged
// with in a ClusterBundle (or must be satisfied in some other fashion). This is
// roughly based on Semantic Import Versioning in the Go Language:
// https://research.swtch.com/vgo-import.
type MinRequirement struct {
	// Name of a component. The component name specified here must match exactly a
	// component name in the `metadata.name` field of another component.
	Component string `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	// The sem-ver apiVersion of the component. The API Version is only a minimum
	// requirement. The assumption any newer component with only backwards
	// compatable changes is acceptable.
	ComponentApiVersion  string   `protobuf:"bytes,2,opt,name=componentApiVersion,proto3" json:"componentApiVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MinRequirement) Reset()         { *m = MinRequirement{} }
func (m *MinRequirement) String() string { return proto.CompactTextString(m) }
func (*MinRequirement) ProtoMessage()    {}
func (*MinRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_500109098fa4d83a, []int{2}
}
func (m *MinRequirement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MinRequirement.Unmarshal(m, b)
}
func (m *MinRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MinRequirement.Marshal(b, m, deterministic)
}
func (m *MinRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinRequirement.Merge(m, src)
}
func (m *MinRequirement) XXX_Size() int {
	return xxx_messageInfo_MinRequirement.Size(m)
}
func (m *MinRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_MinRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_MinRequirement proto.InternalMessageInfo

func (m *MinRequirement) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *MinRequirement) GetComponentApiVersion() string {
	if m != nil {
		return m.ComponentApiVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*ComponentPackage)(nil), "pkg.apis.bundle.v1alpha1.ComponentPackage")
	proto.RegisterType((*ComponentPackageSpec)(nil), "pkg.apis.bundle.v1alpha1.ComponentPackageSpec")
	proto.RegisterType((*MinRequirement)(nil), "pkg.apis.bundle.v1alpha1.MinRequirement")
}

func init() {
	proto.RegisterFile("pkg/apis/bundle/v1alpha1/component_package.proto", fileDescriptor_500109098fa4d83a)
}

var fileDescriptor_500109098fa4d83a = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6f, 0xd4, 0x30,
	0x10, 0xc5, 0xb5, 0xdd, 0xe5, 0x4f, 0x5d, 0x54, 0x90, 0xa9, 0x84, 0x55, 0x55, 0xd5, 0x6a, 0xc5,
	0x21, 0x42, 0xd4, 0x6e, 0xcb, 0x85, 0x63, 0x69, 0x25, 0x90, 0x50, 0x2b, 0x56, 0xa9, 0xc4, 0x81,
	0x4b, 0x70, 0x9c, 0xd9, 0xd4, 0xc4, 0xb1, 0x4d, 0xec, 0xf4, 0x13, 0x72, 0xe4, 0x43, 0xa1, 0xb5,
	0x93, 0x5d, 0x82, 0x9a, 0x03, 0xb7, 0xd8, 0x33, 0xbf, 0xc9, 0x9b, 0xe7, 0x87, 0x4e, 0x6d, 0x55,
	0x32, 0x6e, 0xa5, 0x63, 0x79, 0xab, 0x0b, 0x05, 0xec, 0xfe, 0x8c, 0x2b, 0x7b, 0xc7, 0xcf, 0x98,
	0x30, 0xb5, 0x35, 0x1a, 0xb4, 0xcf, 0x2c, 0x17, 0x15, 0x2f, 0x81, 0xda, 0xc6, 0x78, 0x83, 0x89,
	0xad, 0x4a, 0xba, 0x26, 0x68, 0x24, 0x68, 0x4f, 0x1c, 0x1e, 0x95, 0xc6, 0x94, 0x0a, 0x58, 0xe8,
	0xcb, 0xdb, 0x15, 0x73, 0xbe, 0x69, 0x85, 0x8f, 0xdc, 0xe1, 0xdb, 0xd1, 0x3f, 0xc5, 0x73, 0x26,
	0x4c, 0x5d, 0x1b, 0xdd, 0x75, 0xbf, 0x19, 0xed, 0x36, 0xf9, 0x0f, 0x10, 0x3e, 0xab, 0xc1, 0xf3,
	0xd8, 0xbb, 0xf8, 0x3d, 0x41, 0x2f, 0xae, 0x7a, 0xb5, 0xcb, 0x28, 0x16, 0x1f, 0x23, 0xc4, 0xad,
	0xfc, 0x0a, 0x8d, 0x93, 0x46, 0x93, 0xc9, 0x7c, 0x92, 0xec, 0xa6, 0x7f, 0xdd, 0x60, 0x8c, 0x66,
	0x95, 0xd4, 0x05, 0xd9, 0x09, 0x95, 0xf0, 0x8d, 0x2f, 0xd0, 0xd3, 0xf5, 0xd8, 0x82, 0x7b, 0x4e,
	0xa6, 0xf3, 0x49, 0xb2, 0x77, 0xfe, 0x9a, 0x8e, 0x6d, 0x4b, 0xbf, 0x04, 0x1d, 0x37, 0xe0, 0x79,
	0xba, 0xa1, 0xf0, 0x25, 0x9a, 0x39, 0x0b, 0x82, 0xcc, 0x02, 0x4d, 0xc7, 0xe9, 0x7f, 0xf5, 0xde,
	0x5a, 0x10, 0x69, 0x60, 0x17, 0xbf, 0x76, 0xd0, 0xc1, 0x43, 0x65, 0x4c, 0xd0, 0x93, 0xfb, 0xc1,
	0x3e, 0xfd, 0x11, 0x9f, 0xa2, 0x97, 0x9b, 0xe7, 0xfa, 0xb0, 0xdd, 0x3a, 0xee, 0xf6, 0x50, 0x09,
	0x5f, 0xa3, 0x67, 0x0d, 0xfc, 0x6c, 0x65, 0x03, 0x35, 0x68, 0xef, 0xc8, 0x74, 0x3e, 0x4d, 0xf6,
	0xce, 0x93, 0x71, 0xc1, 0x37, 0x52, 0xa7, 0x5b, 0x20, 0x1d, 0xd0, 0xf8, 0x02, 0x3d, 0x17, 0xaa,
	0x75, 0x1e, 0x9a, 0x2c, 0x3e, 0x8f, 0x23, 0xb3, 0x30, 0xf0, 0x15, 0x8d, 0x99, 0xa0, 0x7d, 0x26,
	0xe8, 0x6d, 0xc8, 0x44, 0xba, 0xdf, 0xf5, 0x47, 0x17, 0x1d, 0x5e, 0xa2, 0x83, 0xe1, 0x84, 0x6c,
	0x25, 0x15, 0x38, 0xf2, 0x28, 0x8c, 0x39, 0x1e, 0xd7, 0xf5, 0x51, 0x2a, 0x48, 0xf1, 0x60, 0xda,
	0xfa, 0xca, 0x2d, 0xbe, 0xa3, 0xfd, 0xa1, 0x66, 0x7c, 0x84, 0x76, 0x37, 0x56, 0x74, 0x0e, 0x6e,
	0x2f, 0xfe, 0xdf, 0xc3, 0xcb, 0xeb, 0x6f, 0x9f, 0x4b, 0xe9, 0xef, 0xda, 0x9c, 0x0a, 0x53, 0xb3,
	0x4f, 0x61, 0xd1, 0x2b, 0x65, 0xda, 0x62, 0xa9, 0xb8, 0x5f, 0x99, 0xa6, 0x66, 0xd5, 0x7b, 0x77,
	0xd2, 0x49, 0x3b, 0xe9, 0x72, 0x3c, 0x96, 0xeb, 0xfc, 0x71, 0xb0, 0xe8, 0xdd, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x91, 0xb1, 0xc2, 0x34, 0x92, 0x03, 0x00, 0x00,
}