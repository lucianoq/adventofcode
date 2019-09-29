1:
	go run main1.go <input | \
    convert -size 900x280 xc:white -font "FreeMono" -pointsize 24 \
      -fill black -annotate +20+30 "@-" -blur 0x6 -sharpen 0x1.0 \
      -units PixelsPerInch -density 72 png:- | \
    tesseract stdin stdout \
   	grep .

2:
	go run main2.go <input

