// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/proxy_template.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// ProxyTemplate defines the desired state of ProxyTemplate
type ProxyTemplate struct {
	// List of Dataplane selectors.
	// +optional
	Selectors []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// Configuration for ProxyTemplate
	Conf                 *ProxyTemplate_Conf `protobuf:"bytes,2,opt,name=conf,proto3" json:"conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProxyTemplate) Reset()         { *m = ProxyTemplate{} }
func (m *ProxyTemplate) String() string { return proto.CompactTextString(m) }
func (*ProxyTemplate) ProtoMessage()    {}
func (*ProxyTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_129e53d675ac14f4, []int{0}
}

func (m *ProxyTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyTemplate.Unmarshal(m, b)
}
func (m *ProxyTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyTemplate.Marshal(b, m, deterministic)
}
func (m *ProxyTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyTemplate.Merge(m, src)
}
func (m *ProxyTemplate) XXX_Size() int {
	return xxx_messageInfo_ProxyTemplate.Size(m)
}
func (m *ProxyTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyTemplate proto.InternalMessageInfo

func (m *ProxyTemplate) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *ProxyTemplate) GetConf() *ProxyTemplate_Conf {
	if m != nil {
		return m.Conf
	}
	return nil
}

type ProxyTemplate_Conf struct {
	// List of imported profiles.
	// +optional
	Imports []string `protobuf:"bytes,1,rep,name=imports,proto3" json:"imports,omitempty"`
	// List of raw xDS resources.
	// +optional
	Resources            []*ProxyTemplateRawResource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ProxyTemplate_Conf) Reset()         { *m = ProxyTemplate_Conf{} }
func (m *ProxyTemplate_Conf) String() string { return proto.CompactTextString(m) }
func (*ProxyTemplate_Conf) ProtoMessage()    {}
func (*ProxyTemplate_Conf) Descriptor() ([]byte, []int) {
	return fileDescriptor_129e53d675ac14f4, []int{0, 0}
}

func (m *ProxyTemplate_Conf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyTemplate_Conf.Unmarshal(m, b)
}
func (m *ProxyTemplate_Conf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyTemplate_Conf.Marshal(b, m, deterministic)
}
func (m *ProxyTemplate_Conf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyTemplate_Conf.Merge(m, src)
}
func (m *ProxyTemplate_Conf) XXX_Size() int {
	return xxx_messageInfo_ProxyTemplate_Conf.Size(m)
}
func (m *ProxyTemplate_Conf) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyTemplate_Conf.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyTemplate_Conf proto.InternalMessageInfo

func (m *ProxyTemplate_Conf) GetImports() []string {
	if m != nil {
		return m.Imports
	}
	return nil
}

func (m *ProxyTemplate_Conf) GetResources() []*ProxyTemplateRawResource {
	if m != nil {
		return m.Resources
	}
	return nil
}

type ProxyTemplateSource struct {
	// Name of a configuration source.
	// +optional
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*ProxyTemplateSource_Profile
	//	*ProxyTemplateSource_Raw
	Type                 isProxyTemplateSource_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ProxyTemplateSource) Reset()         { *m = ProxyTemplateSource{} }
func (m *ProxyTemplateSource) String() string { return proto.CompactTextString(m) }
func (*ProxyTemplateSource) ProtoMessage()    {}
func (*ProxyTemplateSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_129e53d675ac14f4, []int{1}
}

func (m *ProxyTemplateSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyTemplateSource.Unmarshal(m, b)
}
func (m *ProxyTemplateSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyTemplateSource.Marshal(b, m, deterministic)
}
func (m *ProxyTemplateSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyTemplateSource.Merge(m, src)
}
func (m *ProxyTemplateSource) XXX_Size() int {
	return xxx_messageInfo_ProxyTemplateSource.Size(m)
}
func (m *ProxyTemplateSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyTemplateSource.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyTemplateSource proto.InternalMessageInfo

func (m *ProxyTemplateSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isProxyTemplateSource_Type interface {
	isProxyTemplateSource_Type()
}

type ProxyTemplateSource_Profile struct {
	Profile *ProxyTemplateProfileSource `protobuf:"bytes,2,opt,name=profile,proto3,oneof"`
}

type ProxyTemplateSource_Raw struct {
	Raw *ProxyTemplateRawSource `protobuf:"bytes,3,opt,name=raw,proto3,oneof"`
}

func (*ProxyTemplateSource_Profile) isProxyTemplateSource_Type() {}

func (*ProxyTemplateSource_Raw) isProxyTemplateSource_Type() {}

func (m *ProxyTemplateSource) GetType() isProxyTemplateSource_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ProxyTemplateSource) GetProfile() *ProxyTemplateProfileSource {
	if x, ok := m.GetType().(*ProxyTemplateSource_Profile); ok {
		return x.Profile
	}
	return nil
}

func (m *ProxyTemplateSource) GetRaw() *ProxyTemplateRawSource {
	if x, ok := m.GetType().(*ProxyTemplateSource_Raw); ok {
		return x.Raw
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ProxyTemplateSource) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ProxyTemplateSource_Profile)(nil),
		(*ProxyTemplateSource_Raw)(nil),
	}
}

type ProxyTemplateProfileSource struct {
	// Profile name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Profile params if any.
	// +optional
	Params               map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProxyTemplateProfileSource) Reset()         { *m = ProxyTemplateProfileSource{} }
func (m *ProxyTemplateProfileSource) String() string { return proto.CompactTextString(m) }
func (*ProxyTemplateProfileSource) ProtoMessage()    {}
func (*ProxyTemplateProfileSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_129e53d675ac14f4, []int{2}
}

func (m *ProxyTemplateProfileSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyTemplateProfileSource.Unmarshal(m, b)
}
func (m *ProxyTemplateProfileSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyTemplateProfileSource.Marshal(b, m, deterministic)
}
func (m *ProxyTemplateProfileSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyTemplateProfileSource.Merge(m, src)
}
func (m *ProxyTemplateProfileSource) XXX_Size() int {
	return xxx_messageInfo_ProxyTemplateProfileSource.Size(m)
}
func (m *ProxyTemplateProfileSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyTemplateProfileSource.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyTemplateProfileSource proto.InternalMessageInfo

func (m *ProxyTemplateProfileSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProxyTemplateProfileSource) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type ProxyTemplateRawSource struct {
	// List of raw xDS resources.
	// +optional
	Resources            []*ProxyTemplateRawResource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ProxyTemplateRawSource) Reset()         { *m = ProxyTemplateRawSource{} }
func (m *ProxyTemplateRawSource) String() string { return proto.CompactTextString(m) }
func (*ProxyTemplateRawSource) ProtoMessage()    {}
func (*ProxyTemplateRawSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_129e53d675ac14f4, []int{3}
}

func (m *ProxyTemplateRawSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyTemplateRawSource.Unmarshal(m, b)
}
func (m *ProxyTemplateRawSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyTemplateRawSource.Marshal(b, m, deterministic)
}
func (m *ProxyTemplateRawSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyTemplateRawSource.Merge(m, src)
}
func (m *ProxyTemplateRawSource) XXX_Size() int {
	return xxx_messageInfo_ProxyTemplateRawSource.Size(m)
}
func (m *ProxyTemplateRawSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyTemplateRawSource.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyTemplateRawSource proto.InternalMessageInfo

func (m *ProxyTemplateRawSource) GetResources() []*ProxyTemplateRawResource {
	if m != nil {
		return m.Resources
	}
	return nil
}

type ProxyTemplateRawResource struct {
	// The resource's name, to distinguish it from others of the same type of
	// resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource level version. It allows xDS to track the state of individual
	// resources.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// xDS resource.
	Resource             string   `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyTemplateRawResource) Reset()         { *m = ProxyTemplateRawResource{} }
func (m *ProxyTemplateRawResource) String() string { return proto.CompactTextString(m) }
func (*ProxyTemplateRawResource) ProtoMessage()    {}
func (*ProxyTemplateRawResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_129e53d675ac14f4, []int{4}
}

func (m *ProxyTemplateRawResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyTemplateRawResource.Unmarshal(m, b)
}
func (m *ProxyTemplateRawResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyTemplateRawResource.Marshal(b, m, deterministic)
}
func (m *ProxyTemplateRawResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyTemplateRawResource.Merge(m, src)
}
func (m *ProxyTemplateRawResource) XXX_Size() int {
	return xxx_messageInfo_ProxyTemplateRawResource.Size(m)
}
func (m *ProxyTemplateRawResource) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyTemplateRawResource.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyTemplateRawResource proto.InternalMessageInfo

func (m *ProxyTemplateRawResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProxyTemplateRawResource) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ProxyTemplateRawResource) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func init() {
	proto.RegisterType((*ProxyTemplate)(nil), "kuma.mesh.v1alpha1.ProxyTemplate")
	proto.RegisterType((*ProxyTemplate_Conf)(nil), "kuma.mesh.v1alpha1.ProxyTemplate.Conf")
	proto.RegisterType((*ProxyTemplateSource)(nil), "kuma.mesh.v1alpha1.ProxyTemplateSource")
	proto.RegisterType((*ProxyTemplateProfileSource)(nil), "kuma.mesh.v1alpha1.ProxyTemplateProfileSource")
	proto.RegisterMapType((map[string]string)(nil), "kuma.mesh.v1alpha1.ProxyTemplateProfileSource.ParamsEntry")
	proto.RegisterType((*ProxyTemplateRawSource)(nil), "kuma.mesh.v1alpha1.ProxyTemplateRawSource")
	proto.RegisterType((*ProxyTemplateRawResource)(nil), "kuma.mesh.v1alpha1.ProxyTemplateRawResource")
}

func init() {
	proto.RegisterFile("mesh/v1alpha1/proxy_template.proto", fileDescriptor_129e53d675ac14f4)
}

var fileDescriptor_129e53d675ac14f4 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x8b, 0xda, 0x40,
	0x18, 0xee, 0x98, 0x54, 0xcd, 0x2b, 0x85, 0x32, 0x2d, 0x65, 0x08, 0x1e, 0x24, 0x87, 0x22, 0xa5,
	0x8c, 0x68, 0x2f, 0x6d, 0x0e, 0x3d, 0x58, 0x0a, 0xc5, 0x93, 0x8c, 0x3d, 0xf5, 0x52, 0xa6, 0x3a,
	0xa2, 0x98, 0x64, 0x86, 0x49, 0xd4, 0xe6, 0xc7, 0x2d, 0xfb, 0xab, 0xf6, 0xbe, 0x64, 0x92, 0xd1,
	0x0d, 0x1b, 0x71, 0x97, 0xbd, 0xbd, 0x4f, 0xf2, 0x7c, 0xbc, 0x1f, 0x0c, 0x04, 0xb1, 0x48, 0x37,
	0xa3, 0xc3, 0x98, 0x47, 0x6a, 0xc3, 0xc7, 0x23, 0xa5, 0xe5, 0xff, 0xfc, 0x6f, 0x26, 0x62, 0x15,
	0xf1, 0x4c, 0x50, 0xa5, 0x65, 0x26, 0x31, 0xde, 0xed, 0x63, 0x4e, 0x0b, 0x22, 0xb5, 0x44, 0xbf,
	0x5f, 0xd7, 0xa5, 0x22, 0x12, 0xcb, 0x4c, 0xea, 0x52, 0x11, 0xdc, 0x21, 0x78, 0x33, 0x2f, 0xac,
	0x7e, 0x57, 0x4e, 0x38, 0x04, 0xcf, 0x72, 0x52, 0x82, 0x06, 0xce, 0xb0, 0x37, 0xe9, 0xd3, 0xc7,
	0xbe, 0x74, 0x51, 0x91, 0xd8, 0x99, 0x8e, 0x43, 0x70, 0x97, 0x32, 0x59, 0x93, 0xd6, 0x00, 0x0d,
	0x7b, 0x93, 0x8f, 0x4d, 0xb2, 0x5a, 0x18, 0xfd, 0x21, 0x93, 0x35, 0x33, 0x1a, 0x3f, 0x02, 0xb7,
	0x40, 0x98, 0x40, 0x67, 0x1b, 0x2b, 0xa9, 0xb3, 0x32, 0xdd, 0x63, 0x16, 0xe2, 0x19, 0x78, 0x5a,
	0xa4, 0x72, 0xaf, 0x97, 0x22, 0x25, 0x2d, 0xd3, 0xd9, 0xe7, 0xab, 0x11, 0x8c, 0x1f, 0x59, 0x25,
	0x62, 0x67, 0x79, 0x70, 0x83, 0xe0, 0x5d, 0x8d, 0xb7, 0x30, 0x3f, 0x30, 0x06, 0x37, 0xe1, 0xb1,
	0x20, 0x68, 0x80, 0x86, 0x1e, 0x33, 0x35, 0x9e, 0x41, 0x47, 0x69, 0xb9, 0xde, 0x46, 0xa2, 0x1a,
	0x8c, 0x5e, 0x4d, 0x9d, 0x97, 0xfc, 0xd2, 0xf4, 0xd7, 0x2b, 0x66, 0x0d, 0xf0, 0x77, 0x70, 0x34,
	0x3f, 0x12, 0xc7, 0xf8, 0x7c, 0x7a, 0x4a, 0xf7, 0x27, 0x8f, 0x42, 0x38, 0x6d, 0x83, 0x9b, 0xe5,
	0x4a, 0x04, 0xb7, 0x08, 0xfc, 0xcb, 0x89, 0x8d, 0x63, 0x30, 0x68, 0x2b, 0xae, 0x79, 0x6c, 0x77,
	0x17, 0x3e, 0x6f, 0x0a, 0x3a, 0x37, 0xe2, 0x9f, 0x49, 0xa6, 0x73, 0x56, 0x39, 0xf9, 0xdf, 0xa0,
	0xf7, 0xe0, 0x33, 0x7e, 0x0b, 0xce, 0x4e, 0xe4, 0x55, 0x6a, 0x51, 0xe2, 0xf7, 0xf0, 0xfa, 0xc0,
	0xa3, 0x7d, 0xb9, 0x39, 0x8f, 0x95, 0x20, 0x6c, 0x7d, 0x45, 0xc1, 0x0a, 0x3e, 0x34, 0x8f, 0x5a,
	0xbf, 0x33, 0x7a, 0xd9, 0x9d, 0x57, 0x40, 0x2e, 0xd1, 0x1a, 0x97, 0x44, 0xa0, 0x73, 0x10, 0x3a,
	0xdd, 0xca, 0xa4, 0xea, 0xd8, 0x42, 0xec, 0x43, 0xd7, 0xda, 0x9a, 0xf3, 0x79, 0xec, 0x84, 0xa7,
	0xf0, 0xa7, 0x6b, 0xbb, 0xfa, 0xd7, 0x36, 0x0f, 0xeb, 0xcb, 0x7d, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xfb, 0x1c, 0xe2, 0x60, 0xb0, 0x03, 0x00, 0x00,
}
