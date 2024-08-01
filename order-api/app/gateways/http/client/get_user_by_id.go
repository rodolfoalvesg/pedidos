package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c *Client) GetUserByID(ctx context.Context, userID uuid.UUID) error {
	const (
		op  = "Client.User.GetUserByID"
		raw = "api/v1/user-api/users/%s"
	)

	path := fmt.Sprintf(raw, userID)

	req, err := c.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	resp, err := c.DoRequest(req)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("failed to get user by id")
		err = fmt.Errorf("%w: status code %s", err, resp.Status)

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}
