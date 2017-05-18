package main

import (
	"time"
	"fmt"
	"gitlab.ecoworkinc.com/Subspace/softetherlib/softether"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/repository"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/utils"
)

func main() {

	s := softether.SoftEther{IP: "54.227.209.241", Password: "subspace", Hub: "subspace"}
	repo := repository.MysqlProfileSnapshotRepository{
		Host:         "54.227.209.241",
		Account:      "subspace",
		Password:     "subspace",
		DatabaseName: "subspace",
	}

	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for t := range ticker.C {
			fmt.Println("Get session at", t)
			userList, code := s.GetUserList()

			if 0 != code {
				time.Sleep(5 * time.Second)
				continue
			}

			fmt.Println("User List Start---------------")
			profileSnapshots := make([]*model.ProfileSnapshot, 0)
			for _, rawData := range userList {
				userName := rawData["User Name"]

				if userDetail, code := s.GetUserInfo(userName); 0 == code {
					if profile, err := utils.ParseUserGet(s.Hub, userDetail); nil == err {
						profileSnapshots = append(profileSnapshots, profile)

						//Demo usage
						if err := repo.Insert(profile); nil != err {
							fmt.Println(err)
						}
					}
				}
			}

			//Demo usage will duplicate entry
			repo.InsertBatch(profileSnapshots)

		}
	}()

	time.Sleep(30 * time.Second)
}