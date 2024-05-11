package xgorm

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	SpliceIfExit(c *gin.Context, queryType QueryType) *gorm.DB
}

type QueryHelper struct {
	db *gorm.DB
}

func (q *QueryHelper) SpliceIfExit(c *gin.Context, queryType QueryType) *gorm.DB {
	for key, method := range queryType {
		switch method {
		case EQ:
			if query, b := c.GetQuery(key); b {
				q.db = q.db.Where(key+" = ?", query)
			}
			break
		case NEQ:
			if query, b := c.GetQuery(key); b {
				q.db = q.db.Where(key+" <> ?", query)
			}
			break
		case LIKE:
			if query, b := c.GetQuery(key); b {
				q.db = q.db.Where(key+" LIKE ?", "%"+query+"%")
			}
			break
		case LESS:
			if query, b := c.GetQuery(key); b {
				q.db = q.db.Where(key+" < ?", query)
			}
			break
		case LESS_EQ:
			if query, b := c.GetQuery(key); b {
				q.db = q.db.Where(key+" <= ?", query)
			}
			break
		case MORE:
			if query, b := c.GetQuery(key); b {
				q.db = q.db.Where(key+" > ?", query)
			}
			break
		case MORE_EQ:
			if query, b := c.GetQuery(key); b {
				q.db = q.db.Where(key+" >= ?", query)
			}
			break
		case BETWEEN:
			if query, b := c.GetQueryArray(key); b && len(query) == 2 {
				q.db = q.db.Where(key+" BETWEEN ? AND ?", query[0], query[1])
			}
			break
		}

	}
	return q.db
}

func NewQueryHelper(db *gorm.DB) IQueryHelper {
	return &QueryHelper{
		db: db,
	}
}
