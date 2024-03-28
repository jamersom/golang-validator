package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
)

type User struct {
	Username string `json:"username" validate:"required"`
	Age      int    `json:"age" validate:"gte=0,lte=120"`
	CPF      string `json:"cpf" validate:"CPF"`
}

func main() {
	// Criar uma instância do validador
	validate := validator.New()

	// Adicionar validação personalizada para a tag "customTag"
	validate.RegisterValidation("CPF", customValidator)

	// Instância de User com valores inválidos para demonstração
	user := User{
		Username: "",
		Age:      -1,
		CPF:      "000.111.222-88",
	}

	// Validar a estrutura User
	err := validate.Struct(user)
	if err != nil {
		fmt.Println("Erro de validação:", err)
	}
}

// Função de validação personalizada para a tag "customTag"
func customValidator(fl validator.FieldLevel) bool {
	// Verificar se o valor do campo está dentro das condições desejadas
	fieldValue := fl.Field().String()
	if ValidateCPF(fieldValue) {
		return true
	}

	return false
}

// ValidateCPF verifica se um CPF é válido
func ValidateCPF(cpf string) bool {
	// Remover caracteres especiais do CPF
	regexp := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	cleanCPF := regexp.ReplaceAllString(cpf, "")

	// Verificar se o CPF tem 11 dígitos
	if len(cleanCPF) != 11 {
		return false
	}

	// Verificar se todos os dígitos são iguais
	if allEq(cleanCPF) {
		return false
	}

	size := 9
	position := 10

	d := cleanCPF[:size]
	digit := calculateDigit(d, position)

	d = d + digit
	return cleanCPF == d+digit
}

// calculateCPFDigit calcula um dígito verificador de CPF
func calculateCPFDigit(cpf string) int {
	sum := 0
	for i := 0; i < len(cpf); i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (10 - i)
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

func allEq(doc string) bool {
	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

// calculateDigit calculates the next digit for the given document.
func calculateDigit(doc string, position int) string {
	var sum int

	for _, r := range doc {
		sum += toInt(r) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}

// toInt converts a rune to an int.
func toInt(r rune) int {
	return int(r - '0')
}
