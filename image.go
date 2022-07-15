package httpcord

type (
	ImageFormat       string
	MfaLevel          int
	SystemChannelFlag int
)

func (i ImageFormat) String() string {
	return string(i)
}

const (
	JpegImageFormat   ImageFormat = "jpg"
	PngImageFormat    ImageFormat = "png"
	GifImageFormat    ImageFormat = "gif"
	WebpImageFormat   ImageFormat = "webp"
	LottieImageFormat ImageFormat = "json"
)
