package bookcontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-api/config"
	"github.com/mhdianrush/go-api/helper"
	"github.com/mhdianrush/go-api/models"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	var bookResponse []models.BookResponse

	if err := config.DB.Joins("Writer").Find(&books).Find(&bookResponse).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "Book's List", bookResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	// Author's Check
	var author models.Author

	if err := config.DB.First(&author, book.AuthorId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Create Book", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParams)
	if err != nil {
		panic(err)
	}

	var book models.Book
	var bookResponse models.BookResponse

	if err = config.DB.Joins("Writer").First(&book, id).First(&bookResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Book Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "Book Detail", bookResponse)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParams)
	if err != nil {
		panic(err)
	}

	var book models.Book

	if err = config.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Book Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var bookPayload models.Book

	err = json.NewDecoder(r.Body).Decode(&bookPayload)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}

	defer r.Body.Close()

	var author models.Author

	if bookPayload.AuthorId != 0 {
		if err = config.DB.First(&author, bookPayload.AuthorId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Author Not Found", nil)
				return
			}
			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err = config.DB.Where("id = ?", id).Updates(&bookPayload).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 201, "Success Update Book", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParams)
	if err != nil {
		panic(err)
	}

	var book models.Book

	res := config.DB.Delete(&book, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	// means nothing data deleted
	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Book Not Found", nil)
		return
	}

	helper.Response(w, 200, "Success Delete Book", nil)
}
