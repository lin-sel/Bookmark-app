import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MainService } from 'src/app/service/main.service';
import { JsonService } from 'src/app/service/utils/json.service';
import { UtilService } from 'src/app/service/utils/util.service';

@Component({
      selector: 'app-category',
      templateUrl: './category.component.html',
      styleUrls: ['./category.component.css']
})
export class CategoryComponent implements OnInit {

      public loader: any;
      public buttonname: string = "Update";
      public buttonaction: any;
      public content: any;
      public category: FormGroup;
      public categories: any[];
      private role: string = "user";
      constructor(
            private formbuilder: FormBuilder,
            private mainservice: MainService,
            public util: UtilService
      ) {
            this.loader = {
                  loader: "loader",
                  body: "hide"
            }
      }

      ngOnInit() {
            if (!this.mainservice.authUser(this.role)) {
                  alert("Please Login First.");
                  this.util.navigate("login")
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
            }).finally(() => {
                  this.configLoader();
            })
      }

      // Patch Existing Category Data to Form For Update.
      setContent(data) {
            this.patchForm(data)
            this.content = data;
      }


      // Update Selected Category.
      updateCategory() {
            this.configLoader();
            console.log(this.content);
            this.mainservice.updateCategoy(this.category.value, this.content.id).then(data => {
                  alert("Updated");
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            }).finally(() => {
                  this.configLoader();
                  this.util.reload();
            })
      }

      // Delete Selected Category.
      deleteCategory() {
            this.configLoader();
            this.mainservice.deleteCategory(this.content.id).then(data => {
                  alert("Deleted Successfully");
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            }).finally(() => {
                  this.configLoader();
                  this.util.reload();
            })
      }

      // Add New Category.
      addCategory() {
            this.configLoader();
            this.mainservice.addCategory(this.category.value).then(data => {
                  alert("Added");
            }).catch(err => {
                  let error = this.errorParser(err);
                  alert(error);
                  console.log(error)
                  this.isSessionExpire(error);
            }).finally(() => {
                  this.configLoader();
                  this.util.reload();
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
            // let er = this.json.fromStringToJSON(err.error);
            // if (er != undefined) {
            //       return er.error;
            // }
            // return err.error;
            return this.util.errorParser(err);
      }

      // Check Session Expire and Perform Accordingly
      isSessionExpire(s: string) {
            // console.log(this.mainservice.isSessionExpire(s))
            // if (this.mainservice.isSessionExpire(s)) {
            //       this.router.navigate(["login"]);
            // }
            this.util.isSessionExpire(s);
      }

      configLoader() {
            this.util.configLoader(this.loader)
      }

}
