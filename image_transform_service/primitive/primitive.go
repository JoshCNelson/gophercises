package primitive

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Mode defines the shapes used when transforming images
type Mode int

// Modes supported by the primitive package
const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedrect
	ModeBeziers
	ModeRotatedellipse
	ModePolygon
)

// WithMode is an option for the Transform function that will define
// the mode that you want to use. By default, ModeTriangle will be used
func WithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

// func NumShapes(n int) func() []string {
// 	return func() []string {
// 		return []string{"-n", fmt.Sprintf("%d", n)}
// 	}
// }

// Transform will take the provided image and apply a primitive
// transformation to it, then return a reader to the resulting image
func Transform(image io.Reader, ext string, numShapes int, opts ...func() []string) (io.Reader, error) {
	var args []string
	for _, opt := range opts {
		args = append(args, opt()...)
	}
	in, err := tempfile("in_", ext)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary input file")
	}
	defer os.Remove(in.Name())

	out, err := tempfile("in_", ext)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary output file")
	}
	defer os.Remove(out.Name())

	_, err = io.Copy(in, image)
	if err != nil {
		return nil, errors.New("primitive: failed to copy image inout temp input file")
	}

	stdCombo, err := primitive(in.Name(), out.Name(), numShapes, args...)
	if err != nil {
		return nil, fmt.Errorf("primitive: failed to run primitive command. stdcombo=%s", stdCombo)
	}

	// if strings.TrimSpace(stdCombo) == "" {
	// panic(stdCombo)
	// }

	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, out)
	if err != nil {
		return nil, errors.New("primitive: failed to copy output file into byte buffer")
	}

	return b, nil
}

func primitive(inputFile string, outputFile string, numShapes int, args ...string) (string, error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d", inputFile, outputFile, numShapes)
	args = append(strings.Fields(argStr), args...)
	cmd := exec.Command("primitive", args...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}

func tempfile(prefix string, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", prefix)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
