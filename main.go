package main

import (
	"fmt"
	"github.com/davidbyttow/govips/pkg/vips"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
	"net/http"
	"path"
	"regexp"
	"strings"
)

type Resource struct {
	Mode         string
	Url          string
	OriginalUrl  string
	Ext          string
	Host         string
	Cache        bool
	ImageSize    ImageSize
	Token		 string
}

const productRe string = `(?i)((?:\d{2}/){2}([0-9]+)/images/)([0-9]+)/([a-zA-Z0-9_\.-]+)\.(\d+(?:x\d+)?)(@2x)?\.([a-z]{3,4})`

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:4444", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	resource := GetResource(r.URL.Path)
	resource.Token = GetToken(r)
	resp, done := GetUrl(resource.Host + resource.OriginalUrl, w)
	if done {
		return
	}

	switch resource.Mode {
	case "css":
		m := minify.New()
		m.AddFunc("text/css", css.Minify)
		if err := m.Minify("text/css", w, resp.Body); err != nil {
			panic(err)
		}
	case "js":
		m := minify.New()
		m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
		if err := m.Minify("text/javascript", w, resp.Body); err != nil {
			panic(err)
		}
	case "image":
	case "product":
		vips.NewTransform().
			Load(resp.Body).
			ResizeStrategy(vips.ResizeStrategyCrop).
			Resize(resource.ImageSize.Width, resource.ImageSize.Height).
			Quality(90).
			Output(w).
			Apply()
	}
	defer resp.Body.Close()

	//fmt.Printf("%+v\n", resource)
}

func GetUrl(url string, w http.ResponseWriter) (*http.Response, bool) {
	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("failed to get %s: %v", url, err)))
		return nil, true
	}
	if resp.StatusCode/100 != 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("failed to get %s: status %d", url, resp.StatusCode)))
		return nil, true
	}

	return resp, false
}

func GetResource(urlPath string) Resource {
	ext := strings.ToLower(path.Ext(urlPath))
	mode := "other"

	switch ext {
	case ".jpg":
	case ".jpeg":
	case ".png":
		mode = "image"
	case ".js":
		mode = "js"
	case ".css":
		mode = "css"
	}

	resource := Resource{
		Mode: 		 mode,
		Url:  		 urlPath,
		OriginalUrl: urlPath,
		Ext:  		 ext,
		Host: 		 GetOriginalHost("test"),
	}

	if len(regexp.MustCompile(productRe).FindStringIndex(urlPath)) > 0 {
		image := ParseProductImageUrl(urlPath)
		resource.Mode = "product"
		resource.ImageSize = ParseSize(image.Size, image.Dpr)
		resource.Cache = true
		resource.OriginalUrl = GetOriginalProductImagePath(image)
	}

	//fmt.Printf("%+v\n", resource)

	return resource
}

func GetOriginalHost(token string) string {
	return "http://wdevel.site"
}

func GetToken(r *http.Request) string {
	return "test"
}
