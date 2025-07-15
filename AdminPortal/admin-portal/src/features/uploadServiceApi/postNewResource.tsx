import PostNewReosourceHeadersModel from "./requestModels/postNewResourceHeadersModel"

const PostNewResource = async (headers: PostNewReosourceHeadersModel) => {
	const url = "http://localhost:5050/videos"

	const requestHeaders: HeadersInit = new Headers()

	console.log(headers.ContentLength.toString())

	requestHeaders.set("Upload-Length", headers.UploadLength.toString())
	requestHeaders.set("Content-Length", headers.ContentLength.toString())
	requestHeaders.set("File-Name", headers.FileName)


	const resp = await fetch(url, {
		method: "POST",
		headers: requestHeaders
	})

}

export default PostNewResource

