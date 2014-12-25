package imgur

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountService_Account(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/3/account/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("account.json"))
	})

	data, _, _ := mock.Client.AccountService.Account("a")

	wantId := 1
	if !reflect.DeepEqual(data.ID, wantId) {
		t.Errorf("AccountService.Account returned %+v, want %+v", data.ID, wantId)
	}
}

func TestAccountService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/3/account/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, resp, _ := mock.Client.AccountService.Account("a")

	if !reflect.DeepEqual(resp.StatusCode, 401) {
		t.Errorf("AccountService.Account returned %+v, want %+v", resp.StatusCode, 401)
	}

}

func TestAccountService_NotFound(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/3/account/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusNotFound)
	})

	_, resp, _ := mock.Client.AccountService.Account("a")

	if !reflect.DeepEqual(resp.StatusCode, 404) {
		t.Errorf("AccountService.Account returned %+v, want %+v", resp.StatusCode, 404)
	}
}
