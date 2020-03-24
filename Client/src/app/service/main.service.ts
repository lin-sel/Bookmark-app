import { Injectable } from '@angular/core';
import { BookmarkService } from './bookmark/bookmark.service';
import { CategoryService } from './category/category.service';
import { StorageService } from './utils/storage.service';
import { LoginService } from './login/login.service';
import { RegisterService } from './register/register.service';

@Injectable({
      providedIn: 'root'
})
export class MainService {

      constructor(
            private bookmark: BookmarkService,
            private category: CategoryService,
            private storage: StorageService,
            private login: LoginService,
            private register: RegisterService
      ) {
            this.authUser()
      }

      getAllBookmark(check: boolean): Promise<any> {
            return this.bookmark.getAll(check)
      }

      getBookmarkByID(id: string) {
            return this.bookmark.getByID(id);
      }

      deleteBookmark(id: string) {
            return this.bookmark.deleteBookmark(id);
      }

      updateBookmark(data) {
            return this.bookmark.update(data);
      }

      addBookmark(data) {
            return this.bookmark.addBookmark(data);
      }

      authUser() {
            if (this.storage.getByID('userid') == null || this.storage.getByID('token') == null) {
                  return false;
            }
            else {
                  return true;
            }
      }

      isSessionExpire(s: string): boolean {
            let msg: string = s.toLowerCase();
            return msg.includes("session expire")
      }

      getAllCategory(check: boolean): Promise<any> {
            return this.category.getAll(check)
      }

      getCategoryByID(id: string) {
            return this.category.getByID(id);
      }

      deleteCategory(id: string) {
            return this.category.deleteCategory(id);
      }

      updateCategoy(data, id: string) {
            return this.category.update(data, id);
      }

      addCategory(data) {
            return this.category.addCategory(data);
      }


      appLogin(data) {
            return this.login.login(data);
      }

      userRegister(data) {
            return this.register.register(data);
      }


}
