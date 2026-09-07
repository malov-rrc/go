package main

import (
	"bank_transactions/models"
	"bank_transactions/utils"
	"flag"
	"fmt"
	"os"
)

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Fprintln(os.Stderr, "укажите путь к файлу")
	//	os.Exit(1)
	//}
	account := models.Account{Balance: 0}
	fileName := flag.String("file", "transactions.txt", "текстовый файл с транзакциями")
	flag.Parse()
	file := utils.OpenFile(*fileName)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "ошибка при закрытии файла")
			os.Exit(1)
		}
	}(file)
	fileLines := utils.GetNotEmptyFileLines(file)
	fmt.Println(account.MakeTransactions(fileLines))
}
