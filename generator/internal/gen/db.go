package gen

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"go-knowledge/generator/internal/model"
	"go-knowledge/generator/pkg/arg"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"log"
	"net"
	"net/url"
	"os"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (v *ViaSSHDialer) Dial(addr string) (net.Conn, error) {
	return v.client.Dial("tcp", addr)
}

func tunnel(t string) *sql.DB {
	u, err := url.Parse(t)
	if err != nil {
		log.Fatalln("invalid ssh tunnel", err)
	}

	sshPass, _ := u.User.Password()
	var agentClient agent.Agent
	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer conn.Close()
		agentClient = agent.NewClient(conn)
	}

	sshConfig := &ssh.ClientConfig{
		User:            u.User.Username(),
		Auth:            []ssh.AuthMethod{},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}

	if sshPass != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return sshPass, nil
		}))
	}

	if sshconn, err := ssh.Dial("tcp", u.Host, sshConfig); err == nil {
		mysql.RegisterDial("mysql+tcp", (&ViaSSHDialer{sshconn}).Dial)
		db, err := sql.Open("mysql", arg.MySQL)
		if err != nil {
			log.Fatal("sql open fail", err)
			return nil
		}
		return db

	} else {
		log.Fatal("ssh conn fail", err)
	}
	return nil
}

func GetTableSchemas(dsn string, db string, table string) (resp []model.TableSchema, err error) {
	var conn *sql.DB
	if arg.SshTunnel != "" {
		conn = tunnel(arg.SshTunnel)
	} else {
		conn, err = sql.Open("mysql", dsn)
	}

	if err != nil {
		log.Panic("[GetTableSchema] mysql open", err.Error())
		return nil, err
	}

	defer conn.Close()

	q := `SELECT 
TABLE_NAME, COLUMN_NAME, IS_NULLABLE, DATA_TYPE, CHARACTER_MAXIMUM_LENGTH, 
NUMERIC_PRECISION, NUMERIC_SCALE,COLUMN_TYPE,COLUMN_KEY,COLUMN_COMMENT 
FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME LIKE ? ORDER BY TABLE_NAME, ORDINAL_POSITION
`
	rows, err := conn.Query(q, db, table)
	if err != nil {
		return nil, err
	}

	columns := make([]model.TableSchema, 0)
	for rows.Next() {
		cs := model.TableSchema{}
		err := rows.Scan(&cs.TableName, &cs.ColumnName, &cs.IsNullable, &cs.DataType, &cs.CharacterMaximumLength, &cs.NumericPrecision, &cs.NumericScale, &cs.ColumnType, &cs.ColumnKey, &cs.Comment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, cs)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return columns, nil
}
