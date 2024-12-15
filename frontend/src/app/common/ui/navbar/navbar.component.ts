import { Component, inject } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { ActivatedRoute, RouterLink, RouterLinkActive } from '@angular/router';
import { AsyncPipe } from '@angular/common';
import { map } from 'rxjs';

@Component({
  selector: 'app-navbar',
  imports: [RouterLink, RouterLinkActive, AsyncPipe],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.scss'
})
export class NavbarComponent {
  authService = inject(AuthService)

  route = inject(ActivatedRoute)

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
