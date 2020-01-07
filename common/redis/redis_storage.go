package redis

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gsession"
	"github.com/gogf/gf/os/gtimer"
	"time"
)

type RedisStorage struct {
	client        redis.UniversalClient // 兼容单节点和集群模式的Redis连接
	prefix        string                // key前缀
	updatingIdMap *gmap.StrIntMap       // 更新TTL缓存
}

// 初始化缓存管理器的方法
func NewStorageRedis(redis redis.UniversalClient, prefix ...string) *RedisStorage {
	if redis == nil {
		panic("连接信息为nil,无法初始化存储管理器!")
		return nil
	}
	s := &RedisStorage{
		client:        redis,
		updatingIdMap: gmap.NewStrIntMap(true),
	}
	if len(prefix) > 0 && prefix[0] != "" {
		s.prefix = prefix[0]
	}
	// 批量更新TTL
	gtimer.AddSingleton(gsession.DefaultStorageRedisLoopInterval, func() {
		glog.Debug("redis -> 更新SESSION TTL开始!")
		var id string
		var err error
		var ttlSeconds int
		for {
			if id, ttlSeconds = s.updatingIdMap.Pop(); id == "" {
				break
			} else {
				if err = s.doUpdateTTL(id, ttlSeconds); err != nil {
					glog.Error(err)
				}
			}
		}
		glog.Debug("redis -> 更新SESSION TTL结束!")
	})
	return s
}

// 使用JSON形式管理SESSION无需实现该方法
func (s *RedisStorage) New(ttl time.Duration) (id string) {
	return ""
}

// 使用JSON形式管理SESSION无需实现该方法
func (s *RedisStorage) Get(id string, key string) interface{} {
	return nil
}

// 使用JSON形式管理SESSION无需实现该方法
func (s *RedisStorage) GetMap(id string) map[string]interface{} {
	return nil
}

// 使用JSON形式管理SESSION无需实现该方法
func (s *RedisStorage) GetSize(id string) int {
	return -1
}

// 使用JSON形式管理SESSION无需实现该方法
func (s *RedisStorage) Set(id string, key string, value interface{}, ttl time.Duration) error {
	return gsession.ErrorDisabled
}

// 使用JSON形式管理SESSION无需实现该方法
func (s *RedisStorage) SetMap(id string, data map[string]interface{}, ttl time.Duration) error {
	return gsession.ErrorDisabled
}

// 使用JSON形式管理SESSION无需实现该方法
func (s *RedisStorage) Remove(id string, key string) error {
	return gsession.ErrorDisabled
}

// 使用JSON形式管理SESSION无需实现该方法
func (s *RedisStorage) RemoveAll(id string) error {
	return gsession.ErrorDisabled
}

// 使用JSON的形式管理SESSION的获取方法
func (s *RedisStorage) GetSession(id string, ttl time.Duration, data *gmap.StrAnyMap) (*gmap.StrAnyMap, error) {
	glog.Debugf("redis -> 读取: %s, %v", id, ttl)
	cmd := s.client.Get(s.key(id))
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	content, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}
	if len(content) == 0 {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(content, &m); err != nil {
		return nil, err
	}
	if m == nil {
		return nil, nil
	}
	if data == nil {
		return gmap.NewStrAnyMapFrom(m, true), nil
	} else {
		data.Replace(m)
	}
	return data, nil
}

// 使用JSON的形式管理SESSION
func (s *RedisStorage) SetSession(id string, data *gmap.StrAnyMap, ttl time.Duration) error {
	glog.Debugf("redis -> 写入: %s, %v, %v", id, data, ttl)
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cmd := s.client.Do("SETEX", s.key(id), ttl.Seconds(), content)
	return cmd.Err()
}

// 使用延时批量更新TTL
func (s *RedisStorage) UpdateTTL(id string, ttl time.Duration) error {
	if ttl >= gsession.DefaultStorageRedisLoopInterval {
		s.updatingIdMap.Set(id, int(ttl.Seconds()))
	}
	return nil
}

// 实际使用的更新TTL的方法
func (s *RedisStorage) doUpdateTTL(id string, ttlSeconds int) error {
	glog.Debugf("redis -> 更新TTL: %s,%d", id, ttlSeconds)
	cmd := s.client.Do("EXPIRE", s.key(id), ttlSeconds)
	return cmd.Err()
}

// key前缀策略
func (s *RedisStorage) key(id string) string {
	return s.prefix + id
}
