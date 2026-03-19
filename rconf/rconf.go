package rconf

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var refreshLock sync.Mutex
var refreshTimer *time.Timer

////go:embed application.yaml
//var configBytes []byte

func InitConfig[T any](configBytes []byte, tPtr *T) (vv *viper.Viper, err error) {
	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewBuffer(configBytes)); err != nil {
		panic(fmt.Errorf("viper.ReadConfig error: %w", err))
	}
	extPath := v.GetString("app.ext-config-path")
	extFileName := v.GetString("app.ext-config-file")

	slog.Info(fmt.Sprintf("load ext config path: %s, fileName: %s", extPath, extFileName))

	v.SetConfigName(extFileName)
	v.AddConfigPath(v.GetString(extPath))

	// 合并配置（外部配置会覆盖嵌入配置）
	if err := v.MergeInConfig(); err != nil {
		if _, ok := errors.AsType[viper.ConfigFileNotFoundError](err); !ok {
			panic(fmt.Errorf("failed to read external config: %w", err))
		}
		slog.Warn("cannot find the external config file.")
	}

	// 支持环境变量覆盖
	v.AutomaticEnv()
	v.SetEnvPrefix("app_")

	er := v.Unmarshal(tPtr)
	if er != nil {
		panic(er)
	}
	slog.Info(fmt.Sprintf("load config ok, conf: %+v", *tPtr))

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		refreshLock.TryLock()
		defer refreshLock.Unlock()

		if refreshTimer != nil {
			refreshTimer.Stop()
		}
		refreshTimer = time.AfterFunc(5*time.Second, func() {
			slog.Info(fmt.Sprintf("config file changed: %s", e.Name))
			er = v.Unmarshal(tPtr)
			if er != nil {
				slog.Error(fmt.Errorf("reload config fail, err: %w", er).Error())
			} else {
				slog.Info(fmt.Sprintf("reload config ok, conf: %+v", *tPtr))
			}
		})
	})
	return v, nil
}
