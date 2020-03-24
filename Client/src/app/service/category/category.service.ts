import { Injectable } from '@angular/core';
import { LoggerService } from '../utils/logger.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { JsonService } from '../utils/json.service';
import { StorageService } from '../utils/storage.service';
import { Constant } from '../constant';
@Injectable({
      providedIn: 'root'
})
export class CategoryService {

      private categories: any[];
      constructor(
            private _logger: LoggerService,
            private _http: HttpClient,
            private _json: JsonService,
            private _storage: StorageService,
            private _constant: Constant
      ) {
            this.categories = [];
      }


      getAll(check: boolean) {
            return new Promise((resolve, reject) => {
                  if (this.categories.length == 0 || !check) {
                        const header = new HttpHeaders();
                        this._http.get(`${this._constant.BASE}/${this._storage.getByID("userid")}/category`
                              // {
                              //       headers:
                              //       {
                              //             'token': `${this._config.getToken()}`,
                              //             'Content-Type': 'application/json',
                              //       }
                              // }
                        ).toPromise().then((respond: any) => {
                              this._logger.info(respond)
                              this.categories = respond;
                              resolve(respond)
                        }).catch(err => {
                              this._logger.error(err)
                              reject(err)
                        });
                        return;
                  }
                  resolve(this.categories);
            })
      }


      update(data, id) {
            return new Promise((resolve, reject) => {
                  const header = new HttpHeaders();
                  this._http.put(`${this._constant.BASE}/${this._storage.getByID("userid")}/category/${id}`, data
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
            console.log(this.categories.length)
            for (let index = 0; index < this.categories.length; index++) {
                  if (this.categories[index].id == id) {
                        return this.categories[index];
                  }
            }
            return undefined;
      }

      addCategory(data) {
            return new Promise((resolve, reject) => {
                  const header = new HttpHeaders();
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

      deleteCategory(categoryid) {
            return new Promise((resolve, reject) => {
                  const header = new HttpHeaders();
                  this._http.delete(`${this._constant.BASE}/${this._storage.getByID("userid")}/category/${categoryid}`
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
