import axios from "axios"

// Change here
// var url = 'http://localhost:3000';
// var url = "http://10.53.50.53:55555"
var url = "http://10.100.81.223:55555"

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
	console.log(data)
	return axios.post(url + path, data, {
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
