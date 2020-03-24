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
                  if (this.categorywithbookmark.length == 0 || !check) {
                        // let headers = new HttpHeaders().set('token', "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJc3N1ZWRBdCI6MTU4NTA2MTgxNSwidXNlcklEIjoiOTE1NWY0NTUtNjc4MS00NDM4LTg3YWYtNDkyY2MzOTI1NzE2IiwidXNlcm5hbWUiOiJuaWwifQ.Sr6vdRn6jn6rWMXQoEGvOuPNIv1i_MQDxjpB7l_bxgI");
                        this._http.get(`${this._constant.BASE}/${this._storage.getByID("userid")}/category`,
                              // { headers: headers },
                        ).toPromise().then((respond: any) => {
                              this._logger.info(respond)
                              this.categorywithbookmark = respond;
                              resolve(respond)
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
                  this._http.put(`${this._constant.BASE}/${this._storage.getByID("userid")}/bookmark/${data.id}`, data
                        // {
                        //       headers:
                        //       {
                        //             'token': `${this._config.getToken()}`,
                        //             'Content-Type': 'application/json',
                        //       }
                        // }
                  ).toPromise().then((respond: any) => {
                        this._logger.info(respond)
                        this.getAll(false);
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

      addBookmark(data) {
            return new Promise((resolve, reject) => {
                  const headers = new HttpHeaders();
                  this._http.post(`${this._constant.BASE}/${this._storage.getByID("userid")}/bookmark`, data
                        // {
                        //       headers:
                        //       {
                        //             'token': `${this._config.getToken()}`,
                        //             'Content-Type': 'application/json',
                        //       }
                        // }
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
                        // {
                        //       headers:
                        //       {
                        //             'token': `${this._config.getToken()}`,
                        //             'Content-Type': 'application/json',
                        //       }
                        // }
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


}
