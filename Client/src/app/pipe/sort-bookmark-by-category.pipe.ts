import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
      name: 'sortBookmarkByCategory'
})
export class SortBookmarkByCategoryPipe implements PipeTransform {

      transform(categoryid: any, bookmark: any): any {
            if (categoryid == bookmark.id) {
                  return bookmark
            }
            return "";
      }

}
