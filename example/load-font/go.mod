module github.com/mzky/captcha/example/load-font

go 1.20

replace github.com/mzky/captcha => ../../

require (
	github.com/mzky/captcha v0.0.0-00010101000000-000000000000
	golang.org/x/image v0.13.0
)

require github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
