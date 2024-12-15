import { Routes } from '@angular/router';
import { FeedComponent } from './pages/feed/feed.component';
import { HomeComponent } from './pages/home/home.component';
import { LoginComponent } from './pages/login/login.component';
import { NotFoundComponent } from './pages/not-found/not-found.component';

export const routes: Routes = [
    {
        path:'',
        pathMatch: 'full',
        redirectTo: 'feed'
    },
    {
        path: 'feed',
        title: "Twitter Clone",
        component: FeedComponent
    },
    {
        path: 'home',
        title: "Twitter Clone - My Home",
        component: HomeComponent
    },
    {
        path: 'login',
        title: "Twitter Clone - Login",
        component: LoginComponent
    },
    {
        path: '**',
        title: "Twitter Clone - 404",
        component: NotFoundComponent
    }
];
