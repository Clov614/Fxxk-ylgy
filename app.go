package main

import (
	"FxxkYlgy/config"
	_ "FxxkYlgy/config"
	_ "FxxkYlgy/log"
	"fmt"
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

type RetJson struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	Data    int    `json:"data"`
}

var (
	retJson    RetJson
	costTime   int
	cycleCount int
	finishApi  = "https://cat-match.easygame2021.com/sheep/v1/game/game_over?rank_score=1&rank_state=%s&rank_time=%s&rank_role=1&skin=1"
)

func main() {
	log.Info("------羊了个羊无限通关脚本------")
	readConfig := config.ReadConfig()
	if readConfig.Token == "" {
		log.Error("未配置token")
		os.Exit(-1)
	}

	if readConfig.CycleCount <= 0 {
		cycleCount = 1
	}

	RO := &grequests.RequestOptions{
		Headers: map[string]string{
			"Host":       "cat-match.easygame2021.com",
			"User-Agent": config.UserAgent,
			"t":          readConfig.Token,
		},
	}

	cycleCount = readConfig.CycleCount
	if readConfig.CycleCount > 20 {
		cycleCount = 20
	}

	for i := 1; i <= cycleCount; i++ {
		log.Info(fmt.Sprintf("第%d遍游玩", i))
		if readConfig.CostTime <= 0 {
			costTime = GenerateRandInt(1, 3600)
		}
		finishUrl := fmt.Sprintf("https://cat-match.easygame2021.com/sheep/v1/game/game_over?rank_score=1&rank_state=%d&rank_time=%d&rank_role=1&skin=1", cycleCount, costTime)
		response, err := grequests.Get(finishUrl, RO)
		if err != nil {
			log.Error("网络错误，本次跳过")
		}

		_ = response.JSON(retJson)
		if retJson.ErrCode != 0 {
			log.Error(fmt.Sprintf("服务器请求错误，本次跳过，错误代码:%d ,错误信息:%s", retJson.ErrCode, retJson.ErrMsg))
		} else {
			log.Info(fmt.Sprintf("第%d次游玩完毕", i))
		}
	}
	fmt.Scanf("f")
}

func GenerateRandInt(min, max int) int {
	rand.Seed(time.Now().Unix()) //随机种子
	return rand.Intn(max-min) + min
}
