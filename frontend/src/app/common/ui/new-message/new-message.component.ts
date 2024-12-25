import { Component, output } from '@angular/core';

@Component({
  selector: 'app-new-message',
  imports: [],
  templateUrl: './new-message.component.html',
  styleUrl: './new-message.component.scss'
})
export class NewMessageComponent {

  submit = output<string>()

}
