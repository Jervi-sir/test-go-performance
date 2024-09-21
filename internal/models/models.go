package models

type Item struct {
	ID          int
	Name        string
	Price       string
	Sizes       string
	Genders     string
	Images      []string
	LastReposted string
	WilayaCode  string
}

type Shop struct {
	ID          int
	ShopName    string
	ShopImage   string
	Bio         string
	Contacts    string
	MapLocation string
	NbFollowers int
	NbLikes     int
	WilayaID    int
}

type Wilaya struct {
	ID   int
	Name string
	Code int
}