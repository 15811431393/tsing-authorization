package updater

import (
	"errors"
	"strings"

	"local/global"
	"local/updater/jwt_hs256"
	"local/updater/jwt_rs256"

	"github.com/rs/zerolog/log"
)

// 构建更新器实例
func Build(name, config string) (global.UpdaterInstance, error) {
	name = strings.ToUpper(name)
	switch name {
	case "JWT_HS256":
		instance, err := jwt_hs256.New(config)
		if err != nil {
			log.Err(err).Caller().Send()
			return nil, err
		}
		return instance, nil
	case "JWT_RS256":
		instance, err := jwt_rs256.New(config)
		if err != nil {
			log.Err(err).Caller().Send()
			return nil, err
		}
		return instance, nil
	}
	return nil, errors.New("不支持的更新器类型")
}
