import { Injectable } from '@angular/core';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { LoggerService } from '../utils/logger.service';
import { JsonService } from '../utils/json.service';
import { StorageService } from '../utils/storage.service';
import { Constant } from '../constant';

@Injectable({
      providedIn: 'root'
})
export class UserService {

      private user: any;
      constructor(
            private _logger: LoggerService,
            private _http: HttpClient,
            private _json: JsonService,
            private _storage: StorageService,
            private _constant: Constant
      ) {
      }


      get(check: boolean) {
            return new Promise((resolve, reject) => {
                  if (!this.user) {
                        this._http.get(`${this._constant.BASE}/${this._storage.getByID("userid")}/user`,
                              { headers: this.getToken() },
                        ).toPromise().then((respond: any) => {
                              this._logger.info(respond)
                              this.user = respond;
                              resolve(respond)
                              console.log(respond);
                        }).catch(err => {
                              this._logger.error(err)
                              reject(err)
                        });
                        return;
                  }
                  resolve(this.user);
            })
      }


      update(data) {
            return new Promise((resolve, reject) => {
                  this._http.put(`${this._constant.BASE}/${this._storage.getByID("userid")}/user`, data,
                        { headers: this.getToken() },
                  ).toPromise().then((respond: any) => {
                        this._logger.info(respond)
                        this.get(false);
                        resolve(respond)
                  }).catch(err => {
                        this._logger.error(err)
                        reject(err)
                  });
            })
      }


      delete(bookmarkid) {
            return new Promise((resolve, reject) => {
                  const headers = new HttpHeaders();
                  this._http.delete(`${this._constant.BASE}/${this._storage.getByID("userid")}/user`
                        , { headers: this.getToken() },
                  ).toPromise().then((respond: any) => {
                        this._logger.info(respond)
                        this.get(false);
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


