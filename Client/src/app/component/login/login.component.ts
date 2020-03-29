import { Component, OnInit } from '@angular/core';
import { Form, FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/service/login/login.service';
import { LoggerService } from 'src/app/service/utils/logger.service';
import { JsonService } from 'src/app/service/utils/json.service';
import { MainService } from 'src/app/service/main.service';
import { UtilService } from 'src/app/service/utils/util.service';

@Component({
      selector: 'app-login',
      templateUrl: './login.component.html',
      styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

      public login: FormGroup
      public loader: string = 'loader'
      constructor(
            private formbuilder: FormBuilder,
            private util: UtilService,
            private mainservice: MainService,
            private logger: LoggerService,
      ) { }

      ngOnInit() {
            this.initForm()
            this.configLoader()
      }

      initForm() {
            this.login = this.formbuilder.group({
                  username: ['', Validators.required],
                  password: ['', Validators.required]
            });
      }


      // Return Form Controls.
      get f() {
            return this.login.controls;
      }


      // Login to App.
      appLogin() {
            this.configLoader()
            console.log(this.login.value);
            this.mainservice.appLogin(this.login.value).then(() => {
                  this.logger.log("Login done")
                  alert("Login Done")
                  this.navigate("bookmark")
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
            }).finally(() => {
                  this.configLoader();
            });
      }

      // Navigate to Another URL.
      navigate(path: string) {
            this.util.navigate(path);
      }

      // Error Parser.
      errorParser(err) {
            return this.util.errorParser(err);
      }

      configLoader() {
            let obj = {
                  loader: this.loader
            }
            this.util.configLoader(obj)
            this.loader = obj.loader
      }
}
