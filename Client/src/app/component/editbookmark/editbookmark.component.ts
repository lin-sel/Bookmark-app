import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { JsonService } from 'src/app/service/utils/json.service';
import { BookmarkService } from 'src/app/service/bookmark/bookmark.service';
import { MainService } from 'src/app/service/main.service';
import { UtilService } from 'src/app/service/utils/util.service';

@Component({
      selector: 'app-editbookmark',
      templateUrl: './editbookmark.component.html',
      styleUrls: ['./editbookmark.component.css']
})
export class EditbookmarkComponent implements OnInit {

      private url: string = "^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$/gm"
      public bookmark: FormGroup;
      public loader: string = "loader"
      public body: string = "hide";
      constructor(
            private activeroute: ActivatedRoute,
            private formbuilder: FormBuilder,
            private json: JsonService,
            private mainservice: MainService,
            private router: Router,
            private util: UtilService
      ) { }

      ngOnInit() {
            this.initForm();
            if (!this.mainservice.authUser()) {
                  alert("PLease Login First.");
                  this.router.navigate(["login"]);
                  return;
            }
            this.getParam();
      }


      // Get Parameter From URL and Get Bookmark By ID.
      getParam() {
            let id = this.activeroute.snapshot.paramMap.get('id');
            let bookmk = this.mainservice.getBookmarkByID(id);
            if (bookmk == undefined) {
                  alert("Not Found");
                  this.navigate("bookmark");
                  return
            }
            this.patchValue(bookmk)
      }


      // Create Form Object.
      initForm() {
            this.bookmark = this.formbuilder.group({
                  url: ['', [Validators.required, Validators.pattern(/^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$/g)]],
                  tag: ['', Validators.required],
                  label: ['', Validators.required],
                  id: ['', Validators.required],
                  categoryid: ['', Validators.required]
            });
      }

      // Patch Value to Form.
      patchValue(bookmark) {
            this.bookmark.patchValue({
                  url: bookmark.url,
                  tag: bookmark.tag,
                  label: bookmark.label,
                  id: bookmark.id,
                  categoryid: bookmark.categoryid
            });
            this.configLoader();
      }

      // Update Bookmark.
      updateBookmark() {
            this.mainservice.updateBookmark(this.bookmark.value).then(data => {
                  console.log("Updated");
                  alert("Update Done");
                  this.mainservice.getAllBookmark(false);
                  this.navigate("bookmark");
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            }).finally(() => {
                  this.configLoader();
            })
      }

      // Return Control of Form.
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
