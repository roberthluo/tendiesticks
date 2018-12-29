package main 

import (
    "../api/routes"
    "../api/database"
)


func main() {
    routes.Load_config()
    // routes.GetPottestPosts("20")

    routes.TimeSeriesIntraday("", "", "", "", "", "");
    database.SetupDB();
    // routes.TIME_SERIES_INTRADAY("", "", "", "", "", "");
    
}