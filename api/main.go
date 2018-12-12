import (
    "./routes"
    "./database"
)


func main() {
    routes.Load_config()
    // routes.Get_hottest_posts("20")

<<<<<<< HEAD:src/main.go
    routes.TIME_SERIES_INTRADAY("", "", "", "", "", "");
    database.SetupDB();
=======
    // routes.TIME_SERIES_INTRADAY("", "", "", "", "", "");
>>>>>>> 95915fbce29ed5a295409d5e282e7539eaa0b2f0:api/main.go
    
}