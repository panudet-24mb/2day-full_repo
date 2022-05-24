package repositories

import "gorm.io/gorm"

type BankAccount struct {
	AccountID     int `gorm:"primary_key;autoIncrement;not_null"`
	AccountType   int
	AccountNumber string
	AccountName   string
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	// Delete(id string) error
	// FindAll() (bankAccounts []BankAccount, err error)
	// FindByID(id string) (bankAccount BankAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.Table("accounts").AutoMigrate(&BankAccount{})
	return accountRepository{db}
}

func (obj accountRepository) Save(bankAccount BankAccount) error {
	return obj.db.Table("accounts").Save(&bankAccount).Error
}

// func (obj accountRepository) Delete(id string) error {
// 	return obj.db.Table("accounts").Where("id=?", id).Delete(&BankAccount{}).Error
// }

// func (obj accountRepository) FindAll() (bankAccounts []BankAccount, err error) {
// 	err = obj.db.Table("accounts").Find(&bankAccounts).Error
// 	return bankAccounts, err
// }

// func (obj accountRepository) FindByID(id string) (bankAccount BankAccount, err error) {
// 	err = obj.db.Table("accounts").Where("id=?", id).First(&bankAccount).Error
// 	return bankAccount, err
// }
