package core_test

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/abiquo/ojal/core"
)

type battery []*struct {
	Message  string
	Current  interface{}
	Expected interface{}
}

func (b battery) Run(name string, t *testing.T) {
	for _, v := range b {
		if v.Current != v.Expected {
			t.Errorf("%v %v ? %v != %v", name, v.Message, v.Current, v.Expected)
		}
	}
}

type dto struct {
	Name string `json:"name"`
	core.DTO
}

var (
	basic       core.Basic
	oauth       core.Oauth
	environment map[string]string
	none        *core.Link
	self        *core.Link
	edit        *core.Link
	enterprise  core.Linker
	enterprises core.Linker
	name0       string
	name1       string
	enterprise0 *dto
	enterprise1 *dto
	result      *dto
)

func init() {
	environment = map[string]string{
		"ABQ_DISK":         os.Getenv("ABQ_DISK"),
		"ABQ_TEMPLATES":    os.Getenv("ABQ_TEMPLATES"),
		"ABQ_ENDPOINT":     os.Getenv("ABQ_ENDPOINT"),
		"ABQ_USERNAME":     os.Getenv("ABQ_USERNAME"),
		"ABQ_PASSWORD":     os.Getenv("ABQ_PASSWORD"),
		"ABQ_TOKEN":        os.Getenv("ABQ_TOKEN"),
		"ABQ_TOKEN_SECRET": os.Getenv("ABQ_TOKEN_SECRET"),
		"ABQ_API_SECRET":   os.Getenv("ABQ_API_SECRET"),
		"ABQ_API_KEY":      os.Getenv("ABQ_API_KEY"),
		"ABQ_OVA":          os.Getenv("ABQ_OVA"),
	}

	for k, v := range environment {
		if v == "" {
			panic(k + " environment variable should not be empty")
		}
	}

	basic = core.Basic{
		Username: environment["ABQ_USERNAME"],
		Password: environment["ABQ_PASSWORD"],
	}

	oauth = core.Oauth{
		APIKey:      environment["ABQ_API_KEY"],
		APISecret:   environment["ABQ_API_SECRET"],
		TokenSecret: environment["ABQ_TOKEN_SECRET"],
		Token:       environment["ABQ_TOKEN"],
	}
}

func newDTO() core.Resource { return new(dto) }

func TestInit(t *testing.T) {
	err0 := core.Init("https://fail:443/api", core.Basic{})
	err1 := core.Init(environment["ABQ_ENDPOINT"], core.Basic{})
	err2 := core.Init(environment["ABQ_ENDPOINT"], oauth)
	err3 := core.Init(environment["ABQ_ENDPOINT"], basic)

	battery{
		{"err0", err0 == nil, false},
		{"err1", err1 == nil, false},
		{"err2", err2 == nil, true},
		{"err3", err3 == nil, true},
	}.Run("Init", t)

	ts := time.Now().Unix()
	name0 = fmt.Sprint("ztest A ", ts)
	name1 = fmt.Sprint("ztest B ", ts)
	none = core.NewLinkType("none", "none").SetRel("none").SetTitle("none")
	self = core.NewLinkType("self", "self").SetRel("self").SetTitle("self")
	edit = core.NewLinkType("edit", "edit").SetRel("edit").SetTitle("edit")
	enterprise = core.NewLinker("admin/enterprises", "enterprise")
	enterprises = core.NewLinker("admin/enterprises", "enterprises")
	enterprise0 = &dto{Name: name0}
	enterprise1 = &dto{Name: name1}
	result = &dto{}
	core.RegisterMedia("enterprise", "enterprises", newDTO)
	core.RegisterMedia("user", "users", newDTO)
}

func TestLink(t *testing.T) {
	battery{
		{"URL()", none.URL(), core.Resolve("none", nil)},
		{"Media()", none.Media(), core.Media("none")},
		{"Title", none.Title, "none"},
		{"Rel", none.Rel, "none"},
	}.Run("link", t)
}

func TestObject(t *testing.T) {
	object := make(core.Object)

	battery{
		{"Href()", object.Href(), ""},
		{"Media()", object.Media(), core.Media("")},
		{"len(Links())", len(object.Links()), 0},
	}.Run("zero", t)

	object.Add(none)
	battery{
		{"Href()", object.Href(), ""},
		{"Media()", object.Media(), core.Media("")},
		{"len(Links())", len(object.Links()), 1},
	}.Run("zero", t)

	object.Add(self)
	battery{
		{"Href()", object.Href(), self.Href},
		{"Media()", object.Media(), self.Media()},
		{"len(Links())", len(object.Links()), 2},
	}.Run("self", t)

	object.Add(edit)
	battery{
		{"Href()", object.Href(), edit.Href},
		{"Media()", object.Media(), edit.Media()},
		{"len(Links())", len(object.Links()), 3},
	}.Run("edit", t)
}

func TestCall(t *testing.T) {
	post, err := core.Rest(result, core.Post(
		"admin/enterprises",
		"enterprise",
		"enterprise",
		enterprise0,
	))
	battery{
		{"err", err, nil},
		{"Ok()", post.Ok(), true},
		{"Status()", post.Status(), http.StatusCreated},
		{"enterprise.name", result.Name, name0},
	}.Run("post", t)
	href := post.Location()

	put, err := core.Rest(result, core.Put(
		href,
		"enterprise",
		"enterprise",
		enterprise1,
	))
	battery{
		{"err", err, nil},
		{"Ok()", put.Ok(), true},
		{"Status()", put.Status(), http.StatusOK},
		{"enterprise.Name", result.Name, name1},
	}.Run("put", t)

	get, err := core.Rest(result, core.Get(href, "enterprise"))
	battery{
		{"err", err, nil},
		{"Ok()", get.Ok(), true},
		{"Status()", get.Status(), http.StatusOK},
		{"enterprise.Name", result.Name, name1},
	}.Run("get", t)

	delete1, err := core.Rest(nil, core.Delete(href))
	battery{
		{"err", err, nil},
		{"Ok()", delete1.Ok(), true},
		{"Status()", delete1.Status(), http.StatusNoContent},
	}.Run("delete1", t)

	delete2, err := core.Rest(nil, core.Delete(href))
	battery{
		{"err", err == nil, false},
		{"Ok()", delete2.Ok(), false},
		{"Status()", delete2.Status(), http.StatusNotFound},
	}.Run("delete2", t)

	values := url.Values{"idDatacenter": {"1"}}
	query, err := core.Rest(nil, core.Get("admin/rules", "rules").Query(values))
	battery{
		{"err", err, nil},
		{"Ok()", query.Ok(), true},
		{"Status()", query.Status(), http.StatusOK},
		//		{"call.href", call.href, "https://testing:443/api/admin/rules?idDatacenter=1"},
	}.Run("query", t)
}

func TestDTO(t *testing.T) {
	dto := core.NewDTO(none)
	battery{
		{"Href()", dto.URL(), ""},
		{"Type()", dto.Media(), core.Media("")},
	}.Run("none", t)

	dto.Add(self)
	battery{
		{"Href()", dto.URL(), self.URL()},
		{"Type()", dto.Media(), self.Media()},
	}.Run("self", t)

	dto.Add(edit)
	battery{
		{"Href()", dto.URL(), edit.URL()},
		{"Type()", dto.Media(), edit.Media()},
	}.Run("edit", t)
}

func resolve(path string) string {
	return environment["ABQ_ENDPOINT"] + "/" + path
}

func TestResolve(t *testing.T) {
	null := core.Resolve("", nil)
	path := core.Resolve("admin/rules", nil)
	query := core.Resolve("admin/rules", url.Values{"idDatacenter": {"1"}})

	battery{
		{"null", null, resolve("")},
		{"path", path, resolve("admin/rules")},
		{"query", query, resolve("admin/rules?idDatacenter=1")},
	}.Run("resolve", t)
}

func ExampleMedia() {
	fmt.Println(core.Media(""))
	fmt.Println(core.Media("text/plain"))
	fmt.Println(core.Media("enterprise"))

	// Output:
	//
	// text/plain
	// application/vnd.abiquo.enterprise+json
}

func TestCollection(t *testing.T) {
	collection := enterprises.Collection(nil)
	count := 0
	for collection.Next() {
		collection.Item()
		count++
	}

	battery{
		{"Size()", count, collection.Size()},
	}.Run("collection", t)
}

func ExampleCollection_List() {
	for _, e := range enterprises.Collection(nil).List() {
		fmt.Println(e.(*dto).Name)
		break
	}

	// Output:
	// Abiquo
}

func ExampleCollection_Find() {
	for _, name := range []string{"Abiquo", "abq"} {
		finder := func(r core.Resource) bool {
			return r.(*dto).Name == name
		}
		result := enterprises.Collection(nil).Find(finder)
		fmt.Println(result == nil)
	}

	// Output:
	// false
	// true
}

func ExampleCollection_First() {
	e := enterprises.Collection(nil).First()
	fmt.Println(e.(*dto).Name)

	// Output:
	// Abiquo
}

func ExampleRemove() {
	var (
		create1 = core.Create(enterprise, enterprise0)
		create2 = core.Create(enterprise, enterprise0)
		read    = core.Read(enterprise0, enterprise0)
		update  = core.Update(enterprise0, enterprise0)
		remove1 = core.Remove(enterprise0)
		remove2 = core.Remove(enterprise0)
	)
	fmt.Println(create1)
	fmt.Println(create2)
	fmt.Println(read)
	fmt.Println(update)
	fmt.Println(remove1)
	fmt.Println(remove2)

	// Output:
	// <nil>
	// 409 Unexpected status code: ENTERPRISE-4 Duplicate name for an enterprise
	// <nil>
	// <nil>
	// <nil>
	// 404 Unexpected status code: EN-0 The requested enterprise does not exist
}

func TestWalk(t *testing.T) {
	endpoint := core.NewLinker("admin/enterprises/1/users/1", "user")
	username, err1 := endpoint.Walk()
	enterprise, err2 := username.Walk("enterprise")

	battery{
		{"err1", err1, nil},
		{"err2", err2, nil},
		{"nil", enterprise == nil, false},
		{"name", enterprise.(*dto).Name, "Abiquo"},
	}.Run("walk", t)
}

func ExampleResources_Map() {
	var count int
	counter := func(r core.Resource) { count++ }
	collection := enterprises.Collection(nil)
	collection.List().Map(counter)
	fmt.Println(collection.Size() == count)

	// Output:
	// true
}

func ExampleResources_Filter() {
	filter := func(r core.Resource) bool { return r.(*dto).Name == "Abiquo" }
	collection := enterprises.Collection(nil).List().Filter(filter)
	fmt.Println(collection[0].(*dto).Name)
	fmt.Println(len(collection))

	// Output:
	// Abiquo
	// 1
}

func TestUpload(t *testing.T) {
	reply, err := core.Upload(environment["ABQ_TEMPLATES"], environment["ABQ_OVA"], "")

	battery{
		{"err", err, nil},
		{"Ok()", reply.Ok(), true},
		{"Location()", reply.Location() != "", true},
	}.Run("ova", t)

	reply, err = core.Upload(environment["ABQ_TEMPLATES"], environment["ABQ_DISK"], `{
		"categoryName" : "Others",
		"description"  : "ojal ExampleUploadTemplate",
		"disks" : [{
			"bootable"       : true,
			"diskController" : "IDE",
			"diskFileFormat" : "UNKNOWN",
			"diskFilePath"   : "disk1",
			"diskFileSize"   : 1024000,
			"requiredHDInMB" : 1024000,
			"sequence"       : 0
		}],
		"name"             : "uploadTemplate",
		"requiredCpu"      : 1,
		"requiredRamInMB"  : 1024
	}`)

	battery{
		{"err", err, nil},
		{"Ok()", reply.Ok(), true},
		{"Location()", reply.Location() != "", true},
	}.Run("disk", t)
}
