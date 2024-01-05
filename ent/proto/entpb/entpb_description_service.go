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
	description "github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	periodicdescriptiontemplate "github.com/KasumiMercury/uo-patradb-dogtrot/ent/periodicdescriptiontemplate"
	video "github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
	empty "github.com/golang/protobuf/ptypes/empty"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// DescriptionService implements DescriptionServiceServer
type DescriptionService struct {
	client *ent.Client
	UnimplementedDescriptionServiceServer
}

// NewDescriptionService returns a new DescriptionService
func NewDescriptionService(client *ent.Client) *DescriptionService {
	return &DescriptionService{
		client: client,
	}
}

// toProtoDescription transforms the ent type to the pb type
func toProtoDescription(e *ent.Description) (*Description, error) {
	v := &Description{}
	id := e.ID
	v.Id = id
	raw := e.Raw
	v.Raw = raw
	source_id := e.SourceID
	v.SourceId = source_id
	template_confidence := e.TemplateConfidence
	v.TemplateConfidence = template_confidence
	updated_at := timestamppb.New(e.UpdatedAt)
	v.UpdatedAt = updated_at
	variable := wrapperspb.String(e.Variable)
	v.Variable = variable
	if edg := e.Edges.PeriodicTemplate; edg != nil {
		id := edg.ID
		v.PeriodicTemplate = &PeriodicDescriptionTemplate{
			Id: id,
		}
	}
	if edg := e.Edges.Video; edg != nil {
		id := edg.ID
		v.Video = &Video{
			Id: id,
		}
	}
	return v, nil
}

// toProtoDescriptionList transforms a list of ent type to a list of pb type
func toProtoDescriptionList(e []*ent.Description) ([]*Description, error) {
	var pbList []*Description
	for _, entEntity := range e {
		pbEntity, err := toProtoDescription(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements DescriptionServiceServer.Create
func (svc *DescriptionService) Create(ctx context.Context, req *CreateDescriptionRequest) (*Description, error) {
	description := req.GetDescription()
	m, err := svc.createBuilder(description)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoDescription(res)
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

// Get implements DescriptionServiceServer.Get
func (svc *DescriptionService) Get(ctx context.Context, req *GetDescriptionRequest) (*Description, error) {
	var (
		err error
		get *ent.Description
	)
	id := req.GetId()
	switch req.GetView() {
	case GetDescriptionRequest_VIEW_UNSPECIFIED, GetDescriptionRequest_BASIC:
		get, err = svc.client.Description.Get(ctx, id)
	case GetDescriptionRequest_WITH_EDGE_IDS:
		get, err = svc.client.Description.Query().
			Where(description.ID(id)).
			WithPeriodicTemplate(func(query *ent.PeriodicDescriptionTemplateQuery) {
				query.Select(periodicdescriptiontemplate.FieldID)
			}).
			WithVideo(func(query *ent.VideoQuery) {
				query.Select(video.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoDescription(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements DescriptionServiceServer.Update
func (svc *DescriptionService) Update(ctx context.Context, req *UpdateDescriptionRequest) (*Description, error) {
	description := req.GetDescription()
	descriptionID := description.GetId()
	m := svc.client.Description.UpdateOneID(descriptionID)
	descriptionRaw := description.GetRaw()
	m.SetRaw(descriptionRaw)
	descriptionSourceID := description.GetSourceId()
	m.SetSourceID(descriptionSourceID)
	descriptionTemplateConfidence := description.GetTemplateConfidence()
	m.SetTemplateConfidence(descriptionTemplateConfidence)
	descriptionUpdatedAt := runtime.ExtractTime(description.GetUpdatedAt())
	m.SetUpdatedAt(descriptionUpdatedAt)
	if description.GetVariable() != nil {
		descriptionVariable := description.GetVariable().GetValue()
		m.SetVariable(descriptionVariable)
	}
	if description.GetPeriodicTemplate() != nil {
		descriptionPeriodicTemplate := description.GetPeriodicTemplate().GetId()
		m.SetPeriodicTemplateID(descriptionPeriodicTemplate)
	}
	if description.GetVideo() != nil {
		descriptionVideo := description.GetVideo().GetId()
		m.SetVideoID(descriptionVideo)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoDescription(res)
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

// Delete implements DescriptionServiceServer.Delete
func (svc *DescriptionService) Delete(ctx context.Context, req *DeleteDescriptionRequest) (*empty.Empty, error) {
	var err error
	id := req.GetId()
	err = svc.client.Description.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements DescriptionServiceServer.List
func (svc *DescriptionService) List(ctx context.Context, req *ListDescriptionRequest) (*ListDescriptionResponse, error) {
	var (
		err      error
		entList  []*ent.Description
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Description.Query().
		Order(ent.Desc(description.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := string(bytes)
		listQuery = listQuery.
			Where(description.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListDescriptionRequest_VIEW_UNSPECIFIED, ListDescriptionRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListDescriptionRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithPeriodicTemplate(func(query *ent.PeriodicDescriptionTemplateQuery) {
				query.Select(periodicdescriptiontemplate.FieldID)
			}).
			WithVideo(func(query *ent.VideoQuery) {
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
		protoList, err := toProtoDescriptionList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListDescriptionResponse{
			DescriptionList: protoList,
			NextPageToken:   nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements DescriptionServiceServer.BatchCreate
func (svc *DescriptionService) BatchCreate(ctx context.Context, req *BatchCreateDescriptionsRequest) (*BatchCreateDescriptionsResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.DescriptionCreate, len(requests))
	for i, req := range requests {
		description := req.GetDescription()
		var err error
		bulk[i], err = svc.createBuilder(description)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Description.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoDescriptionList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateDescriptionsResponse{
			Descriptions: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *DescriptionService) createBuilder(description *Description) (*ent.DescriptionCreate, error) {
	m := svc.client.Description.Create()
	descriptionRaw := description.GetRaw()
	m.SetRaw(descriptionRaw)
	descriptionSourceID := description.GetSourceId()
	m.SetSourceID(descriptionSourceID)
	descriptionTemplateConfidence := description.GetTemplateConfidence()
	m.SetTemplateConfidence(descriptionTemplateConfidence)
	descriptionUpdatedAt := runtime.ExtractTime(description.GetUpdatedAt())
	m.SetUpdatedAt(descriptionUpdatedAt)
	if description.GetVariable() != nil {
		descriptionVariable := description.GetVariable().GetValue()
		m.SetVariable(descriptionVariable)
	}
	if description.GetPeriodicTemplate() != nil {
		descriptionPeriodicTemplate := description.GetPeriodicTemplate().GetId()
		m.SetPeriodicTemplateID(descriptionPeriodicTemplate)
	}
	if description.GetVideo() != nil {
		descriptionVideo := description.GetVideo().GetId()
		m.SetVideoID(descriptionVideo)
	}
	return m, nil
}
