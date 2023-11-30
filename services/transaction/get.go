package services

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	"fmt"
)

func (s *transactionService) GetAll(paginationDTO dto.GetAllRequest) (dto.PaginationResponse, error) {

	data, err := s.TransactionR.GetAll(paginationDTO)
	if err != nil {
		return dto.PaginationResponse{}, err
	}
	// get current url path

	// search query params
	searchQueryParams := ""

	for _, search := range paginationDTO.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	return data, nil
	// dto.Response{Status: true, Data: data}
}

// 	return data, nil
// }
