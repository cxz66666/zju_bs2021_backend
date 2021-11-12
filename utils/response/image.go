package response

// Image 用于图片展示
func Image(fileName string) Response  {
	return Response{
		Type: TypeImage,
		FileName: fileName,
	}
}