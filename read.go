package memgogo

import (
	"context"

	"github.com/warrior-labs/gomemgo/memgopb"
)

// Read a value from the memgo service.
// {key} is the key of the value to read.
func (c *memgoClient) Read(key string) ([]byte, error) {
	req := &memgopb.ReadRequest{
		Key: key,
	}
	resp, err := c.c.Read(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}
