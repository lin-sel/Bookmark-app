import { Injectable } from '@angular/core';

@Injectable({
      providedIn: 'root'
})
export class JsonService {

      constructor() { }

      fromStringToJSON(data) {
            let output: any
            try {
                  console.log(data);
                  output = JSON.parse(data)
                  console.log(output);
            }
            catch{
                  return undefined;
            }
            return output;
      }

      fromJSONToString(data) {
            return JSON.stringify(data)
      }
}
