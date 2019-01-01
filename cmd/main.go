package main

import (
        // "../api/database"
        "../api/routes"
        // "fmt"
)

func main() {
        routes.Load_config()
        routes.GetHottestPosts("20")

        // routes.TimeSeriesIntraday("", "", "", "", "", "", false)
        // routes.TimeSeriesDaily("", "", "", "", "", "", false)

        // routes.TimeSeriesWeekly("", "", "", "", "", "", false)
        // routes.TimeSeriesWeeklyAdjusted("", "", "", "", "", "", "", false)


        // fmt.Printf("Setup DB. \n")
        // database.SetupDB()
        // routes.TIME_SERIES_INTRADAY("", "", "", "", "", "");

}