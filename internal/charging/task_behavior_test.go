package charging

import (
	"context"
	"errors"
	"testing"
)

type deadlineSearcher struct{}

func (deadlineSearcher) Search(ctx context.Context) error {
	if ctx.Err() == nil {
		return ErrQueryContinued
	}
	return ctx.Err()
}
func TestTaskBehavior(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := SearchHazards(ctx, deadlineSearcher{})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err=%v", err)
	}
}
