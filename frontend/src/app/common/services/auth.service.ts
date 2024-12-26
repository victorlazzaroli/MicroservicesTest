import { HttpClient } from '@angular/common/http';
import { computed, inject, Injectable, signal } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, EMPTY, Observable, retry, tap } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private readonly httpClient = inject(HttpClient)
  private readonly router = inject(Router)


  private _isAuthenticated = signal(true);
  public isAuthenticated = computed(() => this._isAuthenticated())

  public logout() {
    return this.httpClient.post("/api/auth/signout", {})
    .pipe(
      retry(2),
      catchError(() => {
        return EMPTY
      }),
      tap(() => {
        this._isAuthenticated.set(false)
        this.router.navigate(["login"])
      })
    )
  }

  public login(email: string, password: string): Observable<void> {
    return this.httpClient.post<void>("/api/auth/signin", {email, password})
    .pipe(
      catchError(() => {
        this._isAuthenticated.set(false)
        return EMPTY
      }),
      tap(() => {
        this._isAuthenticated.set(true)
      })
    )
  }
}
