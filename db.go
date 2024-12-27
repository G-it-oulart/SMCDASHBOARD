package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

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

func exec_query(query string, args pgx.NamedArgs) ([]byte, error) {
	conn, _ := connect_db()
	var rows pgx.Rows
	rows, query_err := conn.Query(context.Background(), query, args)
	if query_err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", query_err)
		os.Exit(1)
	}
	json:=PgSqlRowsToJson(rows)
	fmt.Println("JSON Result::> ", string(json))
	return json,nil
}
func insert_into_db(linha,cor,current_date string,massa,primer,verniz,esmalte,tingidor float64) error {
	conn, _ := connect_db()
	query := fmt.Sprintf("INSERT INTO dados_pesagens (linha,massa,primer,verniz,esmalte,tingidor,data_pesagem,color_id) VALUES ('%s',%f,%f,%f,%f,%f,'%s',(select color_id from configurados_standards where configurado = '%s'));",linha,massa,primer,verniz,esmalte,tingidor,current_date,cor)
	_,err:= conn.Exec(context.Background(),query)
	if err != nil{
		log.Println("Error Inserting values")
		return err
	}
	return nil
}

func PgSqlRowsToJson(rows pgx.Rows) []byte {
	fieldDescriptions := rows.FieldDescriptions()
	var columns []string
	for _, col := range fieldDescriptions {
		columns = append(columns, string(col.Name))
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
 		values, _ := rows.Values()
	 	for i, v := range values {
	 		valuePtrs[i] = reflect.New(reflect.TypeOf(v)).Interface() // allocate pointer to type
 		}
	 	break
	}
	rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := reflect.ValueOf(valuePtrs[i]).Elem().Interface()
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)

	for rows.Next() {
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := reflect.ValueOf(valuePtrs[i]).Elem().Interface()
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	
	jsonData, _ := json.Marshal(tableData)
	return jsonData
}

func IdentSanit(input_string string) string {
	return pgx.Identifier{input_string}.Sanitize()
}

func filt_dados_pesagens(data_input_init, data_input_end, material_input, material_color string) ([]byte, error) {
	Sql_query := fmt.Sprintf("SELECT dados_pesagens.%s as material FROM dados_pesagens INNER JOIN configurados_standards ON dados_pesagens.color_id = configurados_standards.color_id  WHERE dados_pesagens.%s IS NOT NULL AND configurado = @material_color AND data_pesagem BETWEEN @data_init AND @data_end ORDER BY id ASC; ", IdentSanit(material_input), IdentSanit(material_input))
	filt_list, _ := exec_query(Sql_query, pgx.NamedArgs{"data_init": data_input_init, "data_end": data_input_end, "material_color": material_color})
	return filt_list, nil
}

func return_dates(data_input_init, data_input_end, material_input, material_color string) ([]byte, error) {
	Sql_query := fmt.Sprintf("SELECT data_pesagem as string_value FROM dados_pesagens  INNER JOIN configurados_standards ON dados_pesagens.color_id = configurados_standards.color_id  WHERE dados_pesagens.%s IS NOT NULL AND configurado = @material_color AND data_pesagem BETWEEN @data_init AND @data_end ORDER BY id ASC;" , IdentSanit(material_input))
	filt_list, _ := exec_query(Sql_query, pgx.NamedArgs{"data_init": data_input_init, "data_end": data_input_end, "material_color": material_color})
	return filt_list, nil
}

func return_standards(material_input, material_color string) ([]byte, error) {
	Sql_query := fmt.Sprintf("SELECT %s as material FROM configurados_standards WHERE %s IS NOT NULL AND configurado = @material_color", IdentSanit(material_input),IdentSanit(material_input))
	standards_list, _ := exec_query(Sql_query, pgx.NamedArgs{"material_color": material_color})
	return standards_list, nil
}

func return_color_names(material_color string) ([]byte,) {
	Sql_query := "SELECT configurado as string_value FROM configurados_standards"
	colors_list, _ := exec_query(Sql_query, pgx.NamedArgs{"material_color": ""})
	return colors_list
}
