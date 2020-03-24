import { Component, OnInit } from '@angular/core';
import { Form, FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/service/login/login.service';
import { LoggerService } from 'src/app/service/utils/logger.service';
import { JsonService } from 'src/app/service/utils/json.service';
import { MainService } from 'src/app/service/main.service';

@Component({
      selector: 'app-login',
      templateUrl: './login.component.html',
      styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

      private login: FormGroup
      constructor(
            private formbuilder: FormBuilder,
            private router: Router,
            private mainservice: MainService,
            private logger: LoggerService,
            private json: JsonService
      ) { }

      ngOnInit() {
            this.initForm()
      }

      initForm() {
            this.login = this.formbuilder.group({
                  username: ['nil', Validators.required],
                  password: ['nil', Validators.required]
            });
      }


      // Return Form Controls.
      get f() {
            return this.login.controls;
      }


      // Login to App.
      appLogin() {
            console.log(this.login.value);
            this.mainservice.appLogin(this.login.value).then(() => {
                  this.logger.log("Login done")
                  alert("Login Done")
                  this.navigate("bookmark")
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
            });
      }

      // Navigate to Another URL.
      navigate(path: string) {
            this.router.navigate([path]);
      }

      // Error Parser.
      errorParser(err) {
            let er = this.json.fromStringToJSON(err.error);
            if (er != undefined) {
                  return er.error;
            }
            return err.error;
      }
}
