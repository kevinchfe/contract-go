package cmd

import (
	"contract/pkg/console"
	"contract/pkg/redis"
	"github.com/spf13/cobra"
	"time"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes th Go Playground,but running at our application context",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
	// 存进redis
	redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	// 从redis取出
	console.Success(redis.Redis.Get("hello"))
}
