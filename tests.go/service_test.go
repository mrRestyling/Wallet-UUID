package tests

import (
	"testing"
	"time"
	"wallet/internal/models"
	"wallet/internal/service"

	"github.com/stretchr/testify/assert"
)

type TestService struct {
	service.StorageInt
}

func (s *TestService) Create(user models.User) (string, error) {
	return "Успешно добавлено", nil
}

func (s *TestService) Change(wallet models.Wallet) (string, error) {
	return "Успешно добавлено", nil
}

func (s *TestService) Balance(wallet models.Wallet) (string, error) {
	return "1 рубль", nil
}

func TestCreate(t *testing.T) {

	// Создаем тестовую структуру
	testService := &TestService{}
	svc := service.New(testService)

	// 1. Тест Happy
	testM := models.Wallet{
		WalletID: "4f9a5a4a-2e9b-4d7a-8f4a-2e9b4d7a8f4a",
	}
	// Вызываем метод Create
	result, err := svc.Create(testM)

	// Проверяем результаты
	assert.NoError(t, err)
	assert.Equal(t, "Успешно добавлено", result)

	// 2. Тест с пустым WalletID
	testM = models.Wallet{}
	result, err = svc.Create(testM)
	assert.Error(t, err)
	assert.Equal(t, "Wallet empty", result)

	// 3. Тест с невалидным WalletID
	testM = models.Wallet{
		WalletID: "1002-213-213",
	}
	result, err = svc.Create(testM)
	assert.Error(t, err)
	assert.Equal(t, "invalid UUID format", result)
}

func TestChange(t *testing.T) {

	// Создаем тестовую структуру
	testService := &TestService{}
	svc := service.New(testService)

	// 1. Тест Happy
	testM := models.Wallet{
		WalletID:      "4f9a5a4a-2e9b-4d7a-8f4a-2e9b4d7a8f4a",
		OperationType: "DEPOSIT",
		Amount:        100,
	}

	result, err := svc.Change(testM)
	assert.NoError(t, err)
	assert.Equal(t, "Успешно добавлено", result)

	// 2. Тест с пустым WalletID
	testM = models.Wallet{
		OperationType: "DEPOSIT",
		Amount:        100,
	}
	result, err = svc.Change(testM)
	assert.Error(t, err)
	assert.Equal(t, "Wallet empty", result)

	// 3. Тест с пустым OperationType
	testM = models.Wallet{
		WalletID: "4f9a5a4a-2e9b-4d7a-8f4a-2e9b4d7a8f4a",
		Amount:   100,
	}
	result, err = svc.Change(testM)
	assert.Error(t, err)
	assert.Equal(t, "Operation type empty", result)

	// 4. Тест с другим OperationType
	testM = models.Wallet{
		WalletID:      "4f9a5a4a-2e9b-4d7a-8f4a-2e9b4d7a8f4a",
		OperationType: "DELETE",
		Amount:        100,
	}
	result, err = svc.Change(testM)
	assert.Error(t, err)
	assert.Equal(t, "Operation type empty", result)

	// 5. Тест с пустым Amount
	testM = models.Wallet{
		WalletID:      "4f9a5a4a-2e9b-4d7a-8f4a-2e9b4d7a8f4a",
		OperationType: "DEPOSIT",
	}
	result, err = svc.Change(testM)
	assert.Error(t, err)
	assert.Equal(t, "Amount empty or negative", result)

	// 6. Тест с дополнительными данными
	testM = models.Wallet{
		WalletID:      "4f9a5a4a-2e9b-4d7a-8f4a-2e9b4d7a8f4a",
		OperationType: "DEPOSIT",
		Amount:        100,
		Balance:       0,
		UpdatedTime:   time.Now(),
		CreatedTime:   time.Now(),
	}
	result, err = svc.Change(testM)
	assert.NoError(t, err)
	assert.Equal(t, "Успешно добавлено", result)
}

func TestBalance(t *testing.T) {

	testService := &TestService{}

	srv := service.New(testService)

	// 1. Тест Happy
	testM := models.Wallet{
		WalletID: "4f9a5a4a-2e9b-4d7a-8f4a-2e9b4d7a8f4a",
	}

	result, err := srv.Balance(testM)

	assert.NoError(t, err)
	assert.Equal(t, "1 рубль", result)

	// 2. Тест пустой кошелек
	testM = models.Wallet{
		WalletID: "",
	}

	result, err = srv.Balance(testM)

	assert.Error(t, err)
	assert.Equal(t, "Wallet empty", result)

	// 3. Тест пустой кошелек 2
	testM = models.Wallet{}

	result, err = srv.Balance(testM)

	assert.Error(t, err)
	assert.Equal(t, "Wallet empty", result)

	// 4. Тест формат UUID
	testM = models.Wallet{
		WalletID: "123-321",
	}

	result, err = srv.Balance(testM)

	assert.Error(t, err)
	assert.Equal(t, "invalid UUID format", result)

}
