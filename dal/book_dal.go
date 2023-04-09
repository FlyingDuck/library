package dal

import (
	"fmt"
	"github.com/FlyingDuck/library/model"
	"gorm.io/gorm"
)

type BookDal struct {
}

func NewBookDal() *BookDal {
	return &BookDal{}
}

func (dal *BookDal) FindByModify(offset, limit int) (books []*model.Book, total int64, dalErr error) {
	books = make([]*model.Book, 0)
	result := libraryDB.Offset(offset).Limit(limit).Order("modify_time desc").Where("del = 0").Find(&books)
	if result.Error != nil {
		dalErr = fmt.Errorf("find book failed: %w", result.Error)
		return nil, 0, dalErr
	}

	result = libraryDB.Model(&model.Book{}).Where("del = 0").Count(&total)
	if result.Error != nil {
		dalErr = fmt.Errorf("count book failed: %w", result.Error)
		return nil, 0, dalErr
	}

	return books, total, nil
}

func (dal *BookDal) FindByID(id int64) (book *model.Book, dalErr error) {
	book = &model.Book{}
	result := libraryDB.Where("del = 0").First(book, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return book, nil
		}
		dalErr = fmt.Errorf("find book(ID=%d) failed: %w", id, result.Error)
		return nil, dalErr
	}
	return book, nil
}

func (dal *BookDal) Add(book *model.Book) (ID int64, dalErr error) {
	result := libraryDB.Omit("modify_time", "create_time").Create(book)
	if result.Error != nil {
		dalErr = fmt.Errorf("add new book failed: %w", result.Error)
		return 0, dalErr
	}

	return book.ID, nil
}

func (dal *BookDal) Update(book *model.Book) (dalErr error) {
	result := libraryDB.Model(&model.Book{}).Omit("modify_time", "create_time").Updates(book)
	if result.Error != nil {
		dalErr = fmt.Errorf("update book(ID=%d) failed: %w", book.ID, result.Error)
		return dalErr
	}
	return nil
}

func (dal *BookDal) DelByID(id int64) (dalErr error) {
	result := libraryDB.Model(&model.Book{}).Omit("modify_time", "create_time").Update("del", 1)
	if result.Error != nil {
		dalErr = fmt.Errorf("delete book(ID=%d) failed: %w", id, result.Error)
		return dalErr
	}
	return nil
}
