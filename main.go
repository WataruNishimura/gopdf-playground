package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/signintech/gopdf"
)

const PORT = "50000"

func HandlePDF(ctx *gin.Context) {
	pdf := gopdf.GoPdf{}
	const cm1topx = 37.8 // 1cm to px
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.SetMargins(cm1topx, cm1topx, cm1topx, cm1topx)
	defer pdf.Close()
	pdf.AddPage()

	err := pdf.AddTTFFont("ipa", "ipag.ttf")

	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("ipa", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetLineWidth(gopdf.PageSizeA4.W - cm1topx*2)
	pdf.Br(20)

	strs, err := pdf.SplitText("ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ", gopdf.PageSizeA4.W-cm1topx*2)

	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetY(cm1topx)

	if len(strs) != 0 {
		for _, line := range strs {
			pdf.Text(line)
			pdf.Br(20)
		}
	}

	ctx.Writer.Header().Add("Content-Type", "application/pdf")
	ctx.Writer.Write(pdf.GetBytesPdf())
}

func main() {
	r := gin.Default()
	r.GET("/", HandlePDF)
	r.Run(":" + PORT)
}
