import { Component, OnInit, Input } from '@angular/core';

@Component({
      selector: 'app-popmodel',
      templateUrl: './popmodel.component.html',
      styleUrls: ['./popmodel.component.css']
})
export class PopmodelComponent implements OnInit {
      public keys: string[]
      @Input('content') content: any;
      constructor() { }

      ngOnInit() {
            this.getKeys()
      }

      getKeys() {
            this.keys = Object.keys(this.content)
            console.log(this.content)
      }

}
