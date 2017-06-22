package repository

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/config"
)

const VPN_ACCOUNT_TTL = 5 * time.Minute
const PREFIX_FORMAT = "vpn_account:%s"

type VpnAccountRepository interface {
	SetAccountCache(username string, password string) (err error)
	GetAccountCache(username string) (password string, err error)
	HasAccountCache(username string) (exist bool, err error)
}

type RedisVpnAccountRepository struct {
	Client *redis.Client
}

func GetDefaultVpnAccountRepository() (repo VpnAccountRepository) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.GetDefaultRedisUri(),
		Password: "",                       // no password set
		DB:       config.REDIS_ACCOUNTS_DB, // use default DB
	})
	return RedisVpnAccountRepository{
		Client: client,
	}
}

func InitVpnAccountRepositoryWithHost(host string) (repo VpnAccountRepository) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.FormatRedisUri(host),
		Password: "",                       // no password set
		DB:       config.REDIS_ACCOUNTS_DB, // use default DB
	})
	return RedisVpnAccountRepository{
		Client: client,
	}
}

func (repo RedisVpnAccountRepository) SetAccountCache(username string, password string) (err error) {
	return repo.Client.Set(formatVpnAccountKey(username), password, VPN_ACCOUNT_TTL).Err()
}

func (repo RedisVpnAccountRepository) GetAccountCache(username string) (password string, err error) {
	return repo.Client.Get(formatVpnAccountKey(username)).Result()
}

func (repo RedisVpnAccountRepository) HasAccountCache(username string) (exist bool, err error) {
	count, err := repo.Client.Exists(formatVpnAccountKey(username)).Result()
	return 1 == count, nil
}

func formatVpnAccountKey(username string) string {
	return fmt.Sprintf(PREFIX_FORMAT, username)
}
