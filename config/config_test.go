package config

import (
	"testing"
	"os"
	"path/filepath"
	"fmt"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestAppConfig_Load(t *testing.T) {


	fh := createFileForTest(t, `TELEGRAM_BOT_TOKEN=telegram_bot_token
SLACK_AUTH_TOKEN=slack_auth_token
POSTGRE_HOST = "127.0.0.1"
POSTGRE_PORT = "5432"
POSTGRE_DB_NAME = "notify"
POSTGRE_USER = "root"
POSTGRE_PASSWORD = 123
NG_DEBUG = true
NG_PORT="8080"`)

	path := fh.Name()
	defer func() {
		fh.Close()
		os.Remove(path)
	}()

	AppConf.Load(path)

	assert.Equal(t, "telegram_bot_token", AppConf.Telegram.BotToken)
	assert.Equal(t, "slack_auth_token", AppConf.SlackConf.AuthToken)
	assert.Equal(t, "root", AppConf.Postgre.User)
	assert.Equal(t, ":8080", AppConf.Port)
}


func createFileForTest(t *testing.T, s string) *os.File {
	data := []byte(s)
	path := filepath.Join(os.TempDir(), fmt.Sprintf("file.%d", time.Now().UnixNano()))
	fh, err := os.Create(path)
	if err != nil {
		t.Error(err)
	}
	_, err = fh.Write(data)
	if err != nil {
		t.Error(err)
	}

	return fh
}