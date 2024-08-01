package requests

import (
	"net/http"
	"order-api/app/domain/entities"
)

func ParseQueryParams(r *http.Request) *entities.OrderFilter {
	filter := &entities.OrderFilter{}

	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			value := values[0]
			switch key {
			case "user-id":
				filter.UserID = value
			case "order-id":
				filter.PublicID = value
			case "name":
				filter.Description = value
			}
		}
	}

	return filter
}
