package homework03

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:64;not null"`
	Posts   []Post
	PostNum uint `gorm:"default:0"`
}

type Post struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"not null"`
	Title      string `gorm:"size:80;not null"`
	Content    string `gorm:"size:1200;not null"`
	Comments   []Comment
	HasComment bool `gorm:"default:0"`
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	result := tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_num", gorm.Expr("post_num + 1"))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Printf("用户 %d 不存在或更新失败\n", p.UserID)
	}

	return nil
}

func (c *Comment) BeforeDelete(tx *gorm.DB) error {
	// 在删除前保存 post_id 到上下文
	fmt.Println("comment=", *c)
	tx.Statement.Context = context.WithValue(tx.Statement.Context, "post_id", c.PostID)
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	postID, ok := tx.Statement.Context.Value("post_id").(uint)
	if !ok {
		return nil
	}
	fmt.Println("postId:", postID)
	var count int64
	result := tx.Model(c).Where("post_id = ?", postID).Count(&count)
	if err := result.Error; err != nil {
		return err
	}
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("has_comment", 0).Error
	}
	return nil
}

type Comment struct {
	ID        uint           `gorm:"primaryKey"`
	PostID    uint           `gorm:"not null"`
	Text      string         `gorm:"size:300;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // 启用软删除
}
