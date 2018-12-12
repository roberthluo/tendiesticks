package main

import (
    "./routes"
    "./database"
)


func main() {

    // routes.Get_hottest_posts("20")

    routes.TIME_SERIES_INTRADAY("", "", "", "", "", "");
    database.SetupDB();
    
}