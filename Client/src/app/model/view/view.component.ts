import { Component, OnInit, Input, AfterViewInit } from '@angular/core';

@Component({
      selector: 'app-view',
      templateUrl: './view.component.html',
      styleUrls: ['./view.component.css']
})
export class ViewComponent {

      // private keys: string[]
      @Input('content') content: any;
      @Input('keys') keys: any[];
      constructor() { }


}
