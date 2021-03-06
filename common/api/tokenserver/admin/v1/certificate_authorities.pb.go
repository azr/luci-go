// Code generated by protoc-gen-go.
// source: certificate_authorities.proto
// DO NOT EDIT!

/*
Package admin is a generated protocol buffer package.

It is generated from these files:
	certificate_authorities.proto
	config.proto
	service_accounts.proto

It has these top-level messages:
	ImportConfigRequest
	ImportConfigResponse
	FetchCRLRequest
	FetchCRLResponse
	ListCAsResponse
	GetCAStatusRequest
	GetCAStatusResponse
	IsRevokedCertRequest
	IsRevokedCertResponse
	CheckCertificateRequest
	CheckCertificateResponse
	CRLStatus
	TokenServerConfig
	CertificateAuthorityConfig
	DomainConfig
	CreateServiceAccountRequest
	CreateServiceAccountResponse
*/
package admin

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"
import google_protobuf1 "github.com/luci/luci-go/common/proto/google"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ImportConfigRequest is passed to ImportConfig.
type ImportConfigRequest struct {
	// DevConfig is mapping of {config file name -> config file body}.
	//
	// It is used only on devserver to import some mock config in integration
	// tests. Ignored completely in prod.
	DevConfig map[string]string `protobuf:"bytes,1,rep,name=dev_config,json=devConfig" json:"dev_config,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ImportConfigRequest) Reset()                    { *m = ImportConfigRequest{} }
func (m *ImportConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportConfigRequest) ProtoMessage()               {}
func (*ImportConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ImportConfigRequest) GetDevConfig() map[string]string {
	if m != nil {
		return m.DevConfig
	}
	return nil
}

// ImportConfigResponse is returned by ImportConfig on success.
type ImportConfigResponse struct {
	Revision string `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
}

func (m *ImportConfigResponse) Reset()                    { *m = ImportConfigResponse{} }
func (m *ImportConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ImportConfigResponse) ProtoMessage()               {}
func (*ImportConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// FetchCRLRequest identifies a name of CA to fetch CRL for.
type FetchCRLRequest struct {
	Cn    string `protobuf:"bytes,1,opt,name=cn" json:"cn,omitempty"`
	Force bool   `protobuf:"varint,2,opt,name=force" json:"force,omitempty"`
}

func (m *FetchCRLRequest) Reset()                    { *m = FetchCRLRequest{} }
func (m *FetchCRLRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchCRLRequest) ProtoMessage()               {}
func (*FetchCRLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// FetchCRLResponse is returned by FetchCRL.
type FetchCRLResponse struct {
	CrlStatus *CRLStatus `protobuf:"bytes,1,opt,name=crl_status,json=crlStatus" json:"crl_status,omitempty"`
}

func (m *FetchCRLResponse) Reset()                    { *m = FetchCRLResponse{} }
func (m *FetchCRLResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchCRLResponse) ProtoMessage()               {}
func (*FetchCRLResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FetchCRLResponse) GetCrlStatus() *CRLStatus {
	if m != nil {
		return m.CrlStatus
	}
	return nil
}

// ListCAsResponse is returned by ListCAs.
type ListCAsResponse struct {
	Cn []string `protobuf:"bytes,1,rep,name=cn" json:"cn,omitempty"`
}

func (m *ListCAsResponse) Reset()                    { *m = ListCAsResponse{} }
func (m *ListCAsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListCAsResponse) ProtoMessage()               {}
func (*ListCAsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// GetCAStatusRequest identifies a name of CA to fetch.
type GetCAStatusRequest struct {
	Cn string `protobuf:"bytes,1,opt,name=cn" json:"cn,omitempty"`
}

func (m *GetCAStatusRequest) Reset()                    { *m = GetCAStatusRequest{} }
func (m *GetCAStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCAStatusRequest) ProtoMessage()               {}
func (*GetCAStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// GetCAStatusResponse is returned by GetCAStatus method.
//
// If requested CA doesn't exist, all fields are empty.
type GetCAStatusResponse struct {
	Config     *CertificateAuthorityConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Cert       string                      `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
	Removed    bool                        `protobuf:"varint,3,opt,name=removed" json:"removed,omitempty"`
	Ready      bool                        `protobuf:"varint,4,opt,name=ready" json:"ready,omitempty"`
	AddedRev   string                      `protobuf:"bytes,5,opt,name=added_rev,json=addedRev" json:"added_rev,omitempty"`
	UpdatedRev string                      `protobuf:"bytes,6,opt,name=updated_rev,json=updatedRev" json:"updated_rev,omitempty"`
	RemovedRev string                      `protobuf:"bytes,7,opt,name=removed_rev,json=removedRev" json:"removed_rev,omitempty"`
	CrlStatus  *CRLStatus                  `protobuf:"bytes,8,opt,name=crl_status,json=crlStatus" json:"crl_status,omitempty"`
}

func (m *GetCAStatusResponse) Reset()                    { *m = GetCAStatusResponse{} }
func (m *GetCAStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetCAStatusResponse) ProtoMessage()               {}
func (*GetCAStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetCAStatusResponse) GetConfig() *CertificateAuthorityConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *GetCAStatusResponse) GetCrlStatus() *CRLStatus {
	if m != nil {
		return m.CrlStatus
	}
	return nil
}

// IsRevokedCertRequest contains a name of the CA and a cert serial number.
type IsRevokedCertRequest struct {
	Ca string `protobuf:"bytes,1,opt,name=ca" json:"ca,omitempty"`
	Sn string `protobuf:"bytes,2,opt,name=sn" json:"sn,omitempty"`
}

func (m *IsRevokedCertRequest) Reset()                    { *m = IsRevokedCertRequest{} }
func (m *IsRevokedCertRequest) String() string            { return proto.CompactTextString(m) }
func (*IsRevokedCertRequest) ProtoMessage()               {}
func (*IsRevokedCertRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// IsRevokedCertResponse is returned by IsRevokedCert
type IsRevokedCertResponse struct {
	Revoked bool `protobuf:"varint,1,opt,name=revoked" json:"revoked,omitempty"`
}

func (m *IsRevokedCertResponse) Reset()                    { *m = IsRevokedCertResponse{} }
func (m *IsRevokedCertResponse) String() string            { return proto.CompactTextString(m) }
func (*IsRevokedCertResponse) ProtoMessage()               {}
func (*IsRevokedCertResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// CheckCertificateRequest contains a pem encoded certificate to check.
type CheckCertificateRequest struct {
	CertPem string `protobuf:"bytes,1,opt,name=cert_pem,json=certPem" json:"cert_pem,omitempty"`
}

func (m *CheckCertificateRequest) Reset()                    { *m = CheckCertificateRequest{} }
func (m *CheckCertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckCertificateRequest) ProtoMessage()               {}
func (*CheckCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// CheckCertificateResponse is returned by CheckCertificate.
type CheckCertificateResponse struct {
	IsValid       bool   `protobuf:"varint,1,opt,name=is_valid,json=isValid" json:"is_valid,omitempty"`
	InvalidReason string `protobuf:"bytes,2,opt,name=invalid_reason,json=invalidReason" json:"invalid_reason,omitempty"`
}

func (m *CheckCertificateResponse) Reset()                    { *m = CheckCertificateResponse{} }
func (m *CheckCertificateResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckCertificateResponse) ProtoMessage()               {}
func (*CheckCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// CRLStatus describes the latest known state of imported CRL.
type CRLStatus struct {
	LastUpdateTime    *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	LastFetchTime     *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=last_fetch_time,json=lastFetchTime" json:"last_fetch_time,omitempty"`
	LastFetchEtag     string                      `protobuf:"bytes,3,opt,name=last_fetch_etag,json=lastFetchEtag" json:"last_fetch_etag,omitempty"`
	RevokedCertsCount int64                       `protobuf:"varint,4,opt,name=revoked_certs_count,json=revokedCertsCount" json:"revoked_certs_count,omitempty"`
}

func (m *CRLStatus) Reset()                    { *m = CRLStatus{} }
func (m *CRLStatus) String() string            { return proto.CompactTextString(m) }
func (*CRLStatus) ProtoMessage()               {}
func (*CRLStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CRLStatus) GetLastUpdateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *CRLStatus) GetLastFetchTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastFetchTime
	}
	return nil
}

func init() {
	proto.RegisterType((*ImportConfigRequest)(nil), "tokenserver.admin.ImportConfigRequest")
	proto.RegisterType((*ImportConfigResponse)(nil), "tokenserver.admin.ImportConfigResponse")
	proto.RegisterType((*FetchCRLRequest)(nil), "tokenserver.admin.FetchCRLRequest")
	proto.RegisterType((*FetchCRLResponse)(nil), "tokenserver.admin.FetchCRLResponse")
	proto.RegisterType((*ListCAsResponse)(nil), "tokenserver.admin.ListCAsResponse")
	proto.RegisterType((*GetCAStatusRequest)(nil), "tokenserver.admin.GetCAStatusRequest")
	proto.RegisterType((*GetCAStatusResponse)(nil), "tokenserver.admin.GetCAStatusResponse")
	proto.RegisterType((*IsRevokedCertRequest)(nil), "tokenserver.admin.IsRevokedCertRequest")
	proto.RegisterType((*IsRevokedCertResponse)(nil), "tokenserver.admin.IsRevokedCertResponse")
	proto.RegisterType((*CheckCertificateRequest)(nil), "tokenserver.admin.CheckCertificateRequest")
	proto.RegisterType((*CheckCertificateResponse)(nil), "tokenserver.admin.CheckCertificateResponse")
	proto.RegisterType((*CRLStatus)(nil), "tokenserver.admin.CRLStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for CertificateAuthorities service

type CertificateAuthoritiesClient interface {
	// ImportConfig makes the server read its config from luci-config right now.
	//
	// Note that regularly configs are read in background each 5 min. ImportConfig
	// can be used to force config reread immediately. It will block until configs
	// are read.
	ImportConfig(ctx context.Context, in *ImportConfigRequest, opts ...grpc.CallOption) (*ImportConfigResponse, error)
	// FetchCRL makes the server fetch a CRL for some CA.
	FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error)
	// ListCAs returns a list of Common Names of registered CAs.
	ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error)
	// GetCAStatus returns configuration of some CA defined in the config.
	GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error)
	// IsRevokedCert says whether a certificate serial number is in the CRL.
	IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error)
	// CheckCertificate says whether a certificate is valid or not.
	CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error)
}
type certificateAuthoritiesPRPCClient struct {
	client *prpccommon.Client
}

func NewCertificateAuthoritiesPRPCClient(client *prpccommon.Client) CertificateAuthoritiesClient {
	return &certificateAuthoritiesPRPCClient{client}
}

func (c *certificateAuthoritiesPRPCClient) ImportConfig(ctx context.Context, in *ImportConfigRequest, opts ...grpc.CallOption) (*ImportConfigResponse, error) {
	out := new(ImportConfigResponse)
	err := c.client.Call(ctx, "tokenserver.admin.CertificateAuthorities", "ImportConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error) {
	out := new(FetchCRLResponse)
	err := c.client.Call(ctx, "tokenserver.admin.CertificateAuthorities", "FetchCRL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error) {
	out := new(ListCAsResponse)
	err := c.client.Call(ctx, "tokenserver.admin.CertificateAuthorities", "ListCAs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error) {
	out := new(GetCAStatusResponse)
	err := c.client.Call(ctx, "tokenserver.admin.CertificateAuthorities", "GetCAStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error) {
	out := new(IsRevokedCertResponse)
	err := c.client.Call(ctx, "tokenserver.admin.CertificateAuthorities", "IsRevokedCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error) {
	out := new(CheckCertificateResponse)
	err := c.client.Call(ctx, "tokenserver.admin.CertificateAuthorities", "CheckCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type certificateAuthoritiesClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAuthoritiesClient(cc *grpc.ClientConn) CertificateAuthoritiesClient {
	return &certificateAuthoritiesClient{cc}
}

func (c *certificateAuthoritiesClient) ImportConfig(ctx context.Context, in *ImportConfigRequest, opts ...grpc.CallOption) (*ImportConfigResponse, error) {
	out := new(ImportConfigResponse)
	err := grpc.Invoke(ctx, "/tokenserver.admin.CertificateAuthorities/ImportConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error) {
	out := new(FetchCRLResponse)
	err := grpc.Invoke(ctx, "/tokenserver.admin.CertificateAuthorities/FetchCRL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error) {
	out := new(ListCAsResponse)
	err := grpc.Invoke(ctx, "/tokenserver.admin.CertificateAuthorities/ListCAs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error) {
	out := new(GetCAStatusResponse)
	err := grpc.Invoke(ctx, "/tokenserver.admin.CertificateAuthorities/GetCAStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error) {
	out := new(IsRevokedCertResponse)
	err := grpc.Invoke(ctx, "/tokenserver.admin.CertificateAuthorities/IsRevokedCert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error) {
	out := new(CheckCertificateResponse)
	err := grpc.Invoke(ctx, "/tokenserver.admin.CertificateAuthorities/CheckCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateAuthorities service

type CertificateAuthoritiesServer interface {
	// ImportConfig makes the server read its config from luci-config right now.
	//
	// Note that regularly configs are read in background each 5 min. ImportConfig
	// can be used to force config reread immediately. It will block until configs
	// are read.
	ImportConfig(context.Context, *ImportConfigRequest) (*ImportConfigResponse, error)
	// FetchCRL makes the server fetch a CRL for some CA.
	FetchCRL(context.Context, *FetchCRLRequest) (*FetchCRLResponse, error)
	// ListCAs returns a list of Common Names of registered CAs.
	ListCAs(context.Context, *google_protobuf.Empty) (*ListCAsResponse, error)
	// GetCAStatus returns configuration of some CA defined in the config.
	GetCAStatus(context.Context, *GetCAStatusRequest) (*GetCAStatusResponse, error)
	// IsRevokedCert says whether a certificate serial number is in the CRL.
	IsRevokedCert(context.Context, *IsRevokedCertRequest) (*IsRevokedCertResponse, error)
	// CheckCertificate says whether a certificate is valid or not.
	CheckCertificate(context.Context, *CheckCertificateRequest) (*CheckCertificateResponse, error)
}

func RegisterCertificateAuthoritiesServer(s prpc.Registrar, srv CertificateAuthoritiesServer) {
	s.RegisterService(&_CertificateAuthorities_serviceDesc, srv)
}

func _CertificateAuthorities_ImportConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthoritiesServer).ImportConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.CertificateAuthorities/ImportConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthoritiesServer).ImportConfig(ctx, req.(*ImportConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorities_FetchCRL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCRLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthoritiesServer).FetchCRL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.CertificateAuthorities/FetchCRL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthoritiesServer).FetchCRL(ctx, req.(*FetchCRLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorities_ListCAs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthoritiesServer).ListCAs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.CertificateAuthorities/ListCAs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthoritiesServer).ListCAs(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorities_GetCAStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCAStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthoritiesServer).GetCAStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.CertificateAuthorities/GetCAStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthoritiesServer).GetCAStatus(ctx, req.(*GetCAStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorities_IsRevokedCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsRevokedCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthoritiesServer).IsRevokedCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.CertificateAuthorities/IsRevokedCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthoritiesServer).IsRevokedCert(ctx, req.(*IsRevokedCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorities_CheckCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthoritiesServer).CheckCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.CertificateAuthorities/CheckCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthoritiesServer).CheckCertificate(ctx, req.(*CheckCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAuthorities_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenserver.admin.CertificateAuthorities",
	HandlerType: (*CertificateAuthoritiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportConfig",
			Handler:    _CertificateAuthorities_ImportConfig_Handler,
		},
		{
			MethodName: "FetchCRL",
			Handler:    _CertificateAuthorities_FetchCRL_Handler,
		},
		{
			MethodName: "ListCAs",
			Handler:    _CertificateAuthorities_ListCAs_Handler,
		},
		{
			MethodName: "GetCAStatus",
			Handler:    _CertificateAuthorities_GetCAStatus_Handler,
		},
		{
			MethodName: "IsRevokedCert",
			Handler:    _CertificateAuthorities_IsRevokedCert_Handler,
		},
		{
			MethodName: "CheckCertificate",
			Handler:    _CertificateAuthorities_CheckCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("certificate_authorities.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 770 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x56, 0x92, 0xb6, 0x49, 0x26, 0x6d, 0x9a, 0x6e, 0x7b, 0x7a, 0x72, 0xdc, 0x73, 0xd4, 0x83,
	0xa1, 0xa5, 0x02, 0xe1, 0x8a, 0xf0, 0x2b, 0xe0, 0xa6, 0xa4, 0x01, 0x21, 0x55, 0x02, 0xb9, 0x85,
	0xab, 0x4a, 0x96, 0x6b, 0x6f, 0xd2, 0x55, 0x62, 0x3b, 0x78, 0xd7, 0x96, 0xf2, 0x48, 0xbc, 0x03,
	0x97, 0xbc, 0x0e, 0xef, 0xc0, 0xfe, 0xd9, 0xcd, 0x8f, 0x45, 0xca, 0x9d, 0x67, 0xf6, 0x9b, 0x9f,
	0x7c, 0x33, 0xf3, 0x05, 0xfe, 0xf3, 0x70, 0xcc, 0x48, 0x9f, 0x78, 0x2e, 0xc3, 0x8e, 0x9b, 0xb0,
	0xeb, 0x28, 0x26, 0x8c, 0x60, 0x6a, 0x8d, 0xe3, 0x88, 0x45, 0x68, 0x8b, 0x45, 0x43, 0x1c, 0x52,
	0x1c, 0xa7, 0x38, 0xb6, 0x5c, 0x3f, 0x20, 0xa1, 0xb1, 0x37, 0x88, 0xa2, 0xc1, 0x08, 0x1f, 0x4b,
	0xc0, 0x55, 0xd2, 0x3f, 0xc6, 0xc1, 0x98, 0x4d, 0x14, 0xde, 0xd8, 0x9f, 0x7f, 0x64, 0x24, 0xc0,
	0x94, 0xb9, 0xc1, 0x58, 0x03, 0xd6, 0xbd, 0x28, 0xec, 0x93, 0x81, 0xb2, 0xcc, 0x6f, 0x25, 0xd8,
	0xfe, 0x10, 0x8c, 0xa3, 0x98, 0x75, 0xa5, 0xdb, 0xc6, 0x5f, 0x13, 0x0e, 0x47, 0x17, 0x00, 0x3e,
	0x4e, 0x1d, 0x85, 0x6d, 0x97, 0xfe, 0xaf, 0x1c, 0x35, 0x3a, 0xcf, 0xac, 0x85, 0x5e, 0xac, 0x82,
	0x58, 0xeb, 0x14, 0xa7, 0xca, 0xd1, 0x0b, 0x59, 0x3c, 0xb1, 0xeb, 0x7e, 0x66, 0x1b, 0x6f, 0xa0,
	0x39, 0xfb, 0x88, 0x5a, 0x50, 0x19, 0xe2, 0x09, 0x2f, 0x50, 0x3a, 0xaa, 0xdb, 0xe2, 0x13, 0xed,
	0xc0, 0x6a, 0xea, 0x8e, 0x12, 0xdc, 0x2e, 0x4b, 0x9f, 0x32, 0x5e, 0x95, 0x5f, 0x96, 0xcc, 0x0e,
	0xec, 0xcc, 0x96, 0xa3, 0xe3, 0x88, 0x77, 0x83, 0x0c, 0xa8, 0xc5, 0x38, 0x25, 0x94, 0x44, 0xa1,
	0x4e, 0x94, 0xdb, 0xe6, 0x0b, 0xd8, 0x7c, 0x87, 0x99, 0x77, 0xdd, 0xb5, 0xcf, 0xb2, 0x9f, 0xd6,
	0x84, 0xb2, 0x97, 0x01, 0xf9, 0x97, 0x28, 0xd8, 0x8f, 0x62, 0x4f, 0x15, 0xac, 0xd9, 0xca, 0x30,
	0x3f, 0x42, 0xeb, 0x26, 0x50, 0x17, 0x7a, 0x0d, 0xe0, 0xc5, 0x23, 0x87, 0xb3, 0xc9, 0x12, 0x2a,
	0x33, 0x34, 0x3a, 0xff, 0x16, 0x90, 0xc2, 0x63, 0xce, 0x25, 0xc6, 0xae, 0x73, 0xbc, 0xfa, 0x34,
	0xef, 0xc0, 0xe6, 0x19, 0xa1, 0xac, 0x7b, 0x42, 0xf3, 0x7c, 0x59, 0x27, 0x15, 0xd5, 0x89, 0x79,
	0x0f, 0xd0, 0x7b, 0xcc, 0x11, 0x3a, 0xb8, 0xb8, 0x5f, 0xf3, 0x7b, 0x19, 0xb6, 0x67, 0x60, 0x3a,
	0x5b, 0x0f, 0xd6, 0xf2, 0x71, 0x89, 0xce, 0x1e, 0x15, 0x75, 0x76, 0xb3, 0x6b, 0x27, 0x7a, 0xd5,
	0x26, 0x9a, 0x4d, 0x1d, 0x8c, 0x10, 0xac, 0x88, 0x8d, 0xd4, 0xf4, 0xcb, 0x6f, 0xd4, 0x86, 0x6a,
	0x8c, 0x83, 0x28, 0xc5, 0x7e, 0xbb, 0x22, 0x49, 0xca, 0x4c, 0x41, 0x5e, 0x8c, 0x5d, 0x7f, 0xd2,
	0x5e, 0x51, 0xe4, 0x49, 0x03, 0xed, 0x41, 0xdd, 0xf5, 0x7d, 0xec, 0x3b, 0x7c, 0x0e, 0xed, 0x55,
	0x35, 0x12, 0xe9, 0xb0, 0x71, 0x8a, 0xf6, 0xa1, 0x91, 0x8c, 0x7d, 0xde, 0x81, 0x7a, 0x5e, 0x93,
	0xcf, 0xa0, 0x5d, 0x1a, 0xa0, 0xd3, 0x4b, 0x40, 0x55, 0x01, 0xb4, 0x4b, 0x00, 0x66, 0xe7, 0x50,
	0xfb, 0xb3, 0x39, 0x3c, 0xe7, 0x5b, 0xc4, 0x49, 0x4b, 0x39, 0xda, 0x17, 0x74, 0x4c, 0xd3, 0xec,
	0xe6, 0x34, 0xbb, 0xc2, 0xa6, 0xa1, 0x66, 0x81, 0x7f, 0x99, 0x8f, 0xe1, 0xaf, 0xb9, 0x38, 0xcd,
	0xbb, 0x24, 0x47, 0xba, 0x65, 0xb4, 0x24, 0x47, 0x9a, 0xe6, 0x53, 0xf8, 0xbb, 0x7b, 0x8d, 0xbd,
	0xe1, 0x14, 0xeb, 0x59, 0xb5, 0x7f, 0xa0, 0x26, 0x98, 0x75, 0xc6, 0x38, 0xd0, 0x35, 0xab, 0xc2,
	0xfe, 0x84, 0x03, 0xf3, 0x12, 0xda, 0x8b, 0x51, 0xba, 0x16, 0x0f, 0x23, 0xd4, 0xe1, 0x27, 0x41,
	0xf2, 0x62, 0x84, 0x7e, 0x11, 0x26, 0x3a, 0x80, 0x26, 0x09, 0xe5, 0x0b, 0x67, 0xcd, 0xa5, 0x51,
	0xd6, 0xfb, 0x86, 0xf6, 0xda, 0xd2, 0x69, 0xfe, 0x2c, 0x41, 0x3d, 0xe7, 0x05, 0x9d, 0x42, 0x6b,
	0xe4, 0x52, 0xe6, 0x28, 0xf6, 0x1d, 0xa1, 0x15, 0x7a, 0x7b, 0x0c, 0x4b, 0x09, 0x89, 0x95, 0x09,
	0x89, 0x75, 0x91, 0x09, 0x89, 0xdd, 0x14, 0x31, 0x9f, 0x65, 0x88, 0x70, 0xa2, 0xb7, 0xb0, 0x29,
	0xb3, 0xf4, 0xc5, 0xc1, 0xa8, 0x24, 0xe5, 0xa5, 0x49, 0x36, 0x44, 0x88, 0x3c, 0x31, 0x99, 0xe3,
	0x70, 0x26, 0x07, 0x66, 0xee, 0x40, 0xae, 0x5a, 0x7d, 0x0a, 0xd7, 0xe3, 0x4e, 0x64, 0xc1, 0xb6,
	0xa6, 0xd7, 0x11, 0x84, 0x51, 0x2e, 0x51, 0x49, 0xc8, 0xe4, 0xfa, 0x55, 0xec, 0xad, 0xf8, 0x66,
	0x3e, 0xb4, 0x2b, 0x1e, 0x3a, 0x3f, 0x56, 0x60, 0xb7, 0x60, 0xeb, 0xb9, 0xc0, 0x22, 0x07, 0xd6,
	0xa7, 0xf5, 0x04, 0x1d, 0xde, 0x4e, 0xdf, 0x8c, 0xfb, 0x4b, 0x71, 0x7a, 0x5a, 0xe7, 0x50, 0xcb,
	0x34, 0x04, 0x99, 0x05, 0x41, 0x73, 0xca, 0x64, 0xdc, 0xfd, 0x2d, 0x26, 0x3f, 0xf3, 0xaa, 0xd6,
	0x11, 0xb4, 0xbb, 0x40, 0x6f, 0x4f, 0xfc, 0x13, 0x18, 0x45, 0xb5, 0xe6, 0xb5, 0xe7, 0x12, 0x1a,
	0x53, 0x22, 0x82, 0x0e, 0x0a, 0x42, 0x16, 0xb5, 0xc8, 0x38, 0x5c, 0x06, 0xd3, 0xd9, 0xaf, 0x60,
	0x63, 0xe6, 0x58, 0x50, 0x21, 0x67, 0x05, 0x67, 0x68, 0x1c, 0x2d, 0x07, 0xea, 0x1a, 0x43, 0x68,
	0xcd, 0xdf, 0x09, 0x7a, 0x50, 0xa4, 0x02, 0xc5, 0x27, 0x68, 0x3c, 0xbc, 0x15, 0x56, 0x15, 0xbb,
	0x5a, 0x93, 0x14, 0x3f, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xcc, 0x4c, 0xa1, 0xae, 0x07,
	0x00, 0x00,
}
