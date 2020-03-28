import { Component, OnInit } from '@angular/core';
import { MainService } from 'src/app/service/main.service';
import { JsonService } from 'src/app/service/utils/json.service';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
      selector: 'app-profile',
      templateUrl: './profile.component.html',
      styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

      public userdata: any;
      public user: FormGroup
      public edit: boolean = false;
      constructor(
            private mainservice: MainService,
            private router: Router,
            private json: JsonService,
            private formbuilder: FormBuilder
      ) { }

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
                  email: ['default@gamil.com', Validators.required]
            });
      }

      patchValue(user: any) {
            this.user.patchValue({
                  name: user.name,
                  profile: user.profile,
                  username: user.username,
                  password: user.password,
                  email: 'default@gamil.com'
            });
      }

      // User form Control return
      get f() {
            return this.user.controls;
      }

      // User Profile data set.
      profileIMG() {
            if (this.userdata) {
                  if (this.userdata.profile == null) {
                        return "/assets/images/default.png";
                  }
                  return this.userdata.profile;
            }
      }

      // Update User Data.
      update() {
            // this.setImageProfile()
            this.mainservice.updateUser(this.user.value).then((respond: any) => {
                  alert("Data Updated.");
                  document.location.reload();
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            });
      }

      // setImageProfile() {
      //       let data = this.user.get("profile").value;
      //       let arr = data.split("\\");
      //       console.log(arr)
      //       console.log(arr[arr.length - 1])
      //       this.user.patchValue({
      //             profile: arr[arr.length - 1]
      //       })
      // }

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
            let er = this.json.fromStringToJSON(err.error);
            if (er != undefined) {
                  return er.error;
            }
            return err.error;
      }

      // Check Session Expire and Perform Accordingly
      isSessionExpire(s: string) {
            console.log(this.mainservice.isSessionExpire(s))
            if (this.mainservice.isSessionExpire(s)) {
                  this.router.navigate(["login"]);
            }
      }



}
