package memgogo

import (
	"context"

	"github.com/warrior-labs/gomemgo/memgopb"
)

// Delete a value from the memgo service.
// {key} is the key of the value to delete.
func (c *memgoClient) Delete(key string) error {
	req := &memgopb.DeleteRequest{
		Key: key,
	}
	_, err := c.c.Delete(context.Background(), req)
	return err
}
