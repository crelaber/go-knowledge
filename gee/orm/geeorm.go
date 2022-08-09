package orm

import (
	"database/sql"
	"go-knowledge/gee/orm/dialect"
	"go-knowledge/gee/orm/log"
	"go-knowledge/gee/orm/session"
)

//Engine 的逻辑非常简单，最重要的方法是 NewEngine，NewEngine 主要做了两件事。
//连接数据库，返回 *sql.DB。
//调用 db.Ping()，检查数据库是否能够正常连接。
//另外呢，提供了 Engine 提供了 NewSession() 方法，这样可以通过 Engine 实例创建会话，进而与数据库进行交互了。到这一步，整个 GeeORM 的框架雏形已经出来了。
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	dialect, ok := dialect.GetDialect(driver)

	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}

	e = &Engine{
		db:      db,
		dialect: dialect,
	}
	log.Info("connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}

type TxFunc func(session2 *session.Session) (interface{}, error)

// Transaction 事务相关的操作
func (engine *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p)
		} else if err != nil {
			_ = s.Rollback()
		} else {
			err = s.Commit()
		}
	}()
	return f(s)
}
