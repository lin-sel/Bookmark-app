import { Component, OnInit } from '@angular/core';
import { Form, FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { LoggerService } from 'src/app/service/utils/logger.service';
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
      public isadmin: boolean = false
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
            if (!this.isadmin) {
                  this.mainservice.userLogin(this.login.value).then(() => {
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
                  return
            }
            this.mainservice.adminLogin(this.login.value).then(() => {
                  this.logger.log("Login done")
                  alert("Login Done")
                  this.navigate("admindashboard")
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
