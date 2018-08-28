package main

import (
 db "aze.org/database"
)

func main() {
 defer db.SqlDB.Close()
 router := initRouter()
 router.Run(":8000")
}
