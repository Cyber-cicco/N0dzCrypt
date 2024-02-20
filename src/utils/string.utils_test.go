package utils

import (
	"strings"
	"testing"
)

func TestToPackageName(t *testing.T) {
    p1 := transformIntoPackageName("test")
    p2 := transformIntoPackageName("test-test")
    p3 := transformIntoPackageName("test-test-")
    p4 := transformIntoPackageName("_test-test-")
    p5 := transformIntoPackageName("_test--------test-")
    if p1 != "test" {
        t.Fatalf("erreur pour le test 1 : test")
    }
    if p2 != "testtest" {
        t.Fatalf("erreur pour le test 2 : test-test")
    }
    if p3 != "testtest" {
        t.Fatalf("erreur pour le test 3 : test-test-")
    }
    if p4 != "testtest" {
        t.Fatalf("erreur pour le test 4 : _test-test-")
    }
    if p5 != "testtest" {
        t.Fatalf("erreur pour le test 4 : _test--------test-. Wanted testtest, got %q", p5)
    }
}

func TestGetApplicationName(t *testing.T) {
    a1 := GetApplicationName("test")
    a2 := GetApplicationName("test-test")
    a3 := GetApplicationName("test-test-")
    a4 := GetApplicationName("_test-test-")
    a5 := GetApplicationName("_test--------test-")
    if a1 != "TestApplication" {
        t.Fatalf("erreur pour le test 1 : test. Wanted TestApplicaton, got %q", a1)
    }
    if a2 != "TestTestApplication" {
        t.Fatalf("erreur pour le test 2 : test-test. Wanted TestTestApplicaton, got %q", a2)
    }
    if a3 != "TestTestApplication" {
        t.Fatalf("erreur pour le test 3 : test-test-. Wanted TestTestApplicaton, got %q", a3)
    }
    if a4 != "TestTestApplication" {
        t.Fatalf("erreur pour le test 4 : _test-test-. Wanted TestTestApplicaton, got %q", a4)
    }
    if a5 != "TestTestApplication" {
        t.Fatalf("erreur pour le test 5 : _test--------test-. Wanted TestTestApplicaton, got %q", a5)
    }
}

