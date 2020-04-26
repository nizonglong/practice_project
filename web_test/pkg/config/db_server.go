package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DBServer 表示DB服务器配置
type DBServer struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DBName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

// ConnectString 表示连接数据库的字符串
func (m DBServer) ConnectString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		m.Host, m.Port, m.User, m.Password, m.DBName)
}

// NewGormDB 初始化gormdb连接
func (m DBServer) NewGormDB(openConnection int) (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", m.ConnectString())
	if err != nil {
		return
	}
	// 设置最大链接数
	db.DB().SetMaxOpenConns(openConnection)
	return
}

// NewPostgresDB 初始化postgresdb
func (m DBServer) NewPostgresDB(idleConnection int) (db *sqlx.DB, err error) {
	// 链接数据库
	db, err = sqlx.Open("postgres", m.ConnectString())
	if err != nil {
		return
	}

	// 设置最大空闲连接数
	db.SetMaxIdleConns(idleConnection)
	return
}
