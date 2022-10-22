package twitter

import (
	"doescher.ninja/twitter-service/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildQueryParams(t *testing.T) {
	assert.Equal(t, "", buildQueryParams(QueryOptions{}))
	assert.Equal(t, "?max_results=10", buildQueryParams(QueryOptions{Max: 10}))
	assert.Equal(t, "?max_results=20", buildQueryParams(QueryOptions{Max: 20}))
	assert.Equal(t, "?since_id=1234567890", buildQueryParams(QueryOptions{SinceID: "1234567890"}))
	assert.Equal(t, "?pagination_token=7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g", buildQueryParams(QueryOptions{PaginationToken: "7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g"}))
	assert.Equal(t, "?max_results=10&since_id=1234567890", buildQueryParams(QueryOptions{Max: 10, SinceID: "1234567890"}))
	assert.Equal(t, "?max_results=10&pagination_token=7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g", buildQueryParams(QueryOptions{Max: 10, PaginationToken: "7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g"}))
}

func TestBuildUrl(t *testing.T) {
	assert.Equal(t, config.Const().EndpointTimelineByID+"", buildTimelineUrl(QueryOptions{}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?max_results=10", buildTimelineUrl(QueryOptions{Max: 10}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?max_results=20", buildTimelineUrl(QueryOptions{Max: 20}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?since_id=1234567890", buildTimelineUrl(QueryOptions{SinceID: "1234567890"}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?pagination_token=7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g", buildTimelineUrl(QueryOptions{PaginationToken: "7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g"}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?max_results=10&since_id=1234567890", buildTimelineUrl(QueryOptions{Max: 10, SinceID: "1234567890"}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?max_results=10&pagination_token=7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g", buildTimelineUrl(QueryOptions{Max: 10, PaginationToken: "7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g"}))
}
