package book_inventory_system_domain

type User struct {
	Users []struct {
		UserID       int    `json:"user_id,omitempty"`
		Name         string `json:"name"`
		Password     string `json:"password"`
		LoginStatus  string `json:"login_status"`
		RegisterDate string `json:"register_date"`
	} `json:"users"`
}

type Admin struct {
	Admins []struct {
		AdminID int `json:"admin_id"`
	} `json:"admins"`
}

type Reader struct {
	Readers []struct {
		ReaderID   int   `json:"reader_id"`
		InstanceID []int `json:"instance_id"`
	} `json:"readers"`
}

type Instance struct {
	Instances []struct {
		InstanceID int `json:"instance_id"`
		BookID     int `json:"book_id"`
		Status     int `json:"status"`
	} `json:"instances"`
}

type Book struct {
	Books []struct {
		BookID       int    `json:"book_id"`
		Name         string `json:"name"`
		AuthorID     int    `json:"author_id"`
		GenreID      int    `json:"genre_id"`
		ProductionID int    `json:"production_id"`
		LanguageID   int    `json:"language_id"`
		Description  string `json:"description"`
	} `json:"books"`
}

type Author struct {
	Authors []struct {
		AuthorID     int    `json:"author_id"`
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		Patronymic   string `json:"patronymic"`
		ProductionID int    `json:"production_id"`
	} `json:"authors"`
}

type Production struct {
	Productions []struct {
		ProductionID int    `json:"production_id"`
		Name         string `json:"name"`
	} `json:"productions"`
}

type Genre struct {
	Genres []struct {
		GenreID int    `json:"genre_id"`
		Name    string `json:"name"`
	} `json:"genres"`
}

type Language struct {
	Languages []struct {
		LanguageID int    `json:"language_id"`
		Name       string `json:"name"`
	} `json:"languages"`
}

type AdminMapField struct{}

type AuthorMapField struct {
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Patronymic   string `json:"patronymic"`
	ProductionID int    `json:"production_id"`
}

type BookMapField struct {
	Name         string `json:"name"`
	AuthorID     int    `json:"author_id"`
	GenreID      int    `json:"genre_id"`
	ProductionID int    `json:"production_id"`
	LanguageID   int    `json:"language_id"`
	Description  string `json:"description"`
}

type UserMapField struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	LoginStatus  string `json:"login_status"`
	RegisterDate string `json:"register_date"`
}

type ReaderMapField struct {
	InstanceID []int `json:"instance_id"`
}

type InstanceMapField struct {
	BookID int `json:"book_id"`
	Status int `json:"status"`
}

type ProductionMapField struct {
	Name string `json:"name"`
}

type GenreMapField struct {
	Name string `json:"name"`
}

type LanguageMapField struct {
	Name string `json:"name"`
}
