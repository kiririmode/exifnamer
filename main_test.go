package main

import "testing"

func Test_isJpeg(t *testing.T) {

	testCases := []struct {
		path     string
		expected bool
	}{
		{path: "a.jpeg", expected: true},
		{path: "a.JPEG", expected: true},
		{path: "a.JPG", expected: true},
		{path: "a.jpg", expected: true},
		{path: "a.ajpeg", expected: false},
		{path: "a.jpega", expected: false},
		{path: "ajpeg", expected: false}}

	for _, tc := range testCases {
		actual := isJpeg(tc.path)
		if actual != tc.expected {
			t.Errorf("'%s is JPEG' expected to be %t, but got %t",
				tc.path, tc.expected, actual)
		}
	}
}

func Test_isHeic(t *testing.T) {

	testCases := []struct {
		path     string
		expected bool
	}{
		{path: "a.heic", expected: true},
		{path: "a.HEIC", expected: true},
		{path: "a.aheic", expected: false},
		{path: "a.heica", expected: false},
		{path: "aheic", expected: false}}

	for _, tc := range testCases {
		actual := isHeic(tc.path)
		if actual != tc.expected {
			t.Errorf("'%s is HEIC' expected to be %t, but got %t",
				tc.path, tc.expected, actual)
		}
	}
}

func Test_isTargetImage(t *testing.T) {

	testCases := []struct {
		path     string
		expected bool
	}{
		{path: "a.jpeg", expected: true},
		{path: "a.heic", expected: true},
		{path: "a.gif", expected: false}}

	for _, tc := range testCases {
		actual := IsTargetImage(tc.path)
		if actual != tc.expected {
			t.Errorf("'%s is JPEG or HEIC' expected to be %t, but got %t",
				tc.path, tc.expected, actual)
		}
	}
}
