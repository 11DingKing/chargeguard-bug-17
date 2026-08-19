package charging

import (
	"context"
	"fmt"
)

var ErrQueryContinued = fmt.Errorf("query continued after deadline")

type HazardSearcher interface{ Search(context.Context) error }

func SearchHazards(ctx context.Context, searcher HazardSearcher) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("search cancelled: %w", err)
	}
	if err := searcher.Search(ctx); err != nil {
		return fmt.Errorf("search hazards: %w", err)
	}
	return nil
}
