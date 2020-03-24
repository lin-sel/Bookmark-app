import { Component, OnInit } from '@angular/core';
import { StorageService } from 'src/app/service/utils/storage.service';
import { MainService } from 'src/app/service/main.service';
import { Router } from '@angular/router';

@Component({
      selector: 'app-navbar',
      templateUrl: './navbar.component.html',
      styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

      constructor(
            private storage: StorageService,
            private mainservice: MainService,
            private router: Router
      ) { }

      ngOnInit() {
            // if (!this.mainservice.authUser()) {
            //       alert("PLease Login First.");
            //       this.router.navigate(["login"]);
            //       return;
            // }

            // this.storage.setByID("userid", "9155f455-6781-4438-87af-492cc3925716")
            // this.storage.setByID("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJc3N1ZWRBdCI6MTU4NTA2MTgxNSwidXNlcklEIjoiOTE1NWY0NTUtNjc4MS00NDM4LTg3YWYtNDkyY2MzOTI1NzE2IiwidXNlcm5hbWUiOiJuaWwifQ.Sr6vdRn6jn6rWMXQoEGvOuPNIv1i_MQDxjpB7l_bxgI")
      }

      signOut() {
            this.storage.clear();
      }



}
