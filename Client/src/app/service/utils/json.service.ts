import { Injectable } from '@angular/core';

@Injectable({
      providedIn: 'root'
})
export class JsonService {

      constructor() { }

      fromStringToJSON(data) {
            return JSON.parse(data)
      }

      fromJSONToString(data) {
            return JSON.stringify(data)
      }
}
