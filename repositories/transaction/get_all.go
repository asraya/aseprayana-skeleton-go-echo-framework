package repositories

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	"aseprayana-skeleton-go/entity"
	"fmt"
	"strings"
)

func (r *transactionRepository) GetAll(req dto.GetAllRequest) (dto.PaginationResponse, error) {
	var data []entity.Transaction

	find := r.DB.Offset((req.Page - 1) * req.Limit).Limit(req.Limit).Order(req.SortBy)

	// Menghitung total data yang sesuai dengan pencarian
	var total int64
	if err := r.DB.Model(&entity.Transaction{}).Count(&total).Error; err != nil {
		return dto.PaginationResponse{}, err
	}

	searchs := req.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break

			}
		}
	}
	if err := find.Find(&data).Error; err != nil {
		return dto.PaginationResponse{}, err
	}

	dataPerPage := req.Limit
	dataPassed := (req.Page - 1) * dataPerPage
	nextPageDataStart := dataPassed + dataPerPage - req.Limit
	nextPage := req.Page + 1
	previousPage := req.Page - 1
	numRows := len(data)

	// Buat respons JSON dengan menggunakan struct PaginationResponse
	response := dto.PaginationResponse{
		TotalData:    int(total),
		TotalRows:    numRows,
		Limit:        req.Limit,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		NextPageData: nextPageDataStart,
		Data:         data,
		Searchs:      searchs,
	}

	return response, nil
}
