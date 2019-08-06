1:
	go run main1.go common.go <input

2:
	go run main2.go common.go <input | \
	convert -size 370x100 xc:white -font "FreeMono" -pointsize 12 \
	  -fill black -annotate +10+20  "@-" -blur 0x3 -sharpen 0x1.0 \
	  -units PixelsPerInch -density 72 png:- | \
	tesseract stdin stdout | \
	grep .
