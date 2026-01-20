package homework03

import (
	"testing"

	"github.com/caitoufire-tech/go-homework/testutil"
	"gorm.io/gorm"
)

func TestTask01(t *testing.T) {

	db := testutil.NewTestDB(t, "task_sqlite.db")
	initUserTable(t, db)

	users := NewUser()
	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).CreateInBatches(users, len(users)); err != nil {
		t.Fatalf("CreateInBatches users: %v", err)
	}
	t.Logf("Inserted %d users with posts and comments", len(users))
}

func TestTask02(t *testing.T) {
	db := testutil.NewTestDB(t, "task01_sqlite.db")
	initUserTable(t, db)
	ID := 1
	users := []User{}
	db.Preload("Posts.Comments").Where("ID = ?", ID).Find(&users)
	t.Logf("Find User:%v", users)
}

func initUserTable(t *testing.T, db *gorm.DB) {
	if err := db.AutoMigrate(&User{}, &Comment{}, &Post{}); err != nil {
		t.Fatalf("AutoMigrate User: %v", err)
	}
}

func NewUser() []User {
	return []User{
		{
			Name: "Alice",
			Posts: []Post{
				{
					Title:   "First Post",
					Content: "This is Alice's first post.",
					Comments: []Comment{
						{Text: "Great post!"},
						{Text: "Thanks for sharing."},
					},
				},
			},
		},
		{Name: "Bob",
			Posts: []Post{
				{
					Title:   "Hello World",
					Content: "Bob says hello to the world.",
					Comments: []Comment{
						{Text: "Welcome Bob!"},
						{Text: "Nice to meet you."},
					},
				},
			},
		},
	}
}
