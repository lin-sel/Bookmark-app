import { Injectable } from '@angular/core';
import { JsonService } from './json.service';
import { Router } from '@angular/router';
import { MainService } from '../main.service';

@Injectable({
      providedIn: 'root'
})
export class UtilService {

      constructor(
            private json: JsonService,
            private router: Router,
            private mainservice: MainService
      ) { }

      //ConfigLoader set loader.
      configLoader(obj, param?: any) {
            if (obj.loader == "loader") {
                  obj.loader = "hide"
                  obj.body = "visible"
                  return;
            }
            obj.loader = "loader";
            obj.body = "hide";
      }

      // Error Parser.
      errorParser(err) {
            let er = this.json.fromStringToJSON(err.error);
            console.log(er)
            if (er != undefined) {
                  return er.error;
            }
            return err.error;
      }

      // Navigate to Another URL.
      navigate(path: string) {
            this.router.navigate([path]);
      }

      // Navigate to Another URL.
      navigateWithParam(path: string, param: any) {
            this.router.navigate([path, param]);
      }


      // Check Session Expire and Perform Accordingly
      isSessionExpire(s: string) {
            console.log(this.mainservice.isSessionExpire(s))
            if (this.mainservice.isSessionExpire(s)) {
                  this.router.navigate(["login"]);
            }
      }

      reload() {
            document.location.reload();
      }
}
