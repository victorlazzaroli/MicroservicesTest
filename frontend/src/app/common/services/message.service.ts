import { Injectable, Signal, signal } from '@angular/core';
import { MessageI } from '../ui/message/model/message';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

const mockMessages: MessageI[] = [
  {
    id: 1,
    title: "Welcome to the Forum",
    text: "This is your first post! Feel free to share your thoughts.",
    author: "Admin",
    author_id: 1,
    createdAt: "2024-12-01T10:00:00Z",
    updatedAt: "2024-12-01T10:00:00Z",
    likes: 15,
    dislikes: 2,
  },
  {
    id: 2,
    title: "How to Learn TypeScript",
    text: "Here are some great resources to get started with TypeScript.",
    author: "JohnDoe",
    author_id: 2,
    createdAt: "2024-12-02T14:20:00Z",
    updatedAt: "2024-12-03T08:45:00Z",
    likes: 50,
    dislikes: 5,
  },
  {
    id: 3,
    title: "Favorite Programming Language",
    text: "What's your favorite programming language and why?",
    author: "JaneSmith",
    author_id: 3,
    createdAt: "2024-12-05T09:30:00Z",
    updatedAt: "2024-12-06T16:10:00Z",
    likes: 23,
    dislikes: 8,
  },
  {
    id: 4,
    title: "Weekly Update",
    text: "Here's what's new this week on our platform.",
    author: "Moderator",
    author_id: 4,
    createdAt: "2024-12-10T08:00:00Z",
    updatedAt: "2024-12-10T08:00:00Z",
    likes: 30,
    dislikes: 1,
  },
  {
    id: 5,
    title: "Question About Interfaces",
    text: "Can someone explain the difference between interfaces and types?",
    author: "Coder123",
    author_id: 5,
    createdAt: "2024-12-11T13:45:00Z",
    updatedAt: "2024-12-12T10:30:00Z",
    likes: 40,
    dislikes: 3,
  },
  {
    id: 6,
    title: "Book Recommendations",
    text: "What are some must-read books for software developers?",
    author: "BookLover",
    author_id: 6,
    createdAt: "2024-12-09T17:50:00Z",
    updatedAt: "2024-12-09T17:50:00Z",
    likes: 18,
    dislikes: 4,
  },
  {
    id: 7,
    title: "Bug Report",
    text: "I encountered an issue while using the app. Here's how to reproduce it.",
    author: "BugHunter",
    author_id: 7,
    createdAt: "2024-12-13T15:10:00Z",
    updatedAt: "2024-12-13T16:00:00Z",
    likes: 5,
    dislikes: 0,
  },
  {
    id: 8,
    title: "Tips for Debugging",
    text: "Debugging can be tricky! Here are some tips to make it easier.",
    author: "DebugPro",
    author_id: 8,
    createdAt: "2024-12-07T11:15:00Z",
    updatedAt: "2024-12-07T11:15:00Z",
    likes: 27,
    dislikes: 2,
  },
  {
    id: 9,
    title: "New Feature Suggestions",
    text: "What features would you like to see in the next update?",
    author: "FeatureFan",
    author_id: 9,
    createdAt: "2024-12-04T19:00:00Z",
    updatedAt: "2024-12-04T19:00:00Z",
    likes: 12,
    dislikes: 6,
  },
  {
    id: 10,
    title: "Community Guidelines Reminder",
    text: "Please review the community guidelines to ensure a respectful environment.",
    author: "Admin",
    author_id: 1,
    createdAt: "2024-12-14T09:30:00Z",
    updatedAt: "2024-12-14T09:30:00Z",
    likes: 25,
    dislikes: 1,
  },
];


@Injectable({
  providedIn: 'root'
})
export class MessageService {

  constructor(private readonly httpClient: HttpClient) { }

  messages = signal<MessageI[]>([]);

  getMessages(authorId?: number) {
    if (authorId) {
      this.httpClient.get<MessageI[]>(`api/messages/author/${authorId}`).subscribe(
        (response) => this.messages.set(response)
      )
    } else {
      this.httpClient.get<MessageI[]>(`api/messages/`).subscribe(
        (response) => this.messages.set(response)
      )
    }
  }

  saveMessage(message: Partial<MessageI>): Observable<void> {
    return this.httpClient.post<void>(`api/messages/`, JSON.stringify(message))
  }

}
