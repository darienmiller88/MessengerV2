package utilities

import (
	"fmt"
	"net/http"
	"strings"
)

//Function to check to see if a url is a valid image url, i.e a url that leads directly to an .png or .jpg
func CheckValidImageURL(url string) error {
	// Make a HEAD request to the URL to request header content.
	resp, err := http.Head(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Check if the response status code is within the success range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Check the Content-Type header for image formats
	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "image/png") || strings.HasPrefix(contentType, "image/jpeg") {
		return nil
	}

	return fmt.Errorf("%s is not a valid image url", url)
}