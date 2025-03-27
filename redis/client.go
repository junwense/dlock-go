package redis

import (
	rlock "github.com/meoying/dlock-go/internal/redis"
	"github.com/redis/go-redis/v9"
)

func NewClient(rdb redis.Cmdable) *rlock.Client {
	return rlock.NewClient(rdb)
}
