package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

/*
   @Time : 2018/9/4 10:10 
   @Author : ff
*/

const (
	_USER_TABLE = "user"

	_CATEGORY_TABLE = "category"

	_TOPIC_TABLE = "topic"

	_COMMENT_TABLE = "comment"
)

// User Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

// 分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

// 文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Content         string `orm:"size(5000)"`
	Lables          string
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

// 评论
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}


/**
   注册db
 */
func RegisterDB() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/go-web?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User),new(Category),new(Topic),new(Comment))

}

func AddCategory(title string) error {
	o := orm.NewOrm()

	c := &Category{
		Title: title,
		Created: time.Now(),
		TopicTime: time.Now(),
	}

	qs := o.QueryTable(_CATEGORY_TABLE)
	err := qs.Filter("title", title).One(c)
	if err == nil {
		return err
	}

	cid,err := o.Insert(c)
	if err != nil {
		return err
	}
	beego.Informational("添加分类，返回分类id：",cid)


	return nil
}

func GetCategorys() ([]*Category,error){
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable(_CATEGORY_TABLE)
	size, err := qs.All(&cates)
	beego.Info("查询分类列表条数：",size)
	return cates, err
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func AddTopic(title, category,lable,content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:   title,
		Content: content,
		Lables: lable,
		Category: category,
		Created: time.Now(),
		Updated: time.Now(),
	}
	_, err := o.Insert(topic)
	return err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)

	topic.Lables = strings.Replace(strings.Replace(
		topic.Lables, "#", " ", -1), "$", "", -1)
	return topic, nil
}

func ModifyTopic(tid, title, category,label,content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Lables = label
		topic.Category = category
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return nil
}

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: tidNum}
	_, err = o.Delete(topic)
	return err
}

func GetAllTopics(isDesc bool) (topics []*Topic, err error) {
	o := orm.NewOrm()

	topics = make([]*Topic, 0)

	qs := o.QueryTable(_TOPIC_TABLE)
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	return err
}

func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	replies = make([]*Comment, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err
}

func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	reply := &Comment{Id: ridNum}
	_, err = o.Delete(reply)
	return err
}