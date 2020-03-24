import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { LoggerService } from '../utils/logger.service';
import { JsonService } from '../utils/json.service';
import { StorageService } from '../utils/storage.service';
import { Constant } from '../constant';

@Injectable({
      providedIn: 'root'
})
export class LoginService {

      constructor(
            private _http: HttpClient,
            private _constant: Constant,
            private _logger: LoggerService,
            private _json: JsonService,
            private _storage: StorageService
      ) { }

      login(data) {
            return new Promise((resolve, reject) => {
                  this._http.post(`${this._constant.BASE}/login`, data).toPromise().then((respond: any) => {
                        this._logger.log(respond)
                        this._storage.setByID("token", respond.token)
                        this._storage.setByID("userid", respond.id)
                        resolve()
                  }).catch(err => {
                        reject(err.error)
                        this._logger.error(this._json.fromStringToJSON(err.error).error)
                  })
            })
      }
}
