import { Component, input, model } from '@angular/core';
import { MessageI } from './model/message';
import { DatePipe } from '@angular/common';

@Component({
  selector: 'app-message',
  imports: [DatePipe],
  templateUrl: './message.component.html',
  styleUrl: './message.component.scss'
})
export class MessageComponent {

  message = model.required<MessageI>()

  like() {
    console.log("Like:", this.message()?.id)

    this.message.update((el) => el && {...el, likes: el.likes + 1})
  }
  
  dislike() {
    console.log("Dislike: ", this.message()?.id)
    this.message.update((el) => el && {...el, dislikes: el.dislikes + 1})
  }
}
