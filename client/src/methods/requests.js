import axios from "axios"
import qs from "querystring"

// Change here
var url = 'http://localhost:3000';

const httpGet = (path) => {
	return axios.get(url + path, {
		headers: {
			'Content-Type': 'application/x-www-form-urlencoded'
		}
	})
}

const httpDelete = (path) => {
	return axios.delete(url + path, {
		headers: {
			'Content-Type': 'application/x-www-form-urlencoded'
		}
	})
}

const httpPost = (path, data) => {
	return axios.post(url + path, qs.stringify(data), {
		headers: {
			'Content-Type': 'application/x-www-form-urlencoded'
		}
	})
}

const httpFile = (path, data, func) => {

	var formData = new FormData()
	formData.append("file", data)
	return axios.post(url + path, formData, {
		onUploadProgress: (e) => {
			func(e.loaded / e.total)
		},
		headers: {
			'Content-Type': 'multipart/form-data'
		}
	})
}

export {
	httpGet,
	httpPost,
	httpDelete,
	httpFile
}
