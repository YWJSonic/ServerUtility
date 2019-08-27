package cacheinfo

import (
	"time"

	"gitlab.com/ServerUtility/redigo/redis"
)

// Get redis command
func Get(cachePool *redis.Pool, ket string) (interface{}, error) {
	Conn := cachePool.Get()
	defer Conn.Close()

	return Conn.Do("GET", ket)
}

// Set redis command
func Set(cachePool *redis.Pool, args []interface{}, d time.Duration) (interface{}, error) {
	Conn := cachePool.Get()
	defer Conn.Close()

	if d != 0 {
		args = append(args, "PX", int(d/time.Millisecond))
	}
	return Conn.Do("SET", args...)
}

// Del redis command
func Del(cachePool *redis.Pool, key string) error {
	Conn := cachePool.Get()
	defer Conn.Close()

	_, err := Conn.Do("DEL", key)
	return err
}

// RunExist redis option
func RunExist(cachePool *redis.Pool, key string) (bool, error) {
	Conn := cachePool.Get()
	defer Conn.Close()

	if _, err := Conn.Do("EXISTS", key); err != redis.ErrNil {
		return false, err
	}
	return false, nil

}

// RunDelete redis interface option
func RunDelete(cachePool *redis.Pool, key string) error {
	Conn := cachePool.Get()
	defer Conn.Close()

	_, err := Conn.Do("DEL", key)
	return err
}

// RunFlush redis interface option
func RunFlush(cachePool *redis.Pool) error {
	Conn := cachePool.Get()
	defer Conn.Close()
	_, err := Conn.Do("FLUSHDB")
	return err
}

// RunAdd redis option
func RunAdd(cachePool *redis.Pool, key string, v interface{}, d time.Duration) error {

	args := []interface{}{key, v, "NX"}

	_, err := Set(cachePool, args, d)
	return err
}

// RunReplace redis option
func RunReplace(cachePool *redis.Pool, key string, v interface{}, d time.Duration) error {
	args := []interface{}{key, v, "XX"}

	v, err := Set(cachePool, args, d)
	if err != nil {
		return err
	}
	if v == nil {
		return redis.ErrNil
	}
	return nil
}

// RunSet redis interface option
func RunSet(cachePool *redis.Pool, key string, v interface{}, d time.Duration) error {
	args := []interface{}{key, v}

	_, err := Set(cachePool, args, d)
	return err
}

// GetBool redis interface option
func GetBool(cachePool *redis.Pool, key string) (bool, error) {
	return redis.Bool(Get(cachePool, key))
}

// GetBytes redis interface option
func GetBytes(cachePool *redis.Pool, key string) ([]byte, error) {
	return redis.Bytes(Get(cachePool, key))
}

// GetInt redis interface option
func GetInt(cachePool *redis.Pool, key string) (int, error) {
	return redis.Int(Get(cachePool, key))
}

// GetInt64 redis interface option
func GetInt64(cachePool *redis.Pool, key string) (int64, error) {
	return redis.Int64(Get(cachePool, key))
}

// GetFloat64 redis interface option
func GetFloat64(cachePool *redis.Pool, key string) (float64, error) {
	return redis.Float64(Get(cachePool, key))
}

// GetUint64 redis interface option
func GetUint64(cachePool *redis.Pool, key string) (uint64, error) {
	return redis.Uint64(Get(cachePool, key))
}

// GetString redis interface option
func GetString(cachePool *redis.Pool, key string) (string, error) {
	return redis.String(Get(cachePool, key))
}
