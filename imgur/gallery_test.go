package imgur

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGalleryService_GetAlbum(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/3/gallery", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("gallery_album.json"))
	})

	data, _ := mock.Client.GalleryService.GetAlbum()

	want := "lDRB2"
	if !reflect.DeepEqual(data.ID, want) {
		t.Errorf("GalleryService.Default returned %+v, want %+v", data, want)
	}

	wantLenImages := 11
	if !reflect.DeepEqual(len(data.Images), wantLenImages) {
		t.Errorf("GalleryService.Default returned %+v, want %+v", len(data.Images), wantLenImages)
	}
}

func TestGalleryService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("gallery", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, err := mock.Client.GalleryService.Default()

	if err == nil {
		t.Errorf("GalleryService.Default returned %+v, want error message.", err)
	}
}
