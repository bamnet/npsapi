package npsapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestListParks(t *testing.T) {
	gotKey := ""
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		switch q["start"][0] {
		case "0":
			gotKey = r.Header.Get("X-Api-Key")
			http.ServeFile(w, r, "testdata/list_parks_1.json")
		case "100":
			http.ServeFile(w, r, "testdata/list_parks_2.json")
		default:
			http.ServeFile(w, r, "testdata/list_parks_n.json")
		}
	}))
	defer ts.Close()

	c := NewClient("apikeyhere")
	c.baseURL = ts.URL

	gotParks, err := c.ListParks(context.Background())
	if err != nil {
		t.Errorf("ListParks() got unexpected error: %v", err)
	}

	if wantKey := "apikeyhere"; gotKey != wantKey {
		t.Errorf("ListParks() API Key want %s, got %s", wantKey, gotKey)
	}

	wantParks := []Park{
		{
			ID:          "1A47416F-DAA3-4137-9F30-14AF86B4E547",
			Name:        "African American Civil War Memorial",
			FullName:    "African American Civil War Memorial",
			Code:        "afam",
			States:      []string{"DC"},
			Location:    &LatLng{38.916554, -77.025977},
			Description: "Over 200,000 African-American soldiers...",
			URL:         "https://www.nps.gov/afam/index.htm",
		}, {
			ID:          "BBC5BE3F-E921-40B5-A850-F259D9FAF88C",
			Name:        "Gateway",
			FullName:    "Gateway National Recreation Area",
			Code:        "gate",
			States:      []string{"NY", "NJ"},
			Location:    &LatLng{40.56536246, -73.9171087},
			Description: "Gateway is a large diverse urban park...",
			URL:         "https://www.nps.gov/gate/index.htm",
		},
	}
	if !reflect.DeepEqual(gotParks, wantParks) {
		t.Errorf("ListParks() got %+v, want %+v", gotParks, wantParks)
	}
}

func TestGetPark(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q["parkCode"][0] == "missing" {
			http.ServeFile(w, r, "testdata/get_missing_park.json")
			return
		}
		http.ServeFile(w, r, "testdata/get_park.json")
	}))
	defer ts.Close()

	c := NewClient("apikeyhere")
	c.baseURL = ts.URL

	for _, tc := range []struct {
		parkCode string
		wantErr  error
		wantPark *Park
	}{
		{
			parkCode: "yell",
			wantErr:  nil,
			wantPark: &Park{
				ID:          "F58C6D24-8D10-4573-9826-65D42B8B83AD",
				Name:        "Yellowstone",
				FullName:    "Yellowstone National Park",
				Code:        "yell",
				States:      []string{"ID", "MT", "WY"},
				Location:    &LatLng{44.59824417, -110.5471695},
				Description: "Visit Yellowstone and experience the world's first national park. Marvel at a volcano's hidden power rising up in colorful hot springs, mudpots, and geysers. Explore mountains, forests, and lakes to watch wildlife and witness the drama of the natural world unfold. Discover the history that led to the conservation of our national treasures 'for the benefit and enjoyment of the people.'",
				URL:         "https://www.nps.gov/yell/index.htm",
			},
		}, {
			parkCode: "missing",
			wantErr:  ErrParkNotFound,
			wantPark: nil,
		},
	} {
		gotPark, gotErr := c.GetPark(context.Background(), tc.parkCode)
		if gotErr != tc.wantErr {
			t.Errorf("GetPark(%s) unexpected error, got %v, want %v", tc.parkCode, gotErr, tc.wantErr)
		}
		if !reflect.DeepEqual(gotPark, tc.wantPark) {
			t.Errorf("GetPark(%s) got %v, want %v", tc.parkCode, gotPark, tc.wantPark)
		}
	}
}
