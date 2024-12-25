import { Component, computed, inject } from '@angular/core';
import { MessageService } from '../../common/services/message.service';
import { MessageComponent } from '../../common/ui/message/message.component';
import { MessageI } from '../../common/ui/message/model/message';
import { NewMessageComponent } from "../../common/ui/new-message/new-message.component";
import { switchMap, tap } from 'rxjs';

@Component({
  selector: 'app-home',
  imports: [MessageComponent, NewMessageComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {
  
  messageService = inject(MessageService)

  constructor() {
    this.messageService.getMessages(1)
  }

  createMessage(message: string) {
    this.messageService.saveMessage({
      title: 'Nuovo messaggio',
      text: message,
      author: 'Victor Lazzaroli',
      author_id: 1,
      likes: 0,
      dislikes: 0
    }).pipe(
      tap(() => this.messageService.getMessages(1)))
      .subscribe()
  }
}
