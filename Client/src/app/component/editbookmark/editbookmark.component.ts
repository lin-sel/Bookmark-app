import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { JsonService } from 'src/app/service/utils/json.service';
import { BookmarkService } from 'src/app/service/bookmark/bookmark.service';
import { MainService } from 'src/app/service/main.service';

@Component({
      selector: 'app-editbookmark',
      templateUrl: './editbookmark.component.html',
      styleUrls: ['./editbookmark.component.css']
})
export class EditbookmarkComponent implements OnInit {

      private bookmark: FormGroup;
      constructor(
            private activeroute: ActivatedRoute,
            private formbuilder: FormBuilder,
            private json: JsonService,
            private mainservice: MainService,
            private route: Router
      ) { }

      ngOnInit() {
            if (!this.mainservice.authUser()) {
                  alert("PLease Login First.");
                  this.route.navigate(["login"]);
                  return;
            }
            this.getParam();
      }

      getParam() {
            let id = this.activeroute.snapshot.paramMap.get('id');
            let bookmk = this.mainservice.getBookmarkByID(id);
            if (bookmk == undefined) {
                  alert("Not Found");
                  this.navigate("bookmark");
                  return
            }
            this.initForm(bookmk)
      }

      initForm(bookmark) {
            this.bookmark = this.formbuilder.group({
                  url: [bookmark.url, Validators.required],
                  tag: [bookmark.tag, Validators.required],
                  label: [bookmark.label, Validators.required],
                  id: [bookmark.id, Validators.required],
                  categoryid: [bookmark.categoryid, Validators.required]
            });
      }

      updateBookmark() {
            this.mainservice.updateBookmark(this.bookmark.value).then(data => {
                  console.log("Updated");
                  alert("Update Done");
                  this.navigate("bookmark");
            }).catch(err => {
                  console.log(err);
                  alert(err);
            })
      }

      get f() {
            return this.bookmark.controls;
      }

      navigate(path: string) {
            this.route.navigate([path])
      }
}
