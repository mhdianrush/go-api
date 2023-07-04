package authorcontroller

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
	var author []models.Author

	err := config.DB.Find(&author).Error
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}

	helper.Response(w, 200, "List Author's", author)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}

	defer r.Body.Close()

	if err = config.DB.Create(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}
	helper.Response(w, 200, "Success Create Author", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParams)
	if err != nil {
		panic(err)
	}

	var author models.Author

	if err = config.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "Success Get Detail Author", author)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParams)
	if err != nil {
		panic(err)
	}

	var author models.Author

	if err = config.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}
	defer r.Body.Close()

	// update
	if err = config.DB.Where("id = ?", id).Updates(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 201, "Success Update Author", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParams)
	if err != nil {
		panic(err)
	}

	var author models.Author

	res := config.DB.Delete(&author, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	// means nothing data deleted
	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Author Not Found", nil)
		return
	}

	helper.Response(w, 200, "Success Delete Author", nil)
}
