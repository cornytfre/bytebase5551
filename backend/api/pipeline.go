package api

import "context"

// Pipeline status
type PipelineStatus string

const (
	Pipeline_Open     PipelineStatus = "OPEN"
	Pipeline_Done     PipelineStatus = "DONE"
	Pipeline_Canceled PipelineStatus = "CANCELED"
)

func (e PipelineStatus) String() string {
	switch e {
	case Pipeline_Open:
		return "OPEN"
	case Pipeline_Done:
		return "DONE"
	case Pipeline_Canceled:
		return "CANCELED"
	}
	return ""
}

type Pipeline struct {
	ID int `jsonapi:"primary,pipeline"`

	// Standard fields
	CreatorId   int   `jsonapi:"attr,creatorId"`
	CreatedTs   int64 `jsonapi:"attr,createdTs"`
	UpdaterId   int   `jsonapi:"attr,updaterId"`
	UpdatedTs   int64 `jsonapi:"attr,updatedTs"`
	WorkspaceId int

	// Related fields
	// StageList []*Stage `jsonapi:"relation,stage"`

	// Domain specific fields
	Name   string         `jsonapi:"attr,name"`
	Status PipelineStatus `jsonapi:"attr,status"`
}

type PipelineCreate struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	CreatorId   int
	WorkspaceId int

	// Domain specific fields
	Name string `jsonapi:"attr,name"`
}

type PipelineFind struct {
	ID *int

	// Standard fields
	WorkspaceId *int
}

type PipelinePatch struct {
	ID int `jsonapi:"primary,pipelinePatch"`

	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	UpdaterId   int
	WorkspaceId int

	// Domain specific fields
	Status *PipelineStatus `jsonapi:"attr,status"`
}

type PipelineService interface {
	CreatePipeline(ctx context.Context, create *PipelineCreate) (*Pipeline, error)
	FindPipeline(ctx context.Context, find *PipelineFind) (*Pipeline, error)
	PatchPipelineByID(ctx context.Context, patch *PipelinePatch) (*Pipeline, error)
}
