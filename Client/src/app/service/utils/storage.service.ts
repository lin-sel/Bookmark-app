import { Injectable } from '@angular/core';

@Injectable({
      providedIn: 'root'
})
export class StorageService {

      constructor() { }

      setByID(key: string, value) {
            sessionStorage.setItem(key, value);
      }

      getByID(key: string) {
            return sessionStorage.getItem(key);
      }

      remove(key: string) {
            sessionStorage.removeItem(key);
      }

      clear() {
            sessionStorage.clear();
      }
}
