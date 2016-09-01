package conn_pool
import {
	"fmt"
	"net"
	"errors"
}

type connPool struct {
	conns chan net.Conn
}

func (c *connPool) get() (net.Conn,error){

}

func (c *connPool) put(conn net.Conn) error {

}

func (c *connPool) close() {


}

func (c *connPool) len() int {

}