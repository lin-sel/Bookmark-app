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
      private admin: any;
      constructor(
            private _http: HttpClient,
            private _constant: Constant,
            private _logger: LoggerService,
            private _json: JsonService,
            private _storage: StorageService
      ) { }

      userLogin(data) {
            return new Promise((resolve, reject) => {
                  if (this.user == undefined) {
                        this._http.post(`${this._constant.BASE}/user/login`, data,
                              { headers: this.getToken() }
                        ).toPromise().then((respond: any) => {
                              this.user = respond;
                              // this._logger.log(respond)
                              // this._storage.setByID("token", respond.token)
                              // this._storage.setByID("userid", respond.id)
                              // this._storage.setByID(this._constant.ROLE, "user")
                              this.setSession(respond, "user")
                              resolve()
                        }).catch(err => {
                              reject(err)
                        })
                  }
            })
      }

      setSession(respond, role) {
            this._logger.log(respond);
            this._storage.setByID("token", respond.token)
            this._storage.setByID("userid", respond.id)
            this._storage.setByID(this._constant.ROLE, role)
      }

      adminLogin(data) {
            return new Promise((resolve, reject) => {
                  if (this.admin == undefined) {
                        this._http.post(`${this._constant.BASE}/admin/login`, data,
                              { headers: this.getToken() }
                        ).toPromise().then((respond: any) => {
                              this.user = respond;
                              // this._logger.log(respond)
                              // this._storage.setByID("token", respond.token)
                              // this._storage.setByID("userid", respond.id)
                              // this._storage.setByID(this._constant.ROLE, "admin")
                              this.setSession(respond, "admin")
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

