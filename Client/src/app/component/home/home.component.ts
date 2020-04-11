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
      public bookmarks: any[]
      public categorys: any[]
      public currentcategory: any
      public currentbookmark: any
      public content: any;
      public loader: any;
      public pagesize: number
      public pagenumber: number
      public totalpage: number
      public data: any[];
      public role: string = "user"
      constructor(
            private mainservice: MainService,
            private router: Router,
            public util: UtilService
      ) {
            this.bookmarks = [];
            this.categorys = [];
            this.loader = {
                  loader: "loader",
                  body: "hide"
            }
            this.data = [];
            this.pagesize = 2;
            this.pagenumber = 1;
            this.totalpage = 1;
      }

      ngOnInit() {
            if (!this.mainservice.authUser(this.role)) {
                  alert("Please Login First.");
                  this.router.navigate(["login"]);
                  return;
            }
            this.content = {};
            this.init();
      }

      init() {
            this.getAllCategory();
            // this.getAllBookmark();
      }

      getAllCategory() {
            this.mainservice.getAllCategory(true).then((data: any[]) => {
                  this.categorys = data;
                  this.currentcategory = data[0]
                  console.log("category list => ", this.categorys);
                  console.log("Current Category =>", this.currentcategory)
                  this.setDataObject(data)
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(err);
                  console.log(error)
                  this.isSessionExpire(error);
            }).finally(() => {
                  this.configLoader();
            })
      }

      setDataObject(data: any[]) {
            for (let index = 0; index < data.length; index++) {
                  this.data.push({
                        category: data[index],
                        pagenumber: 1,
                        totalpage: 1,
                        listofbookmarks: Array()
                  });
            }
            this.getBookmarkByCategoryID();
            console.log("After Category List set => ", this.data)
      }

      // getAllBookmark() {
      //       this.configLoader();
      //       if (this.totalpage >= this.pagenumber) {
      //             this.mainservice.getAllBookmark(true, this.pagesize, this.pagenumber).then((response: any) => {
      //                   console.log(response);
      //                   if (this.pagenumber == 1) {
      //                         this.bookmarks = response.listofbookmark;
      //                   } else {
      //                         this.bookmarks.push(response.listofbookmark);
      //                   }
      //                   this.pagenumber = response.totalpage
      //                   this.totalpage = response.totalpage
      //                   this.pagenumber++
      //                   console.log(this.bookmarks)
      //             }).catch(err => {
      //                   let error = this.errorParser(err);
      //                   alert(error);
      //                   console.log(error)
      //                   this.isSessionExpire(error);
      //             }).finally(() => {
      //                   this.configLoader();
      //             });
      //       }
      // }

      // Set Current Category.
      setCurrentCategory(category) {
            console.log("Current Category to be set =>", category)
            this.currentcategory = category.category;
            this.getBookmarkByCategoryID();
            console.log("Current category =>", this.currentcategory);
      }

      // Get Bookmark by Category ID.
      getBookmarkByCategoryID() {
            // this.configLoader();
            let data = this.getBookmark()
            if (data.totalpage >= data.pagenumber) {
                  this.mainservice.getBookmarkByID(this.currentcategory.id, this.pagesize, this.pagenumber).then((response: any) => {
                        console.log("response data: => ", response);
                        if (data.pagenumber == 1) {
                              data.listofbookmarks = response.listofbookmark;
                              data.pagenumber = data.pagenumber++
                              data.totalpage = response.totalpage
                        } else {
                              data.listofbookmarks.push(response.listofbookmark);
                              data.pagenumber = data.pagenumber++
                              data.totalpage = response.totalpage
                        }
                        console.log("After Update =>", this.data)
                  }).catch(err => {
                        let error = this.errorParser(err);
                        alert(error);
                        console.log(error)
                        this.isSessionExpire(error);
                  }).finally(() => {
                        // this.configLoader();
                  });
            }
      }

      // Get Whole Object from data list by Category ID.
      getBookmark() {
            for (let index = 0; index < this.data.length; index++) {
                  if (this.currentcategory.id == this.data[index].category.id) {
                        return this.data[index]
                  }
            }
            return undefined;
      }

      // Get Bookmark list in data list
      getBookmarkList(): any[] {
            let obj = this.getBookmark();
            if (obj == undefined) {
                  return null;
            }
            return obj.listofbookmarks
      }


      getNextLabel() {
            let obj = this.getBookmark();
            if (obj == undefined || obj.totalpage >= obj.pagenumber) {
                  return false
            }
            return true
      }


      // Set Data to Content variable for View.
      viewDetail(data) {
            this.content = data
      }


      // Expose to External URL.
      goToExternalURL() {
            if (this.content['url'].includes("http")) {
                  alert(window.open(this.content['url'], "_blank"))
                  return
            }
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
