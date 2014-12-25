package imgur

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAlbumService_GetAlbum(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/3/album/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("album.json"))
	})

	data, resp, err := mock.Client.AlbumService.GetAlbum("1")

	if err != nil {
		t.Errorf("AlbumService.GetAlbum returned error: %+v", err)
	}

	if !reflect.DeepEqual(resp.StatusCode, http.StatusOK) {
		t.Errorf("AlbumService.GetAlbum returned %+v, want %+v", resp.StatusCode, http.StatusOK)
	}

	wantId := "lDRB2"
	if !reflect.DeepEqual(data.ID, wantId) {
		t.Errorf("AlbumService.GetAlbum returned %+v, want %+v", data.ID, wantId)
	}

	wantImagesSize := 11
	if !reflect.DeepEqual(len(data.Images), wantImagesSize) {
		t.Errorf("AlbumService.GetAlbum returned %+v, want %+v", len(data.Images), wantImagesSize)
	}
}

func TestAlbumService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/3/album/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, resp, _ := mock.Client.AlbumService.GetAlbum("1")

	if !reflect.DeepEqual(resp.StatusCode, 401) {
		t.Errorf("AlbumService.GetAlbum returned %+v, want %+v", resp.StatusCode, 401)
	}

}

func TestAlbumService_NotFound(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/3/album/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusNotFound)
	})

	_, resp, _ := mock.Client.AlbumService.GetAlbum("1")

	if !reflect.DeepEqual(resp.StatusCode, 404) {
		t.Errorf("AlbumService.GetAlbum returned %+v, want %+v", resp.StatusCode, 404)
	}
}
