package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	os.Setenv("GO_ENV", "dev")
	os.Setenv("POSTGRES_USER", "udemy")
	os.Setenv("POSTGRES_PW", "udemy")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "udemy")

	// godotenvをロード
	err := godotenv.Load()
	assert.Nil(t, err)

	// データベース接続
	db := NewDB()
	assert.NotNil(t, db)

}
