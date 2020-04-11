import { Component, OnInit } from '@angular/core';
import { MainService } from 'src/app/service/main.service';

@Component({
      selector: 'app-admindashboard',
      templateUrl: './admindashboard.component.html',
      styleUrls: ['./admindashboard.component.css']
})
export class AdmindashboardComponent implements OnInit {

      private role: string = "admin"
      constructor(
            private mainservice: MainService
      ) { }

      ngOnInit() {
            if (!this.mainservice.authUser(this.role)) {
                  alert("Invalid user")
                  return;
            }
      }

}
