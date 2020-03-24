import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MainService } from 'src/app/service/main.service';

@Component({
      selector: 'app-category',
      templateUrl: './category.component.html',
      styleUrls: ['./category.component.css']
})
export class CategoryComponent implements OnInit {

      private buttonname: string = "Update";
      private content: any;
      private category: FormGroup;
      private categories: any[];
      constructor(
            private formbuilder: FormBuilder,
            private mainservice: MainService,
            private router: Router
      ) { }

      ngOnInit() {
            if (!this.mainservice.authUser()) {
                  alert("PLease Login First.");
                  this.router.navigate(["login"]);
                  return;
            }

            this.initForm();
            this.categories = [];
            this.getAllCategory();
      }

      initForm() {
            this.category = this.formbuilder.group({
                  category: ['', Validators.required]
            });
      }

      patchForm(data) {
            this.category.patchValue({
                  category: data.category
            });
      }

      getAllCategory() {
            this.mainservice.getAllCategory(true).then((data: any[]) => {
                  this.categories = data;
                  console.log(this.categories);
            }).catch(err => {
                  alert(err);
            })
      }

      setContent(data) {
            this.patchForm(data)
            this.content = data;
      }

      updateCategory() {
            console.log(this.content);
            this.mainservice.updateCategoy(this.category.value, this.content.id).then(data => {
                  alert("Updated");
            }).catch(err => {
                  alert(err);
            })
      }

      deleteCategory() {
            this.mainservice.deleteCategory(this.content.id).then(data => {
                  alert("Deleted Successfully");
            }).catch(err => {
                  alert(err);
            })
      }

      addCategory(name: string) {
            this.buttonname = name;
            this.mainservice.addCategory(this.category.value).then(data => {
                  alert("Added");
            }).catch(err => {
                  alert(err);
            })
      }

      get f() {
            return this.category.controls;
      }

}
