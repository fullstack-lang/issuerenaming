// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { WaldoDB } from './waldo-db';

// insertion point for imports
import { FooDB } from './foo-db'

@Injectable({
  providedIn: 'root'
})
export class WaldoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  WaldoServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private waldosUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.waldosUrl = origin + '/api/github.com/fullstack-lang/issuerenaming/go/v1/waldos';
  }

  /** GET waldos from the server */
  getWaldos(): Observable<WaldoDB[]> {
    return this.http.get<WaldoDB[]>(this.waldosUrl)
      .pipe(
        tap(_ => this.log('fetched waldos')),
        catchError(this.handleError<WaldoDB[]>('getWaldos', []))
      );
  }

  /** GET waldo by id. Will 404 if id not found */
  getWaldo(id: number): Observable<WaldoDB> {
    const url = `${this.waldosUrl}/${id}`;
    return this.http.get<WaldoDB>(url).pipe(
      tap(_ => this.log(`fetched waldo id=${id}`)),
      catchError(this.handleError<WaldoDB>(`getWaldo id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new waldo to the server */
  postWaldo(waldodb: WaldoDB): Observable<WaldoDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _Foo_Waldos_reverse = waldodb.Foo_Waldos_reverse
    waldodb.Foo_Waldos_reverse = new FooDB

    return this.http.post<WaldoDB>(this.waldosUrl, waldodb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        waldodb.Foo_Waldos_reverse = _Foo_Waldos_reverse
        this.log(`posted waldodb id=${waldodb.ID}`)
      }),
      catchError(this.handleError<WaldoDB>('postWaldo'))
    );
  }

  /** DELETE: delete the waldodb from the server */
  deleteWaldo(waldodb: WaldoDB | number): Observable<WaldoDB> {
    const id = typeof waldodb === 'number' ? waldodb : waldodb.ID;
    const url = `${this.waldosUrl}/${id}`;

    return this.http.delete<WaldoDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted waldodb id=${id}`)),
      catchError(this.handleError<WaldoDB>('deleteWaldo'))
    );
  }

  /** PUT: update the waldodb on the server */
  updateWaldo(waldodb: WaldoDB): Observable<WaldoDB> {
    const id = typeof waldodb === 'number' ? waldodb : waldodb.ID;
    const url = `${this.waldosUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _Foo_Waldos_reverse = waldodb.Foo_Waldos_reverse
    waldodb.Foo_Waldos_reverse = new FooDB

    return this.http.put<WaldoDB>(url, waldodb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        waldodb.Foo_Waldos_reverse = _Foo_Waldos_reverse
        this.log(`updated waldodb id=${waldodb.ID}`)
      }),
      catchError(this.handleError<WaldoDB>('updateWaldo'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
