import {Component} from 'angular2/core';
import {UploadFormComponent} from './upload-form.component'

@Component({
    selector: 'my-app',
    template:`
    <upload-form></upload-form>
    `,
    directives: [UploadFormComponent]
})
export class AppComponent { }
