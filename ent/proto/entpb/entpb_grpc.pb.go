// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: entpb/entpb.proto

package entpb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	Create(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*Video, error)
	Get(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*Video, error)
	Update(ctx context.Context, in *UpdateVideoRequest, opts ...grpc.CallOption) (*Video, error)
	Delete(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *ListVideoRequest, opts ...grpc.CallOption) (*ListVideoResponse, error)
	BatchCreate(ctx context.Context, in *BatchCreateVideosRequest, opts ...grpc.CallOption) (*BatchCreateVideosResponse, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) Create(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*Video, error) {
	out := new(Video)
	err := c.cc.Invoke(ctx, "/entpb.VideoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) Get(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*Video, error) {
	out := new(Video)
	err := c.cc.Invoke(ctx, "/entpb.VideoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) Update(ctx context.Context, in *UpdateVideoRequest, opts ...grpc.CallOption) (*Video, error) {
	out := new(Video)
	err := c.cc.Invoke(ctx, "/entpb.VideoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) Delete(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/entpb.VideoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) List(ctx context.Context, in *ListVideoRequest, opts ...grpc.CallOption) (*ListVideoResponse, error) {
	out := new(ListVideoResponse)
	err := c.cc.Invoke(ctx, "/entpb.VideoService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) BatchCreate(ctx context.Context, in *BatchCreateVideosRequest, opts ...grpc.CallOption) (*BatchCreateVideosResponse, error) {
	out := new(BatchCreateVideosResponse)
	err := c.cc.Invoke(ctx, "/entpb.VideoService/BatchCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations must embed UnimplementedVideoServiceServer
// for forward compatibility
type VideoServiceServer interface {
	Create(context.Context, *CreateVideoRequest) (*Video, error)
	Get(context.Context, *GetVideoRequest) (*Video, error)
	Update(context.Context, *UpdateVideoRequest) (*Video, error)
	Delete(context.Context, *DeleteVideoRequest) (*empty.Empty, error)
	List(context.Context, *ListVideoRequest) (*ListVideoResponse, error)
	BatchCreate(context.Context, *BatchCreateVideosRequest) (*BatchCreateVideosResponse, error)
	mustEmbedUnimplementedVideoServiceServer()
}

// UnimplementedVideoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (UnimplementedVideoServiceServer) Create(context.Context, *CreateVideoRequest) (*Video, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVideoServiceServer) Get(context.Context, *GetVideoRequest) (*Video, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedVideoServiceServer) Update(context.Context, *UpdateVideoRequest) (*Video, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedVideoServiceServer) Delete(context.Context, *DeleteVideoRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedVideoServiceServer) List(context.Context, *ListVideoRequest) (*ListVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedVideoServiceServer) BatchCreate(context.Context, *BatchCreateVideosRequest) (*BatchCreateVideosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreate not implemented")
}
func (UnimplementedVideoServiceServer) mustEmbedUnimplementedVideoServiceServer() {}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.VideoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).Create(ctx, req.(*CreateVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.VideoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).Get(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.VideoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).Update(ctx, req.(*UpdateVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.VideoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).Delete(ctx, req.(*DeleteVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.VideoService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).List(ctx, req.(*ListVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_BatchCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateVideosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).BatchCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.VideoService/BatchCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).BatchCreate(ctx, req.(*BatchCreateVideosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entpb.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _VideoService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _VideoService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VideoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _VideoService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _VideoService_List_Handler,
		},
		{
			MethodName: "BatchCreate",
			Handler:    _VideoService_BatchCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entpb/entpb.proto",
}
