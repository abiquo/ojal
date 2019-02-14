package core_test

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/abiquo/ojal/abiquo"
	"github.com/abiquo/ojal/core"
)

func verify(t *testing.T, name string, expected, current interface{}) {
	if current != expected {
		t.Errorf("%v FAIL %v != %v", name, expected, current)
	}
}

var (
	ts          = time.Now().Unix()
	basic       core.Basic
	oauth       core.Oauth
	environment map[string]string
	enterprise0 = &abiquo.Enterprise{Name: fmt.Sprint("ztest A ", ts)}
	enterprise1 = &abiquo.Enterprise{Name: fmt.Sprint("ztest B ", ts)}
	result      = &abiquo.Enterprise{}
	none        *core.Link
	self        *core.Link
	edit        *core.Link
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

	core.RegisterMedia("enterprise", "enterprises", func() core.Resource { return new(abiquo.Enterprise) })
	core.RegisterMedia("user", "users", func() core.Resource { return new(abiquo.User) })
	core.Init(environment["ABQ_ENDPOINT"], basic)

	none = core.NewLink("none").SetType("none").SetRel("none").SetTitle("none")
	self = core.NewLink("self").SetType("self").SetRel("self").SetTitle("self")
	edit = core.NewLink("edit").SetType("edit").SetRel("edit").SetTitle("edit")
}

func TestLink(t *testing.T) {
	verify(t, "link.URL()", none.URL(), core.Resolve("none", nil))
	verify(t, "link.Media()", none.Media(), core.Media("none").Media())
	verify(t, "link.Title", none.Title, "none")
	verify(t, "link.Rel", none.Rel, "none")
}

func TestObject(t *testing.T) {
	object := make(core.Object)
	verify(t, "zero.Href()", object.Href(), "")
	verify(t, "zero.Media()", object.Media(), core.Media("").Media())

	object.Add(none)
	verify(t, "none.Href()", object.Href(), "")
	verify(t, "none.Media()", object.Media(), core.Media("").Media())

	object.Add(self)
	verify(t, "self.Href()", object.Href(), self.Href)
	verify(t, "self.Media()", object.Media(), self.Media())

	object.Add(edit)
	verify(t, "edit.Href()", object.Href(), edit.Href)
	verify(t, "edit.Media()", object.Media(), edit.Media())
}

func TestCall(t *testing.T) {
	post, err := core.Rest(result, core.Post(
		core.NewLink("admin/enterprises"),
		core.Media("enterprise"),
		core.Media("enterprise"),
		enterprise0,
	))
	verify(t, "post.err", err, nil)
	verify(t, "post.Ok()", post.Ok(), true)
	verify(t, "post.Status", post.StatusCode, http.StatusCreated)
	verify(t, "enterprise.name", result.Name, enterprise0.Name)

	put, err := core.Rest(result, core.Put(post, post, post, enterprise1))
	verify(t, "put.err", err, nil)
	verify(t, "put.Ok()", put.Ok(), true)
	verify(t, "put.Status", put.StatusCode, http.StatusOK)
	verify(t, "result.Name", result.Name, enterprise1.Name)

	get, err := core.Rest(result, core.Get(post, post))
	verify(t, "get.err", err, nil)
	verify(t, "get.Ok()", get.Ok(), true)
	verify(t, "get.Status", get.StatusCode, http.StatusOK)
	verify(t, "get.enterprise.Name", result.Name, enterprise1.Name)

	delete1, err := core.Rest(nil, core.Delete(post))
	verify(t, "delete1.err", err, nil)
	verify(t, "delete1.Ok()", delete1.Ok(), true)
	verify(t, "delete1.Status", delete1.StatusCode, http.StatusNoContent)

	delete2, err := core.Rest(nil, core.Delete(post))
	verify(t, "delete2.err", err == nil, false)
	verify(t, "delete2.Ok()", delete2.Ok(), false)
	verify(t, "delete2.Status", delete2.StatusCode, http.StatusNotFound)

	q := url.Values{"idDatacenter": {"1"}}
	rules := core.NewLink("admin/rules").SetType("rules")
	query, err := core.Rest(nil, core.Get(rules, rules).Query(q))
	verify(t, "query.err", err, nil)
	verify(t, "query.Ok()", query.Ok(), true)
	verify(t, "query.Status", query.StatusCode, http.StatusOK)
}

func TestDTO(t *testing.T) {
	dto := core.NewDTO(none)
	verify(t, "none.Href()", dto.URL(), "")
	verify(t, "none.Type()", dto.Media(), "")

	dto.Add(self)
	verify(t, "self.Href()", dto.URL(), self.URL())
	verify(t, "self.Type()", dto.Media(), self.Media())

	dto.Add(edit)
	verify(t, "edit.Href()", dto.URL(), edit.URL())
	verify(t, "edit.Type()", dto.Media(), edit.Media())
}

func TestResolve(t *testing.T) {
	null := core.Resolve("", nil)
	path := core.Resolve("admin/rules", nil)
	query := core.Resolve("admin/rules", url.Values{"idDatacenter": {"1"}})
	resolve := func(path string) string {
		return environment["ABQ_ENDPOINT"] + "/" + path
	}

	verify(t, "resolve.null", null, resolve(""))
	verify(t, "resolve.path", path, resolve("admin/rules"))
	verify(t, "resolve.query", query, resolve("admin/rules?idDatacenter=1"))
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
	collection := abiquo.Enterprises().Collection(nil)
	count := 0
	for collection.Next() {
		collection.Item()
		count++
	}
	verify(t, "collection.Size()", count, collection.Size())
}

func ExampleCollection_List() {
	for _, e := range abiquo.Enterprises().Collection(nil).List() {
		fmt.Println(e.Link().Title)
		break
	}

	// Output:
	// Abiquo
}

func ExampleCollection_Find() {
	for _, name := range []string{"Abiquo", "abq"} {
		result := abiquo.Enterprises().Collection(nil).Find(func(r core.Resource) bool {
			return r.Link().Title == name
		})
		fmt.Println(result == nil)
	}

	// Output:
	// false
	// true
}

func ExampleCollection_First() {
	e := abiquo.Enterprises().Collection(nil).First()
	fmt.Println(e.(*abiquo.Enterprise).Name)

	// Output:
	// Abiquo
}

func ExampleDTO_Remove() {
	var (
		endpoint = abiquo.Enterprises().SetType("enterprise")
		create1  = endpoint.Create(enterprise0)
		create2  = endpoint.Create(enterprise0)
		read     = enterprise0.Read(enterprise0)
		update   = enterprise0.Update(enterprise0)
		remove1  = enterprise0.Remove()
		remove2  = enterprise0.Remove()
	)
	fmt.Println(create1)
	fmt.Println(create2)
	fmt.Println(read)
	fmt.Println(update)
	fmt.Println(remove1)
	fmt.Println(remove2)

	// Output:
	// <nil>
	// [{ENTERPRISE-4 Duplicate name for an enterprise}]
	// <nil>
	// <nil>
	// <nil>
	// [{EN-0 The requested enterprise does not exist}]
}

func TestWalk(t *testing.T) {
	endpoint := core.NewLink("admin/enterprises/1/users/1").SetType("user")
	username, err1 := endpoint.Walk()
	verify(t, "walk.err1", err1, nil)

	enterprise, err2 := username.Rel("enterprise").Walk()
	verify(t, "walk.err2", err2, nil)
	verify(t, "walk.nil", enterprise == nil, false)
	verify(t, "walk.Title", enterprise.Link().Title, "Abiquo")
}

func ExampleResources_Map() {
	var count int
	counter := func(r core.Resource) { count++ }
	collection := abiquo.Enterprises().Collection(nil)
	collection.List().Map(counter)
	fmt.Println(collection.Size() == count)

	// Output:
	// true
}

func ExampleResources_Filter() {
	collection := abiquo.Enterprises().Collection(nil)
	filtered := collection.List().Filter(func(r core.Resource) bool {
		return r.(*abiquo.Enterprise).Name == "Abiquo"
	})
	fmt.Println(filtered[0].(*abiquo.Enterprise).Name)
	fmt.Println(len(filtered))

	// Output:
	// Abiquo
	// 1
}

func TestUpload(t *testing.T) {
	reply, err := core.Upload(environment["ABQ_TEMPLATES"], environment["ABQ_OVA"], "")
	verify(t, "ova.err", err, nil)
	verify(t, "ova.Ok()", reply.Ok(), true)
	verify(t, "ova.URL()", reply.URL() != "", true)

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
	verify(t, "disk.err", err, nil)
	verify(t, "disk.Ok()", reply.Ok(), true)
	verify(t, "disk.URL()", reply.URL() != "", true)
}
