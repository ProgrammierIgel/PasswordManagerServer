package addnewaccount

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

//CREATE BY COPILOT

type mockStore struct {
	AddNewAccountFunc func(accountName, password string) error
}

// AddNewPassword implements storage.Store.
func (m *mockStore) AddNewPassword(token string, passwordName string, passwordToAdd string, url string, username string) error {
	panic("unimplemented")
}

// ChangePassword implements storage.Store.
func (m *mockStore) ChangePassword(token string, passwordName string, newSecret string) error {
	panic("unimplemented")
}

// ChangePasswordName implements storage.Store.
func (m *mockStore) ChangePasswordName(token string, passwordName string, newPasswordName string) error {
	panic("unimplemented")
}

// ChangeURL implements storage.Store.
func (m *mockStore) ChangeURL(token string, passwordName string, newURL string) error {
	panic("unimplemented")
}

// ChangeUsername implements storage.Store.
func (m *mockStore) ChangeUsername(token string, passwordName string, newUsername string) error {
	panic("unimplemented")
}

// CheckPassword implements storage.Store.
func (m *mockStore) CheckPassword(account string, password string, token string) error {
	panic("unimplemented")
}

// CheckToken implements storage.Store.
func (m *mockStore) CheckToken(token string) bool {
	panic("unimplemented")
}

// CreateToken implements storage.Store.
func (m *mockStore) CreateToken(accountName string, masterpassword string, extraParam string) (string, error) {
	panic("unimplemented")
}

// DeleteAccount implements storage.Store.
func (m *mockStore) DeleteAccount(name string, token string) error {
	panic("unimplemented")
}

// DeletePassword implements storage.Store.
func (m *mockStore) DeletePassword(token string, passwordName string) error {
	panic("unimplemented")
}

// DevalueAllTokens implements storage.Store.
func (m *mockStore) DevalueAllTokens(password string) error {
	panic("unimplemented")
}

// DevalueAllTokensOfAccount implements storage.Store.
func (m *mockStore) DevalueAllTokensOfAccount(token string) error {
	panic("unimplemented")
}

// DevalueToken implements storage.Store.
func (m *mockStore) DevalueToken(token string) {
	panic("unimplemented")
}

// DisableSync implements storage.Store.
func (m *mockStore) DisableSync(password string) (bool, error) {
	panic("unimplemented")
}

// EnableSync implements storage.Store.
func (m *mockStore) EnableSync(password string) (bool, error) {
	panic("unimplemented")
}

// GetAllPasswordNamesOfAccount implements storage.Store.
func (m *mockStore) GetAllPasswordNamesOfAccount(token string) ([]string, error) {
	panic("unimplemented")
}

// GetPassword implements storage.Store.
func (m *mockStore) GetPassword(token string, passwordName string) (string, error) {
	panic("unimplemented")
}

// GetURL implements storage.Store.
func (m *mockStore) GetURL(token string, passwordName string) (string, error) {
	panic("unimplemented")
}

// GetUsername implements storage.Store.
func (m *mockStore) GetUsername(token string, passwordName string) (string, error) {
	panic("unimplemented")
}

// IsSyncDisabled implements storage.Store.
func (m *mockStore) IsSyncDisabled() bool {
	panic("unimplemented")
}

// SyncFromFile implements storage.Store.
func (m *mockStore) SyncFromFile() error {
	panic("unimplemented")
}

// SyncToFile implements storage.Store.
func (m *mockStore) SyncToFile() error {
	panic("unimplemented")
}

func (m *mockStore) AddNewAccount(accountName, password string) error {
	return m.AddNewAccountFunc(accountName, password)
}

func TestHandle(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockStoreFunc  func(accountName, password string) error
		expectedStatus int
	}{
		{
			name: "Successful account creation",
			requestBody: map[string]string{
				"AccountName": "testAccount",
				"Password":    "testPassword",
			},
			mockStoreFunc: func(accountName, password string) error {
				return nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid JSON in request body",
			requestBody:    "invalid-json",
			mockStoreFunc:  nil,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Error from store",
			requestBody: map[string]string{
				"AccountName": "testAccount",
				"Password":    "testPassword",
			},
			mockStoreFunc: func(accountName, password string) error {
				return errors.New("store error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requestBodyBytes []byte
			if tt.requestBody != nil {
				requestBodyBytes, _ = json.Marshal(tt.requestBody)
			}

			req := httptest.NewRequest(http.MethodPost, "/addNewAccount", bytes.NewReader(requestBodyBytes))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			mockStore := &mockStore{
				AddNewAccountFunc: tt.mockStoreFunc,
			}

			handler := Handle(mockStore)
			handler(rec, req, httprouter.Params{})

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}
		})
	}
}
