package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//--------------------------------------------------------------------------------

func main() {

	// Does some weird Experiments with Images :D

	var dir_src, dir_dst string

	// Run Settings
	dir_src = "/tmp/a"
	dir_dst = "/tmp/b"

	// Starting...
	rand.Seed(time.Now().UnixNano())

	// Process Directory
	processDirectory(dir_src, dir_dst)
}

//--------------------------------------------------------------------------------

func processDirectory(dir_src, dir_dst string) {

	// Reads Directory and processes Files.

	var allowedExtensions []string
	var files []os.FileInfo
	var fi os.FileInfo
	var err error
	var allowed bool
	var fn, fe, fb, e, pathA, pathB, fileNewSuffix string
	var li int

	// Internal Settings
	fileNewSuffix = "" // added to the end of name, before extension
	allowedExtensions = []string{"jpg", "jpeg", "png"}

	// Get List of Files & Directories inside specified Directory
	files, err = ioutil.ReadDir(dir_src)
	if err != nil {
		fmt.Println("Error reading directory:", dir_src) //
	}

	for _, fi = range files {

		// Dir?
		if fi.IsDir() {
			continue
		}

		fn = filepath.Base(fi.Name())
		li = strings.LastIndex(fn, ".")

		// No extension?
		if li == -1 {
			continue
		}

		// File with Extension, split Base and Extension
		fe = fn[li+1:]
		fb = fn[:li]

		// Allowed Extension?
		allowed = false
		for _, e = range allowedExtensions {
			if e == fe {
				allowed = true
				break
			}
		}
		if !allowed {
			continue
		}

		// Allowed File, do the Job
		pathA = dir_src + "/" + fb + "." + fe
		pathB = dir_dst + "/" + fb + fileNewSuffix + "." + fe
		processFile(pathA, pathB, strings.ToLower(fe))
	}

}

//--------------------------------------------------------------------------------

func processFile(path_src, path_dst, ext string) {

	// Processes the File.
	//
	// ext must be lowercased

	var file, file2 *os.File
	var err error
	var img, img2 *image.Image
	var opts_jpg *jpeg.Options
	var coder_png *png.Encoder

	// Open File
	file, err = os.Open(path_src)
	if err != nil {
		fmt.Println("Error opening image:", path_src) //
	}

	// Decode File
	img = new(image.Image)
	if (ext == "jpg") || (ext == "jpeg") {

		*img, err = jpeg.Decode(file)

	} else if ext == "png" {

		*img, err = png.Decode(file)
	}
	if err != nil {
		fmt.Println("Error decoding image:", path_src) //
	}

	// Close opened File
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing file:", path_src) //
	}

	// Process Image
	img2 = new(image.Image)
	fmt.Print(path_src, "\n") // FilePath
	processImage(img, img2)

	// Create output File
	file2, err = os.Create(path_dst)
	if err != nil {
		fmt.Println("Error creating file:", path_dst) //
	}

	// Encode & put Image to File
	if (ext == "jpg") || (ext == "jpeg") {

		opts_jpg = new(jpeg.Options)
		opts_jpg.Quality = 100
		err = jpeg.Encode(file2, *img2, opts_jpg)

	} else if ext == "png" {

		coder_png = new(png.Encoder)
		coder_png.CompressionLevel = png.BestCompression
		err = coder_png.Encode(file2, *img2)

	}
	if err != nil {
		fmt.Println("Error encoding image:", path_dst) //
	}

	// Close opened File
	err = file2.Close()
	if err != nil {
		fmt.Println("Error closing file:", path_dst) //
	}
}

//--------------------------------------------------------------------------------

func processImage(imgA, imgB *image.Image) {

	// De-Colorizes the Image, increases Contrast, colorizes the monochrome Image.
	//
	// On Images which are heavily compressed using lossy Methods (such as JPEG),
	// this Procession may create visible Distortion.

	var rgba *image.RGBA
	var rect image.Rectangle
	var col, pixel_old color.Color
	var wA, hA, x, y, y_max, x_max int
	var r, g, b, a, ch_grey, ch_min, ch_max, ch_min_img, ch_max_img, lt_pix, pix_count uint32
	var nr, ng, nb, na uint8
	var st, k, lightness, k_dark, contrast_max_k, kr, kg, kb, kmin, krnd, kextra float64

	// New Image
	rect = (*imgA).Bounds()
	rgba = image.NewRGBA(rect)

	// Sizes
	wA = (*imgA).Bounds().Dx()
	hA = (*imgA).Bounds().Dy()
	pix_count = uint32(hA * wA)
	y_max = hA - 1
	x_max = wA - 1

	// Initialize internal Parameters
	contrast_max_k = 5
	lightness = 0
	ch_min_img = 65535
	st = 1
	ch_max_img = 0
	kmin = 0.3      // Colorization: minimum random k
	krnd = 1 - kmin // Colorization: base random k
	kextra = 0      // Colorization: extra random k

	// Get Image Statistics
	for y = 0; y < y_max; y++ {
		for x = 0; x < x_max; x++ {
			pixel_old = (*imgA).At(x, y)
			r, g, b, a = pixel_old.RGBA()

			// image's min channel
			if r < ch_min_img {
				ch_min_img = r
			}
			if g < ch_min_img {
				ch_min_img = g
			}
			if b < ch_min_img {
				ch_min_img = b
			}

			// image's max channel
			if r > ch_max_img {
				ch_max_img = r
			}
			if g > ch_max_img {
				ch_max_img = g
			}
			if b > ch_max_img {
				ch_max_img = b
			}

			// lightness
			lt_pix = (r + g + b) / 3
			lightness = lightness + float64(lt_pix)/float64(pix_count)
		}
	}

	// Lightness is in [0; 1]
	lightness = lightness / 65535

	// Experimental Things to make non-linear Contrast
	k_dark = math.Pow(math.Sin(lightness*math.Pi/2), 2) * contrast_max_k

	st = k_dark                                 // Contrast Strength
	k = math.Pow(float64(65535), float64(st-1)) // Contrast Base

	fmt.Printf("LT=%f KD=%f K=%e\n", lightness, k_dark, k) // Show Parameters

	// Random Colorization Koefficient with Limits
	kr = kmin + krnd*rand.Float64() + kextra
	kg = kmin + krnd*rand.Float64() + kextra
	kb = kmin + krnd*rand.Float64() + kextra

	// Modify Image
	for y = 0; y < y_max; y++ {
		for x = 0; x < x_max; x++ {

			// Read Channels of a Pixel
			pixel_old = (*imgA).At(x, y)
			r, g, b, a = pixel_old.RGBA()

			// Get maximum Channel
			ch_max = r
			if g > ch_max {
				ch_max = g
			}
			if b > ch_max {
				ch_max = b
			}

			// Get minimum Channel
			ch_min = r
			if g < ch_min {
				ch_min = g
			}
			if b < ch_min {
				ch_min = b
			}

			// Select control Channel
			//ch_grey = (ch_max + ch_min) / 2  // less contrast, but more smooth
			ch_grey = ch_max

			// Increase Contrast
			ch_max = uint32(math.Pow(float64(ch_grey), float64(st)) / k)

			// Colorize
			nr = uint8(kr * float64(ch_grey) / 256)
			ng = uint8(kg * float64(ch_grey) / 256)
			nb = uint8(kb * float64(ch_grey) / 256)

			// Alpha Channel
			na = uint8(a / 256) // is not Changed

			// Apply changes to Image
			col = color.RGBA{nr, ng, nb, na}
			rgba.Set(x, y, col)
		}
	}

	// RGBA -> Image
	(*imgB) = rgba.SubImage(rect)

}

//--------------------------------------------------------------------------------
