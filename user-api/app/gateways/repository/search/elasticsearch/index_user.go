package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"user-api/app/domain/entities"
)

// IndexUser is a function to index user in elasticsearch
func (r *UserElasticRepository) UserIndex(ctx context.Context, name string, user *entities.User) error {
	const op = "UserElasticRepository.IndexUser"

	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	resp, err := r.ES.Index(
		name,
		strings.NewReader(string(userJSON)),
		r.ES.Index.WithDocumentID(user.PublicID.String()),
		r.ES.Index.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer resp.Body.Close()

	return nil

}
