package elastic

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *UserElasticRepository) DeleteUser(ctx context.Context, publicID uuid.UUID) error {
	const op = "UserElasticRepository.DeleteUser"

	_, err := r.ES.Delete(
		"users",
		publicID.String(),
		r.ES.Delete.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
