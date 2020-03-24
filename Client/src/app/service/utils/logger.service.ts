import { Injectable } from '@angular/core';

@Injectable({
      providedIn: 'root'
})
export class LoggerService {

      constructor() { }

      worn(data) {
            console.warn(data)
      }

      info(data) {
            console.info(data)
      }

      log(data) {
            console.log(data)
      }

      error(data) {
            console.error(data)
      }
}
