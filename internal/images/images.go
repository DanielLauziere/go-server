package images

import (
	// "bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "github.com/chai2010/webp"
	// "github.com/disintegration/imaging"
)

// Error - Error message and debuging object
type Error struct {
	ErrorMessage string `json:"errorMessage"`
}

// ImageAsset - image asset based on database construct
type ImageAsset struct {
	ImageURL string `json:"imageUrl"`
}

// SaveToWebp -
// func SaveToWebp(imageFileString string, webPString string) {

// 	var buf bytes.Buffer
// 	var err error

// 	m, err := imaging.Open(imageFileString)

// 	// Encode lossless webp
// 	if err = webp.Encode(&buf, m, &webp.Options{Lossless: true}); err != nil {
// 		log.Println(err)
// 	}
// 	if err = ioutil.WriteFile(webPString, buf.Bytes(), 0666); err != nil {
// 		log.Println(err)
// 	}

// }

// UploadFile for uploading images
func UploadFile(w http.ResponseWriter, r *http.Request) string {

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return ""
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "temp-*.jpg")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	tempFile.Write(fileBytes)

	// src, err := imaging.Open(tempFile.Name())

	// Resize srcImage to size = 128x128px using the Lanczos filter.
	// resized1200 := imaging.Resize(src, 1200, 0, imaging.Lanczos)
	// resized800 := imaging.Resize(src, 800, 0, imaging.Lanczos)
	// resized500 := imaging.Resize(src, 500, 0, imaging.Lanczos)

	r.ParseForm()
	// x := r.Form.Get("blogName")

	// unusedName := GetUnusedImageName(x)
	unusedName := handler.Filename

	// Save the resulting image as JPEG.
	// err = imaging.Save(resized1200, "static-files/fallbackImages/lg/"+unusedName+".jpg")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("failed to save image: %v", err)
	}

	// SaveToWebp("static-files/fallbackImages/lg/"+unusedName+".jpg", "static-files/images/lg/"+unusedName+".webp")

	// Save the resulting image as JPEG.
	// err = imaging.Save(resized800, "static-files/fallbackImages/md/"+unusedName+".jpg")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("failed to save image: %v", err)
	}

	// SaveToWebp("static-files/fallbackImages/md/"+unusedName+".jpg", "static-files/images/md/"+unusedName+".webp")

	// Save the resulting image as JPEG.
	// err = imaging.Save(resized500, "static-files/fallbackImages/sm/"+unusedName+".jpg")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("failed to save image: %v", err)
	}

	// SaveToWebp("static-files/fallbackImages/sm/"+unusedName+".jpg", "static-files/images/sm/"+unusedName+".webp")

	w.WriteHeader(http.StatusOK)
	return unusedName

}
