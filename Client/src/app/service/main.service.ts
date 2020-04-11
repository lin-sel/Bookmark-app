import { Injectable } from '@angular/core';
import { BookmarkService } from './bookmark/bookmark.service';
import { CategoryService } from './category/category.service';
import { StorageService } from './utils/storage.service';
import { LoginService } from './login/login.service';
import { RegisterService } from './register/register.service';
import { UserService } from './user/user.service';
import { Constant } from './constant';

@Injectable({
      providedIn: 'root'
})
export class MainService {

      constructor(
            private bookmark: BookmarkService,
            private category: CategoryService,
            private storage: StorageService,
            private login: LoginService,
            private register: RegisterService,
            private user: UserService,
            private constant: Constant
      ) {
            // this.authUser()
      }

      getAllBookmark(check: boolean, pagesize: number, pagenumber: number): Promise<any> {
            return this.bookmark.getAll(check, pagesize, pagenumber)
      }

      getBookmarkByID(id: string, pagesize, pagenumber) {
            return this.bookmark.getBookmarkByCategoryID(id, pagesize, pagenumber);
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

      authUser(role: string) {
            console.log(this.storage.getByID('userid'))
            if (this.storage.getByID('userid') == null || this.storage.getByID('token') == null) {
                  if (this.storage.getByID(this.constant.ROLE).toLowerCase().includes(role.toLowerCase()))
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


      userLogin(data) {
            return this.login.userLogin(data);
      }

      adminLogin(data) {
            return this.login.adminLogin(data);
      }

      userRegister(data) {
            return this.register.register(data);
      }

      getUser() {
            return this.user.get(true)
      }

      updateUser(data: any) {
            return this.user.update(data)
      }

      deleteUser(userid) {
            return this.user.delete(userid)
      }


}
