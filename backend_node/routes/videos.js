var express = require('express');
var router = express.Router();
var fs = require('fs')
var http = require('http')
var path = require('path')

/* GET users listing. */
router.get('/record', function (req, res, next) {
	http.get('http://210.212.194.12:8888/feed1.webm', (r) => {
		
		var file = fs.createWriteStream( path.join(__dirname, '../public/videos/') + Number(new Date()).toString() + '.mp4')

		var write = (x) => {
			file.write(x)
		}

		r.on('data', write)

		r.on('close', () => {
			r.removeListener('data', write)
			file.close()
			file.end()
		})

		setTimeout(() => {
			r.emit('close')
			console.log("Done")
			res.end()
		}, 30000)
	})
});

router.get('/:id',(req,res,next)=>{
    if(fs.existsSync(path.join(__dirname, '../public/videos',req.params.id))){
        res.sendFile(path.join(__dirname, '../public/videos',req.params.id));
    }
})

module.exports = router;
