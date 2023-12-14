package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ResponseDto struct {
	Page        int    `json:"page"`
	RowPerPages int    `json:"per_page"`
	TotalRows   int    `json:"total"`
	TotalPages  int    `json:"total_pages"`
	Data        []User `json:"data"`
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://reqres.in/api/users", nil)
	if err != nil {
		log.Printf("failed to request: %v \n", err.Error())
		return
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("err response: %v \n", err.Error())
		return
	}
	bodyByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("err ReadAll: %v \n", err.Error())
		return
	}
	var responseDto ResponseDto
	err = json.Unmarshal(bodyByte, &responseDto)
	if err != nil {
		log.Printf("err unmarshal: %v \n", err.Error())
		return
	}
	fmt.Println("data:")
	for _, v := range responseDto.Data {
		fmt.Println("ID:", v.ID)
		fmt.Println("Name:", v.FirstName, v.LastName)
		fmt.Println("Email:", v.Email)
		fmt.Println()
	}
	fmt.Println("paging:")
	fmt.Println("page:", responseDto.Page)
	fmt.Println("rowsPerPage:", responseDto.RowPerPages)
	fmt.Println("totalRows:", responseDto.TotalRows)
	fmt.Println("totalPages:", responseDto.TotalPages)
}
