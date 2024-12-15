import { Component, inject } from '@angular/core';
import { MessageService } from '../../common/services/message.service';
import { MessageComponent } from '../../common/ui/message/message.component';

@Component({
  selector: 'app-feed',
  imports: [MessageComponent],
  templateUrl: './feed.component.html',
  styleUrl: './feed.component.scss'
})
export class FeedComponent {
  messageService = inject(MessageService)


}
