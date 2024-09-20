package gormx

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var once sync.Once

var singleInstance *gorm.DB

type cfg struct {
	Address      string `json:"address"`
	Port         uint   `json:"port"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	Database     string `json:"database"`
	CharSet      string `json:"charset"`
	MaxIdleConns int    `json:"max-idle-conns"`
	MaxOpenConns int    `json:"max-open-conns"`
	LogMode      string `json:"log-mode"`
}

func GetInstance(opt ...cfg) *gorm.DB {
	if singleInstance == nil && len(opt) > 0 {
		once.Do(
			func() {
				singleInstance = initMysql(&opt[0])
			})
	}

	return singleInstance
}

// InitMysql 初始化数据库并产生数据库全局变量
func initMysql(con *cfg) *gorm.DB {
	//  默认的级别，会打印find找不到模型时的sql语句。
	logMode := logger.Info
	// if con.LogMode == cfg.LogReleaseMode {
	// 	logMode = logger.Silent
	// }

	// Silent 就不会。
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logMode,     // gorm日志模式：silent 可选 Silent，Error，Warn，Info
			Colorful:      true,        // Disable color  true/false
		},
	)

	//"root:123456@tcp(127.0.0.1:3306)/gorm_class?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
	dnsString := getDNS(*con)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dnsString,
		DefaultStringSize:         256,   // string 类型字段的默认长度                                                                            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置                                                                            // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		PrepareStmt: true, // open sql preparestmt

		// Logger: logger.Default.LogMode(logMode), // gorm日志模式：Silent / Info(控制台显示日志) / Error / Warn
		Logger: newLogger,

		DisableForeignKeyConstraintWhenMigrating: true, // 外键约束

		// SkipDefaultTransaction: true, // 禁用默认事务（提高运行速度）
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		// os.Exit(1)
		panic(fmt.Sprintf("gorm.Open err:%v \n", err))
	}

	sqlDB, _ := db.DB()

	// Enable Logger, show detailed log
	sqlDB.SetMaxIdleConns(con.MaxIdleConns)    // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(con.MaxOpenConns)    // 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 设置了连接可复用的最大时间。

	// release环境禁止使用自动迁移
	// if con.LogMode != cfg.LogReleaseMode {
	// 	// 关闭自动迁移，有需要本地调试打开
	// 	// autoMigrate()
	// }

	return db
}

func getDNS(opt cfg) string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		opt.UserName,
		opt.Password,
		opt.Address,
		opt.Port,
		opt.Database,
		opt.CharSet,
	)
}
