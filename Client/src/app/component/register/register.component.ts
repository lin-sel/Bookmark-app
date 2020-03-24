import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { RegisterService } from 'src/app/service/register/register.service';
import { LoggerService } from 'src/app/service/utils/logger.service';

@Component({
      selector: 'app-register',
      templateUrl: './register.component.html',
      styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

      private register: FormGroup
      constructor(
            private formbuilder: FormBuilder,
            private route: Router,
            private registerser: RegisterService,
            private logger: LoggerService
      ) { }

      ngOnInit() {
            this.initForm()
      }

      initForm() {
            this.register = this.formbuilder.group({
                  name: ['', Validators.required],
                  username: ['', Validators.required],
                  password: ['', Validators.required]
            });
      }

      get f() {
            return this.register.controls;
      }

      userRegister() {
            this.registerser.register(this.register.value).then(() => {
                  this.logger.info("Register done");
                  alert("You have Register Successfully Now Login with your username and password.");
                  this.navigate("login");
            }).catch(err => {
                  this.logger.error(err);
                  alert(err)
            });
      }

      navigate(path: string) {
            this.route.navigate([path]);
      }

}
