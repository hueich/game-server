// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/table_service_proto/bigtable_table_service.proto
// DO NOT EDIT!

package google_bigtable_admin_table_v1

import proto "github.com/hueich/game-server/Godeps/_workspace/src/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_bigtable_admin_table_v11 "github.com/hueich/game-server/Godeps/_workspace/src/google.golang.org/cloud/bigtable/internal/table_data_proto"
import google_protobuf1 "github.com/hueich/game-server/Godeps/_workspace/src/google.golang.org/cloud/bigtable/internal/empty"

import (
	context "github.com/hueich/game-server/Godeps/_workspace/src/golang.org/x/net/context"
	grpc "github.com/hueich/game-server/Godeps/_workspace/src/google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for BigtableTableService service

type BigtableTableServiceClient interface {
	// Creates a new table, to be served from a specified cluster.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.Table, error)
	// Lists the names of all tables served from a specified cluster.
	ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error)
	// Gets the schema of the specified table, including its column families.
	GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.Table, error)
	// Permanently deletes a specified table and all of its data.
	DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Changes the name of a specified table.
	// Cannot be used to move tables between clusters, zones, or projects.
	RenameTable(ctx context.Context, in *RenameTableRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Creates a new column family within a specified table.
	CreateColumnFamily(ctx context.Context, in *CreateColumnFamilyRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.ColumnFamily, error)
	// Changes the configuration of a specified column family.
	UpdateColumnFamily(ctx context.Context, in *google_bigtable_admin_table_v11.ColumnFamily, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.ColumnFamily, error)
	// Permanently deletes a specified column family and all of its data.
	DeleteColumnFamily(ctx context.Context, in *DeleteColumnFamilyRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type bigtableTableServiceClient struct {
	cc *grpc.ClientConn
}

func NewBigtableTableServiceClient(cc *grpc.ClientConn) BigtableTableServiceClient {
	return &bigtableTableServiceClient{cc}
}

func (c *bigtableTableServiceClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.Table, error) {
	out := new(google_bigtable_admin_table_v11.Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/CreateTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error) {
	out := new(ListTablesResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/ListTables", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.Table, error) {
	out := new(google_bigtable_admin_table_v11.Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/GetTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/DeleteTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) RenameTable(ctx context.Context, in *RenameTableRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/RenameTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) CreateColumnFamily(ctx context.Context, in *CreateColumnFamilyRequest, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.ColumnFamily, error) {
	out := new(google_bigtable_admin_table_v11.ColumnFamily)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/CreateColumnFamily", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) UpdateColumnFamily(ctx context.Context, in *google_bigtable_admin_table_v11.ColumnFamily, opts ...grpc.CallOption) (*google_bigtable_admin_table_v11.ColumnFamily, error) {
	out := new(google_bigtable_admin_table_v11.ColumnFamily)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/UpdateColumnFamily", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) DeleteColumnFamily(ctx context.Context, in *DeleteColumnFamilyRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/DeleteColumnFamily", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BigtableTableService service

type BigtableTableServiceServer interface {
	// Creates a new table, to be served from a specified cluster.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	CreateTable(context.Context, *CreateTableRequest) (*google_bigtable_admin_table_v11.Table, error)
	// Lists the names of all tables served from a specified cluster.
	ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error)
	// Gets the schema of the specified table, including its column families.
	GetTable(context.Context, *GetTableRequest) (*google_bigtable_admin_table_v11.Table, error)
	// Permanently deletes a specified table and all of its data.
	DeleteTable(context.Context, *DeleteTableRequest) (*google_protobuf1.Empty, error)
	// Changes the name of a specified table.
	// Cannot be used to move tables between clusters, zones, or projects.
	RenameTable(context.Context, *RenameTableRequest) (*google_protobuf1.Empty, error)
	// Creates a new column family within a specified table.
	CreateColumnFamily(context.Context, *CreateColumnFamilyRequest) (*google_bigtable_admin_table_v11.ColumnFamily, error)
	// Changes the configuration of a specified column family.
	UpdateColumnFamily(context.Context, *google_bigtable_admin_table_v11.ColumnFamily) (*google_bigtable_admin_table_v11.ColumnFamily, error)
	// Permanently deletes a specified column family and all of its data.
	DeleteColumnFamily(context.Context, *DeleteColumnFamilyRequest) (*google_protobuf1.Empty, error)
}

func RegisterBigtableTableServiceServer(s *grpc.Server, srv BigtableTableServiceServer) {
	s.RegisterService(&_BigtableTableService_serviceDesc, srv)
}

func _BigtableTableService_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableTableServiceServer).CreateTable(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableTableService_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableTableServiceServer).ListTables(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableTableService_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableTableServiceServer).GetTable(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableTableService_DeleteTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableTableServiceServer).DeleteTable(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableTableService_RenameTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RenameTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableTableServiceServer).RenameTable(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableTableService_CreateColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateColumnFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableTableServiceServer).CreateColumnFamily(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableTableService_UpdateColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_bigtable_admin_table_v11.ColumnFamily)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableTableServiceServer).UpdateColumnFamily(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableTableService_DeleteColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteColumnFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableTableServiceServer).DeleteColumnFamily(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _BigtableTableService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.admin.table.v1.BigtableTableService",
	HandlerType: (*BigtableTableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _BigtableTableService_CreateTable_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _BigtableTableService_ListTables_Handler,
		},
		{
			MethodName: "GetTable",
			Handler:    _BigtableTableService_GetTable_Handler,
		},
		{
			MethodName: "DeleteTable",
			Handler:    _BigtableTableService_DeleteTable_Handler,
		},
		{
			MethodName: "RenameTable",
			Handler:    _BigtableTableService_RenameTable_Handler,
		},
		{
			MethodName: "CreateColumnFamily",
			Handler:    _BigtableTableService_CreateColumnFamily_Handler,
		},
		{
			MethodName: "UpdateColumnFamily",
			Handler:    _BigtableTableService_UpdateColumnFamily_Handler,
		},
		{
			MethodName: "DeleteColumnFamily",
			Handler:    _BigtableTableService_DeleteColumnFamily_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x93, 0xbd, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x61, 0xa9, 0x90, 0xbb, 0x59, 0x88, 0x21, 0x03, 0x43, 0x25, 0x36, 0xe4, 0xa8, 0x65,
	0x42, 0x6c, 0x29, 0x1f, 0x0b, 0x43, 0x55, 0xca, 0x02, 0x43, 0xe4, 0x24, 0x87, 0x65, 0xe4, 0x8f,
	0x10, 0x3b, 0x95, 0x3a, 0xf1, 0x77, 0xb3, 0x41, 0xe2, 0x06, 0x02, 0x54, 0x38, 0x1e, 0x58, 0xac,
	0xda, 0x7e, 0xef, 0xfd, 0x7c, 0x77, 0x0d, 0x7a, 0x60, 0x5a, 0x33, 0x01, 0x84, 0x69, 0x41, 0x15,
	0x23, 0xba, 0x62, 0x71, 0x2e, 0x74, 0x5d, 0xc4, 0x19, 0x67, 0x96, 0x66, 0x02, 0x62, 0xae, 0x2c,
	0x54, 0x8a, 0x8a, 0xb8, 0xdd, 0xa6, 0x06, 0xaa, 0x35, 0xcf, 0x21, 0x2d, 0x2b, 0x6d, 0xf5, 0xa7,
	0x2a, 0xfd, 0x76, 0x49, 0xda, 0x4b, 0x7c, 0xbc, 0xcd, 0xee, 0x44, 0x84, 0x16, 0x92, 0x2b, 0xe2,
	0x7e, 0xaf, 0xa7, 0xd1, 0x2a, 0x94, 0x5d, 0x50, 0x4b, 0x77, 0x83, 0x9b, 0x1b, 0x47, 0x8d, 0xf2,
	0xff, 0xa8, 0x28, 0x95, 0x60, 0x0c, 0x65, 0x60, 0xb6, 0x90, 0x8b, 0xe1, 0x10, 0x90, 0xa5, 0xdd,
	0xb8, 0xd5, 0x99, 0x67, 0x6f, 0x23, 0x74, 0x98, 0x6c, 0x75, 0xab, 0x66, 0xb9, 0x73, 0x10, 0xfc,
	0x8c, 0xc6, 0xf3, 0x0a, 0xa8, 0x75, 0xa7, 0x78, 0x46, 0xfe, 0x6e, 0x20, 0xe9, 0x89, 0x97, 0xf0,
	0x52, 0x83, 0xb1, 0xd1, 0x89, 0xcf, 0xd3, 0xaa, 0x27, 0x7b, 0xb8, 0x46, 0xe8, 0x96, 0x1b, 0xdb,
	0x6e, 0x0d, 0x9e, 0xfa, 0x6c, 0x5f, 0xda, 0x8e, 0x34, 0x0b, 0xb1, 0x98, 0x52, 0x2b, 0xd3, 0x60,
	0x0b, 0x74, 0x70, 0x03, 0xee, 0x18, 0xc7, 0xbe, 0x84, 0x4e, 0x19, 0x5c, 0xdc, 0x23, 0x1a, 0x5f,
	0x82, 0x80, 0xc1, 0x8d, 0xec, 0x89, 0x3b, 0xd6, 0x51, 0xe7, 0x69, 0x67, 0x96, 0xd5, 0x4f, 0xe4,
	0xaa, 0x19, 0xa1, 0x0b, 0x5f, 0x82, 0xa2, 0x72, 0x68, 0x78, 0x4f, 0xec, 0x0f, 0x7f, 0x45, 0xd8,
	0x4d, 0x75, 0xae, 0x45, 0x2d, 0xd5, 0x35, 0x95, 0x5c, 0x6c, 0xf0, 0xf9, 0xb0, 0x7f, 0x42, 0xdf,
	0xd3, 0xa1, 0x4e, 0xbd, 0xd6, 0x9e, 0xe9, 0xe3, 0x01, 0x15, 0xc2, 0xf7, 0x65, 0xf1, 0xf3, 0x01,
	0x41, 0x29, 0xc1, 0x4c, 0x8e, 0xb0, 0x9b, 0x40, 0x58, 0xd1, 0xbf, 0x3d, 0xde, 0xfe, 0x26, 0x09,
	0x9a, 0xe4, 0x5a, 0x7a, 0x92, 0x93, 0x68, 0xd7, 0xe7, 0x69, 0x16, 0x4d, 0xd8, 0x62, 0x3f, 0x1b,
	0xb5, 0xa9, 0x67, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x26, 0xcd, 0xc1, 0xb6, 0x3c, 0x05, 0x00,
	0x00,
}
