// Code generated by protoc-gen-gogo.
// source: server/pkg/metrics/metrics.proto
// DO NOT EDIT!

/*
Package metrics is a generated protocol buffer package.

It is generated from these files:
	server/pkg/metrics/metrics.proto

It has these top-level messages:
	Metrics
*/
package metrics

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Metrics struct {
	ClusterID        string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	PodID            string `protobuf:"bytes,2,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	Nodes            int64  `protobuf:"varint,3,opt,name=nodes,proto3" json:"nodes,omitempty"`
	Version          string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Repos            int64  `protobuf:"varint,5,opt,name=repos,proto3" json:"repos,omitempty"`
	Commits          int64  `protobuf:"varint,6,opt,name=commits,proto3" json:"commits,omitempty"`
	Files            int64  `protobuf:"varint,7,opt,name=files,proto3" json:"files,omitempty"`
	Bytes            int64  `protobuf:"varint,8,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Jobs             int64  `protobuf:"varint,9,opt,name=jobs,proto3" json:"jobs,omitempty"`
	Pipelines        int64  `protobuf:"varint,10,opt,name=pipelines,proto3" json:"pipelines,omitempty"`
	ArchivedCommits  int64  `protobuf:"varint,11,opt,name=archived_commits,json=archivedCommits,proto3" json:"archived_commits,omitempty"`
	CancelledCommits int64  `protobuf:"varint,12,opt,name=cancelled_commits,json=cancelledCommits,proto3" json:"cancelled_commits,omitempty"`
}

func (m *Metrics) Reset()                    { *m = Metrics{} }
func (m *Metrics) String() string            { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()               {}
func (*Metrics) Descriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{0} }

func (m *Metrics) GetClusterID() string {
	if m != nil {
		return m.ClusterID
	}
	return ""
}

func (m *Metrics) GetPodID() string {
	if m != nil {
		return m.PodID
	}
	return ""
}

func (m *Metrics) GetNodes() int64 {
	if m != nil {
		return m.Nodes
	}
	return 0
}

func (m *Metrics) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Metrics) GetRepos() int64 {
	if m != nil {
		return m.Repos
	}
	return 0
}

func (m *Metrics) GetCommits() int64 {
	if m != nil {
		return m.Commits
	}
	return 0
}

func (m *Metrics) GetFiles() int64 {
	if m != nil {
		return m.Files
	}
	return 0
}

func (m *Metrics) GetBytes() int64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *Metrics) GetJobs() int64 {
	if m != nil {
		return m.Jobs
	}
	return 0
}

func (m *Metrics) GetPipelines() int64 {
	if m != nil {
		return m.Pipelines
	}
	return 0
}

func (m *Metrics) GetArchivedCommits() int64 {
	if m != nil {
		return m.ArchivedCommits
	}
	return 0
}

func (m *Metrics) GetCancelledCommits() int64 {
	if m != nil {
		return m.CancelledCommits
	}
	return 0
}

func init() {
	proto.RegisterType((*Metrics)(nil), "Metrics")
}

func init() { proto.RegisterFile("server/pkg/metrics/metrics.proto", fileDescriptorMetrics) }

var fileDescriptorMetrics = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0x80, 0xe9, 0xcf, 0xb6, 0x66, 0x54, 0xac, 0xa1, 0x87, 0x20, 0x42, 0x17, 0x4f, 0x15, 0xc5,
	0x1e, 0x7c, 0x03, 0xdb, 0xcb, 0x1e, 0x04, 0xd9, 0x17, 0x28, 0x6d, 0x32, 0xd6, 0xe8, 0x76, 0x27,
	0x24, 0x6b, 0xc1, 0x97, 0xed, 0xa1, 0x2f, 0xe0, 0x2b, 0x48, 0x26, 0x5d, 0xf5, 0xb4, 0xf3, 0x7d,
	0xf3, 0x2d, 0x0c, 0x81, 0x3c, 0xa0, 0xdf, 0xa1, 0x9f, 0xb9, 0x8f, 0xcd, 0x6c, 0x8b, 0x8d, 0xb7,
	0x3a, 0xb4, 0xdf, 0x07, 0xe7, 0xa9, 0xa1, 0xab, 0xf1, 0x86, 0x36, 0xc4, 0xe3, 0x2c, 0x4e, 0xc9,
	0xde, 0x7c, 0x77, 0x61, 0xf8, 0x9c, 0x3a, 0x79, 0x0f, 0xa0, 0xab, 0xcf, 0xd0, 0xa0, 0x5f, 0x5a,
	0xa3, 0x3a, 0x79, 0x67, 0x2a, 0x9e, 0xce, 0x0f, 0xfb, 0x89, 0x98, 0x27, 0x5b, 0x2c, 0x4a, 0x71,
	0x0c, 0x0a, 0x23, 0x73, 0x18, 0x38, 0x32, 0xb1, 0xec, 0x72, 0x29, 0x0e, 0xfb, 0x49, 0xf6, 0x42,
	0xa6, 0x58, 0x94, 0x99, 0x23, 0x53, 0x18, 0x39, 0x86, 0xac, 0x26, 0x83, 0x41, 0xf5, 0xf2, 0xce,
	0xb4, 0x57, 0x26, 0x90, 0x0a, 0x86, 0x3b, 0xf4, 0xc1, 0x52, 0xad, 0xfa, 0xf1, 0xc7, 0xb2, 0xc5,
	0xd8, 0x7b, 0x74, 0x14, 0x54, 0x96, 0x7a, 0x86, 0xd8, 0x6b, 0xda, 0x6e, 0x6d, 0x13, 0xd4, 0x80,
	0x7d, 0x8b, 0xb1, 0x7f, 0xb5, 0x15, 0x06, 0x35, 0x4c, 0x3d, 0x43, 0xb4, 0xeb, 0xaf, 0x06, 0x83,
	0x3a, 0x49, 0x96, 0x41, 0x4a, 0xe8, 0xbf, 0xd3, 0x3a, 0x28, 0xc1, 0x92, 0x67, 0x79, 0x0d, 0xc2,
	0x59, 0x87, 0x95, 0xad, 0x31, 0x28, 0xe0, 0xc5, 0x9f, 0x90, 0xb7, 0x30, 0x5a, 0x79, 0xfd, 0x66,
	0x77, 0x68, 0x96, 0xed, 0x01, 0xa7, 0x1c, 0x5d, 0xb4, 0x7e, 0x7e, 0x3c, 0xe4, 0x0e, 0x2e, 0xf5,
	0xaa, 0xd6, 0x58, 0x55, 0xff, 0xda, 0x33, 0x6e, 0x47, 0xbf, 0x8b, 0x63, 0xbc, 0x1e, 0xf0, 0xc3,
	0x3f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x46, 0x59, 0x6c, 0x73, 0xb2, 0x01, 0x00, 0x00,
}
