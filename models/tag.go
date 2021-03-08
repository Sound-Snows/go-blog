package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Tag 标签model
type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//GetTags 获取标签列表
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

//GetTagTotal 总条数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

//ExisTagByName 根据ID查询对象
func ExisTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//AddTag 新增标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

//ExistTagByID 查询标签
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//DeleteTag 删除标签
func DeleteTag(id int) bool {
	db.Where("id=?", id).Delete(&Tag{})
	return true
}

//EditTag 修改标签
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id=?", id).Updates(data)

	return true
}

//BeforeCreate is load CreatedOn data
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate is load ModifiedOn data
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
