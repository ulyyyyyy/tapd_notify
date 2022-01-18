package redis

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/ulyyyyyy/tapd_notify/configs"
	"github.com/ulyyyyyy/tapd_notify/internal/config"
	"testing"
)

func initRedis(t *testing.T) {
	fmt.Printf("[active] embed: %s\n", configs.Active)

	if err := config.Load(); err != nil {
		require.Nil(t, err)
	}

	err := Initialize()
	require.Nil(t, err)
}

func TestSet(t *testing.T) {
	initRedis(t)

	err := Set("TITAN:JOB:3", 3)
	require.Nil(t, err)
}

func TestHashSet(t *testing.T) {
	initRedis(t)

	var err error
	err = HashSet("TAPD:HOOK:TEST", map[string]string{"s2": "running"})
	require.Nil(t, err)
	err = HashSet("TAPD:HOOK:TEST", map[string]string{"s1": "dead"})
	require.Nil(t, err)
}

func TestDel(t *testing.T) {
	initRedis(t)

	keys, err := Keys("TAPD:HOOK:TEST")
	require.Nil(t, err)

	t.Log(len(keys))
	t.Log(keys)

	return

	err = Del("TAPD:HOOK:TEST")
	require.Nil(t, err)
}
