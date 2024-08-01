package elastic

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *OrderElasticRepository) DeleteOrder(ctx context.Context, publicID uuid.UUID) error {
	const op = "OrderElasticRepository.DeleteOrder"

	_, err := r.ES.Delete(
		"orders",
		publicID.String(),
		r.ES.Delete.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
