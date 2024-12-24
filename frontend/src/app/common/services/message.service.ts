import { Injectable, signal } from '@angular/core';
import { MessageI } from '../ui/message/model/message';

const mockMessages: MessageI[] = [
  {
    id: 1,
    title: "Welcome to the Forum",
    text: "This is your first post! Feel free to share your thoughts.",
    author: "Admin",
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

  constructor() { }

  messages = signal<MessageI[]>(mockMessages)
}
