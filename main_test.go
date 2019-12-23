package main

import "testing"

func TestGetResource(t *testing.T) {
	t.Run("get Resource for simple image", func(t *testing.T) {
		urlPath := "/wa-data/public/site/themes/hypermarket/img/logo.png"
		want := Resource{
			Url: urlPath,
			OriginalUrl: urlPath,
			Ext: ".png",
			Mode: "image",
			Host: "http://wdevel.site",
		}
		checkResource(urlPath, want, t)
	})

	t.Run("get Resource for product image", func(t *testing.T) {
		urlPath := "/wa-data/public/shop/products/80/00/80/images/53/53.750.JPEG"
		want := Resource{
			Url: urlPath,
			OriginalUrl: "/wa-data/protected/shop/products/80/00/80/images/53.JPEG",
			Ext: ".jpeg",
			Mode: "product",
			Host: "http://wdevel.site",
			Cache:true,
			ImageSize: ImageSize{
				Width:  750,
				Height: 750,
				Mode:   "max",
			},
		}
		checkResource(urlPath, want, t)
	})

	t.Run("get Resource for CSS file", func(t *testing.T) {
		urlPath := "/wa-data/public/site/themes/hypermarket/css/custom.css"
		want := Resource{
			Mode:        "css",
			Url:         urlPath,
			OriginalUrl: urlPath,
			Ext:         ".css",
			Host:        "http://wdevel.site",
			Cache:       false,
			ImageSize:   ImageSize{},
		}
		checkResource(urlPath, want, t)
	})

	t.Run("get Resource for JS file", func(t *testing.T) {
		urlPath := "/wa-data/public/site/themes/hypermarket/js/custom.js"
		want := Resource{
			Mode:        "js",
			Url:         urlPath,
			OriginalUrl: urlPath,
			Ext:         ".js",
			Host:        "http://wdevel.site",
			Cache:       false,
			ImageSize:   ImageSize{},
		}
		checkResource(urlPath, want, t)
	})
}

func checkResource(urlPath string, want Resource, t *testing.T) {
	t.Helper()
	got := GetResource(urlPath)
	if got != want {
		t.Errorf("want %+v, \ngot %+v for %s", want, got, urlPath)
	}
}

func TestGetOriginalHost(t *testing.T) {
	got := GetOriginalHost("test")
	want := "http://wdevel.site"
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

