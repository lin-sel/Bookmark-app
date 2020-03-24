import { Component, OnInit } from '@angular/core';
import { Form, FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/service/login/login.service';
import { LoggerService } from 'src/app/service/utils/logger.service';

@Component({
      selector: 'app-login',
      templateUrl: './login.component.html',
      styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

      private login: FormGroup
      constructor(
            private formbuilder: FormBuilder,
            private route: Router,
            private loginser: LoginService,
            private logger: LoggerService
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

      get f() {
            return this.login.controls;
      }

      appLogin() {
            console.log(this.login.value);
            this.loginser.login(this.login.value).then(() => {
                  this.logger.log("Login done")
                  alert("Login Done")
                  this.navigate("bookmark")
            }).catch(err => {
                  this.logger.error(err)
                  alert(err)
            });
      }

      navigate(path: string) {
            this.route.navigate([path]);
      }

}
