// Code generated by protoc-gen-go-psm. DO NOT EDIT.

package github_spb

import (
	context "context"
	fmt "fmt"

	j5reflect "github.com/pentops/j5/lib/j5reflect"
	j5schema "github.com/pentops/j5/lib/j5schema"
	psm "github.com/pentops/j5/lib/psm"
	sqrlx "github.com/pentops/sqrlx.go/sqrlx"
)

// State Query Service for %sRepo
// QuerySet is the query set for the Repo service.

type RepoPSMQuerySet = psm.StateQuerySet

func NewRepoPSMQuerySet(
	smSpec psm.QuerySpec,
	options psm.StateQueryOptions,
) (*RepoPSMQuerySet, error) {
	return psm.BuildStateQuerySet(smSpec, options)
}

type RepoPSMQuerySpec = psm.QuerySpec

func DefaultRepoPSMQuerySpec(tableSpec psm.QueryTableSpec) RepoPSMQuerySpec {
	return psm.QuerySpec{
		GetMethod: &j5schema.MethodSchema{
			Request:  j5schema.MustObjectSchema((&GetRepoRequest{}).ProtoReflect().Descriptor()),
			Response: j5schema.MustObjectSchema((&GetRepoResponse{}).ProtoReflect().Descriptor()),
		},
		ListMethod: &j5schema.MethodSchema{
			Request:  j5schema.MustObjectSchema((&ListReposRequest{}).ProtoReflect().Descriptor()),
			Response: j5schema.MustObjectSchema((&ListReposResponse{}).ProtoReflect().Descriptor()),
		},
		ListEventsMethod: &j5schema.MethodSchema{
			Request:  j5schema.MustObjectSchema((&ListRepoEventsRequest{}).ProtoReflect().Descriptor()),
			Response: j5schema.MustObjectSchema((&ListRepoEventsResponse{}).ProtoReflect().Descriptor()),
		},
		QueryTableSpec: tableSpec,
		ListRequestFilter: func(reqReflect j5reflect.Object) (map[string]interface{}, error) {
			req, ok := reqReflect.Interface().(*ListReposRequest)
			if !ok {
				return nil, fmt.Errorf("expected *ListReposRequest but got %T", req)
			}
			filter := map[string]interface{}{}
			return filter, nil
		},
		ListEventsRequestFilter: func(reqReflect j5reflect.Object) (map[string]interface{}, error) {
			req, ok := reqReflect.Interface().(*ListRepoEventsRequest)
			if !ok {
				return nil, fmt.Errorf("expected *ListRepoEventsRequest but got %T", req)
			}
			filter := map[string]interface{}{}
			filter["owner"] = req.Owner
			filter["name"] = req.Name
			return filter, nil
		},
	}
}

type RepoQueryServiceImpl struct {
	db       sqrlx.Transactor
	querySet *RepoPSMQuerySet
	UnsafeRepoQueryServiceServer
}

var _ RepoQueryServiceServer = &RepoQueryServiceImpl{}

func NewRepoQueryServiceImpl(db sqrlx.Transactor, querySet *RepoPSMQuerySet) *RepoQueryServiceImpl {
	return &RepoQueryServiceImpl{
		db:       db,
		querySet: querySet,
	}
}

func (s *RepoQueryServiceImpl) GetRepo(ctx context.Context, req *GetRepoRequest) (*GetRepoResponse, error) {
	resObject := &GetRepoResponse{}
	err := s.querySet.Get(ctx, s.db, req.J5Object(), resObject.J5Object())
	if err != nil {
		return nil, err
	}
	return resObject, nil
}

func (s *RepoQueryServiceImpl) ListRepos(ctx context.Context, req *ListReposRequest) (*ListReposResponse, error) {
	resObject := &ListReposResponse{}
	err := s.querySet.List(ctx, s.db, req.J5Object(), resObject.J5Object())
	if err != nil {
		return nil, err
	}
	return resObject, nil
}

func (s *RepoQueryServiceImpl) ListRepoEvents(ctx context.Context, req *ListRepoEventsRequest) (*ListRepoEventsResponse, error) {
	resObject := &ListRepoEventsResponse{}
	err := s.querySet.ListEvents(ctx, s.db, req.J5Object(), resObject.J5Object())
	if err != nil {
		return nil, err
	}
	return resObject, nil
}
