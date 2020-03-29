import { Component, OnInit, Sanitizer } from '@angular/core';
import { MainService } from 'src/app/service/main.service';
import { JsonService } from 'src/app/service/utils/json.service';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UtilService } from 'src/app/service/utils/util.service';
import { DomSanitizer } from '@angular/platform-browser';

@Component({
      selector: 'app-profile',
      templateUrl: './profile.component.html',
      styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

      public userdata: any;
      public user: FormGroup
      public edit: boolean = false;
      public loader: any;
      constructor(
            private mainservice: MainService,
            private formbuilder: FormBuilder,
            private util: UtilService,
            private sanitize: DomSanitizer
      ) {
            this.loader = {
                  loader: "loader",
                  body: "hide"
            }
      }

      ngOnInit() {
            this.initForm();
            this.getUser();
      }

      initForm() {
            this.user = this.formbuilder.group({
                  name: ['', Validators.required],
                  profile: [''],
                  username: ['', Validators.required],
                  password: ['', Validators.required],
                  email: ["", Validators.required]
            });
      }

      patchValue(user: any) {
            this.user.patchValue({
                  name: user.name,
                  username: user.username,
                  password: user.password,
                  email: user.email
            });
      }

      // User form Control return
      get f() {
            return this.user.controls;
      }

      // User Profile data set.
      profileIMG() {
            if (this.userdata) {
                  if (!this.userdata.profile) {
                        console.log("Default")
                        return "/assets/images/default.png";
                  }
                  return this.sanitize.bypassSecurityTrustUrl((atob(this.userdata.profile)));
            }
      }

      // Update User Data.
      update() {
            // this.setImageProfile()
            this.configLoader();
            console.log(this.user.value);
            this.mainservice.updateUser(this.user.value).then((respond: any) => {
                  alert("Data Updated.");
                  document.location.reload();
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            }).finally(() => {
                  this.configLoader();
            });
      }

      onFileChange(event) {
            const reader = new FileReader();
            console.log(event);
            if (event.target.files && event.target.files.length) {
                  let file = event.target.files[0];
                  if (!file.type.includes("image")) {
                        alert("Only Image is Allowed");
                        return
                  }
                  reader.readAsDataURL(file);
                  reader.onload = () => {
                        this.user.patchValue({
                              profile: reader.result
                        });
                  }
                  this.user.controls["profile"].markAsDirty();
            }
      }

      // Update User Data.
      getUser() {
            this.mainservice.getUser().then((respond: any) => {
                  this.patchValue(respond)
                  this.userdata = respond
                  console.log(respond)
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            }).finally(() => {
                  this.configLoader();
            });
      }

      // Toggle Flag between true and false.
      editProfile() {
            if (this.edit) {
                  this.edit = false;
                  return;
            }
            this.edit = true;
      }

      // Error Parser.
      errorParser(err) {
            // let er = this.json.fromStringToJSON(err.error);
            // if (er != undefined) {
            //       return er.error;
            // }
            // return err.error;
            return this.util.errorParser(err);
      }

      // Check Session Expire and Perform Accordingly
      isSessionExpire(s: string) {
            // console.log(this.mainservice.isSessionExpire(s))
            // if (this.mainservice.isSessionExpire(s)) {
            //       this.router.navigate(["login"]);
            // }
            this.util.isSessionExpire(s);
      }

      configLoader() {
            this.util.configLoader(this.loader)
      }



}
