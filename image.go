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
	WebpImageFormat   ImageFormat = "webp"
	GifImageFormat    ImageFormat = "gif"
	LottieImageFormat ImageFormat = "json"
)