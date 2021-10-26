package response


// File用于文件的下载 or 传输
func File(fileName string, fileData []byte) Response  {
	return Response{
		Type: TypeFile,
		File: fileData,
		FileName: fileName,
	}
}