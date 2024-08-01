package requests

import (
	"net/http"
	"user-api/app/domain/entities"
)

func ParseQueryParams(r *http.Request) *entities.UserFilter {
	filter := &entities.UserFilter{}

	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			value := values[0]
			switch key {
			case "public_id":
				filter.PublicID = value
			case "name":
				filter.Name = value
			}
		}
	}

	return filter
}
