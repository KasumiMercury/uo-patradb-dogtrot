// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	ent "github.com/KasumiMercury/uo-patradb-dogtrot/ent"
	streamschedule "github.com/KasumiMercury/uo-patradb-dogtrot/ent/streamschedule"
	empty "github.com/golang/protobuf/ptypes/empty"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// StreamScheduleService implements StreamScheduleServiceServer
type StreamScheduleService struct {
	client *ent.Client
	UnimplementedStreamScheduleServiceServer
}

// NewStreamScheduleService returns a new StreamScheduleService
func NewStreamScheduleService(client *ent.Client) *StreamScheduleService {
	return &StreamScheduleService{
		client: client,
	}
}

// toProtoStreamSchedule transforms the ent type to the pb type
func toProtoStreamSchedule(e *ent.StreamSchedule) (*StreamSchedule, error) {
	v := &StreamSchedule{}
	created_at := timestamppb.New(e.CreatedAt)
	v.CreatedAt = created_at
	id := e.ID
	v.Id = id
	scheduled_at := timestamppb.New(e.ScheduledAt)
	v.ScheduledAt = scheduled_at
	_Title := e.Title
	v.Title = _Title
	updated_at := timestamppb.New(e.UpdatedAt)
	v.UpdatedAt = updated_at
	return v, nil
}

// toProtoStreamScheduleList transforms a list of ent type to a list of pb type
func toProtoStreamScheduleList(e []*ent.StreamSchedule) ([]*StreamSchedule, error) {
	var pbList []*StreamSchedule
	for _, entEntity := range e {
		pbEntity, err := toProtoStreamSchedule(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements StreamScheduleServiceServer.Create
func (svc *StreamScheduleService) Create(ctx context.Context, req *CreateStreamScheduleRequest) (*StreamSchedule, error) {
	streamschedule := req.GetStreamSchedule()
	m, err := svc.createBuilder(streamschedule)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoStreamSchedule(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements StreamScheduleServiceServer.Get
func (svc *StreamScheduleService) Get(ctx context.Context, req *GetStreamScheduleRequest) (*StreamSchedule, error) {
	var (
		err error
		get *ent.StreamSchedule
	)
	id := req.GetId()
	switch req.GetView() {
	case GetStreamScheduleRequest_VIEW_UNSPECIFIED, GetStreamScheduleRequest_BASIC:
		get, err = svc.client.StreamSchedule.Get(ctx, id)
	case GetStreamScheduleRequest_WITH_EDGE_IDS:
		get, err = svc.client.StreamSchedule.Query().
			Where(streamschedule.ID(id)).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoStreamSchedule(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements StreamScheduleServiceServer.Update
func (svc *StreamScheduleService) Update(ctx context.Context, req *UpdateStreamScheduleRequest) (*StreamSchedule, error) {
	streamschedule := req.GetStreamSchedule()
	streamscheduleID := streamschedule.GetId()
	m := svc.client.StreamSchedule.UpdateOneID(streamscheduleID)
	streamscheduleCreatedAt := runtime.ExtractTime(streamschedule.GetCreatedAt())
	m.SetCreatedAt(streamscheduleCreatedAt)
	streamscheduleScheduledAt := runtime.ExtractTime(streamschedule.GetScheduledAt())
	m.SetScheduledAt(streamscheduleScheduledAt)
	streamscheduleTitle := streamschedule.GetTitle()
	m.SetTitle(streamscheduleTitle)
	streamscheduleUpdatedAt := runtime.ExtractTime(streamschedule.GetUpdatedAt())
	m.SetUpdatedAt(streamscheduleUpdatedAt)

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoStreamSchedule(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements StreamScheduleServiceServer.Delete
func (svc *StreamScheduleService) Delete(ctx context.Context, req *DeleteStreamScheduleRequest) (*empty.Empty, error) {
	var err error
	id := req.GetId()
	err = svc.client.StreamSchedule.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements StreamScheduleServiceServer.List
func (svc *StreamScheduleService) List(ctx context.Context, req *ListStreamScheduleRequest) (*ListStreamScheduleResponse, error) {
	var (
		err      error
		entList  []*ent.StreamSchedule
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.StreamSchedule.Query().
		Order(ent.Desc(streamschedule.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := string(bytes)
		listQuery = listQuery.
			Where(streamschedule.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListStreamScheduleRequest_VIEW_UNSPECIFIED, ListStreamScheduleRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListStreamScheduleRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoStreamScheduleList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListStreamScheduleResponse{
			StreamScheduleList: protoList,
			NextPageToken:      nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements StreamScheduleServiceServer.BatchCreate
func (svc *StreamScheduleService) BatchCreate(ctx context.Context, req *BatchCreateStreamSchedulesRequest) (*BatchCreateStreamSchedulesResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.StreamScheduleCreate, len(requests))
	for i, req := range requests {
		streamschedule := req.GetStreamSchedule()
		var err error
		bulk[i], err = svc.createBuilder(streamschedule)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.StreamSchedule.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoStreamScheduleList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateStreamSchedulesResponse{
			StreamSchedules: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *StreamScheduleService) createBuilder(streamschedule *StreamSchedule) (*ent.StreamScheduleCreate, error) {
	m := svc.client.StreamSchedule.Create()
	streamscheduleCreatedAt := runtime.ExtractTime(streamschedule.GetCreatedAt())
	m.SetCreatedAt(streamscheduleCreatedAt)
	streamscheduleScheduledAt := runtime.ExtractTime(streamschedule.GetScheduledAt())
	m.SetScheduledAt(streamscheduleScheduledAt)
	streamscheduleTitle := streamschedule.GetTitle()
	m.SetTitle(streamscheduleTitle)
	streamscheduleUpdatedAt := runtime.ExtractTime(streamschedule.GetUpdatedAt())
	m.SetUpdatedAt(streamscheduleUpdatedAt)
	return m, nil
}