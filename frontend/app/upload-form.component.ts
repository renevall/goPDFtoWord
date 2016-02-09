import {Http, Headers} from 'angular2/http';
import {Component} from 'angular2/core';
import {NgForm}    from 'angular2/common';
import {UPLOAD_DIRECTIVES} from 'ng2-uploader';


@Component({
	selector: "upload-form",
	templateUrl: "app/upload-form.component.html",
	directives: [UPLOAD_DIRECTIVES]
})

export class UploadFormComponent {
	submitted = false;
	uploadFile:any;
	options: Object = {
    url: 'http://localhost:8080/upload'
  };

	constructor(){}

	onSubmit(){
		this.submitted = true;
		console.log("Submitted");
	}

	handleUpload(data){
		console.log("upload triggered");
		if (data && data.response) {
			console.log(data.response)
		//  data = JSON.parse(data.response);
		 this.uploadFile = data;
	 	}
	}
}
