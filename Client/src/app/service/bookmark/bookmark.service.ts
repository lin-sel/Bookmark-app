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

      private categorywithbookmark: any[];
      constructor(
            private _logger: LoggerService,
            private _http: HttpClient,
            private _json: JsonService,
            private _storage: StorageService,
            private _constant: Constant
      ) {
            this.categorywithbookmark = [];
      }


      getAll(check: boolean) {
            return new Promise((resolve, reject) => {
                  if (!check || this.categorywithbookmark.length == 0) {
                        this._http.get(`${this._constant.BASE}/${this._storage.getByID("userid")}/category`,
                              { headers: this.getToken() },
                        ).toPromise().then((respond: any) => {
                              this._logger.info(respond)
                              this.categorywithbookmark = respond;
                              resolve(this.categorywithbookmark);
                              console.log(respond);
                        }).catch(err => {
                              this._logger.error(err)
                              reject(err)
                        });
                        return;
                  }
                  resolve(this.categorywithbookmark);
            })
      }


      update(data) {
            return new Promise((resolve, reject) => {
                  const headers = new HttpHeaders();
                  this._http.put(`${this._constant.BASE}/${this._storage.getByID("userid")}/bookmark/${data.id}`, data,
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

      getByID(id: string) {
            console.log(this.categorywithbookmark.length)
            for (let index = 0; index < this.categorywithbookmark.length; index++) {
                  for (let i = 0; i < this.categorywithbookmark[index].bookmarks.length; i++) {
                        if (this.categorywithbookmark[index].bookmarks[i].id == id) {
                              return this.categorywithbookmark[index].bookmarks[i];
                        }
                  }
            }
            return undefined;
      }

      updateByID(bookmark: any) {
            for (let index = 0; index < this.categorywithbookmark.length; index++) {
                  for (let i = 0; i < this.categorywithbookmark[index].bookmarks.length; i++) {
                        if (this.categorywithbookmark[index].bookmarks[i].id == bookmark.id) {
                              this.categorywithbookmark[index].bookmarks[i] = bookmark;
                        }
                  }
            }
      }

      addBookmark(data) {
            return new Promise((resolve, reject) => {
                  const headers = new HttpHeaders();
                  this._http.post(`${this._constant.BASE}/${this._storage.getByID("userid")}/bookmark`, data
                        , { headers: this.getToken() },
                  ).toPromise().then((respond: any) => {
                        this._logger.info(respond)
                        this.getAll(false);
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
                  this._http.delete(`${this._constant.BASE}/${this._storage.getByID("userid")}/bookmark/${bookmarkid}`
                        , { headers: this.getToken() },
                  ).toPromise().then((respond: any) => {
                        this._logger.info(respond)
                        this.getAll(false);
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


