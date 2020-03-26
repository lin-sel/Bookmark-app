import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { RegisterService } from 'src/app/service/register/register.service';
import { LoggerService } from 'src/app/service/utils/logger.service';
import { JsonService } from 'src/app/service/utils/json.service';
import { MainService } from 'src/app/service/main.service';

@Component({
      selector: 'app-register',
      templateUrl: './register.component.html',
      styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

      public register: FormGroup
      constructor(
            private formbuilder: FormBuilder,
            private route: Router,
            private mainservice: MainService,
            private logger: LoggerService,
            private json: JsonService
      ) { }

      ngOnInit() {
            this.initForm()
      }


      // Create Form Object.
      initForm() {
            this.register = this.formbuilder.group({
                  name: ['', Validators.required],
                  username: ['', Validators.required],
                  password: ['', Validators.required]
            });
      }

      // Return Form Controls.
      get f() {
            return this.register.controls;
      }

      // Register User.
      userRegister() {
            this.mainservice.userRegister(this.register.value).then(() => {
                  this.logger.info("Register done");
                  alert("You have Register Successfully Now Login with your username and password.");
                  this.navigate("login");
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
            });
      }

      // Navigate to Another URL.
      navigate(path: string) {
            this.route.navigate([path]);
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
