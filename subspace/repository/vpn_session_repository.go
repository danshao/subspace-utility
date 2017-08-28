package repository

/**

The user name format is `<USER ID>_<NANOSECOND TIMESTAMP>` and the SoftEther
subspace name will be formatted looks like `SID-<USER ID>_<NANOSECOND TIMESTAMP>-[L2TP]-<NUMBER LOGIN>`.

The subspace key format is `session:<HUB_NAME>:<USER NAME>` and field is `<SESSION_NAME>`
The subspace key will looks like:
sessions:subspace:1_1492575227928 SID-1_1492575227928-[L2TP]-45

*/
import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
	"time"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/config"
)

type SessionRepository interface {
	PutAllSessionIfNotExist(dataSet []*model.Session) (err error)
	ClearAndPutAllSession(dataSet []*model.Session) (err error)
	GetAllSessions() (results []model.Session, err error)
	GetSessionsByUserId(userId int) (results []model.Session, err error)
	GetSessionsByProfileUserName(userName string) (results []model.Session, err error)
	ClearSession() (err error)
}

type RedisSessionRepository struct {
	Hub    string
	Client *redis.Client
}

func InitSessionRepositoryWithHost(host string) (repo SessionRepository) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.FormatRedisUri(host),
		Password: "",                       // no password set
		DB:       config.REDIS_VPN_SESSIONS_DB, // use default DB
	})
	return RedisSessionRepository{
		Hub: config.DEFAULT_HUB_NAME,
		Client: client,
	}
}

const REDIS_TIME_TO_LIVE = 30 * time.Second

func (repo RedisSessionRepository) PutAllSessionIfNotExist(dataSet []*model.Session) (err error) {
	pipe := repo.Client.TxPipeline()
	for _, value := range dataSet {
		redisKey := fmt.Sprintf("sessions:%s:%s", value.Hub, value.UserNameAuthentication)
		redisField := value.SessionName

		sessionJson, _ := json.Marshal(value)
		jsonString := string(sessionJson)

		//Do NOT replace existing subspace that has detail.
		if e := pipe.HSetNX(redisKey, redisField, jsonString).Err(); nil != e {
			fmt.Println("Cache subspace brief fail name:", value.SessionName, e)
		}
		if e := pipe.Expire(redisKey, REDIS_TIME_TO_LIVE).Err(); nil != e {
			fmt.Println("Set expire fail on:", redisKey, e)
		}
	}
	if _, transactionError := pipe.Exec(); nil != transactionError {
		fmt.Println("Batch update session fail.")
		return transactionError
	}
	return
}

func (repo RedisSessionRepository) ClearAndPutAllSession(dataSet []*model.Session) (err error) {
	pipe := repo.Client.TxPipeline()
	pipe.FlushDB()
	for _, value := range dataSet {
		redisKey := fmt.Sprintf("sessions:%s:%s", value.Hub, value.UserNameAuthentication)
		redisField := value.SessionName

		sessionJson, _ := json.Marshal(value)
		jsonString := string(sessionJson)

		if e := pipe.HSet(redisKey, redisField, jsonString).Err(); nil != e {
			fmt.Println("Cache subspace brief fail name:", value.SessionName, e)
		}
		if e := pipe.Expire(redisKey, REDIS_TIME_TO_LIVE).Err(); nil != e {
			fmt.Println("Set expire fail on:", redisKey, e)
		}
	}
	if _, transactionError := pipe.Exec(); nil != transactionError {
		fmt.Println("Batch update session fail.")
		return transactionError
	}
	return
}

func (repo RedisSessionRepository) GetAllSessions() (results []model.Session, err error) {
	return repo.ScanSessions(fmt.Sprintf("sessions:%s:*", repo.Hub))
}

func (repo RedisSessionRepository) GetSessionsByUserId(userId int) (results []model.Session, err error) {
	return repo.ScanSessions(fmt.Sprintf("sessions:%s:%d_*", repo.Hub, userId))
}

func (repo RedisSessionRepository) GetSessionsByProfileUserName(userName string) (results []model.Session, err error) {
	return repo.GetSessionsByKey(fmt.Sprintf("sessions:%s:%s", repo.Hub, userName))
}

func (repo RedisSessionRepository) ScanSessions(pattern string) (results []model.Session, err error) {
	page, _, err := repo.Client.Scan(0, pattern, 10000).Result()
	if nil != err {
		fmt.Println("Scan sessions from redis fail for pattern:", pattern)
		return results, err
	}

	results = make([]model.Session, 0)
	for _, key := range page {
		if sessionsOfProfile, e := repo.GetSessionsByKey(key); nil == e {
			results = append(results, sessionsOfProfile...)
		}

	}

	return results, nil

}

func (repo RedisSessionRepository) GetSessionsByKey(key string) (results []model.Session, err error) {
	results = make([]model.Session, 0)
	values, err := repo.Client.HVals(key).Result()

	if nil != err {
		return results, err
	}

	for _, value := range values {
		var session model.Session
		if e := json.Unmarshal([]byte(value), &session); nil == e {
			results = append(results, session)
		} else {
			fmt.Println("json unmarshal fail", e)
		}
	}
	return results, nil
}

func (repo RedisSessionRepository) ClearSession() (err error) {
	return repo.Client.FlushDb().Err()
}
