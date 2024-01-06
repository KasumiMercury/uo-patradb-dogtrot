// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	ent "github.com/KasumiMercury/uo-patradb-dogtrot/ent"
	video "github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
	videotag "github.com/KasumiMercury/uo-patradb-dogtrot/ent/videotag"
	empty "github.com/golang/protobuf/ptypes/empty"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// VideoTagService implements VideoTagServiceServer
type VideoTagService struct {
	client *ent.Client
	UnimplementedVideoTagServiceServer
}

// NewVideoTagService returns a new VideoTagService
func NewVideoTagService(client *ent.Client) *VideoTagService {
	return &VideoTagService{
		client: client,
	}
}

// toProtoVideoTag transforms the ent type to the pb type
func toProtoVideoTag(e *ent.VideoTag) (*VideoTag, error) {
	v := &VideoTag{}
	id := e.ID
	v.Id = id
	series_numbering := int64(e.SeriesNumbering)
	v.SeriesNumbering = series_numbering
	title := e.Title
	v.Title = title
	for _, edg := range e.Edges.Videos {
		id := edg.ID
		v.Videos = append(v.Videos, &Video{
			Id: id,
		})
	}
	return v, nil
}

// toProtoVideoTagList transforms a list of ent type to a list of pb type
func toProtoVideoTagList(e []*ent.VideoTag) ([]*VideoTag, error) {
	var pbList []*VideoTag
	for _, entEntity := range e {
		pbEntity, err := toProtoVideoTag(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements VideoTagServiceServer.Create
func (svc *VideoTagService) Create(ctx context.Context, req *CreateVideoTagRequest) (*VideoTag, error) {
	videotag := req.GetVideoTag()
	m, err := svc.createBuilder(videotag)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoVideoTag(res)
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

// Get implements VideoTagServiceServer.Get
func (svc *VideoTagService) Get(ctx context.Context, req *GetVideoTagRequest) (*VideoTag, error) {
	var (
		err error
		get *ent.VideoTag
	)
	id := req.GetId()
	switch req.GetView() {
	case GetVideoTagRequest_VIEW_UNSPECIFIED, GetVideoTagRequest_BASIC:
		get, err = svc.client.VideoTag.Get(ctx, id)
	case GetVideoTagRequest_WITH_EDGE_IDS:
		get, err = svc.client.VideoTag.Query().
			Where(videotag.ID(id)).
			WithVideos(func(query *ent.VideoQuery) {
				query.Select(video.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoVideoTag(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements VideoTagServiceServer.Update
func (svc *VideoTagService) Update(ctx context.Context, req *UpdateVideoTagRequest) (*VideoTag, error) {
	videotag := req.GetVideoTag()
	videotagID := videotag.GetId()
	m := svc.client.VideoTag.UpdateOneID(videotagID)
	videotagSeriesNumbering := int(videotag.GetSeriesNumbering())
	m.SetSeriesNumbering(videotagSeriesNumbering)
	videotagTitle := videotag.GetTitle()
	m.SetTitle(videotagTitle)
	for _, item := range videotag.GetVideos() {
		videos := item.GetId()
		m.AddVideoIDs(videos)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoVideoTag(res)
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

// Delete implements VideoTagServiceServer.Delete
func (svc *VideoTagService) Delete(ctx context.Context, req *DeleteVideoTagRequest) (*empty.Empty, error) {
	var err error
	id := req.GetId()
	err = svc.client.VideoTag.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements VideoTagServiceServer.List
func (svc *VideoTagService) List(ctx context.Context, req *ListVideoTagRequest) (*ListVideoTagResponse, error) {
	var (
		err      error
		entList  []*ent.VideoTag
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.VideoTag.Query().
		Order(ent.Desc(videotag.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := string(bytes)
		listQuery = listQuery.
			Where(videotag.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListVideoTagRequest_VIEW_UNSPECIFIED, ListVideoTagRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListVideoTagRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithVideos(func(query *ent.VideoQuery) {
				query.Select(video.FieldID)
			}).
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
		protoList, err := toProtoVideoTagList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListVideoTagResponse{
			VideoTagList:  protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements VideoTagServiceServer.BatchCreate
func (svc *VideoTagService) BatchCreate(ctx context.Context, req *BatchCreateVideoTagsRequest) (*BatchCreateVideoTagsResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.VideoTagCreate, len(requests))
	for i, req := range requests {
		videotag := req.GetVideoTag()
		var err error
		bulk[i], err = svc.createBuilder(videotag)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.VideoTag.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoVideoTagList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateVideoTagsResponse{
			VideoTags: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *VideoTagService) createBuilder(videotag *VideoTag) (*ent.VideoTagCreate, error) {
	m := svc.client.VideoTag.Create()
	videotagSeriesNumbering := int(videotag.GetSeriesNumbering())
	m.SetSeriesNumbering(videotagSeriesNumbering)
	videotagTitle := videotag.GetTitle()
	m.SetTitle(videotagTitle)
	for _, item := range videotag.GetVideos() {
		videos := item.GetId()
		m.AddVideoIDs(videos)
	}
	return m, nil
}
