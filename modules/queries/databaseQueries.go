package databaseQueries

import (
    "fmt"
    "database/sql"
    "log"
    _ "github.com/lib/pq"
   )

func ListTables(db *sql.DB) (named_tables []string) {
   rows, err := db.Query("SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';")

    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    table_names := []string{}
    for rows.Next() {
        var tablename string
        if errors := rows.Scan(&tablename); errors != nil {
                log.Fatal(errors)
        }
        table_names = append(table_names, tablename)

    }
    if errored := rows.Err(); errored != nil {
        log.Fatal(errored)
    }
    return table_names
}

func ListColumns(db *sql.DB, tableName string) (named_columns []string){
   column, err := db.Query(fmt.Sprintf( "SELECT table_name, column_name, data_type from information_schema.columns where table_name = '%s';", tableName))
    if err != nil {
        log.Fatal(err)
    }
    defer column.Close()
    column_names := []string{}
    for column.Next() {
        var table_name string
        var column_name string
        var data_type string
        if errors := column.Scan(&table_name, &column_name, &data_type); errors != nil {
                log.Fatal(errors)
        }
        fmt.Println(column_name)
        column_names = append(column_names, column_name)
    }

    return column_names
}