package repository

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/config"
)

const BACKUP_SNAPSHOT_TTL = 0
const BACKUP_LATEST_KEY = "backup:latest"

type BackupSnapshotRepository interface {
	SetLatestSnapshot(content string) (err error)
	GetLatestSnapshot() (content string, err error)
	
	SetSnapshot(snapshotTime time.Time, content string) (err error)
	GetSnapshot(snapshotTime time.Time) (content string, err error)
	
	HasSnapshot(snapshotTime time.Time) (exist bool, err error)
	ListSnapshot() (keys []string, err error)
}

type RedisBackupSnapshotRepository struct {
	Client *redis.Client
}

func GetDefaultRedisBackupSnapshotRepository() (repo RedisBackupSnapshotRepository) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.GetDefaultRedisUri(),
		Password: "",                               // no password set
		DB:       config.REDIS_BACKUP_SNAPSHOTS_DB, // use default DB
	})
	return RedisBackupSnapshotRepository{
		Client: client,
	}
}

func (repo *RedisBackupSnapshotRepository) SetLatestSnapshot(content string) (err error) {
	return repo.Client.Set(BACKUP_LATEST_KEY, content, BACKUP_SNAPSHOT_TTL).Err()
}

func (repo *RedisBackupSnapshotRepository) GetLatestSnapshot() (content string, err error) {
	return repo.Client.Get(BACKUP_LATEST_KEY).Result()
}

func (repo *RedisBackupSnapshotRepository) SetSnapshot(snapshotTime time.Time, content string) (err error) {
	return repo.Client.Set(formatBackupSnapshotKey(snapshotTime), content, BACKUP_SNAPSHOT_TTL).Err()
}

func (repo *RedisBackupSnapshotRepository) GetSnapshot(snapshotTime time.Time) (content string, err error) {
	return repo.Client.Get(formatBackupSnapshotKey(snapshotTime)).Result()
}

func (repo *RedisBackupSnapshotRepository) HasSnapshot(snapshotTime time.Time) (exist bool, err error) {
	count, err := repo.Client.Exists(formatBackupSnapshotKey(snapshotTime)).Result()
	return 1 == count, nil
}

func (repo *RedisBackupSnapshotRepository) ListSnapshot() (keys []string, err error) {
	return repo.Client.Keys("backup:*").Result()
}

func formatBackupSnapshotKey(t time.Time) string {
	return fmt.Sprintf("backup:history:%s", t.Format("20060102_150405"))
}