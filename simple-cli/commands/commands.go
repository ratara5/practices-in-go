package commands

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/ratara5/simple-cli/expenses"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	str, err := reader. ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.Replace(str, "\r\n", "", 1)

	return str, nil
}

func ShowInConsole(expensesList []float32){
	fmt.Println(contentString(expensesList))
}

func contentString(expensesList []float32) string{
	builder := strings.Builder{} //Captura valores, los mantiene en búfer y al final, permite mostrarlos


	fmt.Println("")
	for i, expense := range expensesList{
		max, min, total, avg := expensesDetails(expensesList)
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))

		if i==len(expensesList)-1{
			builder.WriteString("")
			builder.WriteString("===========================================\n")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", max))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", min))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", avg))
			builder.WriteString("===========================================\n")
		}
	}

	return builder.String()
}

func expensesDetails(expensesList []float32) (max, min, total, avg float32){

	if len(expensesList) == 0{
		return
	}

	min = expenses.Min(expensesList ...)
	max = expenses.Max(expensesList ...)
	total = expenses.Sum(expensesList ...)
	avg = expenses.Average(expensesList ...)

	return max, min, total, avg
}

func Export(fileName string, list []float32) error{
	f, err := os.Create(fileName)
	if err != nil{
		return err
	}

	defer f.Close() //defer aplaza el comando hasta el final de cualquier return

	w := bufio.NewWriter(f)

	_, err = w.WriteString(contentString(list))
	if err != nil{
		return err
	}

	w.Flush() //Flush guarda el archivo en el path indicado
	return nil
}