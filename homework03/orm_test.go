package homework03

import (
	"fmt"
	"testing"

	"github.com/caitoufire-tech/go-homework/testutil"
	"gorm.io/gorm"
)

func TestOrm(t *testing.T) {
	db := testutil.NewTestDB(t, "task_sqlite.db")
	db.AutoMigrate(&User{})
	newUser := User{Name: "Bob", Phone: "11"}
	if err := db.Create(&newUser).Error; err != nil {
		t.Fatalf("create error:%+v", err)
	}
	batchUser := []User{
		{Name: "Sam", Phone: "12"},
		{Name: "Du", Phone: "13"},
	}
	if err := db.Create(&batchUser).Error; err != nil {
		t.Fatalf("create error:%+v", err)
	}
	any := User{}
	if err := db.Where("phone = ?", "12").First(&any).Error; err != nil {
		t.Fatalf("first error:%+v", err)
	}
	t.Logf("any user:%+v", any)
	other := User{}
	db.Take(&other)
	t.Logf("other user:%+v", other)
	anyUsers := []User{}
	if err := db.Where("phone like ?", "%1%").Find(&anyUsers).Error; err != nil {
		t.Fatalf("find where error:%+v", err)
	}
	t.Logf("anyUsers:%+v", anyUsers)

	type UserPhone struct {
		Phone string
	}
	allPhone := []UserPhone{}
	if err := db.Model(&User{}).Select("phone").Scan(&allPhone).Error; err != nil {
		t.Fatalf("select scan error:%+v", err)
	}
	t.Logf("allPhone:%+v", allPhone)

	newUser.Name = "BobNewName also phone nil"
	if err := db.Save(&newUser).Error; err != nil {
		t.Fatalf("save error:%+v", err)
	}
	t.Logf("bob new name:%+v", newUser)
	newBob := User{}
	if err := db.Where("name like ?", "%bob%").First(&newBob).Error; err != nil {
		t.Fatalf("new Bob find error:%+v", err)
	}
	t.Logf("new Bob:%+v", newBob)

	if err := db.Model(&newBob).Update("name", "bob other name").Error; err != nil {
		t.Fatalf("update 1:%+v", err)
	}
	t.Logf("new Bob:%+v", newBob)
	if err := db.Model(&newBob).Updates(map[string]interface{}{"name": "bob other other name", "phone": "12345"}).Error; err != nil {
		t.Fatalf("update 1:%+v", err)
	}
	t.Logf("new Bob:%+v", newBob)

}

func TestExec(t *testing.T) {
	db := testutil.NewTestDB(t, "task_sqlite.db")
	if err := db.Exec("delete from test_users where id = ?", 10).Error; err != nil {
		t.Fatalf("exec fail:%+v", err)
	}
}

func TestOther(t *testing.T) {
	db := testutil.NewTestDB(t, "task_sqlite.db")
	_, err := CreateUser(db, "aloha", "12342123221")
	if err != nil {
		t.Fatalf("create error:%+v", err)
	}
	users, _ := SearchUsersByEmail(db, "2123", 23, 10)
	fmt.Println("users:", users)
}

func CreateUser(db *gorm.DB, name, phone string) (*User, error) {
	db.AutoMigrate(&User{})
	// 你的实现
	newUser := User{Name: name, Phone: phone}
	err := db.Create(&newUser).Error

	return &newUser, err
}

func SearchUsersByEmail(db *gorm.DB, phonePattern string, page, size int) ([]User, error) {
	// 你的实现
	users := []User{}
	db.Scopes(Page(page, size)).
		Model(&User{}).
		Where("phone like ?", "%"+phonePattern+"%").
		Find(&users)
	return users, nil
}

func Page(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(size).Offset(page)
	}
}
