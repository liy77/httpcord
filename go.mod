module httpcord

go 1.18

require (
	github.com/valyala/fasthttp v1.38.0
	golang.org/x/exp v0.0.0-20220713135740-79cabaa25d75
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/klauspost/compress v1.15.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
)

retract (
	[v0.1.0, v0.1.4]
	v0.1.0
)
