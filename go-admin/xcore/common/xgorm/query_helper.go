package xgorm

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xcore/common/xtype/xstring"
)

type OperationType string

const (
	EQ      OperationType = "="
	NEQ                   = "!="
	LIKE                  = "like"
	LESS                  = "<"
	LESS_EQ               = "<="
	MORE                  = ">"
	MORE_EQ               = ">="
	BETWEEN               = "between"
)

type QueryType = map[string]OperationType

type IQueryHelper interface {
	// SpliceQueryIfExit 拼接query参数如果存在
	SpliceQueryIfExit(c *gin.Context, queryType QueryType) *gorm.DB
	SetDB(func(db *gorm.DB))
}

type QueryHelper struct {
	DB *gorm.DB
}

func (q *QueryHelper) SetDB(f func(db *gorm.DB)) {
	f(q.DB)
}

func (q *QueryHelper) SpliceQueryIfExit(c *gin.Context, queryType QueryType) *gorm.DB {
	for key, method := range queryType {
		switch method {
		case EQ:
			if query, b := c.GetQuery(key); b {
				q.DB = q.DB.Where(xstring.LowerCamelCaseToSnake(key)+" = ?", query)
			}
			break
		case NEQ:
			if query, b := c.GetQuery(key); b {
				q.DB = q.DB.Where(xstring.LowerCamelCaseToSnake(key)+" <> ?", query)
			}
			break
		case LIKE:
			if query, b := c.GetQuery(key); b {
				q.DB = q.DB.Where(xstring.LowerCamelCaseToSnake(key)+" LIKE ?", "%"+query+"%")
			}
			break
		case LESS:
			if query, b := c.GetQuery(key); b {
				q.DB = q.DB.Where(xstring.LowerCamelCaseToSnake(key)+" < ?", query)
			}
			break
		case LESS_EQ:
			if query, b := c.GetQuery(key); b {
				q.DB = q.DB.Where(xstring.LowerCamelCaseToSnake(key)+" <= ?", query)
			}
			break
		case MORE:
			if query, b := c.GetQuery(key); b {
				q.DB = q.DB.Where(xstring.LowerCamelCaseToSnake(key)+" > ?", query)
			}
			break
		case MORE_EQ:
			if query, b := c.GetQuery(key); b {
				q.DB = q.DB.Where(xstring.LowerCamelCaseToSnake(key)+" >= ?", query)
			}
			break
		case BETWEEN:
			if query, b := c.GetQueryArray(key); b && len(query) == 2 {
				q.DB = q.DB.Where(xstring.LowerCamelCaseToSnake(key)+" BETWEEN ? AND ?", query[0], query[1])
			}
			break
		}

	}
	return q.DB
}

func NewQueryHelper(db *gorm.DB) IQueryHelper {
	return &QueryHelper{
		DB: db,
	}
}
