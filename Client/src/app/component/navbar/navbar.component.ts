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


      }

      signOut() {
            this.storage.clear();
      }



}
