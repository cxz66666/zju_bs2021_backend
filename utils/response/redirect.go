package response


// Redirect 用于30x重定向
func Redirect(code int, url string) Response {
	return Response{
		Type: TypeRedirect,
		RedirectCode: code,
		RedirectURL: url,
	}
}


