package rpc

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"
	"github.com/multiverse-vcs/go-multiverse/git"
)

// ImportArgs contains the args.
type ImportArgs struct {
	// Name is the name of the repo.
	Name string
	// Type is the repo type.
	Type string
	// URL is the repo address.
	URL string
	// Dir is the repo directory.
	Dir string
}

// ImportReply contains the reply.
type ImportReply struct{}

// Import creates a new repo from an existing git repo.
func (s *Service) Import(args *ImportArgs, reply *ImportReply) error {
	ctx := context.Background()

	if args.Type != "git" {
		return errors.New("unsupported repo format")
	}

	if args.Name == "" {
		return errors.New("name cannot be empty")
	}

	if _, err := s.store.GetCid(args.Name); err == nil {
		return errors.New("repo with name already exists")
	}

	var id cid.Cid
	var err error

	switch {
	case args.URL != "":
		id, err = git.ImportFromURL(ctx, s.client, args.Name, args.URL)
	case args.Dir != "":
		id, err = git.ImportFromFS(ctx, s.client, args.Name, args.Dir)
	default:
		return errors.New("url or dir is required")
	}

	if err != nil {
		return err
	}

	return s.store.PutCid(args.Name, id)
}