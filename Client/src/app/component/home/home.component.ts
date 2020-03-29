import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BookmarkService } from 'src/app/service/bookmark/bookmark.service';
import { MainService } from 'src/app/service/main.service';
import { JsonService } from 'src/app/service/utils/json.service';
import { UtilService } from 'src/app/service/utils/util.service';

@Component({
      selector: 'app-home',
      templateUrl: './home.component.html',
      styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
      public categories: any[]
      public content: any;
      public loader: any;
      constructor(
            private mainservice: MainService,
            private router: Router,
            private util: UtilService
      ) {
            this.categories = [];
            this.loader = {
                  loader: "loader",
                  body: "hide"
            }
      }

      ngOnInit() {
            if (!this.mainservice.authUser()) {
                  alert("Please Login First.");
                  this.router.navigate(["login"]);
                  return;
            }
            this.content = {};
            this.init();
      }

      init() {
            this.getAllBookmark();
      }

      getAllBookmark() {
            this.mainservice.getAllBookmark(true).then((respond: any[]) => {
                  this.categories = respond
                  console.log(this.categories);
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            }).finally(() => {
                  this.configLoader();
            });
      }


      // Set Data to Content variable for View.
      viewDetail(data) {
            this.content = data
      }


      // Expose to External URL.
      goToExternalURL() {
            window.open("https://" + this.content['url'], "_blank")
      }

      // Navigate to another page for add Bookmark.
      addBookmark(id: string) {
            this.navigate("addbookmark", id)
      }

      // Delete Bookmark.
      deleteBookmark(id: string) {
            if (confirm("Are you want to delete?")) {
                  this.mainservice.deleteBookmark(id).then((respond: any[]) => {
                        alert("Bookmark Deleted");
                  }).catch(err => {
                        let error = this.errorParser(err);
                        alert(error);
                        console.log(error)
                        this.isSessionExpire(error);
                  }).finally(() => {
                        this.configLoader();
                        this.util.reload();
                  });
            }
      }


      // Navigate to another url.
      navigate(path: string, id: string) {
            this.util.navigateWithParam(path, id);
            // this.router.navigate([path, id]);
      }

      // Return All Keys Of Object.
      getKeys() {
            return Object.keys(this.content);
      }

      // Error Parser.
      errorParser(err) {
            // let er = this.json.fromStringToJSON(err.error);
            // if (er != undefined) {
            //       return er.error;
            // }
            // return err.error;
            return this.util.errorParser(err)
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
