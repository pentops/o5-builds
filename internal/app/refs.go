package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/elgris/sqrl"
	"github.com/pentops/j5/lib/j5codec"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_pb"
	"github.com/pentops/sqrlx.go/sqrlx"
)

type RefStore struct {
	db sqrlx.Transactor
}

func NewRefStore(db sqrlx.Transactor) (*RefStore, error) {

	return &RefStore{
		db: db,
	}, nil
}

func (rs *RefStore) GetRepo(ctx context.Context, owner string, name string) (*github_pb.RepoState, error) {
	qq := sq.
		Select("state").
		From("repo").
		Where(sq.Eq{
			"owner": owner,
			"name":  name,
		})

	var stateBytes []byte

	err := rs.db.Transact(ctx, &sqrlx.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
		Retryable: true,
	}, func(ctx context.Context, tx sqrlx.Transaction) error {

		err := tx.SelectRow(ctx, qq).Scan(&stateBytes)
		if err != nil {
			return err
		}
		return nil
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("selecting push targets: %w", err)
	}

	repo := &github_pb.RepoState{}
	if err := j5codec.Global.JSONToProto(stateBytes, repo.ProtoReflect()); err != nil {
		return nil, fmt.Errorf("unmarshalling repo state: %w", err)
	}

	return repo, nil
}
