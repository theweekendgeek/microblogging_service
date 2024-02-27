package business

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DataSourceMock struct{}

func (ms DataSourceMock) GetProfileIDs() []string {
	return []string{"12345", "67890"}
}

func TestGetProfileIDs(t *testing.T) {
	bla := DataSourceMock{}
	userIDs := GetUserIDs(bla)
	expected := []string{"12345", "67890"}

	assert.Equal(t, expected, userIDs)
}
