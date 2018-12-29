package main 

import (
    "../api/routes"
    "../api/database"
)


func main() {
    routes.Load_config()
    // routes.Get_hottest_posts("20")

    routes.TIME_SERIES_INTRADAY("", "", "", "", "", "");
    database.SetupDB();
    // routes.TIME_SERIES_INTRADAY("", "", "", "", "", "");
    
}