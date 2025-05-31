package inmemory_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/programmierigel/pwmanager/storage/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("Assert its not nil.", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		assert.NotNil(t, store)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}

func TestAddNewAccount(t *testing.T) {
	t.Run("Creates an account in system", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store1 := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		err = store1.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store1.CreateToken(accountName, accountPassword, "")

		assert.Nil(t, err)
		assert.NotNil(t, token)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})

	t.Run("throws an error if account already exists", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)

		err = store.AddNewAccount(accountName, accountPassword)
		assert.Error(t, err, "account already exists")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}

func TestDeleteAccount(t *testing.T) {
	t.Run("throws an error due create a new token on remove account", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)

		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.NotNil(t, token)
		assert.Nil(t, err)
		err = store.DeleteAccount(accountName, token)
		assert.Nil(t, err)
		err = store.CheckPassword(accountName, accountPassword, "")
		assert.Error(t, err, "unknown password")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})

	t.Run("throws an error if token is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token := "incorrect token"
		err = store.DeleteAccount(accountName, token)
		assert.Error(t, err, "incorrect token")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})

	t.Run("throws an error due delete other account", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)

		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.NotNil(t, token)
		assert.Nil(t, err)
		err = store.DeleteAccount("other account", token)
		assert.Error(t, err, "cant delete other account")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
}

func TestAddNewPassword(t *testing.T) {
	t.Run("Creates a new password", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		passwordToAdd := "secondExamplePassword"
		passwordName := "examplePasswordname"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		err = store.AddNewPassword(token, passwordName, passwordToAdd, "https://www.google.com", "exampeUsername")
		assert.Nil(t, err)
		password, err := store.GetPassword(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, passwordToAdd, password)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})

	t.Run("Throws an error if account already exists", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		passwordToAdd := "secondExamplePassword"
		passwordName := "examplePasswordname"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		err = store.AddNewPassword(token, passwordName, passwordToAdd, "https://www.google.com", "exampeUsername")
		assert.Nil(t, err)
		err = store.AddNewPassword(token, passwordName, passwordToAdd, "https://www.google.com", "exampeUsername")
		assert.Error(t, err, "already exists")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})

	t.Run("Throws an error if token incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		passwordToAdd := "secondExamplePassword"
		passwordName := "examplePasswordname"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token := "incorrect token"
		assert.Nil(t, err)
		err = store.AddNewPassword(token, passwordName, passwordToAdd, "https://www.google.com", "exampeUsername")
		assert.Error(t, err, "incorrect token")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
}

func TestDeletePassword(t *testing.T) {
	t.Run("Check if account after deleting doesn't exists", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		passwordToAdd := "secondExamplePassword"
		passwordName := "examplePasswordname"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		err = store.AddNewPassword(token, passwordName, passwordToAdd, "https://www.google.com", "exampeUsername")
		assert.Nil(t, err)
		err = store.DeletePassword(token, passwordName)
		assert.Nil(t, err)
		password, err := store.GetPassword(token, passwordName)
		assert.Equal(t, password, "")
		assert.Error(t, err, "password on account not found")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	t.Run("throws in error if password is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		passwordToAdd := "secondExamplePassword"
		passwordName := "examplePasswordname"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		err = store.AddNewPassword(token, passwordName, passwordToAdd, "https://www.google.com", "exampeUsername")
		assert.Nil(t, err)
		token = "incorrect token"
		err = store.DeletePassword(token, passwordName)
		assert.Error(t, err, "incorrect token")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})

	t.Run("Throws an error if to deleting account doesn't exists", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName := "exampleAccount"
		accountPassword := "examplePassword"
		passwordName := "examplePasswordname"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		err = store.DeletePassword(token, passwordName)
		assert.Error(t, err, "passwordname does not exist")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
}

func TestGetPassword(t *testing.T) {
	t.Run("Returns the password if token is correct and the password exists.", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, pasword := "anyPasswordName", "anyPassword"
		err = store.AddNewPassword(token, passwordName, pasword, "", "")
		assert.Nil(t, err)
		gettedPassword, err := store.GetPassword(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedPassword, pasword)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Returns an error if the token is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, pasword := "anyPasswordName", "anyPassword"
		err = store.AddNewPassword(token, passwordName, pasword, "", "")
		assert.Nil(t, err)
		token = "anyIncorrectToken"
		gettedPassword, err := store.GetPassword(token, passwordName)
		assert.Error(t, err, "incorrect token")
		assert.Equal(t, gettedPassword, "")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Returns an error if the getted password doesn't exists", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, pasword := "anyPasswordName", "anyPassword"
		err = store.AddNewPassword(token, passwordName, pasword, "", "")
		assert.Nil(t, err)

		gettedPassword, err := store.GetPassword(token, "any unknown password")
		assert.Error(t, err, "password on account not found")
		assert.Equal(t, gettedPassword, "")
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
}

func TestGetURL(t *testing.T) {
	t.Run("Outputs the URL if token and passwordname are correct.", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, passwordURL := "anyPasswordName", "anyURL"
		err = store.AddNewPassword(token, passwordName, "", passwordURL, "")
		assert.Nil(t, err)
		gettedURL, err := store.GetURL(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, passwordURL, gettedURL)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Throws an error if token is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, passwordURL := "anyPasswordName", "anyURL"
		err = store.AddNewPassword(token, passwordName, "", passwordURL, "")
		assert.Nil(t, err)
		token = "any incorrect token"
		gettedURL, err := store.GetURL(token, passwordName)
		assert.Error(t, err, "incorrect token")
		assert.Equal(t, gettedURL, "")
		assert.NotEqual(t, gettedURL, passwordURL)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Throws an error if password doesn't exists.", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, passwordURL := "anyPasswordName", "anyURL"
		err = store.AddNewPassword(token, passwordName, "", passwordURL, "")
		assert.Nil(t, err)
		gettedURL, err := store.GetURL(token, "any incorrect passwordName")
		assert.Error(t, err, "url on account not found")
		assert.Equal(t, gettedURL, "")
		assert.NotEqual(t, gettedURL, passwordURL)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}
func TestGetUsername(t *testing.T) {
	t.Run("Returns the Username if token and passwordname are correct", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, username := "anyPasswordName", "anyUsername"
		err = store.AddNewPassword(token, passwordName, "", "", username)
		assert.Nil(t, err)
		gettedUsername, err := store.GetUsername(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedUsername, username)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	t.Run("Throws an error if token is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, username := "anyPasswordName", "anyUsername"
		err = store.AddNewPassword(token, passwordName, "", "", username)
		assert.Nil(t, err)
		token = "any incorrect token"
		gettedUsername, err := store.GetUsername(token, passwordName)
		assert.Error(t, err, "incorrect password")
		assert.Equal(t, gettedUsername, "")
		assert.NotEqual(t, gettedUsername, username)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})

	t.Run("Throws an error if passwordname isnt registered", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, username := "anyPasswordName", "anyUsername"
		err = store.AddNewPassword(token, passwordName, "", "", username)
		assert.Nil(t, err)
		passwordName = "any not registered passwordname"
		gettedUsername, err := store.GetUsername(token, passwordName)
		assert.Error(t, err, "username on account not found")
		assert.Equal(t, gettedUsername, "")
		assert.NotEqual(t, gettedUsername, username)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}

func TestGetAllPasswordNamesOfAccount(t *testing.T) {
	t.Run("Returns all registered passwordnames on account", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		somePasswordNames := []string{
			"anyPasswordName1",
			"anyPasswordName2",
			"anyPasswordName3",
			"anyPasswordName4",
			"anyPasswordName5",
		}
		for _, passwordName := range somePasswordNames {
			err = store.AddNewPassword(token, passwordName, "", "", "")
			assert.Nil(t, err)
		}
		gettedPasswordNames, err := store.GetAllPasswordNamesOfAccount(token)
		assert.Nil(t, err)
		assert.Equal(t, len(gettedPasswordNames), len(somePasswordNames))
		for _, gettedPasswordName := range gettedPasswordNames {
			isElementInList := func() bool {
				for _, e := range somePasswordNames {
					if e == gettedPasswordName {
						return true
					}

				}
				return false
			}()

			assert.True(t, isElementInList)
		}

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Throws an error if token is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName := "any PasswordName"
		err = store.AddNewPassword(token, passwordName, "", "", "")
		assert.Nil(t, err)
		token = "any invalid token"
		gettedPasswordNames, err := store.GetAllPasswordNamesOfAccount(token)
		assert.Error(t, err, "incorrect token")
		assert.Empty(t, gettedPasswordNames)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Returns a empty list if no account is registered", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)
		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		gettedPasswordNames, err := store.GetAllPasswordNamesOfAccount(token)
		assert.Nil(t, err)
		assert.Empty(t, gettedPasswordNames)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})

}

func TestDisableSync(t *testing.T) {
	t.Run("Disables syncronization", func(t *testing.T) {
		path := "."
		storePassword := "anyPassword"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, storePassword, logger)

		status, err := store.DisableSync(storePassword)
		assert.Nil(t, err)
		assert.True(t, status)
		status = store.IsSyncDisabled()
		assert.True(t, status)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	t.Run("Throws an error if store password is incorrect", func(t *testing.T) {
		path := "."
		storePassword := "anyPassword"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, storePassword, logger)

		status := store.IsSyncDisabled()
		assert.False(t, status)

		wrongPassword := "any wrong password"
		status, err = store.DisableSync(wrongPassword)
		assert.False(t, status)
		assert.Error(t, err, "wrong password")

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Returns true if sync is disabled", func(t *testing.T) {
		path := "."
		password := "any password"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		status, err := store.DisableSync(password)
		assert.True(t, status)
		assert.Nil(t, err)
		status, err = store.DisableSync(password)
		assert.True(t, status)
		assert.Nil(t, err)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}
func TestIsSyncDisabled(t *testing.T) {
	t.Run("Check if sync is disabled after deactivation", func(t *testing.T) {
		path := "."
		password := "any password"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		status, err := store.DisableSync(password)
		assert.Nil(t, err)
		assert.True(t, status)

		gettedStatus := store.IsSyncDisabled()
		assert.Nil(t, err)
		assert.True(t, status)
		assert.Equal(t, gettedStatus, status)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Check if sync is enabled after activation", func(t *testing.T) {
		path := "."
		password := "any password"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		status, err := store.EnableSync(password)
		assert.Nil(t, err)
		assert.False(t, status)

		gettedStatus := store.IsSyncDisabled()
		assert.Nil(t, err)
		assert.False(t, status)
		assert.Equal(t, gettedStatus, status)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Returns that sync is enabled after initalizing from store", func(t *testing.T) {
		path := "."
		password := "any password"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		status := store.IsSyncDisabled()
		assert.False(t, status)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}
func TestEnableSync(t *testing.T) {
	t.Run("sync is enabled if its triggered twice", func(t *testing.T) {
		path := "."
		password := "any password"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		status, err := store.DisableSync(password)
		assert.Nil(t, err)
		assert.True(t, status)

		status = store.IsSyncDisabled()
		assert.True(t, status)

		status, err = store.EnableSync(password)
		assert.Nil(t, err)
		assert.False(t, status)
		status, err = store.EnableSync(password)
		assert.Nil(t, err)
		assert.False(t, status)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("activates the sync if its deactivated", func(t *testing.T) {
		path := "."
		password := "any password"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		status, err := store.DisableSync(password)
		assert.Nil(t, err)
		assert.True(t, status)

		status = store.IsSyncDisabled()
		assert.True(t, status)

		status, err = store.EnableSync(password)
		assert.Nil(t, err)
		assert.False(t, status)
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	t.Run("returns an error if password is incorrect", func(t *testing.T) {
		path := "."
		password := "any password"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		incorrectPassword := "any incorrect password"

		status, err := store.DisableSync(password)
		assert.Nil(t, err)
		assert.True(t, status)

		status = store.IsSyncDisabled()
		assert.True(t, status)

		status, err = store.EnableSync(incorrectPassword)
		assert.True(t, status)
		assert.Error(t, err, "wrong password")

		status = store.IsSyncDisabled()
		assert.True(t, status)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Let the syncronization active if its triggered after initalizing", func(t *testing.T) {
		path := "."
		password := "any password"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		status := store.IsSyncDisabled()
		assert.False(t, status)

		status, err = store.EnableSync(password)
		assert.Nil(t, err)
		assert.False(t, status)

		status = store.IsSyncDisabled()
		assert.False(t, status)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}

func TestChangeUsername(t *testing.T) {
	t.Run("Changes the username if token and passwordname are correct", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, username := "anyPasswordName", "anyUsername"
		err = store.AddNewPassword(token, passwordName, "", "", username)
		assert.Nil(t, err)
		gettedUsername, err := store.GetUsername(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedUsername, username)
		newUsername := "newUsername"
		err = store.ChangeUsername(token, passwordName, newUsername)
		assert.NotEqual(t, username, newUsername)
		assert.Nil(t, err)
		gettedUsername, err = store.GetUsername(token, passwordName)
		assert.Equal(t, gettedUsername, newUsername)
		assert.NotEqual(t, gettedUsername, username)
		assert.Nil(t, err)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Throws an error if token is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, username := "anyPasswordName", "anyUsername"
		err = store.AddNewPassword(token, passwordName, "", "", username)
		assert.Nil(t, err)
		gettedUsername, err := store.GetUsername(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedUsername, username)
		newUsername := "newUsername"
		invalidToken := "invalid token"
		err = store.ChangeUsername(invalidToken, passwordName, newUsername)
		assert.NotEqual(t, username, newUsername)
		assert.Error(t, err, "token incorrect")
		gettedUsername, err = store.GetUsername(token, passwordName)
		assert.NotEqual(t, gettedUsername, newUsername)
		assert.Equal(t, gettedUsername, username)
		assert.Nil(t, err)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	t.Run("Throws an error if passwordname is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, username := "anyPasswordName", "anyUsername"
		err = store.AddNewPassword(token, passwordName, "", "", username)
		assert.Nil(t, err)
		gettedUsername, err := store.GetUsername(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedUsername, username)
		newUsername := "newUsername"
		invalidPasswordName := "invalid passwordname"
		err = store.ChangeUsername(token, invalidPasswordName, newUsername)
		assert.NotEqual(t, username, newUsername)
		assert.Error(t, err, "username on account not found")
		gettedUsername, err = store.GetUsername(token, passwordName)
		assert.NotEqual(t, gettedUsername, newUsername)
		assert.Equal(t, gettedUsername, username)
		assert.Nil(t, err)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
}
func TestChangeURL(t *testing.T) {
	t.Run("changes the url if token and passwordname are correct", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, URL := "anyPasswordName", "anyURL"
		err = store.AddNewPassword(token, passwordName, "", URL, "")
		assert.Nil(t, err)
		gettedURL, err := store.GetURL(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedURL, URL)
		newURL := "newURL"
		err = store.ChangeURL(token, passwordName, newURL)
		assert.NotEqual(t, URL, newURL)
		assert.Nil(t, err)
		gettedURL, err = store.GetURL(token, passwordName)
		assert.Equal(t, gettedURL, newURL)
		assert.NotEqual(t, gettedURL, URL)
		assert.Nil(t, err)
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	t.Run("Throws an error if passwordname is undefined", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, URL := "anyPasswordName", "anyURL"
		err = store.AddNewPassword(token, passwordName, "", URL, "")
		assert.Nil(t, err)
		gettedURL, err := store.GetURL(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedURL, URL)
		newURL := "newURL"
		invalidPasswordName := "invalidPasswordName"
		err = store.ChangeURL(token, invalidPasswordName, newURL)
		assert.NotEqual(t, URL, newURL)
		assert.Error(t, err, "url on account not found")
		gettedURL, err = store.GetURL(token, passwordName)
		assert.NotEqual(t, gettedURL, newURL)
		assert.Equal(t, gettedURL, URL)
		assert.Nil(t, err)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Throws an error if token is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "any AccountName", "any Password"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, URL := "anyPasswordName", "anyURL"
		err = store.AddNewPassword(token, passwordName, "", URL, "")
		assert.Nil(t, err)
		gettedURL, err := store.GetURL(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedURL, URL)
		newURL := "newURL"
		invalidToken := "invalid token"
		err = store.ChangeURL(invalidToken, passwordName, newURL)
		assert.NotEqual(t, URL, newURL)
		assert.Error(t, err, "token incorrect")
		gettedURL, err = store.GetURL(token, passwordName)
		assert.NotEqual(t, gettedURL, newURL)
		assert.Equal(t, gettedURL, URL)
		assert.Nil(t, err)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}

func TestChangePassword(t *testing.T) {
	t.Run("Changes the password if token and passwordname are correct", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, pasword := "anyPasswordName", "anyPassword"
		err = store.AddNewPassword(token, passwordName, pasword, "", "")
		assert.Nil(t, err)
		gettedPassword, err := store.GetPassword(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedPassword, pasword)

		newPassword := "new Password"
		assert.NotEqual(t, newPassword, pasword)
		err = store.ChangePassword(token, passwordName, newPassword)
		assert.Nil(t, err)

		gettedPassword, err = store.GetPassword(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedPassword, newPassword)
		assert.NotEqual(t, gettedPassword, pasword)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("returns an error if the passwordname is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, pasword := "anyPasswordName", "anyPassword"
		err = store.AddNewPassword(token, passwordName, pasword, "", "")
		assert.Nil(t, err)
		gettedPassword, err := store.GetPassword(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedPassword, pasword)

		newPassword := "new Password"
		assert.NotEqual(t, newPassword, pasword)
		invalidPasswordName := "invalid password name"
		err = store.ChangePassword(token, invalidPasswordName, newPassword)
		assert.Error(t, err, "username on account not found")

		gettedPassword, err = store.GetPassword(token, passwordName)
		assert.Nil(t, err)
		assert.NotEqual(t, gettedPassword, newPassword)
		assert.Equal(t, gettedPassword, pasword)
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	t.Run("Throws an error if token is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		passwordName, pasword := "anyPasswordName", "anyPassword"
		err = store.AddNewPassword(token, passwordName, pasword, "", "")
		assert.Nil(t, err)
		gettedPassword, err := store.GetPassword(token, passwordName)
		assert.Nil(t, err)
		assert.Equal(t, gettedPassword, pasword)

		newPassword := "new Password"
		assert.NotEqual(t, newPassword, pasword)
		invalidToken := "invalid token"
		err = store.ChangePassword(invalidToken, passwordName, newPassword)
		assert.Error(t, err, "incorrect token")

		gettedPassword, err = store.GetPassword(token, passwordName)
		assert.Nil(t, err)
		assert.NotEqual(t, gettedPassword, newPassword)
		assert.Equal(t, gettedPassword, pasword)
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})

}

// func TestPasswordName(t *testing.T) {
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// }

func TestCreateToken(t *testing.T) {
	t.Run("Returns a token if accountname and password are correct", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		assert.NotNil(t, token)

		correct := store.CheckToken(token)
		assert.True(t, correct)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	t.Run("Throws an error if password is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		incorrectPassword := "incorrectPassword"
		token, err := store.CreateToken(accountName, incorrectPassword, "")
		assert.Error(t, err, "unknown password or account")
		assert.Equal(t, token, "")

		correct := store.CheckToken(token)
		assert.False(t, correct)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	t.Run("Throws an error if account does not exits", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		unknownAccountname := "unknownAccountname"
		token, err := store.CreateToken(unknownAccountname, accountPassword, "")
		assert.Error(t, err, "unkonown password or account")
		assert.Equal(t, token, "")

		correct := store.CheckToken(token)
		assert.False(t, correct)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Returns never one token twice", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token1, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		assert.NotNil(t, token1)
		token2, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)
		assert.NotNil(t, token2)

		correct := store.CheckToken(token1)
		assert.True(t, correct)
		correct = store.CheckToken(token2)
		assert.True(t, correct)

		assert.NotEqual(t, token1, token2)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Returns never one token twice on 2 accounts", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName1, accountPassword1 := "anyAccountName1", "anyPassword1"
		accountName2, accountPassword2 := "anyAccountName2", "anyPassword2"
		err = store.AddNewAccount(accountName1, accountPassword1)
		assert.Nil(t, err)
		err = store.AddNewAccount(accountName2, accountPassword2)
		assert.Nil(t, err)
		token1, err := store.CreateToken(accountName1, accountPassword1, "")
		assert.Nil(t, err)
		assert.NotNil(t, token1)
		token2, err := store.CreateToken(accountName2, accountPassword2, "")
		assert.Nil(t, err)
		assert.NotNil(t, token2)

		correct := store.CheckToken(token1)
		assert.True(t, correct)
		correct = store.CheckToken(token2)
		assert.True(t, correct)

		assert.NotEqual(t, token1, token2)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}

func TestCheckToken(t *testing.T) {
	t.Run("Returns true if token is valid", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)

		correct := store.CheckToken(token)
		assert.True(t, correct)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Returns false if token is invalid", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		invalidToken := "invalidToken"
		valid := store.CheckToken(invalidToken)
		assert.False(t, valid)
		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Returns false if token is devalued", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)

		correct := store.CheckToken(token)
		assert.True(t, correct)

		store.DevalueToken(token)
		correct = store.CheckToken(token)
		assert.False(t, correct)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	t.Run("Returns false if all tokens of account are devalued", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)

		correct := store.CheckToken(token)
		assert.True(t, correct)

		err = store.DevalueAllTokensOfAccount(token)
		assert.Nil(t, err)
		correct = store.CheckToken(token)
		assert.False(t, correct)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("returns false if all tokens are devalued", func(t *testing.T) {
		path := "."
		password := "anyPassword"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)

		correct := store.CheckToken(token)
		assert.True(t, correct)

		err = store.DevalueAllTokens(password)
		assert.Nil(t, err)
		correct = store.CheckToken(token)
		assert.False(t, correct)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})

	// Test fails because test needs to much time
	/*t.Run("Returns false if the time of the token is up", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)

		correct := store.CheckToken(token)
		assert.True(t, correct)
		//// time.Sleep(time.Second*29)
		time.Sleep(time.Hour + time.Second*5)
		correct = store.CheckToken(token)
		assert.False(t, correct)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))

	})*/
}

func TestDevalueToken(t *testing.T) {
	t.Run("Token is devalued", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)
		token, err := store.CreateToken(accountName, accountPassword, "")
		assert.Nil(t, err)

		correct := store.CheckToken(token)
		assert.True(t, correct)

		store.DevalueToken(token)

		correct = store.CheckToken(token)
		assert.False(t, correct)

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Do nothing if token is incorrect", func(t *testing.T) {
		path := "."
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, "", logger)

		accountName, accountPassword := "anyAccountName", "anyPassword"
		err = store.AddNewAccount(accountName, accountPassword)
		assert.Nil(t, err)

		token := "invalid token"
		store.DevalueToken(token)
		assert.False(t, store.CheckToken(token))

		os.Remove(fmt.Sprintf("%s/secrets.json", path))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}

func TestDevalueAllTokens(t *testing.T) {
	t.Run("Devalues all tokens if password is correct", func(t *testing.T) {
		path := "."
		password := "anyPassword"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		accountName1, accountPassword1 := "anyAccountName1", "anyPassword1"
		err = store.AddNewAccount(accountName1, accountPassword1)
		assert.Nil(t, err)
		token1, err := store.CreateToken(accountName1, accountPassword1, "")
		assert.Nil(t, err)
		correct := store.CheckToken(token1)
		assert.True(t, correct)

		accountName2, accountPassword2 := "anyAccountName2", "anyPassword2"
		err = store.AddNewAccount(accountName2, accountPassword2)
		assert.Nil(t, err)
		token2, err := store.CreateToken(accountName2, accountPassword2, "")
		assert.Nil(t, err)
		correct = store.CheckToken(token2)
		assert.True(t, correct)
		token3, err := store.CreateToken(accountName2, accountPassword2, "")
		assert.Nil(t, err)
		correct = store.CheckToken(token3)
		assert.True(t, correct)

		err = store.DevalueAllTokens(password)
		assert.Nil(t, err)
		correct = store.CheckToken(token1)
		assert.False(t, correct)
		correct = store.CheckToken(token2)
		assert.False(t, correct)
		correct = store.CheckToken(token3)
		assert.False(t, correct)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
	t.Run("Throws an error if password is incorrect", func(t *testing.T) {
		path := "."
		password := "anyPassword"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		accountName1, accountPassword1 := "anyAccountName1", "anyPassword1"
		err = store.AddNewAccount(accountName1, accountPassword1)
		assert.Nil(t, err)
		token1, err := store.CreateToken(accountName1, accountPassword1, "")
		assert.Nil(t, err)
		correct := store.CheckToken(token1)
		assert.True(t, correct)

		accountName2, accountPassword2 := "anyAccountName2", "anyPassword2"
		err = store.AddNewAccount(accountName2, accountPassword2)
		assert.Nil(t, err)
		token2, err := store.CreateToken(accountName2, accountPassword2, "")
		assert.Nil(t, err)
		correct = store.CheckToken(token2)
		assert.True(t, correct)
		token3, err := store.CreateToken(accountName2, accountPassword2, "")
		assert.Nil(t, err)
		correct = store.CheckToken(token3)
		assert.True(t, correct)

		incorrectPassword := "incorrectPassword"
		err = store.DevalueAllTokens(incorrectPassword)
		assert.Error(t, err, "")
		assert.NotEqual(t, password, incorrectPassword)

		correct = store.CheckToken(token1)
		assert.True(t, correct)
		correct = store.CheckToken(token2)
		assert.True(t, correct)
		correct = store.CheckToken(token3)
		assert.True(t, correct)
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})
}

func TestDevalueAllTokensOfAccount(t *testing.T) {
	t.Run("Devalues all tokens of account if token is correct", func(t *testing.T) {
		path := "."
		password := "anyPassword"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		accountName1, accountPassword1 := "anyAccountName1", "anyPassword1"
		accountName2, accountPassword2 := "anyAccountName2", "anyPassword2"
		err = store.AddNewAccount(accountName1, accountPassword1)
		assert.Nil(t, err)
		token1, err := store.CreateToken(accountName1, accountPassword1, "")
		assert.Nil(t, err)
		correct := store.CheckToken(token1)
		assert.True(t, correct)
		token2, err := store.CreateToken(accountName1, accountPassword1, "")
		assert.Nil(t, err)
		correct = store.CheckToken(token2)
		assert.True(t, correct)
		token3, err := store.CreateToken(accountName1, accountPassword1, "")
		assert.Nil(t, err)
		correct = store.CheckToken(token3)
		assert.True(t, correct)

		err = store.AddNewAccount(accountName2, accountPassword2)
		assert.Nil(t, err)
		token4, err := store.CreateToken(accountName2, accountPassword2, "")
		assert.Nil(t, err)
		correct = store.CheckToken(token4)
		assert.True(t, correct)

		err = store.DevalueAllTokensOfAccount(token1)
		assert.Nil(t, err)
		assert.False(t, store.CheckToken(token1))
		assert.False(t, store.CheckToken(token2))
		assert.False(t, store.CheckToken(token3))
		assert.True(t, store.CheckToken(token4))
		os.Remove(fmt.Sprintf("%s/test.log", path))
	})

	t.Run("Devalues no token if token is incorret", func(t *testing.T) {
		path := "."
		password := "anyPassword"
		os.Remove(fmt.Sprintf("%s/secrets.json", path))

		f, err := os.OpenFile(fmt.Sprintf("%s/test.log", path),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		store := inmemory.New(path, password, logger)

		accountName1, accountPassword1 := "anyAccountName1", "anyPassword1"
		accountName2, accountPassword2 := "anyAccountName2", "anyPassword2"
		err = store.AddNewAccount(accountName1, accountPassword1)
		assert.Nil(t, err)
		token1, err := store.CreateToken(accountName1, accountPassword1, "")
		assert.Nil(t, err)
		correct := store.CheckToken(token1)
		assert.True(t, correct)
		token2, err := store.CreateToken(accountName1, accountPassword1, "")
		assert.Nil(t, err)
		correct = store.CheckToken(token2)
		assert.True(t, correct)
		token3, err := store.CreateToken(accountName1, accountPassword1, "")
		assert.Nil(t, err)
		correct = store.CheckToken(token3)
		assert.True(t, correct)

		err = store.AddNewAccount(accountName2, accountPassword2)
		assert.Nil(t, err)
		token4, err := store.CreateToken(accountName2, accountPassword2, "")
		assert.Nil(t, err)
		correct = store.CheckToken(token4)
		assert.True(t, correct)
		invalidToken := "invalidToken"
		err = store.DevalueAllTokensOfAccount(invalidToken)
		assert.Error(t, err, "token incorrect")
		assert.True(t, store.CheckToken(token1))
		assert.True(t, store.CheckToken(token2))
		assert.True(t, store.CheckToken(token3))
		assert.True(t, store.CheckToken(token4))

		os.Remove(fmt.Sprintf("%s/test.log", path))

	})
	// t.Run("", func(t *testing.T) {
	// 	path := "."
	// 	os.Remove(fmt.Sprintf("%s/secrets.json", path))
	// 	store := inmemory.New(path, "")
	// })
	// t.Run("", func(t *testing.T) {
	// 	path := "."
	// 	os.Remove(fmt.Sprintf("%s/secrets.json", path))
	// 	store := inmemory.New(path, "")
	// })
}

// func TestSyncToFile(t *testing.T) {
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// }
// func TestSyncFromFile(t *testing.T) {
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// }
// }
// func TestCheckPassword(t *testing.T) {
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// 	t.Run("", func(t *testing.T) {
// 		path := "."
// 		os.Remove(fmt.Sprintf("%s/secrets.json", path))
// 		store := inmemory.New(path, "")
// 	})
// }
