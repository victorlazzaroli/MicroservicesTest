import { Component, computed, inject } from '@angular/core';
import { MessageService } from '../../common/services/message.service';
import { MessageComponent } from '../../common/ui/message/message.component';
import { MessageI } from '../../common/ui/message/model/message';

@Component({
  selector: 'app-home',
  imports: [MessageComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {
  messageService = inject(MessageService)

  myMessages = computed<MessageI[]>(() => {
     return this.messageService.messages()?.filter(message => message.author.toLowerCase() == "admin") || []
  } )
}
