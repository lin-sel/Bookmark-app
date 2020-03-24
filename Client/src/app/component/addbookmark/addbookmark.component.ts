import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { JsonService } from 'src/app/service/utils/json.service';
import { MainService } from 'src/app/service/main.service';
@Component({
      selector: 'app-addbookmark',
      templateUrl: './addbookmark.component.html',
      styleUrls: ['./addbookmark.component.css']
})
export class AddbookmarkComponent implements OnInit {

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
            this.initForm(id);
      }

      initForm(categoryid) {
            this.bookmark = this.formbuilder.group({
                  url: ['', Validators.required],
                  tag: ['', Validators.required],
                  label: ['', Validators.required],
                  categoryid: [categoryid, Validators.required]
            });
      }

      addBookmark() {
            this.mainservice.addBookmark(this.bookmark.value).then(data => {
                  console.log("New Bookmark Added", data);
                  alert("New Bookmark Added");
                  this.navigate("bookmark");
            }).catch(err => {
                  console.log(err);
                  alert(this.json.fromStringToJSON(err).error);
            })
      }

      get f() {
            return this.bookmark.controls;
      }

      navigate(path: string) {
            this.route.navigate([path])
      }

}
