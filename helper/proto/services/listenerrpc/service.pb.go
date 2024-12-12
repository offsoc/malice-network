// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.20.3
// source: services/listenerrpc/service.proto

package listenerrpc

import (
	clientpb "github.com/chainreactors/malice-network/helper/proto/client/clientpb"
	implantpb "github.com/chainreactors/malice-network/helper/proto/implant/implantpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_listenerrpc_service_proto protoreflect.FileDescriptor

var file_services_listenerrpc_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x72, 0x70,
	0x63, 0x1a, 0x1c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xd1, 0x07, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x50, 0x43, 0x12,
	0x36, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x07, 0x53, 0x79, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x79,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x12, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x35, 0x0a, 0x0f, 0x49, 0x6e, 0x69, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x10, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x10, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x12,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x57,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x19, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2e, 0x43, 0x74, 0x72, 0x6c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x0f,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x37, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x16, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x74, 0x72, 0x6c, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x1a, 0x13, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x74,
	0x72, 0x6c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x0b, 0x53,
	0x74, 0x6f, 0x70, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x74, 0x72, 0x6c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x1a, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x33, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x12, 0x11, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0b, 0x53, 0x70, 0x69, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x53, 0x70, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x16,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x01, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x09, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x11, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x74, 0x72, 0x6c, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x6d, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x68,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_listenerrpc_service_proto_goTypes = []any{
	(*clientpb.RegisterSession)(nil),  // 0: clientpb.RegisterSession
	(*implantpb.SysInfo)(nil),         // 1: modulepb.SysInfo
	(*implantpb.Ping)(nil),            // 2: modulepb.Ping
	(*implantpb.Request)(nil),         // 3: modulepb.Request
	(*clientpb.RegisterListener)(nil), // 4: clientpb.RegisterListener
	(*clientpb.Pipeline)(nil),         // 5: clientpb.Pipeline
	(*clientpb.CtrlPipeline)(nil),     // 6: clientpb.CtrlPipeline
	(*clientpb.Listener)(nil),         // 7: clientpb.Listener
	(*clientpb.Website)(nil),          // 8: clientpb.Website
	(*clientpb.Builder)(nil),          // 9: clientpb.Builder
	(*clientpb.SpiteResponse)(nil),    // 10: clientpb.SpiteResponse
	(*clientpb.JobStatus)(nil),        // 11: clientpb.JobStatus
	(*clientpb.Empty)(nil),            // 12: clientpb.Empty
	(*clientpb.WebsiteResponse)(nil),  // 13: clientpb.WebsiteResponse
	(*clientpb.Pipelines)(nil),        // 14: clientpb.Pipelines
	(*clientpb.SpiteRequest)(nil),     // 15: clientpb.SpiteRequest
	(*clientpb.JobCtrl)(nil),          // 16: clientpb.JobCtrl
}
var file_services_listenerrpc_service_proto_depIdxs = []int32{
	0,  // 0: listenerrpc.ListenerRPC.Register:input_type -> clientpb.RegisterSession
	1,  // 1: listenerrpc.ListenerRPC.SysInfo:input_type -> modulepb.SysInfo
	2,  // 2: listenerrpc.ListenerRPC.Checkin:input_type -> modulepb.Ping
	3,  // 3: listenerrpc.ListenerRPC.InitBindSession:input_type -> modulepb.Request
	4,  // 4: listenerrpc.ListenerRPC.RegisterListener:input_type -> clientpb.RegisterListener
	5,  // 5: listenerrpc.ListenerRPC.RegisterPipeline:input_type -> clientpb.Pipeline
	5,  // 6: listenerrpc.ListenerRPC.RegisterWebsite:input_type -> clientpb.Pipeline
	6,  // 7: listenerrpc.ListenerRPC.StartPipeline:input_type -> clientpb.CtrlPipeline
	6,  // 8: listenerrpc.ListenerRPC.StopPipeline:input_type -> clientpb.CtrlPipeline
	7,  // 9: listenerrpc.ListenerRPC.ListPipelines:input_type -> clientpb.Listener
	6,  // 10: listenerrpc.ListenerRPC.StartWebsite:input_type -> clientpb.CtrlPipeline
	6,  // 11: listenerrpc.ListenerRPC.StopWebsite:input_type -> clientpb.CtrlPipeline
	8,  // 12: listenerrpc.ListenerRPC.UploadWebsite:input_type -> clientpb.Website
	7,  // 13: listenerrpc.ListenerRPC.ListWebsites:input_type -> clientpb.Listener
	9,  // 14: listenerrpc.ListenerRPC.GetArtifact:input_type -> clientpb.Builder
	10, // 15: listenerrpc.ListenerRPC.SpiteStream:input_type -> clientpb.SpiteResponse
	11, // 16: listenerrpc.ListenerRPC.JobStream:input_type -> clientpb.JobStatus
	12, // 17: listenerrpc.ListenerRPC.Register:output_type -> clientpb.Empty
	12, // 18: listenerrpc.ListenerRPC.SysInfo:output_type -> clientpb.Empty
	12, // 19: listenerrpc.ListenerRPC.Checkin:output_type -> clientpb.Empty
	12, // 20: listenerrpc.ListenerRPC.InitBindSession:output_type -> clientpb.Empty
	12, // 21: listenerrpc.ListenerRPC.RegisterListener:output_type -> clientpb.Empty
	12, // 22: listenerrpc.ListenerRPC.RegisterPipeline:output_type -> clientpb.Empty
	13, // 23: listenerrpc.ListenerRPC.RegisterWebsite:output_type -> clientpb.WebsiteResponse
	12, // 24: listenerrpc.ListenerRPC.StartPipeline:output_type -> clientpb.Empty
	12, // 25: listenerrpc.ListenerRPC.StopPipeline:output_type -> clientpb.Empty
	14, // 26: listenerrpc.ListenerRPC.ListPipelines:output_type -> clientpb.Pipelines
	12, // 27: listenerrpc.ListenerRPC.StartWebsite:output_type -> clientpb.Empty
	12, // 28: listenerrpc.ListenerRPC.StopWebsite:output_type -> clientpb.Empty
	12, // 29: listenerrpc.ListenerRPC.UploadWebsite:output_type -> clientpb.Empty
	14, // 30: listenerrpc.ListenerRPC.ListWebsites:output_type -> clientpb.Pipelines
	9,  // 31: listenerrpc.ListenerRPC.GetArtifact:output_type -> clientpb.Builder
	15, // 32: listenerrpc.ListenerRPC.SpiteStream:output_type -> clientpb.SpiteRequest
	16, // 33: listenerrpc.ListenerRPC.JobStream:output_type -> clientpb.JobCtrl
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_listenerrpc_service_proto_init() }
func file_services_listenerrpc_service_proto_init() {
	if File_services_listenerrpc_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_listenerrpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_listenerrpc_service_proto_goTypes,
		DependencyIndexes: file_services_listenerrpc_service_proto_depIdxs,
	}.Build()
	File_services_listenerrpc_service_proto = out.File
	file_services_listenerrpc_service_proto_rawDesc = nil
	file_services_listenerrpc_service_proto_goTypes = nil
	file_services_listenerrpc_service_proto_depIdxs = nil
}