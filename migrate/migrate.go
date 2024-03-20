package main

import (
	"bank-teller-backend/initializers"
	"bank-teller-backend/models"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	// initializers.DB.AutoMigrate(&models.Customer{}, &models.UserAccount{}, &models.Transaction{}, &models.Teller{}, &models.Role{}, &models.User{}, &models.AccountType{}, &models.TransactionType{}, &models.Admin{})
	// initializers.DB.Migrator().DropTable(&models.Customer{}, &models.Account{}, &models.Transaction{}, &models.Teller{})
	seedAdmins()
}

// func seedRoles() {
// 	roles := []string{"Admin", "Teller", "Customer"}

// 	for _, roleName := range roles {
// 		role := models.Role{Name: roleName}

// 		// Check if the role already exists
// 		var existingRole models.Role
// 		if err := initializers.DB.Where("name = ?", roleName).First(&existingRole).Error; err != nil {
// 			// If the role doesn't exist, create it
// 			initializers.DB.Create(&role)
// 		}
// 	}
// }

// func seedAccountTypes() {
// 	accountTypes := []string{"Savings", "Current", "Fixed Deposit", "Joint", "Corporate"}

// 	for _, accountTypeName := range accountTypes {
// 		accountType := models.AccountType{Name: accountTypeName}

// 		// Check if the account type already exists
// 		var existingAccountType models.AccountType
// 		if err := initializers.DB.Where("name = ?", accountTypeName).First(&existingAccountType).Error; err != nil {
// 			// If the account type doesn't exist, create it
// 			initializers.DB.Create(&accountType)
// 		}
// 	}
// }

// func seedTransactionTypes() {
// 	transactionTypes := []string{"Deposit", "Withdrawal"}

// 	for _, transactionTypeName := range transactionTypes {
// 		transactionType := models.TransactionType{Name: transactionTypeName}

// 		// Check if the transaction type already exists
// 		var existingTransactionType models.TransactionType
// 		if err := initializers.DB.Where("name = ?", transactionTypeName).First(&existingTransactionType).Error; err != nil {
// 			// If the transaction type doesn't exist, create it
// 			initializers.DB.Create(&transactionType)
// 		}
// 	}
// }

func seedAdmins() {
	adminRole := models.Role{Name: "Admin"}
	// Check if the admin role already exists
	if err := initializers.DB.Where("name = ?", "Admin").First(&adminRole).Error; err != nil {
		// If not, create it
		adminRole = models.Role{Name: "Admin"}
		initializers.DB.Create(&adminRole)
	}

	// Hash the password
	password := "yourPassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password")
	}

	// Create a new user with the admin role
	adminUser := models.User{
		FirstName: "Admin",
		LastName:  "User",
		Phone:     "1234567890",
		Email:     "admin@example.com",
		Password:  string(hashedPassword),
		RoleID:    adminRole.ID,
	}
	initializers.DB.Create(&adminUser)

	// Create a new admin with the user ID
	admin := models.Admin{
		UserID: adminUser.ID,
	}
	initializers.DB.Create(&admin)
}
