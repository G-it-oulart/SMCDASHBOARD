package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func connect_db() (*pgx.Conn, error) {
	DbUser := os.Getenv("DB_USER_NAME")
	DbPw := os.Getenv("DB_USER_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_DATABASE")
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DbUser, DbPw, DbHost, DbPort, DbName)
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func exec_query(query string, args pgx.NamedArgs) ([]any, error) {
	conn, _ := connect_db()
	var rows pgx.Rows
	rows, query_err := conn.Query(context.Background(), query, args)
	if query_err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", query_err)
		os.Exit(1)
	}
	var list []any
	for rows.Next() {
		var value any
		if err := rows.Scan(&value); err != nil {
			return nil, fmt.Errorf("row scan failed: %v", err)
		}
		list = append(list, value)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}
	log.Printf("Query results: %+v\n", list)
	return list, nil
}

func IdentSanit(input_string string) string {
	return pgx.Identifier{input_string}.Sanitize()
}

func filt_dados_pesagens(data_input_init, data_input_end string, material_input string) ([]any, error) {
	Sql_query := fmt.Sprintf("SELECT %s FROM dados_pesagens where data_pesagem BETWEEN @data_init AND @data_end;", IdentSanit(material_input))
	return exec_query(Sql_query, pgx.NamedArgs{"data_init": data_input_init, "data_end": data_input_end})
}
