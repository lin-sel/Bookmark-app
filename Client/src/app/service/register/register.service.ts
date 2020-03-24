import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { LoggerService } from '../utils/logger.service';
import { JsonService } from '../utils/json.service';
import { Constant } from '../constant';

@Injectable({
      providedIn: 'root'
})
export class RegisterService {

      constructor(
            private _constant: Constant,
            private _http: HttpClient,
            private _logger: LoggerService,
            private _json: JsonService
      ) { }

      register(data) {
            return new Promise((resolve, reject) => {
                  this._http.post(`${this._constant.BASE}/register`, data,
                        {
                              headers: {
                                    'token': ""
                              }
                        }
                  ).toPromise().then(respond => {
                        resolve()
                        this._logger.info("Register Successful")
                  }).catch(err => {
                        reject(err.error)
                        this._logger.error(this._json.fromStringToJSON(err.error).error)
                  });
            });
      }


}
