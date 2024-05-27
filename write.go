package gomemgo

import (
	"context"
	"time"

	"github.com/warrior-labs/gomemgo/memgopb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Write a value to the memgo service.
// {key} is the key of the value to write.
// {value} is the value to write.
// {expiresAt} is the time at which the value should expire (optional).
func (c *memgoClient) Write(key string, value []byte, expiresAt *time.Time) error {
	req := &memgopb.WriteRequest{
		Key:       key,
		Value:     value,
		ExpiresAt: nil,
	}
	if expiresAt != nil {
		req.ExpiresAt = timestamppb.New(*expiresAt)
	}
	_, err := c.c.Write(context.Background(), req)
	return err
}
