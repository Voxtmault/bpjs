package interfaces

import "context"

type Example interface {
	SampleGet(ctx context.Context, query any) ([]*any, error)
	SampleCreate(ctx context.Context, createObj any) error
	SampleUpdate(ctx context.Context, updateObj any) error
	SampleDelete(ctx context.Context, deleteObj any) error
	SampleInjectionUsage(ctx context.Context) error
}
