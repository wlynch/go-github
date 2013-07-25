// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRepositoriesService_ListHooks(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	hooks, err := client.Repositories.ListHooks("o", "r")
	if err != nil {
		t.Errorf("Repositories.ListHooks returned error: %v", err)
	}

	want := []Hook{Hook{ID:1}, Hook{ID:2}}
	if !reflect.DeepEqual(hooks, want) {
		t.Errorf("Repositories.ListHooks returned %+v, want %+v", hooks, want)
	}
}

func TestRepositoriesService_CreateHook(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"id":1,"name":"web","events":["push"],"active":true,"config":{"url":"http://foo.com","content_type":"json"}}`)
	})

	configHook := NewWebHook([]string{"push"}, true,
		HookConfig{"url":"http://foo.com", "content_type":"json"})
	hook, err := client.Repositories.CreateHook("o", "r", configHook)
	if err != nil {
		t.Errorf("Repositories.CreateHook returned error: %+v", err)
		return
	}

	want := Hook{ "", nil, nil, 1, HookOptions{"web", []string{"push"}, true,
		HookConfig{"url":"http://foo.com", "content_type":"json",},},}
		if !reflect.DeepEqual(*hook, want) {
		t.Errorf("Repositories.CreateHook returned %+v, want %+v", *hook, want)
	}
}

func TestRepositoriesService_EditHook(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		fmt.Fprint(w, `{"id":1,"name":"web","events":["push"],"active":true,"config":{"url":"http://foo.com","content_type":"json"}}`)
	})

	configHook := NewWebHook([]string{"push"}, true,
		HookConfig{"url":"http://foo.com",})
	configHook.ID = 1
	hook, err := client.Repositories.EditHook("o", "r", configHook)
	if err != nil {
		t.Errorf("Repositories.EditHook returned error: %+v", err)
		return
	}

	want := Hook{ "", nil, nil, 1, HookOptions{"web", []string{"push"}, true,
		HookConfig{"url":"http://foo.com", "content_type":"json"},},}
		if !reflect.DeepEqual(*hook, want) {
		t.Errorf("Repositories.EditHook returned %+v, want %+v", *hook, want)
	}
}

func TestRepositoriesService_TestPushHook(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/hooks/1/tests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(201)
	})

	err := client.Repositories.TestPushHook("o", "r", 1)
	if err != nil {
		t.Errorf("Repositories.TestPushHook returned error: %+v", err)
		return
	}
}

func TestRepositoriesService_DeleteHook(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(204)
	})

	err := client.Repositories.DeleteHook("o", "r", 1)
	if err != nil {
		t.Errorf("Repositories.DeleteHook returned error: %+v", err)
		return
	}
}

