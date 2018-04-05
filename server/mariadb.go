package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addAgama                = `insert into Agama(IDAgama,NamaAgama,Status,Keterangan,CreateBy, CreateOn)values (?,?,?,?,?,?)`
	selectAgamaByKeterangan = `select IDAgama,NamaAgama,Status,Keterangan,CreateBy from Agama where Keterangan like ?`
	selectAgama             = `select IDAgama,NamaAgama,Status,Keterangan,CreateBy from Agama where Status=1`
	updateAgama             = `update Agama set NamaAgama=?,Status=?,Keterangan=?,UpdatedBy=?,UpdateOn=? where IDAgama=?`
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddAgama(agama Agama) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addAgama, agama.IDAgama, agama.NamaAgama, OnAdd, agama.Keterangan, agama.CreateBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadAgama() (Agamas, error) {
	agama := Agamas{}
	rows, _ := rw.db.Query(selectAgama)
	defer rows.Close()
	for rows.Next() {
		var c Agama
		err := rows.Scan(&c.IDAgama, &c.NamaAgama, &c.Status, &c.Keterangan, &c.CreateBy)
		if err != nil {
			fmt.Println("error query:", err)
			return agama, err
		}
		agama = append(agama, c)
	}
	//fmt.Println("db nya:", customer)
	return agama, nil
}

func (rw *dbReadWriter) UpdateAgama(agm Agama) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateAgama, agm.NamaAgama, agm.Status, agm.Status, agm.UpdateBy, time.Now(), agm.IDAgama)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadAgamaByKeterangan(keterangan string) (Agamas, error) {

	agama := Agamas{}
	rows, _ := rw.db.Query(selectAgamaByKeterangan, keterangan)

	defer rows.Close()

	for rows.Next() {
		var c Agama
		err := rows.Scan(&c.IDAgama, &c.NamaAgama, &c.Status, &c.Keterangan, &c.CreateBy)
		if err != nil {
			fmt.Println("error query:", err)
			return agama, err
		}

		agama = append(agama, c)
	}

	return agama, nil
}
