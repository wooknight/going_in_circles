package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

//MyFirstAPIKey
//6b2b13b6d0c1ef30a19e72f46fef4dd58167094ae535343d3f10264712325dfc

func init() {
	// To get your free API key for metered license, sign up on: https://cloud.unidoc.io
	// Make sure to be using UniPDF v3.19.1 or newer for Metered API key support.
	err := license.SetMeteredKey(`6b2b13b6d0c1ef30a19e72f46fef4dd58167094ae535343d3f10264712325dfc`)
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		panic(err)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: go run pdf_highlight_text.go <input.pdf> <output.pdf> <term> \n")
		os.Exit(0)
	}
	inputPath := os.Args[1]
	outputPath := os.Args[2]
	term := os.Args[3]

	err := highlightWords(inputPath, outputPath, term)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully highlighted the word %s and created %s\n", term, outputPath)
}

// getBoundingBoxes returns the bounding boxes of all instances of search`term` in the Pdf file.
func getBoundingBoxes(page *model.PdfPage, term string) ([]*model.PdfRectangle, error) {
	boundingBoxes := []*model.PdfRectangle{}
	ex, _ := extractor.New(page)
	pageText, _, _, err := ex.ExtractPageText()
	if err != nil {
		return nil, err
	}
	textMarks := pageText.Marks()
	text := pageText.Text()
	indexes := indexAllSubstrings(text, term)
	if len(indexes) == 0 {
		return nil, nil
	}
	for _, start := range indexes {
		end := start + len(term)
		spanMarks, err := textMarks.RangeOffset(start, end)
		if err != nil {
			return nil, err
		}
		bbox, ok := spanMarks.BBox()
		if !ok {
			return nil, fmt.Errorf("spanMarks.BBox has no bounding box. spanMarks=%s", spanMarks)
		}
		boundingBoxes = append(boundingBoxes, &bbox)
	}
	return boundingBoxes, nil

}

// indexAllSubstrings gets the begning indexes where `term` occures inside `text`.
func indexAllSubstrings(text, term string) []int {
	if len(term) == 0 {
		return nil
	}
	var indexes []int
	for start := 0; start < len(text); {
		i := strings.Index(text[start:], term)
		if i < 0 {
			return indexes
		}
		indexes = append(indexes, start+i)
		start += i + len(term)
	}
	return indexes
}

// highlightWords highlights all occurrences of the search term.
func highlightWords(inputPath, outputPath, term string) error {
	reader, f, err := model.NewPdfReaderFromFile(inputPath, nil)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	numPages, err := reader.GetNumPages()
	if err != nil {
		return err
	}
	cr := creator.New()
	for n := 1; n <= numPages; n++ {
		page, err := reader.GetPage(n)
		if err != nil {
			return err
		}
		bBoxes, err := getBoundingBoxes(page, term)
		if err != nil {
			return err
		}

		mediaBox, err := page.GetMediaBox()
		if err != nil {
			return err
		}
		if page.MediaBox == nil {
			page.MediaBox = mediaBox
		}

		if err := cr.AddPage(page); err != nil {
			return err
		}
		h := mediaBox.Ury
		for _, bbox := range bBoxes {
			rect := cr.NewRectangle(bbox.Llx, h-bbox.Lly, bbox.Urx-bbox.Llx, -(bbox.Ury - bbox.Lly))
			rect.SetFillColor(creator.ColorRGBFromHex("#FFFF00"))
			rect.SetBorderWidth(0)
			rect.SetFillOpacity(0.5)
			if err := cr.Draw(rect); err != nil {
				return nil
			}
		}

	}
	cr.SetOutlineTree(reader.GetOutlineTree())

	if err := cr.WriteToFile(outputPath); err != nil {
		return fmt.Errorf("failed to write the output file %s", outputPath)
	}
	return nil
}
