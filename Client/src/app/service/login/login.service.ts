import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { LoggerService } from '../utils/logger.service';
import { JsonService } from '../utils/json.service';
import { StorageService } from '../utils/storage.service';
import { Constant } from '../constant';

@Injectable({
      providedIn: 'root'
})
export class LoginService {

      private user: any;
      constructor(
            private _http: HttpClient,
            private _constant: Constant,
            private _logger: LoggerService,
            private _json: JsonService,
            private _storage: StorageService
      ) { }

      login(data) {
            return new Promise((resolve, reject) => {
                  if (this.user == undefined) {
                        this._http.post(`${this._constant.BASE}/login`, data,
                              { headers: this.getToken() }
                        ).toPromise().then((respond: any) => {
                              this.user = respond;
                              this._logger.log(respond)
                              this._storage.setByID("token", respond.token)
                              this._storage.setByID("userid", respond.id)
                              resolve()
                        }).catch(err => {
                              reject(err)
                        })
                  }
            })
      }

      getToken(): HttpHeaders {
            return new HttpHeaders().set('token', `${this._storage.getByID('token')}`);
      }


      getUser() {
            return this.user;
      }
}

