package main

var testCase1 map[string]string = map[string]string{
	"/a/b/c;/a/b/c;/a/b/c/": "/a/b/c",
	"../a/b;/a/b;../a/b":    "/a/b",
	"/a;/b;/c;/b;/a/;a":     "/a;/b;/c",
}
