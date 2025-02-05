import { Component, inject } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { ActivatedRoute, RouterLink, RouterLinkActive } from '@angular/router';
import { map } from 'rxjs';

@Component({
  selector: 'app-navbar',
  imports: [RouterLink, RouterLinkActive],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.scss'
})
export class NavbarComponent {
  authService = inject(AuthService)
  
  route = inject(ActivatedRoute)
  
  logout() {
    this.authService.logout().subscribe()
  }
  active = this.route.url
  .pipe(
    map(urlSegments => {
      const segments = urlSegments.length
      console.log(urlSegments)
      if (segments >= 1) {
        return urlSegments[segments - 1]?.toString() || '/'
      } else {
        return '/'
      }
    }
  ));



}
