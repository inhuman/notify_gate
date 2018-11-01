package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestAppConfig_Load(t *testing.T) {

	fh := createFileForTest(t, `TELEGRAM_BOT_TOKEN=telegram_bot_token
SLACK_AUTH_TOKEN=slack_auth_token
DB_HOST = "127.0.0.1"
DB_PORT = "5432"
DB_NAME = "notify"
DB_USER = "root"
DB_PASSWORD = 123
NG_DEBUG = true
NG_UI_PORT="8080"`)

	path := fh.Name()
	defer func() {
		fh.Close()
		os.Remove(path)
	}()

	AppConf.Load(path)

	assert.Equal(t, "telegram_bot_token", AppConf.Senders.Telegram.BotToken)
	assert.Equal(t, "slack_auth_token", AppConf.Senders.Slack.AuthToken)
	assert.Equal(t, "root", AppConf.DB.User)
	assert.Equal(t, "8080", AppConf.Port)
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
