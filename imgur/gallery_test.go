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
		fmt.Fprint(w, mock.ReadJSON("gallery.json"))
	})

	data, _, _ := mock.Client.GalleryService.GetAlbum()

	want := "zz6mGQQ"
	if !reflect.DeepEqual(data[0].ID, want) {
		t.Errorf("GalleryService.GetAlbum returned %+v, want %+v", data[0].ID, want)
	}

	wantLenImages := 2
	if !reflect.DeepEqual(len(data), wantLenImages) {
		t.Errorf("GalleryService.GetAlbum returned %+v, want %+v", len(data), wantLenImages)
	}

}

func TestGalleryService_GetTag(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/3/gallery/t/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("gallery_tag.json"))
	})

	data, _, _ := mock.Client.GalleryService.GetTag("a")

	wantName := "funny"
	fmt.Printf("%v\n", data)
	if !reflect.DeepEqual(data.Name, wantName) {
		t.Errorf("GalleryService.GetTag returned %+v, want %+v", data.Name, wantName)
	}

	wantLenImages := 1
	if !reflect.DeepEqual(len(data.Images), wantLenImages) {
		t.Errorf("GalleryService.GetTag returned %+v, want %+v", len(data.Images), wantLenImages)
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

	_, _, err := mock.Client.GalleryService.GetAlbum()

	if err == nil {
		t.Errorf("GalleryService.GetAlbum returned %+v, want error message.", err)
	}
}
