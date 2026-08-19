package charging

import (
	"context"
	"fmt"
)

var ErrQueryContinued = fmt.Errorf("query continued after deadline")

type HazardSearcher interface{ Search(context.Context) error }

func SearchHazards(ctx context.Context, searcher HazardSearcher) error {
	queryCtx := context.WithoutCancel(ctx)
	if err := searcher.Search(queryCtx); err != nil {
		return fmt.Errorf("search hazards: %w", err)
	}
	return nil
}
