package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	_ "github.com/sqlroute/internal/db"
)

type Handler struct {
	DB *sql.DB
}

type SQLTable struct {
	Name    string
	Id      string
	Columns []string
}

func (h *Handler) AllTables(w http.ResponseWriter, r *http.Request) {
	mp, err := h.getTables()
	if err != nil {
		logrus.Errorf("error get tables in all tables handler")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if mp == nil {
		logrus.Errorf("some error in all tables hanbler (mp = nil)")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonEncoded, err := json.Marshal(mp)
	if err != nil {
		logrus.Errorf("error while marshaling map response in all tables handler %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonEncoded)
}

func (h *Handler) LimitTables(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *Handler) TableInformation(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *Handler) UpdateTable(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *Handler) DeleteRow(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *Handler) NewTable(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *Handler) getTables() ([]map[string]SQLTable, error) {
	rows, err := h.DB.Query("SHOW TABLES")
	if err != nil {
		logrus.Infof("error while show tables %v", err)
		return nil, err
	}
	mp := []map[string]SQLTable{}
	for rows.Next() {
		var SQLTablename string
		tmp := map[string]SQLTable{}
		err := rows.Scan(&SQLTablename)
		if err != nil {
			logrus.Infof("error get sql table name %v", err)
			return nil, err
		}
		cols, err := h.DB.Query("SELECT * FROM " + SQLTablename)
		if err != nil {
			logrus.Infof("error while get columns")
			return nil, err
		}
		columns, err := cols.Columns()
		if err != nil {
			logrus.Infof("error while get columns")
			return nil, err
		}
		tmp[SQLTablename] = SQLTable{
			Name:    SQLTablename,
			Columns: columns,
		}
		mp = append(mp, tmp)
	}

	return mp, nil
}
