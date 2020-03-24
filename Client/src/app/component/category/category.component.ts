import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MainService } from 'src/app/service/main.service';
import { JsonService } from 'src/app/service/utils/json.service';

@Component({
      selector: 'app-category',
      templateUrl: './category.component.html',
      styleUrls: ['./category.component.css']
})
export class CategoryComponent implements OnInit {

      private buttonname: string = "Update";
      private buttonaction: any;
      private content: any;
      private category: FormGroup;
      private categories: any[];
      constructor(
            private formbuilder: FormBuilder,
            private mainservice: MainService,
            private router: Router,
            private json: JsonService
      ) { }

      ngOnInit() {
            if (!this.mainservice.authUser()) {
                  alert("Please Login First.");
                  this.router.navigate(["login"]);
                  return;
            }
            this.initForm();
            this.categories = [];
            this.getAllCategory();
            this.buttonaction = this.updateCategory;
      }


      // Create Form Object.
      initForm() {
            this.category = this.formbuilder.group({
                  category: ['', Validators.required]
            });
      }


      // Add Data To Form.
      patchForm(data) {
            this.category.patchValue({
                  category: data.category
            });
      }


      // Get All Category.
      getAllCategory() {
            this.mainservice.getAllCategory(true).then((data: any[]) => {
                  this.categories = data;
                  console.log(this.categories);
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            })
      }

      // Patch Existing Category Data to Form For Update.
      setContent(data) {
            this.patchForm(data)
            this.content = data;
      }


      // Update Selected Category.
      updateCategory() {
            console.log(this.content);
            this.mainservice.updateCategoy(this.category.value, this.content.id).then(data => {
                  alert("Updated");
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            })
      }

      // Delete Selected Category.
      deleteCategory() {
            this.mainservice.deleteCategory(this.content.id).then(data => {
                  alert("Deleted Successfully");
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            })
      }

      // Add New Category.
      addCategory() {
            this.mainservice.addCategory(this.category.value).then(data => {
                  alert("Added");
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            })
      }

      // Add Action to Button for Add Category Event.
      addAction(name: string) {
            this.buttonname = name;
            this.buttonaction = this.addCategory;
      }


      // Return Form Controls.
      get f() {
            return this.category.controls;
      }

      // Error Parser.
      errorParser(err) {
            let er = this.json.fromStringToJSON(err.error);
            if (er != undefined) {
                  return er.error;
            }
            return err.error;
      }

      // Check Session Expire and Perform Accordingly
      isSessionExpire(s: string) {
            console.log(this.mainservice.isSessionExpire(s))
            if (this.mainservice.isSessionExpire(s)) {
                  this.router.navigate(["login"]);
            }
      }

}
