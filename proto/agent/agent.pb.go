/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.6.1
// source: proto/agent/agent.proto

package agent

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	client "sigs.k8s.io/apiserver-network-proxy/konnectivity-client/proto/client"
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

var File_proto_agent_agent_proto protoreflect.FileDescriptor

var file_proto_agent_agent_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x6b, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x31, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x07, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x07, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x73,
	0x69, 0x67, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_agent_agent_proto_goTypes = []interface{}{
	(*client.Packet)(nil), // 0: Packet
}
var file_proto_agent_agent_proto_depIdxs = []int32{
	0, // 0: AgentService.Connect:input_type -> Packet
	0, // 1: AgentService.Connect:output_type -> Packet
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_agent_agent_proto_init() }
func file_proto_agent_agent_proto_init() {
	if File_proto_agent_agent_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_agent_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_agent_agent_proto_goTypes,
		DependencyIndexes: file_proto_agent_agent_proto_depIdxs,
	}.Build()
	File_proto_agent_agent_proto = out.File
	file_proto_agent_agent_proto_rawDesc = nil
	file_proto_agent_agent_proto_goTypes = nil
	file_proto_agent_agent_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentServiceClient interface {
	// Agent Identifier?
	Connect(ctx context.Context, opts ...grpc.CallOption) (AgentService_ConnectClient, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (AgentService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AgentService_serviceDesc.Streams[0], "/AgentService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceConnectClient{stream}
	return x, nil
}

type AgentService_ConnectClient interface {
	Send(*client.Packet) error
	Recv() (*client.Packet, error)
	grpc.ClientStream
}

type agentServiceConnectClient struct {
	grpc.ClientStream
}

func (x *agentServiceConnectClient) Send(m *client.Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentServiceConnectClient) Recv() (*client.Packet, error) {
	m := new(client.Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServiceServer is the server API for AgentService service.
type AgentServiceServer interface {
	// Agent Identifier?
	Connect(AgentService_ConnectServer) error
}

// UnimplementedAgentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (*UnimplementedAgentServiceServer) Connect(AgentService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterAgentServiceServer(s *grpc.Server, srv AgentServiceServer) {
	s.RegisterService(&_AgentService_serviceDesc, srv)
}

func _AgentService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServiceServer).Connect(&agentServiceConnectServer{stream})
}

type AgentService_ConnectServer interface {
	Send(*client.Packet) error
	Recv() (*client.Packet, error)
	grpc.ServerStream
}

type agentServiceConnectServer struct {
	grpc.ServerStream
}

func (x *agentServiceConnectServer) Send(m *client.Packet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentServiceConnectServer) Recv() (*client.Packet, error) {
	m := new(client.Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _AgentService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/agent/agent.proto",
}
