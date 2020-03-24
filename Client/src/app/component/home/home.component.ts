import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BookmarkService } from 'src/app/service/bookmark/bookmark.service';
import { MainService } from 'src/app/service/main.service';
import { JsonService } from 'src/app/service/utils/json.service';

@Component({
      selector: 'app-home',
      templateUrl: './home.component.html',
      styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
      private categories: any[]
      private content: any;
      constructor(
            private mainservice: MainService,
            private router: Router,
            private json: JsonService
      ) { }

      ngOnInit() {
            if (!this.mainservice.authUser()) {
                  alert("PLease Login First.");
                  this.router.navigate(["login"]);
                  return;
            }
            this.content = {};
            this.init();
      }

      init() {
            this.mainservice.getAllBookmark(true).then((respond: any[]) => {
                  this.categories = respond
                  console.log(this.categories);
            }).catch(err => {
                  alert(err.error)
            });
      }

      viewDetail(data) {
            this.content = data
      }

      goToExternalURL() {
            window.open("https://" + this.content['url'], "_blank")
      }

      addBookmark(id: string) {
            this.navigate("addbookmark", id)
      }

      deleteBookmark(id: string) {
            if (confirm("Are you want to delete?")) {
                  this.mainservice.deleteBookmark(id).then((respond: any[]) => {
                        alert("Bookmark Deleted");
                  }).catch(err => {
                        alert(err.error)
                  });
            }
      }

      navigate(path: string, id: string) {
            this.router.navigate([path, id]);
      }

      getKeys() {
            return Object.keys(this.content);
      }

}
