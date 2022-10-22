package twitter

import (
	"doescher.ninja/twitter-service/config"
	"fmt"
	"strings"
)

type QueryOptions struct {
	Max             int
	SinceID         string
	PaginationToken string
}

func buildQueryParams(options QueryOptions) string {
	var queryParams []string
	if options.Max != 0 {
		queryParams = append(queryParams, "max_results="+fmt.Sprintf("%d", options.Max))
	}
	if options.SinceID != "" {
		queryParams = append(queryParams, "since_id="+options.SinceID)
	}
	if options.PaginationToken != "" {
		queryParams = append(queryParams, "pagination_token="+options.PaginationToken)
	}
	queryString := strings.Join(queryParams, "&")
	if len(queryString) > 0 {
		queryString = "?" + queryString
	}
	return queryString
}

func buildTimelineUrl(options QueryOptions) string {
	return config.Const().EndpointTimelineByID + buildQueryParams(options)
}
