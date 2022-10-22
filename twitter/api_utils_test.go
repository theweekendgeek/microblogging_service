package twitter

import (
	"doescher.ninja/twitter-service/config"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: make this easier to read
func TestBuildUrl(t *testing.T) {
	assert.Equal(t, config.Const().EndpointTimelineByID+"?max_results="+fmt.Sprintf("%d", config.Const().MaxTweets), buildTimelineURL(QueryOptions{}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?max_results=20", buildTimelineURL(QueryOptions{Max: 20}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?max_results="+fmt.Sprintf("%d", config.Const().MaxTweets)+"&since_id=1234567890", buildTimelineURL(QueryOptions{SinceID: "1234567890"}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?max_results="+fmt.Sprintf("%d", config.Const().MaxTweets)+"&pagination_token=7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g", buildTimelineURL(QueryOptions{PaginationToken: "7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g"}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?max_results=10&since_id=1234567890", buildTimelineURL(QueryOptions{Max: 10, SinceID: "1234567890"}))
	assert.Equal(t, config.Const().EndpointTimelineByID+"?max_results=10&pagination_token=7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g", buildTimelineURL(QueryOptions{Max: 10, PaginationToken: "7140dibdnow9c7btw423i5yvevms09yqahhym8jsoz48g"}))
}
