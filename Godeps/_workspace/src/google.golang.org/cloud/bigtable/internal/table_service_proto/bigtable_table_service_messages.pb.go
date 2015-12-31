// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/table_service_proto/bigtable_table_service_messages.proto
// DO NOT EDIT!

/*
Package google_bigtable_admin_table_v1 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/cloud/bigtable/internal/table_service_proto/bigtable_table_service_messages.proto
	google.golang.org/cloud/bigtable/internal/table_service_proto/bigtable_table_service.proto

It has these top-level messages:
	CreateTableRequest
	ListTablesRequest
	ListTablesResponse
	GetTableRequest
	DeleteTableRequest
	RenameTableRequest
	CreateColumnFamilyRequest
	DeleteColumnFamilyRequest
*/
package google_bigtable_admin_table_v1

import proto "github.com/hueich/game-server/Godeps/_workspace/src/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_bigtable_admin_table_v11 "github.com/hueich/game-server/Godeps/_workspace/src/google.golang.org/cloud/bigtable/internal/table_data_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateTableRequest struct {
	// The unique name of the cluster in which to create the new table.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The name by which the new table should be referred to within the cluster,
	// e.g. "foobar" rather than "<cluster_name>/tables/foobar".
	TableId string `protobuf:"bytes,2,opt,name=table_id" json:"table_id,omitempty"`
	// The Table to create. The `name` field of the Table and all of its
	// ColumnFamilies must be left blank, and will be populated in the response.
	Table *google_bigtable_admin_table_v11.Table `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	// The optional list of row keys that will be used to initially split the
	// table into several tablets (Tablets are similar to HBase regions).
	// Given two split keys, "s1" and "s2", three tablets will be created,
	// spanning the key ranges: [, s1), [s1, s2), [s2, ).
	//
	// Example:
	//  * Row keys := ["a", "apple", "custom", "customer_1", "customer_2",
	//                 "other", "zz"]
	//  * initial_split_keys := ["apple", "customer_1", "customer_2", "other"]
	//  * Key assignment:
	//    - Tablet 1 [, apple)                => {"a"}.
	//    - Tablet 2 [apple, customer_1)      => {"apple", "custom"}.
	//    - Tablet 3 [customer_1, customer_2) => {"customer_1"}.
	//    - Tablet 4 [customer_2, other)      => {"customer_2"}.
	//    - Tablet 5 [other, )                => {"other", "zz"}.
	InitialSplitKeys []string `protobuf:"bytes,4,rep,name=initial_split_keys" json:"initial_split_keys,omitempty"`
}

func (m *CreateTableRequest) Reset()                    { *m = CreateTableRequest{} }
func (m *CreateTableRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTableRequest) ProtoMessage()               {}
func (*CreateTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateTableRequest) GetTable() *google_bigtable_admin_table_v11.Table {
	if m != nil {
		return m.Table
	}
	return nil
}

type ListTablesRequest struct {
	// The unique name of the cluster for which tables should be listed.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ListTablesRequest) Reset()                    { *m = ListTablesRequest{} }
func (m *ListTablesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTablesRequest) ProtoMessage()               {}
func (*ListTablesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListTablesResponse struct {
	// The tables present in the requested cluster.
	// At present, only the names of the tables are populated.
	Tables []*google_bigtable_admin_table_v11.Table `protobuf:"bytes,1,rep,name=tables" json:"tables,omitempty"`
}

func (m *ListTablesResponse) Reset()                    { *m = ListTablesResponse{} }
func (m *ListTablesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTablesResponse) ProtoMessage()               {}
func (*ListTablesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListTablesResponse) GetTables() []*google_bigtable_admin_table_v11.Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

type GetTableRequest struct {
	// The unique name of the requested table.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetTableRequest) Reset()                    { *m = GetTableRequest{} }
func (m *GetTableRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTableRequest) ProtoMessage()               {}
func (*GetTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DeleteTableRequest struct {
	// The unique name of the table to be deleted.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteTableRequest) Reset()                    { *m = DeleteTableRequest{} }
func (m *DeleteTableRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTableRequest) ProtoMessage()               {}
func (*DeleteTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type RenameTableRequest struct {
	// The current unique name of the table.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The new name by which the table should be referred to within its containing
	// cluster, e.g. "foobar" rather than "<cluster_name>/tables/foobar".
	NewId string `protobuf:"bytes,2,opt,name=new_id" json:"new_id,omitempty"`
}

func (m *RenameTableRequest) Reset()                    { *m = RenameTableRequest{} }
func (m *RenameTableRequest) String() string            { return proto.CompactTextString(m) }
func (*RenameTableRequest) ProtoMessage()               {}
func (*RenameTableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type CreateColumnFamilyRequest struct {
	// The unique name of the table in which to create the new column family.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The name by which the new column family should be referred to within the
	// table, e.g. "foobar" rather than "<table_name>/columnFamilies/foobar".
	ColumnFamilyId string `protobuf:"bytes,2,opt,name=column_family_id" json:"column_family_id,omitempty"`
	// The column family to create. The `name` field must be left blank.
	ColumnFamily *google_bigtable_admin_table_v11.ColumnFamily `protobuf:"bytes,3,opt,name=column_family" json:"column_family,omitempty"`
}

func (m *CreateColumnFamilyRequest) Reset()                    { *m = CreateColumnFamilyRequest{} }
func (m *CreateColumnFamilyRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateColumnFamilyRequest) ProtoMessage()               {}
func (*CreateColumnFamilyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateColumnFamilyRequest) GetColumnFamily() *google_bigtable_admin_table_v11.ColumnFamily {
	if m != nil {
		return m.ColumnFamily
	}
	return nil
}

type DeleteColumnFamilyRequest struct {
	// The unique name of the column family to be deleted.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteColumnFamilyRequest) Reset()                    { *m = DeleteColumnFamilyRequest{} }
func (m *DeleteColumnFamilyRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteColumnFamilyRequest) ProtoMessage()               {}
func (*DeleteColumnFamilyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*CreateTableRequest)(nil), "google.bigtable.admin.table.v1.CreateTableRequest")
	proto.RegisterType((*ListTablesRequest)(nil), "google.bigtable.admin.table.v1.ListTablesRequest")
	proto.RegisterType((*ListTablesResponse)(nil), "google.bigtable.admin.table.v1.ListTablesResponse")
	proto.RegisterType((*GetTableRequest)(nil), "google.bigtable.admin.table.v1.GetTableRequest")
	proto.RegisterType((*DeleteTableRequest)(nil), "google.bigtable.admin.table.v1.DeleteTableRequest")
	proto.RegisterType((*RenameTableRequest)(nil), "google.bigtable.admin.table.v1.RenameTableRequest")
	proto.RegisterType((*CreateColumnFamilyRequest)(nil), "google.bigtable.admin.table.v1.CreateColumnFamilyRequest")
	proto.RegisterType((*DeleteColumnFamilyRequest)(nil), "google.bigtable.admin.table.v1.DeleteColumnFamilyRequest")
}

var fileDescriptor0 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x5d, 0x4f, 0x2a, 0x31,
	0x10, 0xcd, 0x5e, 0xb8, 0xe4, 0x32, 0xd7, 0x0f, 0xec, 0xd3, 0xc2, 0x83, 0x42, 0x13, 0x13, 0x4c,
	0xcc, 0x6e, 0x44, 0xfd, 0x03, 0x60, 0x34, 0x46, 0x4d, 0x08, 0xf2, 0xbe, 0x29, 0xbb, 0xe3, 0xa6,
	0xb1, 0xdb, 0xe2, 0xb6, 0x60, 0xf8, 0x03, 0xc6, 0x9f, 0x2d, 0xdb, 0x5d, 0x11, 0x12, 0x04, 0x7d,
	0xeb, 0xf4, 0x9c, 0x76, 0xce, 0x9c, 0x39, 0x10, 0xc6, 0x4a, 0xc5, 0x02, 0xbd, 0x58, 0x09, 0x26,
	0x63, 0x4f, 0xa5, 0xb1, 0x1f, 0x0a, 0x35, 0x89, 0xfc, 0x11, 0x8f, 0x0d, 0x1b, 0x09, 0xf4, 0xb9,
	0x34, 0x98, 0x4a, 0x26, 0x7c, 0x5b, 0x06, 0x1a, 0xd3, 0x29, 0x0f, 0x31, 0x18, 0xa7, 0xca, 0xa8,
	0x05, 0x2b, 0x58, 0x05, 0x13, 0xd4, 0x9a, 0xc5, 0xa8, 0x3d, 0xcb, 0x22, 0x87, 0x45, 0x93, 0x4f,
	0xb6, 0xc7, 0xa2, 0x84, 0x4b, 0x2f, 0x3f, 0x4f, 0xcf, 0x1a, 0xc3, 0xdf, 0x8a, 0x88, 0x98, 0x61,
	0xeb, 0x15, 0x64, 0x48, 0xde, 0x95, 0xbe, 0x39, 0x40, 0x7a, 0x29, 0x32, 0x83, 0xc3, 0x0c, 0x1a,
	0xe0, 0xcb, 0x04, 0xb5, 0x21, 0x3b, 0x50, 0x96, 0x2c, 0x41, 0xd7, 0x69, 0x3a, 0xed, 0x2a, 0xa9,
	0xc1, 0xbf, 0xfc, 0x21, 0x8f, 0xdc, 0x3f, 0xf6, 0xe6, 0x02, 0xfe, 0xda, 0x1b, 0xb7, 0x34, 0x2f,
	0xff, 0x77, 0x8e, 0xbd, 0xcd, 0xe2, 0x3d, 0xfb, 0x39, 0x69, 0x00, 0xe1, 0x92, 0x1b, 0xce, 0x44,
	0xa0, 0xc7, 0x82, 0x9b, 0xe0, 0x19, 0x67, 0xda, 0x2d, 0x37, 0x4b, 0xed, 0x2a, 0x6d, 0xc1, 0xc1,
	0x3d, 0xd7, 0xc6, 0x12, 0xf5, 0x5a, 0x19, 0xf4, 0x0e, 0xc8, 0x32, 0x45, 0x8f, 0x95, 0xd4, 0x48,
	0x2e, 0xa1, 0x62, 0xdb, 0xe8, 0x39, 0xab, 0xf4, 0x63, 0x2d, 0xf4, 0x08, 0xf6, 0x6f, 0xd0, 0x7c,
	0x3f, 0x34, 0xa5, 0x40, 0xae, 0x50, 0xe0, 0x26, 0x63, 0x68, 0x07, 0xc8, 0x00, 0xb3, 0x7a, 0x83,
	0x79, 0x7b, 0x50, 0x91, 0xf8, 0xba, 0xb0, 0x8e, 0xbe, 0x3b, 0x50, 0xcf, 0x1d, 0xef, 0x29, 0x31,
	0x49, 0xe4, 0x35, 0x4b, 0xb8, 0x98, 0xad, 0x7f, 0xeb, 0x42, 0x2d, 0xb4, 0xa4, 0xe0, 0xc9, 0xb2,
	0xbe, 0x16, 0xd0, 0x83, 0xdd, 0x15, 0xa4, 0x58, 0xc4, 0xe9, 0xb6, 0xe1, 0x97, 0x7b, 0xd2, 0x13,
	0xa8, 0xe7, 0x23, 0x6e, 0x55, 0xd2, 0xbd, 0x05, 0x1a, 0xaa, 0x64, 0xcb, 0xef, 0xdd, 0x56, 0xb7,
	0x00, 0xac, 0x1f, 0x8f, 0x79, 0xd0, 0x1f, 0x8a, 0x9c, 0xf7, 0xb3, 0xc0, 0xf5, 0x9d, 0x51, 0xc5,
	0x26, 0xef, 0xfc, 0x23, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x1c, 0x6f, 0x09, 0x56, 0x03, 0x00, 0x00,
}
