package main

import (
    "seedy/modules/connection"
    "seedy/modules/queries"
    "seedy/modules/ui"
)

func main() {
    db := databaseConnection.ConnectDb()
    tables := databaseQueries.ListTables(db)
    // on selection of table,  use columns to create fake data
//     columns := databaseQueries.ListColumns(db, "merchants")
    seedyUI.ShowMe(tables)
}