package homework03

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:64;not null"`
	Posts []Post
}

type Post struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"not null"`
	Title    string `gorm:"size:80;not null"`
	Content  string `gorm:"size:1200;not null"`
	Comments []Comment
}

type Comment struct {
	ID     uint   `gorm:"primaryKey"`
	PostID uint   `gorm:"not null"`
	Text   string `gorm:"size:300;not null"`
}
