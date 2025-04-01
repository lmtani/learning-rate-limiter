package limiter

import (
	"errors"
	"testing"
	"time"

	"github.com/lmtani/learning-rate-limiter/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRateLimiterStore struct {
	mock.Mock
}

func (m *MockRateLimiterStore) Increment(key string, expire time.Duration, limit int) (int, error) {
	args := m.Called(key, expire)
	return args.Int(0), args.Error(1)
}

func (m *MockRateLimiterStore) Reset(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func TestRateLimiter_ShallPass(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		limitType string
		count     int
		limit     int
		tokenMap  entity.TokenMap
		storeErr  error
		want      bool
	}{
		{
			name:      "IP below limit",
			key:       "192.168.1.1",
			limitType: "ip",
			count:     5,
			limit:     10,
			want:      true,
		},
		{
			name:      "IP at limit",
			key:       "192.168.1.1",
			limitType: "ip",
			count:     10,
			limit:     10,
			want:      true,
		},
		{
			name:      "IP above limit",
			key:       "192.168.1.1",
			limitType: "ip",
			count:     11,
			limit:     10,
			storeErr:  errors.New("limit reached"),
			want:      false,
		},
		{
			name:      "API key below limit without custom limit",
			key:       "api-key-123",
			limitType: "api_key",
			count:     5,
			limit:     10,
			want:      true,
		},
		{
			name:      "API key with custom limit below",
			key:       "api-key-123",
			limitType: "api_key",
			count:     15,
			limit:     10,
			tokenMap:  entity.TokenMap{"api-key-123": 20},
			want:      true,
		},
		{
			name:      "API key with custom limit at limit",
			key:       "api-key-123",
			limitType: "api_key",
			count:     20,
			limit:     10,
			tokenMap:  entity.TokenMap{"api-key-123": 20},
			want:      true,
		},
		{
			name:      "Invalid limit type",
			key:       "some-key",
			limitType: "invalid",
			want:      false,
		},
		{
			name:      "Store error",
			key:       "192.168.1.1",
			limitType: "ip",
			storeErr:  errors.New("storage error"),
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := &MockRateLimiterStore{}

			// Configurar mock para o método Increment
			if tt.limitType == "ip" || tt.limitType == "api_key" {
				mockStore.On("Increment", tt.key, mock.Anything, mock.Anything).Return(tt.count, tt.storeErr)
			}

			rl := NewRateLimiter(tt.limit, time.Minute, mockStore, tt.tokenMap)
			got := rl.ShallPass(tt.key, tt.limitType)

			assert.Equal(t, tt.want, got)
			mockStore.AssertExpectations(t)
		})
	}
}

func TestNewRateLimiter(t *testing.T) {
	mockStore := &MockRateLimiterStore{}
	tokenMap := entity.TokenMap{"key1": 100}

	rl := NewRateLimiter(10, time.Minute, mockStore, tokenMap)

	assert.Equal(t, 10, rl.Limit)
	assert.Equal(t, time.Minute, rl.Expire)
	assert.Equal(t, mockStore, rl.Store)
	assert.Equal(t, tokenMap, rl.TokenMap)
}
