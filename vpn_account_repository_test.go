package main

import (
	"testing"
	"github.com/go-redis/redis"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/config"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/repository"
	"time"
)

func TestRedisVpnAccountRepository(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.FormatRedisUri(config.TEST_HOST),
		Password: "",                       // no password set
		DB:       config.REDIS_ACCOUNTS_DB, // use default DB
	})

	repo := repository.RedisVpnAccountRepository{
		Client: client,
	}
	account := "1_1497421288354092850"
	password := "q4zRV248IT"
	err := repo.SetAccountCache(account, password)
	if nil != err {
		t.Error("Set account into redis fail:", err.Error())
	}
	accountExist, err := repo.HasAccountCache(account)
	if nil != err {
		t.Error("Check account is in redis fail:", err.Error())
	}
	if !accountExist {
		t.Error("Account should in redis but not exist.")
	}

	passwordInRedis, err := repo.GetAccountCache(account)
	if nil != err {
		t.Error("Get account from redis fail:", err.Error())
	}
	if passwordInRedis != password {
		t.Error("password is not match")
	}

	// Sleep to after TTL + 3 seconds
	time.Sleep(repository.VPN_ACCOUNT_TTL + 3 * time.Second)
	accountExistAfterExpire, err := repo.HasAccountCache(account)
	if nil != err {
		t.Error("Check account is in redis fail:", err.Error())
	}
	if accountExistAfterExpire {
		t.Error("Account should expired but exist.")
	}
}