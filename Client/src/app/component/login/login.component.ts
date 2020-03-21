import { Component, OnInit } from '@angular/core';
import { Form, FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

@Component({
      selector: 'app-login',
      templateUrl: './login.component.html',
      styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

      private login: FormGroup
      constructor(
            private formbuilder: FormBuilder
      ) { }

      ngOnInit() {
      }

      initForm() {
            this.login = this.formbuilder.group({
                  username: ['', Validators.required],
                  password: ['', Validators.required]
            })
      }

      get formControl() {
            return this.login.controls
      }

}
