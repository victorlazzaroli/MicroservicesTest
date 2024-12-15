import { computed, Injectable, signal } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private _isAuthenticated = signal(true);

  public isAuthenticated = computed(() => this._isAuthenticated())

  public logout() {
    console.log("Logout")
    this._isAuthenticated.set(false)
  }

  public login() {
    this._isAuthenticated.set(true)
  }
}
