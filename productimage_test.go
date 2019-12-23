package main

import "testing"

func TestParseProductImageUrl(t *testing.T) {
	t.Run("without dpr", func(t *testing.T) {
		urlPath := "/wa-data/public/shop/products/80/00/80/images/53/53.750.JPEG"
		got := ParseProductImageUrl(urlPath)
		want := ProductImage{
			ImagePath: "80/00/80/images/",
			ProductId: "80",
			ImageId:   "53",
			ImageName: "53",
			Size:      "750",
			Dpr:       false,
			Ext:       "JPEG",
		}
		if got != want {
			t.Errorf("want %+v, got %+v", want, got)
		}
	})

	t.Run("with dpr", func(t *testing.T) {
		urlPath := "/wa-data/public/shop/products/80/00/80/images/53/53.750@2x.JPEG"
		got := ParseProductImageUrl(urlPath)
		want := ProductImage{
			ImagePath: "80/00/80/images/",
			ProductId: "80",
			ImageId:   "53",
			ImageName: "53",
			Size:      "750",
			Dpr:       true,
			Ext:       "JPEG",
		}
		if got != want {
			t.Errorf("want %+v, got %+v", want, got)
		}
	})
}

func TestGetOriginalProductImagePath(t *testing.T) {
	t.Run("without filename", func(t *testing.T) {
		image := ProductImage{
			ImagePath: "80/00/80/images/",
			ProductId: "80",
			ImageId:   "53",
			ImageName: "53",
			Size:      "750",
			Dpr:       false,
			Ext:       "JPEG",
		}
		got := GetOriginalProductImagePath(image)
		want := "/wa-data/protected/shop/products/80/00/80/images/53.JPEG"
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("with filename", func(t *testing.T) {
		image := ProductImage{
			ImagePath: "80/00/80/images/",
			ProductId: "80",
			ImageId:   "53",
			ImageName: "TestMe",
			Size:      "750",
			Dpr:       false,
			Ext:       "png",
		}
		got := GetOriginalProductImagePath(image)
		want := "/wa-data/protected/shop/products/80/00/80/images/53.TestMe.png"
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

func TestParseSize(t *testing.T) {

	t.Run("max Mode", func(t *testing.T) {
		dpr := false
		size := "750"
		got := ParseSize(size, dpr)
		want := ImageSize{
			Width:  750,
			Height: 750,
			Mode:   "max",
		}
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("crop Mode", func(t *testing.T) {
		dpr := false
		size := "96x96"
		got := ParseSize(size, dpr)
		want := ImageSize{
			Width:  96,
			Height: 96,
			Mode:   "crop",
		}
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("height Mode", func(t *testing.T) {
		dpr := false
		size := "0x200"
		got := ParseSize(size, dpr)
		want := ImageSize{
			Width:  0,
			Height: 200,
			Mode:   "height",
		}
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("rectangle Mode", func(t *testing.T) {
		dpr := false
		size := "100x200"
		got := ParseSize(size, dpr)
		want := ImageSize{
			Width:  100,
			Height: 200,
			Mode:   "rectangle",
		}
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
