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

import { MetaDB } from './meta-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class MetaService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  MetaServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private metasUrl: string

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
    this.metasUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/metas';
  }

  /** GET metas from the server */
  getMetas(): Observable<MetaDB[]> {
    return this.http.get<MetaDB[]>(this.metasUrl)
      .pipe(
        tap(_ => this.log('fetched metas')),
        catchError(this.handleError<MetaDB[]>('getMetas', []))
      );
  }

  /** GET meta by id. Will 404 if id not found */
  getMeta(id: number): Observable<MetaDB> {
    const url = `${this.metasUrl}/${id}`;
    return this.http.get<MetaDB>(url).pipe(
      tap(_ => this.log(`fetched meta id=${id}`)),
      catchError(this.handleError<MetaDB>(`getMeta id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new meta to the server */
  postMeta(metadb: MetaDB): Observable<MetaDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    metadb.MetaReferences = []

    return this.http.post<MetaDB>(this.metasUrl, metadb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`posted metadb id=${metadb.ID}`)
      }),
      catchError(this.handleError<MetaDB>('postMeta'))
    );
  }

  /** DELETE: delete the metadb from the server */
  deleteMeta(metadb: MetaDB | number): Observable<MetaDB> {
    const id = typeof metadb === 'number' ? metadb : metadb.ID;
    const url = `${this.metasUrl}/${id}`;

    return this.http.delete<MetaDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted metadb id=${id}`)),
      catchError(this.handleError<MetaDB>('deleteMeta'))
    );
  }

  /** PUT: update the metadb on the server */
  updateMeta(metadb: MetaDB): Observable<MetaDB> {
    const id = typeof metadb === 'number' ? metadb : metadb.ID;
    const url = `${this.metasUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    metadb.MetaReferences = []

    return this.http.put<MetaDB>(url, metadb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated metadb id=${metadb.ID}`)
      }),
      catchError(this.handleError<MetaDB>('updateMeta'))
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
