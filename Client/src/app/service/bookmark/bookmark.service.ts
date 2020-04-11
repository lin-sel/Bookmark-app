import { Injectable } from '@angular/core';
import { LoggerService } from '../utils/logger.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { JsonService } from '../utils/json.service';
import { StorageService } from '../utils/storage.service';
import { Constant } from '../constant';

@Injectable({
      providedIn: 'root'
})
export class BookmarkService {

      private bookmarklist: any[];
      constructor(
            private _logger: LoggerService,
            private _http: HttpClient,
            private _json: JsonService,
            private _storage: StorageService,
            private _constant: Constant
      ) {
            this.bookmarklist = [];
      }


      getAll(check: boolean, pagesize: number, pagenumber: number) {
            return new Promise((resolve, reject) => {
                  this._http.get(`${this._constant.BASE}/user/${this._storage.getByID("userid")}/bookmark/${pagesize}/${pagenumber}`,
                        { headers: this.getToken() },
                  ).toPromise().then((respond: any) => {
                        this._logger.info(respond)
                        if (pagenumber == 1) {
                              this.bookmarklist = respond.listofbookmark;
                        } else {
                              this.bookmarklist.push(respond.listofbookmark);
                        }
                        this._logger.info(this.bookmarklist)
                        resolve(respond);
                  }).catch(err => {
                        this._logger.error(err)
                        reject(err)
                  });
            })
      }


      update(data) {
            return new Promise((resolve, reject) => {
                  const headers = new HttpHeaders();
                  this._http.put(`${this._constant.BASE}/user/${this._storage.getByID("userid")}/bookmark/${data.id}`, data,
                        { headers: this.getToken() },
                  ).toPromise().then((respond: any) => {
                        this._logger.info(respond)
                        this.updateByID(data)
                        resolve(respond)
                  }).catch(err => {
                        this._logger.error(err)
                        reject(err)
                  });
            })
      }

      getBookmarkByCategoryID(id, pagesize, pagenumber) {
            return new Promise((resolve, reject) => {
                  this._http.get(`${this._constant.BASE}/user/${this._storage.getByID("userid")}/bookmark/category/${id}/${pagesize}/${pagenumber}`,
                        { headers: this.getToken() },
                  ).toPromise().then((respond: any) => {
                        this._logger.info(respond)
                        if (pagenumber == 1) {
                              this.bookmarklist = respond.listofbookmark;
                        } else {
                              this.bookmarklist.push(respond.listofbookmark);
                        }
                        this._logger.info(this.bookmarklist)
                        resolve(respond);
                  }).catch(err => {
                        this._logger.error(err)
                        reject(err)
                  });
            })
      }

      getByID(id: string) {
            console.log(this.bookmarklist.length)
            for (let index = 0; index < this.bookmarklist.length; index++) {
                  // for (let i = 0; i < this.bookmarklist[index].bookmarks.length; i++) {
                  if (this.bookmarklist[index].id == id) {
                        return this.bookmarklist[index];
                  }
                  // }
            }
            return undefined;
      }

      updateByID(id: string) {
            console.log(this.bookmarklist.length)
            for (let index = 0; index < this.bookmarklist.length; index++) {
                  // for (let i = 0; i < this.bookmarklist[index].bookmarks.length; i++) {
                  if (this.bookmarklist[index].id == id) {
                        return this.bookmarklist[index];
                  }
                  // }
            }
            return undefined;
      }

      addBookmark(data) {
            return new Promise((resolve, reject) => {
                  const headers = new HttpHeaders();
                  this._http.post(`${this._constant.BASE}/user/${this._storage.getByID("userid")}/bookmark`, data
                        , { headers: this.getToken() },
                  ).toPromise().then((respond: any) => {
                        this._logger.info(respond)
                        resolve(respond)
                  }).catch(err => {
                        this._logger.error(err)
                        reject(err)
                  });
            });
      }

      deleteBookmark(bookmarkid) {
            return new Promise((resolve, reject) => {
                  const headers = new HttpHeaders();
                  this._http.delete(`${this._constant.BASE}/user/${this._storage.getByID("userid")}/bookmark/${bookmarkid}`
                        , { headers: this.getToken() },
                  ).toPromise().then((respond: any) => {
                        this._logger.info(respond)
                        resolve(respond)
                  }).catch(err => {
                        this._logger.error(err)
                        reject(err)
                  });
            });
      }


      getToken(): HttpHeaders {
            return new HttpHeaders().set('token', `${this._storage.getByID('token')}`);
      }

}


