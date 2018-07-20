import { Component } from '@angular/core';

/**
 * Generated class for the ToDoListComponent component.
 *
 * See https://angular.io/api/core/Component for more info on Angular
 * Components.
 */
@Component({
  selector: 'to-do-list',
  templateUrl: 'to-do-list.html'
})
export class ToDoListComponent {

  text: string;

  constructor() {
    console.log('Hello ToDoListComponent Component');
    this.text = 'Hello World';
  }

}
