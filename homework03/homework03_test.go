package homework03

import (
	"fmt"
	"testing"

	"github.com/caitoufire-tech/go-homework/testutil"
	"gorm.io/gorm"
)

func TestTask01(t *testing.T) {

	db := testutil.NewTestDB(t, "task_sqlite.db")
	initUserTable(t, db)

	users := NewUser()
	result := db.Session(&gorm.Session{FullSaveAssociations: true}).CreateInBatches(users, len(users))
	if result.Error != nil {
		t.Fatalf("CreateInBatches failed entirely: %v", result.Error)
	}
	t.Logf("Inserted %d users with posts and comments", len(users))
}

func TestTask02_1(t *testing.T) {
	db := testutil.NewTestDB(t, "task_sqlite.db")
	//initUserTable(t, db)
	ID := 1
	users := []User{}
	db.Preload("Posts.Comments").Where("ID = ?", ID).Find(&users)
	if len(users) == 0 {
		t.Fatalf("empty users")
	}
	// 这个为什么没有输出
	t.Logf("users: %+v", users)
	t.Log("") // 增加一个空行日志
	fmt.Printf("【直接打印】users: %+v\n", users)

}

func TestTask02_2(t *testing.T) {
	db := testutil.NewTestDB(t, "task_sqlite.db")

	type MaxPost struct {
		PostID uint
		Cnt    uint
	}
	p := &Post{}
	result := db.Model(&Post{}).Select("test_posts.*,count(1) cnt").
		Joins("left join test_comments c on c.post_id = test_posts.id").
		Group("test_posts.id").
		Order("cnt desc").
		Limit(1).
		First(&p)

	if result.Error != nil {
		t.Fatalf("error:%+v", result.Error)
	}

	fmt.Println(p)
}

func TestTask02_2_WithCount(t *testing.T) {
	db := testutil.NewTestDB(t, "task_sqlite.db")

	// 方法2：查询带计数结果，使用Scan
	type PostWithCount struct {
		ID      uint
		UserID  uint
		Title   string
		Content string
		Cnt     int64
	}

	var result []PostWithCount
	err := db.Model(&Post{}).Select("test_posts.id, test_posts.user_id, test_posts.title, test_posts.content, count(test_comments.id) as cnt").
		Joins("left join test_comments c on c.post_id = test_posts.id").
		Group("test_posts.id").
		Order("cnt desc").
		Limit(1).
		Scan(&result).Error

	if err != nil {
		t.Fatalf("error:%+v", err)
	}

	if len(result) > 0 {
		fmt.Printf("最多评论的文章 - 标题: %s, 评论数: %d\n", result[0].Title, result[0].Cnt)
	}
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
