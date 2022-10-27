package flyweight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleMysqlConnect_Do(t *testing.T) {
	//初始化连接池
	mysqlpool := NewMysqlConnPool(2)
	//获取一个连接
	conn := mysqlpool.Get()
	//执行操作
	fmt.Println(conn.Do())
	//放入连接池
	mysqlpool.Put(conn)
	//OutPut:
	//done
}

// mysql连接数不足的时候，是否新建
func TestNewMysqlConnPool(t *testing.T) {
	//新建只有一个的连接池
	mysqlpool := NewMysqlConnPool(1)
	//获取一个连接，执行完未放入连接池
	conn1 := mysqlpool.Get()
	//又要获取一个连接
	conn2 := mysqlpool.Get()
	//执行
	str := conn2.Do()
	mysqlpool.Put(conn1)
	mysqlpool.Put(conn2)
	assert.Equal(t, "done", str, "fail")
}
