import { Injectable } from '@angular/core';
import { JsonService } from './json.service';
import { Router } from '@angular/router';
import { MainService } from '../main.service';
import { Constant } from '../constant';

@Injectable({
      providedIn: 'root'
})
export class UtilService {

      constructor(
            private json: JsonService,
            private router: Router,
            private mainservice: MainService,
            private constant: Constant
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

      getLocation() {
            let loc = this.replace(document.location.toString(), this.constant.CLIENTURL);
            let prev = this.replace(document.referrer.toString(), this.constant.CLIENTURL);
            loc = loc.replace(/[\/0-9a-zA-Z]{9}\-[0-9a-zA-Z]{4}\-[0-9a-zA-Z]{4}\-[0-9a-zA-Z]{4}\-[0-9a-zA-Z]{12}/g, "")
            prev = prev.replace(/[\/0-9a-zA-Z]{9}\-[0-9a-zA-Z]{4}\-[0-9a-zA-Z]{4}\-[0-9a-zA-Z]{4}\-[0-9a-zA-Z]{12}/g, "")
            if (loc == prev) {
                  return loc
            }
            let path = prev + loc
            return path;
      }

      replace(s: string, substr: string): string {
            return s.replace(substr, "")
      }
}
