package model

func Migration() {
	Db.Set("gorm:table_options", "charset=utf8mb4;").
		AutoMigrate(&User{})
}
