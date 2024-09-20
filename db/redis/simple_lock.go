package redis

import (
	"context"
	"time"

	"github.com/spf13/cast"
)

// Lock attempts to acquire a lock with the given key and a time to live.
// flag is a unique identifier for the locking process.
//
// It returns true if the lock is acquired, false otherwise.
func Lock(key string, ttl time.Duration, flag string) (bool, error) {
	lock, err := GetInstance().SetNX(context.TODO(), key, flag, ttl).Result()
	return lock, err
}

// lua script
const script = `if redis.call("GET",KEYS[1]) == ARGV[1]
then
    return redis.call("DEL",KEYS[1])
else
    return 0
end`

// UnLock releases a shared lock
//
// If flag is not empty, it first checks if it's the lock owner before releasing it,
// preventing the lock from being released by an unauthorized process.
func UnLock(key string, flag string) (bool, error) {
	r, err := GetInstance().Eval(context.TODO(), script, []string{key}, flag).Result()
	if err != nil {
		return false, err
	}
	return cast.ToInt(r) == 1, nil
}
