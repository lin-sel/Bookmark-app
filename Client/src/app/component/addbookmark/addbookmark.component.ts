import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { JsonService } from 'src/app/service/utils/json.service';
import { MainService } from 'src/app/service/main.service';
import { UtilService } from 'src/app/service/utils/util.service';
@Component({
      selector: 'app-addbookmark',
      templateUrl: './addbookmark.component.html',
      styleUrls: ['./addbookmark.component.css']
})
export class AddbookmarkComponent implements OnInit {

      public loader: string = "loader"
      public bookmark: FormGroup;
      public body: string = "hide";
      constructor(
            private activeroute: ActivatedRoute,
            private formbuilder: FormBuilder,
            private mainservice: MainService,
            private util: UtilService
      ) { }

      ngOnInit() {
            if (!this.mainservice.authUser()) {
                  alert("PLease Login First.");
                  this.util.navigate("login");
                  return;
            }
            this.getParam();
      }

      // Get ID from URL.
      getParam() {
            let id = this.activeroute.snapshot.paramMap.get('id');
            this.initForm(id);
      }

      // Form Object Created.
      initForm(categoryid) {
            this.bookmark = this.formbuilder.group({
                  url: ['', [Validators.required, Validators.pattern(/^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$/gm)]],
                  tag: ['', Validators.required],
                  label: ['', Validators.required],
                  categoryid: [categoryid, Validators.required]
            });
            this.configLoader();
      }


      // Add New Bookmark.
      addBookmark() {
            this.mainservice.addBookmark(this.bookmark.value).then(data => {
                  console.log("New Bookmark Added", data);
                  alert("New Bookmark Added");
                  this.navigate("bookmark");
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            }).finally(() => {
                  this.configLoader();
                  this.util.reload();
            })
      }

      // Return Form Controls.
      get f() {
            return this.bookmark.controls;
      }

      // Navigate to Another URL.
      navigate(path: string) {
            this.util.navigate(path)
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
            let obj = {
                  loader: this.loader,
                  body: this.body
            }
            this.util.configLoader(obj)
            this.loader = obj.loader
            this.body = obj.body
      }


}
