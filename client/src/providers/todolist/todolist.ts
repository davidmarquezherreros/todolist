import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';


@Injectable()
export class TodolistProvider {

  constructor(public http: HttpClient) {
    console.log('Hello TodolistProvider Provider');
  }

  public getAll(){

  }
}
